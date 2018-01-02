FROM alpine:latest
MAINTAINER ShiYi <shiyi@fightcoder.com>

RUN apk upgrade --update; \
apk add python; \
apk add g++; \
apk add git; \
apk add go;

ENV GOPATH /workspace:/sandbox:/sandbox/deps

RUN cd /; \
git clone https://github.com/shiyi1996/sandbox.git; \
cd sandbox; \
go build;

WORKDIR /workspace

CMD while true; do sleep 1; done