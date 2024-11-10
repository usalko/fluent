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
	"github.com/usalko/fluent/examples/edge_index/fluent/city"
	"github.com/usalko/fluent/examples/edge_index/fluent/predicate"
	"github.com/usalko/fluent/examples/edge_index/fluent/street"
	"github.com/usalko/fluent/schema/field"
)

// CityUpdate is the builder for updating City entities.
type CityUpdate struct {
	config
	hooks    []Hook
	mutation *CityMutation
}

// Where appends a list predicates to the CityUpdate builder.
func (cu *CityUpdate) Where(ps ...predicate.City) *CityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CityUpdate) SetName(s string) *CityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CityUpdate) SetNillableName(s *string) *CityUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (cu *CityUpdate) AddStreetIDs(ids ...int) *CityUpdate {
	cu.mutation.AddStreetIDs(ids...)
	return cu
}

// AddStreets adds the "streets" edges to the Street entity.
func (cu *CityUpdate) AddStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cu *CityUpdate) Mutation() *CityMutation {
	return cu.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (cu *CityUpdate) ClearStreets() *CityUpdate {
	cu.mutation.ClearStreets()
	return cu
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (cu *CityUpdate) RemoveStreetIDs(ids ...int) *CityUpdate {
	cu.mutation.RemoveStreetIDs(ids...)
	return cu
}

// RemoveStreets removes "streets" edges to Street entities.
func (cu *CityUpdate) RemoveStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(city.Table, city.Columns, sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(city.FieldName, field.TypeString, value)
	}
	if cu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !cu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CityUpdateOne is the builder for updating a single City entity.
type CityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CityMutation
}

// SetName sets the "name" field.
func (cuo *CityUpdateOne) SetName(s string) *CityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CityUpdateOne) SetNillableName(s *string) *CityUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (cuo *CityUpdateOne) AddStreetIDs(ids ...int) *CityUpdateOne {
	cuo.mutation.AddStreetIDs(ids...)
	return cuo
}

// AddStreets adds the "streets" edges to the Street entity.
func (cuo *CityUpdateOne) AddStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cuo *CityUpdateOne) Mutation() *CityMutation {
	return cuo.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (cuo *CityUpdateOne) ClearStreets() *CityUpdateOne {
	cuo.mutation.ClearStreets()
	return cuo
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (cuo *CityUpdateOne) RemoveStreetIDs(ids ...int) *CityUpdateOne {
	cuo.mutation.RemoveStreetIDs(ids...)
	return cuo
}

// RemoveStreets removes "streets" edges to Street entities.
func (cuo *CityUpdateOne) RemoveStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStreetIDs(ids...)
}

// Where appends a list predicates to the CityUpdate builder.
func (cuo *CityUpdateOne) Where(ps ...predicate.City) *CityUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CityUpdateOne) Select(field string, fields ...string) *CityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated City entity.
func (cuo *CityUpdateOne) Save(ctx context.Context) (*City, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CityUpdateOne) SaveX(ctx context.Context) *City {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CityUpdateOne) sqlSave(ctx context.Context) (_node *City, err error) {
	_spec := sqlgraph.NewUpdateSpec(city.Table, city.Columns, sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "City.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, city.FieldID)
		for _, f := range fields {
			if !city.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("fluent: invalid field %q for query", f)}
			}
			if f != city.FieldID {
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
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(city.FieldName, field.TypeString, value)
	}
	if cuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !cuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &City{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
