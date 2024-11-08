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
	"github.com/usalko/fluent/flc/integration/hooks/fluent/card"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/pet"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/predicate"
	"github.com/usalko/fluent/flc/integration/hooks/fluent/user"
	"github.com/usalko/fluent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetVersion sets the "version" field.
func (uu *UserUpdate) SetVersion(i int) *UserUpdate {
	uu.mutation.ResetVersion()
	uu.mutation.SetVersion(i)
	return uu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVersion(i *int) *UserUpdate {
	if i != nil {
		uu.SetVersion(*i)
	}
	return uu
}

// AddVersion adds i to the "version" field.
func (uu *UserUpdate) AddVersion(i int) *UserUpdate {
	uu.mutation.AddVersion(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetWorth sets the "worth" field.
func (uu *UserUpdate) SetWorth(u uint) *UserUpdate {
	uu.mutation.ResetWorth()
	uu.mutation.SetWorth(u)
	return uu
}

// SetNillableWorth sets the "worth" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWorth(u *uint) *UserUpdate {
	if u != nil {
		uu.SetWorth(*u)
	}
	return uu
}

// AddWorth adds u to the "worth" field.
func (uu *UserUpdate) AddWorth(u int) *UserUpdate {
	uu.mutation.AddWorth(u)
	return uu
}

// ClearWorth clears the value of the "worth" field.
func (uu *UserUpdate) ClearWorth() *UserUpdate {
	uu.mutation.ClearWorth()
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetActive sets the "active" field.
func (uu *UserUpdate) SetActive(b bool) *UserUpdate {
	uu.mutation.SetActive(b)
	return uu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetActive(*b)
	}
	return uu
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (uu *UserUpdate) AddCardIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCardIDs(ids...)
	return uu
}

// AddCards adds the "cards" edges to the Card entity.
func (uu *UserUpdate) AddCards(c ...*Card) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCardIDs(ids...)
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPetIDs(ids...)
	return uu
}

// AddPets adds the "pets" edges to the Pet entity.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (uu *UserUpdate) AddFriendIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFriendIDs(ids...)
	return uu
}

// AddFriends adds the "friends" edges to the User entity.
func (uu *UserUpdate) AddFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFriendIDs(ids...)
}

// SetBestFriendID sets the "best_friend" edge to the User entity by ID.
func (uu *UserUpdate) SetBestFriendID(id int) *UserUpdate {
	uu.mutation.SetBestFriendID(id)
	return uu
}

// SetNillableBestFriendID sets the "best_friend" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableBestFriendID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetBestFriendID(*id)
	}
	return uu
}

// SetBestFriend sets the "best_friend" edge to the User entity.
func (uu *UserUpdate) SetBestFriend(u *User) *UserUpdate {
	return uu.SetBestFriendID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (uu *UserUpdate) ClearCards() *UserUpdate {
	uu.mutation.ClearCards()
	return uu
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (uu *UserUpdate) RemoveCardIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCardIDs(ids...)
	return uu
}

// RemoveCards removes "cards" edges to Card entities.
func (uu *UserUpdate) RemoveCards(c ...*Card) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCardIDs(ids...)
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.mutation.ClearPets()
	return uu
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uu *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePetIDs(ids...)
	return uu
}

// RemovePets removes "pets" edges to Pet entities.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
}

// ClearFriends clears all "friends" edges to the User entity.
func (uu *UserUpdate) ClearFriends() *UserUpdate {
	uu.mutation.ClearFriends()
	return uu
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFriendIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFriendIDs(ids...)
	return uu
}

// RemoveFriends removes "friends" edges to User entities.
func (uu *UserUpdate) RemoveFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFriendIDs(ids...)
}

