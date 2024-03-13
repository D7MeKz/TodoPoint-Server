// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todopoint/banking/models/ent/bankaccount"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BankAccountCreate is the builder for creating a BankAccount entity.
type BankAccountCreate struct {
	config
	mutation *BankAccountMutation
	hooks    []Hook
}

// SetBankName sets the "bank_name" field.
func (bac *BankAccountCreate) SetBankName(s string) *BankAccountCreate {
	bac.mutation.SetBankName(s)
	return bac
}

// SetBankAccount sets the "bank_account" field.
func (bac *BankAccountCreate) SetBankAccount(u uuid.UUID) *BankAccountCreate {
	bac.mutation.SetBankAccount(u)
	return bac
}

// SetNillableBankAccount sets the "bank_account" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableBankAccount(u *uuid.UUID) *BankAccountCreate {
	if u != nil {
		bac.SetBankAccount(*u)
	}
	return bac
}

// SetCreatedAt sets the "created_at" field.
func (bac *BankAccountCreate) SetCreatedAt(t time.Time) *BankAccountCreate {
	bac.mutation.SetCreatedAt(t)
	return bac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bac *BankAccountCreate) SetNillableCreatedAt(t *time.Time) *BankAccountCreate {
	if t != nil {
		bac.SetCreatedAt(*t)
	}
	return bac
}

// SetID sets the "id" field.
func (bac *BankAccountCreate) SetID(i int) *BankAccountCreate {
	bac.mutation.SetID(i)
	return bac
}

// Mutation returns the BankAccountMutation object of the builder.
func (bac *BankAccountCreate) Mutation() *BankAccountMutation {
	return bac.mutation
}

// Save creates the BankAccount in the database.
func (bac *BankAccountCreate) Save(ctx context.Context) (*BankAccount, error) {
	bac.defaults()
	return withHooks(ctx, bac.sqlSave, bac.mutation, bac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bac *BankAccountCreate) SaveX(ctx context.Context) *BankAccount {
	v, err := bac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bac *BankAccountCreate) Exec(ctx context.Context) error {
	_, err := bac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bac *BankAccountCreate) ExecX(ctx context.Context) {
	if err := bac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bac *BankAccountCreate) defaults() {
	if _, ok := bac.mutation.BankAccount(); !ok {
		v := bankaccount.DefaultBankAccount()
		bac.mutation.SetBankAccount(v)
	}
	if _, ok := bac.mutation.CreatedAt(); !ok {
		v := bankaccount.DefaultCreatedAt
		bac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bac *BankAccountCreate) check() error {
	if _, ok := bac.mutation.BankName(); !ok {
		return &ValidationError{Name: "bank_name", err: errors.New(`ent: missing required field "BankAccount.bank_name"`)}
	}
	if _, ok := bac.mutation.BankAccount(); !ok {
		return &ValidationError{Name: "bank_account", err: errors.New(`ent: missing required field "BankAccount.bank_account"`)}
	}
	if _, ok := bac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BankAccount.created_at"`)}
	}
	return nil
}

func (bac *BankAccountCreate) sqlSave(ctx context.Context) (*BankAccount, error) {
	if err := bac.check(); err != nil {
		return nil, err
	}
	_node, _spec := bac.createSpec()
	if err := sqlgraph.CreateNode(ctx, bac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bac.mutation.id = &_node.ID
	bac.mutation.done = true
	return _node, nil
}

func (bac *BankAccountCreate) createSpec() (*BankAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &BankAccount{config: bac.config}
		_spec = sqlgraph.NewCreateSpec(bankaccount.Table, sqlgraph.NewFieldSpec(bankaccount.FieldID, field.TypeInt))
	)
	if id, ok := bac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bac.mutation.BankName(); ok {
		_spec.SetField(bankaccount.FieldBankName, field.TypeString, value)
		_node.BankName = value
	}
	if value, ok := bac.mutation.BankAccount(); ok {
		_spec.SetField(bankaccount.FieldBankAccount, field.TypeUUID, value)
		_node.BankAccount = value
	}
	if value, ok := bac.mutation.CreatedAt(); ok {
		_spec.SetField(bankaccount.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// BankAccountCreateBulk is the builder for creating many BankAccount entities in bulk.
type BankAccountCreateBulk struct {
	config
	err      error
	builders []*BankAccountCreate
}

// Save creates the BankAccount entities in the database.
func (bacb *BankAccountCreateBulk) Save(ctx context.Context) ([]*BankAccount, error) {
	if bacb.err != nil {
		return nil, bacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bacb.builders))
	nodes := make([]*BankAccount, len(bacb.builders))
	mutators := make([]Mutator, len(bacb.builders))
	for i := range bacb.builders {
		func(i int, root context.Context) {
			builder := bacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BankAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, bacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, bacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) SaveX(ctx context.Context) []*BankAccount {
	v, err := bacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bacb *BankAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := bacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bacb *BankAccountCreateBulk) ExecX(ctx context.Context) {
	if err := bacb.Exec(ctx); err != nil {
		panic(err)
	}
}
