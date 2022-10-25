FROM golang:1.16-alpine  AS builder



RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct



COPY  . /go/src/coolcar/server


WORKDIR  /go/src/coolcar/server
RUN  go install ./auth/...

# 安装grpc-health-probe
RUN go get github.com/grpc-ecosystem/grpc-health-probe


#多编译阶段

FROM alpine:3.15

COPY --from=builder  /go/bin/auth /bin/auth
COPY --from=builder  /go/bin/grpc-health-probe /bin/grpc-health-probe

ENV ADDR=:8081


EXPOSE 8081


ENTRYPOINT [ "/bin/auth" ]

