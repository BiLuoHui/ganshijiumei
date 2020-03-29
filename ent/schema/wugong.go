package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// WuGong holds the schema definition for the WuGong entity.
type WuGong struct {
	ent.Schema
}

// Fields of the WuGong.
func (WuGong) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("武功名字"),
		field.Int("damage").Default(0).Comment("伤害值"),
		field.Uint8("level").Max(10).Comment("最高等级，为了防止外挂，最高只能10级，你不能打出十一级伤害"),
	}
}

// Edges of the WuGong.
func (WuGong) Edges() []ent.Edge {
	return nil
}

func (WuGong) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
