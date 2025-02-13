// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package a3boil

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TNRoute is an object representing the database table.
type TNRoute struct {
	ID    int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTN  null.Int    `boil:"id_tn" json:"id_tn,omitempty" toml:"id_tn" yaml:"id_tn,omitempty"`
	Route null.String `boil:"route" json:"route,omitempty" toml:"route" yaml:"route,omitempty"`

	R *tnRouteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tnRouteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TNRouteColumns = struct {
	ID    string
	IDTN  string
	Route string
}{
	ID:    "id",
	IDTN:  "id_tn",
	Route: "route",
}

var TNRouteTableColumns = struct {
	ID    string
	IDTN  string
	Route string
}{
	ID:    "tn_routes.id",
	IDTN:  "tn_routes.id_tn",
	Route: "tn_routes.route",
}

// Generated where

var TNRouteWhere = struct {
	ID    whereHelperint
	IDTN  whereHelpernull_Int
	Route whereHelpernull_String
}{
	ID:    whereHelperint{field: "[dbo].[tn_routes].[id]"},
	IDTN:  whereHelpernull_Int{field: "[dbo].[tn_routes].[id_tn]"},
	Route: whereHelpernull_String{field: "[dbo].[tn_routes].[route]"},
}

// TNRouteRels is where relationship names are stored.
var TNRouteRels = struct {
}{}

// tnRouteR is where relationships are stored.
type tnRouteR struct {
}

// NewStruct creates a new relationship struct
func (*tnRouteR) NewStruct() *tnRouteR {
	return &tnRouteR{}
}

// tnRouteL is where Load methods for each relationship are stored.
type tnRouteL struct{}

var (
	tnRouteAllColumns            = []string{"id", "id_tn", "route"}
	tnRouteColumnsWithoutDefault = []string{"id_tn", "route"}
	tnRouteColumnsWithDefault    = []string{"id"}
	tnRoutePrimaryKeyColumns     = []string{"id"}
	tnRouteGeneratedColumns      = []string{"id"}
)

type (
	// TNRouteSlice is an alias for a slice of pointers to TNRoute.
	// This should almost always be used instead of []TNRoute.
	TNRouteSlice []*TNRoute
	// TNRouteHook is the signature for custom TNRoute hook methods
	TNRouteHook func(context.Context, boil.ContextExecutor, *TNRoute) error

	tnRouteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tnRouteType                 = reflect.TypeOf(&TNRoute{})
	tnRouteMapping              = queries.MakeStructMapping(tnRouteType)
	tnRoutePrimaryKeyMapping, _ = queries.BindMapping(tnRouteType, tnRouteMapping, tnRoutePrimaryKeyColumns)
	tnRouteInsertCacheMut       sync.RWMutex
	tnRouteInsertCache          = make(map[string]insertCache)
	tnRouteUpdateCacheMut       sync.RWMutex
	tnRouteUpdateCache          = make(map[string]updateCache)
	tnRouteUpsertCacheMut       sync.RWMutex
	tnRouteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tnRouteAfterSelectMu sync.Mutex
var tnRouteAfterSelectHooks []TNRouteHook

var tnRouteBeforeInsertMu sync.Mutex
var tnRouteBeforeInsertHooks []TNRouteHook
var tnRouteAfterInsertMu sync.Mutex
var tnRouteAfterInsertHooks []TNRouteHook

var tnRouteBeforeUpdateMu sync.Mutex
var tnRouteBeforeUpdateHooks []TNRouteHook
var tnRouteAfterUpdateMu sync.Mutex
var tnRouteAfterUpdateHooks []TNRouteHook

var tnRouteBeforeDeleteMu sync.Mutex
var tnRouteBeforeDeleteHooks []TNRouteHook
var tnRouteAfterDeleteMu sync.Mutex
var tnRouteAfterDeleteHooks []TNRouteHook

var tnRouteBeforeUpsertMu sync.Mutex
var tnRouteBeforeUpsertHooks []TNRouteHook
var tnRouteAfterUpsertMu sync.Mutex
var tnRouteAfterUpsertHooks []TNRouteHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TNRoute) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TNRoute) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TNRoute) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TNRoute) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TNRoute) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TNRoute) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TNRoute) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TNRoute) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TNRoute) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRouteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTNRouteHook registers your hook function for all future operations.
func AddTNRouteHook(hookPoint boil.HookPoint, tnRouteHook TNRouteHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tnRouteAfterSelectMu.Lock()
		tnRouteAfterSelectHooks = append(tnRouteAfterSelectHooks, tnRouteHook)
		tnRouteAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tnRouteBeforeInsertMu.Lock()
		tnRouteBeforeInsertHooks = append(tnRouteBeforeInsertHooks, tnRouteHook)
		tnRouteBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tnRouteAfterInsertMu.Lock()
		tnRouteAfterInsertHooks = append(tnRouteAfterInsertHooks, tnRouteHook)
		tnRouteAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tnRouteBeforeUpdateMu.Lock()
		tnRouteBeforeUpdateHooks = append(tnRouteBeforeUpdateHooks, tnRouteHook)
		tnRouteBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tnRouteAfterUpdateMu.Lock()
		tnRouteAfterUpdateHooks = append(tnRouteAfterUpdateHooks, tnRouteHook)
		tnRouteAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tnRouteBeforeDeleteMu.Lock()
		tnRouteBeforeDeleteHooks = append(tnRouteBeforeDeleteHooks, tnRouteHook)
		tnRouteBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tnRouteAfterDeleteMu.Lock()
		tnRouteAfterDeleteHooks = append(tnRouteAfterDeleteHooks, tnRouteHook)
		tnRouteAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tnRouteBeforeUpsertMu.Lock()
		tnRouteBeforeUpsertHooks = append(tnRouteBeforeUpsertHooks, tnRouteHook)
		tnRouteBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tnRouteAfterUpsertMu.Lock()
		tnRouteAfterUpsertHooks = append(tnRouteAfterUpsertHooks, tnRouteHook)
		tnRouteAfterUpsertMu.Unlock()
	}
}

