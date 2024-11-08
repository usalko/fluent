// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema/edge"
	"github.com/usalko/fluent/schema/field"
)

// An example for an edge-schema (AttachedFile) that uses a custom name for its edges
// (not foreign-keys) and Ent matches by the schema types and not foreign-key names.

// Process holds the edge schema definition of the Process relationship.
type Process struct {
	fluent.Schema
}

// Edges of the Process.
func (Process) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.To("files", File.Type).
			Through("attached_files", AttachedFile.Type).
			Comment("Files that were attached by this process"),
	}
}

// File holds the edge schema definition of the File relationship.
type File struct {
	fluent.Schema
}

// Fields of the File.
func (File) Fields() []fluent.Field {
	return []fluent.Field{
		field.String("name"),
	}
}

// Edges of the File.
func (File) Edges() []fluent.Edge {
	return []fluent.Edge{
		edge.From("processes", Process.Type).
			Ref("files"),
	}
}

// AttachedFile holds the edge schema definition of the File relationship.
type AttachedFile struct {
	fluent.Schema
}

// Fields of the AttachedFile.
func (AttachedFile) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("attach_time").
			Default(time.Now),
		field.Int("f_id"),
		field.Int("proc_id"),
	}
}

// Edges of the AttachedFile.
func (AttachedFile) Edges() []fluent.Edge {
	return []fluent.Edge{
		// Note: the two following edges use different name conventions (e.g., f_id <> file_id), but
		// Ent knows how to resolve this as there is only one usage of each type in this declaration.
		edge.To("fi", File.Type).
			Required().
			Unique().
			Field("f_id"),
		edge.To("proc", Process.Type).
			Required().
			Unique().
			Field("proc_id"),
	}
}
