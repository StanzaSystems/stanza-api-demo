package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc/metadata"

	pb "github.com/StanzaSystems/stanza-api-demo/hubv1"
)

type requests struct {
	tier          string
	customerID    string
	duration      time.Duration
	rate          int
	priorityBoost int
	weight        int
}

func requestQuota(client pb.QuotaServiceClient, reqs requests) {
	// run for 2 secs longer and then disregard last 2 seconds, which can be less accurate as results trail off
	count := reqs.rate*int(reqs.duration.Seconds()) + 2

	var wg sync.WaitGroup
	limiter := rate.NewLimiter(rate.Limit(reqs.rate), int(reqs.rate))

	for i := 0; i < int(count); i++ {
		for !limiter.Allow() {
			time.Sleep(time.Millisecond * 1)
		}

		req := pb.GetTokenRequest{
			S: &pb.DecoratorFeatureSelector{
				DecoratorName: decoratorName,
				Environment:   env,
				Tags: []*pb.Tag{
					{Key: "tier", Value: reqs.tier},
					{Key: "customer_id", Value: reqs.customerID},
				},
			},
			PriorityBoost: &reqs.priorityBoost,
		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			got := doReq(client, &req)
			// todo do meters
			fmt.Printf("%v", got)
		}()
	}

	wg.Wait()
}

func doReq(client pb.QuotaServiceClient, request *pb.GetTokenRequest) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx, "X-Stanza-Key", apikey)

	r, err := client.GetToken(ctx, request)
	if verbose {
		if err != nil {
			log.Printf("Error %v\n", err)
		} else {
			log.Printf("Response %+v\n", r)
		}
	}

	if err != nil {
		return false
	}
	return r.Granted
}
