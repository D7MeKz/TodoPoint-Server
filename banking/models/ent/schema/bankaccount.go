package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// BankAccount holds the schema definition for the BankAccount entity.
type BankAccount struct {
	ent.Schema
}

// Fields of the BankAccount.
func (BankAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("bank_name"),
		field.UUID("bank_account", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the BankAccount.
func (BankAccount) Edges() []ent.Edge {
	return nil
}
