// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mixin

import (
	"time"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// Schema is the default implementation for the fluent.Mixin interface.
// It should be embedded in end-user mixin as follows:
//
//	type M struct {
//		mixin.Schema
//	}
type Schema struct{}

// Fields of the mixin.
func (Schema) Fields() []fluent.Field { return nil }

// Edges of the mixin.
func (Schema) Edges() []fluent.Edge { return nil }

// Indexes of the mixin.
func (Schema) Indexes() []fluent.Index { return nil }

// Hooks of the mixin.
func (Schema) Hooks() []fluent.Hook { return nil }

// Interceptors of the schema.
func (Schema) Interceptors() []fluent.Interceptor { return nil }

// Policy of the mixin.
func (Schema) Policy() fluent.Policy { return nil }

// Annotations of the mixin.
func (Schema) Annotations() []schema.Annotation { return nil }

// schema mixin must implement `Mixin` interface.
var _ fluent.Mixin = (*Schema)(nil)

// CreateTime adds created at time field.
type CreateTime struct{ Schema }

// Fields of the create time mixin.
func (CreateTime) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ fluent.Mixin = (*CreateTime)(nil)

// UpdateTime adds updated at time field.
type UpdateTime struct{ Schema }

// Fields of the update time mixin.
func (UpdateTime) Fields() []fluent.Field {
	return []fluent.Field{
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// update time mixin must implement `Mixin` interface.
var _ fluent.Mixin = (*UpdateTime)(nil)

// Time composes create/update time mixin.
type Time struct{ Schema }

// Fields of the time mixin.
func (Time) Fields() []fluent.Field {
	return append(
		CreateTime{}.Fields(),
		UpdateTime{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ fluent.Mixin = (*Time)(nil)

// AnnotateFields adds field annotations to underlying mixin fields.
func AnnotateFields(m fluent.Mixin, annotations ...schema.Annotation) fluent.Mixin {
	return fieldAnnotator{Mixin: m, annotations: annotations}
}

// AnnotateEdges adds edge annotations to underlying mixin edges.
func AnnotateEdges(m fluent.Mixin, annotations ...schema.Annotation) fluent.Mixin {
	return edgeAnnotator{Mixin: m, annotations: annotations}
}

type fieldAnnotator struct {
	fluent.Mixin
	annotations []schema.Annotation
}

func (a fieldAnnotator) Fields() []fluent.Field {
	fields := a.Mixin.Fields()
	for i := range fields {
		desc := fields[i].Descriptor()
		desc.Annotations = append(desc.Annotations, a.annotations...)
	}
	return fields
}

type edgeAnnotator struct {
	fluent.Mixin
	annotations []schema.Annotation
}

func (a edgeAnnotator) Edges() []fluent.Edge {
	edges := a.Mixin.Edges()
	for i := range edges {
		desc := edges[i].Descriptor()
		desc.Annotations = append(desc.Annotations, a.annotations...)
	}
	return edges
}
