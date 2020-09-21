FROM golang:1.15
WORKDIR /go/src/markman-server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o markman-server -i .

FROM alpine:latest

RUN mkdir /workspace/runtime/log -p 

WORKDIR /workspace 

COPY --from=0 /go/src/markman-server/markman-server markman
COPY app.json.bac app.json
