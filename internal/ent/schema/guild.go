package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Guild struct {
	ent.Schema
}

func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("guild_id").Unique(),
		field.String("welcome_template").Default(""),
		field.String("welcome_channel").Default(""),
	}
}

func (Guild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("auto_role_rules", AutoRoleRule.Type).Ref("guild"),
	}
}
