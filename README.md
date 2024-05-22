# go web项目


## 目录结构

```
.
├── config
│   ├── config.go
│   └── config.toml
├── logs
├── models
├── utils
│   ├── env.go
│   ├── exceptions.go
│   ├── logger.go
│   ├── timeutil.go
│   └── utils.go
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