FROM golang:1.15.6

RUN apt-get update -q && apt-get install -yq ca-certificates

ENV \
  GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /go/src/github.com/egeneralov/solana_exporter
ADD go.mod go.sum /go/src/github.com/egeneralov/solana_exporter/
RUN go mod download -x
ADD . .
RUN go build -v -installsuffix cgo -ldflags="-w -s" -o /go/bin/solana_exporter github.com/egeneralov/solana_exporter/cmd/solana_exporter


FROM debian:buster

RUN apt-get update -q && apt-get install -yq ca-certificates
USER nobody
ENV PATH='/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin'
CMD /go/bin/solana_exporter

COPY --from=0 /go/bin /go/bin
COPY --from=0 /go/src/github.com/egeneralov/solana_exporter/monitor.sh /go/bin/monitor.sh
