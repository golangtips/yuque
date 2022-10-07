FROM golang:1.18-alpine as builder

# 环境变量设置
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

# 设置镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

WORKDIR /app

# 整个源码复制到容器里面
COPY . .

RUN go build -mod=mod -o ./main -v ./cmd/

FROM alpine:3.13

WORKDIR /app

# 复制配置文件
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.toml /app/config.toml

COPY static /app/static
COPY theme /app/theme

RUN chmod +x /app/main

EXPOSE 8080

CMD /app/main