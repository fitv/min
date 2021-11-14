## MIN - 一个基于 Gin 的 API 骨架

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
 - 基于 `Viper` 的配置
 - 基于 `Redis` 缓存器
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

## 使用说明

### 缓存
```go
import "github.com/fitv/min/global"

global.Cache().Set("key", "value", time.Minute)
global.Cache().Get("key") // value
global.Cache().Has("key") // true
global.Cache().TTL("key") // time.Minute
global.Cache().Del("key") // true
```

### 日志
```go
import "github.com/fitv/min/global"

global.Log().Debug("debug")
global.Log().Info("info")
global.Log().Warn("warn")
global.Log().Error("error")
```

### 分页及 API 资源转换
```go
package user

import (
  "context"
  "strconv"

  "github.com/fitv/min/app/resource"
  "github.com/fitv/min/core/response"
  "github.com/fitv/min/global"
  "github.com/gin-gonic/gin"
)

type User struct{}

// Index returns a list of users with pagination
func (User) Index(c *gin.Context) {
  paginator, err := global.Ent().User.
    Query().
    Paginate(context.Background(), c)
  if err != nil {
    response.HandleEntError(c, err)
    return
  }
  resource.NewUserPaginator(c, paginator).Response()
}

// Info returns user information
func (User) Info(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := global.Ent().User.Get(context.Background(), id)
  if err != nil {
    response.HandleEntError(c, err)
    return
  }
  resource.NewUser(c, user).Append(gin.H{
    "role": "staff",
  }).Wrap("user").Response()
}
```

用户列表响应
```json
{
  "current_page": 1,
  "per_page": 15,
  "last_page": 1,
  "total": 2,
  "data": [
    {
      "id": 1,
      "username": "u1"
    },
    {
      "id": 2,
      "username": "u2"
    }
  ]
}
```

用户信息响应
```json
{
  "user": {
    "id":1,
    "name":"u1"
  },
  "meta": {
    "role": "staff"
  }
}
```

### 表单验证
```go
type Auth struct{}

type UserFormLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (Auth) Login(c *gin.Context) {
	var form UserFormLogin

	err := c.ShouldBind(&form)
	if err != nil {
		response.HandleValidatorError(c, err)
		return
	}
  // do login...
}
```
验证失败响应，返回状态码`422`
```json
{
  "message": "姓名为必填字段",
  "errors": {
    "password": "密码为必填字段",
    "username": "姓名为必填字段"
  }
}
```

### 数据库事务
```go
package user

import (
  "context"

  "github.com/fitv/min/ent"
  "github.com/fitv/min/ent/user"
  "github.com/fitv/min/global"
)

type User struct{}

func (User) Update() {
  ctx := context.Background()

  global.DB().WithTx(ctx, func(tx *ent.Tx) error {
    user, err := tx.User.Query().Where(user.ID(1)).ForUpdate().First(ctx)
    if err != nil {
        return err
    }
    // do something...
  })
}
```
