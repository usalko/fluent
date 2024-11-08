// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
		{Name: "t", Type: field.TypeJSON, Nullable: true},
		{Name: "url", Type: field.TypeJSON, Nullable: true},
		{Name: "urls", Type: field.TypeJSON, Nullable: true},
		{Name: "raw", Type: field.TypeJSON, Nullable: true},
		{Name: "dirs", Type: field.TypeJSON},
		{Name: "ints", Type: field.TypeJSON, Nullable: true},
		{Name: "floats", Type: field.TypeJSON, Nullable: true},
		{Name: "strings", Type: field.TypeJSON, Nullable: true},
		{Name: "ints_validate", Type: field.TypeJSON, Nullable: true},
		{Name: "floats_validate", Type: field.TypeJSON, Nullable: true},
		{Name: "strings_validate", Type: field.TypeJSON, Nullable: true},
		{Name: "addr", Type: field.TypeJSON, Nullable: true},
		{Name: "unknown", Type: field.TypeJSON, Nullable: true},
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
