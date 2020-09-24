FROM alpine:latest
RUN mkdir /workspace/runtime/log -p 
WORKDIR /workspace 
COPY .build/markman-server markman
COPY app.json.bac app.json

CMD ["./markman"]