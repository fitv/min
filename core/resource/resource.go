package resource

import (
	"encoding/json"

	"github.com/fitv/min/core/response"
	"github.com/fitv/min/ent"
	"github.com/gin-gonic/gin"
)

const (
	TypeMap = iota
	TypeArray
	TypePaginator
)

var (
	DefaultWrapKey   = "data"
	DefaultAppendKey = "meta"
)

type Resource interface {
	ToMap(*gin.Context) gin.H
	ToArray(*gin.Context) []*JsonResource
	ToPaginator(*gin.Context) *ent.Paginator
}

// JsonResource is a resource that can be marshalled to JSON.
type JsonResource struct {
	ctx          *gin.Context
	resource     Resource
	resourceType int
	append       gin.H
	wrap         string
}

// MissingValue is a placeholder for a missing value.
type MissingValue struct{}

// NewMap returns a new JsonResource with Map type
func NewMap(c *gin.Context, resource Resource) *JsonResource {
	return &JsonResource{ctx: c, resource: resource, resourceType: TypeMap}
}

// NewArray returns a new JsonResource with Array type
func NewArray(c *gin.Context, resource Resource) *JsonResource {
	return &JsonResource{ctx: c, resource: resource, resourceType: TypeArray}
}

// NewPaginator returns a new JsonResource with Paginator type
func NewPaginator(c *gin.Context, resource Resource) *JsonResource {
	return &JsonResource{ctx: c, resource: resource, resourceType: TypePaginator}
}

// Append adds a new key-value pair to the resource
func (r *JsonResource) Append(obj gin.H) *JsonResource {
	r.append = obj
	return r
}

// Warp sets the wrap name for the resource
func (r *JsonResource) Wrap(name string) *JsonResource {
	r.wrap = name
	return r
}

// resolve returns the resource value
func (r *JsonResource) resolve() interface{} {
	switch r.resourceType {
	case TypeMap:
		return r.filter(r.resource.ToMap(r.ctx))
	case TypeArray:
		return r.resource.ToArray(r.ctx)
	case TypePaginator:
		return r.resource.ToPaginator(r.ctx)
	default:
		return r.resource
	}
}

// filter returns the resource value after applying the filters
func (JsonResource) filter(dict gin.H) gin.H {
	data := gin.H{}

	for key, value := range dict {
		switch v := value.(type) {
		case *JsonResource:
			data[key] = v.resolve()
		case *MissingValue:
			// skip
		default:
			data[key] = value
		}
	}
	return data
}

// MarshalJSON implements the json.Marshaler interface.
func (r *JsonResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.resolve())
}

// Response returns the resource as a response.
func (r *JsonResource) Response() {
	obj := r.resolve()

	wrapKey := DefaultWrapKey
	if len(r.wrap) > 0 {
		wrapKey = r.wrap
	}

	if r.append != nil {
		obj = gin.H{
			wrapKey:          obj,
			DefaultAppendKey: r.append,
		}
	}
	response.OK(r.ctx, obj)
}

// When determines if the given condition is true then return the value or executes the given callback.
func When(ok bool, value interface{}) interface{} {
	if !ok {
		return &MissingValue{}
	}
	if fn, ok := value.(func() interface{}); ok {
		return fn()
	}
	return value
}
