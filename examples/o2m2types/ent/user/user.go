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
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
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

// ByPetsCount orders the results by pets count.
func ByPetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPetsStep(), opts...)
	}
}

// ByPets orders the results by pets terms.
func ByPets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PetsTable, PetsColumn),
	)
}
