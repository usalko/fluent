// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
)

type Spec struct {
	fluent.Schema
}

// Edges of the Spec.
func (Spec) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("card", Card.Type),
	}
}
