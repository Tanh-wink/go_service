# go web项目

基于Gin框架的web服务

已实现：日志采集logger，异常处理，统计耗时，企业微信告警，参数验证

## 目录结构

```
.
├── config
│   ├── config.go
│   └── config.toml
├── dto
│   ├── request.go
│   └── response.go
├── logs
├── models
├── utils
│   ├── env.go
│   ├── logger.go
│   ├── timeutil.go
│   ├── utils.go
│   └── webhook.go
├── web
│   ├── controller
│       └── demo.go
│   ├── middleware
│       ├── exceptionHandler.go
│       └── timeHandler.go
│   ├── routers
│       └── router.go
│   └── template
│       └── static
│           └── home.html
├── go.mod
├── go.sum
├── main.go
├── README.md

```

## 环境要求

- go.mod



## 启动 Web 服务

```sh
go run main.go
```