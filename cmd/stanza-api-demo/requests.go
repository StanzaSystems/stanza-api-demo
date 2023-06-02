package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/StanzaSystems/stanza-api-demo/gen/go/stanza/hub/v1"
	"github.com/gin-gonic/gin"
)

type requests struct {
	Tags          string `json:"tags"` //format: foo=bar,baz=quux etc
	Duration      string `json:"duration"`
	duration_time time.Duration
	Rate          int     `json:"rate"`
	PriorityBoost int32   `json:"priority_boost"`
	Weight        float32 `json:"weight"`
	parsedTags    map[string]string
	started       *time.Time
	ended         *time.Time
	APIkey        string `json:"apikey"`
	Environment   string `json:"environment"`
	Decorator     string `json:"decorator"`
}

type RequestRunner struct {
	client  pb.QuotaServiceClient
	history []*requests
	m       *meters
}

func MakeRequestRunner() *RequestRunner {
	conn, err := grpc.Dial(hub, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewQuotaServiceClient(conn)
	return &RequestRunner{
		client:  client,
		history: make([]*requests, 0),
		m:       MakeMeters(),
	}
}

// postRequests starts a set of requests from JSON received in the request body.
func (r *RequestRunner) postRequest(c *gin.Context) {
	fmt.Printf("Post request\n")

	var reqs requests
	reqs.parsedTags = make(map[string]string)

	if err := c.BindJSON(&reqs); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Post request - bad request %+v\n", err)
		return
	}

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
			fmt.Printf("Post request - can't parse tags\n")
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		reqs.parsedTags[st[0]] = st[1]
	}

	dur, err := time.ParseDuration(reqs.Duration)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Post request - can't parse duration %+v\n", err)
		return
	}
	reqs.duration_time = dur

	go r.requestQuota(&reqs)
	c.Writer.WriteHeader(http.StatusOK)
}

func (r *RequestRunner) status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.tmpl", gin.H{"Time": fmt.Sprintf("%v", time.Now())})
}

func (r *RequestRunner) requestQuota(reqs *requests) {
	fmt.Printf("Running requests %+v\n", reqs)

	count := reqs.Rate * int(reqs.duration_time.Seconds())
	r.history = append(r.history, &requests{})
	start := time.Now()
	reqs.started = &start

	var wg sync.WaitGroup
	limiter := rate.NewLimiter(rate.Limit(reqs.Rate), int(reqs.Rate))

	for i := 0; i < int(count); i++ {
		for !limiter.Allow() {
			time.Sleep(time.Millisecond * 1)
		}

		tags := []*pb.Tag{}
		for name, val := range reqs.parsedTags {
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
	reqs.ended = &end

	wg.Wait()

	fmt.Printf("done ")
}

func (r *RequestRunner) doReq(client pb.QuotaServiceClient, request *pb.GetTokenRequest, apikeystr string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	labels := make(map[string]string)
	labels["priorityBoost"] = fmt.Sprintf("%d", *request.PriorityBoost)
	labels["decorator"] = *&request.S.DecoratorName
	labels["environment"] = *&request.S.Environment
	labels["apikey"] = apikeystr

	for _, l := range request.S.GetTags() {
		labels[l.Key] = l.Value
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "X-Stanza-Key", apikeystr)
	resp, err := client.GetToken(ctx, request)

	fmt.Printf("Request, response, error: %+v %+v %+v\n", *request, resp, err)
	if verbose {
		if err != nil {
			log.Printf("Error %v\n", err)
		} else {
			log.Printf("Response %+v\n", r)
		}
	}

	e, ok := status.FromError(err)
	if !ok {
		r.m.GetQuotaErrorCounter(labels).With(labels).Inc()
		return
	}

	if err != nil && e.Code() != codes.ResourceExhausted {
		r.m.GetQuotaErrorCounter(labels).With(labels).Inc()
		return
	}

	if e.Code() == codes.ResourceExhausted {
		r.m.GetQuotaNotGrantedCounter(labels).With(labels).Inc()
		return
	}

	// no error, granted
	r.m.GetQuotaGrantedCounter(labels).With(labels).Inc()

}
