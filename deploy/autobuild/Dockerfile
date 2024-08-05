# 使用最新版本的 golang 官方镜像作为构建阶段的基础镜像
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o markman .

FROM alpine:latest
WORKDIR /workspace
RUN mkdir /workspace/runtime/log -p
COPY --from=builder /app/markman /workspace/markman

ENTRYPOINT [ "./markman" ]