// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
	"github.com/usalko/fluent/schema/index"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	fluent.Schema
}

// Fields of the Media.
func (Media) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("source").
			Optional(),
		field.String("source_uri").
			Optional().
			Comment("source_ui text").
			Annotations(fluent_sql.WithComments(false)),
		field.Text("text").
			Optional().
			Comment("media text"),
	}
}

// Indexes of the Media.
func (Media) Indexes() []fluent.Index {
	return []fluent.Index{
		index.Fields("source", "source_uri").
			Annotations(fluent_sql.PrefixColumn("source", 100)).
			Unique(),
		// MySQL allow indexing text column prefix.
		index.Fields("text").
			Annotations(fluent_sql.Prefix(100)),
	}
}

// Annotations of the Media.
func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("Comment that appears in both the schema and the generated code"),
		fluent_sql.WithComments(true),
		fluent_sql.Check("text <> 'boring'"),
		fluent_sql.Checks(map[string]string{
			"boring_check": "source_uri <> 'github.com/usalko/fluent'",
		}),
	}
}
