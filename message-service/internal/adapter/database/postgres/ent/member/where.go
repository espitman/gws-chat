// Code generated by ent, DO NOT EDIT.

package member

import (
	"entgo.io/ent/dialect/sql"
	"github.com/espitman/gws-chat/message-service/internal/adapter/database/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldID, id))
}

// RoomID applies equality check predicate on the "RoomID" field. It's identical to RoomIDEQ.
func RoomID(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldRoomID, v))
}

// UserID applies equality check predicate on the "UserID" field. It's identical to UserIDEQ.
func UserID(v uint32) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldUserID, v))
}

// RoomIDEQ applies the EQ predicate on the "RoomID" field.
func RoomIDEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldRoomID, v))
}

// RoomIDNEQ applies the NEQ predicate on the "RoomID" field.
func RoomIDNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldRoomID, v))
}

// RoomIDIn applies the In predicate on the "RoomID" field.
func RoomIDIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldRoomID, vs...))
}

// RoomIDNotIn applies the NotIn predicate on the "RoomID" field.
func RoomIDNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldRoomID, vs...))
}

// RoomIDGT applies the GT predicate on the "RoomID" field.
func RoomIDGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldRoomID, v))
}

// RoomIDGTE applies the GTE predicate on the "RoomID" field.
func RoomIDGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldRoomID, v))
}

// RoomIDLT applies the LT predicate on the "RoomID" field.
func RoomIDLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldRoomID, v))
}

// RoomIDLTE applies the LTE predicate on the "RoomID" field.
func RoomIDLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldRoomID, v))
}

// RoomIDContains applies the Contains predicate on the "RoomID" field.
func RoomIDContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldRoomID, v))
}

// RoomIDHasPrefix applies the HasPrefix predicate on the "RoomID" field.
func RoomIDHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldRoomID, v))
}

// RoomIDHasSuffix applies the HasSuffix predicate on the "RoomID" field.
func RoomIDHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldRoomID, v))
}

// RoomIDEqualFold applies the EqualFold predicate on the "RoomID" field.
func RoomIDEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldRoomID, v))
}

// RoomIDContainsFold applies the ContainsFold predicate on the "RoomID" field.
func RoomIDContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldRoomID, v))
}

// UserIDEQ applies the EQ predicate on the "UserID" field.
func UserIDEQ(v uint32) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "UserID" field.
func UserIDNEQ(v uint32) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "UserID" field.
func UserIDIn(vs ...uint32) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "UserID" field.
func UserIDNotIn(vs ...uint32) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "UserID" field.
func UserIDGT(v uint32) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "UserID" field.
func UserIDGTE(v uint32) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "UserID" field.
func UserIDLT(v uint32) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "UserID" field.
func UserIDLTE(v uint32) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldUserID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(sql.NotPredicates(p))
}
