// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluenttask

import (
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/fluent/schema/task"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldPriorities holds the string denoting the priorities field in the database.
	FieldPriorities = "priorities"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldOrderOption holds the string denoting the order_option field in the database.
	FieldOrderOption = "order_option"
	// FieldOp holds the string denoting the op field in the database.
	FieldOp = "op"
	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldPriority,
	FieldPriorities,
	FieldCreatedAt,
	FieldOwner,
	FieldOrder,
	FieldOrderOption,
	FieldOp,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for _, f := range [...]string{FieldName} {
		if column == f {
			return true
		}
	}
	return false
}

var (
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority task.Priority
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultOp holds the default value on creation for the "op" field.
	DefaultOp string
	// OpValidator is a validator for the "op" field. It is called by the builders before save.
	OpValidator func(string) error
)

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOwner orders the results by the owner field.
func ByOwner(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwner, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByOrderOption orders the results by the order_option field.
func ByOrderOption(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderOption, opts...).ToFunc()
}

// ByOp orders the results by the op field.
func ByOp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOp, opts...).ToFunc()
}

// comment from another template.
