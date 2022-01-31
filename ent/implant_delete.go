// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/predicate"
)

// ImplantDelete is the builder for deleting a Implant entity.
type ImplantDelete struct {
	config
	hooks    []Hook
	mutation *ImplantMutation
}

// Where appends a list predicates to the ImplantDelete builder.
func (id *ImplantDelete) Where(ps ...predicate.Implant) *ImplantDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *ImplantDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImplantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			if id.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *ImplantDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *ImplantDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: implant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implant.FieldID,
			},
		},
	}
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, id.driver, _spec)
}

// ImplantDeleteOne is the builder for deleting a single Implant entity.
type ImplantDeleteOne struct {
	id *ImplantDelete
}

// Exec executes the deletion query.
func (ido *ImplantDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{implant.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *ImplantDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
