// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"

	"github.com/google/uuid"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	fluent.Schema
}

// Fields of the Item.
func (Item) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			DefaultFunc(uuid.NewString).
			MaxLen(64),
		field.String("text").
			MaxLen(128).
			Unique().
			Optional(),
	}
}
