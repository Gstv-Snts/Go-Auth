FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=off
RUN cd /build && git clone https://github.com/Gstv-Snts/Go-Auth.git

RUN cd /build/Go-Auth && go build main.go

EXPOSE 8080

ENTRYPOINT [/build/Go-Auth]
