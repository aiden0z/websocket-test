FROM golang:latest as builder

ENV GO111MODULE=on GOPROXY=https://goproxy.io,direct
WORKDIR /opt/ws/
COPY . .
RUN go build main.go

FROM ubuntu

ARG version
LABEL maintainer="Aiden Luo <aiden0xz@gmail.com>" version=${version}
WORKDIR /opt/ws/

COPY --from=builder /opt/ws/main .
COPY static /opt/ws/static
COPY views /opt/ws/views
ENV PATH=/opt/ws:$PATH

CMD ["/opt/ws/main"]