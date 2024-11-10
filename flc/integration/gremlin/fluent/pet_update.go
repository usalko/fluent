// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/usalko/fluent/dialect/gremlin"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/__"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/g"
	"github.com/usalko/fluent/dialect/gremlin/graph/dsl/p"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/pet"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/gremlin/fluent/user"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where appends a list predicates to the PetUpdate builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAge sets the "age" field.
func (pu *PetUpdate) SetAge(f float64) *PetUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(f)
	return pu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (pu *PetUpdate) SetNillableAge(f *float64) *PetUpdate {
	if f != nil {
		pu.SetAge(*f)
	}
	return pu
}

// AddAge adds f to the "age" field.
func (pu *PetUpdate) AddAge(f float64) *PetUpdate {
	pu.mutation.AddAge(f)
	return pu
}

// SetName sets the "name" field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PetUpdate) SetNillableName(s *string) *PetUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetUUID sets the "uuid" field.
func (pu *PetUpdate) SetUUID(u uuid.UUID) *PetUpdate {
	pu.mutation.SetUUID(u)
	return pu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pu *PetUpdate) SetNillableUUID(u *uuid.UUID) *PetUpdate {
	if u != nil {
		pu.SetUUID(*u)
	}
	return pu
}

// ClearUUID clears the value of the "uuid" field.
func (pu *PetUpdate) ClearUUID() *PetUpdate {
	pu.mutation.ClearUUID()
	return pu
}

