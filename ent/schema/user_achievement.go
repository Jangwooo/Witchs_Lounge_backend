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

// UserAchievement holds the schema definition for the UserAchievement entity.
type UserAchievement struct {
	ent.Schema
}

// Mixin of the UserAchievement.
func (UserAchievement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the UserAchievement.
func (UserAchievement) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Comment("유저 ID"),
		field.UUID("achievement_id", uuid.UUID{}).
			Comment("업적 ID"),
		field.Time("unlocked_at").Default(time.Now).
			Comment("업적 달성 시간"),
		field.JSON("progress_data", map[string]interface{}{}).Optional().
			Comment("진행도 데이터"),
	}
}

// Edges of the UserAchievement.
func (UserAchievement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_achievements").
			Field("user_id").
			Unique().
			Required(),
		edge.From("achievement", Achievement.Type).
			Ref("user_achievements").
			Field("achievement_id").
			Unique().
			Required(),
	}
}

// Indexes of the UserAchievement.
func (UserAchievement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "achievement_id").Unique(),
		index.Fields("unlocked_at"),
	}
}
