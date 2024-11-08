// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

// User holds the schema definition for the User entity.
import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/field"
)

// Revision holds the schema definition for the Revision entity.
type Revision struct {
	fluent.Schema
}

// Fields of the Revision.
func (Revision) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id"),
	}
}
