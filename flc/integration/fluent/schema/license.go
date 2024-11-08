// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/mixin"

	"ariga.io/atlas/sql/postgres"
)

// License holds the schema definition for the License entity.
type License struct {
	fluent.Schema
}

func (License) Mixin() []fluent.Mixin {
	return []fluent.Mixin{
		mixin.Time{},
	}
}

// Fields of the License.
func (License) Fields() []fluent.Field {
	return []fluent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeBigSerial,
			}),
	}
}
