// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"github.com/usalko/fluent/dialect/sql/schema"
	"github.com/usalko/fluent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "location", Type: field.TypeBytes, SchemaType: map[string]string{"postgres": "GEOMETRY(Point, 4326)"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}