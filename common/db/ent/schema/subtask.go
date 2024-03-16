package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// SubTask holds the schema definition for the SubTask entity.
type SubTask struct {
	ent.Schema
}

// Fields of the SubTask.
func (SubTask) Fields() []ent.Field {
	return nil
}

// Edges of the SubTask.
func (SubTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("point", Point.Type).Unique(),
		edge.From("task", Task.Type).Ref("subtask"),
	}
}
