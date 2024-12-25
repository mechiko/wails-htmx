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

// OrderMarkConnectingTapCode is an object representing the database table.
type OrderMarkConnectingTapCode struct {
	ID                       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDOrderMarkConnectingTap null.Int    `boil:"id_order_mark_connecting_tap" json:"id_order_mark_connecting_tap,omitempty" toml:"id_order_mark_connecting_tap" yaml:"id_order_mark_connecting_tap,omitempty"`
	SerialNumber             null.String `boil:"serial_number" json:"serial_number,omitempty" toml:"serial_number" yaml:"serial_number,omitempty"`
	Code                     null.String `boil:"code" json:"code,omitempty" toml:"code" yaml:"code,omitempty"`
	ConnectDate              null.String `boil:"connect_date" json:"connect_date,omitempty" toml:"connect_date" yaml:"connect_date,omitempty"`
	ExpirationDate           null.String `boil:"expiration_date" json:"expiration_date,omitempty" toml:"expiration_date" yaml:"expiration_date,omitempty"`
	Status                   null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`

	R *orderMarkConnectingTapCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderMarkConnectingTapCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderMarkConnectingTapCodeColumns = struct {
	ID                       string
	IDOrderMarkConnectingTap string
	SerialNumber             string
	Code                     string
	ConnectDate              string
	ExpirationDate           string
	Status                   string
}{
	ID:                       "id",
	IDOrderMarkConnectingTap: "id_order_mark_connecting_tap",
	SerialNumber:             "serial_number",
	Code:                     "code",
	ConnectDate:              "connect_date",
	ExpirationDate:           "expiration_date",
	Status:                   "status",
}

var OrderMarkConnectingTapCodeTableColumns = struct {
	ID                       string
	IDOrderMarkConnectingTap string
	SerialNumber             string
	Code                     string
	ConnectDate              string
	ExpirationDate           string
	Status                   string
}{
	ID:                       "order_mark_connecting_tap_codes.id",
	IDOrderMarkConnectingTap: "order_mark_connecting_tap_codes.id_order_mark_connecting_tap",
	SerialNumber:             "order_mark_connecting_tap_codes.serial_number",
	Code:                     "order_mark_connecting_tap_codes.code",
	ConnectDate:              "order_mark_connecting_tap_codes.connect_date",
	ExpirationDate:           "order_mark_connecting_tap_codes.expiration_date",
	Status:                   "order_mark_connecting_tap_codes.status",
}

// Generated where

var OrderMarkConnectingTapCodeWhere = struct {
	ID                       whereHelperint
	IDOrderMarkConnectingTap whereHelpernull_Int
	SerialNumber             whereHelpernull_String
	Code                     whereHelpernull_String
	ConnectDate              whereHelpernull_String
	ExpirationDate           whereHelpernull_String
	Status                   whereHelpernull_String
}{
	ID:                       whereHelperint{field: "[dbo].[order_mark_connecting_tap_codes].[id]"},
	IDOrderMarkConnectingTap: whereHelpernull_Int{field: "[dbo].[order_mark_connecting_tap_codes].[id_order_mark_connecting_tap]"},
	SerialNumber:             whereHelpernull_String{field: "[dbo].[order_mark_connecting_tap_codes].[serial_number]"},
	Code:                     whereHelpernull_String{field: "[dbo].[order_mark_connecting_tap_codes].[code]"},
	ConnectDate:              whereHelpernull_String{field: "[dbo].[order_mark_connecting_tap_codes].[connect_date]"},
	ExpirationDate:           whereHelpernull_String{field: "[dbo].[order_mark_connecting_tap_codes].[expiration_date]"},
	Status:                   whereHelpernull_String{field: "[dbo].[order_mark_connecting_tap_codes].[status]"},
}

// OrderMarkConnectingTapCodeRels is where relationship names are stored.
var OrderMarkConnectingTapCodeRels = struct {
}{}

// orderMarkConnectingTapCodeR is where relationships are stored.
type orderMarkConnectingTapCodeR struct {
}

// NewStruct creates a new relationship struct
func (*orderMarkConnectingTapCodeR) NewStruct() *orderMarkConnectingTapCodeR {
	return &orderMarkConnectingTapCodeR{}
}

// orderMarkConnectingTapCodeL is where Load methods for each relationship are stored.
type orderMarkConnectingTapCodeL struct{}

var (
	orderMarkConnectingTapCodeAllColumns            = []string{"id", "id_order_mark_connecting_tap", "serial_number", "code", "connect_date", "expiration_date", "status"}
	orderMarkConnectingTapCodeColumnsWithoutDefault = []string{"id_order_mark_connecting_tap", "serial_number", "code", "connect_date", "expiration_date", "status"}
	orderMarkConnectingTapCodeColumnsWithDefault    = []string{"id"}
	orderMarkConnectingTapCodePrimaryKeyColumns     = []string{"id"}
	orderMarkConnectingTapCodeGeneratedColumns      = []string{"id"}
)

type (
	// OrderMarkConnectingTapCodeSlice is an alias for a slice of pointers to OrderMarkConnectingTapCode.
	// This should almost always be used instead of []OrderMarkConnectingTapCode.
	OrderMarkConnectingTapCodeSlice []*OrderMarkConnectingTapCode
	// OrderMarkConnectingTapCodeHook is the signature for custom OrderMarkConnectingTapCode hook methods
	OrderMarkConnectingTapCodeHook func(context.Context, boil.ContextExecutor, *OrderMarkConnectingTapCode) error

	orderMarkConnectingTapCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderMarkConnectingTapCodeType                 = reflect.TypeOf(&OrderMarkConnectingTapCode{})
	orderMarkConnectingTapCodeMapping              = queries.MakeStructMapping(orderMarkConnectingTapCodeType)
	orderMarkConnectingTapCodePrimaryKeyMapping, _ = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, orderMarkConnectingTapCodePrimaryKeyColumns)
	orderMarkConnectingTapCodeInsertCacheMut       sync.RWMutex
	orderMarkConnectingTapCodeInsertCache          = make(map[string]insertCache)
	orderMarkConnectingTapCodeUpdateCacheMut       sync.RWMutex
	orderMarkConnectingTapCodeUpdateCache          = make(map[string]updateCache)
	orderMarkConnectingTapCodeUpsertCacheMut       sync.RWMutex
	orderMarkConnectingTapCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orderMarkConnectingTapCodeAfterSelectMu sync.Mutex
var orderMarkConnectingTapCodeAfterSelectHooks []OrderMarkConnectingTapCodeHook

var orderMarkConnectingTapCodeBeforeInsertMu sync.Mutex
var orderMarkConnectingTapCodeBeforeInsertHooks []OrderMarkConnectingTapCodeHook
var orderMarkConnectingTapCodeAfterInsertMu sync.Mutex
var orderMarkConnectingTapCodeAfterInsertHooks []OrderMarkConnectingTapCodeHook

var orderMarkConnectingTapCodeBeforeUpdateMu sync.Mutex
var orderMarkConnectingTapCodeBeforeUpdateHooks []OrderMarkConnectingTapCodeHook
var orderMarkConnectingTapCodeAfterUpdateMu sync.Mutex
var orderMarkConnectingTapCodeAfterUpdateHooks []OrderMarkConnectingTapCodeHook

var orderMarkConnectingTapCodeBeforeDeleteMu sync.Mutex
var orderMarkConnectingTapCodeBeforeDeleteHooks []OrderMarkConnectingTapCodeHook
var orderMarkConnectingTapCodeAfterDeleteMu sync.Mutex
var orderMarkConnectingTapCodeAfterDeleteHooks []OrderMarkConnectingTapCodeHook

var orderMarkConnectingTapCodeBeforeUpsertMu sync.Mutex
var orderMarkConnectingTapCodeBeforeUpsertHooks []OrderMarkConnectingTapCodeHook
var orderMarkConnectingTapCodeAfterUpsertMu sync.Mutex
var orderMarkConnectingTapCodeAfterUpsertHooks []OrderMarkConnectingTapCodeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderMarkConnectingTapCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderMarkConnectingTapCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderMarkConnectingTapCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrderMarkConnectingTapCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrderMarkConnectingTapCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrderMarkConnectingTapCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrderMarkConnectingTapCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderMarkConnectingTapCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderMarkConnectingTapCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderMarkConnectingTapCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderMarkConnectingTapCodeHook registers your hook function for all future operations.
func AddOrderMarkConnectingTapCodeHook(hookPoint boil.HookPoint, orderMarkConnectingTapCodeHook OrderMarkConnectingTapCodeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderMarkConnectingTapCodeAfterSelectMu.Lock()
		orderMarkConnectingTapCodeAfterSelectHooks = append(orderMarkConnectingTapCodeAfterSelectHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderMarkConnectingTapCodeBeforeInsertMu.Lock()
		orderMarkConnectingTapCodeBeforeInsertHooks = append(orderMarkConnectingTapCodeBeforeInsertHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderMarkConnectingTapCodeAfterInsertMu.Lock()
		orderMarkConnectingTapCodeAfterInsertHooks = append(orderMarkConnectingTapCodeAfterInsertHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		orderMarkConnectingTapCodeBeforeUpdateMu.Lock()
		orderMarkConnectingTapCodeBeforeUpdateHooks = append(orderMarkConnectingTapCodeBeforeUpdateHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		orderMarkConnectingTapCodeAfterUpdateMu.Lock()
		orderMarkConnectingTapCodeAfterUpdateHooks = append(orderMarkConnectingTapCodeAfterUpdateHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		orderMarkConnectingTapCodeBeforeDeleteMu.Lock()
		orderMarkConnectingTapCodeBeforeDeleteHooks = append(orderMarkConnectingTapCodeBeforeDeleteHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		orderMarkConnectingTapCodeAfterDeleteMu.Lock()
		orderMarkConnectingTapCodeAfterDeleteHooks = append(orderMarkConnectingTapCodeAfterDeleteHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		orderMarkConnectingTapCodeBeforeUpsertMu.Lock()
		orderMarkConnectingTapCodeBeforeUpsertHooks = append(orderMarkConnectingTapCodeBeforeUpsertHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderMarkConnectingTapCodeAfterUpsertMu.Lock()
		orderMarkConnectingTapCodeAfterUpsertHooks = append(orderMarkConnectingTapCodeAfterUpsertHooks, orderMarkConnectingTapCodeHook)
		orderMarkConnectingTapCodeAfterUpsertMu.Unlock()
	}
}

// OneG returns a single orderMarkConnectingTapCode record from the query using the global executor.
func (q orderMarkConnectingTapCodeQuery) OneG(ctx context.Context) (*OrderMarkConnectingTapCode, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single orderMarkConnectingTapCode record from the query.
func (q orderMarkConnectingTapCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderMarkConnectingTapCode, error) {
	o := &OrderMarkConnectingTapCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: failed to execute a one query for order_mark_connecting_tap_codes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all OrderMarkConnectingTapCode records from the query using the global executor.
func (q orderMarkConnectingTapCodeQuery) AllG(ctx context.Context) (OrderMarkConnectingTapCodeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all OrderMarkConnectingTapCode records from the query.
func (q orderMarkConnectingTapCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderMarkConnectingTapCodeSlice, error) {
	var o []*OrderMarkConnectingTapCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "znakboil: failed to assign all query results to OrderMarkConnectingTapCode slice")
	}

	if len(orderMarkConnectingTapCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all OrderMarkConnectingTapCode records in the query using the global executor
func (q orderMarkConnectingTapCodeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all OrderMarkConnectingTapCode records in the query.
func (q orderMarkConnectingTapCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to count order_mark_connecting_tap_codes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q orderMarkConnectingTapCodeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q orderMarkConnectingTapCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: failed to check if order_mark_connecting_tap_codes exists")
	}

	return count > 0, nil
}

// OrderMarkConnectingTapCodes retrieves all the records using an executor.
func OrderMarkConnectingTapCodes(mods ...qm.QueryMod) orderMarkConnectingTapCodeQuery {
	mods = append(mods, qm.From("[dbo].[order_mark_connecting_tap_codes]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[order_mark_connecting_tap_codes].*"})
	}

	return orderMarkConnectingTapCodeQuery{q}
}

// FindOrderMarkConnectingTapCodeG retrieves a single record by ID.
func FindOrderMarkConnectingTapCodeG(ctx context.Context, iD int, selectCols ...string) (*OrderMarkConnectingTapCode, error) {
	return FindOrderMarkConnectingTapCode(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOrderMarkConnectingTapCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderMarkConnectingTapCode(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrderMarkConnectingTapCode, error) {
	orderMarkConnectingTapCodeObj := &OrderMarkConnectingTapCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[order_mark_connecting_tap_codes] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderMarkConnectingTapCodeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: unable to select from order_mark_connecting_tap_codes")
	}

	if err = orderMarkConnectingTapCodeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orderMarkConnectingTapCodeObj, err
	}

	return orderMarkConnectingTapCodeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OrderMarkConnectingTapCode) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderMarkConnectingTapCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_connecting_tap_codes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkConnectingTapCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderMarkConnectingTapCodeInsertCacheMut.RLock()
	cache, cached := orderMarkConnectingTapCodeInsertCache[key]
	orderMarkConnectingTapCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderMarkConnectingTapCodeAllColumns,
			orderMarkConnectingTapCodeColumnsWithDefault,
			orderMarkConnectingTapCodeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, orderMarkConnectingTapCodeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[order_mark_connecting_tap_codes] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[order_mark_connecting_tap_codes] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "znakboil: unable to insert into order_mark_connecting_tap_codes")
	}

	if !cached {
		orderMarkConnectingTapCodeInsertCacheMut.Lock()
		orderMarkConnectingTapCodeInsertCache[key] = cache
		orderMarkConnectingTapCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single OrderMarkConnectingTapCode record using the global executor.
// See Update for more documentation.
func (o *OrderMarkConnectingTapCode) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the OrderMarkConnectingTapCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderMarkConnectingTapCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orderMarkConnectingTapCodeUpdateCacheMut.RLock()
	cache, cached := orderMarkConnectingTapCodeUpdateCache[key]
	orderMarkConnectingTapCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderMarkConnectingTapCodeAllColumns,
			orderMarkConnectingTapCodePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, orderMarkConnectingTapCodeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("znakboil: unable to update order_mark_connecting_tap_codes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[order_mark_connecting_tap_codes] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, orderMarkConnectingTapCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, append(wl, orderMarkConnectingTapCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "znakboil: unable to update order_mark_connecting_tap_codes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by update for order_mark_connecting_tap_codes")
	}

	if !cached {
		orderMarkConnectingTapCodeUpdateCacheMut.Lock()
		orderMarkConnectingTapCodeUpdateCache[key] = cache
		orderMarkConnectingTapCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q orderMarkConnectingTapCodeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q orderMarkConnectingTapCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all for order_mark_connecting_tap_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected for order_mark_connecting_tap_codes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderMarkConnectingTapCodeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderMarkConnectingTapCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkConnectingTapCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[order_mark_connecting_tap_codes] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderMarkConnectingTapCodePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all in orderMarkConnectingTapCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected all in update all orderMarkConnectingTapCode")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OrderMarkConnectingTapCode) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OrderMarkConnectingTapCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no order_mark_connecting_tap_codes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderMarkConnectingTapCodeColumnsWithDefault, o)

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

	orderMarkConnectingTapCodeUpsertCacheMut.RLock()
	cache, cached := orderMarkConnectingTapCodeUpsertCache[key]
	orderMarkConnectingTapCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			orderMarkConnectingTapCodeAllColumns,
			orderMarkConnectingTapCodeColumnsWithDefault,
			orderMarkConnectingTapCodeColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, orderMarkConnectingTapCodeGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(orderMarkConnectingTapCodePrimaryKeyColumns, v) && strmangle.ContainsAny(orderMarkConnectingTapCodeColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("znakboil: unable to upsert order_mark_connecting_tap_codes, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			orderMarkConnectingTapCodeAllColumns,
			orderMarkConnectingTapCodePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, orderMarkConnectingTapCodeGeneratedColumns)

		ret := strmangle.SetComplement(orderMarkConnectingTapCodeAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("znakboil: unable to upsert order_mark_connecting_tap_codes, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[order_mark_connecting_tap_codes]", orderMarkConnectingTapCodePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(orderMarkConnectingTapCodePrimaryKeyColumns))
		copy(whitelist, orderMarkConnectingTapCodePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderMarkConnectingTapCodeType, orderMarkConnectingTapCodeMapping, ret)
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
		return errors.Wrap(err, "znakboil: unable to upsert order_mark_connecting_tap_codes")
	}

	if !cached {
		orderMarkConnectingTapCodeUpsertCacheMut.Lock()
		orderMarkConnectingTapCodeUpsertCache[key] = cache
		orderMarkConnectingTapCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single OrderMarkConnectingTapCode record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OrderMarkConnectingTapCode) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single OrderMarkConnectingTapCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderMarkConnectingTapCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("znakboil: no OrderMarkConnectingTapCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderMarkConnectingTapCodePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[order_mark_connecting_tap_codes] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete from order_mark_connecting_tap_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by delete for order_mark_connecting_tap_codes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q orderMarkConnectingTapCodeQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q orderMarkConnectingTapCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("znakboil: no orderMarkConnectingTapCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from order_mark_connecting_tap_codes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_connecting_tap_codes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OrderMarkConnectingTapCodeSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderMarkConnectingTapCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orderMarkConnectingTapCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkConnectingTapCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[order_mark_connecting_tap_codes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderMarkConnectingTapCodePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from orderMarkConnectingTapCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for order_mark_connecting_tap_codes")
	}

	if len(orderMarkConnectingTapCodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OrderMarkConnectingTapCode) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: no OrderMarkConnectingTapCode provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrderMarkConnectingTapCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderMarkConnectingTapCode(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkConnectingTapCodeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: empty OrderMarkConnectingTapCodeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderMarkConnectingTapCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderMarkConnectingTapCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderMarkConnectingTapCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[order_mark_connecting_tap_codes].* FROM [dbo].[order_mark_connecting_tap_codes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderMarkConnectingTapCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "znakboil: unable to reload all in OrderMarkConnectingTapCodeSlice")
	}

	*o = slice

	return nil
}

// OrderMarkConnectingTapCodeExistsG checks if the OrderMarkConnectingTapCode row exists.
func OrderMarkConnectingTapCodeExistsG(ctx context.Context, iD int) (bool, error) {
	return OrderMarkConnectingTapCodeExists(ctx, boil.GetContextDB(), iD)
}

// OrderMarkConnectingTapCodeExists checks if the OrderMarkConnectingTapCode row exists.
func OrderMarkConnectingTapCodeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[order_mark_connecting_tap_codes] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: unable to check if order_mark_connecting_tap_codes exists")
	}

	return exists, nil
}

// Exists checks if the OrderMarkConnectingTapCode row exists.
func (o *OrderMarkConnectingTapCode) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrderMarkConnectingTapCodeExists(ctx, exec, o.ID)
}
