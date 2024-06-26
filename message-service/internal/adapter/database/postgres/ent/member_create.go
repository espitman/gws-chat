// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/member"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRoomID sets the "RoomID" field.
func (mc *MemberCreate) SetRoomID(s string) *MemberCreate {
	mc.mutation.SetRoomID(s)
	return mc
}

// SetUserID sets the "UserID" field.
func (mc *MemberCreate) SetUserID(u uint32) *MemberCreate {
	mc.mutation.SetUserID(u)
	return mc
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	if _, ok := mc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "RoomID", err: errors.New(`ent: missing required field "Member.RoomID"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "UserID", err: errors.New(`ent: missing required field "Member.UserID"`)}
	}
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	)
	_spec.OnConflict = mc.conflict
	if value, ok := mc.mutation.RoomID(); ok {
		_spec.SetField(member.FieldRoomID, field.TypeString, value)
		_node.RoomID = value
	}
	if value, ok := mc.mutation.UserID(); ok {
		_spec.SetField(member.FieldUserID, field.TypeUint32, value)
		_node.UserID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.Create().
//		SetRoomID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetRoomID(v+v).
//		}).
//		Exec(ctx)
func (mc *MemberCreate) OnConflict(opts ...sql.ConflictOption) *MemberUpsertOne {
	mc.conflict = opts
	return &MemberUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MemberCreate) OnConflictColumns(columns ...string) *MemberUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertOne{
		create: mc,
	}
}

type (
	// MemberUpsertOne is the builder for "upsert"-ing
	//  one Member node.
	MemberUpsertOne struct {
		create *MemberCreate
	}

	// MemberUpsert is the "OnConflict" setter.
	MemberUpsert struct {
		*sql.UpdateSet
	}
)

// SetRoomID sets the "RoomID" field.
func (u *MemberUpsert) SetRoomID(v string) *MemberUpsert {
	u.Set(member.FieldRoomID, v)
	return u
}

// UpdateRoomID sets the "RoomID" field to the value that was provided on create.
func (u *MemberUpsert) UpdateRoomID() *MemberUpsert {
	u.SetExcluded(member.FieldRoomID)
	return u
}

// SetUserID sets the "UserID" field.
func (u *MemberUpsert) SetUserID(v uint32) *MemberUpsert {
	u.Set(member.FieldUserID, v)
	return u
}

// UpdateUserID sets the "UserID" field to the value that was provided on create.
func (u *MemberUpsert) UpdateUserID() *MemberUpsert {
	u.SetExcluded(member.FieldUserID)
	return u
}

// AddUserID adds v to the "UserID" field.
func (u *MemberUpsert) AddUserID(v uint32) *MemberUpsert {
	u.Add(member.FieldUserID, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MemberUpsertOne) UpdateNewValues() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MemberUpsertOne) Ignore() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertOne) DoNothing() *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreate.OnConflict
// documentation for more info.
func (u *MemberUpsertOne) Update(set func(*MemberUpsert)) *MemberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetRoomID sets the "RoomID" field.
func (u *MemberUpsertOne) SetRoomID(v string) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "RoomID" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateRoomID() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateRoomID()
	})
}

// SetUserID sets the "UserID" field.
func (u *MemberUpsertOne) SetUserID(v uint32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "UserID" field.
func (u *MemberUpsertOne) AddUserID(v uint32) *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "UserID" field to the value that was provided on create.
func (u *MemberUpsertOne) UpdateUserID() *MemberUpsertOne {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *MemberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MemberUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MemberUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
	conflict []sql.ConflictOption
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Member.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MemberUpsert) {
//			SetRoomID(v+v).
//		}).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflict(opts ...sql.ConflictOption) *MemberUpsertBulk {
	mcb.conflict = opts
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MemberCreateBulk) OnConflictColumns(columns ...string) *MemberUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MemberUpsertBulk{
		create: mcb,
	}
}

// MemberUpsertBulk is the builder for "upsert"-ing
// a bulk of Member nodes.
type MemberUpsertBulk struct {
	create *MemberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MemberUpsertBulk) UpdateNewValues() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Member.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MemberUpsertBulk) Ignore() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MemberUpsertBulk) DoNothing() *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MemberCreateBulk.OnConflict
// documentation for more info.
func (u *MemberUpsertBulk) Update(set func(*MemberUpsert)) *MemberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MemberUpsert{UpdateSet: update})
	}))
	return u
}

// SetRoomID sets the "RoomID" field.
func (u *MemberUpsertBulk) SetRoomID(v string) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "RoomID" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateRoomID() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateRoomID()
	})
}

// SetUserID sets the "UserID" field.
func (u *MemberUpsertBulk) SetUserID(v uint32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "UserID" field.
func (u *MemberUpsertBulk) AddUserID(v uint32) *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "UserID" field to the value that was provided on create.
func (u *MemberUpsertBulk) UpdateUserID() *MemberUpsertBulk {
	return u.Update(func(s *MemberUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *MemberUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MemberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MemberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MemberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
