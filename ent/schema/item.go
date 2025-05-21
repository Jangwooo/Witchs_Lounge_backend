package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Mixin of the Item.
func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
		field.Text("description").Optional(),
		field.Text("effect_id").Optional().Nillable(),
		field.Enum("type").Values("hat", "cane", "clothes"),
		field.Text("source"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).
			Ref("item"),
	}
}
