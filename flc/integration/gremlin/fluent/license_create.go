// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"time"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/license"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (lc *LicenseCreate) SetCreateTime(t time.Time) *LicenseCreate {
	lc.mutation.SetCreateTime(t)
	return lc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableCreateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetCreateTime(*t)
	}
	return lc
}

// SetUpdateTime sets the "update_time" field.
func (lc *LicenseCreate) SetUpdateTime(t time.Time) *LicenseCreate {
	lc.mutation.SetUpdateTime(t)
	return lc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableUpdateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetUpdateTime(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LicenseCreate) SetID(i int) *LicenseCreate {
	lc.mutation.SetID(i)
	return lc
}

// Mutation returns the LicenseMutation object of the builder.
func (lc *LicenseCreate) Mutation() *LicenseMutation {
	return lc.mutation
}

// Save creates the License in the database.
func (lc *LicenseCreate) Save(ctx context.Context) (*License, error) {
	lc.defaults()
	return withHooks(ctx, lc.gremlinSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LicenseCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LicenseCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LicenseCreate) defaults() {
	if _, ok := lc.mutation.CreateTime(); !ok {
		v := license.DefaultCreateTime()
		lc.mutation.SetCreateTime(v)
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		v := license.DefaultUpdateTime()
		lc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LicenseCreate) check() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`fluent: missing required field "License.create_time"`)}
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`fluent: missing required field "License.update_time"`)}
	}
	return nil
}

func (lc *LicenseCreate) gremlinSave(ctx context.Context) (*License, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := lc.gremlin().Query()
	if err := lc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &License{config: lc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	lc.mutation.id = &rnode.ID
	lc.mutation.done = true
	return rnode, nil
}

func (lc *LicenseCreate) gremlin() *dsl.Traversal {
	v := g.AddV(license.Label)
	if id, ok := lc.mutation.ID(); ok {
		v.Property(dsl.ID, id)
	}
	if value, ok := lc.mutation.CreateTime(); ok {
		v.Property(dsl.Single, license.FieldCreateTime, value)
	}
	if value, ok := lc.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, license.FieldUpdateTime, value)
	}
	return v.ValueMap(true)
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	err      error
	builders []*LicenseCreate
}
