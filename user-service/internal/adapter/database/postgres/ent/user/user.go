// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIsOnline holds the string denoting the isonline field in the database.
	FieldIsOnline = "is_online"
	// Table holds the table name of the user in the database.
	Table = "chat_users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPassword,
	FieldAvatar,
	FieldStatus,
	FieldIsOnline,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAvatar holds the default value on creation for the "Avatar" field.
	DefaultAvatar string
	// DefaultStatus holds the default value on creation for the "Status" field.
	DefaultStatus string
	// DefaultIsOnline holds the default value on creation for the "IsOnline" field.
	DefaultIsOnline bool
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the Name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassword orders the results by the Password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAvatar orders the results by the Avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByStatus orders the results by the Status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByIsOnline orders the results by the IsOnline field.
func ByIsOnline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOnline, opts...).ToFunc()
}
