// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package valid

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// User holds the user schema.
type User struct {
	fluent.Schema
}

func (User) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Group holds the group schema.
type Group struct {
	fluent.Schema
}

func (Group) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("expired_at"),
		field.String("organization"),
	}
}

// Tag holds the tag schema.
type Tag struct {
	fluent.Schema
}

func (Tag) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("text"),
	}
}
