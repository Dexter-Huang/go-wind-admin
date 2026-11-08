# 本地维护的 tx7do 模块

本目录收录后端依赖图中的全部 63 个 `github.com/tx7do/*` 模块，包括直接依赖和间接依赖。
源码来自迁移时 `backend/go.mod` 选定版本的 Go module cache，没有升级依赖。
`modules.json` 记录原模块路径、版本、Go 校验和与本地目录；每个模块保留上游 LICENSE。

`backend/go.mod` 的无版本 `replace` 将这些模块映射到本目录。正常在 `backend/` 下执行
`go build ./...`、`go test ./...` 或 `gow run admin`，都会使用这里的源码。
原 `github.com/tx7do/...` 导入路径保留，以维持生成代码、跨模块引用和 Go internal 包规则。
需要修复库代码时，直接修改本目录对应文件即可。

## 验证与维护

在 `backend/` 执行：

```sh
python scripts/check_local_tx7do.py
go build -mod=readonly ./...
go test -mod=readonly ./...
```

检查脚本确认当前依赖图中的所有 tx7do 模块都使用仓库内源码、目录与清单一致，且许可证存在。
修改子模块后，应从 `backend/` 验证应用；单独进入子模块运行 Go 命令时，主模块的 replace 不生效。
`modules.json` 的校验和用于记录原始发布版本，不代表本地修改后的文件校验和。
后续引入新的 tx7do 模块，必须同时纳入源码、清单和主模块 replace；检查脚本会拒绝遗漏。

本地化覆盖后端依赖图，不包括 `Makefile` 另外安装的独立开发工具（如 gow 和 TypeScript
生成器），也不包括其他作者的依赖。使用 `go install 模块@版本` 会绕开主模块 replace；
已纳入本目录的 redact 生成器应在 `backend/` 执行 `go install github.com/tx7do/go-wind-toolkit/protoc-gen-go-redact`。

## 随源码携带的数据

GeoIP 模块中包含约 119 MiB 的 IP 地址数据库，由上游源码通过 `go:embed` 引用，故一并保留。
Docker 构建上下文会包含本目录，现有 `COPY . /src` 可以直接使用这些本地模块。
源码本地化改变的是维护位置，并不取消原有的许可证及数据文件使用条款。
