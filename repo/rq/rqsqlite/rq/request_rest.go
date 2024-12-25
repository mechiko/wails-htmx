// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package rq

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

// RequestRest is an object representing the database table.
type RequestRest struct {
	ID            int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Fsrarid       null.String `boil:"fsrarid" json:"fsrarid,omitempty" toml:"fsrarid" yaml:"fsrarid,omitempty"`
	RequestID     null.Int64  `boil:"request_id" json:"request_id,omitempty" toml:"request_id" yaml:"request_id,omitempty"`
	CreatedAt     null.String `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.String `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Daterest      null.String `boil:"daterest" json:"daterest,omitempty" toml:"daterest" yaml:"daterest,omitempty"`
	Quantity      null.String `boil:"quantity" json:"quantity,omitempty" toml:"quantity" yaml:"quantity,omitempty"`
	Informf1regid null.String `boil:"informf1regid" json:"informf1regid,omitempty" toml:"informf1regid" yaml:"informf1regid,omitempty"`
	Informf2regid null.String `boil:"informf2regid" json:"informf2regid,omitempty" toml:"informf2regid" yaml:"informf2regid,omitempty"`
	Alccode       null.String `boil:"alccode" json:"alccode,omitempty" toml:"alccode" yaml:"alccode,omitempty"`

	R *requestRestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L requestRestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RequestRestColumns = struct {
	ID            string
	Fsrarid       string
	RequestID     string
	CreatedAt     string
	UpdatedAt     string
	Daterest      string
	Quantity      string
	Informf1regid string
	Informf2regid string
	Alccode       string
}{
	ID:            "id",
	Fsrarid:       "fsrarid",
	RequestID:     "request_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	Daterest:      "daterest",
	Quantity:      "quantity",
	Informf1regid: "informf1regid",
	Informf2regid: "informf2regid",
	Alccode:       "alccode",
}

var RequestRestTableColumns = struct {
	ID            string
	Fsrarid       string
	RequestID     string
	CreatedAt     string
	UpdatedAt     string
	Daterest      string
	Quantity      string
	Informf1regid string
	Informf2regid string
	Alccode       string
}{
	ID:            "request_rest.id",
	Fsrarid:       "request_rest.fsrarid",
	RequestID:     "request_rest.request_id",
	CreatedAt:     "request_rest.created_at",
	UpdatedAt:     "request_rest.updated_at",
	Daterest:      "request_rest.daterest",
	Quantity:      "request_rest.quantity",
	Informf1regid: "request_rest.informf1regid",
	Informf2regid: "request_rest.informf2regid",
	Alccode:       "request_rest.alccode",
}

// Generated where

var RequestRestWhere = struct {
	ID            whereHelperint64
	Fsrarid       whereHelpernull_String
	RequestID     whereHelpernull_Int64
	CreatedAt     whereHelpernull_String
	UpdatedAt     whereHelpernull_String
	Daterest      whereHelpernull_String
	Quantity      whereHelpernull_String
	Informf1regid whereHelpernull_String
	Informf2regid whereHelpernull_String
	Alccode       whereHelpernull_String
}{
	ID:            whereHelperint64{field: "\"request_rest\".\"id\""},
	Fsrarid:       whereHelpernull_String{field: "\"request_rest\".\"fsrarid\""},
	RequestID:     whereHelpernull_Int64{field: "\"request_rest\".\"request_id\""},
	CreatedAt:     whereHelpernull_String{field: "\"request_rest\".\"created_at\""},
	UpdatedAt:     whereHelpernull_String{field: "\"request_rest\".\"updated_at\""},
	Daterest:      whereHelpernull_String{field: "\"request_rest\".\"daterest\""},
	Quantity:      whereHelpernull_String{field: "\"request_rest\".\"quantity\""},
	Informf1regid: whereHelpernull_String{field: "\"request_rest\".\"informf1regid\""},
	Informf2regid: whereHelpernull_String{field: "\"request_rest\".\"informf2regid\""},
	Alccode:       whereHelpernull_String{field: "\"request_rest\".\"alccode\""},
}

// RequestRestRels is where relationship names are stored.
var RequestRestRels = struct {
}{}

// requestRestR is where relationships are stored.
type requestRestR struct {
}

// NewStruct creates a new relationship struct
func (*requestRestR) NewStruct() *requestRestR {
	return &requestRestR{}
}

// requestRestL is where Load methods for each relationship are stored.
type requestRestL struct{}

var (
	requestRestAllColumns            = []string{"id", "fsrarid", "request_id", "created_at", "updated_at", "daterest", "quantity", "informf1regid", "informf2regid", "alccode"}
	requestRestColumnsWithoutDefault = []string{}
	requestRestColumnsWithDefault    = []string{"id", "fsrarid", "request_id", "created_at", "updated_at", "daterest", "quantity", "informf1regid", "informf2regid", "alccode"}
	requestRestPrimaryKeyColumns     = []string{"id"}
	requestRestGeneratedColumns      = []string{"id"}
)

type (
	// RequestRestSlice is an alias for a slice of pointers to RequestRest.
	// This should almost always be used instead of []RequestRest.
	RequestRestSlice []*RequestRest
	// RequestRestHook is the signature for custom RequestRest hook methods
	RequestRestHook func(context.Context, boil.ContextExecutor, *RequestRest) error

	requestRestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	requestRestType                 = reflect.TypeOf(&RequestRest{})
	requestRestMapping              = queries.MakeStructMapping(requestRestType)
	requestRestPrimaryKeyMapping, _ = queries.BindMapping(requestRestType, requestRestMapping, requestRestPrimaryKeyColumns)
	requestRestInsertCacheMut       sync.RWMutex
	requestRestInsertCache          = make(map[string]insertCache)
	requestRestUpdateCacheMut       sync.RWMutex
	requestRestUpdateCache          = make(map[string]updateCache)
	requestRestUpsertCacheMut       sync.RWMutex
	requestRestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var requestRestAfterSelectMu sync.Mutex
var requestRestAfterSelectHooks []RequestRestHook

var requestRestBeforeInsertMu sync.Mutex
var requestRestBeforeInsertHooks []RequestRestHook
var requestRestAfterInsertMu sync.Mutex
var requestRestAfterInsertHooks []RequestRestHook

var requestRestBeforeUpdateMu sync.Mutex
var requestRestBeforeUpdateHooks []RequestRestHook
var requestRestAfterUpdateMu sync.Mutex
var requestRestAfterUpdateHooks []RequestRestHook

var requestRestBeforeDeleteMu sync.Mutex
var requestRestBeforeDeleteHooks []RequestRestHook
var requestRestAfterDeleteMu sync.Mutex
var requestRestAfterDeleteHooks []RequestRestHook

var requestRestBeforeUpsertMu sync.Mutex
var requestRestBeforeUpsertHooks []RequestRestHook
var requestRestAfterUpsertMu sync.Mutex
var requestRestAfterUpsertHooks []RequestRestHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RequestRest) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RequestRest) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RequestRest) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RequestRest) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RequestRest) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RequestRest) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RequestRest) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RequestRest) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RequestRest) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestRestAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRequestRestHook registers your hook function for all future operations.
func AddRequestRestHook(hookPoint boil.HookPoint, requestRestHook RequestRestHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		requestRestAfterSelectMu.Lock()
		requestRestAfterSelectHooks = append(requestRestAfterSelectHooks, requestRestHook)
		requestRestAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		requestRestBeforeInsertMu.Lock()
		requestRestBeforeInsertHooks = append(requestRestBeforeInsertHooks, requestRestHook)
		requestRestBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		requestRestAfterInsertMu.Lock()
		requestRestAfterInsertHooks = append(requestRestAfterInsertHooks, requestRestHook)
		requestRestAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		requestRestBeforeUpdateMu.Lock()
		requestRestBeforeUpdateHooks = append(requestRestBeforeUpdateHooks, requestRestHook)
		requestRestBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		requestRestAfterUpdateMu.Lock()
		requestRestAfterUpdateHooks = append(requestRestAfterUpdateHooks, requestRestHook)
		requestRestAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		requestRestBeforeDeleteMu.Lock()
		requestRestBeforeDeleteHooks = append(requestRestBeforeDeleteHooks, requestRestHook)
		requestRestBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		requestRestAfterDeleteMu.Lock()
		requestRestAfterDeleteHooks = append(requestRestAfterDeleteHooks, requestRestHook)
		requestRestAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		requestRestBeforeUpsertMu.Lock()
		requestRestBeforeUpsertHooks = append(requestRestBeforeUpsertHooks, requestRestHook)
		requestRestBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		requestRestAfterUpsertMu.Lock()
		requestRestAfterUpsertHooks = append(requestRestAfterUpsertHooks, requestRestHook)
		requestRestAfterUpsertMu.Unlock()
	}
}

