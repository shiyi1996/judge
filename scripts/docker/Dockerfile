FROM java:alpine
MAINTAINER ShiYi <shiyi@fightcoder.com>

RUN apk upgrade --update;
RUN apk add g++;
RUN apk add python;
RUN apk add git;
RUN apk add curl;
RUN curl https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz | tar xzf - -C /; \
    mv /go /goroot;

RUN mkdir /lib64; \
    ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2;

ENV GOPATH /workspace:/sandbox:/sandbox/deps
ENV GOROOT /goroot
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOROOT/bin

RUN cd /; \
    git clone https://github.com/shiyi1996/sandbox.git; \
    cd sandbox; \
    go build;

WORKDIR /workspace

CMD while true; do sleep 1; done