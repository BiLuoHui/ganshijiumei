// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/BiLuoHui/ganshijiumei/ent"
)

// The JiangHuRenFunc type is an adapter to allow the use of ordinary
// function as JiangHuRen mutator.
type JiangHuRenFunc func(context.Context, *ent.JiangHuRenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JiangHuRenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.JiangHuRenMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JiangHuRenMutation", m)
	}
	return f(ctx, mv)
}

// The MenPaiFunc type is an adapter to allow the use of ordinary
// function as MenPai mutator.
type MenPaiFunc func(context.Context, *ent.MenPaiMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MenPaiFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MenPaiMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MenPaiMutation", m)
	}
	return f(ctx, mv)
}

// The WeaponFunc type is an adapter to allow the use of ordinary
// function as Weapon mutator.
type WeaponFunc func(context.Context, *ent.WeaponMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WeaponFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WeaponMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WeaponMutation", m)
	}
	return f(ctx, mv)
}

// The WuGongFunc type is an adapter to allow the use of ordinary
// function as WuGong mutator.
type WuGongFunc func(context.Context, *ent.WuGongMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f WuGongFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.WuGongMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.WuGongMutation", m)
	}
	return f(ctx, mv)
}

// On executes the given hook only of the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(op) {
				return nil, fmt.Errorf("%s operation is not allowed", m.Op())
			}
			return next.Mutate(ctx, m)
		})
	}
}
