// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/blob"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
	"github.com/google/uuid"
)

// BlobUpdate is the builder for updating Blob entities.
type BlobUpdate struct {
	config
	hooks    []Hook
	mutation *BlobMutation
}

// Where appends a list predicates to the BlobUpdate builder.
func (bu *BlobUpdate) Where(ps ...predicate.Blob) *BlobUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUUID sets the "uuid" field.
func (bu *BlobUpdate) SetUUID(u uuid.UUID) *BlobUpdate {
	bu.mutation.SetUUID(u)
	return bu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (bu *BlobUpdate) SetNillableUUID(u *uuid.UUID) *BlobUpdate {
	if u != nil {
		bu.SetUUID(*u)
	}
	return bu
}

// SetCount sets the "count" field.
func (bu *BlobUpdate) SetCount(i int) *BlobUpdate {
	bu.mutation.ResetCount()
	bu.mutation.SetCount(i)
	return bu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (bu *BlobUpdate) SetNillableCount(i *int) *BlobUpdate {
	if i != nil {
		bu.SetCount(*i)
	}
	return bu
}

// AddCount adds i to the "count" field.
func (bu *BlobUpdate) AddCount(i int) *BlobUpdate {
	bu.mutation.AddCount(i)
	return bu
}

// SetParentID sets the "parent" edge to the Blob entity by ID.
func (bu *BlobUpdate) SetParentID(id uuid.UUID) *BlobUpdate {
	bu.mutation.SetParentID(id)
	return bu
}

// SetNillableParentID sets the "parent" edge to the Blob entity by ID if the given value is not nil.
func (bu *BlobUpdate) SetNillableParentID(id *uuid.UUID) *BlobUpdate {
	if id != nil {
		bu = bu.SetParentID(*id)
	}
	return bu
}

// SetParent sets the "parent" edge to the Blob entity.
func (bu *BlobUpdate) SetParent(b *Blob) *BlobUpdate {
	return bu.SetParentID(b.ID)
}

// AddLinkIDs adds the "links" edge to the Blob entity by IDs.
func (bu *BlobUpdate) AddLinkIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.AddLinkIDs(ids...)
	return bu
}

// AddLinks adds the "links" edges to the Blob entity.
func (bu *BlobUpdate) AddLinks(b ...*Blob) *BlobUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.AddLinkIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (bu *BlobUpdate) Mutation() *BlobMutation {
	return bu.mutation
}

// ClearParent clears the "parent" edge to the Blob entity.
func (bu *BlobUpdate) ClearParent() *BlobUpdate {
	bu.mutation.ClearParent()
	return bu
}

// ClearLinks clears all "links" edges to the Blob entity.
func (bu *BlobUpdate) ClearLinks() *BlobUpdate {
	bu.mutation.ClearLinks()
	return bu
}

// RemoveLinkIDs removes the "links" edge to Blob entities by IDs.
func (bu *BlobUpdate) RemoveLinkIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.RemoveLinkIDs(ids...)
	return bu
}

