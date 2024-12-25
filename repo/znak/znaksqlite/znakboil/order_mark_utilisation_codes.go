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

// OrderMarkUtilisationCode is an object representing the database table.
type OrderMarkUtilisationCode struct {
	ID                     int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDOrderMarkUtilisation null.Int64  `boil:"id_order_mark_utilisation" json:"id_order_mark_utilisation,omitempty" toml:"id_order_mark_utilisation" yaml:"id_order_mark_utilisation,omitempty"`
	SerialNumber           null.String `boil:"serial_number" json:"serial_number,omitempty" toml:"serial_number" yaml:"serial_number,omitempty"`
	Code                   null.String `boil:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	Status                 null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`

	R *orderMarkUtilisationCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderMarkUtilisationCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderMarkUtilisationCodeColumns = struct {
	ID                     string
	IDOrderMarkUtilisation string
	SerialNumber           string
	Code                   string
	Status                 string
}{
	ID:                     "id",
	IDOrderMarkUtilisation: "id_order_mark_utilisation",
	SerialNumber:           "serial_number",
	Code:                   "code",
	Status:                 "status",
}

var OrderMarkUtilisationCodeTableColumns = struct {
	ID                     string
	IDOrderMarkUtilisation string
	SerialNumber           string
	Code                   string
	Status                 string
}{
	ID:                     "order_mark_utilisation_codes.id",
	IDOrderMarkUtilisation: "order_mark_utilisation_codes.id_order_mark_utilisation",
	SerialNumber:           "order_mark_utilisation_codes.serial_number",
	Code:                   "order_mark_utilisation_codes.code",
	Status:                 "order_mark_utilisation_codes.status",
}

// Generated where

var OrderMarkUtilisationCodeWhere = struct {
	ID                     whereHelperint64
	IDOrderMarkUtilisation whereHelpernull_Int64
	SerialNumber           whereHelpernull_String
	Code                   whereHelpernull_String
	Status                 whereHelpernull_String
}{
	ID:                     whereHelperint64{field: "\"order_mark_utilisation_codes\".\"id\""},
	IDOrderMarkUtilisation: whereHelpernull_Int64{field: "\"order_mark_utilisation_codes\".\"id_order_mark_utilisation\""},
	SerialNumber:           whereHelpernull_String{field: "\"order_mark_utilisation_codes\".\"serial_number\""},
	Code:                   whereHelpernull_String{field: "\"order_mark_utilisation_codes\".\"code\""},
	Status:                 whereHelpernull_String{field: "\"order_mark_utilisation_codes\".\"status\""},
}

// OrderMarkUtilisationCodeRels is where relationship names are stored.
var OrderMarkUtilisationCodeRels = struct {
}{}

// orderMarkUtilisationCodeR is where relationships are stored.
type orderMarkUtilisationCodeR struct {
}

// NewStruct creates a new relationship struct
func (*orderMarkUtilisationCodeR) NewStruct() *orderMarkUtilisationCodeR {
	return &orderMarkUtilisationCodeR{}
}

// orderMarkUtilisationCodeL is where Load methods for each relationship are stored.
type orderMarkUtilisationCodeL struct{}

var (
	orderMarkUtilisationCodeAllColumns            = []string{"id", "id_order_mark_utilisation", "serial_number", "code", "status"}
	orderMarkUtilisationCodeColumnsWithoutDefault = []string{}
	orderMarkUtilisationCodeColumnsWithDefault    = []string{"id", "id_order_mark_utilisation", "serial_number", "code", "status"}
	orderMarkUtilisationCodePrimaryKeyColumns     = []string{"id"}
	orderMarkUtilisationCodeGeneratedColumns      = []string{"id"}
)

type (
	// OrderMarkUtilisationCodeSlice is an alias for a slice of pointers to OrderMarkUtilisationCode.
	// This should almost always be used instead of []OrderMarkUtilisationCode.
	OrderMarkUtilisationCodeSlice []*OrderMarkUtilisationCode
	// OrderMarkUtilisationCodeHook is the signature for custom OrderMarkUtilisationCode hook methods
	OrderMarkUtilisationCodeHook func(context.Context, boil.ContextExecutor, *OrderMarkUtilisationCode) error

	orderMarkUtilisationCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderMarkUtilisationCodeType                 = reflect.TypeOf(&OrderMarkUtilisationCode{})
	orderMarkUtilisationCodeMapping              = queries.MakeStructMapping(orderMarkUtilisationCodeType)
	orderMarkUtilisationCodePrimaryKeyMapping, _ = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, orderMarkUtilisationCodePrimaryKeyColumns)
	orderMarkUtilisationCodeInsertCacheMut       sync.RWMutex
	orderMarkUtilisationCodeInsertCache          = make(map[string]insertCache)
	orderMarkUtilisationCodeUpdateCacheMut       sync.RWMutex
	orderMarkUtilisationCodeUpdateCache          = make(map[string]updateCache)
	orderMarkUtilisationCodeUpsertCacheMut       sync.RWMutex
	orderMarkUtilisationCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderMarkUtilisationCodeAfterSelectMu sync.Mutex
var orderMarkUtilisationCodeAfterSelectHooks []OrderMarkUtilisationCodeHook

var orderMarkUtilisationCodeBeforeInsertMu sync.Mutex
var orderMarkUtilisationCodeBeforeInsertHooks []OrderMarkUtilisationCodeHook
var orderMarkUtilisationCodeAfterInsertMu sync.Mutex
var orderMarkUtilisationCodeAfterInsertHooks []OrderMarkUtilisationCodeHook

var orderMarkUtilisationCodeBeforeUpdateMu sync.Mutex
var orderMarkUtilisationCodeBeforeUpdateHooks []OrderMarkUtilisationCodeHook
var orderMarkUtilisationCodeAfterUpdateMu sync.Mutex
var orderMarkUtilisationCodeAfterUpdateHooks []OrderMarkUtilisationCodeHook

var orderMarkUtilisationCodeBeforeDeleteMu sync.Mutex
var orderMarkUtilisationCodeBeforeDeleteHooks []OrderMarkUtilisationCodeHook
var orderMarkUtilisationCodeAfterDeleteMu sync.Mutex
var orderMarkUtilisationCodeAfterDeleteHooks []OrderMarkUtilisationCodeHook

var orderMarkUtilisationCodeBeforeUpsertMu sync.Mutex
var orderMarkUtilisationCodeBeforeUpsertHooks []OrderMarkUtilisationCodeHook
var orderMarkUtilisationCodeAfterUpsertMu sync.Mutex
var orderMarkUtilisationCodeAfterUpsertHooks []OrderMarkUtilisationCodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderMarkUtilisationCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderMarkUtilisationCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderMarkUtilisationCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderMarkUtilisationCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderMarkUtilisationCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderMarkUtilisationCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderMarkUtilisationCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderMarkUtilisationCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderMarkUtilisationCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkUtilisationCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderMarkUtilisationCodeHook registers your hook function for all future operations.
func AddOrderMarkUtilisationCodeHook(hookPoint boil.HookPoint, orderMarkUtilisationCodeHook OrderMarkUtilisationCodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderMarkUtilisationCodeAfterSelectMu.Lock()
		orderMarkUtilisationCodeAfterSelectHooks = append(orderMarkUtilisationCodeAfterSelectHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderMarkUtilisationCodeBeforeInsertMu.Lock()
		orderMarkUtilisationCodeBeforeInsertHooks = append(orderMarkUtilisationCodeBeforeInsertHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderMarkUtilisationCodeAfterInsertMu.Lock()
		orderMarkUtilisationCodeAfterInsertHooks = append(orderMarkUtilisationCodeAfterInsertHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		orderMarkUtilisationCodeBeforeUpdateMu.Lock()
		orderMarkUtilisationCodeBeforeUpdateHooks = append(orderMarkUtilisationCodeBeforeUpdateHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		orderMarkUtilisationCodeAfterUpdateMu.Lock()
		orderMarkUtilisationCodeAfterUpdateHooks = append(orderMarkUtilisationCodeAfterUpdateHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		orderMarkUtilisationCodeBeforeDeleteMu.Lock()
		orderMarkUtilisationCodeBeforeDeleteHooks = append(orderMarkUtilisationCodeBeforeDeleteHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		orderMarkUtilisationCodeAfterDeleteMu.Lock()
		orderMarkUtilisationCodeAfterDeleteHooks = append(orderMarkUtilisationCodeAfterDeleteHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		orderMarkUtilisationCodeBeforeUpsertMu.Lock()
		orderMarkUtilisationCodeBeforeUpsertHooks = append(orderMarkUtilisationCodeBeforeUpsertHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderMarkUtilisationCodeAfterUpsertMu.Lock()
		orderMarkUtilisationCodeAfterUpsertHooks = append(orderMarkUtilisationCodeAfterUpsertHooks, orderMarkUtilisationCodeHook)
		orderMarkUtilisationCodeAfterUpsertMu.Unlock()
	}
}

// OneG returns a single orderMarkUtilisationCode record from the query using the global executor.
func (q orderMarkUtilisationCodeQuery) OneG(ctx context.Context) (*OrderMarkUtilisationCode, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single orderMarkUtilisationCode record from the query.
func (q orderMarkUtilisationCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderMarkUtilisationCode, error) {
	o := &OrderMarkUtilisationCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: failed to execute a one query for order_mark_utilisation_codes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all OrderMarkUtilisationCode records from the query using the global executor.
func (q orderMarkUtilisationCodeQuery) AllG(ctx context.Context) (OrderMarkUtilisationCodeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all OrderMarkUtilisationCode records from the query.
func (q orderMarkUtilisationCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderMarkUtilisationCodeSlice, error) {
	var o []*OrderMarkUtilisationCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "znakboil: failed to assign all query results to OrderMarkUtilisationCode slice")
	}

	if len(orderMarkUtilisationCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all OrderMarkUtilisationCode records in the query using the global executor
func (q orderMarkUtilisationCodeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all OrderMarkUtilisationCode records in the query.
func (q orderMarkUtilisationCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to count order_mark_utilisation_codes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q orderMarkUtilisationCodeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q orderMarkUtilisationCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: failed to check if order_mark_utilisation_codes exists")
	}

	return count > 0, nil
}

// OrderMarkUtilisationCodes retrieves all the records using an executor.
func OrderMarkUtilisationCodes(mods ...qm.QueryMod) orderMarkUtilisationCodeQuery {
	mods = append(mods, qm.From("\"order_mark_utilisation_codes\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"order_mark_utilisation_codes\".*"})
	}

	return orderMarkUtilisationCodeQuery{q}
}

// FindOrderMarkUtilisationCodeG retrieves a single record by ID.
func FindOrderMarkUtilisationCodeG(ctx context.Context, iD int64, selectCols ...string) (*OrderMarkUtilisationCode, error) {
	return FindOrderMarkUtilisationCode(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOrderMarkUtilisationCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderMarkUtilisationCode(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OrderMarkUtilisationCode, error) {
	orderMarkUtilisationCodeObj := &OrderMarkUtilisationCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order_mark_utilisation_codes\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderMarkUtilisationCodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: unable to select from order_mark_utilisation_codes")
	}

	if err = orderMarkUtilisationCodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderMarkUtilisationCodeObj, err
	}

	return orderMarkUtilisationCodeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OrderMarkUtilisationCode) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderMarkUtilisationCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_utilisation_codes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkUtilisationCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderMarkUtilisationCodeInsertCacheMut.RLock()
	cache, cached := orderMarkUtilisationCodeInsertCache[key]
	orderMarkUtilisationCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderMarkUtilisationCodeAllColumns,
			orderMarkUtilisationCodeColumnsWithDefault,
			orderMarkUtilisationCodeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, orderMarkUtilisationCodeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order_mark_utilisation_codes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order_mark_utilisation_codes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "znakboil: unable to insert into order_mark_utilisation_codes")
	}

	if !cached {
		orderMarkUtilisationCodeInsertCacheMut.Lock()
		orderMarkUtilisationCodeInsertCache[key] = cache
		orderMarkUtilisationCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single OrderMarkUtilisationCode record using the global executor.
// See Update for more documentation.
func (o *OrderMarkUtilisationCode) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the OrderMarkUtilisationCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderMarkUtilisationCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderMarkUtilisationCodeUpdateCacheMut.RLock()
	cache, cached := orderMarkUtilisationCodeUpdateCache[key]
	orderMarkUtilisationCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderMarkUtilisationCodeAllColumns,
			orderMarkUtilisationCodePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, orderMarkUtilisationCodeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("znakboil: unable to update order_mark_utilisation_codes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order_mark_utilisation_codes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, orderMarkUtilisationCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, append(wl, orderMarkUtilisationCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "znakboil: unable to update order_mark_utilisation_codes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by update for order_mark_utilisation_codes")
	}

	if !cached {
		orderMarkUtilisationCodeUpdateCacheMut.Lock()
		orderMarkUtilisationCodeUpdateCache[key] = cache
		orderMarkUtilisationCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q orderMarkUtilisationCodeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q orderMarkUtilisationCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all for order_mark_utilisation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected for order_mark_utilisation_codes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderMarkUtilisationCodeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderMarkUtilisationCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkUtilisationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order_mark_utilisation_codes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkUtilisationCodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all in orderMarkUtilisationCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected all in update all orderMarkUtilisationCode")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OrderMarkUtilisationCode) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderMarkUtilisationCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_utilisation_codes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkUtilisationCodeColumnsWithDefault, o)

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

	orderMarkUtilisationCodeUpsertCacheMut.RLock()
	cache, cached := orderMarkUtilisationCodeUpsertCache[key]
	orderMarkUtilisationCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderMarkUtilisationCodeAllColumns,
			orderMarkUtilisationCodeColumnsWithDefault,
			orderMarkUtilisationCodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orderMarkUtilisationCodeAllColumns,
			orderMarkUtilisationCodePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("znakboil: unable to upsert order_mark_utilisation_codes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderMarkUtilisationCodePrimaryKeyColumns))
			copy(conflict, orderMarkUtilisationCodePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"order_mark_utilisation_codes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderMarkUtilisationCodeType, orderMarkUtilisationCodeMapping, ret)
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
		return errors.Wrap(err, "znakboil: unable to upsert order_mark_utilisation_codes")
	}

	if !cached {
		orderMarkUtilisationCodeUpsertCacheMut.Lock()
		orderMarkUtilisationCodeUpsertCache[key] = cache
		orderMarkUtilisationCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single OrderMarkUtilisationCode record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OrderMarkUtilisationCode) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single OrderMarkUtilisationCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderMarkUtilisationCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("znakboil: no OrderMarkUtilisationCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderMarkUtilisationCodePrimaryKeyMapping)
	sql := "DELETE FROM \"order_mark_utilisation_codes\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete from order_mark_utilisation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by delete for order_mark_utilisation_codes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q orderMarkUtilisationCodeQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q orderMarkUtilisationCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("znakboil: no orderMarkUtilisationCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from order_mark_utilisation_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_utilisation_codes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OrderMarkUtilisationCodeSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderMarkUtilisationCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderMarkUtilisationCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkUtilisationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order_mark_utilisation_codes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkUtilisationCodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from orderMarkUtilisationCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_utilisation_codes")
	}

	if len(orderMarkUtilisationCodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OrderMarkUtilisationCode) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: no OrderMarkUtilisationCode provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrderMarkUtilisationCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderMarkUtilisationCode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkUtilisationCodeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: empty OrderMarkUtilisationCodeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkUtilisationCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderMarkUtilisationCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkUtilisationCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order_mark_utilisation_codes\".* FROM \"order_mark_utilisation_codes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderMarkUtilisationCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "znakboil: unable to reload all in OrderMarkUtilisationCodeSlice")
	}

	*o = slice

	return nil
}

// OrderMarkUtilisationCodeExistsG checks if the OrderMarkUtilisationCode row exists.
func OrderMarkUtilisationCodeExistsG(ctx context.Context, iD int64) (bool, error) {
	return OrderMarkUtilisationCodeExists(ctx, boil.GetContextDB(), iD)
}

// OrderMarkUtilisationCodeExists checks if the OrderMarkUtilisationCode row exists.
func OrderMarkUtilisationCodeExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order_mark_utilisation_codes\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: unable to check if order_mark_utilisation_codes exists")
	}

	return exists, nil
}

// Exists checks if the OrderMarkUtilisationCode row exists.
func (o *OrderMarkUtilisationCode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderMarkUtilisationCodeExists(ctx, exec, o.ID)
}