// OneG returns a single tnRoute record from the query using the global executor.
func (q tnRouteQuery) OneG(ctx context.Context) (*TNRoute, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single tnRoute record from the query.
func (q tnRouteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TNRoute, error) {
	o := &TNRoute{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for tn_routes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TNRoute records from the query using the global executor.
func (q tnRouteQuery) AllG(ctx context.Context) (TNRouteSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TNRoute records from the query.
func (q tnRouteQuery) All(ctx context.Context, exec boil.ContextExecutor) (TNRouteSlice, error) {
	var o []*TNRoute

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TNRoute slice")
	}

	if len(tnRouteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TNRoute records in the query using the global executor
func (q tnRouteQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TNRoute records in the query.
func (q tnRouteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count tn_routes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q tnRouteQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q tnRouteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if tn_routes exists")
	}

	return count > 0, nil
}

// TNRoutes retrieves all the records using an executor.
func TNRoutes(mods ...qm.QueryMod) tnRouteQuery {
	mods = append(mods, qm.From("[dbo].[tn_routes]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[tn_routes].*"})
	}

	return tnRouteQuery{q}
}

// FindTNRouteG retrieves a single record by ID.
func FindTNRouteG(ctx context.Context, iD int, selectCols ...string) (*TNRoute, error) {
	return FindTNRoute(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTNRoute retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTNRoute(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TNRoute, error) {
	tnRouteObj := &TNRoute{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[tn_routes] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tnRouteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from tn_routes")
	}

	if err = tnRouteObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tnRouteObj, err
	}

	return tnRouteObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TNRoute) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TNRoute) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn_routes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnRouteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tnRouteInsertCacheMut.RLock()
	cache, cached := tnRouteInsertCache[key]
	tnRouteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tnRouteAllColumns,
			tnRouteColumnsWithDefault,
			tnRouteColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, tnRouteGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(tnRouteType, tnRouteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tnRouteType, tnRouteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[tn_routes] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[tn_routes] %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryOutput = fmt.Sprintf("OUTPUT INSERTED.[%s] ", strings.Join(returnColumns, "],INSERTED.["))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "a3boil: unable to insert into tn_routes")
	}

	if !cached {
		tnRouteInsertCacheMut.Lock()
		tnRouteInsertCache[key] = cache
		tnRouteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TNRoute record using the global executor.
// See Update for more documentation.
func (o *TNRoute) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TNRoute.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TNRoute) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tnRouteUpdateCacheMut.RLock()
	cache, cached := tnRouteUpdateCache[key]
	tnRouteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tnRouteAllColumns,
			tnRoutePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tnRouteGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update tn_routes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[tn_routes] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, tnRoutePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tnRouteType, tnRouteMapping, append(wl, tnRoutePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update tn_routes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for tn_routes")
	}

	if !cached {
		tnRouteUpdateCacheMut.Lock()
		tnRouteUpdateCache[key] = cache
		tnRouteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tnRouteQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tnRouteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for tn_routes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for tn_routes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TNRouteSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TNRouteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("a3boil: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[tn_routes] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tnRoutePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in tnRoute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all tnRoute")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TNRoute) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TNRoute) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn_routes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnRouteColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tnRouteUpsertCacheMut.RLock()
	cache, cached := tnRouteUpsertCache[key]
	tnRouteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tnRouteAllColumns,
			tnRouteColumnsWithDefault,
			tnRouteColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, tnRouteGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(tnRoutePrimaryKeyColumns, v) && strmangle.ContainsAny(tnRouteColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert tn_routes, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			tnRouteAllColumns,
			tnRoutePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, tnRouteGeneratedColumns)

		ret := strmangle.SetComplement(tnRouteAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert tn_routes, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[tn_routes]", tnRoutePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(tnRoutePrimaryKeyColumns))
		copy(whitelist, tnRoutePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(tnRouteType, tnRouteMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tnRouteType, tnRouteMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // MSSQL doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to upsert tn_routes")
	}

	if !cached {
		tnRouteUpsertCacheMut.Lock()
		tnRouteUpsertCache[key] = cache
		tnRouteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TNRoute record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TNRoute) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TNRoute record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TNRoute) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TNRoute provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tnRoutePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[tn_routes] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from tn_routes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for tn_routes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q tnRouteQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q tnRouteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no tnRouteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tn_routes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn_routes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TNRouteSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TNRouteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tnRouteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[tn_routes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tnRoutePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tnRoute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn_routes")
	}

	if len(tnRouteAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TNRoute) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TNRoute provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TNRoute) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTNRoute(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNRouteSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TNRouteSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNRouteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TNRouteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[tn_routes].* FROM [dbo].[tn_routes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tnRoutePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TNRouteSlice")
	}

	*o = slice

	return nil
}

// TNRouteExistsG checks if the TNRoute row exists.
func TNRouteExistsG(ctx context.Context, iD int) (bool, error) {
	return TNRouteExists(ctx, boil.GetContextDB(), iD)
}

// TNRouteExists checks if the TNRoute row exists.
func TNRouteExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[tn_routes] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if tn_routes exists")
	}

	return exists, nil
}

// Exists checks if the TNRoute row exists.
func (o *TNRoute) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TNRouteExists(ctx, exec, o.ID)
}
