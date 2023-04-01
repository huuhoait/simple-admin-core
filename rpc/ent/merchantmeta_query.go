// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchant"
	"github.com/huuhoait/zero-admin-core/rpc/ent/merchantmeta"
	"github.com/huuhoait/zero-admin-core/rpc/ent/predicate"
)

// MerchantMetaQuery is the builder for querying MerchantMeta entities.
type MerchantMetaQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.MerchantMeta
	withMerchant *MerchantQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MerchantMetaQuery builder.
func (mmq *MerchantMetaQuery) Where(ps ...predicate.MerchantMeta) *MerchantMetaQuery {
	mmq.predicates = append(mmq.predicates, ps...)
	return mmq
}

// Limit the number of records to be returned by this query.
func (mmq *MerchantMetaQuery) Limit(limit int) *MerchantMetaQuery {
	mmq.ctx.Limit = &limit
	return mmq
}

// Offset to start from.
func (mmq *MerchantMetaQuery) Offset(offset int) *MerchantMetaQuery {
	mmq.ctx.Offset = &offset
	return mmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mmq *MerchantMetaQuery) Unique(unique bool) *MerchantMetaQuery {
	mmq.ctx.Unique = &unique
	return mmq
}

// Order specifies how the records should be ordered.
func (mmq *MerchantMetaQuery) Order(o ...OrderFunc) *MerchantMetaQuery {
	mmq.order = append(mmq.order, o...)
	return mmq
}

