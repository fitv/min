package v1

import (
	"context"

	"github.com/fitv/min/core/response"
	"github.com/fitv/min/global"
	"github.com/gin-gonic/gin"
)

type Migrate struct{}

// Index run database migration, you should remove this method in your production environment.
func (Migrate) Index(c *gin.Context) {
	err := global.Ent().Schema.Create(
		context.Background(),
		// schema.WithDropColumn(true),
		// schema.WithDropIndex(true),
	)
	if err != nil {
		response.HandleEntError(c, err)
		return
	}
	response.OK(c, global.Lang().Trans("message.success"))
}
