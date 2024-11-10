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
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/fluent_gql/internal/todo_uuid/fluent/friendship"
	"github.com/usalko/fluent/fluent_gql/internal/todo_uuid/fluent/user"
)

// Friendship is the model entity for the Friendship schema.
type Friendship struct {
	config `json:"-"`
	// ID of the fluent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// FriendID holds the value of the "friend_id" field.
	FriendID uuid.UUID `json:"friend_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FriendshipQuery when eager-loading is set.
	Edges        FriendshipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FriendshipEdges holds the relations/edges for other nodes in the graph.
type FriendshipEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Friend holds the value of the friend edge.
	Friend *User `json:"friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// FriendOrErr returns the Friend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) FriendOrErr() (*User, error) {
	if e.Friend != nil {
		return e.Friend, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Friendship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friendship.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case friendship.FieldID, friendship.FieldUserID, friendship.FieldFriendID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Friendship fields.
func (f *Friendship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friendship.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case friendship.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case friendship.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				f.UserID = *value
			}
		case friendship.FieldFriendID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field friend_id", values[i])
			} else if value != nil {
				f.FriendID = *value
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the fluent.Value that was dynamically selected and assigned to the Friendship.
// This includes values selected through modifiers, order, etc.
func (f *Friendship) Value(name string) (fluent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Friendship entity.
func (f *Friendship) QueryUser() *UserQuery {
	return NewFriendshipClient(f.config).QueryUser(f)
}

// QueryFriend queries the "friend" edge of the Friendship entity.
func (f *Friendship) QueryFriend() *UserQuery {
	return NewFriendshipClient(f.config).QueryFriend(f)
}

// Update returns a builder for updating this Friendship.
// Note that you need to call Friendship.Unwrap() before calling this method if this Friendship
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Friendship) Update() *FriendshipUpdateOne {
	return NewFriendshipClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Friendship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Friendship) Unwrap() *Friendship {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("fluent: Friendship is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Friendship) String() string {
	var builder strings.Builder
	builder.WriteString("Friendship(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", f.UserID))
	builder.WriteString(", ")
	builder.WriteString("friend_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FriendID))
	builder.WriteByte(')')
	return builder.String()
}

// Friendships is a parsable slice of Friendship.
type Friendships []*Friendship
