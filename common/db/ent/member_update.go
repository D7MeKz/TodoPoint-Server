// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todopoint/common/db/ent/member"
	"todopoint/common/db/ent/pointinfo"
	"todopoint/common/db/ent/predicate"
	"todopoint/common/db/ent/task"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmail(s *string) *MemberUpdate {
	if s != nil {
		mu.SetEmail(*s)
	}
	return mu
}

// SetUsername sets the "username" field.
func (mu *MemberUpdate) SetUsername(s string) *MemberUpdate {
	mu.mutation.SetUsername(s)
	return mu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableUsername(s *string) *MemberUpdate {
	if s != nil {
		mu.SetUsername(*s)
	}
	return mu
}

// ClearUsername clears the value of the "username" field.
func (mu *MemberUpdate) ClearUsername() *MemberUpdate {
	mu.mutation.ClearUsername()
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mu *MemberUpdate) SetNillablePassword(s *string) *MemberUpdate {
	if s != nil {
		mu.SetPassword(*s)
	}
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MemberUpdate) SetCreatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableCreatedAt(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// AddPointInfoIDs adds the "point_info" edge to the PointInfo entity by IDs.
func (mu *MemberUpdate) AddPointInfoIDs(ids ...int) *MemberUpdate {
	mu.mutation.AddPointInfoIDs(ids...)
	return mu
}

// AddPointInfo adds the "point_info" edges to the PointInfo entity.
func (mu *MemberUpdate) AddPointInfo(p ...*PointInfo) *MemberUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddPointInfoIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (mu *MemberUpdate) AddTaskIDs(ids ...int) *MemberUpdate {
	mu.mutation.AddTaskIDs(ids...)
	return mu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (mu *MemberUpdate) AddTasks(t ...*Task) *MemberUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTaskIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearPointInfo clears all "point_info" edges to the PointInfo entity.
func (mu *MemberUpdate) ClearPointInfo() *MemberUpdate {
	mu.mutation.ClearPointInfo()
	return mu
}

// RemovePointInfoIDs removes the "point_info" edge to PointInfo entities by IDs.
func (mu *MemberUpdate) RemovePointInfoIDs(ids ...int) *MemberUpdate {
	mu.mutation.RemovePointInfoIDs(ids...)
	return mu
}

// RemovePointInfo removes "point_info" edges to PointInfo entities.
func (mu *MemberUpdate) RemovePointInfo(p ...*PointInfo) *MemberUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemovePointInfoIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (mu *MemberUpdate) ClearTasks() *MemberUpdate {
	mu.mutation.ClearTasks()
	return mu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (mu *MemberUpdate) RemoveTaskIDs(ids ...int) *MemberUpdate {
	mu.mutation.RemoveTaskIDs(ids...)
	return mu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (mu *MemberUpdate) RemoveTasks(t ...*Task) *MemberUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
	}
	if mu.mutation.UsernameCleared() {
		_spec.ClearField(member.FieldUsername, field.TypeString)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
	}
	if mu.mutation.PointInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedPointInfoIDs(); len(nodes) > 0 && !mu.mutation.PointInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PointInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !mu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmail(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetEmail(*s)
	}
	return muo
}

// SetUsername sets the "username" field.
func (muo *MemberUpdateOne) SetUsername(s string) *MemberUpdateOne {
	muo.mutation.SetUsername(s)
	return muo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableUsername(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetUsername(*s)
	}
	return muo
}

// ClearUsername clears the value of the "username" field.
func (muo *MemberUpdateOne) ClearUsername() *MemberUpdateOne {
	muo.mutation.ClearUsername()
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillablePassword(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetPassword(*s)
	}
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MemberUpdateOne) SetCreatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableCreatedAt(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// AddPointInfoIDs adds the "point_info" edge to the PointInfo entity by IDs.
func (muo *MemberUpdateOne) AddPointInfoIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.AddPointInfoIDs(ids...)
	return muo
}

// AddPointInfo adds the "point_info" edges to the PointInfo entity.
func (muo *MemberUpdateOne) AddPointInfo(p ...*PointInfo) *MemberUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddPointInfoIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (muo *MemberUpdateOne) AddTaskIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.AddTaskIDs(ids...)
	return muo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (muo *MemberUpdateOne) AddTasks(t ...*Task) *MemberUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTaskIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearPointInfo clears all "point_info" edges to the PointInfo entity.
func (muo *MemberUpdateOne) ClearPointInfo() *MemberUpdateOne {
	muo.mutation.ClearPointInfo()
	return muo
}

// RemovePointInfoIDs removes the "point_info" edge to PointInfo entities by IDs.
func (muo *MemberUpdateOne) RemovePointInfoIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.RemovePointInfoIDs(ids...)
	return muo
}

// RemovePointInfo removes "point_info" edges to PointInfo entities.
func (muo *MemberUpdateOne) RemovePointInfo(p ...*PointInfo) *MemberUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemovePointInfoIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (muo *MemberUpdateOne) ClearTasks() *MemberUpdateOne {
	muo.mutation.ClearTasks()
	return muo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (muo *MemberUpdateOne) RemoveTaskIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.RemoveTaskIDs(ids...)
	return muo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (muo *MemberUpdateOne) RemoveTasks(t ...*Task) *MemberUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Member.password": %w`, err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Username(); ok {
		_spec.SetField(member.FieldUsername, field.TypeString, value)
	}
	if muo.mutation.UsernameCleared() {
		_spec.ClearField(member.FieldUsername, field.TypeString)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
	}
	if muo.mutation.PointInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedPointInfoIDs(); len(nodes) > 0 && !muo.mutation.PointInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PointInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.PointInfoTable,
			Columns: []string{member.PointInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !muo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.TasksTable,
			Columns: []string{member.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
