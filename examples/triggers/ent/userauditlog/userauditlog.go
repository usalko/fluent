// Code generated by ent, DO NOT EDIT.

package userauditlog

import (
	"github.com/usalko/fluent/dialect/sql"
)

const (
	// Label holds the string label denoting the userauditlog type in the database.
	Label = "user_audit_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOperationType holds the string denoting the operation_type field in the database.
	FieldOperationType = "operation_type"
	// FieldOperationTime holds the string denoting the operation_time field in the database.
	FieldOperationTime = "operation_time"
	// FieldOldValue holds the string denoting the old_value field in the database.
	FieldOldValue = "old_value"
	// FieldNewValue holds the string denoting the new_value field in the database.
	FieldNewValue = "new_value"
	// Table holds the table name of the userauditlog in the database.
	Table = "user_audit_logs"
)

// Columns holds all SQL columns for userauditlog fields.
var Columns = []string{
	FieldID,
	FieldOperationType,
	FieldOperationTime,
	FieldOldValue,
	FieldNewValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the UserAuditLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOperationType orders the results by the operation_type field.
func ByOperationType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperationType, opts...).ToFunc()
}

// ByOperationTime orders the results by the operation_time field.
func ByOperationTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperationTime, opts...).ToFunc()
}

// ByOldValue orders the results by the old_value field.
func ByOldValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOldValue, opts...).ToFunc()
}

// ByNewValue orders the results by the new_value field.
func ByNewValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNewValue, opts...).ToFunc()
}
