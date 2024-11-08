// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"strings"
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/fluent/license"
)

// License is the model entity for the License schema.
type License struct {
	config `json:"-"`
	// ID of the fluent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*License) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case license.FieldID:
			values[i] = new(sql.NullInt64)
		case license.FieldCreateTime, license.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the License fields.
func (l *License) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case license.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case license.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				l.CreateTime = value.Time
			}
		case license.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				l.UpdateTime = value.Time
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the License.
// This includes values selected through modifiers, order, etc.
func (l *License) Value(name string) (fluent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this License.
// Note that you need to call License.Unwrap() before calling this method if this License
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *License) Update() *LicenseUpdateOne {
	return NewLicenseClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the License entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *License) Unwrap() *License {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: License is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *License) String() string {
	var builder strings.Builder
	builder.WriteString("License(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("create_time=")
	builder.WriteString(l.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(l.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Licenses is a parsable slice of License.
type Licenses []*License