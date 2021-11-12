## MIN - 一个基于 GIN 的 API 骨架

[English](README.md)

## 使用的第三方包
- [Gin](https://github.com/gin-gonic/gin)
- [ent](https://entgo.io/ent)
- [viper](https://github.com/spf13/viper)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [go-redis](https://github.com/go-redis/redis)
- [validator](https://github.com/go-playground/validator)

## 特性
 - 快速构建 `RESTful API`
 - 基于 `JWT` 的用户授权
 - 基于 `Viper` 的文件配置
 - 缓存
 - 日志
 - `ent` 模型分页器
 - API 资源转换器
 - 多语言翻译支持

## 安装
克隆仓库
```
git clone https://github.com/fitv/min.git
```

复制配置文件
```
cp config.example.yaml config.yaml
```

运行应用
```
go run main.go
```

构建应用
```
go build -o min main.go
```

API文档地址
```
http://127.0.0.1:3000/apidoc
```

运行数据库迁移
```
http://127.0.0.1:3000/api/v1/migrate
```
