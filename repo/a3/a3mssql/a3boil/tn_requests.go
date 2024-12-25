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

// TNRequest is an object representing the database table.
type TNRequest struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTN         null.Int    `boil:"id_tn" json:"id_tn,omitempty" toml:"id_tn" yaml:"id_tn,omitempty"`
	RequestDate  null.String `boil:"request_date" json:"request_date,omitempty" toml:"request_date" yaml:"request_date,omitempty"`
	RequestRegID null.String `boil:"request_reg_id" json:"request_reg_id,omitempty" toml:"request_reg_id" yaml:"request_reg_id,omitempty"`
	Status       null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID      null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`

	R *tnRequestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tnRequestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TNRequestColumns = struct {
	ID           string
	IDTN         string
	RequestDate  string
	RequestRegID string
	Status       string
	ReplyID      string
}{
	ID:           "id",
	IDTN:         "id_tn",
	RequestDate:  "request_date",
	RequestRegID: "request_reg_id",
	Status:       "status",
	ReplyID:      "reply_id",
}

var TNRequestTableColumns = struct {
	ID           string
	IDTN         string
	RequestDate  string
	RequestRegID string
	Status       string
	ReplyID      string
}{
	ID:           "tn_requests.id",
	IDTN:         "tn_requests.id_tn",
	RequestDate:  "tn_requests.request_date",
	RequestRegID: "tn_requests.request_reg_id",
	Status:       "tn_requests.status",
	ReplyID:      "tn_requests.reply_id",
}

// Generated where

var TNRequestWhere = struct {
	ID           whereHelperint
	IDTN         whereHelpernull_Int
	RequestDate  whereHelpernull_String
	RequestRegID whereHelpernull_String
	Status       whereHelpernull_String
	ReplyID      whereHelpernull_String
}{
	ID:           whereHelperint{field: "[dbo].[tn_requests].[id]"},
	IDTN:         whereHelpernull_Int{field: "[dbo].[tn_requests].[id_tn]"},
	RequestDate:  whereHelpernull_String{field: "[dbo].[tn_requests].[request_date]"},
	RequestRegID: whereHelpernull_String{field: "[dbo].[tn_requests].[request_reg_id]"},
	Status:       whereHelpernull_String{field: "[dbo].[tn_requests].[status]"},
	ReplyID:      whereHelpernull_String{field: "[dbo].[tn_requests].[reply_id]"},
}

// TNRequestRels is where relationship names are stored.
var TNRequestRels = struct {
}{}

// tnRequestR is where relationships are stored.
type tnRequestR struct {
}

// NewStruct creates a new relationship struct
func (*tnRequestR) NewStruct() *tnRequestR {
	return &tnRequestR{}
}

// tnRequestL is where Load methods for each relationship are stored.
type tnRequestL struct{}

var (
	tnRequestAllColumns            = []string{"id", "id_tn", "request_date", "request_reg_id", "status", "reply_id"}
	tnRequestColumnsWithoutDefault = []string{"id_tn", "request_date", "request_reg_id", "status", "reply_id"}
	tnRequestColumnsWithDefault    = []string{"id"}
	tnRequestPrimaryKeyColumns     = []string{"id"}
	tnRequestGeneratedColumns      = []string{"id"}
)

type (
	// TNRequestSlice is an alias for a slice of pointers to TNRequest.
	// This should almost always be used instead of []TNRequest.
	TNRequestSlice []*TNRequest
	// TNRequestHook is the signature for custom TNRequest hook methods
	TNRequestHook func(context.Context, boil.ContextExecutor, *TNRequest) error

	tnRequestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tnRequestType                 = reflect.TypeOf(&TNRequest{})
	tnRequestMapping              = queries.MakeStructMapping(tnRequestType)
	tnRequestPrimaryKeyMapping, _ = queries.BindMapping(tnRequestType, tnRequestMapping, tnRequestPrimaryKeyColumns)
	tnRequestInsertCacheMut       sync.RWMutex
	tnRequestInsertCache          = make(map[string]insertCache)
	tnRequestUpdateCacheMut       sync.RWMutex
	tnRequestUpdateCache          = make(map[string]updateCache)
	tnRequestUpsertCacheMut       sync.RWMutex
	tnRequestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tnRequestAfterSelectMu sync.Mutex
var tnRequestAfterSelectHooks []TNRequestHook

var tnRequestBeforeInsertMu sync.Mutex
var tnRequestBeforeInsertHooks []TNRequestHook
var tnRequestAfterInsertMu sync.Mutex
var tnRequestAfterInsertHooks []TNRequestHook

var tnRequestBeforeUpdateMu sync.Mutex
var tnRequestBeforeUpdateHooks []TNRequestHook
var tnRequestAfterUpdateMu sync.Mutex
var tnRequestAfterUpdateHooks []TNRequestHook

var tnRequestBeforeDeleteMu sync.Mutex
var tnRequestBeforeDeleteHooks []TNRequestHook
var tnRequestAfterDeleteMu sync.Mutex
var tnRequestAfterDeleteHooks []TNRequestHook

var tnRequestBeforeUpsertMu sync.Mutex
var tnRequestBeforeUpsertHooks []TNRequestHook
var tnRequestAfterUpsertMu sync.Mutex
var tnRequestAfterUpsertHooks []TNRequestHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TNRequest) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TNRequest) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TNRequest) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TNRequest) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TNRequest) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TNRequest) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TNRequest) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TNRequest) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TNRequest) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnRequestAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTNRequestHook registers your hook function for all future operations.
func AddTNRequestHook(hookPoint boil.HookPoint, tnRequestHook TNRequestHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tnRequestAfterSelectMu.Lock()
		tnRequestAfterSelectHooks = append(tnRequestAfterSelectHooks, tnRequestHook)
		tnRequestAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tnRequestBeforeInsertMu.Lock()
		tnRequestBeforeInsertHooks = append(tnRequestBeforeInsertHooks, tnRequestHook)
		tnRequestBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tnRequestAfterInsertMu.Lock()
		tnRequestAfterInsertHooks = append(tnRequestAfterInsertHooks, tnRequestHook)
		tnRequestAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tnRequestBeforeUpdateMu.Lock()
		tnRequestBeforeUpdateHooks = append(tnRequestBeforeUpdateHooks, tnRequestHook)
		tnRequestBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tnRequestAfterUpdateMu.Lock()
		tnRequestAfterUpdateHooks = append(tnRequestAfterUpdateHooks, tnRequestHook)
		tnRequestAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tnRequestBeforeDeleteMu.Lock()
		tnRequestBeforeDeleteHooks = append(tnRequestBeforeDeleteHooks, tnRequestHook)
		tnRequestBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tnRequestAfterDeleteMu.Lock()
		tnRequestAfterDeleteHooks = append(tnRequestAfterDeleteHooks, tnRequestHook)
		tnRequestAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tnRequestBeforeUpsertMu.Lock()
		tnRequestBeforeUpsertHooks = append(tnRequestBeforeUpsertHooks, tnRequestHook)
		tnRequestBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tnRequestAfterUpsertMu.Lock()
		tnRequestAfterUpsertHooks = append(tnRequestAfterUpsertHooks, tnRequestHook)
		tnRequestAfterUpsertMu.Unlock()
	}
}

// OneG returns a single tnRequest record from the query using the global executor.
func (q tnRequestQuery) OneG(ctx context.Context) (*TNRequest, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single tnRequest record from the query.
func (q tnRequestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TNRequest, error) {
	o := &TNRequest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for tn_requests")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TNRequest records from the query using the global executor.
func (q tnRequestQuery) AllG(ctx context.Context) (TNRequestSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TNRequest records from the query.
func (q tnRequestQuery) All(ctx context.Context, exec boil.ContextExecutor) (TNRequestSlice, error) {
	var o []*TNRequest

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TNRequest slice")
	}

	if len(tnRequestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TNRequest records in the query using the global executor
func (q tnRequestQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TNRequest records in the query.
func (q tnRequestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count tn_requests rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q tnRequestQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q tnRequestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if tn_requests exists")
	}

	return count > 0, nil
}

// TNRequests retrieves all the records using an executor.
func TNRequests(mods ...qm.QueryMod) tnRequestQuery {
	mods = append(mods, qm.From("[dbo].[tn_requests]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[tn_requests].*"})
	}

	return tnRequestQuery{q}
}

// FindTNRequestG retrieves a single record by ID.
func FindTNRequestG(ctx context.Context, iD int, selectCols ...string) (*TNRequest, error) {
	return FindTNRequest(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTNRequest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTNRequest(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TNRequest, error) {
	tnRequestObj := &TNRequest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[tn_requests] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tnRequestObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from tn_requests")
	}

	if err = tnRequestObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tnRequestObj, err
	}

	return tnRequestObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TNRequest) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TNRequest) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn_requests provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnRequestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tnRequestInsertCacheMut.RLock()
	cache, cached := tnRequestInsertCache[key]
	tnRequestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tnRequestAllColumns,
			tnRequestColumnsWithDefault,
			tnRequestColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, tnRequestGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(tnRequestType, tnRequestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tnRequestType, tnRequestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[tn_requests] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[tn_requests] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into tn_requests")
	}

	if !cached {
		tnRequestInsertCacheMut.Lock()
		tnRequestInsertCache[key] = cache
		tnRequestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TNRequest record using the global executor.
// See Update for more documentation.
func (o *TNRequest) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TNRequest.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TNRequest) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tnRequestUpdateCacheMut.RLock()
	cache, cached := tnRequestUpdateCache[key]
	tnRequestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tnRequestAllColumns,
			tnRequestPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tnRequestGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update tn_requests, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[tn_requests] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, tnRequestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tnRequestType, tnRequestMapping, append(wl, tnRequestPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update tn_requests row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for tn_requests")
	}

	if !cached {
		tnRequestUpdateCacheMut.Lock()
		tnRequestUpdateCache[key] = cache
		tnRequestUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tnRequestQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tnRequestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for tn_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for tn_requests")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TNRequestSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TNRequestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[tn_requests] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tnRequestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in tnRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all tnRequest")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TNRequest) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TNRequest) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn_requests provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnRequestColumnsWithDefault, o)

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

	tnRequestUpsertCacheMut.RLock()
	cache, cached := tnRequestUpsertCache[key]
	tnRequestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tnRequestAllColumns,
			tnRequestColumnsWithDefault,
			tnRequestColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, tnRequestGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(tnRequestPrimaryKeyColumns, v) && strmangle.ContainsAny(tnRequestColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert tn_requests, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			tnRequestAllColumns,
			tnRequestPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, tnRequestGeneratedColumns)

		ret := strmangle.SetComplement(tnRequestAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert tn_requests, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[tn_requests]", tnRequestPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(tnRequestPrimaryKeyColumns))
		copy(whitelist, tnRequestPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(tnRequestType, tnRequestMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tnRequestType, tnRequestMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert tn_requests")
	}

	if !cached {
		tnRequestUpsertCacheMut.Lock()
		tnRequestUpsertCache[key] = cache
		tnRequestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TNRequest record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TNRequest) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TNRequest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TNRequest) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TNRequest provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tnRequestPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[tn_requests] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from tn_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for tn_requests")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q tnRequestQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q tnRequestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no tnRequestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tn_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn_requests")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TNRequestSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TNRequestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tnRequestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[tn_requests] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tnRequestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tnRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn_requests")
	}

	if len(tnRequestAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TNRequest) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TNRequest provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TNRequest) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTNRequest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNRequestSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TNRequestSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNRequestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TNRequestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[tn_requests].* FROM [dbo].[tn_requests] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tnRequestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TNRequestSlice")
	}

	*o = slice

	return nil
}

// TNRequestExistsG checks if the TNRequest row exists.
func TNRequestExistsG(ctx context.Context, iD int) (bool, error) {
	return TNRequestExists(ctx, boil.GetContextDB(), iD)
}

// TNRequestExists checks if the TNRequest row exists.
func TNRequestExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[tn_requests] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if tn_requests exists")
	}

	return exists, nil
}

// Exists checks if the TNRequest row exists.
func (o *TNRequest) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TNRequestExists(ctx, exec, o.ID)
}
