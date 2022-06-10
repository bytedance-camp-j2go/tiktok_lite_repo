# tiktok-lite

>  ##### J2Go team's collaborative project

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/bytedance-camp-j2go/tiktok_lite_repo">
    <img src="https://avatars.githubusercontent.com/u/104659389?s=200&v=4" alt="Logo" width="80" height="80">
  </a>


  <h3 align="center">tiktok-lite</h3>
  <p align="center">
    J2GO 小组协作项目，支持多存储方式的“抖声”服务端。
    <br />
    <a href="https://r4qkym36wq.feishu.cn/docx/doxcnU7rEwbhTeauZfotujaRZTd"><strong>本项目的汇报文档 »</strong></a>
    <br />
    <br />
    <a href="https://dy.evlic.cn">Demo地址</a>
    ·
    <a href="https://github.com/bytedance-camp-j2go/tiktok_lite_repo/issues">报告Bug</a>
    ·
    <a href="https://github.com/bytedance-camp-j2go/tiktok_lite_repo/issues">提出新特性</a>
  </p>


</p>

## 目录

- [上手指南](#上手指南)
  - [项目文件结构](#项目文件结构)
  - [启动项目有关配置](#启动项目有关配置)
- [项目依赖说明](#项目依赖说明)
- [作者](#作者)
- [鸣谢](#鸣谢)

## 上手指南

### 项目文件结构

> 包名遵守以下规则:
>
> - 只由小写字母组成。不包含大写字母和下划线等字符
> - 简短并包含一定的上下文信息
> - 不要与标准库同名
> - 不使用常用变量名作为包名
> - 使用单数而不是复数
> - 谨慎地使用缩写

```
├── bootstrap           	初始化相关函数包
├── config              	定义配置
├── config.{GO_ENV}.yaml	本地配置文件
├── config.yaml         	配置文件模版
├── dao                 	数据库 CRUD 函数包
├── driver              	实现存储驱动
├── global              	全局变量
├── handler             	路由的处理方法 作用等同于 MVC controller
├── logs                	日志默认输出路径
├── middleware          	中间件包
├── model               	数据库实体定义、实体生成相关函数包
├── response            	定义返回体结构，并附带通用的返回处理方法
├── router              	定义路由
└── util                	工具类
```

### 启动项目有关配置

1. 在启动项目前配置终端环境变量 `GO_ENV={{env}}`
2. 项目初次启动后会生成 `config.{{env}}.yaml` 配置文件，请根据 `config.yaml` 对其中数据库配置进行重写（必须）

### 项目依赖说明

- [Go-Cache](https://github.com/fanjindong/go-cache)，在视频流接口中存在非常多次的用户信息查询。
  - 开发过程中使用的是外置数据库+逻辑外键
  - 在一次网络环境较差的情况下我花了将近 10s 才获取到 5 条视频的数据
  - 使用 go-cache 本地 user-info，10s 过期
  - 消除大部分重复的 SQL 请求

- GORM，项目中重要的 SQL 框架
  - 定义了数据库连接池，复用数据库连接提高性能
  - 规范使用( where c = ? ) 使用 Gorm 提供的预编译功能，有效防止 SQL 注入，项目中预留对接口输入参数的检查函数

- Gin 项目 Web 框架，编写中间件非常简单

- Viper 项目配置管理，采用环境变量重载默认配置的方式完成配置读入

- go-snowflak 简易实现的雪花算法，用于生成全局唯一 id

- qiniu/go-sdk 用于实现七牛对象存储

- jwt-go 用于 jwt 生成、解密 token

- go-redis 用于操作 Redis

- go.uber.org/zap 高性能的日志库

- moul.io/zapgorm2 用于替换 gorm 默认的日志 -> zap

### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

| Github 主页                               | 联系方式                   |
| ----------------------------------------- | -------------------------- |
| [evlic](https://github.com/evlic)         | email: evlic.kr@icloud.com |
| [isdongrl](https://github.com/isdongrl)   | email:                     |
| [niyaoquna](https://github.com/niyaoquna) | email:                     |

### 鸣谢

- [字节跳动后端青训营](https://youthcamp.bytedance.com/)

<!-- links -->

[your-project-path]:shaojintian/Best_README_template
[contributors-shield]: https://img.shields.io/github/contributors/shaojintian/Best_README_template.svg?style=flat-square
[contributors-url]: https://github.com/shaojintian/Best_README_template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/shaojintian/Best_README_template.svg?style=flat-square
[forks-url]: https://github.com/shaojintian/Best_README_template/network/members
[stars-shield]: https://img.shields.io/github/stars/shaojintian/Best_README_template.svg?style=flat-square
[stars-url]: https://github.com/shaojintian/Best_README_template/stargazers
[issues-shield]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/shaojintian/Best_README_template.svg
[license-shield]: https://img.shields.io/github/license/shaojintian/Best_README_template.svg?style=flat-square
[license-url]: https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian