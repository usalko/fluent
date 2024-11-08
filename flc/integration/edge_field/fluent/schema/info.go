// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"encoding/json"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// Info holds the schema definition for the Info entity.
type Info struct {
	fluent.Schema
}

func (Info) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id"),
		field.JSON("content", json.RawMessage{}),
	}
}

func (Info) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("user", User.Type).
			Unique().
			StorageKey(edge.Column("id")),
	}
}
