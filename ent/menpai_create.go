// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BiLuoHui/ganshijiumei/ent/menpai"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MenPaiCreate is the builder for creating a MenPai entity.
type MenPaiCreate struct {
	config
	mutation *MenPaiMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (mpc *MenPaiCreate) SetCreatedAt(t time.Time) *MenPaiCreate {
	mpc.mutation.SetCreatedAt(t)
	return mpc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (mpc *MenPaiCreate) SetNillableCreatedAt(t *time.Time) *MenPaiCreate {
	if t != nil {
		mpc.SetCreatedAt(*t)
	}
	return mpc
}

// SetUpdatedAt sets the updated_at field.
func (mpc *MenPaiCreate) SetUpdatedAt(t time.Time) *MenPaiCreate {
	mpc.mutation.SetUpdatedAt(t)
	return mpc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (mpc *MenPaiCreate) SetNillableUpdatedAt(t *time.Time) *MenPaiCreate {
	if t != nil {
		mpc.SetUpdatedAt(*t)
	}
	return mpc
}

// SetName sets the name field.
func (mpc *MenPaiCreate) SetName(s string) *MenPaiCreate {
	mpc.mutation.SetName(s)
	return mpc
}

// SetAddress sets the address field.
func (mpc *MenPaiCreate) SetAddress(s string) *MenPaiCreate {
	mpc.mutation.SetAddress(s)
	return mpc
}

// SetNillableAddress sets the address field if the given value is not nil.
func (mpc *MenPaiCreate) SetNillableAddress(s *string) *MenPaiCreate {
	if s != nil {
		mpc.SetAddress(*s)
	}
	return mpc
}

// Save creates the MenPai in the database.
func (mpc *MenPaiCreate) Save(ctx context.Context) (*MenPai, error) {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		v := menpai.DefaultCreatedAt()
		mpc.mutation.SetCreatedAt(v)
	}
	if _, ok := mpc.mutation.UpdatedAt(); !ok {
		v := menpai.DefaultUpdatedAt()
		mpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mpc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if _, ok := mpc.mutation.Address(); !ok {
		v := menpai.DefaultAddress
		mpc.mutation.SetAddress(v)
	}
	var (
		err  error
		node *MenPai
	)
	if len(mpc.hooks) == 0 {
		node, err = mpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenPaiMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mpc.mutation = mutation
			node, err = mpc.sqlSave(ctx)
			return node, err
		})
		for i := len(mpc.hooks) - 1; i >= 0; i-- {
			mut = mpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mpc *MenPaiCreate) SaveX(ctx context.Context) *MenPai {
	v, err := mpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mpc *MenPaiCreate) sqlSave(ctx context.Context) (*MenPai, error) {
	var (
		mp    = &MenPai{config: mpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: menpai.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menpai.FieldID,
			},
		}
	)
	if value, ok := mpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menpai.FieldCreatedAt,
		})
		mp.CreatedAt = value
	}
	if value, ok := mpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menpai.FieldUpdatedAt,
		})
		mp.UpdatedAt = value
	}
	if value, ok := mpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menpai.FieldName,
		})
		mp.Name = value
	}
	if value, ok := mpc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menpai.FieldAddress,
		})
		mp.Address = value
	}
	if err := sqlgraph.CreateNode(ctx, mpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	mp.ID = int(id)
	return mp, nil
}