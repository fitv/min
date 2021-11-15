package response

import (
	"fmt"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/fitv/min/core/lang"
	"github.com/fitv/min/ent"
	"github.com/fitv/min/global"
	"github.com/fitv/min/util/str"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	// regexEntNotFoundLabel match ent label from ent not found error
	regexEntNotFoundLabel = regexp.MustCompile(`^ent: (\w+) not found$`)
	// regexValidatorLabel match validator label from struct name, eg: "xxFrom", "xxFromFoo"
	regexValidatorLabel = regexp.MustCompile(`^(\w+)Form`)
)

func OK(c *gin.Context, obj interface{}) {
	if message, ok := obj.(string); ok {
		obj = gin.H{
			"message": message,
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, obj)
}

func BadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": lang.Trans("message.unauthorized"),
	})
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"message": lang.Trans("message.forbidden"),
	})
}

func NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": message,
	})
}

func UnprocessableEntity(c *gin.Context, message string, errors map[string]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"errors":  errors,
	})
}

func ServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": lang.Trans("message.server_error"),
	})
}

// HandleEntError handle Ent error
func HandleEntError(c *gin.Context, err error) {
	switch err.(type) {
	case *ent.NotFoundError:
		label := "message"
		matches := regexEntNotFoundLabel.FindStringSubmatch(err.Error())
		if len(matches) > 1 {
			label = matches[1]
		}
		NotFound(c, lang.Trans(label+".not_found"))
	default:
		global.Log().Error(fmt.Errorf("ent error: %w\n%s", err, string(debug.Stack())))
		ServerError(c)
	}
}

// HandleValidatorError handle Validator error
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok || len(errs) == 0 {
		BadRequest(c, lang.Trans("message.validate_failed"))
		return
	}

	var label, message string
	errors := make(map[string]string)
	matches := regexValidatorLabel.FindStringSubmatch(errs[0].StructNamespace())
	if len(matches) > 1 {
		label = str.ToSnakeCase(matches[1])
	}

	for i, err := range errs {
		field := str.ToSnakeCase(err.Field())
		name := lang.Trans(label + "." + field)

		errors[field] = strings.Replace(err.Translate(global.Trans()), err.Field(), name, 1)
		if i == 0 {
			message = errors[field]
		}
	}
	UnprocessableEntity(c, message, errors)
}
