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
	"github.com/usalko/fluent/examples/edge_index/fluent/city"
	"github.com/usalko/fluent/examples/edge_index/fluent/predicate"
	"github.com/usalko/fluent/examples/edge_index/fluent/street"
)

const (
	// Operation types.
	OpCreate    = fluent.OpCreate
	OpDelete    = fluent.OpDelete
	OpDeleteOne = fluent.OpDeleteOne
	OpUpdate    = fluent.OpUpdate
	OpUpdateOne = fluent.OpUpdateOne

	// Node types.
	TypeCity   = "City"
	TypeStreet = "Street"
)

// CityMutation represents an operation that mutates the City nodes in the graph.
type CityMutation struct {
	config
	op             Op
	typ            string
	id             *int
	name           *string
	clearedFields  map[string]struct{}
	streets        map[int]struct{}
	removedstreets map[int]struct{}
	clearedstreets bool
	done           bool
	oldValue       func(context.Context) (*City, error)
	predicates     []predicate.City
}

var _ fluent.Mutation = (*CityMutation)(nil)

// cityOption allows management of the mutation configuration using functional options.
type cityOption func(*CityMutation)

// newCityMutation creates new mutation for the City entity.
func newCityMutation(c config, op Op, opts ...cityOption) *CityMutation {
	m := &CityMutation{
		config:        c,
		op:            op,
		typ:           TypeCity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCityID sets the ID field of the mutation.
func withCityID(id int) cityOption {
	return func(m *CityMutation) {
		var (
			err   error
			once  sync.Once
			value *City
		)
		m.oldValue = func(ctx context.Context) (*City, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().City.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCity sets the old City of the mutation.
func withCity(node *City) cityOption {
	return func(m *CityMutation) {
		m.oldValue = func(context.Context) (*City, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (fluent.Tx), a transactional client is returned.
func (m CityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CityMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CityMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().City.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CityMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CityMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *CityMutation) ResetName() {
	m.name = nil
}

// AddStreetIDs adds the "streets" edge to the Street entity by ids.
func (m *CityMutation) AddStreetIDs(ids ...int) {
	if m.streets == nil {
		m.streets = make(map[int]struct{})
	}
	for i := range ids {
		m.streets[ids[i]] = struct{}{}
	}
}

// ClearStreets clears the "streets" edge to the Street entity.
func (m *CityMutation) ClearStreets() {
	m.clearedstreets = true
}

// StreetsCleared reports if the "streets" edge to the Street entity was cleared.
func (m *CityMutation) StreetsCleared() bool {
	return m.clearedstreets
}

// RemoveStreetIDs removes the "streets" edge to the Street entity by IDs.
func (m *CityMutation) RemoveStreetIDs(ids ...int) {
	if m.removedstreets == nil {
		m.removedstreets = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.streets, ids[i])
		m.removedstreets[ids[i]] = struct{}{}
	}
}

// RemovedStreets returns the removed IDs of the "streets" edge to the Street entity.
func (m *CityMutation) RemovedStreetsIDs() (ids []int) {
	for id := range m.removedstreets {
		ids = append(ids, id)
	}
	return
}

// StreetsIDs returns the "streets" edge IDs in the mutation.
func (m *CityMutation) StreetsIDs() (ids []int) {
	for id := range m.streets {
		ids = append(ids, id)
	}
	return
}

// ResetStreets resets all changes to the "streets" edge.
func (m *CityMutation) ResetStreets() {
	m.streets = nil
	m.clearedstreets = false
	m.removedstreets = nil
}

// Where appends a list predicates to the CityMutation builder.
func (m *CityMutation) Where(ps ...predicate.City) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CityMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CityMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.City, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CityMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CityMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (City).
func (m *CityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CityMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, city.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CityMutation) Field(name string) (fluent.Value, bool) {
	switch name {
	case city.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CityMutation) OldField(ctx context.Context, name string) (fluent.Value, error) {
	switch name {
	case city.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown City field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) SetField(name string, value fluent.Value) error {
	switch name {
	case city.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CityMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CityMutation) AddedField(name string) (fluent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) AddField(name string, value fluent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown City numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CityMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CityMutation) ClearField(name string) error {
	return fmt.Errorf("unknown City nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CityMutation) ResetField(name string) error {
	switch name {
	case city.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CityMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.streets != nil {
		edges = append(edges, city.EdgeStreets)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CityMutation) AddedIDs(name string) []fluent.Value {
	switch name {
	case city.EdgeStreets:
		ids := make([]fluent.Value, 0, len(m.streets))
		for id := range m.streets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedstreets != nil {
		edges = append(edges, city.EdgeStreets)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CityMutation) RemovedIDs(name string) []fluent.Value {
	switch name {
	case city.EdgeStreets:
		ids := make([]fluent.Value, 0, len(m.removedstreets))
		for id := range m.removedstreets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedstreets {
		edges = append(edges, city.EdgeStreets)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CityMutation) EdgeCleared(name string) bool {
	switch name {
	case city.EdgeStreets:
		return m.clearedstreets
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CityMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown City unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CityMutation) ResetEdge(name string) error {
	switch name {
	case city.EdgeStreets:
		m.ResetStreets()
		return nil
	}
	return fmt.Errorf("unknown City edge %s", name)
}

// StreetMutation represents an operation that mutates the Street nodes in the graph.
type StreetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	city          *int
	clearedcity   bool
	done          bool
	oldValue      func(context.Context) (*Street, error)
	predicates    []predicate.Street
}

var _ fluent.Mutation = (*StreetMutation)(nil)

// streetOption allows management of the mutation configuration using functional options.
type streetOption func(*StreetMutation)

// newStreetMutation creates new mutation for the Street entity.
func newStreetMutation(c config, op Op, opts ...streetOption) *StreetMutation {
	m := &StreetMutation{
		config:        c,
		op:            op,
		typ:           TypeStreet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStreetID sets the ID field of the mutation.
func withStreetID(id int) streetOption {
	return func(m *StreetMutation) {
		var (
			err   error
			once  sync.Once
			value *Street
		)
		m.oldValue = func(ctx context.Context) (*Street, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Street.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStreet sets the old Street of the mutation.
func withStreet(node *Street) streetOption {
	return func(m *StreetMutation) {
		m.oldValue = func(context.Context) (*Street, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (fluent.Tx), a transactional client is returned.
func (m StreetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StreetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StreetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StreetMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Street.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *StreetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *StreetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Street entity.
// If the Street object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StreetMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *StreetMutation) ResetName() {
	m.name = nil
}

// SetCityID sets the "city" edge to the City entity by id.
func (m *StreetMutation) SetCityID(id int) {
	m.city = &id
}

// ClearCity clears the "city" edge to the City entity.
func (m *StreetMutation) ClearCity() {
	m.clearedcity = true
}

// CityCleared reports if the "city" edge to the City entity was cleared.
func (m *StreetMutation) CityCleared() bool {
	return m.clearedcity
}

// CityID returns the "city" edge ID in the mutation.
func (m *StreetMutation) CityID() (id int, exists bool) {
	if m.city != nil {
		return *m.city, true
	}
	return
}

// CityIDs returns the "city" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// CityID instead. It exists only for internal usage by the builders.
func (m *StreetMutation) CityIDs() (ids []int) {
	if id := m.city; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCity resets all changes to the "city" edge.
func (m *StreetMutation) ResetCity() {
	m.city = nil
	m.clearedcity = false
}

// Where appends a list predicates to the StreetMutation builder.
func (m *StreetMutation) Where(ps ...predicate.Street) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StreetMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StreetMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Street, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StreetMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StreetMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Street).
func (m *StreetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StreetMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, street.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StreetMutation) Field(name string) (fluent.Value, bool) {
	switch name {
	case street.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StreetMutation) OldField(ctx context.Context, name string) (fluent.Value, error) {
	switch name {
	case street.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Street field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StreetMutation) SetField(name string, value fluent.Value) error {
	switch name {
	case street.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Street field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StreetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StreetMutation) AddedField(name string) (fluent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StreetMutation) AddField(name string, value fluent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Street numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StreetMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StreetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StreetMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Street nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StreetMutation) ResetField(name string) error {
	switch name {
	case street.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Street field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StreetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.city != nil {
		edges = append(edges, street.EdgeCity)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StreetMutation) AddedIDs(name string) []fluent.Value {
	switch name {
	case street.EdgeCity:
		if id := m.city; id != nil {
			return []fluent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StreetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StreetMutation) RemovedIDs(name string) []fluent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StreetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcity {
		edges = append(edges, street.EdgeCity)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StreetMutation) EdgeCleared(name string) bool {
	switch name {
	case street.EdgeCity:
		return m.clearedcity
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StreetMutation) ClearEdge(name string) error {
	switch name {
	case street.EdgeCity:
		m.ClearCity()
		return nil
	}
	return fmt.Errorf("unknown Street unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StreetMutation) ResetEdge(name string) error {
	switch name {
	case street.EdgeCity:
		m.ResetCity()
		return nil
	}
	return fmt.Errorf("unknown Street edge %s", name)
}
