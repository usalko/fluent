// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	uuidc "github.com/usalko/fluent/flc/integration/custom_id/uuidcompatible"
	"github.com/usalko/fluent/schema/field"
)

type LinkInformation struct {
	Name string
	Link string
}

// Link holds the schema definition for the Link entity.
type Link struct {
	fluent.Schema
}

// Fields of the IntSid.
func (Link) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuidc.UUIDC{}).Default(uuidc.NewUUIDC),
		field.JSON("link_information", map[string]LinkInformation{}).
			Default(map[string]LinkInformation{
				"ent": {
					Name: "ent",
					Link: "https://github.com/usalko/fluent/",
				},
			}),
	}
}
