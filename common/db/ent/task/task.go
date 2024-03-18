// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldTotalStatus holds the string denoting the total_status field in the database.
	FieldTotalStatus = "total_status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// EdgeSubtask holds the string denoting the subtask edge name in mutations.
	EdgeSubtask = "subtask"
	// EdgeSuccessPoint holds the string denoting the success_point edge name in mutations.
	EdgeSuccessPoint = "success_point"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// SubtaskTable is the table that holds the subtask relation/edge.
	SubtaskTable = "tasks"
	// SubtaskInverseTable is the table name for the SubTask entity.
	// It exists in this package in order to avoid circular dependency with the "subtask" package.
	SubtaskInverseTable = "sub_tasks"
	// SubtaskColumn is the table column denoting the subtask relation/edge.
	SubtaskColumn = "task_subtask"
	// SuccessPointTable is the table that holds the success_point relation/edge.
	SuccessPointTable = "points"
	// SuccessPointInverseTable is the table name for the Point entity.
	// It exists in this package in order to avoid circular dependency with the "point" package.
	SuccessPointInverseTable = "points"
	// SuccessPointColumn is the table column denoting the success_point relation/edge.
	SuccessPointColumn = "task_success_point"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "member_tasks"
	// UserInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	UserInverseTable = "members"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldTotalStatus,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_subtask",
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"member_id", "task_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt time.Time
)

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByTotalStatus orders the results by the total_status field.
func ByTotalStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// BySubtaskField orders the results by subtask field.
func BySubtaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubtaskStep(), sql.OrderByField(field, opts...))
	}
}

// BySuccessPointField orders the results by success_point field.
func BySuccessPointField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSuccessPointStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubtaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubtaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SubtaskTable, SubtaskColumn),
	)
}
func newSuccessPointStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SuccessPointInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SuccessPointTable, SuccessPointColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
