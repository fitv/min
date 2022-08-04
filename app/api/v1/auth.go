package v1

import (
	"context"

	"github.com/fitv/min/app/resource"
	"github.com/fitv/min/core/auth"
	"github.com/fitv/min/core/response"
	"github.com/fitv/min/ent/user"
	"github.com/fitv/min/global"
	"github.com/fitv/min/util/hash"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

// UserFormLogin user login form
type UserFormLogin struct {
	Username string `json:"username" binding:"required,min=3,max=12"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// Login user login
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
		response.BadRequest(c, global.Lang().Trans("auth.invalid_credentials"))
		return
	}

	token, err := auth.SignToken(user.ID)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OK(c, NewAccessToken(token))
}

// Register user register
func (Auth) Register(c *gin.Context) {
	var form UserFormLogin

	err := c.ShouldBind(&form)
	if err != nil {
		response.HandleValidatorError(c, err)
		return
	}

	user, err := global.Ent().User.Create().
		SetUsername(form.Username).
		SetPassword(form.Password).
		Save(context.Background())
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	resource.NewUser(c, user).Response()
}

// RefreshToken refresh JWT token
func (Auth) Refresh(c *gin.Context) {
	token, err := auth.SignToken(auth.MustUID(c))
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OK(c, NewAccessToken(token))
}

// Profile get current authorized user info
func (Auth) Profile(c *gin.Context) {
	user, err := global.Ent().User.Get(context.Background(), auth.MustUID(c))
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	resource.NewUser(c, user).Response()
}
