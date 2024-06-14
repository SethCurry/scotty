// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/scotty/internal/ent/predicate"
	"github.com/SethCurry/scotty/internal/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetDiscordID sets the "discord_id" field.
func (uu *UserUpdate) SetDiscordID(s string) *UserUpdate {
	uu.mutation.SetDiscordID(s)
	return uu
}

// SetNillableDiscordID sets the "discord_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDiscordID(s *string) *UserUpdate {
	if s != nil {
		uu.SetDiscordID(*s)
	}
	return uu
}

// SetRank sets the "rank" field.
func (uu *UserUpdate) SetRank(i int8) *UserUpdate {
	uu.mutation.ResetRank()
	uu.mutation.SetRank(i)
	return uu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRank(i *int8) *UserUpdate {
	if i != nil {
		uu.SetRank(*i)
	}
	return uu
}

// AddRank adds i to the "rank" field.
func (uu *UserUpdate) AddRank(i int8) *UserUpdate {
	uu.mutation.AddRank(i)
	return uu
}

// SetFinalsID sets the "finals_id" field.
func (uu *UserUpdate) SetFinalsID(s string) *UserUpdate {
	uu.mutation.SetFinalsID(s)
	return uu
}

// SetNillableFinalsID sets the "finals_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFinalsID(s *string) *UserUpdate {
	if s != nil {
		uu.SetFinalsID(*s)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.DiscordID(); ok {
		_spec.SetField(user.FieldDiscordID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.FinalsID(); ok {
		_spec.SetField(user.FieldFinalsID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetDiscordID sets the "discord_id" field.
func (uuo *UserUpdateOne) SetDiscordID(s string) *UserUpdateOne {
	uuo.mutation.SetDiscordID(s)
	return uuo
}

// SetNillableDiscordID sets the "discord_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDiscordID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDiscordID(*s)
	}
	return uuo
}

// SetRank sets the "rank" field.
func (uuo *UserUpdateOne) SetRank(i int8) *UserUpdateOne {
	uuo.mutation.ResetRank()
	uuo.mutation.SetRank(i)
	return uuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRank(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetRank(*i)
	}
	return uuo
}

// AddRank adds i to the "rank" field.
func (uuo *UserUpdateOne) AddRank(i int8) *UserUpdateOne {
	uuo.mutation.AddRank(i)
	return uuo
}

// SetFinalsID sets the "finals_id" field.
func (uuo *UserUpdateOne) SetFinalsID(s string) *UserUpdateOne {
	uuo.mutation.SetFinalsID(s)
	return uuo
}

// SetNillableFinalsID sets the "finals_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFinalsID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFinalsID(*s)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.DiscordID(); ok {
		_spec.SetField(user.FieldDiscordID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Rank(); ok {
		_spec.SetField(user.FieldRank, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedRank(); ok {
		_spec.AddField(user.FieldRank, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.FinalsID(); ok {
		_spec.SetField(user.FieldFinalsID, field.TypeString, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
