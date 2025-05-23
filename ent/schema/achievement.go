package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

// Mixin of the Achievement.
func (Achievement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Comment("업적 이름"),
		field.Text("description").
			Comment("업적 설명"),
		field.Text("icon_url").Optional().
			Comment("업적 아이콘 URL"),
		field.Enum("type").Values("score", "combo", "accuracy", "play_count", "special").
			Comment("업적 타입"),
		field.JSON("conditions", map[string]interface{}{}).
			Comment("달성 조건"),
		field.JSON("rewards", map[string]interface{}{}).
			Comment("보상"),
		field.Int("points").Default(0).
			Comment("업적 포인트"),
		field.Bool("is_hidden").Default(false).
			Comment("숨김 업적 여부"),
		field.Bool("is_active").Default(true).
			Comment("활성 여부"),
	}
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_achievements", UserAchievement.Type),
	}
}

// Indexes of the Achievement.
func (Achievement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("is_active"),
		index.Fields("is_hidden"),
	}
}
