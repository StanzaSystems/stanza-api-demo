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
 * Grafana, for displaying graphs of what the demo is observing

### Ports

 The demo will attempt to map Grafana's endpoint to port 3000 on your machine. If port 3000 is in use, 
 you can edit the `docker-compose.yaml` file at the root of this repo and change that to another port.
 For example, to use port `3001` make the following change:

```
  grafana:
    build: ./grafana
    ports:
      - '3001:3000'
```

You will then need to access Grafana on whichever port you have specified, rather than 3000 as used in the examples below.

### Grafana and CLI 

Find the Grafana container at [http://localhost:3000](http://localhost:3000). Here you can see graphs showing the Stanza API's behaviour - how many requests are granted, denied, errors, and latency. Initially there will be no data there (until we run some requests).

You can run sequences of commands against the Stanza API using the CLI provided (examples below).
`docker exec stanza-api-demo-cli-1  /stanza-api-cli`

## Decorators and Rate Limits in Action

One of Stanza's core concepts is the Decorator. Decorators are used to guard a resource that can become overloaded. 
You can configure Stanza Decorators with a rate and a burst.
The rate is the number of requests that the Decorator can serve steady state. Burst, if higher than the rate, allows temporary periods of higher usage,
but the average number of requests cannot exceed the steady state rate.  

Observe this by running the Stanza API demo as follows:
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=30s --rate=150 --tags=tier=paid,customer_id=paid-customer-1
```

Our demo quota sets a rate limit of 100 requests per second for each customer in the `paid` tier. We are requesting above that rate (150 qps).
In [Grafana](http://localhost:3000/d/W23Z3R_Vk/stanza-api-demo?orgId=1) you will see the rate of granted requests
rise to around 100 per second, while the rate not granted rises to 50 per second. This sequence of requests will run for 30 seconds. 
It will take a few seconds for metrics to be scraped and displayed. 

## Request Prioritization

Stanza performs dynamic prioritisation of requests. Stanza has 11 request priority levels. 0 is the highest priority and 10 is the lowest.
By default, requests have priority `5`. We can boost or reduce priority of each request. 

You can see priority boosting in action by running two overlapping sets of requests, as follows:
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=paid,customer_id=paid-customer-1
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=30s --rate=100 --tags=tier=paid,customer_id=paid-customer-1 --priority_boost=5
```

These should show up as two separate graph lines in [Grafana](http://localhost:3000/d/W23Z3R_Vk/stanza-api-demo?orgId=1).
You'll also see a separate line for the sum of all granted requests.
You will see that most of the boosted requests are granted.
Because available quota is 100 requests per second, and this is consumed by the boosted requests, the non-boosted requests are not granted.
After 30 seconds, the boosted requests stop and the default-priority requests will be granted - you should see the lines cross on the graph.

### Child Quotas

In many applications, we want to be able to make some guarantees about isolation and fairness between different workloads.
We can't do this through request prioritisation alone.

In this demo, we have set up a Decorator with the following configuration:

```
{
  "enabled": true,
  "tagConfig": {
    "tagName": "tier"
  },
  "childQuotaConfigs": {
    "tier=enterprise": {
      "rateLimitConfig": {
        "rate": 1500,
        "burst": 1500
      },
      "tagConfig": {
        "tagName": "customer_id",
        "tagOptional": true
      },
      "childQuotaConfigs": {
        "customer_id=a-large-customer": {
          "rateLimitConfig": {
            "rate": 50,
            "burst": 80
          }
        },
        "customer_id=a-larger-customer": {
          "rateLimitConfig": {
            "rate": 150,
            "burst": 200
          }
        },
        "customer_id=a-small-customer": {
          "rateLimitConfig": {
            "rate": 20,
            "burst": 50
          }
        }
      },
      "childDefaultConfigs": {
        "customer_id": {
          "rate": 20,
          "burst": 40,
          "bestEffortBurst": true
        }
      }
    },
    "tier=free": {
      "rateLimitConfig": {
        "rate": 10,
        "burst": 10
      }
    },
    "tier=paid": {
      "rateLimitConfig": {
        "rate": 100,
        "burst": 100
      },
      "tagConfig": {
        "tagName": "customer_id"
      },
      "childDefaultConfigs": {
        "customer_id": {
          "rate": 10,
          "burst": 20,
          "bestEffortBurst": true
        }
      }
    }
  }
}
```


There are three customer tiers here: free, paid, and enterprise. They are specified as tags (labels) which are flexible - you can define
any set of tags that works for your application.

### Free Tier and Weights
The free tier gets a total of 10 qps, shared between all customers in that tier. In fact, we don't use `customer_id` in this tier to achieve fairness (just to demonstrate the flexibility of what is possible).

Run this sequence of requests: 
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=30s --rate=100 --tags=tier=free
```
You'll see that the requests granted tops out at 10 qps.

Requests can have varying weights. Try
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=30s --rate=100 --tags=tier=free --weight=5
```

This will grant only 2 qps because each request has weight 5.

```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=30s --rate=100 --tags=tier=free --weight=0.5
```

This will grant 10 qps because each request has weight 0.5.
When using weights with Stanza your application can estimate upfront and then update when the request has completed and the full cost is known.


### Paid Tier and Best Effort Burst

The paid tier gets a total of 100 qps, and each customer within that tier gets a steady-state rate of 10qps and can occasionally burst to 20qps (if they average 10qps or lower).
We don't specify limits for each customer - we only specify one default configuration for the tier, so every customer gets the same rate limit.

If the entire paid tier is oversubscribed, customers may not get 10 qps each, but Stanza will allocate as fairly as possible.
Customers in the paid tier are allowed to `bestEffortBurst`. This means that they will get more than 10 qps allocated, as long as there is available capacity at the paid tier level.

Run these commands: 
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=free
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=paid,customer_id=paid-customer-1
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=paid,customer_id=paid-customer-2
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=paid,customer_id=paid-customer-3
```

This runs requests against both the free tier and the paid tier. You should see a 110 qps being granted - 10 for the free tier and 100 for the paid tier, split between the 3 customers.

### Enterprise Tier and Varying Per-customer limits

The enterprise tier as a whole has a limit of 1500 qps.
Several customers have specifically defined rate limits, and there is a default configuration which is used for any customer_ids that do not have a specified limit.
The specifically-defined customers do not have `bestEffortBurst` enabled (just for demonstration purposes), so they must stick to their assigned rate limits.

In this scenario you will see that the enterprise customer `a-larger-customer` is strictly limited to 150 qps while
`default-ent-customer` is allowed to burst to 200 qps because the enterprise tier has spare capacity.

Run these commands: 
```
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=free
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=100 --tags=tier=paid,customer_id=paid-customer-1
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=200 --tags=tier=enterprise,customer_id=a-larger-customer
docker exec stanza-api-demo-cli-1  /stanza-api-cli --duration=60s --rate=200 --tags=tier=enterprise,customer_id=default-ent-customer
```

### Try your own scenarios

You can use the `docker exec stanza-api-demo-cli-1  /stanza-api-cli` tool to run any set of requests you choose against the Stanza demo. 

## Using your own custom Stanza API Key and Config
Currently the demo is set up to use a pre-loaded API key and decorator configuration, but
soon you will be able to set up your own API key and experiment with your own decorator configurations.

## Performance and Load
Stanza is currently in eval/alpha and our API demo is hosted only in one region (us-east-2). 
When fully launched we will run in several regions globally to reduce latency, but for now, you will experience some latency if you are not located near us-east-2. 
If located at a significant distance from us-east-2 you may see some queueing and occasional timeouts in the client if you send (multiple hundreds per second). This is mainly due to queuing at the client. This won't be an issue when we're out of eval mode.
