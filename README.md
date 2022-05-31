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

```
├── bootstrap			初始化
├── config				定义
├── config.dev.yaml		本地配置文件
├── config.yaml			配置文件模版
├── config_test.go		测试配置初始化，输出配置结果
├── dao					DAO
├── driver				实现存储驱动
├── global				全局变量
├── go.mod
├── go.sum
├── handler				路由的处理方法 作用等同于 MVC controller
├── logs				日志默认输出路径
├── main.go	
├── middleware			定义中间件
├── model				数据库实体，类似于 JAVA PO
├── response			定义返回体	DTO
├── router				定义路由
└── util				工具类
```

## 配置文件描述

加载配置文件时会从系统环境变量读取 GO_ENV 变量作为配置文件名称

1. `config.{GO_ENV}.yaml`   最高优先级
2. `config.def.yaml`        默认加载

- 配置文件生效优先顺序 > `config.{GO_ENV}.yaml` >> `config.yaml`
- 两者相同部分以 `config.yaml` 为准:

