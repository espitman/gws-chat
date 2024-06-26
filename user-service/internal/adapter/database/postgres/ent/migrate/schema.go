// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatUsersColumns holds the columns for the "chat_users" table.
	ChatUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "status", Type: field.TypeString, Default: "available"},
		{Name: "is_online", Type: field.TypeBool, Default: false},
	}
	// ChatUsersTable holds the schema information for the "chat_users" table.
	ChatUsersTable = &schema.Table{
		Name:       "chat_users",
		Columns:    ChatUsersColumns,
		PrimaryKey: []*schema.Column{ChatUsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatUsersTable,
	}
)

func init() {
	ChatUsersTable.Annotation = &entsql.Annotation{
		Table: "chat_users",
	}
}
