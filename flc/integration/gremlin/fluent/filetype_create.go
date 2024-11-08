// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"

	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/p"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/filetype"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	mutation *FileTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ftc *FileTypeCreate) SetName(s string) *FileTypeCreate {
	ftc.mutation.SetName(s)
	return ftc
}

// SetType sets the "type" field.
func (ftc *FileTypeCreate) SetType(f filetype.Type) *FileTypeCreate {
	ftc.mutation.SetType(f)
	return ftc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableType(f *filetype.Type) *FileTypeCreate {
	if f != nil {
		ftc.SetType(*f)
	}
	return ftc
}

// SetState sets the "state" field.
func (ftc *FileTypeCreate) SetState(f filetype.State) *FileTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableState(f *filetype.State) *FileTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftc *FileTypeCreate) AddFileIDs(ids ...string) *FileTypeCreate {
	ftc.mutation.AddFileIDs(ids...)
	return ftc
}

// AddFiles adds the "files" edges to the File entity.
func (ftc *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftc *FileTypeCreate) Mutation() *FileTypeMutation {
	return ftc.mutation
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	ftc.defaults()
	return withHooks(ctx, ftc.gremlinSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FileTypeCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FileTypeCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FileTypeCreate) defaults() {
	if _, ok := ftc.mutation.GetType(); !ok {
		v := filetype.DefaultType
		ftc.mutation.SetType(v)
	}
	if _, ok := ftc.mutation.State(); !ok {
		v := filetype.DefaultState
		ftc.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FileTypeCreate) check() error {
	if _, ok := ftc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FileType.name"`)}
	}
	if _, ok := ftc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FileType.type"`)}
	}
	if v, ok := ftc.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "FileType.state"`)}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

func (ftc *FileTypeCreate) gremlinSave(ctx context.Context) (*FileType, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &FileType{config: ftc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	ftc.mutation.id = &rnode.ID
	ftc.mutation.done = true
	return rnode, nil
}

func (ftc *FileTypeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(filetype.Label)
	if value, ok := ftc.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	if value, ok := ftc.mutation.GetType(); ok {
		v.Property(dsl.Single, filetype.FieldType, value)
	}
	if value, ok := ftc.mutation.State(); ok {
		v.Property(dsl.Single, filetype.FieldState, value)
	}
	for _, id := range ftc.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// FileTypeCreateBulk is the builder for creating many FileType entities in bulk.
type FileTypeCreateBulk struct {
	config
	err      error
	builders []*FileTypeCreate
}
