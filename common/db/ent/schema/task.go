package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.Int("total_status"),
		field.Time("created_at").Default(time.Now()),
		field.Time("modified_at").Default(time.Now()),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subtask", SubTask.Type).Unique(),
		edge.To("success_point", Point.Type).Unique(),
		edge.From("user", Member.Type).Ref("tasks"),
	}
}
