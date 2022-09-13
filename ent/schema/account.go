package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner").
			NotEmpty().
			Immutable(),
		field.Int64("balance").
			Default(0),
		field.String("currency").
			NotEmpty().
			Default("EUR"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("username", User.Type).
			Ref("accounts").
			Unique().
			Required(),
	}
}
