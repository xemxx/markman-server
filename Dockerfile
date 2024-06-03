FROM alpine:latest
RUN mkdir /app/runtime/log -p
WORKDIR /app
COPY .build/markman-server markman
COPY app.yaml.default app.yaml

ENTRYPOINT [ "./markman" ]