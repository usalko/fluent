// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/dialect/sql/sqlgraph"
	"github.com/usalko/fluent/flc/integration/custom_id/fluent/mixin_id"
	"github.com/usalko/fluent/schema/field"
)

// MixinIDCreate is the builder for creating a MixinID entity.
type MixinIDCreate struct {
	config
	mutation *MixinIDMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSomeField sets the "some_field" field.
func (mi *MixinIDCreate) SetSomeField(s string) *MixinIDCreate {
	mi.mutation.SetSomeField(s)
	return mi
}

// SetMixinField sets the "mixin_field" field.
func (mi *MixinIDCreate) SetMixinField(s string) *MixinIDCreate {
	mi.mutation.SetMixinField(s)
	return mi
}

// SetID sets the "id" field.
func (mi *MixinIDCreate) SetID(uu uuid.UUID) *MixinIDCreate {
	mi.mutation.SetID(uu)
	return mi
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mi *MixinIDCreate) SetNillableID(uu *uuid.UUID) *MixinIDCreate {
	if uu != nil {
		mi.SetID(*uu)
	}
	return mi
}

// Mutation returns the MixinIDMutation object of the builder.
func (mi *MixinIDCreate) Mutation() *MixinIDMutation {
	return mi.mutation
}

// Save creates the MixinID in the database.
func (mi *MixinIDCreate) Save(ctx context.Context) (*MixinID, error) {
	mi.defaults()
	return withHooks(ctx, mi.sqlSave, mi.mutation, mi.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mi *MixinIDCreate) SaveX(ctx context.Context) *MixinID {
	v, err := mi.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mi *MixinIDCreate) Exec(ctx context.Context) error {
	_, err := mi.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mi *MixinIDCreate) ExecX(ctx context.Context) {
	if err := mi.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mi *MixinIDCreate) defaults() {
	if _, ok := mi.mutation.ID(); !ok {
		v := mixin_id.DefaultID()
		mi.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mi *MixinIDCreate) check() error {
	if _, ok := mi.mutation.SomeField(); !ok {
		return &ValidationError{Name: "some_field", err: errors.New(`fluent: missing required field "MixinID.some_field"`)}
	}
	if _, ok := mi.mutation.MixinField(); !ok {
		return &ValidationError{Name: "mixin_field", err: errors.New(`fluent: missing required field "MixinID.mixin_field"`)}
	}
	return nil
}

func (mi *MixinIDCreate) sqlSave(ctx context.Context) (*MixinID, error) {
	if err := mi.check(); err != nil {
		return nil, err
	}
	_node, _spec := mi.createSpec()
	if err := sqlgraph.CreateNode(ctx, mi.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mi.mutation.id = &_node.ID
	mi.mutation.done = true
	return _node, nil
}

func (mi *MixinIDCreate) createSpec() (*MixinID, *sqlgraph.CreateSpec) {
	var (
		_node = &MixinID{config: mi.config}
		_spec = sqlgraph.NewCreateSpec(mixin_id.Table, sqlgraph.NewFieldSpec(mixin_id.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = mi.conflict
	if id, ok := mi.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mi.mutation.SomeField(); ok {
		_spec.SetField(mixin_id.FieldSomeField, field.TypeString, value)
		_node.SomeField = value
	}
	if value, ok := mi.mutation.MixinField(); ok {
		_spec.SetField(mixin_id.FieldMixinField, field.TypeString, value)
		_node.MixinField = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MixinID.Create().
//		SetSomeField(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *fluent.MixinIDUpsert) {
//			SetSomeField(v+v).
//		}).
//		Exec(ctx)
func (mi *MixinIDCreate) OnConflict(opts ...sql.ConflictOption) *MixinIDUpsertOne {
	mi.conflict = opts
	return &MixinIDUpsertOne{
		create: mi,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mi *MixinIDCreate) OnConflictColumns(columns ...string) *MixinIDUpsertOne {
	mi.conflict = append(mi.conflict, sql.ConflictColumns(columns...))
	return &MixinIDUpsertOne{
		create: mi,
	}
}

type (
	// MixinIDUpsertOne is the builder for "upsert"-ing
	//  one MixinID node.
	MixinIDUpsertOne struct {
		create *MixinIDCreate
	}

	// MixinIDUpsert is the "OnConflict" setter.
	MixinIDUpsert struct {
		*sql.UpdateSet
	}
)

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsert) SetSomeField(v string) *MixinIDUpsert {
	u.Set(mixin_id.FieldSomeField, v)
	return u
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsert) UpdateSomeField() *MixinIDUpsert {
	u.SetExcluded(mixin_id.FieldSomeField)
	return u
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsert) SetMixinField(v string) *MixinIDUpsert {
	u.Set(mixin_id.FieldMixinField, v)
	return u
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsert) UpdateMixinField() *MixinIDUpsert {
	u.SetExcluded(mixin_id.FieldMixinField)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mixin_id.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MixinIDUpsertOne) UpdateNewValues() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mixin_id.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MixinID.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MixinIDUpsertOne) Ignore() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MixinIDUpsertOne) DoNothing() *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MixinIDCreate.OnConflict
// documentation for more info.
func (u *MixinIDUpsertOne) Update(set func(*MixinIDUpsert)) *MixinIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MixinIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsertOne) SetSomeField(v string) *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetSomeField(v)
	})
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsertOne) UpdateSomeField() *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateSomeField()
	})
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsertOne) SetMixinField(v string) *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetMixinField(v)
	})
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsertOne) UpdateMixinField() *MixinIDUpsertOne {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateMixinField()
	})
}

