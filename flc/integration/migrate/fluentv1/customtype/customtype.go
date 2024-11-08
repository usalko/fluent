// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package customtype

import (
	"github.com/usalko/fluent/dialect/sql"
)

const (
	// Label holds the string label denoting the customtype type in the database.
	Label = "custom_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCustom holds the string denoting the custom field in the database.
	FieldCustom = "custom"
	// Table holds the table name of the customtype in the database.
	Table = "custom_types"
)

// Columns holds all SQL columns for customtype fields.
var Columns = []string{
	FieldID,
	FieldCustom,
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

// OrderOption defines the ordering options for the CustomType queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCustom orders the results by the custom field.
func ByCustom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustom, opts...).ToFunc()
}
