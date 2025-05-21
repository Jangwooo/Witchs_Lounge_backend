package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Mixin of the Character.
func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
		field.Text("description").Optional(),
		field.Text("source"),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).
			Ref("character"),
		edge.To("records", Record.Type),
	}
}
