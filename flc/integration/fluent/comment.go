// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/fluent/comment"
	schemadir "github.com/usalko/fluent/flc/integration/fluent/schema/dir"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the fluent.
	ID int `json:"id,omitempty"`
	// UniqueInt holds the value of the "unique_int" field.
	UniqueInt int `json:"unique_int,omitempty"`
	// UniqueFloat holds the value of the "unique_float" field.
	UniqueFloat float64 `json:"unique_float,omitempty"`
	// NillableInt holds the value of the "nillable_int" field.
	NillableInt *int `json:"nillable_int,omitempty"`
	// Table holds the value of the "table" field.
	Table string `json:"table,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir schemadir.Dir `json:"dir,omitempty"`
	// Client holds the value of the "client" field.
	Client       string `json:"client,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldDir:
			values[i] = new([]byte)
		case comment.FieldUniqueFloat:
			values[i] = new(sql.NullFloat64)
		case comment.FieldID, comment.FieldUniqueInt, comment.FieldNillableInt:
			values[i] = new(sql.NullInt64)
		case comment.FieldTable, comment.FieldClient:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldUniqueInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unique_int", values[i])
			} else if value.Valid {
				c.UniqueInt = int(value.Int64)
			}
		case comment.FieldUniqueFloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field unique_float", values[i])
			} else if value.Valid {
				c.UniqueFloat = value.Float64
			}
		case comment.FieldNillableInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int", values[i])
			} else if value.Valid {
				c.NillableInt = new(int)
				*c.NillableInt = int(value.Int64)
			}
		case comment.FieldTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table", values[i])
			} else if value.Valid {
				c.Table = value.String
			}
		case comment.FieldDir:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field dir", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Dir); err != nil {
					return fmt.Errorf("unmarshal field dir: %w", err)
				}
			}
		case comment.FieldClient:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client", values[i])
			} else if value.Valid {
				c.Client = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (fluent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("fluent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("unique_int=")
	builder.WriteString(fmt.Sprintf("%v", c.UniqueInt))
	builder.WriteString(", ")
	builder.WriteString("unique_float=")
	builder.WriteString(fmt.Sprintf("%v", c.UniqueFloat))
	builder.WriteString(", ")
	if v := c.NillableInt; v != nil {
		builder.WriteString("nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("table=")
	builder.WriteString(c.Table)
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(fmt.Sprintf("%v", c.Dir))
	builder.WriteString(", ")
	builder.WriteString("client=")
	builder.WriteString(c.Client)
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment
