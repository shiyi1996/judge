FROM alpine:latest
MAINTAINER ShiYi <shiyi@fightcoder.com>

RUN apk upgrade --update; \
apk add python; \
apk add g++; \
apk add git; \
apk add go;

ADD . /workspace

ENV GOPATH /workspace:/workspace/fightcoder-sandbox:/workspace/fightcoder-sandbox/deps:


WORKDIR /workspace

CMD while true; do sleep 1; done