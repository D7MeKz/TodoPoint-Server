// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"todopoint/banking/models/ent/bankaccount"
	"todopoint/banking/models/ent/schema"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bankaccountFields := schema.BankAccount{}.Fields()
	_ = bankaccountFields
	// bankaccountDescBankAccount is the schema descriptor for bank_account field.
	bankaccountDescBankAccount := bankaccountFields[3].Descriptor()
	// bankaccount.DefaultBankAccount holds the default value on creation for the bank_account field.
	bankaccount.DefaultBankAccount = bankaccountDescBankAccount.Default.(func() uuid.UUID)
	// bankaccountDescCreatedAt is the schema descriptor for created_at field.
	bankaccountDescCreatedAt := bankaccountFields[4].Descriptor()
	// bankaccount.DefaultCreatedAt holds the default value on creation for the created_at field.
	bankaccount.DefaultCreatedAt = bankaccountDescCreatedAt.Default.(time.Time)
}