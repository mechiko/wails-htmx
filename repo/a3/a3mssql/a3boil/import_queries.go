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

// ImportQuery is an object representing the database table.
type ImportQuery struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDImportReports null.Int    `boil:"id_import_reports" json:"id_import_reports,omitempty" toml:"id_import_reports" yaml:"id_import_reports,omitempty"`
	QueryDate       null.String `boil:"query_date" json:"query_date,omitempty" toml:"query_date" yaml:"query_date,omitempty"`
	QueryRegID      null.String `boil:"query_reg_id" json:"query_reg_id,omitempty" toml:"query_reg_id" yaml:"query_reg_id,omitempty"`
	Status          null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID         null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`

	R *importQueryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L importQueryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ImportQueryColumns = struct {
	ID              string
	IDImportReports string
	QueryDate       string
	QueryRegID      string
	Status          string
	ReplyID         string
}{
	ID:              "id",
	IDImportReports: "id_import_reports",
	QueryDate:       "query_date",
	QueryRegID:      "query_reg_id",
	Status:          "status",
	ReplyID:         "reply_id",
}

var ImportQueryTableColumns = struct {
	ID              string
	IDImportReports string
	QueryDate       string
	QueryRegID      string
	Status          string
	ReplyID         string
}{
	ID:              "import_queries.id",
	IDImportReports: "import_queries.id_import_reports",
	QueryDate:       "import_queries.query_date",
	QueryRegID:      "import_queries.query_reg_id",
	Status:          "import_queries.status",
	ReplyID:         "import_queries.reply_id",
}

// Generated where

var ImportQueryWhere = struct {
	ID              whereHelperint
	IDImportReports whereHelpernull_Int
	QueryDate       whereHelpernull_String
	QueryRegID      whereHelpernull_String
	Status          whereHelpernull_String
	ReplyID         whereHelpernull_String
}{
	ID:              whereHelperint{field: "[dbo].[import_queries].[id]"},
	IDImportReports: whereHelpernull_Int{field: "[dbo].[import_queries].[id_import_reports]"},
	QueryDate:       whereHelpernull_String{field: "[dbo].[import_queries].[query_date]"},
	QueryRegID:      whereHelpernull_String{field: "[dbo].[import_queries].[query_reg_id]"},
	Status:          whereHelpernull_String{field: "[dbo].[import_queries].[status]"},
	ReplyID:         whereHelpernull_String{field: "[dbo].[import_queries].[reply_id]"},
}

// ImportQueryRels is where relationship names are stored.
var ImportQueryRels = struct {
}{}

// importQueryR is where relationships are stored.
type importQueryR struct {
}

// NewStruct creates a new relationship struct
func (*importQueryR) NewStruct() *importQueryR {
	return &importQueryR{}
}

// importQueryL is where Load methods for each relationship are stored.
type importQueryL struct{}

var (
	importQueryAllColumns            = []string{"id", "id_import_reports", "query_date", "query_reg_id", "status", "reply_id"}
	importQueryColumnsWithoutDefault = []string{"id_import_reports", "query_date", "query_reg_id", "status", "reply_id"}
	importQueryColumnsWithDefault    = []string{"id"}
	importQueryPrimaryKeyColumns     = []string{"id"}
	importQueryGeneratedColumns      = []string{"id"}
)

type (
	// ImportQuerySlice is an alias for a slice of pointers to ImportQuery.
	// This should almost always be used instead of []ImportQuery.
	ImportQuerySlice []*ImportQuery
	// ImportQueryHook is the signature for custom ImportQuery hook methods
	ImportQueryHook func(context.Context, boil.ContextExecutor, *ImportQuery) error

	importQueryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	importQueryType                 = reflect.TypeOf(&ImportQuery{})
	importQueryMapping              = queries.MakeStructMapping(importQueryType)
	importQueryPrimaryKeyMapping, _ = queries.BindMapping(importQueryType, importQueryMapping, importQueryPrimaryKeyColumns)
	importQueryInsertCacheMut       sync.RWMutex
	importQueryInsertCache          = make(map[string]insertCache)
	importQueryUpdateCacheMut       sync.RWMutex
	importQueryUpdateCache          = make(map[string]updateCache)
	importQueryUpsertCacheMut       sync.RWMutex
	importQueryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var importQueryAfterSelectMu sync.Mutex
var importQueryAfterSelectHooks []ImportQueryHook

var importQueryBeforeInsertMu sync.Mutex
var importQueryBeforeInsertHooks []ImportQueryHook
var importQueryAfterInsertMu sync.Mutex
var importQueryAfterInsertHooks []ImportQueryHook

var importQueryBeforeUpdateMu sync.Mutex
var importQueryBeforeUpdateHooks []ImportQueryHook
var importQueryAfterUpdateMu sync.Mutex
var importQueryAfterUpdateHooks []ImportQueryHook

var importQueryBeforeDeleteMu sync.Mutex
var importQueryBeforeDeleteHooks []ImportQueryHook
var importQueryAfterDeleteMu sync.Mutex
var importQueryAfterDeleteHooks []ImportQueryHook

var importQueryBeforeUpsertMu sync.Mutex
var importQueryBeforeUpsertHooks []ImportQueryHook
var importQueryAfterUpsertMu sync.Mutex
var importQueryAfterUpsertHooks []ImportQueryHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ImportQuery) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ImportQuery) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ImportQuery) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ImportQuery) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ImportQuery) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ImportQuery) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ImportQuery) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ImportQuery) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ImportQuery) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range importQueryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddImportQueryHook registers your hook function for all future operations.
func AddImportQueryHook(hookPoint boil.HookPoint, importQueryHook ImportQueryHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		importQueryAfterSelectMu.Lock()
		importQueryAfterSelectHooks = append(importQueryAfterSelectHooks, importQueryHook)
		importQueryAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		importQueryBeforeInsertMu.Lock()
		importQueryBeforeInsertHooks = append(importQueryBeforeInsertHooks, importQueryHook)
		importQueryBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		importQueryAfterInsertMu.Lock()
		importQueryAfterInsertHooks = append(importQueryAfterInsertHooks, importQueryHook)
		importQueryAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		importQueryBeforeUpdateMu.Lock()
		importQueryBeforeUpdateHooks = append(importQueryBeforeUpdateHooks, importQueryHook)
		importQueryBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		importQueryAfterUpdateMu.Lock()
		importQueryAfterUpdateHooks = append(importQueryAfterUpdateHooks, importQueryHook)
		importQueryAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		importQueryBeforeDeleteMu.Lock()
		importQueryBeforeDeleteHooks = append(importQueryBeforeDeleteHooks, importQueryHook)
		importQueryBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		importQueryAfterDeleteMu.Lock()
		importQueryAfterDeleteHooks = append(importQueryAfterDeleteHooks, importQueryHook)
		importQueryAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		importQueryBeforeUpsertMu.Lock()
		importQueryBeforeUpsertHooks = append(importQueryBeforeUpsertHooks, importQueryHook)
		importQueryBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		importQueryAfterUpsertMu.Lock()
		importQueryAfterUpsertHooks = append(importQueryAfterUpsertHooks, importQueryHook)
		importQueryAfterUpsertMu.Unlock()
	}
}

// OneG returns a single importQuery record from the query using the global executor.
func (q importQueryQuery) OneG(ctx context.Context) (*ImportQuery, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single importQuery record from the query.
func (q importQueryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ImportQuery, error) {
	o := &ImportQuery{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for import_queries")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ImportQuery records from the query using the global executor.
func (q importQueryQuery) AllG(ctx context.Context) (ImportQuerySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ImportQuery records from the query.
func (q importQueryQuery) All(ctx context.Context, exec boil.ContextExecutor) (ImportQuerySlice, error) {
	var o []*ImportQuery

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ImportQuery slice")
	}

	if len(importQueryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ImportQuery records in the query using the global executor
func (q importQueryQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ImportQuery records in the query.
func (q importQueryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count import_queries rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q importQueryQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q importQueryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if import_queries exists")
	}

	return count > 0, nil
}

// ImportQueries retrieves all the records using an executor.
func ImportQueries(mods ...qm.QueryMod) importQueryQuery {
	mods = append(mods, qm.From("[dbo].[import_queries]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[import_queries].*"})
	}

	return importQueryQuery{q}
}

// FindImportQueryG retrieves a single record by ID.
func FindImportQueryG(ctx context.Context, iD int, selectCols ...string) (*ImportQuery, error) {
	return FindImportQuery(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindImportQuery retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindImportQuery(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ImportQuery, error) {
	importQueryObj := &ImportQuery{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[import_queries] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, importQueryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from import_queries")
	}

	if err = importQueryObj.doAfterSelectHooks(ctx, exec); err != nil {
		return importQueryObj, err
	}

	return importQueryObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ImportQuery) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ImportQuery) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no import_queries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(importQueryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	importQueryInsertCacheMut.RLock()
	cache, cached := importQueryInsertCache[key]
	importQueryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			importQueryAllColumns,
			importQueryColumnsWithDefault,
			importQueryColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, importQueryGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(importQueryType, importQueryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(importQueryType, importQueryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[import_queries] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[import_queries] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into import_queries")
	}

	if !cached {
		importQueryInsertCacheMut.Lock()
		importQueryInsertCache[key] = cache
		importQueryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ImportQuery record using the global executor.
// See Update for more documentation.
func (o *ImportQuery) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ImportQuery.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ImportQuery) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	importQueryUpdateCacheMut.RLock()
	cache, cached := importQueryUpdateCache[key]
	importQueryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			importQueryAllColumns,
			importQueryPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, importQueryGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update import_queries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[import_queries] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, importQueryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(importQueryType, importQueryMapping, append(wl, importQueryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update import_queries row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for import_queries")
	}

	if !cached {
		importQueryUpdateCacheMut.Lock()
		importQueryUpdateCache[key] = cache
		importQueryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q importQueryQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q importQueryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for import_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for import_queries")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ImportQuerySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ImportQuerySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), importQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[import_queries] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, importQueryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in importQuery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all importQuery")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ImportQuery) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ImportQuery) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no import_queries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(importQueryColumnsWithDefault, o)

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

	importQueryUpsertCacheMut.RLock()
	cache, cached := importQueryUpsertCache[key]
	importQueryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			importQueryAllColumns,
			importQueryColumnsWithDefault,
			importQueryColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, importQueryGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(importQueryPrimaryKeyColumns, v) && strmangle.ContainsAny(importQueryColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert import_queries, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			importQueryAllColumns,
			importQueryPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, importQueryGeneratedColumns)

		ret := strmangle.SetComplement(importQueryAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert import_queries, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[import_queries]", importQueryPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(importQueryPrimaryKeyColumns))
		copy(whitelist, importQueryPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(importQueryType, importQueryMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(importQueryType, importQueryMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert import_queries")
	}

	if !cached {
		importQueryUpsertCacheMut.Lock()
		importQueryUpsertCache[key] = cache
		importQueryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ImportQuery record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ImportQuery) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ImportQuery record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ImportQuery) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ImportQuery provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), importQueryPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[import_queries] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from import_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for import_queries")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q importQueryQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q importQueryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no importQueryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from import_queries")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for import_queries")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ImportQuerySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ImportQuerySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(importQueryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), importQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[import_queries] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, importQueryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from importQuery slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for import_queries")
	}

	if len(importQueryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ImportQuery) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ImportQuery provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ImportQuery) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindImportQuery(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ImportQuerySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ImportQuerySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ImportQuerySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ImportQuerySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), importQueryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[import_queries].* FROM [dbo].[import_queries] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, importQueryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ImportQuerySlice")
	}

	*o = slice

	return nil
}

// ImportQueryExistsG checks if the ImportQuery row exists.
func ImportQueryExistsG(ctx context.Context, iD int) (bool, error) {
	return ImportQueryExists(ctx, boil.GetContextDB(), iD)
}

// ImportQueryExists checks if the ImportQuery row exists.
func ImportQueryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[import_queries] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if import_queries exists")
	}

	return exists, nil
}

// Exists checks if the ImportQuery row exists.
func (o *ImportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ImportQueryExists(ctx, exec, o.ID)
}