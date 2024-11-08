// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	fluent.Schema
}

// Fields of the Card.
func (Card) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("number"),
	}
}
