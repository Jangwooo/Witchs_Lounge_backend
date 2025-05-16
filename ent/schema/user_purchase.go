package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserPurchase는 유저와 상품 간의 N:M 관계를 위한 중간 테이블입니다.
type UserPurchase struct {
	ent.Schema
}

// Fields of the UserPurchase.
func (UserPurchase) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("product_id", uuid.UUID{}),
		field.Time("purchase_date").
			Default(time.Now),
		// 구매 관련 추가 필드들...
	}
}

// Edges of the UserPurchase.
func (UserPurchase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("product", Product.Type).
			Field("product_id").
			Unique().
			Required(),
	}
}
