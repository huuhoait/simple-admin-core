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
	"github.com/suyuan32/simple-admin-core/rpc/ent/audit"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
)

// AuditUpdate is the builder for updating Audit entities.
type AuditUpdate struct {
	config
	hooks    []Hook
	mutation *AuditMutation
}

// Where appends a list predicates to the AuditUpdate builder.
func (au *AuditUpdate) Where(ps ...predicate.Audit) *AuditUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AuditUpdate) SetUpdatedAt(t time.Time) *AuditUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetObjectName sets the "object_name" field.
func (au *AuditUpdate) SetObjectName(s string) *AuditUpdate {
	au.mutation.SetObjectName(s)
	return au
}

// SetActionName sets the "action_name" field.
func (au *AuditUpdate) SetActionName(s string) *AuditUpdate {
	au.mutation.SetActionName(s)
	return au
}

// SetChangedData sets the "changed_data" field.
func (au *AuditUpdate) SetChangedData(s string) *AuditUpdate {
	au.mutation.SetChangedData(s)
	return au
}

// Mutation returns the AuditMutation object of the builder.
func (au *AuditUpdate) Mutation() *AuditMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuditUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks[int, AuditMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuditUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuditUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuditUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AuditUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := audit.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AuditUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(audit.Table, audit.Columns, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(audit.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.ObjectName(); ok {
		_spec.SetField(audit.FieldObjectName, field.TypeString, value)
	}
	if value, ok := au.mutation.ActionName(); ok {
		_spec.SetField(audit.FieldActionName, field.TypeString, value)
	}
	if value, ok := au.mutation.ChangedData(); ok {
		_spec.SetField(audit.FieldChangedData, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true

	return n, nil
}

// AuditUpdateOne is the builder for updating a single Audit entity.
type AuditUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuditMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AuditUpdateOne) SetUpdatedAt(t time.Time) *AuditUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetObjectName sets the "object_name" field.
func (auo *AuditUpdateOne) SetObjectName(s string) *AuditUpdateOne {
	auo.mutation.SetObjectName(s)
	return auo
}

// SetActionName sets the "action_name" field.
func (auo *AuditUpdateOne) SetActionName(s string) *AuditUpdateOne {
	auo.mutation.SetActionName(s)
	return auo
}

// SetChangedData sets the "changed_data" field.
func (auo *AuditUpdateOne) SetChangedData(s string) *AuditUpdateOne {
	auo.mutation.SetChangedData(s)
	return auo
}

// Mutation returns the AuditMutation object of the builder.
func (auo *AuditUpdateOne) Mutation() *AuditMutation {
	return auo.mutation
}

// Where appends a list predicates to the AuditUpdate builder.
func (auo *AuditUpdateOne) Where(ps ...predicate.Audit) *AuditUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuditUpdateOne) Select(field string, fields ...string) *AuditUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Audit entity.
func (auo *AuditUpdateOne) Save(ctx context.Context) (*Audit, error) {
	auo.defaults()
	return withHooks[*Audit, AuditMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuditUpdateOne) SaveX(ctx context.Context) *Audit {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuditUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuditUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AuditUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := audit.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AuditUpdateOne) sqlSave(ctx context.Context) (_node *Audit, err error) {
	_spec := sqlgraph.NewUpdateSpec(audit.Table, audit.Columns, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Audit.id" for update`)}
	}

	//0

	//1

	//2

	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audit.FieldID)
		for _, f := range fields {
			if !audit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != audit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(audit.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.ObjectName(); ok {
		_spec.SetField(audit.FieldObjectName, field.TypeString, value)
	}
	if value, ok := auo.mutation.ActionName(); ok {
		_spec.SetField(audit.FieldActionName, field.TypeString, value)
	}
	if value, ok := auo.mutation.ChangedData(); ok {
		_spec.SetField(audit.FieldChangedData, field.TypeString, value)
	}
	_node = &Audit{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}

	//0

	//1

	//2

	auo.mutation.done = true

	return _node, nil
}
