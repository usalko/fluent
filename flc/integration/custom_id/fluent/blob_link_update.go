// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/blob"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/blob_link"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// BlobLinkUpdate is the builder for updating BlobLink entities.
type BlobLinkUpdate struct {
	config
	hooks    []Hook
	mutation *BlobLinkMutation
}

// Where appends a list predicates to the BlobLinkUpdate builder.
func (blu *BlobLinkUpdate) Where(ps ...predicate.BlobLink) *BlobLinkUpdate {
	blu.mutation.Where(ps...)
	return blu
}

// SetCreatedAt sets the "created_at" field.
func (blu *BlobLinkUpdate) SetCreatedAt(tt time.Time) *BlobLinkUpdate {
	blu.mutation.SetCreatedAt(tt)
	return blu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (blu *BlobLinkUpdate) SetNillableCreatedAt(tt *time.Time) *BlobLinkUpdate {
	if tt != nil {
		blu.SetCreatedAt(*tt)
	}
	return blu
}

// SetBlobID sets the "blob_id" field.
func (blu *BlobLinkUpdate) SetBlobID(uu uuid.UUID) *BlobLinkUpdate {
	blu.mutation.SetBlobID(uu)
	return blu
}

// SetNillableBlobID sets the "blob_id" field if the given value is not nil.
func (blu *BlobLinkUpdate) SetNillableBlobID(uu *uuid.UUID) *BlobLinkUpdate {
	if uu != nil {
		blu.SetBlobID(*uu)
	}
	return blu
}

// SetLinkID sets the "link_id" field.
func (blu *BlobLinkUpdate) SetLinkID(uu uuid.UUID) *BlobLinkUpdate {
	blu.mutation.SetLinkID(uu)
	return blu
}

// SetNillableLinkID sets the "link_id" field if the given value is not nil.
func (blu *BlobLinkUpdate) SetNillableLinkID(uu *uuid.UUID) *BlobLinkUpdate {
	if uu != nil {
		blu.SetLinkID(*uu)
	}
	return blu
}

// SetBlob sets the "blob" edge to the Blob entity.
func (blu *BlobLinkUpdate) SetBlob(b *Blob) *BlobLinkUpdate {
	return blu.SetBlobID(b.ID)
}

// SetLink sets the "link" edge to the Blob entity.
func (blu *BlobLinkUpdate) SetLink(b *Blob) *BlobLinkUpdate {
	return blu.SetLinkID(b.ID)
}

// Mutation returns the BlobLinkMutation object of the builder.
func (blu *BlobLinkUpdate) Mutation() *BlobLinkMutation {
	return blu.mutation
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (blu *BlobLinkUpdate) ClearBlob() *BlobLinkUpdate {
	blu.mutation.ClearBlob()
	return blu
}

// ClearLink clears the "link" edge to the Blob entity.
func (blu *BlobLinkUpdate) ClearLink() *BlobLinkUpdate {
	blu.mutation.ClearLink()
	return blu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (blu *BlobLinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, blu.sqlSave, blu.mutation, blu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (blu *BlobLinkUpdate) SaveX(ctx context.Context) int {
	affected, err := blu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (blu *BlobLinkUpdate) Exec(ctx context.Context) error {
	_, err := blu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blu *BlobLinkUpdate) ExecX(ctx context.Context) {
	if err := blu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blu *BlobLinkUpdate) check() error {
	if blu.mutation.BlobCleared() && len(blu.mutation.BlobIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "BlobLink.blob"`)
	}
	if blu.mutation.LinkCleared() && len(blu.mutation.LinkIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "BlobLink.link"`)
	}
	return nil
}

func (blu *BlobLinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := blu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(blob_link.Table, blob_link.Columns, sqlgraph.NewFieldSpec(blob_link.FieldBlobID, field.TypeUUID), sqlgraph.NewFieldSpec(blob_link.FieldLinkID, field.TypeUUID))
	if ps := blu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := blu.mutation.CreatedAt(); ok {
		_spec.SetField(blob_link.FieldCreatedAt, field.TypeTime, value)
	}
	if blu.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.BlobTable,
			Columns: []string{blob_link.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.BlobTable,
			Columns: []string{blob_link.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if blu.mutation.LinkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.LinkTable,
			Columns: []string{blob_link.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := blu.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.LinkTable,
			Columns: []string{blob_link.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, blu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob_link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	blu.mutation.done = true
	return n, nil
}

// BlobLinkUpdateOne is the builder for updating a single BlobLink entity.
type BlobLinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlobLinkMutation
}

// SetCreatedAt sets the "created_at" field.
func (bluo *BlobLinkUpdateOne) SetCreatedAt(tt time.Time) *BlobLinkUpdateOne {
	bluo.mutation.SetCreatedAt(tt)
	return bluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bluo *BlobLinkUpdateOne) SetNillableCreatedAt(tt *time.Time) *BlobLinkUpdateOne {
	if tt != nil {
		bluo.SetCreatedAt(*tt)
	}
	return bluo
}

// SetBlobID sets the "blob_id" field.
func (bluo *BlobLinkUpdateOne) SetBlobID(uu uuid.UUID) *BlobLinkUpdateOne {
	bluo.mutation.SetBlobID(uu)
	return bluo
}

// SetNillableBlobID sets the "blob_id" field if the given value is not nil.
func (bluo *BlobLinkUpdateOne) SetNillableBlobID(uu *uuid.UUID) *BlobLinkUpdateOne {
	if uu != nil {
		bluo.SetBlobID(*uu)
	}
	return bluo
}

// SetLinkID sets the "link_id" field.
func (bluo *BlobLinkUpdateOne) SetLinkID(uu uuid.UUID) *BlobLinkUpdateOne {
	bluo.mutation.SetLinkID(uu)
	return bluo
}

// SetNillableLinkID sets the "link_id" field if the given value is not nil.
func (bluo *BlobLinkUpdateOne) SetNillableLinkID(uu *uuid.UUID) *BlobLinkUpdateOne {
	if uu != nil {
		bluo.SetLinkID(*uu)
	}
	return bluo
}

// SetBlob sets the "blob" edge to the Blob entity.
func (bluo *BlobLinkUpdateOne) SetBlob(b *Blob) *BlobLinkUpdateOne {
	return bluo.SetBlobID(b.ID)
}

// SetLink sets the "link" edge to the Blob entity.
func (bluo *BlobLinkUpdateOne) SetLink(b *Blob) *BlobLinkUpdateOne {
	return bluo.SetLinkID(b.ID)
}

// Mutation returns the BlobLinkMutation object of the builder.
func (bluo *BlobLinkUpdateOne) Mutation() *BlobLinkMutation {
	return bluo.mutation
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (bluo *BlobLinkUpdateOne) ClearBlob() *BlobLinkUpdateOne {
	bluo.mutation.ClearBlob()
	return bluo
}

// ClearLink clears the "link" edge to the Blob entity.
func (bluo *BlobLinkUpdateOne) ClearLink() *BlobLinkUpdateOne {
	bluo.mutation.ClearLink()
	return bluo
}

// Where appends a list predicates to the BlobLinkUpdate builder.
func (bluo *BlobLinkUpdateOne) Where(ps ...predicate.BlobLink) *BlobLinkUpdateOne {
	bluo.mutation.Where(ps...)
	return bluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bluo *BlobLinkUpdateOne) Select(field string, fields ...string) *BlobLinkUpdateOne {
	bluo.fields = append([]string{field}, fields...)
	return bluo
}

// Save executes the query and returns the updated BlobLink entity.
func (bluo *BlobLinkUpdateOne) Save(ctx context.Context) (*BlobLink, error) {
	return withHooks(ctx, bluo.sqlSave, bluo.mutation, bluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bluo *BlobLinkUpdateOne) SaveX(ctx context.Context) *BlobLink {
	node, err := bluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bluo *BlobLinkUpdateOne) Exec(ctx context.Context) error {
	_, err := bluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bluo *BlobLinkUpdateOne) ExecX(ctx context.Context) {
	if err := bluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bluo *BlobLinkUpdateOne) check() error {
	if bluo.mutation.BlobCleared() && len(bluo.mutation.BlobIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "BlobLink.blob"`)
	}
	if bluo.mutation.LinkCleared() && len(bluo.mutation.LinkIDs()) > 0 {
		return errors.New(`fluent: clearing a required unique edge "BlobLink.link"`)
	}
	return nil
}

func (bluo *BlobLinkUpdateOne) sqlSave(ctx context.Context) (_node *BlobLink, err error) {
	if err := bluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(blob_link.Table, blob_link.Columns, sqlgraph.NewFieldSpec(blob_link.FieldBlobID, field.TypeUUID), sqlgraph.NewFieldSpec(blob_link.FieldLinkID, field.TypeUUID))
	if id, ok := bluo.mutation.BlobID(); !ok {
		return nil, &ValidationError{Name: "blob_id", err: errors.New(`fluent: missing "BlobLink.blob_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := bluo.mutation.LinkID(); !ok {
		return nil, &ValidationError{Name: "link_id", err: errors.New(`fluent: missing "BlobLink.link_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := bluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !blob_link.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := bluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bluo.mutation.CreatedAt(); ok {
		_spec.SetField(blob_link.FieldCreatedAt, field.TypeTime, value)
	}
	if bluo.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.BlobTable,
			Columns: []string{blob_link.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.BlobTable,
			Columns: []string{blob_link.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bluo.mutation.LinkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.LinkTable,
			Columns: []string{blob_link.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bluo.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blob_link.LinkTable,
			Columns: []string{blob_link.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BlobLink{config: bluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob_link.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bluo.mutation.done = true
	return _node, nil
}