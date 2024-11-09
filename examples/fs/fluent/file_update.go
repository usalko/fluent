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
	"github.com/usalko/fluent/examples/fs/fluent/file"
	"github.com/usalko/fluent/examples/fs/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FileUpdate) SetNillableName(s *string) *FileUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetDeleted sets the "deleted" field.
func (fu *FileUpdate) SetDeleted(b bool) *FileUpdate {
	fu.mutation.SetDeleted(b)
	return fu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (fu *FileUpdate) SetNillableDeleted(b *bool) *FileUpdate {
	if b != nil {
		fu.SetDeleted(*b)
	}
	return fu
}

// SetParentID sets the "parent_id" field.
func (fu *FileUpdate) SetParentID(i int) *FileUpdate {
	fu.mutation.SetParentID(i)
	return fu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (fu *FileUpdate) SetNillableParentID(i *int) *FileUpdate {
	if i != nil {
		fu.SetParentID(*i)
	}
	return fu
}

// ClearParentID clears the value of the "parent_id" field.
func (fu *FileUpdate) ClearParentID() *FileUpdate {
	fu.mutation.ClearParentID()
	return fu
}

// SetParent sets the "parent" edge to the File entity.
func (fu *FileUpdate) SetParent(f *File) *FileUpdate {
	return fu.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the File entity by IDs.
func (fu *FileUpdate) AddChildIDs(ids ...int) *FileUpdate {
	fu.mutation.AddChildIDs(ids...)
	return fu
}

// AddChildren adds the "children" edges to the File entity.
func (fu *FileUpdate) AddChildren(f ...*File) *FileUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddChildIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearParent clears the "parent" edge to the File entity.
func (fu *FileUpdate) ClearParent() *FileUpdate {
	fu.mutation.ClearParent()
	return fu
}

// ClearChildren clears all "children" edges to the File entity.
func (fu *FileUpdate) ClearChildren() *FileUpdate {
	fu.mutation.ClearChildren()
	return fu
}

// RemoveChildIDs removes the "children" edge to File entities by IDs.
func (fu *FileUpdate) RemoveChildIDs(ids ...int) *FileUpdate {
	fu.mutation.RemoveChildIDs(ids...)
	return fu
}

// RemoveChildren removes "children" edges to File entities.
func (fu *FileUpdate) RemoveChildren(f ...*File) *FileUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Deleted(); ok {
		_spec.SetField(file.FieldDeleted, field.TypeBool, value)
	}
	if fu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ParentTable,
			Columns: []string{file.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ParentTable,
			Columns: []string{file.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableName(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetDeleted sets the "deleted" field.
func (fuo *FileUpdateOne) SetDeleted(b bool) *FileUpdateOne {
	fuo.mutation.SetDeleted(b)
	return fuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableDeleted(b *bool) *FileUpdateOne {
	if b != nil {
		fuo.SetDeleted(*b)
	}
	return fuo
}

// SetParentID sets the "parent_id" field.
func (fuo *FileUpdateOne) SetParentID(i int) *FileUpdateOne {
	fuo.mutation.SetParentID(i)
	return fuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableParentID(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetParentID(*i)
	}
	return fuo
}

// ClearParentID clears the value of the "parent_id" field.
func (fuo *FileUpdateOne) ClearParentID() *FileUpdateOne {
	fuo.mutation.ClearParentID()
	return fuo
}

// SetParent sets the "parent" edge to the File entity.
func (fuo *FileUpdateOne) SetParent(f *File) *FileUpdateOne {
	return fuo.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the File entity by IDs.
func (fuo *FileUpdateOne) AddChildIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.AddChildIDs(ids...)
	return fuo
}

// AddChildren adds the "children" edges to the File entity.
func (fuo *FileUpdateOne) AddChildren(f ...*File) *FileUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddChildIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearParent clears the "parent" edge to the File entity.
func (fuo *FileUpdateOne) ClearParent() *FileUpdateOne {
	fuo.mutation.ClearParent()
	return fuo
}

// ClearChildren clears all "children" edges to the File entity.
func (fuo *FileUpdateOne) ClearChildren() *FileUpdateOne {
	fuo.mutation.ClearChildren()
	return fuo
}

// RemoveChildIDs removes the "children" edge to File entities by IDs.
func (fuo *FileUpdateOne) RemoveChildIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.RemoveChildIDs(ids...)
	return fuo
}

// RemoveChildren removes "children" edges to File entities.
func (fuo *FileUpdateOne) RemoveChildren(f ...*File) *FileUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Deleted(); ok {
		_spec.SetField(file.FieldDeleted, field.TypeBool, value)
	}
	if fuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ParentTable,
			Columns: []string{file.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ParentTable,
			Columns: []string{file.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.ChildrenTable,
			Columns: []string{file.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
