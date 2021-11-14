package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	gen "github.com/fitv/min/ent"
	"github.com/fitv/min/ent/hook"
	"github.com/fitv/min/util/hash"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		CharsetMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					if password, ok := m.Password(); ok && len(password) <= 20 {
						hashPassword, err := hash.Make(password)
						if err != nil {
							return nil, fmt.Errorf("failed to hash password: %w", err)
						}
						m.SetPassword(string(hashPassword))
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
