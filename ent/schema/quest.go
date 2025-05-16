package schema

import "entgo.io/ent"

// Quest holds the schema definition for the Quest entity.
type Quest struct {
	ent.Schema
}

// Fields of the Quest.
func (Quest) Fields() []ent.Field {
	return nil
}

// Edges of the Quest.
func (Quest) Edges() []ent.Edge {
	return nil
}
