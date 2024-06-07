// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"terminal/ent/hosts"
	"terminal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostsDelete is the builder for deleting a Hosts entity.
type HostsDelete struct {
	config
	hooks    []Hook
	mutation *HostsMutation
}

// Where appends a list predicates to the HostsDelete builder.
func (hd *HostsDelete) Where(ps ...predicate.Hosts) *HostsDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HostsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hd.sqlExec, hd.mutation, hd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HostsDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HostsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hosts.Table, sqlgraph.NewFieldSpec(hosts.FieldID, field.TypeInt))
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hd.mutation.done = true
	return affected, err
}

// HostsDeleteOne is the builder for deleting a single Hosts entity.
type HostsDeleteOne struct {
	hd *HostsDelete
}

// Where appends a list predicates to the HostsDelete builder.
func (hdo *HostsDeleteOne) Where(ps ...predicate.Hosts) *HostsDeleteOne {
	hdo.hd.mutation.Where(ps...)
	return hdo
}

// Exec executes the deletion query.
func (hdo *HostsDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hosts.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HostsDeleteOne) ExecX(ctx context.Context) {
	if err := hdo.Exec(ctx); err != nil {
		panic(err)
	}
}