// Exec executes the query.
func (u *MixinIDUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("fluent: missing options for MixinIDCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MixinIDUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MixinIDUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("fluent: MixinIDUpsertOne.ID is not supported by MySQL driver. Use MixinIDUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MixinIDUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MixinIDCreateBulk is the builder for creating many MixinID entities in bulk.
type MixinIDCreateBulk struct {
	config
	err      error
	builders []*MixinIDCreate
	conflict []sql.ConflictOption
}

// Save creates the MixinID entities in the database.
func (mib *MixinIDCreateBulk) Save(ctx context.Context) ([]*MixinID, error) {
	if mib.err != nil {
		return nil, mib.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mib.builders))
	nodes := make([]*MixinID, len(mib.builders))
	mutators := make([]Mutator, len(mib.builders))
	for i := range mib.builders {
		func(i int, root context.Context) {
			builder := mib.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MixinIDMutation)
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
					_, err = mutators[i+1].Mutate(root, mib.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mib.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mib.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mib.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mib *MixinIDCreateBulk) SaveX(ctx context.Context) []*MixinID {
	v, err := mib.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mib *MixinIDCreateBulk) Exec(ctx context.Context) error {
	_, err := mib.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mib *MixinIDCreateBulk) ExecX(ctx context.Context) {
	if err := mib.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MixinID.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *fluent.MixinIDUpsert) {
//			SetSomeField(v+v).
//		}).
//		Exec(ctx)
func (mib *MixinIDCreateBulk) OnConflict(opts ...sql.ConflictOption) *MixinIDUpsertBulk {
	mib.conflict = opts
	return &MixinIDUpsertBulk{
		create: mib,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mib *MixinIDCreateBulk) OnConflictColumns(columns ...string) *MixinIDUpsertBulk {
	mib.conflict = append(mib.conflict, sql.ConflictColumns(columns...))
	return &MixinIDUpsertBulk{
		create: mib,
	}
}

// MixinIDUpsertBulk is the builder for "upsert"-ing
// a bulk of MixinID nodes.
type MixinIDUpsertBulk struct {
	create *MixinIDCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mixin_id.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MixinIDUpsertBulk) UpdateNewValues() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mixin_id.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MixinID.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MixinIDUpsertBulk) Ignore() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MixinIDUpsertBulk) DoNothing() *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MixinIDCreateBulk.OnConflict
// documentation for more info.
func (u *MixinIDUpsertBulk) Update(set func(*MixinIDUpsert)) *MixinIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MixinIDUpsert{UpdateSet: update})
	}))
	return u
}

// SetSomeField sets the "some_field" field.
func (u *MixinIDUpsertBulk) SetSomeField(v string) *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetSomeField(v)
	})
}

// UpdateSomeField sets the "some_field" field to the value that was provided on create.
func (u *MixinIDUpsertBulk) UpdateSomeField() *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateSomeField()
	})
}

// SetMixinField sets the "mixin_field" field.
func (u *MixinIDUpsertBulk) SetMixinField(v string) *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.SetMixinField(v)
	})
}

// UpdateMixinField sets the "mixin_field" field to the value that was provided on create.
func (u *MixinIDUpsertBulk) UpdateMixinField() *MixinIDUpsertBulk {
	return u.Update(func(s *MixinIDUpsert) {
		s.UpdateMixinField()
	})
}

// Exec executes the query.
func (u *MixinIDUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("fluent: OnConflict was set for builder %d. Set it on the MixinIDCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("fluent: missing options for MixinIDCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MixinIDUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}