// QueryMerchant chains the current query on the "merchant" edge.
func (mmq *MerchantMetaQuery) QueryMerchant() *MerchantQuery {
	query := (&MerchantClient{config: mmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantmeta.Table, merchantmeta.FieldID, selector),
			sqlgraph.To(merchant.Table, merchant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, merchantmeta.MerchantTable, merchantmeta.MerchantColumn),
		)
		fromU = sqlgraph.SetNeighbors(mmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MerchantMeta entity from the query.
// Returns a *NotFoundError when no MerchantMeta was found.
func (mmq *MerchantMetaQuery) First(ctx context.Context) (*MerchantMeta, error) {
	nodes, err := mmq.Limit(1).All(setContextOp(ctx, mmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{merchantmeta.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mmq *MerchantMetaQuery) FirstX(ctx context.Context) *MerchantMeta {
	node, err := mmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MerchantMeta ID from the query.
// Returns a *NotFoundError when no MerchantMeta ID was found.
func (mmq *MerchantMetaQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mmq.Limit(1).IDs(setContextOp(ctx, mmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{merchantmeta.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mmq *MerchantMetaQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := mmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MerchantMeta entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MerchantMeta entity is found.
// Returns a *NotFoundError when no MerchantMeta entities are found.
func (mmq *MerchantMetaQuery) Only(ctx context.Context) (*MerchantMeta, error) {
	nodes, err := mmq.Limit(2).All(setContextOp(ctx, mmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{merchantmeta.Label}
	default:
		return nil, &NotSingularError{merchantmeta.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mmq *MerchantMetaQuery) OnlyX(ctx context.Context) *MerchantMeta {
	node, err := mmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MerchantMeta ID in the query.
// Returns a *NotSingularError when more than one MerchantMeta ID is found.
// Returns a *NotFoundError when no entities are found.
func (mmq *MerchantMetaQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mmq.Limit(2).IDs(setContextOp(ctx, mmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{merchantmeta.Label}
	default:
		err = &NotSingularError{merchantmeta.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mmq *MerchantMetaQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := mmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MerchantMetaSlice.
func (mmq *MerchantMetaQuery) All(ctx context.Context) ([]*MerchantMeta, error) {
	ctx = setContextOp(ctx, mmq.ctx, "All")
	if err := mmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MerchantMeta, *MerchantMetaQuery]()
	return withInterceptors[[]*MerchantMeta](ctx, mmq, qr, mmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mmq *MerchantMetaQuery) AllX(ctx context.Context) []*MerchantMeta {
	nodes, err := mmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MerchantMeta IDs.
func (mmq *MerchantMetaQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if mmq.ctx.Unique == nil && mmq.path != nil {
		mmq.Unique(true)
	}
	ctx = setContextOp(ctx, mmq.ctx, "IDs")
	if err = mmq.Select(merchantmeta.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mmq *MerchantMetaQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := mmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mmq *MerchantMetaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mmq.ctx, "Count")
	if err := mmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mmq, querierCount[*MerchantMetaQuery](), mmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mmq *MerchantMetaQuery) CountX(ctx context.Context) int {
	count, err := mmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mmq *MerchantMetaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mmq.ctx, "Exist")
	switch _, err := mmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mmq *MerchantMetaQuery) ExistX(ctx context.Context) bool {
	exist, err := mmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MerchantMetaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mmq *MerchantMetaQuery) Clone() *MerchantMetaQuery {
	if mmq == nil {
		return nil
	}
	return &MerchantMetaQuery{
		config:       mmq.config,
		ctx:          mmq.ctx.Clone(),
		order:        append([]OrderFunc{}, mmq.order...),
		inters:       append([]Interceptor{}, mmq.inters...),
		predicates:   append([]predicate.MerchantMeta{}, mmq.predicates...),
		withMerchant: mmq.withMerchant.Clone(),
		// clone intermediate query.
		sql:  mmq.sql.Clone(),
		path: mmq.path,
	}
}

// WithMerchant tells the query-builder to eager-load the nodes that are connected to
// the "merchant" edge. The optional arguments are used to configure the query builder of the edge.
func (mmq *MerchantMetaQuery) WithMerchant(opts ...func(*MerchantQuery)) *MerchantMetaQuery {
	query := (&MerchantClient{config: mmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mmq.withMerchant = query
	return mmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy string `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MerchantMeta.Query().
//		GroupBy(merchantmeta.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mmq *MerchantMetaQuery) GroupBy(field string, fields ...string) *MerchantMetaGroupBy {
	mmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MerchantMetaGroupBy{build: mmq}
	grbuild.flds = &mmq.ctx.Fields
	grbuild.label = merchantmeta.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy string `json:"created_by,omitempty"`
//	}
//
//	client.MerchantMeta.Query().
//		Select(merchantmeta.FieldCreatedBy).
//		Scan(ctx, &v)
func (mmq *MerchantMetaQuery) Select(fields ...string) *MerchantMetaSelect {
	mmq.ctx.Fields = append(mmq.ctx.Fields, fields...)
	sbuild := &MerchantMetaSelect{MerchantMetaQuery: mmq}
	sbuild.label = merchantmeta.Label
	sbuild.flds, sbuild.scan = &mmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MerchantMetaSelect configured with the given aggregations.
func (mmq *MerchantMetaQuery) Aggregate(fns ...AggregateFunc) *MerchantMetaSelect {
	return mmq.Select().Aggregate(fns...)
}

func (mmq *MerchantMetaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mmq); err != nil {
				return err
			}
		}
	}
	for _, f := range mmq.ctx.Fields {
		if !merchantmeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mmq.path != nil {
		prev, err := mmq.path(ctx)
		if err != nil {
			return err
		}
		mmq.sql = prev
	}
	return nil
}

func (mmq *MerchantMetaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MerchantMeta, error) {
	var (
		nodes       = []*MerchantMeta{}
		_spec       = mmq.querySpec()
		loadedTypes = [1]bool{
			mmq.withMerchant != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MerchantMeta).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MerchantMeta{config: mmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mmq.withMerchant; query != nil {
		if err := mmq.loadMerchant(ctx, query, nodes, nil,
			func(n *MerchantMeta, e *Merchant) { n.Edges.Merchant = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mmq *MerchantMetaQuery) loadMerchant(ctx context.Context, query *MerchantQuery, nodes []*MerchantMeta, init func(*MerchantMeta), assign func(*MerchantMeta, *Merchant)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*MerchantMeta)
	for i := range nodes {
		fk := nodes[i].MerchantID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(merchant.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "merchant_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mmq *MerchantMetaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mmq.querySpec()
	_spec.Node.Columns = mmq.ctx.Fields
	if len(mmq.ctx.Fields) > 0 {
		_spec.Unique = mmq.ctx.Unique != nil && *mmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mmq.driver, _spec)
}

func (mmq *MerchantMetaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(merchantmeta.Table, merchantmeta.Columns, sqlgraph.NewFieldSpec(merchantmeta.FieldID, field.TypeUint64))
	_spec.From = mmq.sql
	if unique := mmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mmq.path != nil {
		_spec.Unique = true
	}
	if fields := mmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantmeta.FieldID)
		for i := range fields {
			if fields[i] != merchantmeta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mmq *MerchantMetaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mmq.driver.Dialect())
	t1 := builder.Table(merchantmeta.Table)
	columns := mmq.ctx.Fields
	if len(columns) == 0 {
		columns = merchantmeta.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mmq.sql != nil {
		selector = mmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mmq.ctx.Unique != nil && *mmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mmq.predicates {
		p(selector)
	}
	for _, p := range mmq.order {
		p(selector)
	}
	if offset := mmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MerchantMetaGroupBy is the group-by builder for MerchantMeta entities.
type MerchantMetaGroupBy struct {
	selector
	build *MerchantMetaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mmgb *MerchantMetaGroupBy) Aggregate(fns ...AggregateFunc) *MerchantMetaGroupBy {
	mmgb.fns = append(mmgb.fns, fns...)
	return mmgb
}

// Scan applies the selector query and scans the result into the given value.
func (mmgb *MerchantMetaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mmgb.build.ctx, "GroupBy")
	if err := mmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantMetaQuery, *MerchantMetaGroupBy](ctx, mmgb.build, mmgb, mmgb.build.inters, v)
}

func (mmgb *MerchantMetaGroupBy) sqlScan(ctx context.Context, root *MerchantMetaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mmgb.fns))
	for _, fn := range mmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mmgb.flds)+len(mmgb.fns))
		for _, f := range *mmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MerchantMetaSelect is the builder for selecting fields of MerchantMeta entities.
type MerchantMetaSelect struct {
	*MerchantMetaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mms *MerchantMetaSelect) Aggregate(fns ...AggregateFunc) *MerchantMetaSelect {
	mms.fns = append(mms.fns, fns...)
	return mms
}

// Scan applies the selector query and scans the result into the given value.
func (mms *MerchantMetaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mms.ctx, "Select")
	if err := mms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantMetaQuery, *MerchantMetaSelect](ctx, mms.MerchantMetaQuery, mms, mms.inters, v)
}

func (mms *MerchantMetaSelect) sqlScan(ctx context.Context, root *MerchantMetaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mms.fns))
	for _, fn := range mms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
