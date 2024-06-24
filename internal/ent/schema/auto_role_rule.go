package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AutoRoleRule struct {
	ent.Schema
}

func (AutoRoleRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id").Unique(),
	}
}

func (AutoRoleRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("guild", Guild.Type).Unique(),
	}
}
