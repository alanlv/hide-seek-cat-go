![](https://camo.githubusercontent.com/2b507540e2681c1a25698f246b9dca69c30548ed66a7323075b0224cbb1bf058/68747470733a2f2f676f6c616e672e6f72672f646f632f676f706865722f6669766579656172732e6a7067)
# About 躲猫猫🐱
- 一个专门为年轻人打造的社交短视频平台
- 用Gin打造的Go后端RESTFul API服务，根据后续业务的扩展，将逐步升级为分布式、微服务，充分发挥go的高并发能力
- 目前的技术选型：
    - 躲猫猫官网：React(性能优化的最佳实践)
    - APP(ios android web)；Flutter2.2(拥抱空安全，多端运行)
    - 后端服务：Koa2(正逐渐用Go替换)
- 未来的计划：
    - 主打**社交** + **短视频**
    - Flutter2.2重构APP
    - Go助力后端服务，包含RabbitMQ、分布式、微服务、Docker、CI/CD、OSS
    - Ffmpeg助力短视频，WebRTC助力实时音视频通话，Tensorflow2助力推荐系统, Python助力大数据分析
    - 躲猫猫短视频下载器：C(windows)
    - 躲猫猫视频播放器：C++(QT)
      
# Development environment

- Go SDK env
```
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\10991\AppData\Local\go-build
set GOENV=C:\Users\10991\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\10991\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\10991\go
set GOPRIVATE=
set GOPROXY=https://goproxy.cn,direct
set GOROOT=D:\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=D:\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.16.3
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\10991\AppData\Local\Temp\go-build1440782241=/tmp/go-build -gno-record-gcc-switches
```
- IDE: GoLand2020
- PostMan: 接口测试
- jmeter: 压力测试
- Go Module, please visit：https://goproxy.cn/ or https://mirrors.aliyun.com/goproxy/

# Libraries
1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架 
1. [Viper](https://github.com/spf13/viper): 配置文件读取
1. [GORM](https://github.com/go-gorm/gorm): 数据库ORM
1. [validator.v9](https://github.com/go-playground/validator): 参数校验
1. [UUID](github.com/satori/go.uuid): 生成唯一ID


# Community

## Community channels

If you have questions, or would like any assistance regarding the use of this source code, please join our community channels, your question will be answered more quickly, and it will be the most suitable place. This repository is exclusive for opening issues, and requesting resources, but feel free to be part of Hide-Seek-Cat Community.

| **QQ**                                                                                                                   | **WeChat**                                                                                                                 | **Telephone**                                                                                                          |
| :-------------------------------------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------- |
| 1099129793 | yl1099129793 | 13795950539 |

## How to contribute

_Want to contribute to the project? We will be proud to highlight you as one of our collaborators. Here are some points where you can contribute and make this product (and Flutter) even better._

- Helping to translate the readme into other languages.
- Adding documentation to the readme (a lot of go's functions haven't been documented yet).
- Write articles or make videos teaching how to use go (they will be inserted in the Readme and in the future in our Wiki).
- Offering PRs for code/tests.
- Including new functions.

Any contribution is welcome!

## Articles and videos

- [YouTube](https://www.youtube.com/channel/UClg53fJlRO-5GAwGoHjxP0A)
- [BiliBili](https://space.bilibili.com/355529756)
- [微信公众号(尹哥)](https://mp.weixin.qq.com/s?__biz=MzU2NzkxMzg2NQ==&mid=2247486880&idx=1&sn=d0092b467c18210798467b3aaf5121bb&chksm=fc94b146cbe33850ee820e42d1b483b3207652d1ea5410630caf3ce69c519b5147d717e2b259&token=812813886&lang=zh_CN#rd)