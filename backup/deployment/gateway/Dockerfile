# 启动编译环境
FROM golang:1.16-alpine AS builder

# 配置编译环境
RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中
COPY  .  /go/src/coolcar/server

# 编译
WORKDIR /go/src/coolcar/server
#跨平台编译GOOS=linux,GOOS=windows
#CGO_ENABLED=0 设置go可执行文件与系统的静态链接、动态链接等
RUN  GOOS=linux  go  install ./gateway/...


#vim /etc/hosts
#添加如下内容
#140.82.114.4    github.com
#199.232.5.194   github.global.ssl.fastly.net

#docker  multi stage 多阶段编译 --from中的builder对应的就是顶部第一阶段中的builder
FROM alpine:3.15
#ADD 适合网络文件，但是官方不建议，会需要更多曾的镜像，更将以使用 wget之后再进行操作
#ADD 还适合将压缩文件tar.gz 拷贝到镜像中，自动将压缩包从source复制到镜像中，并对压缩包进行解压缩到dest中
COPY --from=builder /go/bin/gateway /bin/gateway

ENV ADDR=:8080

EXPOSE 8080

#CMD 与 ENTRYPOINT  不变的写在 ENTRYPOINT中，动态可变、可被用户输入所覆盖的参数写在CMD中
#CMD ["executable"]

# 设置服务入口
ENTRYPOINT [ "/bin/gateway" ]