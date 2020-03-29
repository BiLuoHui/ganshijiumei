package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// JiangHuRen holds the schema definition for the JiangHuRen entity.
type JiangHuRen struct {
	ent.Schema
}

// Fields of the JiangHuRen.
func (JiangHuRen) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Comment("江湖名号一定要独一无二"),
		field.Uint("age").
			Comment("姑娘你芳龄几何"),
		field.Bool("sex").
			Comment("性别"),
	}
}

// Edges of the JiangHuRen.
func (JiangHuRen) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("weapon", Weapon.Type).
			Unique(),
		edge.From("menpai", MenPai.Type).
			Ref("disciples").
			Unique().
			Comment("所属门派"),
		edge.To("spouse", JiangHuRen.Type).
			Unique().
			Comment("夫妻"),
		edge.To("apprentices", JiangHuRen.Type).
			From("master").
			Unique().
			Comment("师徒"),
		edge.To("following", JiangHuRen.Type).
			From("followers").
			Comment("关注与粉丝"),
		edge.To("friends", JiangHuRen.Type).
			Comment("朋友"),
	}
}

func (JiangHuRen) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
