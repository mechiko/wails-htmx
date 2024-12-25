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

// ProductionProductsMark is an object representing the database table.
type ProductionProductsMark struct {
	ID                        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDProductionProducts      null.Int    `boil:"id_production_products" json:"id_production_products,omitempty" toml:"id_production_products" yaml:"id_production_products,omitempty"`
	IDProductionProductsBoxes null.Int    `boil:"id_production_products_boxes" json:"id_production_products_boxes,omitempty" toml:"id_production_products_boxes" yaml:"id_production_products_boxes,omitempty"`
	Mark                      null.String `boil:"mark" json:"mark,omitempty" toml:"mark" yaml:"mark,omitempty"`

	R *productionProductsMarkR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productionProductsMarkL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductionProductsMarkColumns = struct {
	ID                        string
	IDProductionProducts      string
	IDProductionProductsBoxes string
	Mark                      string
}{
	ID:                        "id",
	IDProductionProducts:      "id_production_products",
	IDProductionProductsBoxes: "id_production_products_boxes",
	Mark:                      "mark",
}

var ProductionProductsMarkTableColumns = struct {
	ID                        string
	IDProductionProducts      string
	IDProductionProductsBoxes string
	Mark                      string
}{
	ID:                        "production_products_marks.id",
	IDProductionProducts:      "production_products_marks.id_production_products",
	IDProductionProductsBoxes: "production_products_marks.id_production_products_boxes",
	Mark:                      "production_products_marks.mark",
}

// Generated where

var ProductionProductsMarkWhere = struct {
	ID                        whereHelperint
	IDProductionProducts      whereHelpernull_Int
	IDProductionProductsBoxes whereHelpernull_Int
	Mark                      whereHelpernull_String
}{
	ID:                        whereHelperint{field: "[dbo].[production_products_marks].[id]"},
	IDProductionProducts:      whereHelpernull_Int{field: "[dbo].[production_products_marks].[id_production_products]"},
	IDProductionProductsBoxes: whereHelpernull_Int{field: "[dbo].[production_products_marks].[id_production_products_boxes]"},
	Mark:                      whereHelpernull_String{field: "[dbo].[production_products_marks].[mark]"},
}

// ProductionProductsMarkRels is where relationship names are stored.
var ProductionProductsMarkRels = struct {
}{}

// productionProductsMarkR is where relationships are stored.
type productionProductsMarkR struct {
}

// NewStruct creates a new relationship struct
func (*productionProductsMarkR) NewStruct() *productionProductsMarkR {
	return &productionProductsMarkR{}
}

// productionProductsMarkL is where Load methods for each relationship are stored.
type productionProductsMarkL struct{}

var (
	productionProductsMarkAllColumns            = []string{"id", "id_production_products", "id_production_products_boxes", "mark"}
	productionProductsMarkColumnsWithoutDefault = []string{"id_production_products", "id_production_products_boxes", "mark"}
	productionProductsMarkColumnsWithDefault    = []string{"id"}
	productionProductsMarkPrimaryKeyColumns     = []string{"id"}
	productionProductsMarkGeneratedColumns      = []string{"id"}
)

type (
	// ProductionProductsMarkSlice is an alias for a slice of pointers to ProductionProductsMark.
	// This should almost always be used instead of []ProductionProductsMark.
	ProductionProductsMarkSlice []*ProductionProductsMark
	// ProductionProductsMarkHook is the signature for custom ProductionProductsMark hook methods
	ProductionProductsMarkHook func(context.Context, boil.ContextExecutor, *ProductionProductsMark) error

	productionProductsMarkQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productionProductsMarkType                 = reflect.TypeOf(&ProductionProductsMark{})
	productionProductsMarkMapping              = queries.MakeStructMapping(productionProductsMarkType)
	productionProductsMarkPrimaryKeyMapping, _ = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, productionProductsMarkPrimaryKeyColumns)
	productionProductsMarkInsertCacheMut       sync.RWMutex
	productionProductsMarkInsertCache          = make(map[string]insertCache)
	productionProductsMarkUpdateCacheMut       sync.RWMutex
	productionProductsMarkUpdateCache          = make(map[string]updateCache)
	productionProductsMarkUpsertCacheMut       sync.RWMutex
	productionProductsMarkUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productionProductsMarkAfterSelectMu sync.Mutex
var productionProductsMarkAfterSelectHooks []ProductionProductsMarkHook

var productionProductsMarkBeforeInsertMu sync.Mutex
var productionProductsMarkBeforeInsertHooks []ProductionProductsMarkHook
var productionProductsMarkAfterInsertMu sync.Mutex
var productionProductsMarkAfterInsertHooks []ProductionProductsMarkHook

var productionProductsMarkBeforeUpdateMu sync.Mutex
var productionProductsMarkBeforeUpdateHooks []ProductionProductsMarkHook
var productionProductsMarkAfterUpdateMu sync.Mutex
var productionProductsMarkAfterUpdateHooks []ProductionProductsMarkHook

var productionProductsMarkBeforeDeleteMu sync.Mutex
var productionProductsMarkBeforeDeleteHooks []ProductionProductsMarkHook
var productionProductsMarkAfterDeleteMu sync.Mutex
var productionProductsMarkAfterDeleteHooks []ProductionProductsMarkHook

var productionProductsMarkBeforeUpsertMu sync.Mutex
var productionProductsMarkBeforeUpsertHooks []ProductionProductsMarkHook
var productionProductsMarkAfterUpsertMu sync.Mutex
var productionProductsMarkAfterUpsertHooks []ProductionProductsMarkHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductionProductsMark) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductionProductsMark) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductionProductsMark) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductionProductsMark) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductionProductsMark) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductionProductsMark) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductionProductsMark) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductionProductsMark) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductionProductsMark) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionProductsMarkAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductionProductsMarkHook registers your hook function for all future operations.
func AddProductionProductsMarkHook(hookPoint boil.HookPoint, productionProductsMarkHook ProductionProductsMarkHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productionProductsMarkAfterSelectMu.Lock()
		productionProductsMarkAfterSelectHooks = append(productionProductsMarkAfterSelectHooks, productionProductsMarkHook)
		productionProductsMarkAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productionProductsMarkBeforeInsertMu.Lock()
		productionProductsMarkBeforeInsertHooks = append(productionProductsMarkBeforeInsertHooks, productionProductsMarkHook)
		productionProductsMarkBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productionProductsMarkAfterInsertMu.Lock()
		productionProductsMarkAfterInsertHooks = append(productionProductsMarkAfterInsertHooks, productionProductsMarkHook)
		productionProductsMarkAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productionProductsMarkBeforeUpdateMu.Lock()
		productionProductsMarkBeforeUpdateHooks = append(productionProductsMarkBeforeUpdateHooks, productionProductsMarkHook)
		productionProductsMarkBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productionProductsMarkAfterUpdateMu.Lock()
		productionProductsMarkAfterUpdateHooks = append(productionProductsMarkAfterUpdateHooks, productionProductsMarkHook)
		productionProductsMarkAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productionProductsMarkBeforeDeleteMu.Lock()
		productionProductsMarkBeforeDeleteHooks = append(productionProductsMarkBeforeDeleteHooks, productionProductsMarkHook)
		productionProductsMarkBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productionProductsMarkAfterDeleteMu.Lock()
		productionProductsMarkAfterDeleteHooks = append(productionProductsMarkAfterDeleteHooks, productionProductsMarkHook)
		productionProductsMarkAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productionProductsMarkBeforeUpsertMu.Lock()
		productionProductsMarkBeforeUpsertHooks = append(productionProductsMarkBeforeUpsertHooks, productionProductsMarkHook)
		productionProductsMarkBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productionProductsMarkAfterUpsertMu.Lock()
		productionProductsMarkAfterUpsertHooks = append(productionProductsMarkAfterUpsertHooks, productionProductsMarkHook)
		productionProductsMarkAfterUpsertMu.Unlock()
	}
}

