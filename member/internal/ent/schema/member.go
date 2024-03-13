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
		field.String("email"),
		field.String("username"),
		field.String("password"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
	//return []ent.Edge{
	//	edge.To("user_id", Task.Type),
	//}
}
