# wechat SDK for Go

## 概述
| 模块    | 描述                     |
|--------:|:-------------------------|
| mp      | 微信公众平台         |
| work    | 企业微信         |
| util    | 公共文件                 |
| examples| Demo文件                 |

## 安装
go get -u github.com/yijizhichang/wechat-sdk

##分支说明
###master分支
早期分支，主要支持微信公众号相关接口，内部封装了日志输出及打印，便于发现跟微信交互时，产生的错误；
不方便之处：sdk的日志与应用的日志是分离的，不便于日志的统一管理。
###simple分支
简化包功能
缓存的方式由应用自己实现sdk中的cache接口来传入，方便应用与包共用一套缓存配置
取消包内打日志，直接将错误返回，由应用程序统一处理。
支持企业微信相关接口

持续更新，推荐使用simple分支

##develop,work分支
为开发版，内测版分支

## 详细文档说明

- [微信公众号](README.MP.md)