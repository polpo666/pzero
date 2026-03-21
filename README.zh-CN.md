# pzero

`jzero` 的个人化重命名分支，用于自行发布以及后续定制模板体系。

**简体中文** | **[English](README.md)**

## 简介

`pzero` 基于 [go-zero 框架](https://github.com/zeromicro/go-zero) 及其 `goctl` 工具链，支持初始化 `api` / `gateway` / `rpc` 项目，并可根据 `api`、`proto`、`sql` 等描述文件生成服务端与客户端代码。

当前这个仓库的目标是：

* 以 `github.com/polpo666/pzero` 对外发布
* 将 CLI 与二进制统一为 `pzero`
* 保留现有模板和代码生成能力，方便后续继续按你的项目需求定制
* 暂时不处理 `docs` 站点相关内容

## 快速开始

```shell
go install github.com/polpo666/pzero/cmd/pzero@latest

pzero check
pzero new your_project
cd your_project
go mod tidy
go run main.go server
```

## Docker

```shell
docker run --rm -v ${PWD}:/app ghcr.io/polpo666/pzero:latest new your_project
```

## 模板

内置模板系统已经保留，后续你可以直接定制 `cmd/pzero/.template`，或者再维护你自己的远程模板仓库。

## 仓库信息

* 源码仓库: https://github.com/polpo666/pzero
* 示例代码: 当前先看本仓库的 `examples` 目录，后续如果需要再拆分独立 examples 仓库
