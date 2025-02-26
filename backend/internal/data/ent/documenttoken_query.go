// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/documenttoken"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// DocumentTokenQuery is the builder for querying DocumentToken entities.
type DocumentTokenQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.DocumentToken
	withDocument *DocumentQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DocumentTokenQuery builder.
func (dtq *DocumentTokenQuery) Where(ps ...predicate.DocumentToken) *DocumentTokenQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit adds a limit step to the query.
func (dtq *DocumentTokenQuery) Limit(limit int) *DocumentTokenQuery {
	dtq.limit = &limit
	return dtq
}

// Offset adds an offset step to the query.
func (dtq *DocumentTokenQuery) Offset(offset int) *DocumentTokenQuery {
	dtq.offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DocumentTokenQuery) Unique(unique bool) *DocumentTokenQuery {
	dtq.unique = &unique
	return dtq
}

// Order adds an order step to the query.
func (dtq *DocumentTokenQuery) Order(o ...OrderFunc) *DocumentTokenQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// QueryDocument chains the current query on the "document" edge.
func (dtq *DocumentTokenQuery) QueryDocument() *DocumentQuery {
	query := &DocumentQuery{config: dtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(documenttoken.Table, documenttoken.FieldID, selector),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, documenttoken.DocumentTable, documenttoken.DocumentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DocumentToken entity from the query.
// Returns a *NotFoundError when no DocumentToken was found.
func (dtq *DocumentTokenQuery) First(ctx context.Context) (*DocumentToken, error) {
	nodes, err := dtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{documenttoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DocumentTokenQuery) FirstX(ctx context.Context) *DocumentToken {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DocumentToken ID from the query.
// Returns a *NotFoundError when no DocumentToken ID was found.
func (dtq *DocumentTokenQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{documenttoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DocumentTokenQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DocumentToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DocumentToken entity is found.
// Returns a *NotFoundError when no DocumentToken entities are found.
func (dtq *DocumentTokenQuery) Only(ctx context.Context) (*DocumentToken, error) {
	nodes, err := dtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{documenttoken.Label}
	default:
		return nil, &NotSingularError{documenttoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DocumentTokenQuery) OnlyX(ctx context.Context) *DocumentToken {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DocumentToken ID in the query.
// Returns a *NotSingularError when more than one DocumentToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DocumentTokenQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{documenttoken.Label}
	default:
		err = &NotSingularError{documenttoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DocumentTokenQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DocumentTokens.
func (dtq *DocumentTokenQuery) All(ctx context.Context) ([]*DocumentToken, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DocumentTokenQuery) AllX(ctx context.Context) []*DocumentToken {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DocumentToken IDs.
func (dtq *DocumentTokenQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dtq.Select(documenttoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DocumentTokenQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DocumentTokenQuery) Count(ctx context.Context) (int, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DocumentTokenQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DocumentTokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := dtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DocumentTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DocumentTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DocumentTokenQuery) Clone() *DocumentTokenQuery {
	if dtq == nil {
		return nil
	}
	return &DocumentTokenQuery{
		config:       dtq.config,
		limit:        dtq.limit,
		offset:       dtq.offset,
		order:        append([]OrderFunc{}, dtq.order...),
		predicates:   append([]predicate.DocumentToken{}, dtq.predicates...),
		withDocument: dtq.withDocument.Clone(),
		// clone intermediate query.
		sql:    dtq.sql.Clone(),
		path:   dtq.path,
		unique: dtq.unique,
	}
}

// WithDocument tells the query-builder to eager-load the nodes that are connected to
// the "document" edge. The optional arguments are used to configure the query builder of the edge.
func (dtq *DocumentTokenQuery) WithDocument(opts ...func(*DocumentQuery)) *DocumentTokenQuery {
	query := &DocumentQuery{config: dtq.config}
	for _, opt := range opts {
		opt(query)
	}
	dtq.withDocument = query
	return dtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DocumentToken.Query().
//		GroupBy(documenttoken.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dtq *DocumentTokenQuery) GroupBy(field string, fields ...string) *DocumentTokenGroupBy {
	grbuild := &DocumentTokenGroupBy{config: dtq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dtq.sqlQuery(ctx), nil
	}
	grbuild.label = documenttoken.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DocumentToken.Query().
//		Select(documenttoken.FieldCreatedAt).
//		Scan(ctx, &v)
func (dtq *DocumentTokenQuery) Select(fields ...string) *DocumentTokenSelect {
	dtq.fields = append(dtq.fields, fields...)
	selbuild := &DocumentTokenSelect{DocumentTokenQuery: dtq}
	selbuild.label = documenttoken.Label
	selbuild.flds, selbuild.scan = &dtq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a DocumentTokenSelect configured with the given aggregations.
func (dtq *DocumentTokenQuery) Aggregate(fns ...AggregateFunc) *DocumentTokenSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DocumentTokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dtq.fields {
		if !documenttoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DocumentTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DocumentToken, error) {
	var (
		nodes       = []*DocumentToken{}
		withFKs     = dtq.withFKs
		_spec       = dtq.querySpec()
		loadedTypes = [1]bool{
			dtq.withDocument != nil,
		}
	)
	if dtq.withDocument != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, documenttoken.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DocumentToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DocumentToken{config: dtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dtq.withDocument; query != nil {
		if err := dtq.loadDocument(ctx, query, nodes, nil,
			func(n *DocumentToken, e *Document) { n.Edges.Document = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dtq *DocumentTokenQuery) loadDocument(ctx context.Context, query *DocumentQuery, nodes []*DocumentToken, init func(*DocumentToken), assign func(*DocumentToken, *Document)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DocumentToken)
	for i := range nodes {
		if nodes[i].document_document_tokens == nil {
			continue
		}
		fk := *nodes[i].document_document_tokens
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(document.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "document_document_tokens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dtq *DocumentTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	_spec.Node.Columns = dtq.fields
	if len(dtq.fields) > 0 {
		_spec.Unique = dtq.unique != nil && *dtq.unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DocumentTokenQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (dtq *DocumentTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   documenttoken.Table,
			Columns: documenttoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: documenttoken.FieldID,
			},
		},
		From:   dtq.sql,
		Unique: true,
	}
	if unique := dtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, documenttoken.FieldID)
		for i := range fields {
			if fields[i] != documenttoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DocumentTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(documenttoken.Table)
	columns := dtq.fields
	if len(columns) == 0 {
		columns = documenttoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.unique != nil && *dtq.unique {
		selector.Distinct()
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DocumentTokenGroupBy is the group-by builder for DocumentToken entities.
type DocumentTokenGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DocumentTokenGroupBy) Aggregate(fns ...AggregateFunc) *DocumentTokenGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dtgb *DocumentTokenGroupBy) Scan(ctx context.Context, v any) error {
	query, err := dtgb.path(ctx)
	if err != nil {
		return err
	}
	dtgb.sql = query
	return dtgb.sqlScan(ctx, v)
}

func (dtgb *DocumentTokenGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range dtgb.fields {
		if !documenttoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dtgb *DocumentTokenGroupBy) sqlQuery() *sql.Selector {
	selector := dtgb.sql.Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dtgb.fields)+len(dtgb.fns))
		for _, f := range dtgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dtgb.fields...)...)
}

// DocumentTokenSelect is the builder for selecting fields of DocumentToken entities.
type DocumentTokenSelect struct {
	*DocumentTokenQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DocumentTokenSelect) Aggregate(fns ...AggregateFunc) *DocumentTokenSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DocumentTokenSelect) Scan(ctx context.Context, v any) error {
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	dts.sql = dts.DocumentTokenQuery.sqlQuery(ctx)
	return dts.sqlScan(ctx, v)
}

func (dts *DocumentTokenSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(dts.sql))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		dts.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		dts.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := dts.sql.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
