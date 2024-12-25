// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package znakboil

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

// OrderMarkAggregationCode is an object representing the database table.
type OrderMarkAggregationCode struct {
	ID                     int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDOrderMarkAggregation null.Int64  `boil:"id_order_mark_aggregation" json:"id_order_mark_aggregation,omitempty" toml:"id_order_mark_aggregation" yaml:"id_order_mark_aggregation,omitempty"`
	SerialNumber           null.String `boil:"serial_number" json:"serial_number,omitempty" toml:"serial_number" yaml:"serial_number,omitempty"`
	Code                   null.String `boil:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	Status                 null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`

	R *orderMarkAggregationCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderMarkAggregationCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderMarkAggregationCodeColumns = struct {
	ID                     string
	IDOrderMarkAggregation string
	SerialNumber           string
	Code                   string
	Status                 string
}{
	ID:                     "id",
	IDOrderMarkAggregation: "id_order_mark_aggregation",
	SerialNumber:           "serial_number",
	Code:                   "code",
	Status:                 "status",
}

var OrderMarkAggregationCodeTableColumns = struct {
	ID                     string
	IDOrderMarkAggregation string
	SerialNumber           string
	Code                   string
	Status                 string
}{
	ID:                     "order_mark_aggregation_codes.id",
	IDOrderMarkAggregation: "order_mark_aggregation_codes.id_order_mark_aggregation",
	SerialNumber:           "order_mark_aggregation_codes.serial_number",
	Code:                   "order_mark_aggregation_codes.code",
	Status:                 "order_mark_aggregation_codes.status",
}

// Generated where

var OrderMarkAggregationCodeWhere = struct {
	ID                     whereHelperint64
	IDOrderMarkAggregation whereHelpernull_Int64
	SerialNumber           whereHelpernull_String
	Code                   whereHelpernull_String
	Status                 whereHelpernull_String
}{
	ID:                     whereHelperint64{field: "\"order_mark_aggregation_codes\".\"id\""},
	IDOrderMarkAggregation: whereHelpernull_Int64{field: "\"order_mark_aggregation_codes\".\"id_order_mark_aggregation\""},
	SerialNumber:           whereHelpernull_String{field: "\"order_mark_aggregation_codes\".\"serial_number\""},
	Code:                   whereHelpernull_String{field: "\"order_mark_aggregation_codes\".\"code\""},
	Status:                 whereHelpernull_String{field: "\"order_mark_aggregation_codes\".\"status\""},
}

// OrderMarkAggregationCodeRels is where relationship names are stored.
var OrderMarkAggregationCodeRels = struct {
}{}

// orderMarkAggregationCodeR is where relationships are stored.
type orderMarkAggregationCodeR struct {
}

// NewStruct creates a new relationship struct
func (*orderMarkAggregationCodeR) NewStruct() *orderMarkAggregationCodeR {
	return &orderMarkAggregationCodeR{}
}

// orderMarkAggregationCodeL is where Load methods for each relationship are stored.
type orderMarkAggregationCodeL struct{}

var (
	orderMarkAggregationCodeAllColumns            = []string{"id", "id_order_mark_aggregation", "serial_number", "code", "status"}
	orderMarkAggregationCodeColumnsWithoutDefault = []string{}
	orderMarkAggregationCodeColumnsWithDefault    = []string{"id", "id_order_mark_aggregation", "serial_number", "code", "status"}
	orderMarkAggregationCodePrimaryKeyColumns     = []string{"id"}
	orderMarkAggregationCodeGeneratedColumns      = []string{"id"}
)

type (
	// OrderMarkAggregationCodeSlice is an alias for a slice of pointers to OrderMarkAggregationCode.
	// This should almost always be used instead of []OrderMarkAggregationCode.
	OrderMarkAggregationCodeSlice []*OrderMarkAggregationCode
	// OrderMarkAggregationCodeHook is the signature for custom OrderMarkAggregationCode hook methods
	OrderMarkAggregationCodeHook func(context.Context, boil.ContextExecutor, *OrderMarkAggregationCode) error

	orderMarkAggregationCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderMarkAggregationCodeType                 = reflect.TypeOf(&OrderMarkAggregationCode{})
	orderMarkAggregationCodeMapping              = queries.MakeStructMapping(orderMarkAggregationCodeType)
	orderMarkAggregationCodePrimaryKeyMapping, _ = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, orderMarkAggregationCodePrimaryKeyColumns)
	orderMarkAggregationCodeInsertCacheMut       sync.RWMutex
	orderMarkAggregationCodeInsertCache          = make(map[string]insertCache)
	orderMarkAggregationCodeUpdateCacheMut       sync.RWMutex
	orderMarkAggregationCodeUpdateCache          = make(map[string]updateCache)
	orderMarkAggregationCodeUpsertCacheMut       sync.RWMutex
	orderMarkAggregationCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderMarkAggregationCodeAfterSelectMu sync.Mutex
var orderMarkAggregationCodeAfterSelectHooks []OrderMarkAggregationCodeHook

var orderMarkAggregationCodeBeforeInsertMu sync.Mutex
var orderMarkAggregationCodeBeforeInsertHooks []OrderMarkAggregationCodeHook
var orderMarkAggregationCodeAfterInsertMu sync.Mutex
var orderMarkAggregationCodeAfterInsertHooks []OrderMarkAggregationCodeHook

var orderMarkAggregationCodeBeforeUpdateMu sync.Mutex
var orderMarkAggregationCodeBeforeUpdateHooks []OrderMarkAggregationCodeHook
var orderMarkAggregationCodeAfterUpdateMu sync.Mutex
var orderMarkAggregationCodeAfterUpdateHooks []OrderMarkAggregationCodeHook

var orderMarkAggregationCodeBeforeDeleteMu sync.Mutex
var orderMarkAggregationCodeBeforeDeleteHooks []OrderMarkAggregationCodeHook
var orderMarkAggregationCodeAfterDeleteMu sync.Mutex
var orderMarkAggregationCodeAfterDeleteHooks []OrderMarkAggregationCodeHook

var orderMarkAggregationCodeBeforeUpsertMu sync.Mutex
var orderMarkAggregationCodeBeforeUpsertHooks []OrderMarkAggregationCodeHook
var orderMarkAggregationCodeAfterUpsertMu sync.Mutex
var orderMarkAggregationCodeAfterUpsertHooks []OrderMarkAggregationCodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderMarkAggregationCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderMarkAggregationCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderMarkAggregationCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderMarkAggregationCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderMarkAggregationCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderMarkAggregationCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderMarkAggregationCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderMarkAggregationCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderMarkAggregationCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkAggregationCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderMarkAggregationCodeHook registers your hook function for all future operations.
func AddOrderMarkAggregationCodeHook(hookPoint boil.HookPoint, orderMarkAggregationCodeHook OrderMarkAggregationCodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderMarkAggregationCodeAfterSelectMu.Lock()
		orderMarkAggregationCodeAfterSelectHooks = append(orderMarkAggregationCodeAfterSelectHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderMarkAggregationCodeBeforeInsertMu.Lock()
		orderMarkAggregationCodeBeforeInsertHooks = append(orderMarkAggregationCodeBeforeInsertHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderMarkAggregationCodeAfterInsertMu.Lock()
		orderMarkAggregationCodeAfterInsertHooks = append(orderMarkAggregationCodeAfterInsertHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		orderMarkAggregationCodeBeforeUpdateMu.Lock()
		orderMarkAggregationCodeBeforeUpdateHooks = append(orderMarkAggregationCodeBeforeUpdateHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		orderMarkAggregationCodeAfterUpdateMu.Lock()
		orderMarkAggregationCodeAfterUpdateHooks = append(orderMarkAggregationCodeAfterUpdateHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		orderMarkAggregationCodeBeforeDeleteMu.Lock()
		orderMarkAggregationCodeBeforeDeleteHooks = append(orderMarkAggregationCodeBeforeDeleteHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		orderMarkAggregationCodeAfterDeleteMu.Lock()
		orderMarkAggregationCodeAfterDeleteHooks = append(orderMarkAggregationCodeAfterDeleteHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		orderMarkAggregationCodeBeforeUpsertMu.Lock()
		orderMarkAggregationCodeBeforeUpsertHooks = append(orderMarkAggregationCodeBeforeUpsertHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderMarkAggregationCodeAfterUpsertMu.Lock()
		orderMarkAggregationCodeAfterUpsertHooks = append(orderMarkAggregationCodeAfterUpsertHooks, orderMarkAggregationCodeHook)
		orderMarkAggregationCodeAfterUpsertMu.Unlock()
	}
}

// OneG returns a single orderMarkAggregationCode record from the query using the global executor.
func (q orderMarkAggregationCodeQuery) OneG(ctx context.Context) (*OrderMarkAggregationCode, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single orderMarkAggregationCode record from the query.
func (q orderMarkAggregationCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderMarkAggregationCode, error) {
	o := &OrderMarkAggregationCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: failed to execute a one query for order_mark_aggregation_codes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all OrderMarkAggregationCode records from the query using the global executor.
func (q orderMarkAggregationCodeQuery) AllG(ctx context.Context) (OrderMarkAggregationCodeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all OrderMarkAggregationCode records from the query.
func (q orderMarkAggregationCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderMarkAggregationCodeSlice, error) {
	var o []*OrderMarkAggregationCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "znakboil: failed to assign all query results to OrderMarkAggregationCode slice")
	}

	if len(orderMarkAggregationCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all OrderMarkAggregationCode records in the query using the global executor
func (q orderMarkAggregationCodeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all OrderMarkAggregationCode records in the query.
func (q orderMarkAggregationCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to count order_mark_aggregation_codes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q orderMarkAggregationCodeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q orderMarkAggregationCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: failed to check if order_mark_aggregation_codes exists")
	}

	return count > 0, nil
}

// OrderMarkAggregationCodes retrieves all the records using an executor.
func OrderMarkAggregationCodes(mods ...qm.QueryMod) orderMarkAggregationCodeQuery {
	mods = append(mods, qm.From("\"order_mark_aggregation_codes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"order_mark_aggregation_codes\".*"})
	}

	return orderMarkAggregationCodeQuery{q}
}

// FindOrderMarkAggregationCodeG retrieves a single record by ID.
func FindOrderMarkAggregationCodeG(ctx context.Context, iD int64, selectCols ...string) (*OrderMarkAggregationCode, error) {
	return FindOrderMarkAggregationCode(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOrderMarkAggregationCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderMarkAggregationCode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OrderMarkAggregationCode, error) {
	orderMarkAggregationCodeObj := &OrderMarkAggregationCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order_mark_aggregation_codes\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderMarkAggregationCodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: unable to select from order_mark_aggregation_codes")
	}

	if err = orderMarkAggregationCodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderMarkAggregationCodeObj, err
	}

	return orderMarkAggregationCodeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OrderMarkAggregationCode) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderMarkAggregationCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_aggregation_codes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkAggregationCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderMarkAggregationCodeInsertCacheMut.RLock()
	cache, cached := orderMarkAggregationCodeInsertCache[key]
	orderMarkAggregationCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderMarkAggregationCodeAllColumns,
			orderMarkAggregationCodeColumnsWithDefault,
			orderMarkAggregationCodeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, orderMarkAggregationCodeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order_mark_aggregation_codes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order_mark_aggregation_codes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "znakboil: unable to insert into order_mark_aggregation_codes")
	}

	if !cached {
		orderMarkAggregationCodeInsertCacheMut.Lock()
		orderMarkAggregationCodeInsertCache[key] = cache
		orderMarkAggregationCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single OrderMarkAggregationCode record using the global executor.
// See Update for more documentation.
func (o *OrderMarkAggregationCode) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the OrderMarkAggregationCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderMarkAggregationCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderMarkAggregationCodeUpdateCacheMut.RLock()
	cache, cached := orderMarkAggregationCodeUpdateCache[key]
	orderMarkAggregationCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderMarkAggregationCodeAllColumns,
			orderMarkAggregationCodePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, orderMarkAggregationCodeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("znakboil: unable to update order_mark_aggregation_codes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order_mark_aggregation_codes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, orderMarkAggregationCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, append(wl, orderMarkAggregationCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "znakboil: unable to update order_mark_aggregation_codes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by update for order_mark_aggregation_codes")
	}

	if !cached {
		orderMarkAggregationCodeUpdateCacheMut.Lock()
		orderMarkAggregationCodeUpdateCache[key] = cache
		orderMarkAggregationCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q orderMarkAggregationCodeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q orderMarkAggregationCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all for order_mark_aggregation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected for order_mark_aggregation_codes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderMarkAggregationCodeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderMarkAggregationCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("znakboil: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkAggregationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order_mark_aggregation_codes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkAggregationCodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all in orderMarkAggregationCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected all in update all orderMarkAggregationCode")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OrderMarkAggregationCode) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderMarkAggregationCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_aggregation_codes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkAggregationCodeColumnsWithDefault, o)

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

	orderMarkAggregationCodeUpsertCacheMut.RLock()
	cache, cached := orderMarkAggregationCodeUpsertCache[key]
	orderMarkAggregationCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderMarkAggregationCodeAllColumns,
			orderMarkAggregationCodeColumnsWithDefault,
			orderMarkAggregationCodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orderMarkAggregationCodeAllColumns,
			orderMarkAggregationCodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("znakboil: unable to upsert order_mark_aggregation_codes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderMarkAggregationCodePrimaryKeyColumns))
			copy(conflict, orderMarkAggregationCodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"order_mark_aggregation_codes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderMarkAggregationCodeType, orderMarkAggregationCodeMapping, ret)
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
		return errors.Wrap(err, "znakboil: unable to upsert order_mark_aggregation_codes")
	}

	if !cached {
		orderMarkAggregationCodeUpsertCacheMut.Lock()
		orderMarkAggregationCodeUpsertCache[key] = cache
		orderMarkAggregationCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single OrderMarkAggregationCode record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OrderMarkAggregationCode) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single OrderMarkAggregationCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderMarkAggregationCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("znakboil: no OrderMarkAggregationCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderMarkAggregationCodePrimaryKeyMapping)
	sql := "DELETE FROM \"order_mark_aggregation_codes\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete from order_mark_aggregation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by delete for order_mark_aggregation_codes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q orderMarkAggregationCodeQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q orderMarkAggregationCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("znakboil: no orderMarkAggregationCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from order_mark_aggregation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_aggregation_codes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OrderMarkAggregationCodeSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderMarkAggregationCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderMarkAggregationCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkAggregationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order_mark_aggregation_codes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkAggregationCodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from orderMarkAggregationCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_aggregation_codes")
	}

	if len(orderMarkAggregationCodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OrderMarkAggregationCode) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: no OrderMarkAggregationCode provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrderMarkAggregationCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderMarkAggregationCode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkAggregationCodeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: empty OrderMarkAggregationCodeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkAggregationCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderMarkAggregationCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkAggregationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order_mark_aggregation_codes\".* FROM \"order_mark_aggregation_codes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkAggregationCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "znakboil: unable to reload all in OrderMarkAggregationCodeSlice")
	}

	*o = slice

	return nil
}

// OrderMarkAggregationCodeExistsG checks if the OrderMarkAggregationCode row exists.
func OrderMarkAggregationCodeExistsG(ctx context.Context, iD int64) (bool, error) {
	return OrderMarkAggregationCodeExists(ctx, boil.GetContextDB(), iD)
}

// OrderMarkAggregationCodeExists checks if the OrderMarkAggregationCode row exists.
func OrderMarkAggregationCodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order_mark_aggregation_codes\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: unable to check if order_mark_aggregation_codes exists")
	}

	return exists, nil
}

// Exists checks if the OrderMarkAggregationCode row exists.
func (o *OrderMarkAggregationCode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderMarkAggregationCodeExists(ctx, exec, o.ID)
}