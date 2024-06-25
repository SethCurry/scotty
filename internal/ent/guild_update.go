// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SethCurry/scotty/internal/ent/autorolerule"
	"github.com/SethCurry/scotty/internal/ent/guild"
	"github.com/SethCurry/scotty/internal/ent/predicate"
)

// GuildUpdate is the builder for updating Guild entities.
type GuildUpdate struct {
	config
	hooks    []Hook
	mutation *GuildMutation
}

// Where appends a list predicates to the GuildUpdate builder.
func (gu *GuildUpdate) Where(ps ...predicate.Guild) *GuildUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GuildUpdate) SetName(s string) *GuildUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableName(s *string) *GuildUpdate {
	if s != nil {
		gu.SetName(*s)
	}
	return gu
}

// SetGuildID sets the "guild_id" field.
func (gu *GuildUpdate) SetGuildID(s string) *GuildUpdate {
	gu.mutation.SetGuildID(s)
	return gu
}

// SetNillableGuildID sets the "guild_id" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableGuildID(s *string) *GuildUpdate {
	if s != nil {
		gu.SetGuildID(*s)
	}
	return gu
}

// SetWelcomeTemplate sets the "welcome_template" field.
func (gu *GuildUpdate) SetWelcomeTemplate(s string) *GuildUpdate {
	gu.mutation.SetWelcomeTemplate(s)
	return gu
}

// SetNillableWelcomeTemplate sets the "welcome_template" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableWelcomeTemplate(s *string) *GuildUpdate {
	if s != nil {
		gu.SetWelcomeTemplate(*s)
	}
	return gu
}

// SetWelcomeChannel sets the "welcome_channel" field.
func (gu *GuildUpdate) SetWelcomeChannel(s string) *GuildUpdate {
	gu.mutation.SetWelcomeChannel(s)
	return gu
}

// SetNillableWelcomeChannel sets the "welcome_channel" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableWelcomeChannel(s *string) *GuildUpdate {
	if s != nil {
		gu.SetWelcomeChannel(*s)
	}
	return gu
}

// AddAutoRoleRuleIDs adds the "auto_role_rules" edge to the AutoRoleRule entity by IDs.
func (gu *GuildUpdate) AddAutoRoleRuleIDs(ids ...int) *GuildUpdate {
	gu.mutation.AddAutoRoleRuleIDs(ids...)
	return gu
}

// AddAutoRoleRules adds the "auto_role_rules" edges to the AutoRoleRule entity.
func (gu *GuildUpdate) AddAutoRoleRules(a ...*AutoRoleRule) *GuildUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return gu.AddAutoRoleRuleIDs(ids...)
}

// Mutation returns the GuildMutation object of the builder.
func (gu *GuildUpdate) Mutation() *GuildMutation {
	return gu.mutation
}

// ClearAutoRoleRules clears all "auto_role_rules" edges to the AutoRoleRule entity.
func (gu *GuildUpdate) ClearAutoRoleRules() *GuildUpdate {
	gu.mutation.ClearAutoRoleRules()
	return gu
}

// RemoveAutoRoleRuleIDs removes the "auto_role_rules" edge to AutoRoleRule entities by IDs.
func (gu *GuildUpdate) RemoveAutoRoleRuleIDs(ids ...int) *GuildUpdate {
	gu.mutation.RemoveAutoRoleRuleIDs(ids...)
	return gu
}

// RemoveAutoRoleRules removes "auto_role_rules" edges to AutoRoleRule entities.
func (gu *GuildUpdate) RemoveAutoRoleRules(a ...*AutoRoleRule) *GuildUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return gu.RemoveAutoRoleRuleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GuildUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GuildUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GuildUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GuildUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(guild.Table, guild.Columns, sqlgraph.NewFieldSpec(guild.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(guild.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.GuildID(); ok {
		_spec.SetField(guild.FieldGuildID, field.TypeString, value)
	}
	if value, ok := gu.mutation.WelcomeTemplate(); ok {
		_spec.SetField(guild.FieldWelcomeTemplate, field.TypeString, value)
	}
	if value, ok := gu.mutation.WelcomeChannel(); ok {
		_spec.SetField(guild.FieldWelcomeChannel, field.TypeString, value)
	}
	if gu.mutation.AutoRoleRulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedAutoRoleRulesIDs(); len(nodes) > 0 && !gu.mutation.AutoRoleRulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.AutoRoleRulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GuildUpdateOne is the builder for updating a single Guild entity.
type GuildUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GuildMutation
}

// SetName sets the "name" field.
func (guo *GuildUpdateOne) SetName(s string) *GuildUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableName(s *string) *GuildUpdateOne {
	if s != nil {
		guo.SetName(*s)
	}
	return guo
}

// SetGuildID sets the "guild_id" field.
func (guo *GuildUpdateOne) SetGuildID(s string) *GuildUpdateOne {
	guo.mutation.SetGuildID(s)
	return guo
}

// SetNillableGuildID sets the "guild_id" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableGuildID(s *string) *GuildUpdateOne {
	if s != nil {
		guo.SetGuildID(*s)
	}
	return guo
}

// SetWelcomeTemplate sets the "welcome_template" field.
func (guo *GuildUpdateOne) SetWelcomeTemplate(s string) *GuildUpdateOne {
	guo.mutation.SetWelcomeTemplate(s)
	return guo
}

// SetNillableWelcomeTemplate sets the "welcome_template" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableWelcomeTemplate(s *string) *GuildUpdateOne {
	if s != nil {
		guo.SetWelcomeTemplate(*s)
	}
	return guo
}