// ClearBestFriend clears the "best_friend" edge to the User entity.
func (uu *UserUpdate) ClearBestFriend() *UserUpdate {
	uu.mutation.ClearBestFriend()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedVersion(); ok {
		_spec.AddField(user.FieldVersion, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Worth(); ok {
		_spec.SetField(user.FieldWorth, field.TypeUint, value)
	}
	if value, ok := uu.mutation.AddedWorth(); ok {
		_spec.AddField(user.FieldWorth, field.TypeUint, value)
	}
	if uu.mutation.WorthCleared() {
		_spec.ClearField(user.FieldWorth, field.TypeUint)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if uu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCardsIDs(); len(nodes) > 0 && !uu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !uu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.BestFriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetVersion sets the "version" field.
func (uuo *UserUpdateOne) SetVersion(i int) *UserUpdateOne {
	uuo.mutation.ResetVersion()
	uuo.mutation.SetVersion(i)
	return uuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVersion(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetVersion(*i)
	}
	return uuo
}

// AddVersion adds i to the "version" field.
func (uuo *UserUpdateOne) AddVersion(i int) *UserUpdateOne {
	uuo.mutation.AddVersion(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetWorth sets the "worth" field.
func (uuo *UserUpdateOne) SetWorth(u uint) *UserUpdateOne {
	uuo.mutation.ResetWorth()
	uuo.mutation.SetWorth(u)
	return uuo
}

// SetNillableWorth sets the "worth" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWorth(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetWorth(*u)
	}
	return uuo
}

// AddWorth adds u to the "worth" field.
func (uuo *UserUpdateOne) AddWorth(u int) *UserUpdateOne {
	uuo.mutation.AddWorth(u)
	return uuo
}

// ClearWorth clears the value of the "worth" field.
func (uuo *UserUpdateOne) ClearWorth() *UserUpdateOne {
	uuo.mutation.ClearWorth()
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetActive sets the "active" field.
func (uuo *UserUpdateOne) SetActive(b bool) *UserUpdateOne {
	uuo.mutation.SetActive(b)
	return uuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetActive(*b)
	}
	return uuo
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (uuo *UserUpdateOne) AddCardIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCardIDs(ids...)
	return uuo
}

// AddCards adds the "cards" edges to the Card entity.
func (uuo *UserUpdateOne) AddCards(c ...*Card) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCardIDs(ids...)
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPetIDs(ids...)
	return uuo
}

// AddPets adds the "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFriendIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFriendIDs(ids...)
	return uuo
}

// AddFriends adds the "friends" edges to the User entity.
func (uuo *UserUpdateOne) AddFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFriendIDs(ids...)
}

// SetBestFriendID sets the "best_friend" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetBestFriendID(id int) *UserUpdateOne {
	uuo.mutation.SetBestFriendID(id)
	return uuo
}

// SetNillableBestFriendID sets the "best_friend" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBestFriendID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetBestFriendID(*id)
	}
	return uuo
}

// SetBestFriend sets the "best_friend" edge to the User entity.
func (uuo *UserUpdateOne) SetBestFriend(u *User) *UserUpdateOne {
	return uuo.SetBestFriendID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (uuo *UserUpdateOne) ClearCards() *UserUpdateOne {
	uuo.mutation.ClearCards()
	return uuo
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (uuo *UserUpdateOne) RemoveCardIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCardIDs(ids...)
	return uuo
}

// RemoveCards removes "cards" edges to Card entities.
func (uuo *UserUpdateOne) RemoveCards(c ...*Card) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCardIDs(ids...)
}

// ClearPets clears all "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.mutation.ClearPets()
	return uuo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePetIDs(ids...)
	return uuo
}

// RemovePets removes "pets" edges to Pet entities.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
}

// ClearFriends clears all "friends" edges to the User entity.
func (uuo *UserUpdateOne) ClearFriends() *UserUpdateOne {
	uuo.mutation.ClearFriends()
	return uuo
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFriendIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFriendIDs(ids...)
	return uuo
}

// RemoveFriends removes "friends" edges to User entities.
func (uuo *UserUpdateOne) RemoveFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFriendIDs(ids...)
}

// ClearBestFriend clears the "best_friend" edge to the User entity.
func (uuo *UserUpdateOne) ClearBestFriend() *UserUpdateOne {
	uuo.mutation.ClearBestFriend()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedVersion(); ok {
		_spec.AddField(user.FieldVersion, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Worth(); ok {
		_spec.SetField(user.FieldWorth, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.AddedWorth(); ok {
		_spec.AddField(user.FieldWorth, field.TypeUint, value)
	}
	if uuo.mutation.WorthCleared() {
		_spec.ClearField(user.FieldWorth, field.TypeUint)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.Active(); ok {
		_spec.SetField(user.FieldActive, field.TypeBool, value)
	}
	if uuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCardsIDs(); len(nodes) > 0 && !uuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CardsTable,
			Columns: []string{user.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !uuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.BestFriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
