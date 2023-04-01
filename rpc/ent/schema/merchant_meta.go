package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/zero-tools/orm/ent/mixins"
)

type MerchantMeta struct {
	ent.Schema
}

func (MerchantMeta) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("the title shown in the ui | 展示名称 （建议配合i18n）"),
		field.String("key").Comment("key | 键"),
		field.String("value").Comment("value | 值"),
		field.Uint64("merchant_id").Optional().Comment("Merchant ID | 字典ID"),
	}
}

func (MerchantMeta) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (MerchantMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", Merchant.Type).Field("merchant_id").Ref("merchant_meta").Unique(),
	}
}

func (MerchantMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_merchant_meta"},
	}
}
