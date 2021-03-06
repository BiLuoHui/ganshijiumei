// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BiLuoHui/ganshijiumei/ent/jianghuren"
	"github.com/BiLuoHui/ganshijiumei/ent/menpai"
	"github.com/BiLuoHui/ganshijiumei/ent/weapon"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// JiangHuRenCreate is the builder for creating a JiangHuRen entity.
type JiangHuRenCreate struct {
	config
	mutation *JiangHuRenMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (jhrc *JiangHuRenCreate) SetCreatedAt(t time.Time) *JiangHuRenCreate {
	jhrc.mutation.SetCreatedAt(t)
	return jhrc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableCreatedAt(t *time.Time) *JiangHuRenCreate {
	if t != nil {
		jhrc.SetCreatedAt(*t)
	}
	return jhrc
}

// SetUpdatedAt sets the updated_at field.
func (jhrc *JiangHuRenCreate) SetUpdatedAt(t time.Time) *JiangHuRenCreate {
	jhrc.mutation.SetUpdatedAt(t)
	return jhrc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableUpdatedAt(t *time.Time) *JiangHuRenCreate {
	if t != nil {
		jhrc.SetUpdatedAt(*t)
	}
	return jhrc
}

// SetName sets the name field.
func (jhrc *JiangHuRenCreate) SetName(s string) *JiangHuRenCreate {
	jhrc.mutation.SetName(s)
	return jhrc
}

// SetAge sets the age field.
func (jhrc *JiangHuRenCreate) SetAge(u uint) *JiangHuRenCreate {
	jhrc.mutation.SetAge(u)
	return jhrc
}

// SetSex sets the sex field.
func (jhrc *JiangHuRenCreate) SetSex(b bool) *JiangHuRenCreate {
	jhrc.mutation.SetSex(b)
	return jhrc
}

// SetWeaponID sets the weapon edge to Weapon by id.
func (jhrc *JiangHuRenCreate) SetWeaponID(id int) *JiangHuRenCreate {
	jhrc.mutation.SetWeaponID(id)
	return jhrc
}

// SetNillableWeaponID sets the weapon edge to Weapon by id if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableWeaponID(id *int) *JiangHuRenCreate {
	if id != nil {
		jhrc = jhrc.SetWeaponID(*id)
	}
	return jhrc
}

// SetWeapon sets the weapon edge to Weapon.
func (jhrc *JiangHuRenCreate) SetWeapon(w *Weapon) *JiangHuRenCreate {
	return jhrc.SetWeaponID(w.ID)
}

// SetMenpaiID sets the menpai edge to MenPai by id.
func (jhrc *JiangHuRenCreate) SetMenpaiID(id int) *JiangHuRenCreate {
	jhrc.mutation.SetMenpaiID(id)
	return jhrc
}

// SetNillableMenpaiID sets the menpai edge to MenPai by id if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableMenpaiID(id *int) *JiangHuRenCreate {
	if id != nil {
		jhrc = jhrc.SetMenpaiID(*id)
	}
	return jhrc
}

// SetMenpai sets the menpai edge to MenPai.
func (jhrc *JiangHuRenCreate) SetMenpai(m *MenPai) *JiangHuRenCreate {
	return jhrc.SetMenpaiID(m.ID)
}

// SetSpouseID sets the spouse edge to JiangHuRen by id.
func (jhrc *JiangHuRenCreate) SetSpouseID(id int) *JiangHuRenCreate {
	jhrc.mutation.SetSpouseID(id)
	return jhrc
}

// SetNillableSpouseID sets the spouse edge to JiangHuRen by id if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableSpouseID(id *int) *JiangHuRenCreate {
	if id != nil {
		jhrc = jhrc.SetSpouseID(*id)
	}
	return jhrc
}

// SetSpouse sets the spouse edge to JiangHuRen.
func (jhrc *JiangHuRenCreate) SetSpouse(j *JiangHuRen) *JiangHuRenCreate {
	return jhrc.SetSpouseID(j.ID)
}

// SetMasterID sets the master edge to JiangHuRen by id.
func (jhrc *JiangHuRenCreate) SetMasterID(id int) *JiangHuRenCreate {
	jhrc.mutation.SetMasterID(id)
	return jhrc
}

// SetNillableMasterID sets the master edge to JiangHuRen by id if the given value is not nil.
func (jhrc *JiangHuRenCreate) SetNillableMasterID(id *int) *JiangHuRenCreate {
	if id != nil {
		jhrc = jhrc.SetMasterID(*id)
	}
	return jhrc
}

// SetMaster sets the master edge to JiangHuRen.
func (jhrc *JiangHuRenCreate) SetMaster(j *JiangHuRen) *JiangHuRenCreate {
	return jhrc.SetMasterID(j.ID)
}

