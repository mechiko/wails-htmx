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

// ProductionResourcesPallet is an object representing the database table.
type ProductionResourcesPallet struct {
	ID                    int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDProductionResources null.Int    `boil:"id_production_resources" json:"id_production_resources,omitempty" toml:"id_production_resources" yaml:"id_production_resources,omitempty"`
	PalletNumber          null.String `boil:"pallet_number" json:"pallet_number,omitempty" toml:"pallet_number" yaml:"pallet_number,omitempty"`

	R *productionResourcesPalletR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productionResourcesPalletL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductionResourcesPalletColumns = struct {
	ID                    string
	IDProductionResources string
	PalletNumber          string
}{
	ID:                    "id",
	IDProductionResources: "id_production_resources",
	PalletNumber:          "pallet_number",
}

var ProductionResourcesPalletTableColumns = struct {
	ID                    string
	IDProductionResources string
	PalletNumber          string
}{
	ID:                    "production_resources_pallets.id",
	IDProductionResources: "production_resources_pallets.id_production_resources",
	PalletNumber:          "production_resources_pallets.pallet_number",
}

// Generated where

var ProductionResourcesPalletWhere = struct {
	ID                    whereHelperint
	IDProductionResources whereHelpernull_Int
	PalletNumber          whereHelpernull_String
}{
	ID:                    whereHelperint{field: "[dbo].[production_resources_pallets].[id]"},
	IDProductionResources: whereHelpernull_Int{field: "[dbo].[production_resources_pallets].[id_production_resources]"},
	PalletNumber:          whereHelpernull_String{field: "[dbo].[production_resources_pallets].[pallet_number]"},
}

// ProductionResourcesPalletRels is where relationship names are stored.
var ProductionResourcesPalletRels = struct {
}{}

// productionResourcesPalletR is where relationships are stored.
type productionResourcesPalletR struct {
}

// NewStruct creates a new relationship struct
func (*productionResourcesPalletR) NewStruct() *productionResourcesPalletR {
	return &productionResourcesPalletR{}
}

// productionResourcesPalletL is where Load methods for each relationship are stored.
type productionResourcesPalletL struct{}

var (
	productionResourcesPalletAllColumns            = []string{"id", "id_production_resources", "pallet_number"}
	productionResourcesPalletColumnsWithoutDefault = []string{"id_production_resources", "pallet_number"}
	productionResourcesPalletColumnsWithDefault    = []string{"id"}
	productionResourcesPalletPrimaryKeyColumns     = []string{"id"}
	productionResourcesPalletGeneratedColumns      = []string{"id"}
)

type (
	// ProductionResourcesPalletSlice is an alias for a slice of pointers to ProductionResourcesPallet.
	// This should almost always be used instead of []ProductionResourcesPallet.
	ProductionResourcesPalletSlice []*ProductionResourcesPallet
	// ProductionResourcesPalletHook is the signature for custom ProductionResourcesPallet hook methods
	ProductionResourcesPalletHook func(context.Context, boil.ContextExecutor, *ProductionResourcesPallet) error

	productionResourcesPalletQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productionResourcesPalletType                 = reflect.TypeOf(&ProductionResourcesPallet{})
	productionResourcesPalletMapping              = queries.MakeStructMapping(productionResourcesPalletType)
	productionResourcesPalletPrimaryKeyMapping, _ = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, productionResourcesPalletPrimaryKeyColumns)
	productionResourcesPalletInsertCacheMut       sync.RWMutex
	productionResourcesPalletInsertCache          = make(map[string]insertCache)
	productionResourcesPalletUpdateCacheMut       sync.RWMutex
	productionResourcesPalletUpdateCache          = make(map[string]updateCache)
	productionResourcesPalletUpsertCacheMut       sync.RWMutex
	productionResourcesPalletUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productionResourcesPalletAfterSelectMu sync.Mutex
var productionResourcesPalletAfterSelectHooks []ProductionResourcesPalletHook

var productionResourcesPalletBeforeInsertMu sync.Mutex
var productionResourcesPalletBeforeInsertHooks []ProductionResourcesPalletHook
var productionResourcesPalletAfterInsertMu sync.Mutex
var productionResourcesPalletAfterInsertHooks []ProductionResourcesPalletHook

var productionResourcesPalletBeforeUpdateMu sync.Mutex
var productionResourcesPalletBeforeUpdateHooks []ProductionResourcesPalletHook
var productionResourcesPalletAfterUpdateMu sync.Mutex
var productionResourcesPalletAfterUpdateHooks []ProductionResourcesPalletHook

var productionResourcesPalletBeforeDeleteMu sync.Mutex
var productionResourcesPalletBeforeDeleteHooks []ProductionResourcesPalletHook
var productionResourcesPalletAfterDeleteMu sync.Mutex
var productionResourcesPalletAfterDeleteHooks []ProductionResourcesPalletHook

var productionResourcesPalletBeforeUpsertMu sync.Mutex
var productionResourcesPalletBeforeUpsertHooks []ProductionResourcesPalletHook
var productionResourcesPalletAfterUpsertMu sync.Mutex
var productionResourcesPalletAfterUpsertHooks []ProductionResourcesPalletHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductionResourcesPallet) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductionResourcesPallet) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductionResourcesPallet) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductionResourcesPallet) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductionResourcesPallet) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductionResourcesPallet) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductionResourcesPallet) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductionResourcesPallet) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductionResourcesPallet) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionResourcesPalletAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductionResourcesPalletHook registers your hook function for all future operations.
func AddProductionResourcesPalletHook(hookPoint boil.HookPoint, productionResourcesPalletHook ProductionResourcesPalletHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productionResourcesPalletAfterSelectMu.Lock()
		productionResourcesPalletAfterSelectHooks = append(productionResourcesPalletAfterSelectHooks, productionResourcesPalletHook)
		productionResourcesPalletAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productionResourcesPalletBeforeInsertMu.Lock()
		productionResourcesPalletBeforeInsertHooks = append(productionResourcesPalletBeforeInsertHooks, productionResourcesPalletHook)
		productionResourcesPalletBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productionResourcesPalletAfterInsertMu.Lock()
		productionResourcesPalletAfterInsertHooks = append(productionResourcesPalletAfterInsertHooks, productionResourcesPalletHook)
		productionResourcesPalletAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productionResourcesPalletBeforeUpdateMu.Lock()
		productionResourcesPalletBeforeUpdateHooks = append(productionResourcesPalletBeforeUpdateHooks, productionResourcesPalletHook)
		productionResourcesPalletBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productionResourcesPalletAfterUpdateMu.Lock()
		productionResourcesPalletAfterUpdateHooks = append(productionResourcesPalletAfterUpdateHooks, productionResourcesPalletHook)
		productionResourcesPalletAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productionResourcesPalletBeforeDeleteMu.Lock()
		productionResourcesPalletBeforeDeleteHooks = append(productionResourcesPalletBeforeDeleteHooks, productionResourcesPalletHook)
		productionResourcesPalletBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productionResourcesPalletAfterDeleteMu.Lock()
		productionResourcesPalletAfterDeleteHooks = append(productionResourcesPalletAfterDeleteHooks, productionResourcesPalletHook)
		productionResourcesPalletAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productionResourcesPalletBeforeUpsertMu.Lock()
		productionResourcesPalletBeforeUpsertHooks = append(productionResourcesPalletBeforeUpsertHooks, productionResourcesPalletHook)
		productionResourcesPalletBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productionResourcesPalletAfterUpsertMu.Lock()
		productionResourcesPalletAfterUpsertHooks = append(productionResourcesPalletAfterUpsertHooks, productionResourcesPalletHook)
		productionResourcesPalletAfterUpsertMu.Unlock()
	}
}

