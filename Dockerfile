FROM golang:1.15

ENV CGO_ENABLED 0

# Allow Go to retreive the dependencies for the build step
RUN apk add --no-cache git

# Get Delve from a GOPATH not from a Go Modules project
WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /t-blog/
ADD . /t-blog/

RUN go build -o /t-blog/srv .

EXPOSE 8888 40000

CMD ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/t-blog/srv"]