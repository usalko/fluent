// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package zoo

import (
	"github.com/usalko/fluent/dialect/sql"
)

const (
	// Label holds the string label denoting the zoo type in the database.
	Label = "zoo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the zoo in the database.
	Table = "zoos"
)

// Columns holds all SQL columns for zoo fields.
var Columns = []string{
	FieldID,
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

// OrderOption defines the ordering options for the Zoo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
