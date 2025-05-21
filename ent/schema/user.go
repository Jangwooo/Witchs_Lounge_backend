package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/witchs-lounge_backend/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.GlobalMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("nickname"),
		field.Text("steam_id").Immutable().Unique(),
		field.Text("steam_avatar_url").Optional(),
		field.Text("steam_default_language").Default("ko"),
		field.Time("last_login_at").Default(time.Now),

		field.Text("customize_data").Default("{}").Nillable().Optional(),
		field.Text("save_data").Default("{}").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Add edges here
		edge.To("purchased_products", Product.Type).
			Through("user_purchases", UserPurchase.Type),
		edge.To("records", Record.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("steam_id"),
	}
}
