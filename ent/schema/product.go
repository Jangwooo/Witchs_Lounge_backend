package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product는 상점에서 판매하는 상품 엔티티입니다.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Float("price").Positive(),
		field.Enum("type").Values("hat", "cane", "clothes", "character"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type).Unique().
			StorageKey(edge.Column("item_id")),
		edge.To("character", Character.Type).Unique().
			StorageKey(edge.Column("character_id")),
		edge.From("purchasers", User.Type).
			Ref("purchased_products").
			Through("user_purchases", UserPurchase.Type),
	}
}