// SetNickname sets the "nickname" field.
func (pu *PetUpdate) SetNickname(s string) *PetUpdate {
	pu.mutation.SetNickname(s)
	return pu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (pu *PetUpdate) SetNillableNickname(s *string) *PetUpdate {
	if s != nil {
		pu.SetNickname(*s)
	}
	return pu
}

// ClearNickname clears the value of the "nickname" field.
func (pu *PetUpdate) ClearNickname() *PetUpdate {
	pu.mutation.ClearNickname()
	return pu
}

// SetTrained sets the "trained" field.
func (pu *PetUpdate) SetTrained(b bool) *PetUpdate {
	pu.mutation.SetTrained(b)
	return pu
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (pu *PetUpdate) SetNillableTrained(b *bool) *PetUpdate {
	if b != nil {
		pu.SetTrained(*b)
	}
	return pu
}

// SetOptionalTime sets the "optional_time" field.
func (pu *PetUpdate) SetOptionalTime(t time.Time) *PetUpdate {
	pu.mutation.SetOptionalTime(t)
	return pu
}

// SetNillableOptionalTime sets the "optional_time" field if the given value is not nil.
func (pu *PetUpdate) SetNillableOptionalTime(t *time.Time) *PetUpdate {
	if t != nil {
		pu.SetOptionalTime(*t)
	}
	return pu
}

// ClearOptionalTime clears the value of the "optional_time" field.
func (pu *PetUpdate) ClearOptionalTime() *PetUpdate {
	pu.mutation.ClearOptionalTime()
	return pu
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (pu *PetUpdate) SetTeamID(id string) *PetUpdate {
	pu.mutation.SetTeamID(id)
	return pu
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableTeamID(id *string) *PetUpdate {
	if id != nil {
		pu = pu.SetTeamID(*id)
	}
	return pu
}

// SetTeam sets the "team" edge to the User entity.
func (pu *PetUpdate) SetTeam(u *User) *PetUpdate {
	return pu.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *PetUpdate) SetOwnerID(id string) *PetUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *string) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pu *PetUpdate) Mutation() *PetMutation {
	return pu.mutation
}

// ClearTeam clears the "team" edge to the User entity.
func (pu *PetUpdate) ClearTeam() *PetUpdate {
	pu.mutation.ClearTeam()
	return pu
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.gremlinSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pu.gremlin().Query()
	if err := pu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	pu.mutation.done = true
	return res.ReadInt()
}

func (pu *PetUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(pet.Label)
	for _, p := range pu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := pu.mutation.Age(); ok {
		v.Property(dsl.Single, pet.FieldAge, value)
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		v.Property(dsl.Single, pet.FieldAge, __.Union(__.Values(pet.FieldAge), __.Constant(value)).Sum())
	}
	if value, ok := pu.mutation.Name(); ok {
		v.Property(dsl.Single, pet.FieldName, value)
	}
	if value, ok := pu.mutation.UUID(); ok {
		v.Property(dsl.Single, pet.FieldUUID, value)
	}
	if value, ok := pu.mutation.Nickname(); ok {
		v.Property(dsl.Single, pet.FieldNickname, value)
	}
	if value, ok := pu.mutation.Trained(); ok {
		v.Property(dsl.Single, pet.FieldTrained, value)
	}
	if value, ok := pu.mutation.OptionalTime(); ok {
		v.Property(dsl.Single, pet.FieldOptionalTime, value)
	}
	var properties []any
	if pu.mutation.UUIDCleared() {
		properties = append(properties, pet.FieldUUID)
	}
	if pu.mutation.NicknameCleared() {
		properties = append(properties, pet.FieldNickname)
	}
	if pu.mutation.OptionalTimeCleared() {
		properties = append(properties, pet.FieldOptionalTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if pu.mutation.TeamCleared() {
		tr := rv.Clone().InE(user.TeamLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range pu.mutation.TeamIDs() {
		v.AddE(user.TeamLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(pet.Label, user.TeamLabel, id)),
		})
	}
	if pu.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.PetsLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range pu.mutation.OwnerIDs() {
		v.AddE(user.PetsLabel).From(g.V(id)).InV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PetMutation
}

// SetAge sets the "age" field.
func (puo *PetUpdateOne) SetAge(f float64) *PetUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(f)
	return puo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableAge(f *float64) *PetUpdateOne {
	if f != nil {
		puo.SetAge(*f)
	}
	return puo
}

// AddAge adds f to the "age" field.
func (puo *PetUpdateOne) AddAge(f float64) *PetUpdateOne {
	puo.mutation.AddAge(f)
	return puo
}

// SetName sets the "name" field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableName(s *string) *PetUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetUUID sets the "uuid" field.
func (puo *PetUpdateOne) SetUUID(u uuid.UUID) *PetUpdateOne {
	puo.mutation.SetUUID(u)
	return puo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableUUID(u *uuid.UUID) *PetUpdateOne {
	if u != nil {
		puo.SetUUID(*u)
	}
	return puo
}

// ClearUUID clears the value of the "uuid" field.
func (puo *PetUpdateOne) ClearUUID() *PetUpdateOne {
	puo.mutation.ClearUUID()
	return puo
}

// SetNickname sets the "nickname" field.
func (puo *PetUpdateOne) SetNickname(s string) *PetUpdateOne {
	puo.mutation.SetNickname(s)
	return puo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableNickname(s *string) *PetUpdateOne {
	if s != nil {
		puo.SetNickname(*s)
	}
	return puo
}

// ClearNickname clears the value of the "nickname" field.
func (puo *PetUpdateOne) ClearNickname() *PetUpdateOne {
	puo.mutation.ClearNickname()
	return puo
}

// SetTrained sets the "trained" field.
func (puo *PetUpdateOne) SetTrained(b bool) *PetUpdateOne {
	puo.mutation.SetTrained(b)
	return puo
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableTrained(b *bool) *PetUpdateOne {
	if b != nil {
		puo.SetTrained(*b)
	}
	return puo
}

// SetOptionalTime sets the "optional_time" field.
func (puo *PetUpdateOne) SetOptionalTime(t time.Time) *PetUpdateOne {
	puo.mutation.SetOptionalTime(t)
	return puo
}

// SetNillableOptionalTime sets the "optional_time" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOptionalTime(t *time.Time) *PetUpdateOne {
	if t != nil {
		puo.SetOptionalTime(*t)
	}
	return puo
}

// ClearOptionalTime clears the value of the "optional_time" field.
func (puo *PetUpdateOne) ClearOptionalTime() *PetUpdateOne {
	puo.mutation.ClearOptionalTime()
	return puo
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (puo *PetUpdateOne) SetTeamID(id string) *PetUpdateOne {
	puo.mutation.SetTeamID(id)
	return puo
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableTeamID(id *string) *PetUpdateOne {
	if id != nil {
		puo = puo.SetTeamID(*id)
	}
	return puo
}

// SetTeam sets the "team" edge to the User entity.
func (puo *PetUpdateOne) SetTeam(u *User) *PetUpdateOne {
	return puo.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *PetUpdateOne) SetOwnerID(id string) *PetUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *string) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (puo *PetUpdateOne) Mutation() *PetMutation {
	return puo.mutation
}

// ClearTeam clears the "team" edge to the User entity.
func (puo *PetUpdateOne) ClearTeam() *PetUpdateOne {
	puo.mutation.ClearTeam()
	return puo
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// Where appends a list predicates to the PetUpdate builder.
func (puo *PetUpdateOne) Where(ps ...predicate.Pet) *PetUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pet entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, puo.gremlinSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) gremlinSave(ctx context.Context) (*Pet, error) {
	res := &gremlin.Response{}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`fluent: missing "Pet.id" for update`)}
	}
	query, bindings := puo.gremlin(id).Query()
	if err := puo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	puo.mutation.done = true
	pe := &Pet{config: puo.config}
	if err := pe.FromResponse(res); err != nil {
		return nil, err
	}
	return pe, nil
}

func (puo *PetUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := puo.mutation.Age(); ok {
		v.Property(dsl.Single, pet.FieldAge, value)
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		v.Property(dsl.Single, pet.FieldAge, __.Union(__.Values(pet.FieldAge), __.Constant(value)).Sum())
	}
	if value, ok := puo.mutation.Name(); ok {
		v.Property(dsl.Single, pet.FieldName, value)
	}
	if value, ok := puo.mutation.UUID(); ok {
		v.Property(dsl.Single, pet.FieldUUID, value)
	}
	if value, ok := puo.mutation.Nickname(); ok {
		v.Property(dsl.Single, pet.FieldNickname, value)
	}
	if value, ok := puo.mutation.Trained(); ok {
		v.Property(dsl.Single, pet.FieldTrained, value)
	}
	if value, ok := puo.mutation.OptionalTime(); ok {
		v.Property(dsl.Single, pet.FieldOptionalTime, value)
	}
	var properties []any
	if puo.mutation.UUIDCleared() {
		properties = append(properties, pet.FieldUUID)
	}
	if puo.mutation.NicknameCleared() {
		properties = append(properties, pet.FieldNickname)
	}
	if puo.mutation.OptionalTimeCleared() {
		properties = append(properties, pet.FieldOptionalTime)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if puo.mutation.TeamCleared() {
		tr := rv.Clone().InE(user.TeamLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range puo.mutation.TeamIDs() {
		v.AddE(user.TeamLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(pet.Label, user.TeamLabel, id)),
		})
	}
	if puo.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.PetsLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range puo.mutation.OwnerIDs() {
		v.AddE(user.PetsLabel).From(g.V(id)).InV()
	}
	if len(puo.fields) > 0 {
		fields := make([]any, 0, len(puo.fields)+1)
		fields = append(fields, true)
		for _, f := range puo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
