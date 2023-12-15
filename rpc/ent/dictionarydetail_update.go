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
	"github.com/huuhoait/simple-admin-core/rpc/ent/dictionary"
	"github.com/huuhoait/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/huuhoait/simple-admin-core/rpc/ent/predicate"
)

// DictionaryDetailUpdate is the builder for updating DictionaryDetail entities.
type DictionaryDetailUpdate struct {
	config
	hooks    []Hook
	mutation *DictionaryDetailMutation
}

// Where appends a list predicates to the DictionaryDetailUpdate builder.
func (ddu *DictionaryDetailUpdate) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetUpdatedAt sets the "updated_at" field.
func (ddu *DictionaryDetailUpdate) SetUpdatedAt(t time.Time) *DictionaryDetailUpdate {
	ddu.mutation.SetUpdatedAt(t)
	return ddu
}

// SetStatus sets the "status" field.
func (ddu *DictionaryDetailUpdate) SetStatus(u uint8) *DictionaryDetailUpdate {
	ddu.mutation.ResetStatus()
	ddu.mutation.SetStatus(u)
	return ddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableStatus(u *uint8) *DictionaryDetailUpdate {
	if u != nil {
		ddu.SetStatus(*u)
	}
	return ddu
}

// AddStatus adds u to the "status" field.
func (ddu *DictionaryDetailUpdate) AddStatus(u int8) *DictionaryDetailUpdate {
	ddu.mutation.AddStatus(u)
	return ddu
}

// ClearStatus clears the value of the "status" field.
func (ddu *DictionaryDetailUpdate) ClearStatus() *DictionaryDetailUpdate {
	ddu.mutation.ClearStatus()
	return ddu
}

// SetSort sets the "sort" field.
func (ddu *DictionaryDetailUpdate) SetSort(u uint32) *DictionaryDetailUpdate {
	ddu.mutation.ResetSort()
	ddu.mutation.SetSort(u)
	return ddu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableSort(u *uint32) *DictionaryDetailUpdate {
	if u != nil {
		ddu.SetSort(*u)
	}
	return ddu
}

// AddSort adds u to the "sort" field.
func (ddu *DictionaryDetailUpdate) AddSort(u int32) *DictionaryDetailUpdate {
	ddu.mutation.AddSort(u)
	return ddu
}

// SetTitle sets the "title" field.
func (ddu *DictionaryDetailUpdate) SetTitle(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetTitle(s)
	return ddu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableTitle(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetTitle(*s)
	}
	return ddu
}

// SetKey sets the "key" field.
func (ddu *DictionaryDetailUpdate) SetKey(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetKey(s)
	return ddu
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableKey(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetKey(*s)
	}
	return ddu
}

// SetValue sets the "value" field.
func (ddu *DictionaryDetailUpdate) SetValue(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetValue(s)
	return ddu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableValue(s *string) *DictionaryDetailUpdate {
	if s != nil {
		ddu.SetValue(*s)
	}
	return ddu
}

// SetDictionaryID sets the "dictionary_id" field.
func (ddu *DictionaryDetailUpdate) SetDictionaryID(u uint64) *DictionaryDetailUpdate {
	ddu.mutation.SetDictionaryID(u)
	return ddu
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableDictionaryID(u *uint64) *DictionaryDetailUpdate {
	if u != nil {
		ddu.SetDictionaryID(*u)
	}
	return ddu
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (ddu *DictionaryDetailUpdate) ClearDictionaryID() *DictionaryDetailUpdate {
	ddu.mutation.ClearDictionaryID()
	return ddu
}

// SetDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID.
func (ddu *DictionaryDetailUpdate) SetDictionariesID(id uint64) *DictionaryDetailUpdate {
	ddu.mutation.SetDictionariesID(id)
	return ddu
}

// SetNillableDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableDictionariesID(id *uint64) *DictionaryDetailUpdate {
	if id != nil {
		ddu = ddu.SetDictionariesID(*id)
	}
	return ddu
}

// SetDictionaries sets the "dictionaries" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) SetDictionaries(d *Dictionary) *DictionaryDetailUpdate {
	return ddu.SetDictionariesID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (ddu *DictionaryDetailUpdate) Mutation() *DictionaryDetailMutation {
	return ddu.mutation
}

// ClearDictionaries clears the "dictionaries" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) ClearDictionaries() *DictionaryDetailUpdate {
	ddu.mutation.ClearDictionaries()
	return ddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DictionaryDetailUpdate) Save(ctx context.Context) (int, error) {
	ddu.defaults()
	return withHooks(ctx, ddu.sqlSave, ddu.mutation, ddu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DictionaryDetailUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddu *DictionaryDetailUpdate) defaults() {
	if _, ok := ddu.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		ddu.mutation.SetUpdatedAt(v)
	}
}

func (ddu *DictionaryDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionarydetail.Table, dictionarydetail.Columns, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64))
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddu.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ddu.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ddu.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if ddu.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeUint8)
	}
	if value, ok := ddu.mutation.Sort(); ok {
		_spec.SetField(dictionarydetail.FieldSort, field.TypeUint32, value)
	}
	if value, ok := ddu.mutation.AddedSort(); ok {
		_spec.AddField(dictionarydetail.FieldSort, field.TypeUint32, value)
	}
	if value, ok := ddu.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if ddu.mutation.DictionariesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionariesTable,
			Columns: []string{dictionarydetail.DictionariesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ddu.mutation.DictionariesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionariesTable,
			Columns: []string{dictionarydetail.DictionariesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ddu.mutation.done = true
	return n, nil
}

// DictionaryDetailUpdateOne is the builder for updating a single DictionaryDetail entity.
type DictionaryDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DictionaryDetailMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dduo *DictionaryDetailUpdateOne) SetUpdatedAt(t time.Time) *DictionaryDetailUpdateOne {
	dduo.mutation.SetUpdatedAt(t)
	return dduo
}

