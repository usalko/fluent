// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// GroupInfo holds the schema for the group-info entity.
type GroupInfo struct {
	fluent.Schema
}

// Fields of the group.
func (GroupInfo) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("desc"),
		field.Int("max_users").
			Default(1e4),
	}
}

// Edges of the group.
func (GroupInfo) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("groups", Group.Type).
			Ref("info"),
	}
}
