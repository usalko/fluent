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
	"github.com/usalko/fluent/flc/integration/fluent/comment"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	schemadir "github.com/usalko/fluent/flc/integration/fluent/schema/dir"
	"github.com/usalko/fluent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUniqueInt sets the "unique_int" field.
func (cu *CommentUpdate) SetUniqueInt(i int) *CommentUpdate {
	cu.mutation.ResetUniqueInt()
	cu.mutation.SetUniqueInt(i)
	return cu
}

// SetNillableUniqueInt sets the "unique_int" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUniqueInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetUniqueInt(*i)
	}
	return cu
}

// AddUniqueInt adds i to the "unique_int" field.
func (cu *CommentUpdate) AddUniqueInt(i int) *CommentUpdate {
	cu.mutation.AddUniqueInt(i)
	return cu
}

// SetUniqueFloat sets the "unique_float" field.
func (cu *CommentUpdate) SetUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.ResetUniqueFloat()
	cu.mutation.SetUniqueFloat(f)
	return cu
}

// SetNillableUniqueFloat sets the "unique_float" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUniqueFloat(f *float64) *CommentUpdate {
	if f != nil {
		cu.SetUniqueFloat(*f)
	}
	return cu
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cu *CommentUpdate) AddUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.AddUniqueFloat(f)
	return cu
}

// SetNillableInt sets the "nillable_int" field.
func (cu *CommentUpdate) SetNillableInt(i int) *CommentUpdate {
	cu.mutation.ResetNillableInt()
	cu.mutation.SetNillableInt(i)
	return cu
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableNillableInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetNillableInt(*i)
	}
	return cu
}

// AddNillableInt adds i to the "nillable_int" field.
func (cu *CommentUpdate) AddNillableInt(i int) *CommentUpdate {
	cu.mutation.AddNillableInt(i)
	return cu
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cu *CommentUpdate) ClearNillableInt() *CommentUpdate {
	cu.mutation.ClearNillableInt()
	return cu
}

