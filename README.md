# stanza-api-demo
todo describe

 * stanza-api-demo should export metrics and provide a rest interface to run sets of requests
 * cli should run in own container you can exec into to hit the api demo
 * grafana graphs of requests permitted/denied, with labels (tags, priority)
 * graph of total weights also
 * latencies?
 * way to do weight adjustments via cli
 * demo should do proper deadlines etc
 * read through for todos etc

## Running the Stanza Demo

Clone this repository to a machine that has `Docker` and `docker-compose` installed.

In the root of the repo run `docker-compose up -d`

Find the Grafana container at [http://localhost:3000](http://localhost:3000). Here you can see graphs showing the Stanza API's behaviour - how many requests are granted, denied, errors, and latency.

You can run sequences of commands against the Stanza API using the CLI provided.
`docker exec stanza-api-demo-cli-1  /stanza-api-cli`
 
 TODO arguments

The Stanza runner status page is at http://localhost:9278/status. Here you can see a record of previous and currently running sets of requests.


## Use your own custom Stanza API Key and Config
Currently the demo is set up to use a pre-loaded API key and decorator configuration, but
soon you will be able to set up your own API key and experiment with your own decorator confgurations.

## Performance and Load

Stanza is currently in eval/beta; our API demo is hosted only in one region (us-east-2). 
When fully launched we will run in several regions to reduce latency, but for now, you will experience some latency if you are not located near us-east-2. 
If located at a significant distance from us-east-2 you may see some queueing and occasional timeouts in the client if you send a significant number of requests (multiple hundreds per second). This is mainly due to queuing at the client.