FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN cd /build && git clone https://github.com/nathancrjackson/wimi-go.git

RUN cd /build/wimi-go/src && go build webserver.go && mv /build/wimi-go/src/webserver /build

EXPOSE 8081

ENTRYPOINT [ "/build/webserver" ]
