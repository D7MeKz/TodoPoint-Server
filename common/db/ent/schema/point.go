package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Point holds the schema definition for the Point entity.
type Point struct {
	ent.Schema
}

// Fields of the Point.
func (Point) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("point"),
		field.String("type"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Point.
func (Point) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subtask", SubTask.Type).Ref("point").Unique(),
		edge.From("task", Task.Type).Ref("point").Unique(),
		edge.From("point_info", PointInfo.Type).Ref("points").Unique(),
	}
}
