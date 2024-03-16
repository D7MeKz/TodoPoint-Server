// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/pointinfo"
	"todopoint/common/db/ent/subtask"
	"todopoint/common/db/ent/task"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PointCreate is the builder for creating a Point entity.
type PointCreate struct {
	config
	mutation *PointMutation
	hooks    []Hook
}

// SetPoint sets the "point" field.
func (pc *PointCreate) SetPoint(i int) *PointCreate {
	pc.mutation.SetPoint(i)
	return pc
}

// SetType sets the "type" field.
func (pc *PointCreate) SetType(s string) *PointCreate {
	pc.mutation.SetType(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PointCreate) SetCreatedAt(t time.Time) *PointCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PointCreate) SetNillableCreatedAt(t *time.Time) *PointCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PointCreate) SetID(i int) *PointCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetSubtaskID sets the "subtask" edge to the SubTask entity by ID.
func (pc *PointCreate) SetSubtaskID(id int) *PointCreate {
	pc.mutation.SetSubtaskID(id)
	return pc
}

// SetNillableSubtaskID sets the "subtask" edge to the SubTask entity by ID if the given value is not nil.
func (pc *PointCreate) SetNillableSubtaskID(id *int) *PointCreate {
	if id != nil {
		pc = pc.SetSubtaskID(*id)
	}
	return pc
}

// SetSubtask sets the "subtask" edge to the SubTask entity.
func (pc *PointCreate) SetSubtask(s *SubTask) *PointCreate {
	return pc.SetSubtaskID(s.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (pc *PointCreate) SetTaskID(id int) *PointCreate {
	pc.mutation.SetTaskID(id)
	return pc
}

// SetNillableTaskID sets the "task" edge to the Task entity by ID if the given value is not nil.
func (pc *PointCreate) SetNillableTaskID(id *int) *PointCreate {
	if id != nil {
		pc = pc.SetTaskID(*id)
	}
	return pc
}

// SetTask sets the "task" edge to the Task entity.
func (pc *PointCreate) SetTask(t *Task) *PointCreate {
	return pc.SetTaskID(t.ID)
}

// SetPointInfoID sets the "point_info" edge to the PointInfo entity by ID.
func (pc *PointCreate) SetPointInfoID(id int) *PointCreate {
	pc.mutation.SetPointInfoID(id)
	return pc
}

// SetNillablePointInfoID sets the "point_info" edge to the PointInfo entity by ID if the given value is not nil.
func (pc *PointCreate) SetNillablePointInfoID(id *int) *PointCreate {
	if id != nil {
		pc = pc.SetPointInfoID(*id)
	}
	return pc
}

// SetPointInfo sets the "point_info" edge to the PointInfo entity.
func (pc *PointCreate) SetPointInfo(p *PointInfo) *PointCreate {
	return pc.SetPointInfoID(p.ID)
}

// Mutation returns the PointMutation object of the builder.
func (pc *PointCreate) Mutation() *PointMutation {
	return pc.mutation
}

// Save creates the Point in the database.
func (pc *PointCreate) Save(ctx context.Context) (*Point, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PointCreate) SaveX(ctx context.Context) *Point {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PointCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PointCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PointCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := point.DefaultCreatedAt
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PointCreate) check() error {
	if _, ok := pc.mutation.Point(); !ok {
		return &ValidationError{Name: "point", err: errors.New(`ent: missing required field "Point.point"`)}
	}
	if _, ok := pc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Point.type"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Point.created_at"`)}
	}
	return nil
}

func (pc *PointCreate) sqlSave(ctx context.Context) (*Point, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PointCreate) createSpec() (*Point, *sqlgraph.CreateSpec) {
	var (
		_node = &Point{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(point.Table, sqlgraph.NewFieldSpec(point.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Point(); ok {
		_spec.SetField(point.FieldPoint, field.TypeInt, value)
		_node.Point = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(point.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(point.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := pc.mutation.SubtaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   point.SubtaskTable,
			Columns: []string{point.SubtaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sub_task_point = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   point.TaskTable,
			Columns: []string{point.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_success_point = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PointInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   point.PointInfoTable,
			Columns: []string{point.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.point_info_points = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PointCreateBulk is the builder for creating many Point entities in bulk.
type PointCreateBulk struct {
	config
	err      error
	builders []*PointCreate
}

// Save creates the Point entities in the database.
func (pcb *PointCreateBulk) Save(ctx context.Context) ([]*Point, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Point, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PointMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PointCreateBulk) SaveX(ctx context.Context) []*Point {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PointCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PointCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}