// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// RelationshipInfo holds the schema definition for the RelationshipInfo entity.
type RelationshipInfo struct {
	fluent.Schema
}

// Fields of the RelationshipInfo.
func (RelationshipInfo) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("text"),
	}
}
