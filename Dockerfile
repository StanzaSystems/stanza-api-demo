# Build Stage
FROM golang:1.20 AS build

LABEL app="stanza-api-demo"
LABEL REPO="https://github.com/StanzaSystems/stanza-api-demo"

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o /go/src/app/bin/stanza-api-demo ./cmd/stanza-api-demo

# Final Stage
FROM gcr.io/distroless/base-debian11

# Metrics and HTTP port
EXPOSE 9277/TCP

WORKDIR /
COPY --from=build /go/src/app/bin/stanza-api-demo /stanza-api-demo

CMD ["/stanza-api-demo", "-verbose"]