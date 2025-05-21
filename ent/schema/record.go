package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Record holds the schema definition for the Record entity.
type Record struct {
	ent.Schema
}

// Mixin of the Record.
func (Record) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Record.
func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("music_id", uuid.UUID{}),
		field.UUID("stage_id", uuid.UUID{}),
		field.UUID("character_id", uuid.UUID{}),
		field.Int("score"),
		field.Int("perfect_count").Default(0),
		field.Int("good_count").Default(0),
		field.Int("bad_count").Default(0),
		field.Int("miss_count").Default(0),
		field.Time("played_at").Default(time.Now),
		field.Float("accuracy").Default(0),
		field.Text("additional_info").Optional(),
	}
}

// Edges of the Record.
func (Record) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("records").
			Field("user_id").
			Unique().
			Required(),
		edge.From("music", Music.Type).
			Ref("records").
			Field("music_id").
			Unique().
			Required(),
		edge.From("stage", Stage.Type).
			Ref("records").
			Field("stage_id").
			Unique().
			Required(),
		edge.From("character", Character.Type).
			Ref("records").
			Field("character_id").
			Unique().
			Required(),
	}
}

// Indexes of the Record.
func (Record) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "music_id", "stage_id"),
		index.Fields("user_id"),
		index.Fields("music_id"),
		index.Fields("stage_id"),
		index.Fields("played_at"),
	}
}
