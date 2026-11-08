# Podman Compose 本地中间件

本机项目中间件统一由 `backend/podman-compose.middleware.yaml` 管理，项目名
`gwa-middleware`。后端继续在 Windows 本地运行，所有发布端口绑定 127.0.0.1。

| 服务 | 地址 | 持久化 |
| --- | --- | --- |
| PostgreSQL 16 + pgvector | localhost:5432，库 gwa，用户 postgres | gwa-compose-pgvector-data |
| Redis 7 | localhost:6379 | gwa-redis-data（原卷，AOF） |
| SeaweedFS S3 | localhost:9000；Master UI http://localhost:9333 | gwa-storage-data（原卷） |
| r-nacos 0.8.8 | http://localhost:10848；SDK 8848、9848 | gwa-compose-rnacos-data |
| Jaeger 2.23.0 | http://localhost:16686；OTLP 4317、4318 | gwa-compose-jaeger-data（Badger，48 小时保留） |

## 日常管理

在 `backend/` 下执行：

```powershell
podman compose --env-file .local/middleware.env -f podman-compose.middleware.yaml up -d
podman compose --env-file .local/middleware.env -f podman-compose.middleware.yaml ps
podman compose --env-file .local/middleware.env -f podman-compose.middleware.yaml logs --tail 100
podman compose --env-file .local/middleware.env -f podman-compose.middleware.yaml stop
```

本机已生成私有 `.local/middleware.env`，不可提交；字段模板为
`backend/middleware.env.example`。数据库、Redis、S3 仍沿用原本机开发凭证；
r-nacos 使用 `.local/rnacos/credentials.json` 中的随机密码，已开启 API 鉴权。
SeaweedFS 凭证使用 `scripts/docker/seaweedfs/s3.json` 的开发设置。

在全新机器部署时，先按模板创建私有 env 文件，并创建复用卷：

```powershell
podman volume create gwa-redis-data
podman volume create gwa-storage-data
```

再运行 `up -d`。PostgreSQL 首次启动创建 gwa 库；pgvector 扩展由应用启动迁移
或数据库管理员执行 `CREATE EXTENSION IF NOT EXISTS vector` 创建。

## 本机迁移记录（2026-11-08）

原 `gwa-postgres`、`gwa-redis`、`gwa-storage` 容器已停止保留，不要与 Compose
实例同时启动。原 PostgreSQL 是 Alpine 发行版，改用 pgvector 镜像时使用逻辑
备份恢复到独立新卷，未把原数据库目录直接挂给不同发行版。备份文件位于
`backend/.local/gwa-before-compose.dump`。旧 `gwa-postgres-data` 卷保持原状。
Redis 与对象存储复用原数据卷，因此它们不是独立回滚副本。

此前 Windows 原生 r-nacos 进程已经停止；原目录与数据库保留在
`backend/.local/rnacos/`，活跃配置已导入容器版 r-nacos。
后端以 `--config-source=nacos --conf .local/rnacos/configs` 启动，
从 `.local/rnacos/configs/config.yaml` 引导，监听 7788 / 7789。
远程 YAML 已添加 OTLP gRPC 指向 `localhost:4317`，Jaeger 可查询到
`admin-service`。远程配置变更仍按重启后端生效。

验证内容：Compose 五服务运行、PostgreSQL / Redis 健康检查通过、恢复原用户
数据、pgvector 0.8.7 向量类型可用、S3 完整临时对象读写清理、后端验证码接口
正常、r-nacos 鉴权读写正常、Jaeger 收到实际后端链路。

项目任务调度使用 Redis 上的 Asynq，无需另起任务队列服务。MySQL 是数据库
替代选项；AI 模型提供商需在管理页配置实际服务，pgvector 已准备好。