// OneG returns a single requestRest record from the query using the global executor.
func (q requestRestQuery) OneG(ctx context.Context) (*RequestRest, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single requestRest record from the query.
func (q requestRestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RequestRest, error) {
	o := &RequestRest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: failed to execute a one query for request_rest")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all RequestRest records from the query using the global executor.
func (q requestRestQuery) AllG(ctx context.Context) (RequestRestSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all RequestRest records from the query.
func (q requestRestQuery) All(ctx context.Context, exec boil.ContextExecutor) (RequestRestSlice, error) {
	var o []*RequestRest

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rq: failed to assign all query results to RequestRest slice")
	}

	if len(requestRestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all RequestRest records in the query using the global executor
func (q requestRestQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all RequestRest records in the query.
func (q requestRestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to count request_rest rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q requestRestQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q requestRestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rq: failed to check if request_rest exists")
	}

	return count > 0, nil
}

// RequestRests retrieves all the records using an executor.
func RequestRests(mods ...qm.QueryMod) requestRestQuery {
	mods = append(mods, qm.From("\"request_rest\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"request_rest\".*"})
	}

	return requestRestQuery{q}
}

// FindRequestRestG retrieves a single record by ID.
func FindRequestRestG(ctx context.Context, iD int64, selectCols ...string) (*RequestRest, error) {
	return FindRequestRest(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindRequestRest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRequestRest(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RequestRest, error) {
	requestRestObj := &RequestRest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"request_rest\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, requestRestObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: unable to select from request_rest")
	}

	if err = requestRestObj.doAfterSelectHooks(ctx, exec); err != nil {
		return requestRestObj, err
	}

	return requestRestObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *RequestRest) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RequestRest) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no request_rest provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestRestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	requestRestInsertCacheMut.RLock()
	cache, cached := requestRestInsertCache[key]
	requestRestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			requestRestAllColumns,
			requestRestColumnsWithDefault,
			requestRestColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, requestRestGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(requestRestType, requestRestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(requestRestType, requestRestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"request_rest\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"request_rest\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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
		return errors.Wrap(err, "rq: unable to insert into request_rest")
	}

	if !cached {
		requestRestInsertCacheMut.Lock()
		requestRestInsertCache[key] = cache
		requestRestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single RequestRest record using the global executor.
// See Update for more documentation.
func (o *RequestRest) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the RequestRest.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RequestRest) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	requestRestUpdateCacheMut.RLock()
	cache, cached := requestRestUpdateCache[key]
	requestRestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			requestRestAllColumns,
			requestRestPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, requestRestGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("rq: unable to update request_rest, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"request_rest\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, requestRestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(requestRestType, requestRestMapping, append(wl, requestRestPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rq: unable to update request_rest row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by update for request_rest")
	}

	if !cached {
		requestRestUpdateCacheMut.Lock()
		requestRestUpdateCache[key] = cache
		requestRestUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q requestRestQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q requestRestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all for request_rest")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected for request_rest")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RequestRestSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RequestRestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("rq: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestRestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"request_rest\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestRestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all in requestRest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected all in update all requestRest")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *RequestRest) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RequestRest) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no request_rest provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestRestColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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

	requestRestUpsertCacheMut.RLock()
	cache, cached := requestRestUpsertCache[key]
	requestRestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			requestRestAllColumns,
			requestRestColumnsWithDefault,
			requestRestColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			requestRestAllColumns,
			requestRestPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("rq: unable to upsert request_rest, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(requestRestPrimaryKeyColumns))
			copy(conflict, requestRestPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"request_rest\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(requestRestType, requestRestMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(requestRestType, requestRestMapping, ret)
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
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "rq: unable to upsert request_rest")
	}

	if !cached {
		requestRestUpsertCacheMut.Lock()
		requestRestUpsertCache[key] = cache
		requestRestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single RequestRest record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *RequestRest) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single RequestRest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RequestRest) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rq: no RequestRest provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), requestRestPrimaryKeyMapping)
	sql := "DELETE FROM \"request_rest\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete from request_rest")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by delete for request_rest")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q requestRestQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q requestRestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rq: no requestRestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from request_rest")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for request_rest")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RequestRestSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RequestRestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(requestRestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestRestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"request_rest\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestRestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from requestRest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for request_rest")
	}

	if len(requestRestAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *RequestRest) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: no RequestRest provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RequestRest) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRequestRest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RequestRestSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: empty RequestRestSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RequestRestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RequestRestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestRestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"request_rest\".* FROM \"request_rest\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestRestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rq: unable to reload all in RequestRestSlice")
	}

	*o = slice

	return nil
}

// RequestRestExistsG checks if the RequestRest row exists.
func RequestRestExistsG(ctx context.Context, iD int64) (bool, error) {
	return RequestRestExists(ctx, boil.GetContextDB(), iD)
}

// RequestRestExists checks if the RequestRest row exists.
func RequestRestExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"request_rest\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rq: unable to check if request_rest exists")
	}

	return exists, nil
}

// Exists checks if the RequestRest row exists.
func (o *RequestRest) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RequestRestExists(ctx, exec, o.ID)
}