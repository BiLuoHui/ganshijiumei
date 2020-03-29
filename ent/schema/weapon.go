package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Weapon holds the schema definition for the Weapon entity.
type Weapon struct {
	ent.Schema
}

// Fields of the Weapon.
func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("level").
			Default(0),
		field.Int("category").
			Default(0).
			Comment("0-其他，1-刀，2-剑"),
	}
}

// Edges of the Weapon.
func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", JiangHuRen.Type).
			Ref("weapon").
			Unique().
			Required(),
	}
}

func (Weapon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
