// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/member"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/message"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/predicate"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/room"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMember  = "Member"
	TypeMessage = "Message"
	TypeRoom    = "Room"
)

// MemberMutation represents an operation that mutates the Member nodes in the graph.
type MemberMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_RoomID       *string
	_UserID       *uint32
	add_UserID    *int32
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

// SetRoomID sets the "RoomID" field.
func (m *MemberMutation) SetRoomID(s string) {
	m._RoomID = &s
}

// RoomID returns the value of the "RoomID" field in the mutation.
func (m *MemberMutation) RoomID() (r string, exists bool) {
	v := m._RoomID
	if v == nil {
		return
	}
	return *v, true
}

// OldRoomID returns the old "RoomID" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldRoomID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRoomID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRoomID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoomID: %w", err)
	}
	return oldValue.RoomID, nil
}

// ResetRoomID resets all changes to the "RoomID" field.
func (m *MemberMutation) ResetRoomID() {
	m._RoomID = nil
}

// SetUserID sets the "UserID" field.
func (m *MemberMutation) SetUserID(u uint32) {
	m._UserID = &u
	m.add_UserID = nil
}

// UserID returns the value of the "UserID" field in the mutation.
func (m *MemberMutation) UserID() (r uint32, exists bool) {
	v := m._UserID
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "UserID" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldUserID(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds u to the "UserID" field.
func (m *MemberMutation) AddUserID(u int32) {
	if m.add_UserID != nil {
		*m.add_UserID += u
	} else {
		m.add_UserID = &u
	}
}

// AddedUserID returns the value that was added to the "UserID" field in this mutation.
func (m *MemberMutation) AddedUserID() (r int32, exists bool) {
	v := m.add_UserID
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "UserID" field.
func (m *MemberMutation) ResetUserID() {
	m._UserID = nil
	m.add_UserID = nil
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
	fields := make([]string, 0, 2)
	if m._RoomID != nil {
		fields = append(fields, member.FieldRoomID)
	}
	if m._UserID != nil {
		fields = append(fields, member.FieldUserID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case member.FieldRoomID:
		return m.RoomID()
	case member.FieldUserID:
		return m.UserID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case member.FieldRoomID:
		return m.OldRoomID(ctx)
	case member.FieldUserID:
		return m.OldUserID(ctx)
	}
	return nil, fmt.Errorf("unknown Member field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case member.FieldRoomID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoomID(v)
		return nil
	case member.FieldUserID:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MemberMutation) AddedFields() []string {
	var fields []string
	if m.add_UserID != nil {
		fields = append(fields, member.FieldUserID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MemberMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case member.FieldUserID:
		return m.AddedUserID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	case member.FieldUserID:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Member numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MemberMutation) ClearedFields() []string {
	return nil
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
	return fmt.Errorf("unknown Member nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MemberMutation) ResetField(name string) error {
	switch name {
	case member.FieldRoomID:
		m.ResetRoomID()
		return nil
	case member.FieldUserID:
		m.ResetUserID()
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

// MessageMutation represents an operation that mutates the Message nodes in the graph.
type MessageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_RoomID       *string
	_UserID       *uint32
	add_UserID    *int32
	_Body         *string
	_Time         *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Message, error)
	predicates    []predicate.Message
}

var _ ent.Mutation = (*MessageMutation)(nil)

// messageOption allows management of the mutation configuration using functional options.
type messageOption func(*MessageMutation)

// newMessageMutation creates new mutation for the Message entity.
func newMessageMutation(c config, op Op, opts ...messageOption) *MessageMutation {
	m := &MessageMutation{
		config:        c,
		op:            op,
		typ:           TypeMessage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessageID sets the ID field of the mutation.
func withMessageID(id int) messageOption {
	return func(m *MessageMutation) {
		var (
			err   error
			once  sync.Once
			value *Message
		)
		m.oldValue = func(ctx context.Context) (*Message, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Message.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessage sets the old Message of the mutation.
func withMessage(node *Message) messageOption {
	return func(m *MessageMutation) {
		m.oldValue = func(context.Context) (*Message, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Message.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRoomID sets the "RoomID" field.
func (m *MessageMutation) SetRoomID(s string) {
	m._RoomID = &s
}

// RoomID returns the value of the "RoomID" field in the mutation.
func (m *MessageMutation) RoomID() (r string, exists bool) {
	v := m._RoomID
	if v == nil {
		return
	}
	return *v, true
}

// OldRoomID returns the old "RoomID" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldRoomID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRoomID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRoomID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoomID: %w", err)
	}
	return oldValue.RoomID, nil
}

// ResetRoomID resets all changes to the "RoomID" field.
func (m *MessageMutation) ResetRoomID() {
	m._RoomID = nil
}

// SetUserID sets the "UserID" field.
func (m *MessageMutation) SetUserID(u uint32) {
	m._UserID = &u
	m.add_UserID = nil
}

// UserID returns the value of the "UserID" field in the mutation.
func (m *MessageMutation) UserID() (r uint32, exists bool) {
	v := m._UserID
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "UserID" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldUserID(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds u to the "UserID" field.
func (m *MessageMutation) AddUserID(u int32) {
	if m.add_UserID != nil {
		*m.add_UserID += u
	} else {
		m.add_UserID = &u
	}
}

// AddedUserID returns the value that was added to the "UserID" field in this mutation.
func (m *MessageMutation) AddedUserID() (r int32, exists bool) {
	v := m.add_UserID
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "UserID" field.
func (m *MessageMutation) ResetUserID() {
	m._UserID = nil
	m.add_UserID = nil
}

// SetBody sets the "Body" field.
func (m *MessageMutation) SetBody(s string) {
	m._Body = &s
}

// Body returns the value of the "Body" field in the mutation.
func (m *MessageMutation) Body() (r string, exists bool) {
	v := m._Body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "Body" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "Body" field.
func (m *MessageMutation) ResetBody() {
	m._Body = nil
}

// SetTime sets the "Time" field.
func (m *MessageMutation) SetTime(t time.Time) {
	m._Time = &t
}

// Time returns the value of the "Time" field in the mutation.
func (m *MessageMutation) Time() (r time.Time, exists bool) {
	v := m._Time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "Time" field's value of the Message entity.
// If the Message object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessageMutation) OldTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// ResetTime resets all changes to the "Time" field.
func (m *MessageMutation) ResetTime() {
	m._Time = nil
}

// Where appends a list predicates to the MessageMutation builder.
func (m *MessageMutation) Where(ps ...predicate.Message) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MessageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MessageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Message, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MessageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MessageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Message).
func (m *MessageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessageMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m._RoomID != nil {
		fields = append(fields, message.FieldRoomID)
	}
	if m._UserID != nil {
		fields = append(fields, message.FieldUserID)
	}
	if m._Body != nil {
		fields = append(fields, message.FieldBody)
	}
	if m._Time != nil {
		fields = append(fields, message.FieldTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case message.FieldRoomID:
		return m.RoomID()
	case message.FieldUserID:
		return m.UserID()
	case message.FieldBody:
		return m.Body()
	case message.FieldTime:
		return m.Time()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case message.FieldRoomID:
		return m.OldRoomID(ctx)
	case message.FieldUserID:
		return m.OldUserID(ctx)
	case message.FieldBody:
		return m.OldBody(ctx)
	case message.FieldTime:
		return m.OldTime(ctx)
	}
	return nil, fmt.Errorf("unknown Message field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case message.FieldRoomID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoomID(v)
		return nil
	case message.FieldUserID:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case message.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case message.FieldTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessageMutation) AddedFields() []string {
	var fields []string
	if m.add_UserID != nil {
		fields = append(fields, message.FieldUserID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessageMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case message.FieldUserID:
		return m.AddedUserID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessageMutation) AddField(name string, value ent.Value) error {
	switch name {
	case message.FieldUserID:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	}
	return fmt.Errorf("unknown Message numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessageMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessageMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Message nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessageMutation) ResetField(name string) error {
	switch name {
	case message.FieldRoomID:
		m.ResetRoomID()
		return nil
	case message.FieldUserID:
		m.ResetUserID()
		return nil
	case message.FieldBody:
		m.ResetBody()
		return nil
	case message.FieldTime:
		m.ResetTime()
		return nil
	}
	return fmt.Errorf("unknown Message field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessageMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessageMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessageMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessageMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Message unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessageMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Message edge %s", name)
}

// RoomMutation represents an operation that mutates the Room nodes in the graph.
type RoomMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_RoomID       *string
	_Users        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Room, error)
	predicates    []predicate.Room
}

var _ ent.Mutation = (*RoomMutation)(nil)

// roomOption allows management of the mutation configuration using functional options.
type roomOption func(*RoomMutation)

// newRoomMutation creates new mutation for the Room entity.
func newRoomMutation(c config, op Op, opts ...roomOption) *RoomMutation {
	m := &RoomMutation{
		config:        c,
		op:            op,
		typ:           TypeRoom,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRoomID sets the ID field of the mutation.
func withRoomID(id int) roomOption {
	return func(m *RoomMutation) {
		var (
			err   error
			once  sync.Once
			value *Room
		)
		m.oldValue = func(ctx context.Context) (*Room, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Room.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRoom sets the old Room of the mutation.
func withRoom(node *Room) roomOption {
	return func(m *RoomMutation) {
		m.oldValue = func(context.Context) (*Room, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RoomMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RoomMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RoomMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RoomMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Room.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetRoomID sets the "RoomID" field.
func (m *RoomMutation) SetRoomID(s string) {
	m._RoomID = &s
}

// RoomID returns the value of the "RoomID" field in the mutation.
func (m *RoomMutation) RoomID() (r string, exists bool) {
	v := m._RoomID
	if v == nil {
		return
	}
	return *v, true
}

// OldRoomID returns the old "RoomID" field's value of the Room entity.
// If the Room object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoomMutation) OldRoomID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRoomID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRoomID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRoomID: %w", err)
	}
	return oldValue.RoomID, nil
}

// ResetRoomID resets all changes to the "RoomID" field.
func (m *RoomMutation) ResetRoomID() {
	m._RoomID = nil
}

// SetUsers sets the "Users" field.
func (m *RoomMutation) SetUsers(s string) {
	m._Users = &s
}

// Users returns the value of the "Users" field in the mutation.
func (m *RoomMutation) Users() (r string, exists bool) {
	v := m._Users
	if v == nil {
		return
	}
	return *v, true
}

// OldUsers returns the old "Users" field's value of the Room entity.
// If the Room object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoomMutation) OldUsers(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsers is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsers requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsers: %w", err)
	}
	return oldValue.Users, nil
}

// ResetUsers resets all changes to the "Users" field.
func (m *RoomMutation) ResetUsers() {
	m._Users = nil
}

// Where appends a list predicates to the RoomMutation builder.
func (m *RoomMutation) Where(ps ...predicate.Room) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RoomMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RoomMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Room, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RoomMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RoomMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Room).
func (m *RoomMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RoomMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m._RoomID != nil {
		fields = append(fields, room.FieldRoomID)
	}
	if m._Users != nil {
		fields = append(fields, room.FieldUsers)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RoomMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case room.FieldRoomID:
		return m.RoomID()
	case room.FieldUsers:
		return m.Users()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RoomMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case room.FieldRoomID:
		return m.OldRoomID(ctx)
	case room.FieldUsers:
		return m.OldUsers(ctx)
	}
	return nil, fmt.Errorf("unknown Room field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoomMutation) SetField(name string, value ent.Value) error {
	switch name {
	case room.FieldRoomID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRoomID(v)
		return nil
	case room.FieldUsers:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsers(v)
		return nil
	}
	return fmt.Errorf("unknown Room field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RoomMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RoomMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoomMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Room numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RoomMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RoomMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RoomMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Room nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RoomMutation) ResetField(name string) error {
	switch name {
	case room.FieldRoomID:
		m.ResetRoomID()
		return nil
	case room.FieldUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown Room field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RoomMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RoomMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RoomMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RoomMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RoomMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RoomMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RoomMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Room unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RoomMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Room edge %s", name)
}
