package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DataSet holds the schema definition for the DataSet entity.
type EventLogArg struct {
	ent.Schema
}

// Fields of the DataSet.
func (EventLogArg) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
		field.String("value"),
	}
}

// Edges of the DataSet.
func (EventLogArg) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: Make these required on next migrate
		edge.From("eventLog", EventLog.Type).Ref("args").
			Unique().Required(),
	}
}
