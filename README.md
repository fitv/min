## MIN - an API skeleton based on Gin

[English](README.md) | [中文](README_zh.md)

## Packages used
- [Gin](https://github.com/gin-gonic/gin)
- [ent](https://entgo.io/ent)
- [viper](https://github.com/spf13/viper)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [go-redis](https://github.com/go-redis/redis)
- [validator](https://github.com/go-playground/validator)

## Features
 - Easily build your `RESTful API`
 - `JWT` Authorization
 - Configuration
 - Cache
 - Logger
 - `ent` Paginator
 - API Resource Transformer
 - Multi-language translation support

## Installation
Clone repository
```
git clone https://github.com/fitv/min.git
```

Copy the configuration file
```
cp config.example.yml config.yml
```

Run Application
```
go run main.go
```

Build Application
```
go build -o min main.go
```

Open the API document URL in your browser
```
http://127.0.0.1:3000/apidoc
```

Run database migration
```
http://127.0.0.1:3000/api/v1/migrate
```
