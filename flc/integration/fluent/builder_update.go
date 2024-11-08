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
	"github.com/usalko/fluent/flc/integration/fluent/builder"
	"github.com/usalko/fluent/flc/integration/fluent/predicate"
	"github.com/usalko/fluent/schema/field"
)

// BuilderUpdate is the builder for updating Builder entities.
type BuilderUpdate struct {
	config
	hooks     []Hook
	mutation  *BuilderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the BuilderUpdate builder.
func (bu *BuilderUpdate) Where(ps ...predicate.Builder) *BuilderUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// Mutation returns the BuilderMutation object of the builder.
func (bu *BuilderUpdate) Mutation() *BuilderMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BuilderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BuilderUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BuilderUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BuilderUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bu *BuilderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BuilderUpdate {
	bu.modifiers = append(bu.modifiers, modifiers...)
	return bu
}

func (bu *BuilderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(builder.Table, builder.Columns, sqlgraph.NewFieldSpec(builder.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(bu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{builder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BuilderUpdateOne is the builder for updating a single Builder entity.
type BuilderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *BuilderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Mutation returns the BuilderMutation object of the builder.
func (buo *BuilderUpdateOne) Mutation() *BuilderMutation {
	return buo.mutation
}

// Where appends a list predicates to the BuilderUpdate builder.
func (buo *BuilderUpdateOne) Where(ps ...predicate.Builder) *BuilderUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BuilderUpdateOne) Select(field string, fields ...string) *BuilderUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Builder entity.
func (buo *BuilderUpdateOne) Save(ctx context.Context) (*Builder, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BuilderUpdateOne) SaveX(ctx context.Context) *Builder {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BuilderUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BuilderUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (buo *BuilderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BuilderUpdateOne {
	buo.modifiers = append(buo.modifiers, modifiers...)
	return buo
}

func (buo *BuilderUpdateOne) sqlSave(ctx context.Context) (_node *Builder, err error) {
	_spec := sqlgraph.NewUpdateSpec(builder.Table, builder.Columns, sqlgraph.NewFieldSpec(builder.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Builder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, builder.FieldID)
		for _, f := range fields {
			if !builder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != builder.FieldID {
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
	_spec.AddModifiers(buo.modifiers...)
	_node = &Builder{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{builder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
