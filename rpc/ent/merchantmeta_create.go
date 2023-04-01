// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchant"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchantmeta"
)

// MerchantMetaCreate is the builder for creating a MerchantMeta entity.
type MerchantMetaCreate struct {
	config
	mutation *MerchantMetaMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (mmc *MerchantMetaCreate) SetCreatedBy(s string) *MerchantMetaCreate {
	mmc.mutation.SetCreatedBy(s)
	return mmc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableCreatedBy(s *string) *MerchantMetaCreate {
	if s != nil {
		mmc.SetCreatedBy(*s)
	}
	return mmc
}

// SetCreatedAt sets the "created_at" field.
func (mmc *MerchantMetaCreate) SetCreatedAt(t time.Time) *MerchantMetaCreate {
	mmc.mutation.SetCreatedAt(t)
	return mmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableCreatedAt(t *time.Time) *MerchantMetaCreate {
	if t != nil {
		mmc.SetCreatedAt(*t)
	}
	return mmc
}

// SetUpdatedBy sets the "updated_by" field.
func (mmc *MerchantMetaCreate) SetUpdatedBy(s string) *MerchantMetaCreate {
	mmc.mutation.SetUpdatedBy(s)
	return mmc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableUpdatedBy(s *string) *MerchantMetaCreate {
	if s != nil {
		mmc.SetUpdatedBy(*s)
	}
	return mmc
}

// SetUpdatedAt sets the "updated_at" field.
func (mmc *MerchantMetaCreate) SetUpdatedAt(t time.Time) *MerchantMetaCreate {
	mmc.mutation.SetUpdatedAt(t)
	return mmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableUpdatedAt(t *time.Time) *MerchantMetaCreate {
	if t != nil {
		mmc.SetUpdatedAt(*t)
	}
	return mmc
}

// SetStatus sets the "status" field.
func (mmc *MerchantMetaCreate) SetStatus(u uint8) *MerchantMetaCreate {
	mmc.mutation.SetStatus(u)
	return mmc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableStatus(u *uint8) *MerchantMetaCreate {
	if u != nil {
		mmc.SetStatus(*u)
	}
	return mmc
}

// SetSort sets the "sort" field.
func (mmc *MerchantMetaCreate) SetSort(u uint32) *MerchantMetaCreate {
	mmc.mutation.SetSort(u)
	return mmc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableSort(u *uint32) *MerchantMetaCreate {
	if u != nil {
		mmc.SetSort(*u)
	}
	return mmc
}

// SetTitle sets the "title" field.
func (mmc *MerchantMetaCreate) SetTitle(s string) *MerchantMetaCreate {
	mmc.mutation.SetTitle(s)
	return mmc
}

// SetKey sets the "key" field.
func (mmc *MerchantMetaCreate) SetKey(s string) *MerchantMetaCreate {
	mmc.mutation.SetKey(s)
	return mmc
}

// SetValue sets the "value" field.
func (mmc *MerchantMetaCreate) SetValue(s string) *MerchantMetaCreate {
	mmc.mutation.SetValue(s)
	return mmc
}

// SetMerchantID sets the "merchant_id" field.
func (mmc *MerchantMetaCreate) SetMerchantID(u uint64) *MerchantMetaCreate {
	mmc.mutation.SetMerchantID(u)
	return mmc
}

// SetNillableMerchantID sets the "merchant_id" field if the given value is not nil.
func (mmc *MerchantMetaCreate) SetNillableMerchantID(u *uint64) *MerchantMetaCreate {
	if u != nil {
		mmc.SetMerchantID(*u)
	}
	return mmc
}

// SetID sets the "id" field.
func (mmc *MerchantMetaCreate) SetID(u uint64) *MerchantMetaCreate {
	mmc.mutation.SetID(u)
	return mmc
}

// SetMerchant sets the "merchant" edge to the Merchant entity.
func (mmc *MerchantMetaCreate) SetMerchant(m *Merchant) *MerchantMetaCreate {
	return mmc.SetMerchantID(m.ID)
}

// Mutation returns the MerchantMetaMutation object of the builder.
func (mmc *MerchantMetaCreate) Mutation() *MerchantMetaMutation {
	return mmc.mutation
}

// Save creates the MerchantMeta in the database.
func (mmc *MerchantMetaCreate) Save(ctx context.Context) (*MerchantMeta, error) {
	mmc.defaults()
	return withHooks[*MerchantMeta, MerchantMetaMutation](ctx, mmc.sqlSave, mmc.mutation, mmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mmc *MerchantMetaCreate) SaveX(ctx context.Context) *MerchantMeta {
	v, err := mmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmc *MerchantMetaCreate) Exec(ctx context.Context) error {
	_, err := mmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmc *MerchantMetaCreate) ExecX(ctx context.Context) {
	if err := mmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mmc *MerchantMetaCreate) defaults() {
	if _, ok := mmc.mutation.CreatedAt(); !ok {
		v := merchantmeta.DefaultCreatedAt()
		mmc.mutation.SetCreatedAt(v)
	}
	if _, ok := mmc.mutation.UpdatedAt(); !ok {
		v := merchantmeta.DefaultUpdatedAt()
		mmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mmc.mutation.Status(); !ok {
		v := merchantmeta.DefaultStatus
		mmc.mutation.SetStatus(v)
	}
	if _, ok := mmc.mutation.Sort(); !ok {
		v := merchantmeta.DefaultSort
		mmc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mmc *MerchantMetaCreate) check() error {
	if _, ok := mmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MerchantMeta.created_at"`)}
	}
	if _, ok := mmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MerchantMeta.updated_at"`)}
	}
	if _, ok := mmc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "MerchantMeta.sort"`)}
	}
	if _, ok := mmc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "MerchantMeta.title"`)}
	}
	if _, ok := mmc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "MerchantMeta.key"`)}
	}
	if _, ok := mmc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "MerchantMeta.value"`)}
	}
	return nil
}

func (mmc *MerchantMetaCreate) sqlSave(ctx context.Context) (*MerchantMeta, error) {
	if err := mmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	mmc.mutation.id = &_node.ID
	mmc.mutation.done = true
	return _node, nil
}

func (mmc *MerchantMetaCreate) createSpec() (*MerchantMeta, *sqlgraph.CreateSpec) {
	var (
		_node = &MerchantMeta{config: mmc.config}
		_spec = sqlgraph.NewCreateSpec(merchantmeta.Table, sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64))
	)
	if id, ok := mmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mmc.mutation.CreatedBy(); ok {
		_spec.SetField(merchantmeta.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := mmc.mutation.CreatedAt(); ok {
		_spec.SetField(merchantmeta.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mmc.mutation.UpdatedBy(); ok {
		_spec.SetField(merchantmeta.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := mmc.mutation.UpdatedAt(); ok {
		_spec.SetField(merchantmeta.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mmc.mutation.Status(); ok {
		_spec.SetField(merchantmeta.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := mmc.mutation.Sort(); ok {
		_spec.SetField(merchantmeta.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := mmc.mutation.Title(); ok {
		_spec.SetField(merchantmeta.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := mmc.mutation.Key(); ok {
		_spec.SetField(merchantmeta.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := mmc.mutation.Value(); ok {
		_spec.SetField(merchantmeta.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := mmc.mutation.MerchantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   merchantmeta.MerchantTable,
			Columns: []string{merchantmeta.MerchantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MerchantID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MerchantMetaCreateBulk is the builder for creating many MerchantMeta entities in bulk.
type MerchantMetaCreateBulk struct {
	config
	builders []*MerchantMetaCreate
}

// Save creates the MerchantMeta entities in the database.
func (mmcb *MerchantMetaCreateBulk) Save(ctx context.Context) ([]*MerchantMeta, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mmcb.builders))
	nodes := make([]*MerchantMeta, len(mmcb.builders))
	mutators := make([]Mutator, len(mmcb.builders))
	for i := range mmcb.builders {
		func(i int, root context.Context) {
			builder := mmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantMetaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mmcb *MerchantMetaCreateBulk) SaveX(ctx context.Context) []*MerchantMeta {
	v, err := mmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmcb *MerchantMetaCreateBulk) Exec(ctx context.Context) error {
	_, err := mmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmcb *MerchantMetaCreateBulk) ExecX(ctx context.Context) {
	if err := mmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
