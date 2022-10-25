FROM golang:1.16-alpine AS builder

ARG APPNAME
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /go/src/coolcar/server

WORKDIR /go/src/coolcar/server

RUN GOOS=linux CGO_ENABLED=0 go install ./$APPNAME/...

#进入多阶段编译
FROM alpine:3.15

ARG APPNAME

COPY --from=builder /go/bin/$APPNAME /bin/$APPNAME

ENV ADDR ':8081'
ENV APP_NAME $APPNAME


EXPOSE 8081
EXPOSE 8080

ENTRYPOINT ["/bin/gateway"]