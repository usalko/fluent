// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "github.com/usalko/fluent"

// Builder holds the schema definition for the Builder entity.
type Builder struct {
	fluent.Schema
}

// Fields of the Builder.
func (Builder) Fields() []fluent.Field {
	return nil
}

// Edges of the Builder.
func (Builder) Edges() []fluent.Edge {
	return nil
}