// SetStatus sets the "status" field.
func (dduo *DictionaryDetailUpdateOne) SetStatus(u uint8) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetStatus()
	dduo.mutation.SetStatus(u)
	return dduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableStatus(u *uint8) *DictionaryDetailUpdateOne {
	if u != nil {
		dduo.SetStatus(*u)
	}
	return dduo
}

// AddStatus adds u to the "status" field.
func (dduo *DictionaryDetailUpdateOne) AddStatus(u int8) *DictionaryDetailUpdateOne {
	dduo.mutation.AddStatus(u)
	return dduo
}

// ClearStatus clears the value of the "status" field.
func (dduo *DictionaryDetailUpdateOne) ClearStatus() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearStatus()
	return dduo
}

// SetSort sets the "sort" field.
func (dduo *DictionaryDetailUpdateOne) SetSort(u uint32) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetSort()
	dduo.mutation.SetSort(u)
	return dduo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableSort(u *uint32) *DictionaryDetailUpdateOne {
	if u != nil {
		dduo.SetSort(*u)
	}
	return dduo
}

// AddSort adds u to the "sort" field.
func (dduo *DictionaryDetailUpdateOne) AddSort(u int32) *DictionaryDetailUpdateOne {
	dduo.mutation.AddSort(u)
	return dduo
}

// SetTitle sets the "title" field.
func (dduo *DictionaryDetailUpdateOne) SetTitle(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetTitle(s)
	return dduo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableTitle(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetTitle(*s)
	}
	return dduo
}

// SetKey sets the "key" field.
func (dduo *DictionaryDetailUpdateOne) SetKey(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetKey(s)
	return dduo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableKey(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetKey(*s)
	}
	return dduo
}

// SetValue sets the "value" field.
func (dduo *DictionaryDetailUpdateOne) SetValue(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetValue(s)
	return dduo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableValue(s *string) *DictionaryDetailUpdateOne {
	if s != nil {
		dduo.SetValue(*s)
	}
	return dduo
}

// SetDictionaryID sets the "dictionary_id" field.
func (dduo *DictionaryDetailUpdateOne) SetDictionaryID(u uint64) *DictionaryDetailUpdateOne {
	dduo.mutation.SetDictionaryID(u)
	return dduo
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableDictionaryID(u *uint64) *DictionaryDetailUpdateOne {
	if u != nil {
		dduo.SetDictionaryID(*u)
	}
	return dduo
}

// ClearDictionaryID clears the value of the "dictionary_id" field.
func (dduo *DictionaryDetailUpdateOne) ClearDictionaryID() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDictionaryID()
	return dduo
}

// SetDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID.
func (dduo *DictionaryDetailUpdateOne) SetDictionariesID(id uint64) *DictionaryDetailUpdateOne {
	dduo.mutation.SetDictionariesID(id)
	return dduo
}

// SetNillableDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableDictionariesID(id *uint64) *DictionaryDetailUpdateOne {
	if id != nil {
		dduo = dduo.SetDictionariesID(*id)
	}
	return dduo
}

// SetDictionaries sets the "dictionaries" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) SetDictionaries(d *Dictionary) *DictionaryDetailUpdateOne {
	return dduo.SetDictionariesID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (dduo *DictionaryDetailUpdateOne) Mutation() *DictionaryDetailMutation {
	return dduo.mutation
}

// ClearDictionaries clears the "dictionaries" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) ClearDictionaries() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDictionaries()
	return dduo
}

// Where appends a list predicates to the DictionaryDetailUpdate builder.
func (dduo *DictionaryDetailUpdateOne) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailUpdateOne {
	dduo.mutation.Where(ps...)
	return dduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DictionaryDetailUpdateOne) Select(field string, fields ...string) *DictionaryDetailUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DictionaryDetail entity.
func (dduo *DictionaryDetailUpdateOne) Save(ctx context.Context) (*DictionaryDetail, error) {
	dduo.defaults()
	return withHooks(ctx, dduo.sqlSave, dduo.mutation, dduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) SaveX(ctx context.Context) *DictionaryDetail {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DictionaryDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dduo *DictionaryDetailUpdateOne) defaults() {
	if _, ok := dduo.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		dduo.mutation.SetUpdatedAt(v)
	}
}

func (dduo *DictionaryDetailUpdateOne) sqlSave(ctx context.Context) (_node *DictionaryDetail, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionarydetail.Table, dictionarydetail.Columns, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64))
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DictionaryDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionarydetail.FieldID)
		for _, f := range fields {
			if !dictionarydetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dictionarydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dduo.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dduo.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := dduo.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if dduo.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeUint8)
	}
	if value, ok := dduo.mutation.Sort(); ok {
		_spec.SetField(dictionarydetail.FieldSort, field.TypeUint32, value)
	}
	if value, ok := dduo.mutation.AddedSort(); ok {
		_spec.AddField(dictionarydetail.FieldSort, field.TypeUint32, value)
	}
	if value, ok := dduo.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if dduo.mutation.DictionariesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionariesTable,
			Columns: []string{dictionarydetail.DictionariesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dduo.mutation.DictionariesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionariesTable,
			Columns: []string{dictionarydetail.DictionariesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DictionaryDetail{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dduo.mutation.done = true
	return _node, nil
}
