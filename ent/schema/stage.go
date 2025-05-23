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
		field.UUID("music_id", uuid.UUID{}).
			Comment("음악 ID"),
		field.Text("level_name").NotEmpty().
			Comment("난이도 이름 (Easy, Normal, Hard, Expert)"),
		field.Int("difficulty").
			Comment("난이도 수치 (1-10)"),
		field.Text("level_address").NotEmpty().
			Comment("채보 파일 경로"),
		field.Text("jacket_address").NotEmpty().
			Comment("난이도별 재킷 이미지 경로"),
		field.Int("total_notes").
			Comment("총 노트 수"),
		field.Int("max_combo").
			Comment("최대 콤보"),
		field.Bool("is_active").Default(true).
			Comment("활성 여부"),
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
