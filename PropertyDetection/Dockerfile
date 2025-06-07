# 第一阶段：构建可执行文件
FROM golang:1.24-alpine as builder
# 设置工作目录
WORKDIR /app
# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./
# 设置 Go 模块代理
ENV GOPROXY=https://mirrors.aliyun.com/goproxy
# 下载依赖
RUN go mod download
# 复制项目文件
COPY . .
# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o property-detection .
# 第二阶段：创建轻量级镜像
FROM alpine:latest
# 安装必要的依赖
RUN apk --no-cache add ca-certificates
#RUN apk add --no-cache git
# 设置工作目录
WORKDIR /root/
# 从 builder 阶段复制可执行文件和配置文件
COPY --from=builder /app/property-detection .
COPY --from=builder /app/config.yaml .
# 暴露端口
EXPOSE 7888
# 启动服务
CMD ["./property-detection"]