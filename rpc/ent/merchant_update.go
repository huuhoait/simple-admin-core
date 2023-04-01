// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchant"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchantmeta"
	"github.com/huuhoait/zero-admin-core/rpc/ent/predicate"
)

// MerchantUpdate is the builder for updating Merchant entities.
type MerchantUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantMutation
}

// Where appends a list predicates to the MerchantUpdate builder.
func (mu *MerchantUpdate) Where(ps ...predicate.Merchant) *MerchantUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedBy sets the "created_by" field.
func (mu *MerchantUpdate) SetCreatedBy(s string) *MerchantUpdate {
	mu.mutation.SetCreatedBy(s)
	return mu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableCreatedBy(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetCreatedBy(*s)
	}
	return mu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (mu *MerchantUpdate) ClearCreatedBy() *MerchantUpdate {
	mu.mutation.ClearCreatedBy()
	return mu
}

// SetUpdatedBy sets the "updated_by" field.
func (mu *MerchantUpdate) SetUpdatedBy(s string) *MerchantUpdate {
	mu.mutation.SetUpdatedBy(s)
	return mu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableUpdatedBy(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetUpdatedBy(*s)
	}
	return mu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mu *MerchantUpdate) ClearUpdatedBy() *MerchantUpdate {
	mu.mutation.ClearUpdatedBy()
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MerchantUpdate) SetUpdatedAt(t time.Time) *MerchantUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MerchantUpdate) SetStatus(u uint8) *MerchantUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(u)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableStatus(u *uint8) *MerchantUpdate {
	if u != nil {
		mu.SetStatus(*u)
	}
	return mu
}

// AddStatus adds u to the "status" field.
func (mu *MerchantUpdate) AddStatus(u int8) *MerchantUpdate {
	mu.mutation.AddStatus(u)
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *MerchantUpdate) ClearStatus() *MerchantUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetSort sets the "sort" field.
func (mu *MerchantUpdate) SetSort(u uint32) *MerchantUpdate {
	mu.mutation.ResetSort()
	mu.mutation.SetSort(u)
	return mu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableSort(u *uint32) *MerchantUpdate {
	if u != nil {
		mu.SetSort(*u)
	}
	return mu
}

// AddSort adds u to the "sort" field.
func (mu *MerchantUpdate) AddSort(u int32) *MerchantUpdate {
	mu.mutation.AddSort(u)
	return mu
}

// SetName sets the "name" field.
func (mu *MerchantUpdate) SetName(s string) *MerchantUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetLeader sets the "leader" field.
func (mu *MerchantUpdate) SetLeader(s string) *MerchantUpdate {
	mu.mutation.SetLeader(s)
	return mu
}

// SetPhone sets the "phone" field.
func (mu *MerchantUpdate) SetPhone(s string) *MerchantUpdate {
	mu.mutation.SetPhone(s)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MerchantUpdate) SetEmail(s string) *MerchantUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetRemark sets the "remark" field.
func (mu *MerchantUpdate) SetRemark(s string) *MerchantUpdate {
	mu.mutation.SetRemark(s)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MerchantUpdate) SetParentID(u uint64) *MerchantUpdate {
	mu.mutation.SetParentID(u)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableParentID(u *uint64) *MerchantUpdate {
	if u != nil {
		mu.SetParentID(*u)
	}
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MerchantUpdate) ClearParentID() *MerchantUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetParent sets the "parent" edge to the Merchant entity.
func (mu *MerchantUpdate) SetParent(m *Merchant) *MerchantUpdate {
	return mu.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Merchant entity by IDs.
func (mu *MerchantUpdate) AddChildIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Merchant entity.
func (mu *MerchantUpdate) AddChildren(m ...*Merchant) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// AddMerchantMetumIDs adds the "merchant_meta" edge to the MerchantMeta entity by IDs.
func (mu *MerchantUpdate) AddMerchantMetumIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.AddMerchantMetumIDs(ids...)
	return mu
}

// AddMerchantMeta adds the "merchant_meta" edges to the MerchantMeta entity.
func (mu *MerchantUpdate) AddMerchantMeta(m ...*MerchantMeta) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddMerchantMetumIDs(ids...)
}

// Mutation returns the MerchantMutation object of the builder.
func (mu *MerchantUpdate) Mutation() *MerchantMutation {
	return mu.mutation
}

// ClearParent clears the "parent" edge to the Merchant entity.
func (mu *MerchantUpdate) ClearParent() *MerchantUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearChildren clears all "children" edges to the Merchant entity.
func (mu *MerchantUpdate) ClearChildren() *MerchantUpdate {
	mu.mutation.ClearChildren()
	return mu
}

// RemoveChildIDs removes the "children" edge to Merchant entities by IDs.
func (mu *MerchantUpdate) RemoveChildIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Merchant entities.
func (mu *MerchantUpdate) RemoveChildren(m ...*Merchant) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// ClearMerchantMeta clears all "merchant_meta" edges to the MerchantMeta entity.
func (mu *MerchantUpdate) ClearMerchantMeta() *MerchantUpdate {
	mu.mutation.ClearMerchantMeta()
	return mu
}

// RemoveMerchantMetumIDs removes the "merchant_meta" edge to MerchantMeta entities by IDs.
func (mu *MerchantUpdate) RemoveMerchantMetumIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.RemoveMerchantMetumIDs(ids...)
	return mu
}

