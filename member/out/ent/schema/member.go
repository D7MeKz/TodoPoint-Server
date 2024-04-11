package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("email").Unique().NotEmpty(),
		field.String("username").Optional().NotEmpty(),
		field.String("password").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}