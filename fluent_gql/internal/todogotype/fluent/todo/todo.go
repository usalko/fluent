// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by flc, DO NOT EDIT.

package todo

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldInit holds the string denoting the init field in the database.
	FieldInit = "init"
	// FieldCustom holds the string denoting the custom field in the database.
	FieldCustom = "custom"
	// FieldCustomp holds the string denoting the customp field in the database.
	FieldCustomp = "customp"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeSecret holds the string denoting the secret edge name in mutations.
	EdgeSecret = "secret"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "todos"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "todo_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "todos"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "todo_children"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "todos"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_id"
	// SecretTable is the table that holds the secret relation/edge.
	SecretTable = "todos"
	// SecretInverseTable is the table name for the VerySecret entity.
	// It exists in this package in order to avoid circular dependency with the "verysecret" package.
	SecretInverseTable = "very_secrets"
	// SecretColumn is the table column denoting the secret relation/edge.
	SecretColumn = "todo_secret"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldStatus,
	FieldPriority,
	FieldText,
	FieldBlob,
	FieldInit,
	FieldCustom,
	FieldCustomp,
	FieldValue,
	FieldCategoryID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "todos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"todo_children",
	"todo_secret",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID string
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusPending    Status = "PENDING"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusInProgress, StatusCompleted, StatusPending:
		return nil
	default:
		return fmt.Errorf("todo: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Todo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByCategoryID orders the results by the category_id field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoryField orders the results by category field.
func ByCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// BySecretField orders the results by secret field.
func BySecretField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSecretStep(), sql.OrderByField(field, opts...))
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
	)
}
func newSecretStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SecretInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SecretTable, SecretColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