// AddApprenticeIDs adds the apprentices edge to JiangHuRen by ids.
func (jhrc *JiangHuRenCreate) AddApprenticeIDs(ids ...int) *JiangHuRenCreate {
	jhrc.mutation.AddApprenticeIDs(ids...)
	return jhrc
}

// AddApprentices adds the apprentices edges to JiangHuRen.
func (jhrc *JiangHuRenCreate) AddApprentices(j ...*JiangHuRen) *JiangHuRenCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jhrc.AddApprenticeIDs(ids...)
}

// AddFollowerIDs adds the followers edge to JiangHuRen by ids.
func (jhrc *JiangHuRenCreate) AddFollowerIDs(ids ...int) *JiangHuRenCreate {
	jhrc.mutation.AddFollowerIDs(ids...)
	return jhrc
}

// AddFollowers adds the followers edges to JiangHuRen.
func (jhrc *JiangHuRenCreate) AddFollowers(j ...*JiangHuRen) *JiangHuRenCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jhrc.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to JiangHuRen by ids.
func (jhrc *JiangHuRenCreate) AddFollowingIDs(ids ...int) *JiangHuRenCreate {
	jhrc.mutation.AddFollowingIDs(ids...)
	return jhrc
}

// AddFollowing adds the following edges to JiangHuRen.
func (jhrc *JiangHuRenCreate) AddFollowing(j ...*JiangHuRen) *JiangHuRenCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jhrc.AddFollowingIDs(ids...)
}

// AddFriendIDs adds the friends edge to JiangHuRen by ids.
func (jhrc *JiangHuRenCreate) AddFriendIDs(ids ...int) *JiangHuRenCreate {
	jhrc.mutation.AddFriendIDs(ids...)
	return jhrc
}

// AddFriends adds the friends edges to JiangHuRen.
func (jhrc *JiangHuRenCreate) AddFriends(j ...*JiangHuRen) *JiangHuRenCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jhrc.AddFriendIDs(ids...)
}

// Save creates the JiangHuRen in the database.
func (jhrc *JiangHuRenCreate) Save(ctx context.Context) (*JiangHuRen, error) {
	if _, ok := jhrc.mutation.CreatedAt(); !ok {
		v := jianghuren.DefaultCreatedAt()
		jhrc.mutation.SetCreatedAt(v)
	}
	if _, ok := jhrc.mutation.UpdatedAt(); !ok {
		v := jianghuren.DefaultUpdatedAt()
		jhrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := jhrc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if _, ok := jhrc.mutation.Age(); !ok {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if _, ok := jhrc.mutation.Sex(); !ok {
		return nil, errors.New("ent: missing required field \"sex\"")
	}
	var (
		err  error
		node *JiangHuRen
	)
	if len(jhrc.hooks) == 0 {
		node, err = jhrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JiangHuRenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jhrc.mutation = mutation
			node, err = jhrc.sqlSave(ctx)
			return node, err
		})
		for i := len(jhrc.hooks) - 1; i >= 0; i-- {
			mut = jhrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jhrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jhrc *JiangHuRenCreate) SaveX(ctx context.Context) *JiangHuRen {
	v, err := jhrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jhrc *JiangHuRenCreate) sqlSave(ctx context.Context) (*JiangHuRen, error) {
	var (
		jhr   = &JiangHuRen{config: jhrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: jianghuren.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jianghuren.FieldID,
			},
		}
	)
	if value, ok := jhrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: jianghuren.FieldCreatedAt,
		})
		jhr.CreatedAt = value
	}
	if value, ok := jhrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: jianghuren.FieldUpdatedAt,
		})
		jhr.UpdatedAt = value
	}
	if value, ok := jhrc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: jianghuren.FieldName,
		})
		jhr.Name = value
	}
	if value, ok := jhrc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: jianghuren.FieldAge,
		})
		jhr.Age = value
	}
	if value, ok := jhrc.mutation.Sex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: jianghuren.FieldSex,
		})
		jhr.Sex = value
	}
	if nodes := jhrc.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   jianghuren.WeaponTable,
			Columns: []string{jianghuren.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: weapon.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.MenpaiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jianghuren.MenpaiTable,
			Columns: []string{jianghuren.MenpaiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menpai.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   jianghuren.SpouseTable,
			Columns: []string{jianghuren.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.MasterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jianghuren.MasterTable,
			Columns: []string{jianghuren.MasterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.ApprenticesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jianghuren.ApprenticesTable,
			Columns: []string{jianghuren.ApprenticesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   jianghuren.FollowersTable,
			Columns: jianghuren.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   jianghuren.FollowingTable,
			Columns: jianghuren.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jhrc.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   jianghuren.FriendsTable,
			Columns: jianghuren.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: jianghuren.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, jhrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	jhr.ID = int(id)
	return jhr, nil
}
