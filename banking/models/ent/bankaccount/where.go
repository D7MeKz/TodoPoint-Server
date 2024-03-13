// Code generated by ent, DO NOT EDIT.

package bankaccount

import (
	"time"
	"todopoint/banking/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLTE(FieldID, id))
}

// BankName applies equality check predicate on the "bank_name" field. It's identical to BankNameEQ.
func BankName(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldBankName, v))
}

// BankAccount applies equality check predicate on the "bank_account" field. It's identical to BankAccountEQ.
func BankAccount(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldBankAccount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// BankNameEQ applies the EQ predicate on the "bank_name" field.
func BankNameEQ(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldBankName, v))
}

// BankNameNEQ applies the NEQ predicate on the "bank_name" field.
func BankNameNEQ(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNEQ(FieldBankName, v))
}

// BankNameIn applies the In predicate on the "bank_name" field.
func BankNameIn(vs ...string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldIn(FieldBankName, vs...))
}

// BankNameNotIn applies the NotIn predicate on the "bank_name" field.
func BankNameNotIn(vs ...string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNotIn(FieldBankName, vs...))
}

// BankNameGT applies the GT predicate on the "bank_name" field.
func BankNameGT(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGT(FieldBankName, v))
}

// BankNameGTE applies the GTE predicate on the "bank_name" field.
func BankNameGTE(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGTE(FieldBankName, v))
}

// BankNameLT applies the LT predicate on the "bank_name" field.
func BankNameLT(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLT(FieldBankName, v))
}

// BankNameLTE applies the LTE predicate on the "bank_name" field.
func BankNameLTE(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLTE(FieldBankName, v))
}

// BankNameContains applies the Contains predicate on the "bank_name" field.
func BankNameContains(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldContains(FieldBankName, v))
}

// BankNameHasPrefix applies the HasPrefix predicate on the "bank_name" field.
func BankNameHasPrefix(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldHasPrefix(FieldBankName, v))
}

// BankNameHasSuffix applies the HasSuffix predicate on the "bank_name" field.
func BankNameHasSuffix(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldHasSuffix(FieldBankName, v))
}

// BankNameEqualFold applies the EqualFold predicate on the "bank_name" field.
func BankNameEqualFold(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEqualFold(FieldBankName, v))
}

// BankNameContainsFold applies the ContainsFold predicate on the "bank_name" field.
func BankNameContainsFold(v string) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldContainsFold(FieldBankName, v))
}

// BankAccountEQ applies the EQ predicate on the "bank_account" field.
func BankAccountEQ(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldBankAccount, v))
}

// BankAccountNEQ applies the NEQ predicate on the "bank_account" field.
func BankAccountNEQ(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNEQ(FieldBankAccount, v))
}

// BankAccountIn applies the In predicate on the "bank_account" field.
func BankAccountIn(vs ...uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldIn(FieldBankAccount, vs...))
}

// BankAccountNotIn applies the NotIn predicate on the "bank_account" field.
func BankAccountNotIn(vs ...uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNotIn(FieldBankAccount, vs...))
}

// BankAccountGT applies the GT predicate on the "bank_account" field.
func BankAccountGT(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGT(FieldBankAccount, v))
}

// BankAccountGTE applies the GTE predicate on the "bank_account" field.
func BankAccountGTE(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGTE(FieldBankAccount, v))
}

// BankAccountLT applies the LT predicate on the "bank_account" field.
func BankAccountLT(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLT(FieldBankAccount, v))
}

// BankAccountLTE applies the LTE predicate on the "bank_account" field.
func BankAccountLTE(v uuid.UUID) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLTE(FieldBankAccount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BankAccount {
	return predicate.BankAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BankAccount) predicate.BankAccount {
	return predicate.BankAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BankAccount) predicate.BankAccount {
	return predicate.BankAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BankAccount) predicate.BankAccount {
	return predicate.BankAccount(sql.NotPredicates(p))
}
