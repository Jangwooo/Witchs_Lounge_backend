package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Music holds the schema definition for the Music entity.
type Music struct {
	ent.Schema
}

// Fields of the Music.
func (Music) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Text("name"),
		field.Text("music_source"),
		field.Text("jacket_source"),
		field.Float("duration"),
		field.Text("Author"),
	}
}

// Edges of the Music.
func (Music) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stages", Stage.Type),
		edge.To("records", Record.Type),
	}
}
