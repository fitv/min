package v1

import (
	"context"

	"github.com/fitv/min/app/resource"
	"github.com/fitv/min/core/auth"
	"github.com/fitv/min/core/lang"
	"github.com/fitv/min/core/response"
	"github.com/fitv/min/ent/user"
	"github.com/fitv/min/global"
	"github.com/fitv/min/util/hash"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

type UserFormLogin struct {
	Username string `json:"username" binding:"required,min=3,max=12"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

func (Auth) Login(c *gin.Context) {
	var form UserFormLogin

	err := c.ShouldBind(&form)
	if err != nil {
		response.HandleValidatorError(c, err)
		return
	}

	user, err := global.Ent().User.
		Query().
		Where(user.Username(form.Username)).
		First(context.Background())
	if err != nil {
		response.HandleEntError(c, err)
		return
	}

	if !hash.Check(user.Password, form.Password) {
		response.BadRequest(c, lang.Trans("auth.invalid_credentials"))
		return
	}

	token, err := auth.SignToken(user.ID)
	if err != nil {
		response.BadRequest(c, lang.Trans("auth.error"))
		return
	}
	response.OK(c, NewAccessToken(token))
}

func (Auth) Register(c *gin.Context) {
	var form UserFormLogin

	err := c.ShouldBind(&form)
	if err != nil {
		response.HandleValidatorError(c, err)
		return
	}

	exist, err := global.Ent().User.
		Query().
		Where(user.Username(form.Username)).
		Exist(context.Background())
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	if exist {
		response.BadRequest(c, lang.Trans("auth.username_existed"))
		return
	}

	hashPassword, err := hash.Make(form.Password)
	if err != nil {
		response.ServerError(c)
		return
	}

	user, err := global.Ent().User.Create().
		SetUsername(form.Username).
		SetPassword(string(hashPassword)).
		Save(context.Background())
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	resource.NewUser(c, user).Response()
}

func (Auth) Refresh(c *gin.Context) {
	token, err := auth.SignToken(auth.MustUID(c))
	if err != nil {
		response.BadRequest(c, lang.Trans("auth.error"))
		return
	}
	response.OK(c, NewAccessToken(token))
}

func (Auth) Profile(c *gin.Context) {
	user, err := global.Ent().User.Get(context.Background(), auth.MustUID(c))
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	resource.NewUser(c, user).Response()
}
