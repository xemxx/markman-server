FROM alpine:latest

RUN mkdir /workspace/runtime/log -p 

WORKDIR /workspace 

ADD .build/markman-server markman
ADD app.json.bac app.json
