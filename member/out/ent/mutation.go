// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"todopoint/member/out/ent/member"
	"todopoint/member/out/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMember = "Member"
)

// MemberMutation represents an operation that mutates the Member nodes in the graph.
type MemberMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	username      *string
	password      *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Member, error)
	predicates    []predicate.Member
}

var _ ent.Mutation = (*MemberMutation)(nil)

// memberOption allows management of the mutation configuration using functional options.
type memberOption func(*MemberMutation)

// newMemberMutation creates new mutation for the Member entity.
func newMemberMutation(c config, op Op, opts ...memberOption) *MemberMutation {
	m := &MemberMutation{
		config:        c,
		op:            op,
		typ:           TypeMember,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMemberID sets the ID field of the mutation.
func withMemberID(id int) memberOption {
	return func(m *MemberMutation) {
		var (
			err   error
			once  sync.Once
			value *Member
		)
		m.oldValue = func(ctx context.Context) (*Member, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Member.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMember sets the old Member of the mutation.
func withMember(node *Member) memberOption {
	return func(m *MemberMutation) {
		m.oldValue = func(context.Context) (*Member, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MemberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MemberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Member entities.
func (m *MemberMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MemberMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MemberMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Member.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *MemberMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *MemberMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *MemberMutation) ResetEmail() {
	m.email = nil
}

// SetUsername sets the "username" field.
func (m *MemberMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *MemberMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ClearUsername clears the value of the "username" field.
func (m *MemberMutation) ClearUsername() {
	m.username = nil
	m.clearedFields[member.FieldUsername] = struct{}{}
}

// UsernameCleared returns if the "username" field was cleared in this mutation.
func (m *MemberMutation) UsernameCleared() bool {
	_, ok := m.clearedFields[member.FieldUsername]
	return ok
}

// ResetUsername resets all changes to the "username" field.
func (m *MemberMutation) ResetUsername() {
	m.username = nil
	delete(m.clearedFields, member.FieldUsername)
}

// SetPassword sets the "password" field.
func (m *MemberMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *MemberMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *MemberMutation) ResetPassword() {
	m.password = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *MemberMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MemberMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *MemberMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the MemberMutation builder.
func (m *MemberMutation) Where(ps ...predicate.Member) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MemberMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MemberMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Member, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MemberMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MemberMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Member).
func (m *MemberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MemberMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.email != nil {
		fields = append(fields, member.FieldEmail)
	}
	if m.username != nil {
		fields = append(fields, member.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, member.FieldPassword)
	}
	if m.created_at != nil {
		fields = append(fields, member.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case member.FieldEmail:
		return m.Email()
	case member.FieldUsername:
		return m.Username()
	case member.FieldPassword:
		return m.Password()
	case member.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case member.FieldEmail:
		return m.OldEmail(ctx)
	case member.FieldUsername:
		return m.OldUsername(ctx)
	case member.FieldPassword:
		return m.OldPassword(ctx)
	case member.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Member field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case member.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case member.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case member.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case member.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MemberMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MemberMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Member numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MemberMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(member.FieldUsername) {
		fields = append(fields, member.FieldUsername)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MemberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MemberMutation) ClearField(name string) error {
	switch name {
	case member.FieldUsername:
		m.ClearUsername()
		return nil
	}
	return fmt.Errorf("unknown Member nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MemberMutation) ResetField(name string) error {
	switch name {
	case member.FieldEmail:
		m.ResetEmail()
		return nil
	case member.FieldUsername:
		m.ResetUsername()
		return nil
	case member.FieldPassword:
		m.ResetPassword()
		return nil
	case member.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MemberMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MemberMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MemberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MemberMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MemberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MemberMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MemberMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Member unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MemberMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Member edge %s", name)
}