package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Stage holds the schema definition for the Stage entity.
type Stage struct {
	ent.Schema
}

// Mixin of the Stage.
func (Stage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Stage.
func (Stage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("music_id", uuid.UUID{}),
		field.Text("level_name").NotEmpty(),
		field.Text("level_address").NotEmpty(),
		field.Text("jacket_address").NotEmpty(),
	}
}

// Edges of the Stage.
func (Stage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("music", Music.Type).
			Ref("stages").
			Field("music_id").
			Required().
			Unique(),
		edge.To("records", Record.Type),
	}
}
