//go:build !hidegroups

// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package buildflags

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

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
