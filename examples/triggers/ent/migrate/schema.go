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
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserAuditLogsColumns holds the columns for the "user_audit_logs" table.
	UserAuditLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "operation_type", Type: field.TypeString},
		{Name: "operation_time", Type: field.TypeString},
		{Name: "old_value", Type: field.TypeString, Nullable: true},
		{Name: "new_value", Type: field.TypeString, Nullable: true},
	}
	// UserAuditLogsTable holds the schema information for the "user_audit_logs" table.
	UserAuditLogsTable = &schema.Table{
		Name:       "user_audit_logs",
		Columns:    UserAuditLogsColumns,
		PrimaryKey: []*schema.Column{UserAuditLogsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
		UserAuditLogsTable,
	}
)

func init() {
}