// SetWelcomeChannel sets the "welcome_channel" field.
func (guo *GuildUpdateOne) SetWelcomeChannel(s string) *GuildUpdateOne {
	guo.mutation.SetWelcomeChannel(s)
	return guo
}

// SetNillableWelcomeChannel sets the "welcome_channel" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableWelcomeChannel(s *string) *GuildUpdateOne {
	if s != nil {
		guo.SetWelcomeChannel(*s)
	}
	return guo
}

// AddAutoRoleRuleIDs adds the "auto_role_rules" edge to the AutoRoleRule entity by IDs.
func (guo *GuildUpdateOne) AddAutoRoleRuleIDs(ids ...int) *GuildUpdateOne {
	guo.mutation.AddAutoRoleRuleIDs(ids...)
	return guo
}

// AddAutoRoleRules adds the "auto_role_rules" edges to the AutoRoleRule entity.
func (guo *GuildUpdateOne) AddAutoRoleRules(a ...*AutoRoleRule) *GuildUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return guo.AddAutoRoleRuleIDs(ids...)
}

// Mutation returns the GuildMutation object of the builder.
func (guo *GuildUpdateOne) Mutation() *GuildMutation {
	return guo.mutation
}

// ClearAutoRoleRules clears all "auto_role_rules" edges to the AutoRoleRule entity.
func (guo *GuildUpdateOne) ClearAutoRoleRules() *GuildUpdateOne {
	guo.mutation.ClearAutoRoleRules()
	return guo
}

// RemoveAutoRoleRuleIDs removes the "auto_role_rules" edge to AutoRoleRule entities by IDs.
func (guo *GuildUpdateOne) RemoveAutoRoleRuleIDs(ids ...int) *GuildUpdateOne {
	guo.mutation.RemoveAutoRoleRuleIDs(ids...)
	return guo
}

// RemoveAutoRoleRules removes "auto_role_rules" edges to AutoRoleRule entities.
func (guo *GuildUpdateOne) RemoveAutoRoleRules(a ...*AutoRoleRule) *GuildUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return guo.RemoveAutoRoleRuleIDs(ids...)
}

// Where appends a list predicates to the GuildUpdate builder.
func (guo *GuildUpdateOne) Where(ps ...predicate.Guild) *GuildUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GuildUpdateOne) Select(field string, fields ...string) *GuildUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Guild entity.
func (guo *GuildUpdateOne) Save(ctx context.Context) (*Guild, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GuildUpdateOne) SaveX(ctx context.Context) *Guild {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GuildUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GuildUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GuildUpdateOne) sqlSave(ctx context.Context) (_node *Guild, err error) {
	_spec := sqlgraph.NewUpdateSpec(guild.Table, guild.Columns, sqlgraph.NewFieldSpec(guild.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Guild.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guild.FieldID)
		for _, f := range fields {
			if !guild.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != guild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(guild.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.GuildID(); ok {
		_spec.SetField(guild.FieldGuildID, field.TypeString, value)
	}
	if value, ok := guo.mutation.WelcomeTemplate(); ok {
		_spec.SetField(guild.FieldWelcomeTemplate, field.TypeString, value)
	}
	if value, ok := guo.mutation.WelcomeChannel(); ok {
		_spec.SetField(guild.FieldWelcomeChannel, field.TypeString, value)
	}
	if guo.mutation.AutoRoleRulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedAutoRoleRulesIDs(); len(nodes) > 0 && !guo.mutation.AutoRoleRulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.AutoRoleRulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   guild.AutoRoleRulesTable,
			Columns: []string{guild.AutoRoleRulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(autorolerule.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Guild{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
