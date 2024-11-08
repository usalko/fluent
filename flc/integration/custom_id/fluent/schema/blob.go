// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent/schema"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/fluent_sql"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"

	"github.com/google/uuid"
)

// Blob holds the schema definition for the Blob entity.
type Blob struct {
	fluent.Schema
}

// Fields of the Blob.
func (Blob) Fields() []fluent.Field {
	return []fluent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(fluent_sql.Annotation{
				Default: "uuid_generate_v4()",
			}).
			Unique(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int("count").
			Default(0),
	}
}

// Edges of the Blob.
func (Blob) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("parent", Blob.Type).
			Unique(),
		edge.To("links", Blob.Type).
			Through("blob_links", BlobLink.Type),
	}
}

// BlobLink holds the edge schema definition for blob links.
type BlobLink struct {
	fluent.Schema
}

// Annotations of the BlobLink.
func (BlobLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("blob_id", "link_id"),
	}
}

// Fields of the BlobLink.
func (BlobLink) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.UUID("blob_id", uuid.UUID{}),
		field.UUID("link_id", uuid.UUID{}),
	}
}

// Edges of the BlobLink.
func (BlobLink) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("blob", Blob.Type).
			Field("blob_id").
			Required().
			Unique(),
		edge.To("link", Blob.Type).
			Field("link_id").
			Required().
			Unique(),
	}
}
