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
	"github.com/usalko/fluent/flc/integration/fluent/api"
	"github.com/usalko/fluent/schema/field"
)

// APICreate is the builder for creating a Api entity.
type APICreate struct {
	config
	mutation *APIMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the APIMutation object of the builder.
func (a *APICreate) Mutation() *APIMutation {
	return a.mutation
}

// Save creates the Api in the database.
func (a *APICreate) Save(ctx context.Context) (*Api, error) {
	return withHooks(ctx, a.sqlSave, a.mutation, a.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (a *APICreate) SaveX(ctx context.Context) *Api {
	v, err := a.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (a *APICreate) Exec(ctx context.Context) error {
	_, err := a.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (a *APICreate) ExecX(ctx context.Context) {
	if err := a.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (a *APICreate) check() error {
	return nil
}

func (a *APICreate) sqlSave(ctx context.Context) (*Api, error) {
	if err := a.check(); err != nil {
		return nil, err
	}
	_node, _spec := a.createSpec()
	if err := sqlgraph.CreateNode(ctx, a.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	a.mutation.id = &_node.ID
	a.mutation.done = true
	return _node, nil
}

func (a *APICreate) createSpec() (*Api, *sqlgraph.CreateSpec) {
	var (
		_node = &Api{config: a.config}
		_spec = sqlgraph.NewCreateSpec(api.Table, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt))
	)
	_spec.OnConflict = a.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Api.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (a *APICreate) OnConflict(opts ...sql.ConflictOption) *ApiUpsertOne {
	a.conflict = opts
	return &ApiUpsertOne{
		create: a,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Api.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (a *APICreate) OnConflictColumns(columns ...string) *ApiUpsertOne {
	a.conflict = append(a.conflict, sql.ConflictColumns(columns...))
	return &ApiUpsertOne{
		create: a,
	}
}

type (
	// ApiUpsertOne is the builder for "upsert"-ing
	//  one Api node.
	ApiUpsertOne struct {
		create *APICreate
	}

	// ApiUpsert is the "OnConflict" setter.
	ApiUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Api.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ApiUpsertOne) UpdateNewValues() *ApiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Api.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ApiUpsertOne) Ignore() *ApiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiUpsertOne) DoNothing() *ApiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the APICreate.OnConflict
// documentation for more info.
func (u *ApiUpsertOne) Update(set func(*ApiUpsert)) *ApiUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ApiUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("fluent: missing options for APICreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ApiUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ApiUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// APICreateBulk is the builder for creating many Api entities in bulk.
type APICreateBulk struct {
	config
	err      error
	builders []*APICreate
	conflict []sql.ConflictOption
}

// Save creates the Api entities in the database.
func (ab *APICreateBulk) Save(ctx context.Context) ([]*Api, error) {
	if ab.err != nil {
		return nil, ab.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ab.builders))
	nodes := make([]*Api, len(ab.builders))
	mutators := make([]Mutator, len(ab.builders))
	for i := range ab.builders {
		func(i int, root context.Context) {
			builder := ab.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*APIMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ab.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ab.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ab.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ab.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ab *APICreateBulk) SaveX(ctx context.Context) []*Api {
	v, err := ab.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ab *APICreateBulk) Exec(ctx context.Context) error {
	_, err := ab.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ab *APICreateBulk) ExecX(ctx context.Context) {
	if err := ab.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Api.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (ab *APICreateBulk) OnConflict(opts ...sql.ConflictOption) *ApiUpsertBulk {
	ab.conflict = opts
	return &ApiUpsertBulk{
		create: ab,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Api.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ab *APICreateBulk) OnConflictColumns(columns ...string) *ApiUpsertBulk {
	ab.conflict = append(ab.conflict, sql.ConflictColumns(columns...))
	return &ApiUpsertBulk{
		create: ab,
	}
}

// ApiUpsertBulk is the builder for "upsert"-ing
// a bulk of Api nodes.
type ApiUpsertBulk struct {
	create *APICreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Api.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ApiUpsertBulk) UpdateNewValues() *ApiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Api.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ApiUpsertBulk) Ignore() *ApiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ApiUpsertBulk) DoNothing() *ApiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the APICreateBulk.OnConflict
// documentation for more info.
func (u *ApiUpsertBulk) Update(set func(*ApiUpsert)) *ApiUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ApiUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ApiUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("fluent: OnConflict was set for builder %d. Set it on the APICreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("fluent: missing options for APICreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ApiUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
