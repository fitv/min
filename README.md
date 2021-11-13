## MIN - an API skeleton based on Gin

[简体中文](README_zh.md)

## Packages used
- [Gin](https://github.com/gin-gonic/gin)
- [ent](https://entgo.io/ent)
- [viper](https://github.com/spf13/viper)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [go-redis](https://github.com/go-redis/redis)
- [validator](https://github.com/go-playground/validator)

## Features
 - Easily build your `RESTful APIs`
 - `JWT` Authorization
 - Configuration
 - Cache based on `Redis`
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
cp config.example.yaml config.yaml
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

global.Cache().Set("key", "value", time.Minute)
global.Cache().Get("key")
global.Cache().Has("key")
global.Cache().TTL("key")
global.Cache().Del("key")
```

### Logger
```go
import "github.com/fitv/min/global"

global.Log().Debug("debug")
global.Log().Info("info")
global.Log().Warn("warn")
global.Log().Error("error")
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

User Index output
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

User Info output
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
  ctx := context.Background()

  global.DB().WithTx(ctx, func(tx *ent.Tx) error {
    user, err := tx.User.Query().Where(user.ID(1)).ForUpdate().First(ctx)
    if err != nil {
        return err
    }
    // do something...
    return nil
  })
}
```
