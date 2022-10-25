FROM golang:1.16-alpine  AS builder



RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct



COPY  . /go/src/coolcar/server


WORKDIR  /go/src/coolcar/server
RUN  go install ./blob/...



#多编译阶段

FROM alpine:3.15

COPY --from=builder  /go/bin/blob /bin/blob

ENV ADDR=:8081


EXPOSE 8081


ENTRYPOINT [ "/bin/blob" ]

