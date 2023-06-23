package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	gokitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	demo "github.com/StanzaSystems/stanza-api-demo/demo"

	"github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp"
	configpb "github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp/grpc_gcp"
	pb "github.com/StanzaSystems/stanza-api-demo/gen/go/stanza/hub/v1"
	"github.com/gin-gonic/gin"
)

type RequestRunner struct {
	client  pb.QuotaServiceClient
	history []*demo.Requests
	m       *meters
}

func MakeRequestRunner() *RequestRunner {
	config := &tls.Config{} // use default system CA
	creds := credentials.NewTLS(config)
	if hub_insecure {
		creds = insecure.NewCredentials()
	}

	apiConfig := &configpb.ApiConfig{
		ChannelPool: &configpb.ChannelPoolConfig{
			MaxSize:                          10,
			MaxConcurrentStreamsLowWatermark: 25,
		},
	}
	c, err := protojson.Marshal(apiConfig)
	if err != nil {
		log.Fatalf("cannot json encode config: %v", err)
	}
	jsonCfg := string(c)

	conn, err := grpc.Dial(hub,
		grpc.WithTransportCredentials(creds),
		grpc.WithDefaultServiceConfig(
			fmt.Sprintf(
				`{"loadBalancingConfig": [{"%s":%s}]}`,
				grpcgcp.Name,
				jsonCfg,
			),
		),
		grpc.WithChainUnaryInterceptor(
			// Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			grpcgcp.GCPUnaryClientInterceptor,
			grpc_prometheus.UnaryClientInterceptor,
			//logging.UnaryClientInterceptor(interceptorLogger(rpcLogger), logging.WithFieldsFromContext(logTraceID)),
		),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewQuotaServiceClient(conn)
	return &RequestRunner{
		client:  client,
		history: make([]*demo.Requests, 0),
		m:       MakeMeters(),
	}
}

// postRequests starts a set of requests from JSON received in the request body.
func (r *RequestRunner) postRequest(c *gin.Context) {
	reqs := demo.Requests{}

	if err := c.BindJSON(&reqs); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	reqs.ParsedTags = make(map[string]string)

	if reqs.APIkey == "" {
		reqs.APIkey = apikey_default
	}
	if reqs.Decorator == "" {
		reqs.Decorator = decoratorName
	}
	if reqs.Environment == "" {
		reqs.Environment = env
	}

	splitTags := strings.Split(reqs.Tags, ",")
	for _, tagKV := range splitTags {
		if len(tagKV) == 0 {
			continue
		}
		st := strings.Split(tagKV, "=")
		if len(st) != 2 {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		reqs.ParsedTags[st[0]] = st[1]
	}

	if reqs.Rate > 2000 {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	dur, err := time.ParseDuration(reqs.Duration)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	reqs.Duration_time = dur

	go r.requestQuota(reqs)
	c.Writer.WriteHeader(http.StatusOK)
}

func (r *RequestRunner) status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.tmpl", gin.H{"Time": fmt.Sprintf("%v", time.Now())})
}

func (r *RequestRunner) requestQuota(reqs demo.Requests) {
	count := reqs.Rate * int(reqs.Duration_time.Seconds())
	r.history = append(r.history, &demo.Requests{})
	start := time.Now()
	reqs.Started = &start

	var wg sync.WaitGroup
	limiter := rate.NewLimiter(rate.Limit(reqs.Rate), int(reqs.Rate))

	for i := 0; i < int(count); i++ {
		for !limiter.Allow() {
			time.Sleep(time.Millisecond * 1)
		}

		tags := []*pb.Tag{}
		for name, val := range reqs.ParsedTags {
			tags = append(
				tags,
				&pb.Tag{Key: name, Value: val},
			)
		}

		req := pb.GetTokenRequest{
			S: &pb.DecoratorFeatureSelector{
				DecoratorName: reqs.Decorator,
				Environment:   reqs.Environment,
				Tags:          tags,
			},
			Weight:        &reqs.Weight,
			PriorityBoost: &reqs.PriorityBoost,
		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			go r.doReq(r.client, &req, reqs.APIkey)
		}()
	}

	end := time.Now()
	reqs.Ended = &end

	wg.Wait()
}

func (r *RequestRunner) doReq(client pb.QuotaServiceClient, request *pb.GetTokenRequest, apikeystr string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	labels := make(map[string]string)
	labels["priorityBoost"] = fmt.Sprintf("%d", *request.PriorityBoost)
	labels["environment"] = request.S.Environment
	labels["tags"] = tagsToStr(request.S.GetTags())

	ctx = metadata.AppendToOutgoingContext(ctx, "X-Stanza-Key", apikeystr)

	start := time.Now()
	resp, err := client.GetToken(ctx, request)
	duration := time.Since(start)

	r.m.latency.Observe(float64(duration.Seconds()))
	if duration.Milliseconds() > 300 {
		if err != nil {
			fmt.Printf("ERROR %+v\n", err)
		}
	}

	if verbose {
		if err != nil {
			log.Printf("Error %v, duration %+v\n", err, duration)
		} else {
			log.Printf("Response %+v, duration %+v\n", resp, duration)
		}
	}

	e, ok := status.FromError(err)
	if !ok {
		r.m.GetQuotaErrorCounter().With(labels).Inc()
		return
	}

	if err != nil && e.Code() != codes.ResourceExhausted {
		r.m.GetQuotaErrorCounter().With(labels).Inc()
		return
	}

	if resp.Granted {
		r.m.GetQuotaGrantedCounter().With(labels).Inc()
	} else {
		r.m.GetQuotaNotGrantedCounter().With(labels).Inc()
	}
}

func tagsToStr(tags []*pb.Tag) string {
	result := ""
	m := make(map[string]string)

	keys := make([]string, 0)
	for _, t := range tags {
		keys = append(keys, t.Key)
		m[t.Key] = t.Value
	}
	sort.Strings(keys)
	for _, k := range keys {
		result = result + k + "=" + m[k] + ","
	}

	if len(result) > 0 {
		return result[0 : len(result)-1]
	} else {
		return result
	}
}

func interceptorLogger(l gokitlog.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = level.Debug(l).Log(largs...)
		case logging.LevelInfo:
			_ = level.Info(l).Log(largs...)
		case logging.LevelWarn:
			_ = level.Warn(l).Log(largs...)
		case logging.LevelError:
			_ = level.Error(l).Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