// RemoveMerchantMeta removes "merchant_meta" edges to MerchantMeta entities.
func (mu *MerchantUpdate) RemoveMerchantMeta(m ...*MerchantMeta) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveMerchantMetumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MerchantUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks[int, MerchantMutation](ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MerchantUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MerchantUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MerchantUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MerchantUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := merchant.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MerchantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedBy(); ok {
		_spec.SetField(merchant.FieldCreatedBy, field.TypeString, value)
	}
	if mu.mutation.CreatedByCleared() {
		_spec.ClearField(merchant.FieldCreatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.UpdatedBy(); ok {
		_spec.SetField(merchant.FieldUpdatedBy, field.TypeString, value)
	}
	if mu.mutation.UpdatedByCleared() {
		_spec.ClearField(merchant.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(merchant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(merchant.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.AddField(merchant.FieldStatus, field.TypeUint8, value)
	}
	if mu.mutation.StatusCleared() {
		_spec.ClearField(merchant.FieldStatus, field.TypeUint8)
	}
	if value, ok := mu.mutation.Sort(); ok {
		_spec.SetField(merchant.FieldSort, field.TypeUint32, value)
	}
	if value, ok := mu.mutation.AddedSort(); ok {
		_spec.AddField(merchant.FieldSort, field.TypeUint32, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(merchant.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Leader(); ok {
		_spec.SetField(merchant.FieldLeader, field.TypeString, value)
	}
	if value, ok := mu.mutation.Phone(); ok {
		_spec.SetField(merchant.FieldPhone, field.TypeString, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(merchant.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Remark(); ok {
		_spec.SetField(merchant.FieldRemark, field.TypeString, value)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchant.ParentTable,
			Columns: []string{merchant.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchant.ParentTable,
			Columns: []string{merchant.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MerchantMetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedMerchantMetaIDs(); len(nodes) > 0 && !mu.mutation.MerchantMetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MerchantMetaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true

	return n, nil
}

// MerchantUpdateOne is the builder for updating a single Merchant entity.
type MerchantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantMutation
}

// SetCreatedBy sets the "created_by" field.
func (muo *MerchantUpdateOne) SetCreatedBy(s string) *MerchantUpdateOne {
	muo.mutation.SetCreatedBy(s)
	return muo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableCreatedBy(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetCreatedBy(*s)
	}
	return muo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (muo *MerchantUpdateOne) ClearCreatedBy() *MerchantUpdateOne {
	muo.mutation.ClearCreatedBy()
	return muo
}

// SetUpdatedBy sets the "updated_by" field.
func (muo *MerchantUpdateOne) SetUpdatedBy(s string) *MerchantUpdateOne {
	muo.mutation.SetUpdatedBy(s)
	return muo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableUpdatedBy(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetUpdatedBy(*s)
	}
	return muo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (muo *MerchantUpdateOne) ClearUpdatedBy() *MerchantUpdateOne {
	muo.mutation.ClearUpdatedBy()
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MerchantUpdateOne) SetUpdatedAt(t time.Time) *MerchantUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetStatus sets the "status" field.
func (muo *MerchantUpdateOne) SetStatus(u uint8) *MerchantUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(u)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableStatus(u *uint8) *MerchantUpdateOne {
	if u != nil {
		muo.SetStatus(*u)
	}
	return muo
}

// AddStatus adds u to the "status" field.
func (muo *MerchantUpdateOne) AddStatus(u int8) *MerchantUpdateOne {
	muo.mutation.AddStatus(u)
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *MerchantUpdateOne) ClearStatus() *MerchantUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetSort sets the "sort" field.
func (muo *MerchantUpdateOne) SetSort(u uint32) *MerchantUpdateOne {
	muo.mutation.ResetSort()
	muo.mutation.SetSort(u)
	return muo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableSort(u *uint32) *MerchantUpdateOne {
	if u != nil {
		muo.SetSort(*u)
	}
	return muo
}

// AddSort adds u to the "sort" field.
func (muo *MerchantUpdateOne) AddSort(u int32) *MerchantUpdateOne {
	muo.mutation.AddSort(u)
	return muo
}

// SetName sets the "name" field.
func (muo *MerchantUpdateOne) SetName(s string) *MerchantUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetLeader sets the "leader" field.
func (muo *MerchantUpdateOne) SetLeader(s string) *MerchantUpdateOne {
	muo.mutation.SetLeader(s)
	return muo
}

// SetPhone sets the "phone" field.
func (muo *MerchantUpdateOne) SetPhone(s string) *MerchantUpdateOne {
	muo.mutation.SetPhone(s)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MerchantUpdateOne) SetEmail(s string) *MerchantUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetRemark sets the "remark" field.
func (muo *MerchantUpdateOne) SetRemark(s string) *MerchantUpdateOne {
	muo.mutation.SetRemark(s)
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MerchantUpdateOne) SetParentID(u uint64) *MerchantUpdateOne {
	muo.mutation.SetParentID(u)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableParentID(u *uint64) *MerchantUpdateOne {
	if u != nil {
		muo.SetParentID(*u)
	}
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MerchantUpdateOne) ClearParentID() *MerchantUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetParent sets the "parent" edge to the Merchant entity.
func (muo *MerchantUpdateOne) SetParent(m *Merchant) *MerchantUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Merchant entity by IDs.
func (muo *MerchantUpdateOne) AddChildIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Merchant entity.
func (muo *MerchantUpdateOne) AddChildren(m ...*Merchant) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// AddMerchantMetumIDs adds the "merchant_meta" edge to the MerchantMeta entity by IDs.
func (muo *MerchantUpdateOne) AddMerchantMetumIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.AddMerchantMetumIDs(ids...)
	return muo
}

// AddMerchantMeta adds the "merchant_meta" edges to the MerchantMeta entity.
func (muo *MerchantUpdateOne) AddMerchantMeta(m ...*MerchantMeta) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddMerchantMetumIDs(ids...)
}

// Mutation returns the MerchantMutation object of the builder.
func (muo *MerchantUpdateOne) Mutation() *MerchantMutation {
	return muo.mutation
}

// ClearParent clears the "parent" edge to the Merchant entity.
func (muo *MerchantUpdateOne) ClearParent() *MerchantUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearChildren clears all "children" edges to the Merchant entity.
func (muo *MerchantUpdateOne) ClearChildren() *MerchantUpdateOne {
	muo.mutation.ClearChildren()
	return muo
}

// RemoveChildIDs removes the "children" edge to Merchant entities by IDs.
func (muo *MerchantUpdateOne) RemoveChildIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Merchant entities.
func (muo *MerchantUpdateOne) RemoveChildren(m ...*Merchant) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
}

// ClearMerchantMeta clears all "merchant_meta" edges to the MerchantMeta entity.
func (muo *MerchantUpdateOne) ClearMerchantMeta() *MerchantUpdateOne {
	muo.mutation.ClearMerchantMeta()
	return muo
}

// RemoveMerchantMetumIDs removes the "merchant_meta" edge to MerchantMeta entities by IDs.
func (muo *MerchantUpdateOne) RemoveMerchantMetumIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.RemoveMerchantMetumIDs(ids...)
	return muo
}

// RemoveMerchantMeta removes "merchant_meta" edges to MerchantMeta entities.
func (muo *MerchantUpdateOne) RemoveMerchantMeta(m ...*MerchantMeta) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveMerchantMetumIDs(ids...)
}

// Where appends a list predicates to the MerchantUpdate builder.
func (muo *MerchantUpdateOne) Where(ps ...predicate.Merchant) *MerchantUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MerchantUpdateOne) Select(field string, fields ...string) *MerchantUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Merchant entity.
func (muo *MerchantUpdateOne) Save(ctx context.Context) (*Merchant, error) {
	muo.defaults()
	return withHooks[*Merchant, MerchantMutation](ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MerchantUpdateOne) SaveX(ctx context.Context) *Merchant {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MerchantUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MerchantUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MerchantUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := merchant.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MerchantUpdateOne) sqlSave(ctx context.Context) (_node *Merchant, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Merchant.id" for update`)}
	}

	//0

	//1

	//2

	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchant.FieldID)
		for _, f := range fields {
			if !merchant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedBy(); ok {
		_spec.SetField(merchant.FieldCreatedBy, field.TypeString, value)
	}
	if muo.mutation.CreatedByCleared() {
		_spec.ClearField(merchant.FieldCreatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.UpdatedBy(); ok {
		_spec.SetField(merchant.FieldUpdatedBy, field.TypeString, value)
	}
	if muo.mutation.UpdatedByCleared() {
		_spec.ClearField(merchant.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(merchant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(merchant.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.AddField(merchant.FieldStatus, field.TypeUint8, value)
	}
	if muo.mutation.StatusCleared() {
		_spec.ClearField(merchant.FieldStatus, field.TypeUint8)
	}
	if value, ok := muo.mutation.Sort(); ok {
		_spec.SetField(merchant.FieldSort, field.TypeUint32, value)
	}
	if value, ok := muo.mutation.AddedSort(); ok {
		_spec.AddField(merchant.FieldSort, field.TypeUint32, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(merchant.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Leader(); ok {
		_spec.SetField(merchant.FieldLeader, field.TypeString, value)
	}
	if value, ok := muo.mutation.Phone(); ok {
		_spec.SetField(merchant.FieldPhone, field.TypeString, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(merchant.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Remark(); ok {
		_spec.SetField(merchant.FieldRemark, field.TypeString, value)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchant.ParentTable,
			Columns: []string{merchant.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchant.ParentTable,
			Columns: []string{merchant.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.ChildrenTable,
			Columns: []string{merchant.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MerchantMetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedMerchantMetaIDs(); len(nodes) > 0 && !muo.mutation.MerchantMetaCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MerchantMetaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.MerchantMetaTable,
			Columns: []string{merchant.MerchantMetaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Merchant{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}

	//0

	//1

	//2

	muo.mutation.done = true

	return _node, nil
}
