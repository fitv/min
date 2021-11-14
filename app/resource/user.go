package resource

import (
	"github.com/fitv/min/core/resource"
	"github.com/fitv/min/ent"
	"github.com/gin-gonic/gin"
)

type User struct {
	user      *ent.User
	users     []*ent.User
	paginator *ent.Paginator
}

func NewUser(c *gin.Context, user *ent.User) *resource.JsonResource {
	return resource.NewMap(c, &User{user: user})
}

func NewUsers(c *gin.Context, users []*ent.User) *resource.JsonResource {
	return resource.NewArray(c, &User{users: users})
}

func NewUserPaginator(c *gin.Context, paginator *ent.Paginator) *resource.JsonResource {
	return resource.NewPaginator(c, &User{paginator: paginator})
}

func (r *User) ToMap(c *gin.Context) gin.H {
	return gin.H{
		"id":       r.user.ID,
		"username": r.user.Username,
		// "tags": resource.When(r.user.Edges.Tags != nil, NewTags(c, r.user.Edges.Tags)),
	}
}

func (r *User) ToArray(c *gin.Context) []*resource.JsonResource {
	rs := make([]*resource.JsonResource, len(r.users))
	for i, user := range r.users {
		rs[i] = NewUser(c, user)
	}
	return rs
}

func (r *User) ToPaginator(c *gin.Context) *ent.Paginator {
	r.paginator.Data = NewUsers(c, r.paginator.Data.([]*ent.User))
	return r.paginator
}
