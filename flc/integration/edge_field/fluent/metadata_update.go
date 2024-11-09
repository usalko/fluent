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
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/metadata"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/edge_field/fluent/user"
	"github.com/usalko/fluent/schema/field"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks    []Hook
	mutation *MetadataMutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetAge sets the "age" field.
func (mu *MetadataUpdate) SetAge(i int) *MetadataUpdate {
	mu.mutation.ResetAge()
	mu.mutation.SetAge(i)
	return mu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableAge(i *int) *MetadataUpdate {
	if i != nil {
		mu.SetAge(*i)
	}
	return mu
}

// AddAge adds i to the "age" field.
func (mu *MetadataUpdate) AddAge(i int) *MetadataUpdate {
	mu.mutation.AddAge(i)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MetadataUpdate) SetParentID(i int) *MetadataUpdate {
	mu.mutation.SetParentID(i)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MetadataUpdate) SetNillableParentID(i *int) *MetadataUpdate {
	if i != nil {
		mu.SetParentID(*i)
	}
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MetadataUpdate) ClearParentID() *MetadataUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mu *MetadataUpdate) SetUserID(id int) *MetadataUpdate {
	mu.mutation.SetUserID(id)
	return mu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mu *MetadataUpdate) SetNillableUserID(id *int) *MetadataUpdate {
	if id != nil {
		mu = mu.SetUserID(*id)
	}
	return mu
}

// SetUser sets the "user" edge to the User entity.
func (mu *MetadataUpdate) SetUser(u *User) *MetadataUpdate {
	return mu.SetUserID(u.ID)
}

// AddChildIDs adds the "children" edge to the Metadata entity by IDs.
func (mu *MetadataUpdate) AddChildIDs(ids ...int) *MetadataUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Metadata entity.
func (mu *MetadataUpdate) AddChildren(m ...*Metadata) *MetadataUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// SetParent sets the "parent" edge to the Metadata entity.
func (mu *MetadataUpdate) SetParent(m *Metadata) *MetadataUpdate {
	return mu.SetParentID(m.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mu *MetadataUpdate) ClearUser() *MetadataUpdate {
	mu.mutation.ClearUser()
	return mu
}

// ClearChildren clears all "children" edges to the Metadata entity.
func (mu *MetadataUpdate) ClearChildren() *MetadataUpdate {
	mu.mutation.ClearChildren()
	return mu
}

// RemoveChildIDs removes the "children" edge to Metadata entities by IDs.
func (mu *MetadataUpdate) RemoveChildIDs(ids ...int) *MetadataUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Metadata entities.
func (mu *MetadataUpdate) RemoveChildren(m ...*Metadata) *MetadataUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Metadata entity.
func (mu *MetadataUpdate) ClearParent() *MetadataUpdate {
	mu.mutation.ClearParent()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Age(); ok {
		_spec.SetField(metadata.FieldAge, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedAge(); ok {
		_spec.AddField(metadata.FieldAge, field.TypeInt, value)
	}
	if mu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   metadata.ParentTable,
			Columns: []string{metadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   metadata.ParentTable,
			Columns: []string{metadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetadataMutation
}

// SetAge sets the "age" field.
func (muo *MetadataUpdateOne) SetAge(i int) *MetadataUpdateOne {
	muo.mutation.ResetAge()
	muo.mutation.SetAge(i)
	return muo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableAge(i *int) *MetadataUpdateOne {
	if i != nil {
		muo.SetAge(*i)
	}
	return muo
}

// AddAge adds i to the "age" field.
func (muo *MetadataUpdateOne) AddAge(i int) *MetadataUpdateOne {
	muo.mutation.AddAge(i)
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MetadataUpdateOne) SetParentID(i int) *MetadataUpdateOne {
	muo.mutation.SetParentID(i)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableParentID(i *int) *MetadataUpdateOne {
	if i != nil {
		muo.SetParentID(*i)
	}
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MetadataUpdateOne) ClearParentID() *MetadataUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (muo *MetadataUpdateOne) SetUserID(id int) *MetadataUpdateOne {
	muo.mutation.SetUserID(id)
	return muo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (muo *MetadataUpdateOne) SetNillableUserID(id *int) *MetadataUpdateOne {
	if id != nil {
		muo = muo.SetUserID(*id)
	}
	return muo
}

// SetUser sets the "user" edge to the User entity.
func (muo *MetadataUpdateOne) SetUser(u *User) *MetadataUpdateOne {
	return muo.SetUserID(u.ID)
}

// AddChildIDs adds the "children" edge to the Metadata entity by IDs.
func (muo *MetadataUpdateOne) AddChildIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Metadata entity.
func (muo *MetadataUpdateOne) AddChildren(m ...*Metadata) *MetadataUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// SetParent sets the "parent" edge to the Metadata entity.
func (muo *MetadataUpdateOne) SetParent(m *Metadata) *MetadataUpdateOne {
	return muo.SetParentID(m.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (muo *MetadataUpdateOne) ClearUser() *MetadataUpdateOne {
	muo.mutation.ClearUser()
	return muo
}

// ClearChildren clears all "children" edges to the Metadata entity.
func (muo *MetadataUpdateOne) ClearChildren() *MetadataUpdateOne {
	muo.mutation.ClearChildren()
	return muo
}

// RemoveChildIDs removes the "children" edge to Metadata entities by IDs.
func (muo *MetadataUpdateOne) RemoveChildIDs(ids ...int) *MetadataUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Metadata entities.
func (muo *MetadataUpdateOne) RemoveChildren(m ...*Metadata) *MetadataUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Metadata entity.
func (muo *MetadataUpdateOne) ClearParent() *MetadataUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// Where appends a list predicates to the MetadataUpdate builder.
func (muo *MetadataUpdateOne) Where(ps ...predicate.Metadata) *MetadataUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Age(); ok {
		_spec.SetField(metadata.FieldAge, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedAge(); ok {
		_spec.AddField(metadata.FieldAge, field.TypeInt, value)
	}
	if muo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   metadata.ParentTable,
			Columns: []string{metadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   metadata.ParentTable,
			Columns: []string{metadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
