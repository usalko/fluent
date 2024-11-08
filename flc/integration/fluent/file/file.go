// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package file

import (
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSetID holds the string denoting the set_id field in the database.
	FieldSetID = "set_id"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "fsize"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldOp holds the string denoting the op field in the database.
	FieldOp = "op"
	// FieldFieldID holds the string denoting the field_id field in the database.
	FieldFieldID = "field_id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeField holds the string denoting the field edge name in mutations.
	EdgeField = "field"
	// Table holds the table name of the file in the database.
	Table = "files"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "files"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_files"
	// TypeTable is the table that holds the type relation/edge.
	TypeTable = "files"
	// TypeInverseTable is the table name for the FileType entity.
	// It exists in this package in order to avoid circular dependency with the "filetype" package.
	TypeInverseTable = "file_types"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "file_type_files"
	// FieldTable is the table that holds the field relation/edge.
	FieldTable = "field_types"
	// FieldInverseTable is the table name for the FieldType entity.
	// It exists in this package in order to avoid circular dependency with the "fieldtype" package.
	FieldInverseTable = "field_types"
	// FieldColumn is the table column denoting the field relation/edge.
	FieldColumn = "file_field"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldSetID,
	FieldSize,
	FieldName,
	FieldUser,
	FieldGroup,
	FieldOp,
	FieldFieldID,
	FieldCreateTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"file_type_files",
	"group_files",
	"user_files",
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

var (
	// SetIDValidator is a validator for the "set_id" field. It is called by the builders before save.
	SetIDValidator func(int) error
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize int
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int) error
)

// OrderOption defines the ordering options for the File queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySetID orders the results by the set_id field.
func BySetID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSetID, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUser orders the results by the user field.
func ByUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUser, opts...).ToFunc()
}

// ByGroup orders the results by the group field.
func ByGroup(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroup, opts...).ToFunc()
}

// ByOp orders the results by the op field.
func ByOp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOp, opts...).ToFunc()
}

// ByFieldID orders the results by the field_id field.
func ByFieldID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByTypeField orders the results by type field.
func ByTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByFieldCount orders the results by field count.
func ByFieldCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFieldStep(), opts...)
	}
}

// ByField orders the results by field terms.
func ByField(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFieldStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TypeTable, TypeColumn),
	)
}
func newFieldStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FieldInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FieldTable, FieldColumn),
	)
}

// comment from another template.