// OneG returns a single productionProductsMark record from the query using the global executor.
func (q productionProductsMarkQuery) OneG(ctx context.Context) (*ProductionProductsMark, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productionProductsMark record from the query.
func (q productionProductsMarkQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductionProductsMark, error) {
	o := &ProductionProductsMark{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for production_products_marks")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductionProductsMark records from the query using the global executor.
func (q productionProductsMarkQuery) AllG(ctx context.Context) (ProductionProductsMarkSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductionProductsMark records from the query.
func (q productionProductsMarkQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductionProductsMarkSlice, error) {
	var o []*ProductionProductsMark

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ProductionProductsMark slice")
	}

	if len(productionProductsMarkAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductionProductsMark records in the query using the global executor
func (q productionProductsMarkQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductionProductsMark records in the query.
func (q productionProductsMarkQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count production_products_marks rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productionProductsMarkQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productionProductsMarkQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if production_products_marks exists")
	}

	return count > 0, nil
}

// ProductionProductsMarks retrieves all the records using an executor.
func ProductionProductsMarks(mods ...qm.QueryMod) productionProductsMarkQuery {
	mods = append(mods, qm.From("[dbo].[production_products_marks]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[production_products_marks].*"})
	}

	return productionProductsMarkQuery{q}
}

// FindProductionProductsMarkG retrieves a single record by ID.
func FindProductionProductsMarkG(ctx context.Context, iD int, selectCols ...string) (*ProductionProductsMark, error) {
	return FindProductionProductsMark(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductionProductsMark retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductionProductsMark(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProductionProductsMark, error) {
	productionProductsMarkObj := &ProductionProductsMark{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[production_products_marks] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productionProductsMarkObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from production_products_marks")
	}

	if err = productionProductsMarkObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productionProductsMarkObj, err
	}

	return productionProductsMarkObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductionProductsMark) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductionProductsMark) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_products_marks provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionProductsMarkColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productionProductsMarkInsertCacheMut.RLock()
	cache, cached := productionProductsMarkInsertCache[key]
	productionProductsMarkInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productionProductsMarkAllColumns,
			productionProductsMarkColumnsWithDefault,
			productionProductsMarkColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productionProductsMarkGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[production_products_marks] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[production_products_marks] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into production_products_marks")
	}

	if !cached {
		productionProductsMarkInsertCacheMut.Lock()
		productionProductsMarkInsertCache[key] = cache
		productionProductsMarkInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductionProductsMark record using the global executor.
// See Update for more documentation.
func (o *ProductionProductsMark) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductionProductsMark.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductionProductsMark) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productionProductsMarkUpdateCacheMut.RLock()
	cache, cached := productionProductsMarkUpdateCache[key]
	productionProductsMarkUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productionProductsMarkAllColumns,
			productionProductsMarkPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productionProductsMarkGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update production_products_marks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[production_products_marks] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, productionProductsMarkPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, append(wl, productionProductsMarkPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update production_products_marks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for production_products_marks")
	}

	if !cached {
		productionProductsMarkUpdateCacheMut.Lock()
		productionProductsMarkUpdateCache[key] = cache
		productionProductsMarkUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productionProductsMarkQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productionProductsMarkQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for production_products_marks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for production_products_marks")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductionProductsMarkSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductionProductsMarkSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionProductsMarkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[production_products_marks] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productionProductsMarkPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in productionProductsMark slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all productionProductsMark")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductionProductsMark) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductionProductsMark) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_products_marks provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionProductsMarkColumnsWithDefault, o)

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

	productionProductsMarkUpsertCacheMut.RLock()
	cache, cached := productionProductsMarkUpsertCache[key]
	productionProductsMarkUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productionProductsMarkAllColumns,
			productionProductsMarkColumnsWithDefault,
			productionProductsMarkColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, productionProductsMarkGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(productionProductsMarkPrimaryKeyColumns, v) && strmangle.ContainsAny(productionProductsMarkColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert production_products_marks, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			productionProductsMarkAllColumns,
			productionProductsMarkPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, productionProductsMarkGeneratedColumns)

		ret := strmangle.SetComplement(productionProductsMarkAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert production_products_marks, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[production_products_marks]", productionProductsMarkPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(productionProductsMarkPrimaryKeyColumns))
		copy(whitelist, productionProductsMarkPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productionProductsMarkType, productionProductsMarkMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert production_products_marks")
	}

	if !cached {
		productionProductsMarkUpsertCacheMut.Lock()
		productionProductsMarkUpsertCache[key] = cache
		productionProductsMarkUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductionProductsMark record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductionProductsMark) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductionProductsMark record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductionProductsMark) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ProductionProductsMark provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productionProductsMarkPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[production_products_marks] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from production_products_marks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for production_products_marks")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productionProductsMarkQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productionProductsMarkQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no productionProductsMarkQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from production_products_marks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_products_marks")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductionProductsMarkSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductionProductsMarkSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productionProductsMarkBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionProductsMarkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[production_products_marks] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionProductsMarkPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from productionProductsMark slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_products_marks")
	}

	if len(productionProductsMarkAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductionProductsMark) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ProductionProductsMark provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductionProductsMark) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductionProductsMark(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionProductsMarkSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ProductionProductsMarkSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionProductsMarkSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductionProductsMarkSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionProductsMarkPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[production_products_marks].* FROM [dbo].[production_products_marks] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productionProductsMarkPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ProductionProductsMarkSlice")
	}

	*o = slice

	return nil
}

// ProductionProductsMarkExistsG checks if the ProductionProductsMark row exists.
func ProductionProductsMarkExistsG(ctx context.Context, iD int) (bool, error) {
	return ProductionProductsMarkExists(ctx, boil.GetContextDB(), iD)
}

// ProductionProductsMarkExists checks if the ProductionProductsMark row exists.
func ProductionProductsMarkExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[production_products_marks] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if production_products_marks exists")
	}

	return exists, nil
}

// Exists checks if the ProductionProductsMark row exists.
func (o *ProductionProductsMark) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductionProductsMarkExists(ctx, exec, o.ID)
}