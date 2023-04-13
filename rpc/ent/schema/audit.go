package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Audit struct {
	ent.Schema
}

func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("object_name").Comment("Effect Object"),
		field.String("action_name").Comment("Effect action"),
		field.String("changed_data").Comment("Changed Data"),
	}
}

func (Audit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		//mixins.StatusMixin{},
	}
}

func (Audit) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Audit) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Audit) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "sys_audits"}}
}
