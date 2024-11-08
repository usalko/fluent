// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// base schema for sharing fields and edges.
type base struct {
	fluent.Schema
}

func (base) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("base_field"),
	}
}

// User holds the user schema.
type User struct {
	base
}

func (u User) Fields() []fluent.Field {
	return append(
		u.base.Fields(),
		field.String("user_field"),
	)
}
