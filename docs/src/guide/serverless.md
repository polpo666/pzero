---
title: 插件指南
icon: arcticons:game-plugins
star: true
order: 5.4
---

pzero 生成的新 frame 项目已经不再把基于 serverless 的插件脚手架作为推荐路径。

## 当前状态

* 新的 pzero frame 模板不再把 `--serverless` 作为主要插件工作流来文档化
* 旧的 jzero 项目中可能仍然保留历史 serverless/plugin 结构，但新项目应以 `pzero new` 当前生成的模板为准
* 如果新项目需要扩展能力，建议基于现有模板里的 `plugins` 集成点和普通模块边界来设计，而不是继续沿用旧的 serverless 目录布局

## 推荐做法

对于新项目，使用正常的 frame 初始化和代码生成流程：

```bash
pzero new your_project --frame api
cd your_project
go mod tidy
pzero add api demo
pzero gen
```

如果你维护的是仍然使用旧版 serverless/plugin 结构的历史 jzero 项目，可以把那套布局视为兼容模式，而不是 pzero 当前默认架构。