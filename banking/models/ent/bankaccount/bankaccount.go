// Code generated by ent, DO NOT EDIT.

package bankaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the bankaccount type in the database.
	Label = "bank_account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "bank_name"
	// FieldBankAccount holds the string denoting the bank_account field in the database.
	FieldBankAccount = "bank_account"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the bankaccount in the database.
	Table = "bank_accounts"
)

// Columns holds all SQL columns for bankaccount fields.
var Columns = []string{
	FieldID,
	FieldBankName,
	FieldBankAccount,
	FieldCreatedAt,
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
	// DefaultBankAccount holds the default value on creation for the "bank_account" field.
	DefaultBankAccount func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the BankAccount queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBankName orders the results by the bank_name field.
func ByBankName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankName, opts...).ToFunc()
}

// ByBankAccount orders the results by the bank_account field.
func ByBankAccount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBankAccount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
