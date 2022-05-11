# tiktok_lite_repo
> J2Go team's collaborative project

## 目录结构

> 包名遵守以下规则: 
>
> - 只由小写字母组成。不包含大写字母和下划线等字符
> - 简短并包含一定的上下文信息
> - 不要与标准库同名
> - 不使用常用变量名作为包名
> - 使用单数而不是复数
> - 谨慎地使用缩写



## 配置文件描述

加载配置文件时会从系统环境变量读取 GO_ENV 变量作为配置文件名称
1. `config.{GO_ENV}.yaml`   最高优先级
2. `config.def.yaml`        默认加载

- 配置文件读取顺序 > `config.{GO_ENV}.yaml` >> `config.yaml` 
- 两者相同部分以 `config.yaml` 为准:

