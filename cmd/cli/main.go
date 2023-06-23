package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	demo "github.com/StanzaSystems/stanza-api-demo/demo"
)

var (
	durationStr   string
	rate          int
	tags          string
	priorityBoost int
	weight        float64
)

func main() {
	flag.StringVar(&durationStr, "duration", "60s", "How long to keep running queries against Stanza.")
	flag.IntVar(&rate, "rate", 100, "How many requests per second to send to Stanza.")
	flag.StringVar(&tags, "tags", "", "Tags to apply to requests sent to Stanza. Format is a comma-separated list of tagName=value. Example: tier=paid,customer_id=paid-customer-2")
	flag.IntVar(&priorityBoost, "priority_boost", 0, "How much to boost the priority of requests sent to Stanza (5 is the highest possible boost, -5 lowest.")
	flag.Float64Var(&weight, "weight", 1, "Weight of requests to send to Stanza.")

	flag.Parse()

	reqs := demo.Requests{
		Duration:      durationStr,
		Rate:          rate,
		Tags:          tags,
		Weight:        float32(weight),
		PriorityBoost: int32(priorityBoost),
	}

	data, _ := json.Marshal(reqs)

	r, err := http.NewRequest("POST", "http://demo:9278/run", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		fmt.Print("Requests running - see Grafana http://localhost:3000 for dashboard\n")
	} else {
		fmt.Printf("Something went wrong - see docker logs.\n")
	}
}