// OneG returns a single productionResourcesPallet record from the query using the global executor.
func (q productionResourcesPalletQuery) OneG(ctx context.Context) (*ProductionResourcesPallet, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productionResourcesPallet record from the query.
func (q productionResourcesPalletQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductionResourcesPallet, error) {
	o := &ProductionResourcesPallet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for production_resources_pallets")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductionResourcesPallet records from the query using the global executor.
func (q productionResourcesPalletQuery) AllG(ctx context.Context) (ProductionResourcesPalletSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductionResourcesPallet records from the query.
func (q productionResourcesPalletQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductionResourcesPalletSlice, error) {
	var o []*ProductionResourcesPallet

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ProductionResourcesPallet slice")
	}

	if len(productionResourcesPalletAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductionResourcesPallet records in the query using the global executor
func (q productionResourcesPalletQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductionResourcesPallet records in the query.
func (q productionResourcesPalletQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count production_resources_pallets rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productionResourcesPalletQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productionResourcesPalletQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if production_resources_pallets exists")
	}

	return count > 0, nil
}

// ProductionResourcesPallets retrieves all the records using an executor.
func ProductionResourcesPallets(mods ...qm.QueryMod) productionResourcesPalletQuery {
	mods = append(mods, qm.From("[dbo].[production_resources_pallets]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[production_resources_pallets].*"})
	}

	return productionResourcesPalletQuery{q}
}

// FindProductionResourcesPalletG retrieves a single record by ID.
func FindProductionResourcesPalletG(ctx context.Context, iD int, selectCols ...string) (*ProductionResourcesPallet, error) {
	return FindProductionResourcesPallet(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductionResourcesPallet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductionResourcesPallet(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProductionResourcesPallet, error) {
	productionResourcesPalletObj := &ProductionResourcesPallet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[production_resources_pallets] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productionResourcesPalletObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from production_resources_pallets")
	}

	if err = productionResourcesPalletObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productionResourcesPalletObj, err
	}

	return productionResourcesPalletObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductionResourcesPallet) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductionResourcesPallet) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_resources_pallets provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionResourcesPalletColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productionResourcesPalletInsertCacheMut.RLock()
	cache, cached := productionResourcesPalletInsertCache[key]
	productionResourcesPalletInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productionResourcesPalletAllColumns,
			productionResourcesPalletColumnsWithDefault,
			productionResourcesPalletColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productionResourcesPalletGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[production_resources_pallets] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[production_resources_pallets] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into production_resources_pallets")
	}

	if !cached {
		productionResourcesPalletInsertCacheMut.Lock()
		productionResourcesPalletInsertCache[key] = cache
		productionResourcesPalletInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductionResourcesPallet record using the global executor.
// See Update for more documentation.
func (o *ProductionResourcesPallet) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductionResourcesPallet.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductionResourcesPallet) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productionResourcesPalletUpdateCacheMut.RLock()
	cache, cached := productionResourcesPalletUpdateCache[key]
	productionResourcesPalletUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productionResourcesPalletAllColumns,
			productionResourcesPalletPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productionResourcesPalletGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update production_resources_pallets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[production_resources_pallets] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, productionResourcesPalletPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, append(wl, productionResourcesPalletPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update production_resources_pallets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for production_resources_pallets")
	}

	if !cached {
		productionResourcesPalletUpdateCacheMut.Lock()
		productionResourcesPalletUpdateCache[key] = cache
		productionResourcesPalletUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productionResourcesPalletQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productionResourcesPalletQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for production_resources_pallets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for production_resources_pallets")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductionResourcesPalletSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductionResourcesPalletSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionResourcesPalletPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[production_resources_pallets] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productionResourcesPalletPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in productionResourcesPallet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all productionResourcesPallet")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductionResourcesPallet) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductionResourcesPallet) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_resources_pallets provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionResourcesPalletColumnsWithDefault, o)

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

	productionResourcesPalletUpsertCacheMut.RLock()
	cache, cached := productionResourcesPalletUpsertCache[key]
	productionResourcesPalletUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productionResourcesPalletAllColumns,
			productionResourcesPalletColumnsWithDefault,
			productionResourcesPalletColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, productionResourcesPalletGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(productionResourcesPalletPrimaryKeyColumns, v) && strmangle.ContainsAny(productionResourcesPalletColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert production_resources_pallets, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			productionResourcesPalletAllColumns,
			productionResourcesPalletPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, productionResourcesPalletGeneratedColumns)

		ret := strmangle.SetComplement(productionResourcesPalletAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert production_resources_pallets, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[production_resources_pallets]", productionResourcesPalletPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(productionResourcesPalletPrimaryKeyColumns))
		copy(whitelist, productionResourcesPalletPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productionResourcesPalletType, productionResourcesPalletMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert production_resources_pallets")
	}

	if !cached {
		productionResourcesPalletUpsertCacheMut.Lock()
		productionResourcesPalletUpsertCache[key] = cache
		productionResourcesPalletUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductionResourcesPallet record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductionResourcesPallet) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductionResourcesPallet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductionResourcesPallet) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ProductionResourcesPallet provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productionResourcesPalletPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[production_resources_pallets] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from production_resources_pallets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for production_resources_pallets")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productionResourcesPalletQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productionResourcesPalletQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no productionResourcesPalletQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from production_resources_pallets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_resources_pallets")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductionResourcesPalletSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductionResourcesPalletSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productionResourcesPalletBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionResourcesPalletPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[production_resources_pallets] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionResourcesPalletPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from productionResourcesPallet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_resources_pallets")
	}

	if len(productionResourcesPalletAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductionResourcesPallet) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ProductionResourcesPallet provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductionResourcesPallet) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductionResourcesPallet(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionResourcesPalletSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ProductionResourcesPalletSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionResourcesPalletSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductionResourcesPalletSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionResourcesPalletPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[production_resources_pallets].* FROM [dbo].[production_resources_pallets] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionResourcesPalletPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ProductionResourcesPalletSlice")
	}

	*o = slice

	return nil
}

// ProductionResourcesPalletExistsG checks if the ProductionResourcesPallet row exists.
func ProductionResourcesPalletExistsG(ctx context.Context, iD int) (bool, error) {
	return ProductionResourcesPalletExists(ctx, boil.GetContextDB(), iD)
}

// ProductionResourcesPalletExists checks if the ProductionResourcesPallet row exists.
func ProductionResourcesPalletExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[production_resources_pallets] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if production_resources_pallets exists")
	}

	return exists, nil
}

// Exists checks if the ProductionResourcesPallet row exists.
func (o *ProductionResourcesPallet) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductionResourcesPalletExists(ctx, exec, o.ID)
}
