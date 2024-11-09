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

package fluent

import (
	"fmt"
	"strings"

	"github.com/usalko/fluent/fluent_gql/internal/todo/fluent/workspace"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
)

// Workspace is the model entity for the Workspace schema.
type Workspace struct {
	config `json:"-"`
	// ID of the fluent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workspace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workspace.FieldID:
			values[i] = new(sql.NullInt64)
		case workspace.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workspace fields.
func (w *Workspace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workspace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case workspace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Workspace.
// This includes values selected through modifiers, order, etc.
func (w *Workspace) Value(name string) (fluent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Workspace.
// Note that you need to call Workspace.Unwrap() before calling this method if this Workspace
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workspace) Update() *WorkspaceUpdateOne {
	return NewWorkspaceClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workspace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workspace) Unwrap() *Workspace {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("fluent: Workspace is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workspace) String() string {
	var builder strings.Builder
	builder.WriteString("Workspace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Workspaces is a parsable slice of Workspace.
type Workspaces []*Workspace
