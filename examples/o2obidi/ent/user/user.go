// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeSpouse holds the string denoting the spouse edge name in mutations.
	EdgeSpouse = "spouse"
	// Table holds the table name of the user in the database.
	Table = "users"
	// SpouseTable is the table that holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "user_spouse"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_spouse",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySpouseField orders the results by spouse field.
func BySpouseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSpouseStep(), sql.OrderByField(field, opts...))
	}
}
func newSpouseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
	)
}
