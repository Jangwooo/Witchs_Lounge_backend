// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/witchs-lounge_backend/ent/predicate"
	"github.com/witchs-lounge_backend/ent/userachievement"
)

// UserAchievementDelete is the builder for deleting a UserAchievement entity.
type UserAchievementDelete struct {
	config
	hooks    []Hook
	mutation *UserAchievementMutation
}

// Where appends a list predicates to the UserAchievementDelete builder.
func (uad *UserAchievementDelete) Where(ps ...predicate.UserAchievement) *UserAchievementDelete {
	uad.mutation.Where(ps...)
	return uad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uad *UserAchievementDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uad.sqlExec, uad.mutation, uad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uad *UserAchievementDelete) ExecX(ctx context.Context) int {
	n, err := uad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uad *UserAchievementDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userachievement.Table, sqlgraph.NewFieldSpec(userachievement.FieldID, field.TypeUUID))
	if ps := uad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uad.mutation.done = true
	return affected, err
}

// UserAchievementDeleteOne is the builder for deleting a single UserAchievement entity.
type UserAchievementDeleteOne struct {
	uad *UserAchievementDelete
}

// Where appends a list predicates to the UserAchievementDelete builder.
func (uado *UserAchievementDeleteOne) Where(ps ...predicate.UserAchievement) *UserAchievementDeleteOne {
	uado.uad.mutation.Where(ps...)
	return uado
}

// Exec executes the deletion query.
func (uado *UserAchievementDeleteOne) Exec(ctx context.Context) error {
	n, err := uado.uad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userachievement.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uado *UserAchievementDeleteOne) ExecX(ctx context.Context) {
	if err := uado.Exec(ctx); err != nil {
		panic(err)
	}
}
