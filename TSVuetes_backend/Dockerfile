# 使用官方的 Go 镜像
FROM golang:1.23.2-alpine AS build

# 设置工作目录
WORKDIR /app

# 复制所有源代码到工作目录，包括 config 文件夹
COPY . .

# 下载依赖并构建 Go 应用
RUN go mod tidy
RUN go build -o main .

# 使用轻量级的 Alpine 镜像作为运行环境
FROM alpine:latest
WORKDIR /root/

# 创建 config 目录来存放配置文件
RUN mkdir -p /root/config

# 从构建阶段（build）复制编译后的二进制文件到最终运行阶段
COPY --from=build /app/main .

# 复制 config 文件夹中的 config.yml 到容器中的 /root/config 目录
COPY --from=build /app/config/config.yml /root/config/config.yml

# 暴露后端服务的端口
EXPOSE 8080

# 运行二进制文件
CMD ["./main"]