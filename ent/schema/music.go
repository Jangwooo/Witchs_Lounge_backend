package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Music holds the schema definition for the Music entity.
type Music struct {
	ent.Schema
}

// Mixin of the Music.
func (Music) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Music.
func (Music) Fields() []ent.Field {
	return []ent.Field{
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
