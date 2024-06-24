// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AutoRoleRulesColumns holds the columns for the "auto_role_rules" table.
	AutoRoleRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_id", Type: field.TypeString, Unique: true},
		{Name: "auto_role_rule_guild", Type: field.TypeInt, Nullable: true},
	}
	// AutoRoleRulesTable holds the schema information for the "auto_role_rules" table.
	AutoRoleRulesTable = &schema.Table{
		Name:       "auto_role_rules",
		Columns:    AutoRoleRulesColumns,
		PrimaryKey: []*schema.Column{AutoRoleRulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auto_role_rules_guilds_guild",
				Columns:    []*schema.Column{AutoRoleRulesColumns[2]},
				RefColumns: []*schema.Column{GuildsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GuildsColumns holds the columns for the "guilds" table.
	GuildsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "guild_id", Type: field.TypeString, Unique: true},
	}
	// GuildsTable holds the schema information for the "guilds" table.
	GuildsTable = &schema.Table{
		Name:       "guilds",
		Columns:    GuildsColumns,
		PrimaryKey: []*schema.Column{GuildsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "discord_id", Type: field.TypeString, Unique: true},
		{Name: "ranked_score", Type: field.TypeInt, Default: 0},
		{Name: "finals_id", Type: field.TypeString, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AutoRoleRulesTable,
		GuildsTable,
		UsersTable,
	}
)

func init() {
	AutoRoleRulesTable.ForeignKeys[0].RefTable = GuildsTable
}
