# create 
FROM golang:1.13
WORKDIR /go/src/github.com/first-golang-k8s-dockerApp

COPY cli ./cli

RUN CGO_ENABLED=0 GOOS=linux go install ./cli/server

# execute

FROM alpine:latest
WORKDIR /app/

COPY --from=0 /go/bin/server ./binary

CMD ./binary
