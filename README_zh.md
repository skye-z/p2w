# P2W - Page To What

[English](README.md)

P2W 是一个网页转PDF或图片的工具, 支持通过命令行和HTTP接口调用

[![](https://img.shields.io/badge/Go-1.20+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/P2W-1.0.0-green)](control)
[![CodeQL](https://github.com/skye-z/p2w/workflows/CodeQL/badge.svg)](https://github.com/skye-z/p2w/security/code-scanning)

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=bugs)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)

## 安装

如果你不使用Docker, 那么在使用前你需要预先安装*Chrome*浏览器.

### 独立执行文件

首先你需要根据操作系统下载P2W的可执行文件.

如果您使用的是Linux, 请将可执行文件放置在`/usr/local/bin`目录下.

如果您使用的是Windows, 请将可执行文件放置在环境变量`PATH`中任意目录下.

### Docker

> 镜像暂未发布

```shell
docker run -d -p 12800:12800 --name p2w skye-z/p2w
```

## 使用

### 通过命令行使用

P2W提供 `pdf` 和 `image` 两个命令

* 公共 Flags
    * --url/-u: 待转换的网址
    * --path/-p: 转换后输出路径(默认 `./`)
    * --code/-c: 转换任务标识代码
    * --send/-s: 转换后发送地址
* `image` 专用 Flags
    * --element/-e: 截取元素
    * --quality/-q: 图片质量(默认90)

> 请注意, `path` 与 `send` 只能二选一, 优先 `send` 标签

```shell
# 输出PDF到当前路径
p2w pdf -u="https://github.com" -p="./"
# 发送PDF到指定地址
p2w pdf -u="https://github.com" -c="github" -s="http://localhost:8080/test"

# 输出完整页面图片到当前路径
p2w image -u="https://github.com" -p="./" -q="90"
# 发送完整页面图片到指定地址
p2w image -u="https://github.com" -q="90" -c="github" -s="http://localhost:8080/test"

# 输出指定元素图片到当前路径
p2w image -u="https://github.com" -p="./" -q="90" -e=".application-main"
# 发送指定元素图片到指定地址
p2w image -u="https://github.com" -q="90" -e=".application-main" -c="github" -s="http://localhost:8080/test"
```

### 通过HTTP接口使用

P2W提供 `/api/pdf` 和 `/api/img` 两个 `GET` 接口

* 公共参数
    * url: 待转换的网址
    * code: 转换任务标识代码
    * send: 转换后发送地址
* `image` 专用参数
    * element: 截取元素
    * quality: 图片质量(默认90)

请注意, 需要使用命令行命令 `server` 启动HTTP服务器才可访问

``` shell
p2w server -p="12800"
```

## 打包编译

### 编译可执行程序

```shell
go mod download
go mod tidy
# 为当前平台打包
go build -o p2w -ldflags '-s -w'

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o p2w -ldflags '-s -w'
```

### 打包 Docker 镜像

```shell
# 先编译Linux版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# 然后在目录下构建镜像
docker build -t skye-z/p2w:1.0.0 .
```