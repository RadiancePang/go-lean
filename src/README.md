# 线上推理平台
## 项目简介
机器学习线预测服务

## 依赖初始化
go mod init 模块名
go mod download 依赖下载
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

## 依赖包
>gin-gonic 

>zap

> lumberjack

> yaml

> gorm

## 包结构

> engine 服务运行相关

> app 应用相关代码