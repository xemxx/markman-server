FROM alpine:latest
RUN mkdir /workspace/runtime/log -p
WORKDIR /workspace
COPY .build/markman-server markman
COPY docker/app.json app.json

ENTRYPOINT [ "./markman" ]