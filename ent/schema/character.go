package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
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
