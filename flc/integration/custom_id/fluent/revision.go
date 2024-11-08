// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"strings"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/revision"
)

// Revision is the model entity for the Revision schema.
type Revision struct {
	config
	// ID of the fluent.
	ID           string `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Revision) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Revision fields.
func (r *Revision) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case revision.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Revision.
// This includes values selected through modifiers, order, etc.
func (r *Revision) Value(name string) (fluent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Revision.
// Note that you need to call Revision.Unwrap() before calling this method if this Revision
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Revision) Update() *RevisionUpdateOne {
	return NewRevisionClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Revision entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Revision) Unwrap() *Revision {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Revision is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Revision) String() string {
	var builder strings.Builder
	builder.WriteString("Revision(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Revisions is a parsable slice of Revision.
type Revisions []*Revision
