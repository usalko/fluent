// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package buildflags

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
