# 社区高并发秒杀系统 (Full-Stack Seckill Project)

[![Go version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Vue version](https://img.shields.io/badge/vue-3-brightgreen.svg)](https://vuejs.org/)
[![Docker](https://img.shields.io/badge/docker-compose-blue.svg)](https://www.docker.com/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

这是一个从零到一完整实现的、前后端分离的全栈项目，旨在模拟真实世界中的高并发商品秒杀场景。项目采用现代化的技术栈，后端使用 Go 语言以发挥其高并发性能优势，前端则采用 Vue3 全家桶构建流畅的用户交互体验。

## ✨ 项目亮点 (Features)

* **高并发秒杀架构**: 采用业界标准的 **`Redis 预减库存 + RabbitMQ 异步下单`** 方案，通过 Redis 原子操作解决超卖问题，并利用消息队列**削峰填谷**，有效保护数据库，极大提升系统并发处理能力和用户响应速度。
* **前后端分离**: 基于 **Gin** 框架构建高性能 RESTful API，**Vue3** + **Vite** 负责构建用户界面，实现了彻底的前后端分离开发模式。
* **JWT 用户认证**: 使用 **JSON Web Token (JWT)** 实现无状态的用户认证和授权，通过自定义 Gin 中间件对需要授权的接口进行有效保护。
* **数据一致性**: 在基础秒杀逻辑中，通过 **GORM** 的**数据库事务**保证“扣减库存”和“创建订单”操作的原子性；在异步架构中，通过消息队列的确认机制（ACK）和消费者逻辑保证订单数据最终落地。
* **现代化前端工程**：采用 **Vue3** + **TypeScript** 构建类型安全的前端应用，使用 **Pinia** 进行全局状态管理（如用户 Token），**Vue Router** 管理前端路由并设置导航守卫，**Element Plus** 快速构建美观的 UI 界面。
* **完整的容器化**：项目所有服务（Go 后端, Vue/Nginx 前端, MySQL, Redis, RabbitMQ）均被容器化，通过 **Docker Compose** 进行编排，实现了真正的**一键部署和环境隔离**，极大简化了开发和运维流程。


## 🛠️ 技术栈 (Technology Stack)

| 类别               | 技术                                                                 |
| ------------------ | -------------------------------------------------------------------- |
| **后端 (Backend)** | Go, Gin, GORM, `golang-jwt`, `bcrypt`                                |
| **前端 (Frontend)** | Vue3, Vite, TypeScript, Pinia, Vue Router, Axios, Element Plus, Nginx |
| **数据库与中间件** | MySQL 8, Redis, RabbitMQ                                             |
| **部署 (Deployment)**| Docker, Docker Compose                                               |

## 🚀 快速开始 (Getting Started)

### 环境要求

* Git
* Docker (`>= 20.10`)
* Docker Compose

### 部署步骤

1.  **克隆项目到本地**
    ```bash
    git clone [https://github.com/your-username/seckill-project.git](https://github.com/your-username/seckill-project.git)
    cd seckill-project
    ```

2.  **修改配置**
    在部署前，你需要设置自己的数据库密码。
    - 打开 `docker-compose.yml` 文件，修改 `mysql` 服务下的 `MYSQL_ROOT_PASSWORD` 和 `MYSQL_PASSWORD`。
    - 打开 `backend/config/config.yaml` 文件，将 `dsn` 配置中的密码修改为上一步设置的 `MYSQL_PASSWORD`。

3.  **一键启动所有服务**
    在项目根目录下，执行以下命令：
    ```bash
    docker-compose up --build -d
    ```
    * `--build`：会强制 Docker 根据 `Dockerfile` 构建新的 `backend` 和 `frontend` 镜像。
    * `-d`：表示在后台运行。
    * 第一次启动会需要几分钟时间来下载所有基础镜像和构建项目。

4.  **访问应用**
    * **前端应用**: 打开浏览器，访问 `http://localhost`
    * **RabbitMQ 管理后台**: 访问 `http://localhost:15672` (用户名/密码: `guest`/`guest`)
    * **数据库**: 可以使用 Navicat, DBeaver 等工具连接 `localhost:3306` (用户名/密码: `seckill_user`/你在 `docker-compose.yml` 中设置的密码)

## 📦 API 接口文档

| 接口路径                     | 方法 | 功能             | 是否需要认证 |
| ---------------------------- | ---- | ---------------- | ------------ |
| `/api/v1/register`           | POST | 用户注册         | 否           |
| `/api/v1/login`              | POST | 用户登录         | 否           |
| `/api/v1/profile`            | GET  | 获取用户信息     | 是           |
| `/api/v1/products`           | POST | 创建新商品       | 是           |
| `/api/v1/activities`         | GET  | 获取秒杀活动列表 | 否           |
| `/api/v1/activities`         | POST | 创建秒杀活动     | 是           |
| `/api/v1/seckill/{id}`       | POST | 执行秒杀         | 是           |

## 📈 未来可拓展方向 (TODO)

-   [ ] **订单状态查询**: 开发“我的订单”页面，供用户查询自己的秒杀订单状态。
-   [ ] **WebSocket 实时通知**: 当后台消费者成功创建订单后，通过 WebSocket 向前端用户推送实时成功通知。
-   [ ] **更精细的权限控制**: 引入 Casbin 等库，实现基于角色的访问控制（RBAC），区分普通用户和管理员权限。
-   [ ] **部署到云服务器**: 编写 shell 脚本，实现项目在云服务器上的自动化部署，并配置 HTTPS。

## 📄 开源许可证 (License)

本项目采用 [MIT License](https://opensource.org/licenses/MIT) 开源。
