package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PointInfo holds the schema definition for the PointInfo entity.
type PointInfo struct {
	ent.Schema
}

// Fields of the PointInfo.
func (PointInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int64("total_points"),
		field.Time("modified_at").Default(time.Now()),
	}
}

// Edges of the PointInfo.
func (PointInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("points", Point.Type),
		// PointDetails Only one
		edge.From("user_id", Member.Type).Ref("point_info").Unique(),
	}
}
