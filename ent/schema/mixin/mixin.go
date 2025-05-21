package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// GlobalSchemaMixin adds common fields like custom ID.
type GlobalMixin struct {
	mixin.Schema
}

// Fields of the GlobalSchemaMixin.
func (GlobalMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique().
			Comment("Global custom UUID ID"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Created time"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Updated time"),
	}
}

// Mixin 적용
func (GlobalMixin) Mixin() []ent.Mixin {
	return nil
}
