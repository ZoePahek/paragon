// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
)

// TargetQuery is the builder for querying Target entities.
type TargetQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Target
	// eager-loading edges.
	withTasks       *TaskQuery
	withTags        *TagQuery
	withCredentials *CredentialQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (tq *TargetQuery) Where(ps ...predicate.Target) *TargetQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TargetQuery) Limit(limit int) *TargetQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TargetQuery) Offset(offset int) *TargetQuery {
	tq.offset = &offset
	return tq
}

// Order adds an order step to the query.
func (tq *TargetQuery) Order(o ...Order) *TargetQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryTasks chains the current query on the tasks edge.
func (tq *TargetQuery) QueryTasks() *TaskQuery {
	query := &TaskQuery{config: tq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(target.Table, target.FieldID, tq.sqlQuery()),
		sqlgraph.To(task.Table, task.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, target.TasksTable, target.TasksColumn),
	)
	query.sql = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
	return query
}

// QueryTags chains the current query on the tags edge.
func (tq *TargetQuery) QueryTags() *TagQuery {
	query := &TagQuery{config: tq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(target.Table, target.FieldID, tq.sqlQuery()),
		sqlgraph.To(tag.Table, tag.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, target.TagsTable, target.TagsPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
	return query
}

// QueryCredentials chains the current query on the credentials edge.
func (tq *TargetQuery) QueryCredentials() *CredentialQuery {
	query := &CredentialQuery{config: tq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(target.Table, target.FieldID, tq.sqlQuery()),
		sqlgraph.To(credential.Table, credential.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, target.CredentialsTable, target.CredentialsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
	return query
}

// First returns the first Target entity in the query. Returns *ErrNotFound when no target was found.
func (tq *TargetQuery) First(ctx context.Context) (*Target, error) {
	ts, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, &ErrNotFound{target.Label}
	}
	return ts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TargetQuery) FirstX(ctx context.Context) *Target {
	t, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return t
}

// FirstID returns the first Target id in the query. Returns *ErrNotFound when no id was found.
func (tq *TargetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{target.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (tq *TargetQuery) FirstXID(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Target entity in the query, returns an error if not exactly one entity was returned.
func (tq *TargetQuery) Only(ctx context.Context) (*Target, error) {
	ts, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ts) {
	case 1:
		return ts[0], nil
	case 0:
		return nil, &ErrNotFound{target.Label}
	default:
		return nil, &ErrNotSingular{target.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TargetQuery) OnlyX(ctx context.Context) *Target {
	t, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// OnlyID returns the only Target id in the query, returns an error if not exactly one id was returned.
func (tq *TargetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{target.Label}
	default:
		err = &ErrNotSingular{target.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (tq *TargetQuery) OnlyXID(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Targets.
func (tq *TargetQuery) All(ctx context.Context) ([]*Target, error) {
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TargetQuery) AllX(ctx context.Context) []*Target {
	ts, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ts
}

// IDs executes the query and returns a list of Target ids.
func (tq *TargetQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(target.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TargetQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TargetQuery) Count(ctx context.Context) (int, error) {
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TargetQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TargetQuery) Exist(ctx context.Context) (bool, error) {
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TargetQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TargetQuery) Clone() *TargetQuery {
	return &TargetQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]Order{}, tq.order...),
		unique:     append([]string{}, tq.unique...),
		predicates: append([]predicate.Target{}, tq.predicates...),
		// clone intermediate query.
		sql: tq.sql.Clone(),
	}
}

//  WithTasks tells the query-builder to eager-loads the nodes that are connected to
// the "tasks" edge. The optional arguments used to configure the query builder of the edge.
func (tq *TargetQuery) WithTasks(opts ...func(*TaskQuery)) *TargetQuery {
	query := &TaskQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTasks = query
	return tq
}

//  WithTags tells the query-builder to eager-loads the nodes that are connected to
// the "tags" edge. The optional arguments used to configure the query builder of the edge.
func (tq *TargetQuery) WithTags(opts ...func(*TagQuery)) *TargetQuery {
	query := &TagQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withTags = query
	return tq
}

//  WithCredentials tells the query-builder to eager-loads the nodes that are connected to
// the "credentials" edge. The optional arguments used to configure the query builder of the edge.
func (tq *TargetQuery) WithCredentials(opts ...func(*CredentialQuery)) *TargetQuery {
	query := &CredentialQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withCredentials = query
	return tq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Target.Query().
//		GroupBy(target.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TargetQuery) GroupBy(field string, fields ...string) *TargetGroupBy {
	group := &TargetGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = tq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.Target.Query().
//		Select(target.FieldName).
//		Scan(ctx, &v)
//
func (tq *TargetQuery) Select(field string, fields ...string) *TargetSelect {
	selector := &TargetSelect{config: tq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = tq.sqlQuery()
	return selector
}

func (tq *TargetQuery) sqlAll(ctx context.Context) ([]*Target, error) {
	var (
		nodes []*Target
		_spec = tq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Target{config: tq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withTasks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Target)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Task(func(s *sql.Selector) {
			s.Where(sql.InValues(target.TasksColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.target_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "target_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "target_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Tasks = append(node.Edges.Tasks, n)
		}
	}

	if query := tq.withTags; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Target, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Target)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   target.TagsTable,
				Columns: target.TagsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(target.TagsPrimaryKey[0], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(eout.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, tq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "tags": %v`, err)
		}
		query.Where(tag.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tags = append(nodes[i].Edges.Tags, n)
			}
		}
	}

	if query := tq.withCredentials; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Target)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Credential(func(s *sql.Selector) {
			s.Where(sql.InValues(target.CredentialsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.target_credential_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "target_credential_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "target_credential_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Credentials = append(node.Edges.Credentials, n)
		}
	}

	return nodes, nil
}

func (tq *TargetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TargetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (tq *TargetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TargetQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(target.Table)
	selector := builder.Select(t1.Columns(target.Columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(target.Columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TargetGroupBy is the builder for group-by Target entities.
type TargetGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TargetGroupBy) Aggregate(fns ...Aggregate) *TargetGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scan the result into the given value.
func (tgb *TargetGroupBy) Scan(ctx context.Context, v interface{}) error {
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TargetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (tgb *TargetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TargetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TargetGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (tgb *TargetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TargetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TargetGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (tgb *TargetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TargetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TargetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (tgb *TargetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TargetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TargetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TargetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tgb.sqlQuery().Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TargetGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql
	columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
	columns = append(columns, tgb.fields...)
	for _, fn := range tgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tgb.fields...)
}

// TargetSelect is the builder for select fields of Target entities.
type TargetSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ts *TargetSelect) Scan(ctx context.Context, v interface{}) error {
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TargetSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ts *TargetSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TargetSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TargetSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ts *TargetSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TargetSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TargetSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ts *TargetSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TargetSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TargetSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ts *TargetSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TargetSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TargetSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TargetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sqlQuery().Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ts *TargetSelect) sqlQuery() sql.Querier {
	selector := ts.sql
	selector.Select(selector.Columns(ts.fields...)...)
	return selector
}
