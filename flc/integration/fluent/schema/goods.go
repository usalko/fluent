// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "github.com/usalko/fluent"

// Goods holds the schema definition for the Goods entity.
type Goods struct {
	fluent.Schema
}

// Fields of the Goods.
func (Goods) Fields() []fluent.Field {
	return nil
}

// Edges of the Goods.
func (Goods) Edges() []fluent.Edge {
	return nil
}
