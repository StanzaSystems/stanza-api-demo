# Stanza API Demo

[Stanza](https://www.stanza.systems/) is a service that helps you protect your user experience during times when your services are heavily loaded.

We have some frontend capabilities that can let you gracefully degrade your site in real-time in response to overload - [contact us](https://www.stanza.systems/contact) if you'd like a demo of these. 
This repo provides a demo of some of the capabilities of Stanza's APIs that are designed to be integrated in your backend code. 

## Running the Stanza Demo

Clone this repository to a machine that has `docker` and `docker-compose` installed.

In the root of the repo run `docker build` and then `docker-compose up -d`.

This will run several containers, including:
 * A CLI
 * A server which runs requests against Stanza's demo service and exports metrics
 * Grafana, for displaying graphs of 

Find the Grafana container at [http://localhost:3000](http://localhost:3000). Here you can see graphs showing the Stanza API's behaviour - how many requests are granted, denied, errors, and latency. Initially there will be no data there.

You can run sequences of commands against the Stanza API using the CLI provided.
`docker exec stanza-api-demo-cli-1  /stanza-api-cli`

## Decorators and Rate Limits in Action

One of Stanza's core concepts is the Decorator. Decorators are used to guard a resource that can become overloaded. 
You can configure Stanza Decorators with a rate and a burst.
The rate is the number of requests that the Decorator can serve steady state. Burst, if higher than the rate, allows temporary periods of higher usage,
but the average number of requests cannot exceed the steady state rate.  

Observe this by running the Stanza API demo with the following parameters:
TODO fix
  * "duration": "30s",
  * "rate": 150,
  * "tags": "tier=paid,customer_id=paid-customer-1"

Our demo quota sets a rate limit of 100 requests per second for each customer in the `paid` tier.
In [Grafana](http://localhost:3000/d/W23Z3R_Vk/stanza-api-demo?orgId=1) you will see the rate of granted requests
rise to around 100 per second, while the rate not granted rises to 50 per second. This sequence of requests will run for 30 seconds. 

## Request Prioritization

Stanza performs dynamic prioritisation of requests. Stanza has 11 request priority levels. 0 is the highest priority and 10 is the lowest.
By default, requests have priority `5`. We can boost or reduce priority of each request. 

You can see priority boosting in action by running two overlapping sets of requests.
todo fix to cli
First set of requests:
  * "duration": "60s",
  * "rate": 100,
  * "tags": "tier=paid,customer_id=paid-customer-1"

Second set:s
  * "duration": "60s",
  * "rate": 100,
  *  "tags": "tier=paid,customer_id=paid-customer-1",
  *  "priority_boost": 5

These should show up as two separate graph lines in [Grafana](http://localhost:3000/d/W23Z3R_Vk/stanza-api-demo?orgId=1).
You will see that most of the boosted requests are granted, after a very short initial period where Stanza adjusts to the new request pattern. 
Because available quota is 100 requests per second, and this is consumed by the boosted requests, the non-boosted requests are not granted.

### Child Quotas

In many applications, we want to be able to make some guarantees about isolation and fairness between different workloads.
We can't do this through request prioritisation alone.

In this demo, we have set up a Decorator with the following configuration:



config := ipb.QuotaConfig{
		Enabled: proto.Bool(true),
		TagConfig: &ipb.TagConfig{
			TagName:     "tier",
			TagOptional: false,
		},
		ChildQuotaConfigs: map[string]*ipb.QuotaConfig{
			"tier=free": {
				RateLimitConfig: &ipb.RateLimitConfig{
					Rate:  proto.Uint32(10),
					Burst: proto.Uint32(10),
				},
			},
			"tier=paid": {
				RateLimitConfig: &ipb.RateLimitConfig{
					Rate:  proto.Uint32(100),
					Burst: proto.Uint32(100),
				},
				TagConfig: &ipb.TagConfig{TagName: "customer_id"},
				ChildDefaultConfigs: map[string]*ipb.RateLimitConfig{
					"customer_id": {
						Rate:            proto.Uint32(10),
						Burst:           proto.Uint32(20),
						BestEffortBurst: true,
					},
				},
			},
			"tier=enterprise": {
				RateLimitConfig: &ipb.RateLimitConfig{
					Rate:  proto.Uint32(1500),
					Burst: proto.Uint32(1500),
				},
				ChildQuotaConfigs: map[string]*ipb.QuotaConfig{
					"customer_id=a-small-customer":  {RateLimitConfig: &ipb.RateLimitConfig{Rate: proto.Uint32(20), Burst: proto.Uint32(50)}},
					"customer_id=a-large-customer":  {RateLimitConfig: &ipb.RateLimitConfig{Rate: proto.Uint32(50), Burst: proto.Uint32(80)}},
					"customer_id=a-larger-customer": {RateLimitConfig: &ipb.RateLimitConfig{Rate: proto.Uint32(150), Burst: proto.Uint32(200)}},
				},
				TagConfig: &ipb.TagConfig{TagName: "customer_id", TagOptional: true},
				ChildDefaultConfigs: map[string]*ipb.RateLimitConfig{
					"customer_id": {
						Rate:            proto.Uint32(20),
						Burst:           proto.Uint32(40),
						BestEffortBurst: true,
					},
				},
			},
		},
	}

 
 TODO arguments

The Stanza runner status page is at http://localhost:9278/status. Here you can see a record of previous and currently running sets of requests.

## Using your own custom Stanza API Key and Config
Currently the demo is set up to use a pre-loaded API key and decorator configuration, but
soon you will be able to set up your own API key and experiment with your own decorator confgurations.

## Performance and Load
Stanza is currently in eval/alpha and our API demo is hosted only in one region (us-east-2). 
When fully launched we will run in several regions globally to reduce latency, but for now, you will experience some latency if you are not located near us-east-2. 
If located at a significant distance from us-east-2 you may see some queueing and occasional timeouts in the client if you send (multiple hundreds per second). This is mainly due to queuing at the client. This won't be an issue when we're out of eval mode.