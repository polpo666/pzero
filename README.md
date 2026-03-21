# pzero

Fork of `jzero`, renamed for personal publishing and custom template development.

**English** | **[简体中文](README.zh-CN.md)**

## Overview

`pzero` is based on the [go-zero framework](https://github.com/zeromicro/go-zero) and its `goctl` tooling. It initializes `api` / `gateway` / `rpc` projects and generates server or client code from descriptive files such as `api`, `proto`, and `sql`.

Current repository target:

* Publish under `github.com/polpo666/pzero`
* Use `pzero` as the CLI command and binary name
* Keep the code generation and template system, so it can be customized later for your own internal project scaffolds
* Leave the `docs` site out of this rename scope for now

## Quick Start

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

## Templates

The built-in template system is preserved. You can later customize the templates under `cmd/pzero/.template` or host your own remote templates repository.

## Repository

* Source: https://github.com/polpo666/pzero
* Examples: see the local `examples` directory or maintain a separate examples repository later if needed
