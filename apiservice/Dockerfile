FROM alpine:3.14

RUN apk update
RUN rm -rf /var/lib/apt/lists/*

WORKDIR /go

COPY api .

ENTRYPOINT  /go/api