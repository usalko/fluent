// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/flc/integration/custom_id/sid"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// IntSID holds the schema definition for the IntSID entity.
type IntSID struct {
	fluent.Schema
}

// Fields of the IntSid.
func (IntSID) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int64("id").
			GoType(sid.New()).
			Unique().
			Immutable(),
	}
}

// Edges of the IntSid.
func (IntSID) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("parent", IntSID.Type).
			Unique(),
		edge.From("children", IntSID.Type).
			Ref("parent"),
	}
}
