// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"terminal/ent/folders"
	"terminal/ent/hosts"
	"terminal/ent/keys"
	"terminal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostsUpdate is the builder for updating Hosts entities.
type HostsUpdate struct {
	config
	hooks    []Hook
	mutation *HostsMutation
}

// Where appends a list predicates to the HostsUpdate builder.
func (hu *HostsUpdate) Where(ps ...predicate.Hosts) *HostsUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetLabel sets the "label" field.
func (hu *HostsUpdate) SetLabel(s string) *HostsUpdate {
	hu.mutation.SetLabel(s)
	return hu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (hu *HostsUpdate) SetNillableLabel(s *string) *HostsUpdate {
	if s != nil {
		hu.SetLabel(*s)
	}
	return hu
}

// SetUsername sets the "username" field.
func (hu *HostsUpdate) SetUsername(s string) *HostsUpdate {
	hu.mutation.SetUsername(s)
	return hu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (hu *HostsUpdate) SetNillableUsername(s *string) *HostsUpdate {
	if s != nil {
		hu.SetUsername(*s)
	}
	return hu
}

// SetAddress sets the "address" field.
func (hu *HostsUpdate) SetAddress(s string) *HostsUpdate {
	hu.mutation.SetAddress(s)
	return hu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (hu *HostsUpdate) SetNillableAddress(s *string) *HostsUpdate {
	if s != nil {
		hu.SetAddress(*s)
	}
	return hu
}

// SetPort sets the "port" field.
func (hu *HostsUpdate) SetPort(u uint) *HostsUpdate {
	hu.mutation.ResetPort()
	hu.mutation.SetPort(u)
	return hu
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (hu *HostsUpdate) SetNillablePort(u *uint) *HostsUpdate {
	if u != nil {
		hu.SetPort(*u)
	}
	return hu
}

// AddPort adds u to the "port" field.
func (hu *HostsUpdate) AddPort(u int) *HostsUpdate {
	hu.mutation.AddPort(u)
	return hu
}

// SetPassword sets the "password" field.
func (hu *HostsUpdate) SetPassword(s string) *HostsUpdate {
	hu.mutation.SetPassword(s)
	return hu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (hu *HostsUpdate) SetNillablePassword(s *string) *HostsUpdate {
	if s != nil {
		hu.SetPassword(*s)
	}
	return hu
}

// ClearPassword clears the value of the "password" field.
func (hu *HostsUpdate) ClearPassword() *HostsUpdate {
	hu.mutation.ClearPassword()
	return hu
}

// SetFolderID sets the "folder" edge to the Folders entity by ID.
func (hu *HostsUpdate) SetFolderID(id int) *HostsUpdate {
	hu.mutation.SetFolderID(id)
	return hu
}

// SetFolder sets the "folder" edge to the Folders entity.
func (hu *HostsUpdate) SetFolder(f *Folders) *HostsUpdate {
	return hu.SetFolderID(f.ID)
}

// SetKeyID sets the "key" edge to the Keys entity by ID.
func (hu *HostsUpdate) SetKeyID(id int) *HostsUpdate {
	hu.mutation.SetKeyID(id)
	return hu
}

// SetNillableKeyID sets the "key" edge to the Keys entity by ID if the given value is not nil.
func (hu *HostsUpdate) SetNillableKeyID(id *int) *HostsUpdate {
	if id != nil {
		hu = hu.SetKeyID(*id)
	}
	return hu
}

// SetKey sets the "key" edge to the Keys entity.
func (hu *HostsUpdate) SetKey(k *Keys) *HostsUpdate {
	return hu.SetKeyID(k.ID)
}

// Mutation returns the HostsMutation object of the builder.
func (hu *HostsUpdate) Mutation() *HostsMutation {
	return hu.mutation
}

// ClearFolder clears the "folder" edge to the Folders entity.
func (hu *HostsUpdate) ClearFolder() *HostsUpdate {
	hu.mutation.ClearFolder()
	return hu
}

// ClearKey clears the "key" edge to the Keys entity.
func (hu *HostsUpdate) ClearKey() *HostsUpdate {
	hu.mutation.ClearKey()
	return hu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostsUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostsUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostsUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HostsUpdate) check() error {
	if v, ok := hu.mutation.Label(); ok {
		if err := hosts.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "Hosts.label": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Username(); ok {
		if err := hosts.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Hosts.username": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Address(); ok {
		if err := hosts.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Hosts.address": %w`, err)}
		}
	}
	if v, ok := hu.mutation.Port(); ok {
		if err := hosts.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Hosts.port": %w`, err)}
		}
	}
	if _, ok := hu.mutation.FolderID(); hu.mutation.FolderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Hosts.folder"`)
	}
	return nil
}

func (hu *HostsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hosts.Table, hosts.Columns, sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.Label(); ok {
		_spec.SetField(hosts.FieldLabel, field.TypeString, value)
	}
	if value, ok := hu.mutation.Username(); ok {
		_spec.SetField(hosts.FieldUsername, field.TypeString, value)
	}
	if value, ok := hu.mutation.Address(); ok {
		_spec.SetField(hosts.FieldAddress, field.TypeString, value)
	}
	if value, ok := hu.mutation.Port(); ok {
		_spec.SetField(hosts.FieldPort, field.TypeUint, value)
	}
	if value, ok := hu.mutation.AddedPort(); ok {
		_spec.AddField(hosts.FieldPort, field.TypeUint, value)
	}
	if value, ok := hu.mutation.Password(); ok {
		_spec.SetField(hosts.FieldPassword, field.TypeString, value)
	}
	if hu.mutation.PasswordCleared() {
		_spec.ClearField(hosts.FieldPassword, field.TypeString)
	}
	if hu.mutation.FolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hosts.FolderTable,
			Columns: []string{hosts.FolderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.FolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hosts.FolderTable,
			Columns: []string{hosts.FolderColumn},
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
	if hu.mutation.KeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hosts.KeyTable,
			Columns: []string{hosts.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(keys.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hosts.KeyTable,
			Columns: []string{hosts.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(keys.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hosts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HostsUpdateOne is the builder for updating a single Hosts entity.
type HostsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostsMutation
}

// SetLabel sets the "label" field.
func (huo *HostsUpdateOne) SetLabel(s string) *HostsUpdateOne {
	huo.mutation.SetLabel(s)
	return huo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (huo *HostsUpdateOne) SetNillableLabel(s *string) *HostsUpdateOne {
	if s != nil {
		huo.SetLabel(*s)
	}
	return huo
}

// SetUsername sets the "username" field.
func (huo *HostsUpdateOne) SetUsername(s string) *HostsUpdateOne {
	huo.mutation.SetUsername(s)
	return huo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (huo *HostsUpdateOne) SetNillableUsername(s *string) *HostsUpdateOne {
	if s != nil {
		huo.SetUsername(*s)
	}
	return huo
}

// SetAddress sets the "address" field.
func (huo *HostsUpdateOne) SetAddress(s string) *HostsUpdateOne {
	huo.mutation.SetAddress(s)
	return huo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (huo *HostsUpdateOne) SetNillableAddress(s *string) *HostsUpdateOne {
	if s != nil {
		huo.SetAddress(*s)
	}
	return huo
}

// SetPort sets the "port" field.
func (huo *HostsUpdateOne) SetPort(u uint) *HostsUpdateOne {
	huo.mutation.ResetPort()
	huo.mutation.SetPort(u)
	return huo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (huo *HostsUpdateOne) SetNillablePort(u *uint) *HostsUpdateOne {
	if u != nil {
		huo.SetPort(*u)
	}
	return huo
}

// AddPort adds u to the "port" field.
func (huo *HostsUpdateOne) AddPort(u int) *HostsUpdateOne {
	huo.mutation.AddPort(u)
	return huo
}

// SetPassword sets the "password" field.
func (huo *HostsUpdateOne) SetPassword(s string) *HostsUpdateOne {
	huo.mutation.SetPassword(s)
	return huo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (huo *HostsUpdateOne) SetNillablePassword(s *string) *HostsUpdateOne {
	if s != nil {
		huo.SetPassword(*s)
	}
	return huo
}

// ClearPassword clears the value of the "password" field.
func (huo *HostsUpdateOne) ClearPassword() *HostsUpdateOne {
	huo.mutation.ClearPassword()
	return huo
}

// SetFolderID sets the "folder" edge to the Folders entity by ID.
func (huo *HostsUpdateOne) SetFolderID(id int) *HostsUpdateOne {
	huo.mutation.SetFolderID(id)
	return huo
}

// SetFolder sets the "folder" edge to the Folders entity.
func (huo *HostsUpdateOne) SetFolder(f *Folders) *HostsUpdateOne {
	return huo.SetFolderID(f.ID)
}

// SetKeyID sets the "key" edge to the Keys entity by ID.
func (huo *HostsUpdateOne) SetKeyID(id int) *HostsUpdateOne {
	huo.mutation.SetKeyID(id)
	return huo
}

// SetNillableKeyID sets the "key" edge to the Keys entity by ID if the given value is not nil.
func (huo *HostsUpdateOne) SetNillableKeyID(id *int) *HostsUpdateOne {
	if id != nil {
		huo = huo.SetKeyID(*id)
	}
	return huo
}

// SetKey sets the "key" edge to the Keys entity.
func (huo *HostsUpdateOne) SetKey(k *Keys) *HostsUpdateOne {
	return huo.SetKeyID(k.ID)
}

// Mutation returns the HostsMutation object of the builder.
func (huo *HostsUpdateOne) Mutation() *HostsMutation {
	return huo.mutation
}

// ClearFolder clears the "folder" edge to the Folders entity.
func (huo *HostsUpdateOne) ClearFolder() *HostsUpdateOne {
	huo.mutation.ClearFolder()
	return huo
}

// ClearKey clears the "key" edge to the Keys entity.
func (huo *HostsUpdateOne) ClearKey() *HostsUpdateOne {
	huo.mutation.ClearKey()
	return huo
}

// Where appends a list predicates to the HostsUpdate builder.
func (huo *HostsUpdateOne) Where(ps ...predicate.Hosts) *HostsUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HostsUpdateOne) Select(field string, fields ...string) *HostsUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Hosts entity.
func (huo *HostsUpdateOne) Save(ctx context.Context) (*Hosts, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostsUpdateOne) SaveX(ctx context.Context) *Hosts {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostsUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostsUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HostsUpdateOne) check() error {
	if v, ok := huo.mutation.Label(); ok {
		if err := hosts.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "Hosts.label": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Username(); ok {
		if err := hosts.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Hosts.username": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Address(); ok {
		if err := hosts.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Hosts.address": %w`, err)}
		}
	}
	if v, ok := huo.mutation.Port(); ok {
		if err := hosts.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Hosts.port": %w`, err)}
		}
	}
	if _, ok := huo.mutation.FolderID(); huo.mutation.FolderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Hosts.folder"`)
	}
	return nil
}

func (huo *HostsUpdateOne) sqlSave(ctx context.Context) (_node *Hosts, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hosts.Table, hosts.Columns, sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Hosts.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hosts.FieldID)
		for _, f := range fields {
			if !hosts.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hosts.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.Label(); ok {
		_spec.SetField(hosts.FieldLabel, field.TypeString, value)
	}
	if value, ok := huo.mutation.Username(); ok {
		_spec.SetField(hosts.FieldUsername, field.TypeString, value)
	}
	if value, ok := huo.mutation.Address(); ok {
		_spec.SetField(hosts.FieldAddress, field.TypeString, value)
	}
	if value, ok := huo.mutation.Port(); ok {
		_spec.SetField(hosts.FieldPort, field.TypeUint, value)
	}
	if value, ok := huo.mutation.AddedPort(); ok {
		_spec.AddField(hosts.FieldPort, field.TypeUint, value)
	}
	if value, ok := huo.mutation.Password(); ok {
		_spec.SetField(hosts.FieldPassword, field.TypeString, value)
	}
	if huo.mutation.PasswordCleared() {
		_spec.ClearField(hosts.FieldPassword, field.TypeString)
	}
	if huo.mutation.FolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hosts.FolderTable,
			Columns: []string{hosts.FolderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(folders.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.FolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hosts.FolderTable,
			Columns: []string{hosts.FolderColumn},
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
	if huo.mutation.KeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hosts.KeyTable,
			Columns: []string{hosts.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(keys.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hosts.KeyTable,
			Columns: []string{hosts.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(keys.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Hosts{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hosts.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
