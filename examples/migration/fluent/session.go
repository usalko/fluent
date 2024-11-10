// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/examples/migration/fluent/session"
	"github.com/usalko/fluent/examples/migration/fluent/session_device"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the fluent.
	ID uuid.UUID `json:"id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// IssuedAt holds the value of the "issued_at" field.
	IssuedAt time.Time `json:"issued_at,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Method holds the value of the "method" field.
	Method map[string]interface{} `json:"method,omitempty"`
	// DeviceID holds the value of the "device_id" field.
	DeviceID uuid.UUID `json:"device_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges        SessionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Device holds the value of the device edge.
	Device *SessionDevice `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) DeviceOrErr() (*SessionDevice, error) {
	if e.Device != nil {
		return e.Device, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: session_device.Label}
	}
	return nil, &NotLoadedError{edge: "device"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldMethod:
			values[i] = new([]byte)
		case session.FieldActive:
			values[i] = new(sql.NullBool)
		case session.FieldToken:
			values[i] = new(sql.NullString)
		case session.FieldIssuedAt, session.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case session.FieldID, session.FieldDeviceID:
			values[i] = new(uuid.UUID)
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
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case session.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				s.Active = value.Bool
			}
		case session.FieldIssuedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field issued_at", values[i])
			} else if value.Valid {
				s.IssuedAt = value.Time
			}
		case session.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				s.ExpiresAt = value.Time
			}
		case session.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				s.Token = value.String
			}
		case session.FieldMethod:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Method); err != nil {
					return fmt.Errorf("unmarshal field method: %w", err)
				}
			}
		case session.FieldDeviceID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value != nil {
				s.DeviceID = *value
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
func (s *Session) QueryDevice() *SessionDeviceQuery {
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
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", s.Active))
	builder.WriteString(", ")
	builder.WriteString("issued_at=")
	builder.WriteString(s.IssuedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(s.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(s.Token)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(fmt.Sprintf("%v", s.Method))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", s.DeviceID))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session
