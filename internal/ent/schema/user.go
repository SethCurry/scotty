package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("discord_id").Unique(),
		field.Int("ranked_score").Default(0),
		field.String("finals_id").Optional().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
