// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todopoint/common/db/ent/member"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/subtask"
	"todopoint/common/db/ent/task"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (tc *TaskCreate) SetTitle(s string) *TaskCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetTotalStatus sets the "total_status" field.
func (tc *TaskCreate) SetTotalStatus(i int) *TaskCreate {
	tc.mutation.SetTotalStatus(i)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetModifiedAt sets the "modified_at" field.
func (tc *TaskCreate) SetModifiedAt(t time.Time) *TaskCreate {
	tc.mutation.SetModifiedAt(t)
	return tc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableModifiedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetModifiedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(i int) *TaskCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetSubtaskID sets the "subtask" edge to the SubTask entity by ID.
func (tc *TaskCreate) SetSubtaskID(id int) *TaskCreate {
	tc.mutation.SetSubtaskID(id)
	return tc
}

// SetNillableSubtaskID sets the "subtask" edge to the SubTask entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableSubtaskID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetSubtaskID(*id)
	}
	return tc
}

// SetSubtask sets the "subtask" edge to the SubTask entity.
func (tc *TaskCreate) SetSubtask(s *SubTask) *TaskCreate {
	return tc.SetSubtaskID(s.ID)
}

// SetSuccessPointID sets the "success_point" edge to the Point entity by ID.
func (tc *TaskCreate) SetSuccessPointID(id int) *TaskCreate {
	tc.mutation.SetSuccessPointID(id)
	return tc
}

// SetNillableSuccessPointID sets the "success_point" edge to the Point entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableSuccessPointID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetSuccessPointID(*id)
	}
	return tc
}

// SetSuccessPoint sets the "success_point" edge to the Point entity.
func (tc *TaskCreate) SetSuccessPoint(p *Point) *TaskCreate {
	return tc.SetSuccessPointID(p.ID)
}

// AddUserIDIDs adds the "user_id" edge to the Member entity by IDs.
func (tc *TaskCreate) AddUserIDIDs(ids ...int) *TaskCreate {
	tc.mutation.AddUserIDIDs(ids...)
	return tc
}

// AddUserID adds the "user_id" edges to the Member entity.
func (tc *TaskCreate) AddUserID(m ...*Member) *TaskCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tc.AddUserIDIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.ModifiedAt(); !ok {
		v := task.DefaultModifiedAt
		tc.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Task.title"`)}
	}
	if _, ok := tc.mutation.TotalStatus(); !ok {
		return &ValidationError{Name: "total_status", err: errors.New(`ent: missing required field "Task.total_status"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "Task.modified_at"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tc.mutation.TotalStatus(); ok {
		_spec.SetField(task.FieldTotalStatus, field.TypeInt, value)
		_node.TotalStatus = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.ModifiedAt(); ok {
		_spec.SetField(task.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if nodes := tc.mutation.SubtaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.SubtaskTable,
			Columns: []string{task.SubtaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_subtask = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.SuccessPointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   task.SuccessPointTable,
			Columns: []string{task.SuccessPointColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(point.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.UserIDTable,
			Columns: task.UserIDPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}