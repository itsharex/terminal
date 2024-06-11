// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"terminal/ent/folders"
	"terminal/ent/hosts"
	"terminal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FoldersUpdate is the builder for updating Folders entities.
type FoldersUpdate struct {
	config
	hooks    []Hook
	mutation *FoldersMutation
}

// Where appends a list predicates to the FoldersUpdate builder.
func (fu *FoldersUpdate) Where(ps ...predicate.Folders) *FoldersUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetLabel sets the "label" field.
func (fu *FoldersUpdate) SetLabel(s string) *FoldersUpdate {
	fu.mutation.SetLabel(s)
	return fu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (fu *FoldersUpdate) SetNillableLabel(s *string) *FoldersUpdate {
	if s != nil {
		fu.SetLabel(*s)
	}
	return fu
}

// SetParentID sets the "parent_id" field.
func (fu *FoldersUpdate) SetParentID(i int) *FoldersUpdate {
	fu.mutation.SetParentID(i)
	return fu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (fu *FoldersUpdate) SetNillableParentID(i *int) *FoldersUpdate {
	if i != nil {
		fu.SetParentID(*i)
	}
	return fu
}

// ClearParentID clears the value of the "parent_id" field.
func (fu *FoldersUpdate) ClearParentID() *FoldersUpdate {
	fu.mutation.ClearParentID()
	return fu
}

// SetParent sets the "parent" edge to the Folders entity.
func (fu *FoldersUpdate) SetParent(f *Folders) *FoldersUpdate {
	return fu.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the Folders entity by IDs.
func (fu *FoldersUpdate) AddChildIDs(ids ...int) *FoldersUpdate {
	fu.mutation.AddChildIDs(ids...)
	return fu
}

// AddChildren adds the "children" edges to the Folders entity.
func (fu *FoldersUpdate) AddChildren(f ...*Folders) *FoldersUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddChildIDs(ids...)
}

// SetHostID sets the "host" edge to the Hosts entity by ID.
func (fu *FoldersUpdate) SetHostID(id int) *FoldersUpdate {
	fu.mutation.SetHostID(id)
	return fu
}

// SetNillableHostID sets the "host" edge to the Hosts entity by ID if the given value is not nil.
func (fu *FoldersUpdate) SetNillableHostID(id *int) *FoldersUpdate {
	if id != nil {
		fu = fu.SetHostID(*id)
	}
	return fu
}

// SetHost sets the "host" edge to the Hosts entity.
func (fu *FoldersUpdate) SetHost(h *Hosts) *FoldersUpdate {
	return fu.SetHostID(h.ID)
}

// Mutation returns the FoldersMutation object of the builder.
func (fu *FoldersUpdate) Mutation() *FoldersMutation {
	return fu.mutation
}

// ClearParent clears the "parent" edge to the Folders entity.
func (fu *FoldersUpdate) ClearParent() *FoldersUpdate {
	fu.mutation.ClearParent()
	return fu
}

// ClearChildren clears all "children" edges to the Folders entity.
func (fu *FoldersUpdate) ClearChildren() *FoldersUpdate {
	fu.mutation.ClearChildren()
	return fu
}

// RemoveChildIDs removes the "children" edge to Folders entities by IDs.
func (fu *FoldersUpdate) RemoveChildIDs(ids ...int) *FoldersUpdate {
	fu.mutation.RemoveChildIDs(ids...)
	return fu
}

// RemoveChildren removes "children" edges to Folders entities.
func (fu *FoldersUpdate) RemoveChildren(f ...*Folders) *FoldersUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveChildIDs(ids...)
}

// ClearHost clears the "host" edge to the Hosts entity.
func (fu *FoldersUpdate) ClearHost() *FoldersUpdate {
	fu.mutation.ClearHost()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FoldersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FoldersUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FoldersUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FoldersUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FoldersUpdate) check() error {
	if v, ok := fu.mutation.Label(); ok {
		if err := folders.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "Folders.label": %w`, err)}
		}
	}
	return nil
}

func (fu *FoldersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(folders.Table, folders.Columns, sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Label(); ok {
		_spec.SetField(folders.FieldLabel, field.TypeString, value)
	}
	if fu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folders.ParentTable,
			Columns: []string{folders.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folders.ParentTable,
			Columns: []string{folders.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   folders.HostTable,
			Columns: []string{folders.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   folders.HostTable,
			Columns: []string{folders.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folders.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FoldersUpdateOne is the builder for updating a single Folders entity.
type FoldersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FoldersMutation
}

// SetLabel sets the "label" field.
func (fuo *FoldersUpdateOne) SetLabel(s string) *FoldersUpdateOne {
	fuo.mutation.SetLabel(s)
	return fuo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (fuo *FoldersUpdateOne) SetNillableLabel(s *string) *FoldersUpdateOne {
	if s != nil {
		fuo.SetLabel(*s)
	}
	return fuo
}

// SetParentID sets the "parent_id" field.
func (fuo *FoldersUpdateOne) SetParentID(i int) *FoldersUpdateOne {
	fuo.mutation.SetParentID(i)
	return fuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (fuo *FoldersUpdateOne) SetNillableParentID(i *int) *FoldersUpdateOne {
	if i != nil {
		fuo.SetParentID(*i)
	}
	return fuo
}

// ClearParentID clears the value of the "parent_id" field.
func (fuo *FoldersUpdateOne) ClearParentID() *FoldersUpdateOne {
	fuo.mutation.ClearParentID()
	return fuo
}

// SetParent sets the "parent" edge to the Folders entity.
func (fuo *FoldersUpdateOne) SetParent(f *Folders) *FoldersUpdateOne {
	return fuo.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the Folders entity by IDs.
func (fuo *FoldersUpdateOne) AddChildIDs(ids ...int) *FoldersUpdateOne {
	fuo.mutation.AddChildIDs(ids...)
	return fuo
}

// AddChildren adds the "children" edges to the Folders entity.
func (fuo *FoldersUpdateOne) AddChildren(f ...*Folders) *FoldersUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddChildIDs(ids...)
}

// SetHostID sets the "host" edge to the Hosts entity by ID.
func (fuo *FoldersUpdateOne) SetHostID(id int) *FoldersUpdateOne {
	fuo.mutation.SetHostID(id)
	return fuo
}

// SetNillableHostID sets the "host" edge to the Hosts entity by ID if the given value is not nil.
func (fuo *FoldersUpdateOne) SetNillableHostID(id *int) *FoldersUpdateOne {
	if id != nil {
		fuo = fuo.SetHostID(*id)
	}
	return fuo
}

// SetHost sets the "host" edge to the Hosts entity.
func (fuo *FoldersUpdateOne) SetHost(h *Hosts) *FoldersUpdateOne {
	return fuo.SetHostID(h.ID)
}

// Mutation returns the FoldersMutation object of the builder.
func (fuo *FoldersUpdateOne) Mutation() *FoldersMutation {
	return fuo.mutation
}

// ClearParent clears the "parent" edge to the Folders entity.
func (fuo *FoldersUpdateOne) ClearParent() *FoldersUpdateOne {
	fuo.mutation.ClearParent()
	return fuo
}

// ClearChildren clears all "children" edges to the Folders entity.
func (fuo *FoldersUpdateOne) ClearChildren() *FoldersUpdateOne {
	fuo.mutation.ClearChildren()
	return fuo
}

// RemoveChildIDs removes the "children" edge to Folders entities by IDs.
func (fuo *FoldersUpdateOne) RemoveChildIDs(ids ...int) *FoldersUpdateOne {
	fuo.mutation.RemoveChildIDs(ids...)
	return fuo
}

// RemoveChildren removes "children" edges to Folders entities.
func (fuo *FoldersUpdateOne) RemoveChildren(f ...*Folders) *FoldersUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveChildIDs(ids...)
}

// ClearHost clears the "host" edge to the Hosts entity.
func (fuo *FoldersUpdateOne) ClearHost() *FoldersUpdateOne {
	fuo.mutation.ClearHost()
	return fuo
}

// Where appends a list predicates to the FoldersUpdate builder.
func (fuo *FoldersUpdateOne) Where(ps ...predicate.Folders) *FoldersUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FoldersUpdateOne) Select(field string, fields ...string) *FoldersUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Folders entity.
func (fuo *FoldersUpdateOne) Save(ctx context.Context) (*Folders, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FoldersUpdateOne) SaveX(ctx context.Context) *Folders {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FoldersUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FoldersUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FoldersUpdateOne) check() error {
	if v, ok := fuo.mutation.Label(); ok {
		if err := folders.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "Folders.label": %w`, err)}
		}
	}
	return nil
}

func (fuo *FoldersUpdateOne) sqlSave(ctx context.Context) (_node *Folders, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(folders.Table, folders.Columns, sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Folders.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, folders.FieldID)
		for _, f := range fields {
			if !folders.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != folders.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Label(); ok {
		_spec.SetField(folders.FieldLabel, field.TypeString, value)
	}
	if fuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folders.ParentTable,
			Columns: []string{folders.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folders.ParentTable,
			Columns: []string{folders.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folders.ChildrenTable,
			Columns: []string{folders.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   folders.HostTable,
			Columns: []string{folders.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   folders.HostTable,
			Columns: []string{folders.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Folders{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folders.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}