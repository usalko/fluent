// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"strings"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/device"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/schema"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/session"
)

// Session is the model entity for the Session schema.
type Session struct {
	config
	// ID of the fluent.
	ID schema.ID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges           SessionEdges `json:"edges"`
	device_sessions *schema.ID
	selectValues    sql.SelectValues
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) DeviceOrErr() (*Device, error) {
	if e.Device != nil {
		return e.Device, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "device"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			values[i] = new(schema.ID)
		case session.ForeignKeys[0]: // device_sessions
			values[i] = &sql.NullScanner{S: new(schema.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			if value, ok := values[i].(*schema.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field device_sessions", values[i])
			} else if value.Valid {
				s.device_sessions = new(schema.ID)
				*s.device_sessions = *value.S.(*schema.ID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Session.
// This includes values selected through modifiers, order, etc.
func (s *Session) Value(name string) (fluent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryDevice queries the "device" edge of the Session entity.
func (s *Session) QueryDevice() *DeviceQuery {
	return NewSessionClient(s.config).QueryDevice(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return NewSessionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("fluent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session
