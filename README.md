## MIN - an API skeleton based on Gin

[简体中文](README_zh.md)

## Packages used
- [Gin](https://github.com/gin-gonic/gin)
- [ent](https://entgo.io/ent)
- [viper](https://github.com/spf13/viper)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [go-redis](https://github.com/go-redis/redis)
- [validator](https://github.com/go-playground/validator)
- [swagger-ui](https://github.com/swagger-api/swagger-ui)

## Features
 - Easily build your `RESTful APIs`
 - `JWT` Authorization
 - Configuration
 - Cache based on `Redis`
 - `ent` Paginator
 - API Resource Transformer
 - Multi-language translation support
 - Swagger UI (render YAML file)

## Installation
Clone repository
```
git clone https://github.com/fitv/min.git
```

Copy the configuration file
```
cd min && cp config.example.yaml config.yaml
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

## Usage

### Cache
```go
import "github.com/fitv/min/global"

global.Cache().Set(ctx, "key", "value", time.Minute)
global.Cache().Get(ctx, "key") // value
global.Cache().Has(ctx, "key") // true
global.Cache().TTL(ctx, "key") // time.Minute
global.Cache().Del(ctx, "key") // true
```

### Logger
```go
import "github.com/fitv/min/global"

global.Log().Debug("debug")
global.Log().Info("info")
global.Log().Warn("warn")
global.Log().Error("error")
```

### Multi-language translation
```go
import "github.com/fitv/min/global"

global.Lang().Trans("hello.world") // world
global.Lang().Locale("zh").Trans("hello.world") // 世界
```

### Paginator & API Resource Transformer
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
    Paginate(ctx, c)
  if err != nil {
    response.HandleEntError(c, err)
    return
  }
  resource.NewUserPaginator(c, paginator).Response()
}

// Info returns user information
func (User) Info(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := global.Ent().User.Get(ctx, id)
  if err != nil {
    response.HandleEntError(c, err)
    return
  }
  resource.NewUser(c, user).Append(gin.H{
    "role": "staff",
  }).Wrap("user").Response()
}
```

User Index response
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

User Info response
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

### From Validator
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
Verification failure response, return status code `422`
```json
{
  "message": "username is a required field",
  "errors": {
    "password": "password is a required field",
    "username": "username is a required field"
  }
}
```

### Database Transaction
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
  global.DB().WithTx(ctx, func(tx *ent.Tx) error {
    user, err := tx.User.Query().Where(user.ID(1)).ForUpdate().First(ctx)
    if err != nil {
        return err
    }
    // do update...
  })
}
```
