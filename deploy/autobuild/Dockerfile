# 使用最新版本的 golang 官方镜像作为构建阶段的基础镜像
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN ls -lh
RUN go mod download
RUN go get .
RUN go build -o markman-server .

FROM alpine:latest
RUN mkdir /workspace/runtime/log -p
WORKDIR /workspace
COPY --from=builder /app/markman-server markman

ENTRYPOINT [ "/workspace/markman" ]