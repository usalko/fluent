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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/usalko/fluent/fluent_gql/internal/todofed/fluent/category"
	"github.com/usalko/fluent/fluent_gql/internal/todofed/fluent/schema/schematype"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the fluent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Status holds the value of the "status" field.
	Status category.Status `json:"status,omitempty"`
	// Config holds the value of the "config" field.
	Config *schematype.CategoryConfig `json:"config,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
	// Count holds the value of the "count" field.
	Count uint64 `json:"count,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges        CategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo `json:"todos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedTodos map[string][]*Todo
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldStrings:
			values[i] = new([]byte)
		case category.FieldConfig:
			values[i] = new(schematype.CategoryConfig)
		case category.FieldID, category.FieldDuration, category.FieldCount:
			values[i] = new(sql.NullInt64)
		case category.FieldText, category.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case category.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case category.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = category.Status(value.String)
			}
		case category.FieldConfig:
			if value, ok := values[i].(*schematype.CategoryConfig); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil {
				c.Config = value
			}
		case category.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				c.Duration = time.Duration(value.Int64)
			}
		case category.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				c.Count = uint64(value.Int64)
			}
		case category.FieldStrings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field strings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Strings); err != nil {
					return fmt.Errorf("unmarshal field strings: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Category.
// This includes values selected through modifiers, order, etc.
func (c *Category) Value(name string) (fluent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryTodos queries the "todos" edge of the Category entity.
func (c *Category) QueryTodos() *TodoQuery {
	return NewCategoryClient(c.config).QueryTodos(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return NewCategoryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("fluent: Category is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("text=")
	builder.WriteString(c.Text)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", c.Config))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", c.Duration))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", c.Count))
	builder.WriteString(", ")
	builder.WriteString("strings=")
	builder.WriteString(fmt.Sprintf("%v", c.Strings))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (c Category) IsEntity() {}

// NamedTodos returns the Todos named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Category) NamedTodos(name string) ([]*Todo, error) {
	if c.Edges.namedTodos == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedTodos[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Category) appendNamedTodos(name string, edges ...*Todo) {
	if c.Edges.namedTodos == nil {
		c.Edges.namedTodos = make(map[string][]*Todo)
	}
	if len(edges) == 0 {
		c.Edges.namedTodos[name] = []*Todo{}
	} else {
		c.Edges.namedTodos[name] = append(c.Edges.namedTodos[name], edges...)
	}
}

// Categories is a parsable slice of Category.
type Categories []*Category
