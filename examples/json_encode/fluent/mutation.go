// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package fluent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/usalko/fluent"
	"github.com/usalko/fluent/dialect/sql"
	"github.com/usalko/fluent/examples/json_encode/fluent/card"
	"github.com/usalko/fluent/examples/json_encode/fluent/pet"
	"github.com/usalko/fluent/examples/json_encode/fluent/predicate"
	"github.com/usalko/fluent/examples/json_encode/fluent/user"
)

const (
	// Operation types.
	OpCreate    = fluent.OpCreate
	OpDelete    = fluent.OpDelete
	OpDeleteOne = fluent.OpDeleteOne
	OpUpdate    = fluent.OpUpdate
	OpUpdateOne = fluent.OpUpdateOne

	// Node types.
	TypeCard = "Card"
	TypePet  = "Pet"
	TypeUser = "User"
)

// CardMutation represents an operation that mutates the Card nodes in the graph.
type CardMutation struct {
	config
	op            Op
	typ           string
	id            *int
	number        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Card, error)
	predicates    []predicate.Card
}

var _ fluent.Mutation = (*CardMutation)(nil)

// cardOption allows management of the mutation configuration using functional options.
type cardOption func(*CardMutation)

// newCardMutation creates new mutation for the Card entity.
func newCardMutation(c config, op Op, opts ...cardOption) *CardMutation {
	m := &CardMutation{
		config:        c,
		op:            op,
		typ:           TypeCard,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCardID sets the ID field of the mutation.
func withCardID(id int) cardOption {
	return func(m *CardMutation) {
		var (
			err   error
			once  sync.Once
			value *Card
		)
		m.oldValue = func(ctx context.Context) (*Card, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Card.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCard sets the old Card of the mutation.
func withCard(node *Card) cardOption {
	return func(m *CardMutation) {
		m.oldValue = func(context.Context) (*Card, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (fluent.Tx), a transactional client is returned.
func (m CardMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CardMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CardMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CardMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Card.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetNumber sets the "number" field.
func (m *CardMutation) SetNumber(s string) {
	m.number = &s
}

// Number returns the value of the "number" field in the mutation.
func (m *CardMutation) Number() (r string, exists bool) {
	v := m.number
	if v == nil {
		return
	}
	return *v, true
}

// OldNumber returns the old "number" field's value of the Card entity.
// If the Card object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CardMutation) OldNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNumber: %w", err)
	}
	return oldValue.Number, nil
}

// ResetNumber resets all changes to the "number" field.
func (m *CardMutation) ResetNumber() {
	m.number = nil
}

// Where appends a list predicates to the CardMutation builder.
func (m *CardMutation) Where(ps ...predicate.Card) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CardMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CardMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Card, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CardMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CardMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Card).
func (m *CardMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CardMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.number != nil {
		fields = append(fields, card.FieldNumber)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CardMutation) Field(name string) (fluent.Value, bool) {
	switch name {
	case card.FieldNumber:
		return m.Number()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CardMutation) OldField(ctx context.Context, name string) (fluent.Value, error) {
	switch name {
	case card.FieldNumber:
		return m.OldNumber(ctx)
	}
	return nil, fmt.Errorf("unknown Card field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) SetField(name string, value fluent.Value) error {
	switch name {
	case card.FieldNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNumber(v)
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CardMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CardMutation) AddedField(name string) (fluent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CardMutation) AddField(name string, value fluent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Card numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CardMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CardMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CardMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Card nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CardMutation) ResetField(name string) error {
	switch name {
	case card.FieldNumber:
		m.ResetNumber()
		return nil
	}
	return fmt.Errorf("unknown Card field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CardMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CardMutation) AddedIDs(name string) []fluent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CardMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CardMutation) RemovedIDs(name string) []fluent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CardMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CardMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CardMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Card unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CardMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Card edge %s", name)
}

// PetMutation represents an operation that mutates the Pet nodes in the graph.
type PetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	owner         *int
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Pet, error)
	predicates    []predicate.Pet
}

var _ fluent.Mutation = (*PetMutation)(nil)

// petOption allows management of the mutation configuration using functional options.
type petOption func(*PetMutation)

// newPetMutation creates new mutation for the Pet entity.
func newPetMutation(c config, op Op, opts ...petOption) *PetMutation {
	m := &PetMutation{
		config:        c,
		op:            op,
		typ:           TypePet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPetID sets the ID field of the mutation.
func withPetID(id int) petOption {
	return func(m *PetMutation) {
		var (
			err   error
			once  sync.Once
			value *Pet
		)
		m.oldValue = func(ctx context.Context) (*Pet, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Pet.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPet sets the old Pet of the mutation.
func withPet(node *Pet) petOption {
	return func(m *PetMutation) {
		m.oldValue = func(context.Context) (*Pet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (fluent.Tx), a transactional client is returned.
func (m PetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PetMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Pet.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *PetMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *PetMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *PetMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *PetMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *PetMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *PetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PetMutation) ResetName() {
	m.name = nil
}

// SetOwnerID sets the "owner_id" field.
func (m *PetMutation) SetOwnerID(i int) {
	m.owner = &i
}

// OwnerID returns the value of the "owner_id" field in the mutation.
func (m *PetMutation) OwnerID() (r int, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old "owner_id" field's value of the Pet entity.
// If the Pet object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PetMutation) OldOwnerID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ResetOwnerID resets all changes to the "owner_id" field.
func (m *PetMutation) ResetOwnerID() {
	m.owner = nil
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *PetMutation) ClearOwner() {
	m.clearedowner = true
	m.clearedFields[pet.FieldOwnerID] = struct{}{}
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *PetMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// OwnerID instead. It exists only for internal usage by the builders.
func (m *PetMutation) OwnerIDs() (ids []int) {
	if id := m.owner; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *PetMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Where appends a list predicates to the PetMutation builder.
func (m *PetMutation) Where(ps ...predicate.Pet) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PetMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PetMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Pet, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PetMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PetMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Pet).
func (m *PetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PetMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.age != nil {
		fields = append(fields, pet.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, pet.FieldName)
	}
	if m.owner != nil {
		fields = append(fields, pet.FieldOwnerID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PetMutation) Field(name string) (fluent.Value, bool) {
	switch name {
	case pet.FieldAge:
		return m.Age()
	case pet.FieldName:
		return m.Name()
	case pet.FieldOwnerID:
		return m.OwnerID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PetMutation) OldField(ctx context.Context, name string) (fluent.Value, error) {
	switch name {
	case pet.FieldAge:
		return m.OldAge(ctx)
	case pet.FieldName:
		return m.OldName(ctx)
	case pet.FieldOwnerID:
		return m.OldOwnerID(ctx)
	}
	return nil, fmt.Errorf("unknown Pet field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PetMutation) SetField(name string, value fluent.Value) error {
	switch name {
	case pet.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case pet.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case pet.FieldOwnerID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PetMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, pet.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PetMutation) AddedField(name string) (fluent.Value, bool) {
	switch name {
	case pet.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PetMutation) AddField(name string, value fluent.Value) error {
	switch name {
	case pet.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Pet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Pet nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PetMutation) ResetField(name string) error {
	switch name {
	case pet.FieldAge:
		m.ResetAge()
		return nil
	case pet.FieldName:
		m.ResetName()
		return nil
	case pet.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	}
	return fmt.Errorf("unknown Pet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, pet.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PetMutation) AddedIDs(name string) []fluent.Value {
	switch name {
	case pet.EdgeOwner:
		if id := m.owner; id != nil {
			return []fluent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PetMutation) RemovedIDs(name string) []fluent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, pet.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PetMutation) EdgeCleared(name string) bool {
	switch name {
	case pet.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PetMutation) ClearEdge(name string) error {
	switch name {
	case pet.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Pet unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PetMutation) ResetEdge(name string) error {
	switch name {
	case pet.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Pet edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	pets          map[int]struct{}
	removedpets   map[int]struct{}
	clearedpets   bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ fluent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (fluent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddPetIDs adds the "pets" edge to the Pet entity by ids.
func (m *UserMutation) AddPetIDs(ids ...int) {
	if m.pets == nil {
		m.pets = make(map[int]struct{})
	}
	for i := range ids {
		m.pets[ids[i]] = struct{}{}
	}
}

// ClearPets clears the "pets" edge to the Pet entity.
func (m *UserMutation) ClearPets() {
	m.clearedpets = true
}

// PetsCleared reports if the "pets" edge to the Pet entity was cleared.
func (m *UserMutation) PetsCleared() bool {
	return m.clearedpets
}

// RemovePetIDs removes the "pets" edge to the Pet entity by IDs.
func (m *UserMutation) RemovePetIDs(ids ...int) {
	if m.removedpets == nil {
		m.removedpets = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.pets, ids[i])
		m.removedpets[ids[i]] = struct{}{}
	}
}

// RemovedPets returns the removed IDs of the "pets" edge to the Pet entity.
func (m *UserMutation) RemovedPetsIDs() (ids []int) {
	for id := range m.removedpets {
		ids = append(ids, id)
	}
	return
}

// PetsIDs returns the "pets" edge IDs in the mutation.
func (m *UserMutation) PetsIDs() (ids []int) {
	for id := range m.pets {
		ids = append(ids, id)
	}
	return
}

// ResetPets resets all changes to the "pets" edge.
func (m *UserMutation) ResetPets() {
	m.pets = nil
	m.clearedpets = false
	m.removedpets = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (fluent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (fluent.Value, error) {
	switch name {
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value fluent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (fluent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value fluent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.pets != nil {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []fluent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]fluent.Value, 0, len(m.pets))
		for id := range m.pets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedpets != nil {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []fluent.Value {
	switch name {
	case user.EdgePets:
		ids := make([]fluent.Value, 0, len(m.removedpets))
		for id := range m.removedpets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpets {
		edges = append(edges, user.EdgePets)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePets:
		return m.clearedpets
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePets:
		m.ResetPets()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
