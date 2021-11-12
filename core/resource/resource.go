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

	for k, v := range dict {
		if r, ok := v.(*JsonResource); ok {
			data[k] = r.resolve()
		} else if _, ok := v.(*MissingValue); ok {
			continue
		} else {
			data[k] = v
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

	wrapKey := "data"
	if len(r.wrap) > 0 {
		wrapKey = r.wrap
	}

	if r.append != nil {
		obj = gin.H{
			wrapKey: obj,
			"meta":  r.append,
		}
	}
	response.OK(r.ctx, obj)
}

// When determines if the given condition is true and executes the given callbacks.
func When(ok bool, fn interface{}) interface{} {
	if !ok {
		return &MissingValue{}
	}

	if f, ok := fn.(func() interface{}); ok {
		return f()
	}
	return fn
}
