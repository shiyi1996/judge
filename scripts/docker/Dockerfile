FROM java:alpine
MAINTAINER ShiYi <shiyi@fightcoder.com>

RUN apk upgrade --update;
RUN apk add g++;
RUN apk add python;
RUN apk add curl;
RUN apk add go;
RUN apk add git;

 ENV GOPATH /workspace:/sandbox:/sandbox/deps

 RUN cd /; \
 git clone https://github.com/shiyi1996/sandbox.git;  \
 cd sandbox; \
 go build;

 WORKDIR /workspace

 CMD while true; do sleep 1; done