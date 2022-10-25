FROM golang:1.16-alpine AS  builder

ARG APPNAME

RUN go env -w  GO111MODULE=on

RUN go env -w  GOPROXY=https://goproxy.io,https://goproxy.cn,direct

COPY . /go/src/coolcar/server

WORKDIR /go/src/coolcar/server

RUN GOOS=linux CGO_ENABLED=0 go install ./$APPNAME/...
# 安装grpc-health-probe
RUN go get github.com/grpc-ecosystem/grpc-health-probe


FROM alpine:3.15

ARG APPNAME

COPY --from=builder  /go/bin/$APPNAME /bin/$APPNAME
COPY --from=builder  /go/bin/grpc-health-probe /bin/grpc-health-probe


ENV ADDR ':8081'

ENV APP_NAME $APPNAME

EXPOSE 8081

ENTRYPOINT ["/bin/rental"]