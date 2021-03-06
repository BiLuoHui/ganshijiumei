// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/BiLuoHui/ganshijiumei/ent/predicate"
	"github.com/BiLuoHui/ganshijiumei/ent/wugong"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// WuGongUpdate is the builder for updating WuGong entities.
type WuGongUpdate struct {
	config
	hooks      []Hook
	mutation   *WuGongMutation
	predicates []predicate.WuGong
}

// Where adds a new predicate for the builder.
func (wgu *WuGongUpdate) Where(ps ...predicate.WuGong) *WuGongUpdate {
	wgu.predicates = append(wgu.predicates, ps...)
	return wgu
}

// SetUpdatedAt sets the updated_at field.
func (wgu *WuGongUpdate) SetUpdatedAt(t time.Time) *WuGongUpdate {
	wgu.mutation.SetUpdatedAt(t)
	return wgu
}

// SetName sets the name field.
func (wgu *WuGongUpdate) SetName(s string) *WuGongUpdate {
	wgu.mutation.SetName(s)
	return wgu
}

// SetDamage sets the damage field.
func (wgu *WuGongUpdate) SetDamage(i int) *WuGongUpdate {
	wgu.mutation.ResetDamage()
	wgu.mutation.SetDamage(i)
	return wgu
}

// SetNillableDamage sets the damage field if the given value is not nil.
func (wgu *WuGongUpdate) SetNillableDamage(i *int) *WuGongUpdate {
	if i != nil {
		wgu.SetDamage(*i)
	}
	return wgu
}

// AddDamage adds i to damage.
func (wgu *WuGongUpdate) AddDamage(i int) *WuGongUpdate {
	wgu.mutation.AddDamage(i)
	return wgu
}

// SetLevel sets the level field.
func (wgu *WuGongUpdate) SetLevel(u uint8) *WuGongUpdate {
	wgu.mutation.ResetLevel()
	wgu.mutation.SetLevel(u)
	return wgu
}

// AddLevel adds u to level.
func (wgu *WuGongUpdate) AddLevel(u uint8) *WuGongUpdate {
	wgu.mutation.AddLevel(u)
	return wgu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (wgu *WuGongUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := wgu.mutation.UpdatedAt(); !ok {
		v := wugong.UpdateDefaultUpdatedAt()
		wgu.mutation.SetUpdatedAt(v)
	}
	if v, ok := wgu.mutation.Level(); ok {
		if err := wugong.LevelValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"level\": %v", err)
		}
	}
	var (
		err      error
		affected int
	)
	if len(wgu.hooks) == 0 {
		affected, err = wgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WuGongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wgu.mutation = mutation
			affected, err = wgu.sqlSave(ctx)
			return affected, err
		})
		for i := len(wgu.hooks) - 1; i >= 0; i-- {
			mut = wgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wgu *WuGongUpdate) SaveX(ctx context.Context) int {
	affected, err := wgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wgu *WuGongUpdate) Exec(ctx context.Context) error {
	_, err := wgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wgu *WuGongUpdate) ExecX(ctx context.Context) {
	if err := wgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wgu *WuGongUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wugong.Table,
			Columns: wugong.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wugong.FieldID,
			},
		},
	}
	if ps := wgu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wugong.FieldUpdatedAt,
		})
	}
	if value, ok := wgu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wugong.FieldName,
		})
	}
	if value, ok := wgu.mutation.Damage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wugong.FieldDamage,
		})
	}
	if value, ok := wgu.mutation.AddedDamage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wugong.FieldDamage,
		})
	}
	if value, ok := wgu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: wugong.FieldLevel,
		})
	}
	if value, ok := wgu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: wugong.FieldLevel,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wugong.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WuGongUpdateOne is the builder for updating a single WuGong entity.
type WuGongUpdateOne struct {
	config
	hooks    []Hook
	mutation *WuGongMutation
}

// SetUpdatedAt sets the updated_at field.
func (wguo *WuGongUpdateOne) SetUpdatedAt(t time.Time) *WuGongUpdateOne {
	wguo.mutation.SetUpdatedAt(t)
	return wguo
}

// SetName sets the name field.
func (wguo *WuGongUpdateOne) SetName(s string) *WuGongUpdateOne {
	wguo.mutation.SetName(s)
	return wguo
}

// SetDamage sets the damage field.
func (wguo *WuGongUpdateOne) SetDamage(i int) *WuGongUpdateOne {
	wguo.mutation.ResetDamage()
	wguo.mutation.SetDamage(i)
	return wguo
}

// SetNillableDamage sets the damage field if the given value is not nil.
func (wguo *WuGongUpdateOne) SetNillableDamage(i *int) *WuGongUpdateOne {
	if i != nil {
		wguo.SetDamage(*i)
	}
	return wguo
}

// AddDamage adds i to damage.
func (wguo *WuGongUpdateOne) AddDamage(i int) *WuGongUpdateOne {
	wguo.mutation.AddDamage(i)
	return wguo
}

// SetLevel sets the level field.
func (wguo *WuGongUpdateOne) SetLevel(u uint8) *WuGongUpdateOne {
	wguo.mutation.ResetLevel()
	wguo.mutation.SetLevel(u)
	return wguo
}

// AddLevel adds u to level.
func (wguo *WuGongUpdateOne) AddLevel(u uint8) *WuGongUpdateOne {
	wguo.mutation.AddLevel(u)
	return wguo
}

// Save executes the query and returns the updated entity.
func (wguo *WuGongUpdateOne) Save(ctx context.Context) (*WuGong, error) {
	if _, ok := wguo.mutation.UpdatedAt(); !ok {
		v := wugong.UpdateDefaultUpdatedAt()
		wguo.mutation.SetUpdatedAt(v)
	}
	if v, ok := wguo.mutation.Level(); ok {
		if err := wugong.LevelValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"level\": %v", err)
		}
	}
	var (
		err  error
		node *WuGong
	)
	if len(wguo.hooks) == 0 {
		node, err = wguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WuGongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wguo.mutation = mutation
			node, err = wguo.sqlSave(ctx)
			return node, err
		})
		for i := len(wguo.hooks) - 1; i >= 0; i-- {
			mut = wguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wguo *WuGongUpdateOne) SaveX(ctx context.Context) *WuGong {
	wg, err := wguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return wg
}

// Exec executes the query on the entity.
func (wguo *WuGongUpdateOne) Exec(ctx context.Context) error {
	_, err := wguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wguo *WuGongUpdateOne) ExecX(ctx context.Context) {
	if err := wguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wguo *WuGongUpdateOne) sqlSave(ctx context.Context) (wg *WuGong, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wugong.Table,
			Columns: wugong.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wugong.FieldID,
			},
		},
	}
	id, ok := wguo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing WuGong.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := wguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: wugong.FieldUpdatedAt,
		})
	}
	if value, ok := wguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wugong.FieldName,
		})
	}
	if value, ok := wguo.mutation.Damage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wugong.FieldDamage,
		})
	}
	if value, ok := wguo.mutation.AddedDamage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wugong.FieldDamage,
		})
	}
	if value, ok := wguo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: wugong.FieldLevel,
		})
	}
	if value, ok := wguo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: wugong.FieldLevel,
		})
	}
	wg = &WuGong{config: wguo.config}
	_spec.Assign = wg.assignValues
	_spec.ScanValues = wg.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wugong.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return wg, nil
}
