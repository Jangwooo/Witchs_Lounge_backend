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
		field.Text("name").
			Comment("곡 제목"),
		field.Text("artist").
			Comment("아티스트"),
		field.Text("composer").Optional().
			Comment("작곡가"),
		field.Text("music_source").
			Comment("음악 파일 경로"),
		field.Text("jacket_source").
			Comment("재킷 이미지 경로"),
		field.Float("duration").
			Comment("곡 길이(초)"),
		field.Float("bpm").
			Comment("BPM"),
		field.Text("genre").Optional().
			Comment("장르"),
		field.Text("description").Optional().
			Comment("곡 설명"),
		field.Bool("is_featured").Default(false).
			Comment("추천곡 여부"),
		field.Bool("is_free").Default(true).
			Comment("무료곡 여부"),
		field.Int("unlock_level").Default(1).
			Comment("해금 레벨"),
		field.Time("release_date").Optional().Nillable().
			Comment("출시일"),
		field.Bool("is_active").Default(true).
			Comment("활성 여부"),
	}
}

// Edges of the Music.
func (Music) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stages", Stage.Type),
		edge.To("records", Record.Type),
	}
}
