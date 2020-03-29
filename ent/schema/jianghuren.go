package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// JiangHuRen holds the schema definition for the JiangHuRen entity.
type JiangHuRen struct {
	ent.Schema
}

// Fields of the JiangHuRen.
func (JiangHuRen) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Comment("江湖名号一定要独一无二"),
		field.Uint("age").Comment("姑娘你芳龄几何"),
	}
}

// Edges of the JiangHuRen.
func (JiangHuRen) Edges() []ent.Edge {
	return nil
}

func (JiangHuRen) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
