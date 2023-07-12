FROM alpine:latest
RUN mkdir /workspace/runtime/log -p
WORKDIR /workspace
COPY .build/markman-server markman
COPY app.yaml.default app.yaml

ENTRYPOINT [ "./markman" ]