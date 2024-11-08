// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"fmt"
	"strings"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/examples/viewcomposite/fluent/cleanuser"
)

// CleanUser is the model entity for the CleanUser schema.
type CleanUser struct {
	config `json:"-"`
	// ID holds the value of the "id" field.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PublicInfo holds the value of the "public_info" field.
	PublicInfo   string `json:"public_info,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CleanUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cleanuser.FieldID:
			values[i] = new(sql.NullInt64)
		case cleanuser.FieldName, cleanuser.FieldPublicInfo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CleanUser fields.
func (cu *CleanUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cleanuser.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				cu.ID = int(value.Int64)
			}
		case cleanuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cu.Name = value.String
			}
		case cleanuser.FieldPublicInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_info", values[i])
			} else if value.Valid {
				cu.PublicInfo = value.String
			}
		default:
			cu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the CleanUser.
// This includes values selected through modifiers, order, etc.
func (cu *CleanUser) Value(name string) (fluent.Value, error) {
	return cu.selectValues.Get(name)
}

// Unwrap unwraps the CleanUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cu *CleanUser) Unwrap() *CleanUser {
	_tx, ok := cu.config.driver.(*txDriver)
	if !ok {
		panic("ent: CleanUser is not a transactional entity")
	}
	cu.config.driver = _tx.drv
	return cu
}

// String implements the fmt.Stringer.
func (cu *CleanUser) String() string {
	var builder strings.Builder
	builder.WriteString("CleanUser(")
	builder.WriteString("id=")
	builder.WriteString(fmt.Sprintf("%v", cu.ID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cu.Name)
	builder.WriteString(", ")
	builder.WriteString("public_info=")
	builder.WriteString(cu.PublicInfo)
	builder.WriteByte(')')
	return builder.String()
}

// CleanUsers is a parsable slice of CleanUser.
type CleanUsers []*CleanUser
