// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/subtask"
	"todopoint/common/db/ent/task"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubTaskCreate is the builder for creating a SubTask entity.
type SubTaskCreate struct {
	config
	mutation *SubTaskMutation
	hooks    []Hook
}

// SetPointID sets the "point" edge to the Point entity by ID.
func (stc *SubTaskCreate) SetPointID(id int) *SubTaskCreate {
	stc.mutation.SetPointID(id)
	return stc
}

// SetNillablePointID sets the "point" edge to the Point entity by ID if the given value is not nil.
func (stc *SubTaskCreate) SetNillablePointID(id *int) *SubTaskCreate {
	if id != nil {
		stc = stc.SetPointID(*id)
	}
	return stc
}

// SetPoint sets the "point" edge to the Point entity.
func (stc *SubTaskCreate) SetPoint(p *Point) *SubTaskCreate {
	return stc.SetPointID(p.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (stc *SubTaskCreate) SetTaskID(id int) *SubTaskCreate {
	stc.mutation.SetTaskID(id)
	return stc
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (stc *SubTaskCreate) SetNillableTaskID(id *int) *SubTaskCreate {
	if id != nil {
		stc = stc.SetTaskID(*id)
	}
	return stc
}

// SetTask sets the "task" edge to the Task entity.
func (stc *SubTaskCreate) SetTask(t *Task) *SubTaskCreate {
	return stc.SetTaskID(t.ID)
}

// Mutation returns the SubTaskMutation object of the builder.
func (stc *SubTaskCreate) Mutation() *SubTaskMutation {
	return stc.mutation
}

// Save creates the SubTask in the database.
func (stc *SubTaskCreate) Save(ctx context.Context) (*SubTask, error) {
	return withHooks(ctx, stc.sqlSave, stc.mutation, stc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (stc *SubTaskCreate) SaveX(ctx context.Context) *SubTask {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *SubTaskCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *SubTaskCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SubTaskCreate) check() error {
	return nil
}

func (stc *SubTaskCreate) sqlSave(ctx context.Context) (*SubTask, error) {
	if err := stc.check(); err != nil {
		return nil, err
	}
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	stc.mutation.id = &_node.ID
	stc.mutation.done = true
	return _node, nil
}

func (stc *SubTaskCreate) createSpec() (*SubTask, *sqlgraph.CreateSpec) {
	var (
		_node = &SubTask{config: stc.config}
		_spec = sqlgraph.NewCreateSpec(subtask.Table, sqlgraph.NewFieldSpec(subtask.FieldID, field.TypeInt))
	)
	if nodes := stc.mutation.PointIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   subtask.PointTable,
			Columns: []string{subtask.PointColumn},
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
	if nodes := stc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subtask.TaskTable,
			Columns: []string{subtask.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_subtask = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubTaskCreateBulk is the builder for creating many SubTask entities in bulk.
type SubTaskCreateBulk struct {
	config
	err      error
	builders []*SubTaskCreate
}

// Save creates the SubTask entities in the database.
func (stcb *SubTaskCreateBulk) Save(ctx context.Context) ([]*SubTask, error) {
	if stcb.err != nil {
		return nil, stcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SubTask, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubTaskMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *SubTaskCreateBulk) SaveX(ctx context.Context) []*SubTask {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *SubTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *SubTaskCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}
