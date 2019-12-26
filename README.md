# go-lean
## 项目简介
go学习项目

## 依赖初始化

>go mod init 模块名

>go mod download 依赖下载

## 配置信息
ymal
ini
## 日志说明

* EngineLogger 引擎相关日志记录

* GinLogger gin路由请求相关日志记录


```
EngineLogger = NewLogger("src/logs/app_engine.log", zapcore.InfoLevel, 128, 0, 0, true, "EngineApp")
GinLogger = NewLogger("src/logs/gin_engine.log", zapcore.InfoLevel, 128, 0, 0, true, "EngineGin")
```

#### 日志参数
* filePath 日志文件路径
* level 日志级别
* maxSize 每个日志文件保存的最大尺寸 单位：M
* maxBackups 日志文件最多保存多少个备份
* maxAge 文件最多保存多少天
* compress 是否压缩
* serviceName 服务名



#### 日志调用

* logger.EngineLogger.Error 错误日志
* logger.EngineLogger.Panic panic日志
* logger.EngineLogger.Info info日志

## 技术栈
>gin-gonic 

>zap

> lumberjack

> yaml

> gorm

## 包结构
> components 项目组件

> common 常量和实体

> config 配置信息

> controller 接口路由

> openapi api接口实现

> repository 数据库操作

> thirds 三方插件

