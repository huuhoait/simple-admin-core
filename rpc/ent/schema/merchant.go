package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/zero-tools/orm/ent/mixins"
)

type Merchant struct {
	ent.Schema
}

func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Merchant name "),
		field.String("leader").Comment("Merchant leader"),
		field.String("phone").Comment("Merchant's phone number "),
		field.String("email").Comment("Merchant's email"),
		field.String("remark").Comment("Remark"),
		field.Uint64("parent_id").Optional().Default(0).Comment("Parent Merchant ID | 父级部门ID"),
	}
}

func (Merchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseIDMixin{},
		mixins.StatusMixin{},
		mixins.SortMixin{},
	}
}

func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Merchant.Type).From("parent").Unique().Field("parent_id"),
		edge.To("merchant_meta", MerchantMeta.Type),
	}
}

func (Merchant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_Merchants"},
	}
}

/*
func (Merchant) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(VersionHook(), ent.OpUpdateOne),
	}
}


func VersionHook() ent.Hook {

	return func(next ent.Mutator) ent.Mutator {
		// A hook that validates the "version" field is incremented by 1 on each update.
		// Note that this is just a dummy example, and it doesn't promise consistency in
		// the database.
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			fmt.Println("Hook update")
			// Add an SQL predicate that validates the "version" column is equal
			// to "oldV" (it wasn't changed during the mutation by other process).
			return next.Mutate(ctx, m)
		})
	}
}*/
