package schema

import "entgo.io/ent"

// Transfer holds the schema definition for the Transfer entity.
type Transfer struct {
	ent.Schema
}

// Fields of the Transfer.
func (Transfer) Fields() []ent.Field {
	return nil
}

// Edges of the Transfer.
func (Transfer) Edges() []ent.Edge {
	return nil
}
