package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// MenPai holds the schema definition for the MenPai entity.
type MenPai struct {
	ent.Schema
}

// Fields of the MenPai.
func (MenPai) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Comment("门派名号"),
		field.String("address").Default("").Comment("门派府台尊地"),
	}
}

// Edges of the MenPai.
func (MenPai) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("disciples", JiangHuRen.Type).
			Comment("门人弟子"),
	}
}

func (MenPai) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
