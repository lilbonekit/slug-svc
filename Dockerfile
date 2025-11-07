FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/lilbonekit/slug-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/ /go/src/github.com/lilbonekit/slug-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/ /usr/local/bin/
RUN apk add --no-cache ca-certificates

ENTRYPOINT [""]
