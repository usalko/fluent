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

package verysecret

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql"
)

const (
	// Label holds the string label denoting the verysecret type in the database.
	Label = "very_secret"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// Table holds the table name of the verysecret in the database.
	Table = "very_secrets"
)

// Columns holds all SQL columns for verysecret fields.
var Columns = []string{
	FieldID,
	FieldPassword,
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

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the VerySecret queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}