// RemoveLinks removes "links" edges to Blob entities.
func (bu *BlobUpdate) RemoveLinks(b ...*Blob) *BlobUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.RemoveLinkIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlobUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlobUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlobUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BlobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blob.Table, blob.Columns, sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.SetField(blob.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := bu.mutation.Count(); ok {
		_spec.SetField(blob.FieldCount, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedCount(); ok {
		_spec.AddField(blob.FieldCount, field.TypeInt, value)
	}
	if bu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		createE := &BlobLinkCreate{config: bu.config, mutation: newBlobLinkMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedLinksIDs(); len(nodes) > 0 && !bu.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: bu.config, mutation: newBlobLinkMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: bu.config, mutation: newBlobLinkMutation(bu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlobUpdateOne is the builder for updating a single Blob entity.
type BlobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlobMutation
}

// SetUUID sets the "uuid" field.
func (buo *BlobUpdateOne) SetUUID(u uuid.UUID) *BlobUpdateOne {
	buo.mutation.SetUUID(u)
	return buo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableUUID(u *uuid.UUID) *BlobUpdateOne {
	if u != nil {
		buo.SetUUID(*u)
	}
	return buo
}

// SetCount sets the "count" field.
func (buo *BlobUpdateOne) SetCount(i int) *BlobUpdateOne {
	buo.mutation.ResetCount()
	buo.mutation.SetCount(i)
	return buo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableCount(i *int) *BlobUpdateOne {
	if i != nil {
		buo.SetCount(*i)
	}
	return buo
}

// AddCount adds i to the "count" field.
func (buo *BlobUpdateOne) AddCount(i int) *BlobUpdateOne {
	buo.mutation.AddCount(i)
	return buo
}

// SetParentID sets the "parent" edge to the Blob entity by ID.
func (buo *BlobUpdateOne) SetParentID(id uuid.UUID) *BlobUpdateOne {
	buo.mutation.SetParentID(id)
	return buo
}

// SetNillableParentID sets the "parent" edge to the Blob entity by ID if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableParentID(id *uuid.UUID) *BlobUpdateOne {
	if id != nil {
		buo = buo.SetParentID(*id)
	}
	return buo
}

// SetParent sets the "parent" edge to the Blob entity.
func (buo *BlobUpdateOne) SetParent(b *Blob) *BlobUpdateOne {
	return buo.SetParentID(b.ID)
}

// AddLinkIDs adds the "links" edge to the Blob entity by IDs.
func (buo *BlobUpdateOne) AddLinkIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.AddLinkIDs(ids...)
	return buo
}

// AddLinks adds the "links" edges to the Blob entity.
func (buo *BlobUpdateOne) AddLinks(b ...*Blob) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.AddLinkIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (buo *BlobUpdateOne) Mutation() *BlobMutation {
	return buo.mutation
}

// ClearParent clears the "parent" edge to the Blob entity.
func (buo *BlobUpdateOne) ClearParent() *BlobUpdateOne {
	buo.mutation.ClearParent()
	return buo
}

// ClearLinks clears all "links" edges to the Blob entity.
func (buo *BlobUpdateOne) ClearLinks() *BlobUpdateOne {
	buo.mutation.ClearLinks()
	return buo
}

// RemoveLinkIDs removes the "links" edge to Blob entities by IDs.
func (buo *BlobUpdateOne) RemoveLinkIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.RemoveLinkIDs(ids...)
	return buo
}

// RemoveLinks removes "links" edges to Blob entities.
func (buo *BlobUpdateOne) RemoveLinks(b ...*Blob) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.RemoveLinkIDs(ids...)
}

// Where appends a list predicates to the BlobUpdate builder.
func (buo *BlobUpdateOne) Where(ps ...predicate.Blob) *BlobUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlobUpdateOne) Select(field string, fields ...string) *BlobUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blob entity.
func (buo *BlobUpdateOne) Save(ctx context.Context) (*Blob, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlobUpdateOne) SaveX(ctx context.Context) *Blob {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlobUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlobUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BlobUpdateOne) sqlSave(ctx context.Context) (_node *Blob, err error) {
	_spec := sqlgraph.NewUpdateSpec(blob.Table, blob.Columns, sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blob.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blob.FieldID)
		for _, f := range fields {
			if !blob.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UUID(); ok {
		_spec.SetField(blob.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := buo.mutation.Count(); ok {
		_spec.SetField(blob.FieldCount, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedCount(); ok {
		_spec.AddField(blob.FieldCount, field.TypeInt, value)
	}
	if buo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		createE := &BlobLinkCreate{config: buo.config, mutation: newBlobLinkMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedLinksIDs(); len(nodes) > 0 && !buo.mutation.LinksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: buo.config, mutation: newBlobLinkMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: buo.config, mutation: newBlobLinkMutation(buo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Blob{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
