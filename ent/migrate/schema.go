// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// JiangHuRensColumns holds the columns for the "jiang_hu_rens" table.
	JiangHuRensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeUint},
	}
	// JiangHuRensTable holds the schema information for the "jiang_hu_rens" table.
	JiangHuRensTable = &schema.Table{
		Name:        "jiang_hu_rens",
		Columns:     JiangHuRensColumns,
		PrimaryKey:  []*schema.Column{JiangHuRensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MenPaisColumns holds the columns for the "men_pais" table.
	MenPaisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "address", Type: field.TypeString, Default: ""},
	}
	// MenPaisTable holds the schema information for the "men_pais" table.
	MenPaisTable = &schema.Table{
		Name:        "men_pais",
		Columns:     MenPaisColumns,
		PrimaryKey:  []*schema.Column{MenPaisColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// WuGongsColumns holds the columns for the "wu_gongs" table.
	WuGongsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "damage", Type: field.TypeInt},
		{Name: "level", Type: field.TypeUint8},
	}
	// WuGongsTable holds the schema information for the "wu_gongs" table.
	WuGongsTable = &schema.Table{
		Name:        "wu_gongs",
		Columns:     WuGongsColumns,
		PrimaryKey:  []*schema.Column{WuGongsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JiangHuRensTable,
		MenPaisTable,
		WuGongsTable,
	}
)

func init() {
}
