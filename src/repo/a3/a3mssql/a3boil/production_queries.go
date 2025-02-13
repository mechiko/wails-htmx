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

// ProductionQuery is an object representing the database table.
type ProductionQuery struct {
	ID                  int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDProductionReports null.Int    `boil:"id_production_reports" json:"id_production_reports,omitempty" toml:"id_production_reports" yaml:"id_production_reports,omitempty"`
	QueryDate           null.String `boil:"query_date" json:"query_date,omitempty" toml:"query_date" yaml:"query_date,omitempty"`
	QueryRegID          null.String `boil:"query_reg_id" json:"query_reg_id,omitempty" toml:"query_reg_id" yaml:"query_reg_id,omitempty"`
	Status              null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID             null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`

	R *productionQueryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productionQueryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductionQueryColumns = struct {
	ID                  string
	IDProductionReports string
	QueryDate           string
	QueryRegID          string
	Status              string
	ReplyID             string
}{
	ID:                  "id",
	IDProductionReports: "id_production_reports",
	QueryDate:           "query_date",
	QueryRegID:          "query_reg_id",
	Status:              "status",
	ReplyID:             "reply_id",
}

var ProductionQueryTableColumns = struct {
	ID                  string
	IDProductionReports string
	QueryDate           string
	QueryRegID          string
	Status              string
	ReplyID             string
}{
	ID:                  "production_queries.id",
	IDProductionReports: "production_queries.id_production_reports",
	QueryDate:           "production_queries.query_date",
	QueryRegID:          "production_queries.query_reg_id",
	Status:              "production_queries.status",
	ReplyID:             "production_queries.reply_id",
}

// Generated where

var ProductionQueryWhere = struct {
	ID                  whereHelperint
	IDProductionReports whereHelpernull_Int
	QueryDate           whereHelpernull_String
	QueryRegID          whereHelpernull_String
	Status              whereHelpernull_String
	ReplyID             whereHelpernull_String
}{
	ID:                  whereHelperint{field: "[dbo].[production_queries].[id]"},
	IDProductionReports: whereHelpernull_Int{field: "[dbo].[production_queries].[id_production_reports]"},
	QueryDate:           whereHelpernull_String{field: "[dbo].[production_queries].[query_date]"},
	QueryRegID:          whereHelpernull_String{field: "[dbo].[production_queries].[query_reg_id]"},
	Status:              whereHelpernull_String{field: "[dbo].[production_queries].[status]"},
	ReplyID:             whereHelpernull_String{field: "[dbo].[production_queries].[reply_id]"},
}

// ProductionQueryRels is where relationship names are stored.
var ProductionQueryRels = struct {
}{}

// productionQueryR is where relationships are stored.
type productionQueryR struct {
}

// NewStruct creates a new relationship struct
func (*productionQueryR) NewStruct() *productionQueryR {
	return &productionQueryR{}
}

// productionQueryL is where Load methods for each relationship are stored.
type productionQueryL struct{}

var (
	productionQueryAllColumns            = []string{"id", "id_production_reports", "query_date", "query_reg_id", "status", "reply_id"}
	productionQueryColumnsWithoutDefault = []string{"id_production_reports", "query_date", "query_reg_id", "status", "reply_id"}
	productionQueryColumnsWithDefault    = []string{"id"}
	productionQueryPrimaryKeyColumns     = []string{"id"}
	productionQueryGeneratedColumns      = []string{"id"}
)

type (
	// ProductionQuerySlice is an alias for a slice of pointers to ProductionQuery.
	// This should almost always be used instead of []ProductionQuery.
	ProductionQuerySlice []*ProductionQuery
	// ProductionQueryHook is the signature for custom ProductionQuery hook methods
	ProductionQueryHook func(context.Context, boil.ContextExecutor, *ProductionQuery) error

	productionQueryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productionQueryType                 = reflect.TypeOf(&ProductionQuery{})
	productionQueryMapping              = queries.MakeStructMapping(productionQueryType)
	productionQueryPrimaryKeyMapping, _ = queries.BindMapping(productionQueryType, productionQueryMapping, productionQueryPrimaryKeyColumns)
	productionQueryInsertCacheMut       sync.RWMutex
	productionQueryInsertCache          = make(map[string]insertCache)
	productionQueryUpdateCacheMut       sync.RWMutex
	productionQueryUpdateCache          = make(map[string]updateCache)
	productionQueryUpsertCacheMut       sync.RWMutex
	productionQueryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productionQueryAfterSelectMu sync.Mutex
var productionQueryAfterSelectHooks []ProductionQueryHook

var productionQueryBeforeInsertMu sync.Mutex
var productionQueryBeforeInsertHooks []ProductionQueryHook
var productionQueryAfterInsertMu sync.Mutex
var productionQueryAfterInsertHooks []ProductionQueryHook

var productionQueryBeforeUpdateMu sync.Mutex
var productionQueryBeforeUpdateHooks []ProductionQueryHook
var productionQueryAfterUpdateMu sync.Mutex
var productionQueryAfterUpdateHooks []ProductionQueryHook

var productionQueryBeforeDeleteMu sync.Mutex
var productionQueryBeforeDeleteHooks []ProductionQueryHook
var productionQueryAfterDeleteMu sync.Mutex
var productionQueryAfterDeleteHooks []ProductionQueryHook

var productionQueryBeforeUpsertMu sync.Mutex
var productionQueryBeforeUpsertHooks []ProductionQueryHook
var productionQueryAfterUpsertMu sync.Mutex
var productionQueryAfterUpsertHooks []ProductionQueryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductionQuery) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductionQuery) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductionQuery) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductionQuery) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductionQuery) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductionQuery) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductionQuery) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductionQuery) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductionQuery) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionQueryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductionQueryHook registers your hook function for all future operations.
func AddProductionQueryHook(hookPoint boil.HookPoint, productionQueryHook ProductionQueryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productionQueryAfterSelectMu.Lock()
		productionQueryAfterSelectHooks = append(productionQueryAfterSelectHooks, productionQueryHook)
		productionQueryAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productionQueryBeforeInsertMu.Lock()
		productionQueryBeforeInsertHooks = append(productionQueryBeforeInsertHooks, productionQueryHook)
		productionQueryBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productionQueryAfterInsertMu.Lock()
		productionQueryAfterInsertHooks = append(productionQueryAfterInsertHooks, productionQueryHook)
		productionQueryAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productionQueryBeforeUpdateMu.Lock()
		productionQueryBeforeUpdateHooks = append(productionQueryBeforeUpdateHooks, productionQueryHook)
		productionQueryBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productionQueryAfterUpdateMu.Lock()
		productionQueryAfterUpdateHooks = append(productionQueryAfterUpdateHooks, productionQueryHook)
		productionQueryAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productionQueryBeforeDeleteMu.Lock()
		productionQueryBeforeDeleteHooks = append(productionQueryBeforeDeleteHooks, productionQueryHook)
		productionQueryBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productionQueryAfterDeleteMu.Lock()
		productionQueryAfterDeleteHooks = append(productionQueryAfterDeleteHooks, productionQueryHook)
		productionQueryAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productionQueryBeforeUpsertMu.Lock()
		productionQueryBeforeUpsertHooks = append(productionQueryBeforeUpsertHooks, productionQueryHook)
		productionQueryBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productionQueryAfterUpsertMu.Lock()
		productionQueryAfterUpsertHooks = append(productionQueryAfterUpsertHooks, productionQueryHook)
		productionQueryAfterUpsertMu.Unlock()
	}
}

// OneG returns a single productionQuery record from the query using the global executor.
func (q productionQueryQuery) OneG(ctx context.Context) (*ProductionQuery, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productionQuery record from the query.
func (q productionQueryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductionQuery, error) {
	o := &ProductionQuery{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for production_queries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductionQuery records from the query using the global executor.
func (q productionQueryQuery) AllG(ctx context.Context) (ProductionQuerySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductionQuery records from the query.
func (q productionQueryQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductionQuerySlice, error) {
	var o []*ProductionQuery

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ProductionQuery slice")
	}

	if len(productionQueryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductionQuery records in the query using the global executor
func (q productionQueryQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductionQuery records in the query.
func (q productionQueryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count production_queries rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productionQueryQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productionQueryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if production_queries exists")
	}

	return count > 0, nil
}

// ProductionQueries retrieves all the records using an executor.
func ProductionQueries(mods ...qm.QueryMod) productionQueryQuery {
	mods = append(mods, qm.From("[dbo].[production_queries]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[production_queries].*"})
	}

	return productionQueryQuery{q}
}

// FindProductionQueryG retrieves a single record by ID.
func FindProductionQueryG(ctx context.Context, iD int, selectCols ...string) (*ProductionQuery, error) {
	return FindProductionQuery(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductionQuery retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductionQuery(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProductionQuery, error) {
	productionQueryObj := &ProductionQuery{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[production_queries] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productionQueryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from production_queries")
	}

	if err = productionQueryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productionQueryObj, err
	}

	return productionQueryObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductionQuery) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductionQuery) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_queries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionQueryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productionQueryInsertCacheMut.RLock()
	cache, cached := productionQueryInsertCache[key]
	productionQueryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productionQueryAllColumns,
			productionQueryColumnsWithDefault,
			productionQueryColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productionQueryGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productionQueryType, productionQueryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productionQueryType, productionQueryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[production_queries] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[production_queries] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into production_queries")
	}

	if !cached {
		productionQueryInsertCacheMut.Lock()
		productionQueryInsertCache[key] = cache
		productionQueryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductionQuery record using the global executor.
// See Update for more documentation.
func (o *ProductionQuery) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductionQuery.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductionQuery) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productionQueryUpdateCacheMut.RLock()
	cache, cached := productionQueryUpdateCache[key]
	productionQueryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productionQueryAllColumns,
			productionQueryPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productionQueryGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update production_queries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[production_queries] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, productionQueryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productionQueryType, productionQueryMapping, append(wl, productionQueryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update production_queries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for production_queries")
	}

	if !cached {
		productionQueryUpdateCacheMut.Lock()
		productionQueryUpdateCache[key] = cache
		productionQueryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productionQueryQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productionQueryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for production_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for production_queries")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductionQuerySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductionQuerySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[production_queries] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productionQueryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in productionQuery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all productionQuery")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductionQuery) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductionQuery) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_queries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionQueryColumnsWithDefault, o)

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

	productionQueryUpsertCacheMut.RLock()
	cache, cached := productionQueryUpsertCache[key]
	productionQueryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productionQueryAllColumns,
			productionQueryColumnsWithDefault,
			productionQueryColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, productionQueryGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(productionQueryPrimaryKeyColumns, v) && strmangle.ContainsAny(productionQueryColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert production_queries, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			productionQueryAllColumns,
			productionQueryPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, productionQueryGeneratedColumns)

		ret := strmangle.SetComplement(productionQueryAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert production_queries, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[production_queries]", productionQueryPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(productionQueryPrimaryKeyColumns))
		copy(whitelist, productionQueryPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(productionQueryType, productionQueryMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productionQueryType, productionQueryMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert production_queries")
	}

	if !cached {
		productionQueryUpsertCacheMut.Lock()
		productionQueryUpsertCache[key] = cache
		productionQueryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductionQuery record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductionQuery) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductionQuery record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductionQuery) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ProductionQuery provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productionQueryPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[production_queries] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from production_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for production_queries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productionQueryQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productionQueryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no productionQueryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from production_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_queries")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductionQuerySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductionQuerySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productionQueryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[production_queries] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionQueryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from productionQuery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_queries")
	}

	if len(productionQueryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductionQuery) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ProductionQuery provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductionQuery) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductionQuery(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionQuerySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ProductionQuerySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionQuerySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductionQuerySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[production_queries].* FROM [dbo].[production_queries] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionQueryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ProductionQuerySlice")
	}

	*o = slice

	return nil
}

// ProductionQueryExistsG checks if the ProductionQuery row exists.
func ProductionQueryExistsG(ctx context.Context, iD int) (bool, error) {
	return ProductionQueryExists(ctx, boil.GetContextDB(), iD)
}

// ProductionQueryExists checks if the ProductionQuery row exists.
func ProductionQueryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[production_queries] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if production_queries exists")
	}

	return exists, nil
}

// Exists checks if the ProductionQuery row exists.
func (o *ProductionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductionQueryExists(ctx, exec, o.ID)
}
