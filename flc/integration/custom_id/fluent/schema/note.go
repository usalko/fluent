// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/google/uuid"
	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

type NoteID string

// Note holds the schema definition for the Note entity.
type Note struct {
	fluent.Schema
}

// Fields of the Note.
func (Note) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("id").
			GoType(NoteID("")).
			MaxLen(36).
			NotEmpty().
			Unique().
			Immutable().
			DefaultFunc(func() NoteID {
				return NoteID(uuid.NewString())
			}),
		field.String("text").
			Optional(),
	}
}

// Edges of the Note.
func (Note) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("children", Note.Type).
			From("parent").
			Unique(),
	}
}
