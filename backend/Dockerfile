# file: backend/Dockerfile

# --- 第一阶段: 构建阶段 ---
# 使用官方的 Go 镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download

# 复制所有源代码
COPY . .

# 构建应用
# CGO_ENABLED=0: 禁用 CGO，构建静态二进制文件
# GOOS=linux: 确保我们构建的是 Linux 可执行文件，因为我们的最终镜像将是 Alpine Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- 第二阶段: 运行阶段 ---
# 使用一个非常小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY ./config ./config

# 暴露端口
EXPOSE 8080

# 容器启动时运行的命令
CMD ["./main"]