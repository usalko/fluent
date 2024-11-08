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
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/note"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/schema"
	"github.com/usalko/fluent/schema/field"
)

// NoteUpdate is the builder for updating Note entities.
type NoteUpdate struct {
	config
	hooks    []Hook
	mutation *NoteMutation
}

// Where appends a list predicates to the NoteUpdate builder.
func (nu *NoteUpdate) Where(ps ...predicate.Note) *NoteUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetText sets the "text" field.
func (nu *NoteUpdate) SetText(s string) *NoteUpdate {
	nu.mutation.SetText(s)
	return nu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableText(s *string) *NoteUpdate {
	if s != nil {
		nu.SetText(*s)
	}
	return nu
}

// ClearText clears the value of the "text" field.
func (nu *NoteUpdate) ClearText() *NoteUpdate {
	nu.mutation.ClearText()
	return nu
}

// SetParentID sets the "parent" edge to the Note entity by ID.
func (nu *NoteUpdate) SetParentID(id schema.NoteID) *NoteUpdate {
	nu.mutation.SetParentID(id)
	return nu
}

// SetNillableParentID sets the "parent" edge to the Note entity by ID if the given value is not nil.
func (nu *NoteUpdate) SetNillableParentID(id *schema.NoteID) *NoteUpdate {
	if id != nil {
		nu = nu.SetParentID(*id)
	}
	return nu
}

// SetParent sets the "parent" edge to the Note entity.
func (nu *NoteUpdate) SetParent(n *Note) *NoteUpdate {
	return nu.SetParentID(n.ID)
}

// AddChildIDs adds the "children" edge to the Note entity by IDs.
func (nu *NoteUpdate) AddChildIDs(ids ...schema.NoteID) *NoteUpdate {
	nu.mutation.AddChildIDs(ids...)
	return nu
}

// AddChildren adds the "children" edges to the Note entity.
func (nu *NoteUpdate) AddChildren(n ...*Note) *NoteUpdate {
	ids := make([]schema.NoteID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nu.AddChildIDs(ids...)
}

// Mutation returns the NoteMutation object of the builder.
func (nu *NoteUpdate) Mutation() *NoteMutation {
	return nu.mutation
}

// ClearParent clears the "parent" edge to the Note entity.
func (nu *NoteUpdate) ClearParent() *NoteUpdate {
	nu.mutation.ClearParent()
	return nu
}

// ClearChildren clears all "children" edges to the Note entity.
func (nu *NoteUpdate) ClearChildren() *NoteUpdate {
	nu.mutation.ClearChildren()
	return nu
}

// RemoveChildIDs removes the "children" edge to Note entities by IDs.
func (nu *NoteUpdate) RemoveChildIDs(ids ...schema.NoteID) *NoteUpdate {
	nu.mutation.RemoveChildIDs(ids...)
	return nu
}

// RemoveChildren removes "children" edges to Note entities.
func (nu *NoteUpdate) RemoveChildren(n ...*Note) *NoteUpdate {
	ids := make([]schema.NoteID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NoteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NoteUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NoteUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NoteUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Text(); ok {
		_spec.SetField(note.FieldText, field.TypeString, value)
	}
	if nu.mutation.TextCleared() {
		_spec.ClearField(note.FieldText, field.TypeString)
	}
	if nu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.ParentTable,
			Columns: []string{note.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.ParentTable,
			Columns: []string{note.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !nu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NoteUpdateOne is the builder for updating a single Note entity.
type NoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NoteMutation
}

// SetText sets the "text" field.
func (nuo *NoteUpdateOne) SetText(s string) *NoteUpdateOne {
	nuo.mutation.SetText(s)
	return nuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableText(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetText(*s)
	}
	return nuo
}

// ClearText clears the value of the "text" field.
func (nuo *NoteUpdateOne) ClearText() *NoteUpdateOne {
	nuo.mutation.ClearText()
	return nuo
}

// SetParentID sets the "parent" edge to the Note entity by ID.
func (nuo *NoteUpdateOne) SetParentID(id schema.NoteID) *NoteUpdateOne {
	nuo.mutation.SetParentID(id)
	return nuo
}

// SetNillableParentID sets the "parent" edge to the Note entity by ID if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableParentID(id *schema.NoteID) *NoteUpdateOne {
	if id != nil {
		nuo = nuo.SetParentID(*id)
	}
	return nuo
}

// SetParent sets the "parent" edge to the Note entity.
func (nuo *NoteUpdateOne) SetParent(n *Note) *NoteUpdateOne {
	return nuo.SetParentID(n.ID)
}

// AddChildIDs adds the "children" edge to the Note entity by IDs.
func (nuo *NoteUpdateOne) AddChildIDs(ids ...schema.NoteID) *NoteUpdateOne {
	nuo.mutation.AddChildIDs(ids...)
	return nuo
}

// AddChildren adds the "children" edges to the Note entity.
func (nuo *NoteUpdateOne) AddChildren(n ...*Note) *NoteUpdateOne {
	ids := make([]schema.NoteID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nuo.AddChildIDs(ids...)
}

// Mutation returns the NoteMutation object of the builder.
func (nuo *NoteUpdateOne) Mutation() *NoteMutation {
	return nuo.mutation
}

// ClearParent clears the "parent" edge to the Note entity.
func (nuo *NoteUpdateOne) ClearParent() *NoteUpdateOne {
	nuo.mutation.ClearParent()
	return nuo
}

// ClearChildren clears all "children" edges to the Note entity.
func (nuo *NoteUpdateOne) ClearChildren() *NoteUpdateOne {
	nuo.mutation.ClearChildren()
	return nuo
}

// RemoveChildIDs removes the "children" edge to Note entities by IDs.
func (nuo *NoteUpdateOne) RemoveChildIDs(ids ...schema.NoteID) *NoteUpdateOne {
	nuo.mutation.RemoveChildIDs(ids...)
	return nuo
}

// RemoveChildren removes "children" edges to Note entities.
func (nuo *NoteUpdateOne) RemoveChildren(n ...*Note) *NoteUpdateOne {
	ids := make([]schema.NoteID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the NoteUpdate builder.
func (nuo *NoteUpdateOne) Where(ps ...predicate.Note) *NoteUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NoteUpdateOne) Select(field string, fields ...string) *NoteUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Note entity.
func (nuo *NoteUpdateOne) Save(ctx context.Context) (*Note, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NoteUpdateOne) SaveX(ctx context.Context) *Note {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NoteUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NoteUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NoteUpdateOne) sqlSave(ctx context.Context) (_node *Note, err error) {
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Note.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, note.FieldID)
		for _, f := range fields {
			if !note.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != note.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Text(); ok {
		_spec.SetField(note.FieldText, field.TypeString, value)
	}
	if nuo.mutation.TextCleared() {
		_spec.ClearField(note.FieldText, field.TypeString)
	}
	if nuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.ParentTable,
			Columns: []string{note.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.ParentTable,
			Columns: []string{note.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !nuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   note.ChildrenTable,
			Columns: []string{note.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(note.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Note{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}