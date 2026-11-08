# r-nacos 配置中心

后端通过 `--config-source=local|nacos` 明确选择配置来源，默认 `local`。
配置中心插件在 admin 入口注册，无需开启服务注册发现。切换后需重启服务。

从 `backend/` 使用编译后的 admin 二进制启动：

```powershell
# 仅加载本地 YAML，不连接 Nacos（即使目录中存在 Nacos 连接配置）
.\.local\rnacos\admin-source.exe --config-source=local --conf app/admin/service/configs

# 仅加载 Nacos 业务配置，本地 config.yaml 只提供连接信息
.\.local\rnacos\admin-source.exe --config-source=nacos --conf .local/rnacos/configs
```

`--conf` 是目录路径。Nacos 模式必须有 `config.yaml`，且其中
`config.type` 为 `nacos` 并提供 `config.nacos`；缺少连接配置或模式拼错会报错。
Nacos 模式不会合并同目录中的本地业务 YAML；不会因缺少连接配置而切回 local。

2026-11-09 已验证来源隔离测试、启动参数传递测试、后端完整构建，
并分别以 local / nacos 模式启动完整后端，验证码接口均正常。

## 配置组织

每个环境使用独立 Namespace ID，Group 为 `GO_WIND_ADMIN`，Data ID 为
`admin-service.yaml`。将 `backend/app/admin/service/configs/` 下六份 YAML 合并，
保留 `server`、`data`、`client`、`authn`、`authz`、`logger`、`oss` 顶层键。
不要上传本地引导文件 `config.yaml`。

本地引导配置参考 `backend/configs-nacos.example/config.yaml`。部署时替换凭证，
放入私有目录，启动 admin 二进制时传 `--config-source=nacos --conf <私有目录>`。该目录可以仅含
`config.yaml`，此时所有业务运行配置来自远程。配置加载器不展开 `${ENV}`。

r-nacos 默认 API 端口 8848、客户端 gRPC 端口 9848、控制台端口 10848。
客户端须能访问前两个端口。启用 `RNACOS_ENABLE_OPEN_API_AUTH=true`，设置初始
管理账号密码并持久化数据目录；初始账号变量仅在首次初始化时有效。

## 生效与故障语义

当前应用仅在启动时 Scan 配置并创建基础设施组件。远程发布后应重启 admin；
订阅收到变化不等于数据库连接、监听地址和日志器已经更新。

SDK 提供本地快照相关选项，Nacos 模式不回退到本地业务 YAML。
若希望无远程配置时拒绝启动，应额外验证必需字段和空内容。不要把 SDK 缓存
当成受控发布机制；缓存目录也可能含敏感信息，应与引导凭证一起限制访问。

`GOWIND_CRYPTO_KEY`、RSA 私钥等继续由部署环境注入，部分值在远程配置加载前
已经读取。前端构建配置与数据库中的业务/租户配置继续使用原有管理方式。

## 本地测试实例

本地二进制、数据、引导配置、凭证和日志放在 `backend/.local/rnacos/`，
通过本仓 `.git/info/exclude` 排除，不提交真实配置和凭证。
目前已迁移为 Podman Compose 容器，原生进程已停止，见
[Podman 中间件部署](./podman-middleware.md)。实例绑定 `127.0.0.1`，控制台为 <http://127.0.0.1:10848/>；测试账号凭证记录在
该目录的 `credentials.json`。启动目录必须保持不变，避免切换到新的数据库。

官方资料：<https://github.com/nacos-group/r-nacos>。

### 2026-11-08 实测结果

以下为引入显式开关前的历史测试；现行 Nacos 模式不再合并本地业务 YAML。

Windows 官方二进制 r-nacos v0.8.8，配置插件 v0.2.5（Nacos Go SDK v2.3.5）。

| 场景 | 结果 |
| --- | --- |
| 控制台访问 | HTTP 200 |
| 匿名读取远程配置 | HTTP 403 |
| 鉴权发布、读取合并 YAML | 成功 |
| 本地仅有引导文件启动完整 admin | HTTP 7788、SSE 7789、Asynq 启动成功 |
| 调用 `/admin/v1/captcha` | 返回 captchaId、imageBase64 |
| 本地端口 19990，远程端口 19991 | 加载结果 19991，远程优先 |
| 远程 YAML 语法错误 | 配置加载失败，探针退出码 1 |
| 不可达端口且禁用快照 | 配置加载失败，探针退出码 1 |
| 不可达端口且已有快照，使用默认缓存选项 | 加载成功，返回已缓存端口 |
| 在线发布端口从 7788 改为 7790 | 配置探针读到 7790，既有 admin 仍监听 7788 |

端口变更测试结束已恢复远程配置为 7788。缓存与不可达测试使用独立探针，
没有停止正在运行的 r-nacos。React、Vue Element 类型检查通过；Vben 的
`pnpm run typecheck` 因 Corepack 下载 pnpm 时 TLS 连接重置未能执行，随后
通过 `npm run typecheck` 执行同一 package script，检查通过，未修改工具链版本。
后端完整构建通过，`go test ./app/admin/service/cmd/server` 通过（该包无测试文件）。

从 `backend/` 重启测试后端的命令（先停止已有测试后端，避免端口冲突）：

```powershell
.\.local\rnacos\admin-source.exe --config-source=nacos --conf .local/rnacos/configs
```