// SetTable sets the "table" field.
func (cu *CommentUpdate) SetTable(s string) *CommentUpdate {
	cu.mutation.SetTable(s)
	return cu
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableTable(s *string) *CommentUpdate {
	if s != nil {
		cu.SetTable(*s)
	}
	return cu
}

// ClearTable clears the value of the "table" field.
func (cu *CommentUpdate) ClearTable() *CommentUpdate {
	cu.mutation.ClearTable()
	return cu
}

// SetDir sets the "dir" field.
func (cu *CommentUpdate) SetDir(s schemadir.Dir) *CommentUpdate {
	cu.mutation.SetDir(s)
	return cu
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDir(s *schemadir.Dir) *CommentUpdate {
	if s != nil {
		cu.SetDir(*s)
	}
	return cu
}

// ClearDir clears the value of the "dir" field.
func (cu *CommentUpdate) ClearDir() *CommentUpdate {
	cu.mutation.ClearDir()
	return cu
}

// SetClient sets the "client" field.
func (cu *CommentUpdate) SetClient(s string) *CommentUpdate {
	cu.mutation.SetClient(s)
	return cu
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableClient(s *string) *CommentUpdate {
	if s != nil {
		cu.SetClient(*s)
	}
	return cu
}

// ClearClient clears the value of the "client" field.
func (cu *CommentUpdate) ClearClient() *CommentUpdate {
	cu.mutation.ClearClient()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CommentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UniqueInt(); ok {
		_spec.SetField(comment.FieldUniqueInt, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedUniqueInt(); ok {
		_spec.AddField(comment.FieldUniqueInt, field.TypeInt, value)
	}
	if value, ok := cu.mutation.UniqueFloat(); ok {
		_spec.SetField(comment.FieldUniqueFloat, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedUniqueFloat(); ok {
		_spec.AddField(comment.FieldUniqueFloat, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.NillableInt(); ok {
		_spec.SetField(comment.FieldNillableInt, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedNillableInt(); ok {
		_spec.AddField(comment.FieldNillableInt, field.TypeInt, value)
	}
	if cu.mutation.NillableIntCleared() {
		_spec.ClearField(comment.FieldNillableInt, field.TypeInt)
	}
	if value, ok := cu.mutation.Table(); ok {
		_spec.SetField(comment.FieldTable, field.TypeString, value)
	}
	if cu.mutation.TableCleared() {
		_spec.ClearField(comment.FieldTable, field.TypeString)
	}
	if value, ok := cu.mutation.Dir(); ok {
		_spec.SetField(comment.FieldDir, field.TypeJSON, value)
	}
	if cu.mutation.DirCleared() {
		_spec.ClearField(comment.FieldDir, field.TypeJSON)
	}
	if value, ok := cu.mutation.GetClient(); ok {
		_spec.SetField(comment.FieldClient, field.TypeString, value)
	}
	if cu.mutation.ClientCleared() {
		_spec.ClearField(comment.FieldClient, field.TypeString)
	}
	_spec.AddModifiers(cu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUniqueInt sets the "unique_int" field.
func (cuo *CommentUpdateOne) SetUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetUniqueInt()
	cuo.mutation.SetUniqueInt(i)
	return cuo
}

// SetNillableUniqueInt sets the "unique_int" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUniqueInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetUniqueInt(*i)
	}
	return cuo
}

// AddUniqueInt adds i to the "unique_int" field.
func (cuo *CommentUpdateOne) AddUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.AddUniqueInt(i)
	return cuo
}

// SetUniqueFloat sets the "unique_float" field.
func (cuo *CommentUpdateOne) SetUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.ResetUniqueFloat()
	cuo.mutation.SetUniqueFloat(f)
	return cuo
}

// SetNillableUniqueFloat sets the "unique_float" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUniqueFloat(f *float64) *CommentUpdateOne {
	if f != nil {
		cuo.SetUniqueFloat(*f)
	}
	return cuo
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cuo *CommentUpdateOne) AddUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.AddUniqueFloat(f)
	return cuo
}

// SetNillableInt sets the "nillable_int" field.
func (cuo *CommentUpdateOne) SetNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetNillableInt()
	cuo.mutation.SetNillableInt(i)
	return cuo
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableNillableInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetNillableInt(*i)
	}
	return cuo
}

// AddNillableInt adds i to the "nillable_int" field.
func (cuo *CommentUpdateOne) AddNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.AddNillableInt(i)
	return cuo
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cuo *CommentUpdateOne) ClearNillableInt() *CommentUpdateOne {
	cuo.mutation.ClearNillableInt()
	return cuo
}

// SetTable sets the "table" field.
func (cuo *CommentUpdateOne) SetTable(s string) *CommentUpdateOne {
	cuo.mutation.SetTable(s)
	return cuo
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableTable(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetTable(*s)
	}
	return cuo
}

// ClearTable clears the value of the "table" field.
func (cuo *CommentUpdateOne) ClearTable() *CommentUpdateOne {
	cuo.mutation.ClearTable()
	return cuo
}

// SetDir sets the "dir" field.
func (cuo *CommentUpdateOne) SetDir(s schemadir.Dir) *CommentUpdateOne {
	cuo.mutation.SetDir(s)
	return cuo
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDir(s *schemadir.Dir) *CommentUpdateOne {
	if s != nil {
		cuo.SetDir(*s)
	}
	return cuo
}

// ClearDir clears the value of the "dir" field.
func (cuo *CommentUpdateOne) ClearDir() *CommentUpdateOne {
	cuo.mutation.ClearDir()
	return cuo
}

// SetClient sets the "client" field.
func (cuo *CommentUpdateOne) SetClient(s string) *CommentUpdateOne {
	cuo.mutation.SetClient(s)
	return cuo
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableClient(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetClient(*s)
	}
	return cuo
}

// ClearClient clears the value of the "client" field.
func (cuo *CommentUpdateOne) ClearClient() *CommentUpdateOne {
	cuo.mutation.ClearClient()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CommentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UniqueInt(); ok {
		_spec.SetField(comment.FieldUniqueInt, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedUniqueInt(); ok {
		_spec.AddField(comment.FieldUniqueInt, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.UniqueFloat(); ok {
		_spec.SetField(comment.FieldUniqueFloat, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedUniqueFloat(); ok {
		_spec.AddField(comment.FieldUniqueFloat, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.NillableInt(); ok {
		_spec.SetField(comment.FieldNillableInt, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedNillableInt(); ok {
		_spec.AddField(comment.FieldNillableInt, field.TypeInt, value)
	}
	if cuo.mutation.NillableIntCleared() {
		_spec.ClearField(comment.FieldNillableInt, field.TypeInt)
	}
	if value, ok := cuo.mutation.Table(); ok {
		_spec.SetField(comment.FieldTable, field.TypeString, value)
	}
	if cuo.mutation.TableCleared() {
		_spec.ClearField(comment.FieldTable, field.TypeString)
	}
	if value, ok := cuo.mutation.Dir(); ok {
		_spec.SetField(comment.FieldDir, field.TypeJSON, value)
	}
	if cuo.mutation.DirCleared() {
		_spec.ClearField(comment.FieldDir, field.TypeJSON)
	}
	if value, ok := cuo.mutation.GetClient(); ok {
		_spec.SetField(comment.FieldClient, field.TypeString, value)
	}
	if cuo.mutation.ClientCleared() {
		_spec.ClearField(comment.FieldClient, field.TypeString)
	}
	_spec.AddModifiers(cuo.modifiers...)
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}