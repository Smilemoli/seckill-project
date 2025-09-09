# file: frontend/Dockerfile

# --- 第一阶段: 构建阶段 ---
FROM node:18-alpine AS builder

WORKDIR /app

# 复制 package.json 和 package-lock.json (或 yarn.lock)
COPY package*.json ./
# 安装依赖
RUN npm install

# 复制所有源代码
COPY . .
# 打包构建前端项目
RUN npm run build

# --- 第二阶段: 运行阶段 ---
FROM nginx:1.25-alpine

# 将构建阶段生成的静态文件复制到 Nginx 的默认托管目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制我们自定义的 Nginx 配置文件
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露 80 端口
EXPOSE 80

# Nginx 默认会启动，我们不需要 CMD