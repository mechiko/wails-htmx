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

// ProductCardsImage is an object representing the database table.
type ProductCardsImage struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDProductCards null.Int    `boil:"id_product_cards" json:"id_product_cards,omitempty" toml:"id_product_cards" yaml:"id_product_cards,omitempty"`
	PhotoType      null.String `boil:"photo_type" json:"photo_type,omitempty" toml:"photo_type" yaml:"photo_type,omitempty"`
	PhotoDate      null.String `boil:"photo_date" json:"photo_date,omitempty" toml:"photo_date" yaml:"photo_date,omitempty"`
	PhotoURL       null.String `boil:"photo_url" json:"photo_url,omitempty" toml:"photo_url" yaml:"photo_url,omitempty"`
	Barcode        null.String `boil:"barcode" json:"barcode,omitempty" toml:"barcode" yaml:"barcode,omitempty"`

	R *productCardsImageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productCardsImageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductCardsImageColumns = struct {
	ID             string
	IDProductCards string
	PhotoType      string
	PhotoDate      string
	PhotoURL       string
	Barcode        string
}{
	ID:             "id",
	IDProductCards: "id_product_cards",
	PhotoType:      "photo_type",
	PhotoDate:      "photo_date",
	PhotoURL:       "photo_url",
	Barcode:        "barcode",
}

var ProductCardsImageTableColumns = struct {
	ID             string
	IDProductCards string
	PhotoType      string
	PhotoDate      string
	PhotoURL       string
	Barcode        string
}{
	ID:             "product_cards_images.id",
	IDProductCards: "product_cards_images.id_product_cards",
	PhotoType:      "product_cards_images.photo_type",
	PhotoDate:      "product_cards_images.photo_date",
	PhotoURL:       "product_cards_images.photo_url",
	Barcode:        "product_cards_images.barcode",
}

// Generated where

var ProductCardsImageWhere = struct {
	ID             whereHelperint
	IDProductCards whereHelpernull_Int
	PhotoType      whereHelpernull_String
	PhotoDate      whereHelpernull_String
	PhotoURL       whereHelpernull_String
	Barcode        whereHelpernull_String
}{
	ID:             whereHelperint{field: "[dbo].[product_cards_images].[id]"},
	IDProductCards: whereHelpernull_Int{field: "[dbo].[product_cards_images].[id_product_cards]"},
	PhotoType:      whereHelpernull_String{field: "[dbo].[product_cards_images].[photo_type]"},
	PhotoDate:      whereHelpernull_String{field: "[dbo].[product_cards_images].[photo_date]"},
	PhotoURL:       whereHelpernull_String{field: "[dbo].[product_cards_images].[photo_url]"},
	Barcode:        whereHelpernull_String{field: "[dbo].[product_cards_images].[barcode]"},
}

// ProductCardsImageRels is where relationship names are stored.
var ProductCardsImageRels = struct {
}{}

// productCardsImageR is where relationships are stored.
type productCardsImageR struct {
}

// NewStruct creates a new relationship struct
func (*productCardsImageR) NewStruct() *productCardsImageR {
	return &productCardsImageR{}
}

// productCardsImageL is where Load methods for each relationship are stored.
type productCardsImageL struct{}

var (
	productCardsImageAllColumns            = []string{"id", "id_product_cards", "photo_type", "photo_date", "photo_url", "barcode"}
	productCardsImageColumnsWithoutDefault = []string{"id_product_cards", "photo_type", "photo_date", "photo_url", "barcode"}
	productCardsImageColumnsWithDefault    = []string{"id"}
	productCardsImagePrimaryKeyColumns     = []string{"id"}
	productCardsImageGeneratedColumns      = []string{"id"}
)

type (
	// ProductCardsImageSlice is an alias for a slice of pointers to ProductCardsImage.
	// This should almost always be used instead of []ProductCardsImage.
	ProductCardsImageSlice []*ProductCardsImage
	// ProductCardsImageHook is the signature for custom ProductCardsImage hook methods
	ProductCardsImageHook func(context.Context, boil.ContextExecutor, *ProductCardsImage) error

	productCardsImageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productCardsImageType                 = reflect.TypeOf(&ProductCardsImage{})
	productCardsImageMapping              = queries.MakeStructMapping(productCardsImageType)
	productCardsImagePrimaryKeyMapping, _ = queries.BindMapping(productCardsImageType, productCardsImageMapping, productCardsImagePrimaryKeyColumns)
	productCardsImageInsertCacheMut       sync.RWMutex
	productCardsImageInsertCache          = make(map[string]insertCache)
	productCardsImageUpdateCacheMut       sync.RWMutex
	productCardsImageUpdateCache          = make(map[string]updateCache)
	productCardsImageUpsertCacheMut       sync.RWMutex
	productCardsImageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productCardsImageAfterSelectMu sync.Mutex
var productCardsImageAfterSelectHooks []ProductCardsImageHook

var productCardsImageBeforeInsertMu sync.Mutex
var productCardsImageBeforeInsertHooks []ProductCardsImageHook
var productCardsImageAfterInsertMu sync.Mutex
var productCardsImageAfterInsertHooks []ProductCardsImageHook

var productCardsImageBeforeUpdateMu sync.Mutex
var productCardsImageBeforeUpdateHooks []ProductCardsImageHook
var productCardsImageAfterUpdateMu sync.Mutex
var productCardsImageAfterUpdateHooks []ProductCardsImageHook

var productCardsImageBeforeDeleteMu sync.Mutex
var productCardsImageBeforeDeleteHooks []ProductCardsImageHook
var productCardsImageAfterDeleteMu sync.Mutex
var productCardsImageAfterDeleteHooks []ProductCardsImageHook

var productCardsImageBeforeUpsertMu sync.Mutex
var productCardsImageBeforeUpsertHooks []ProductCardsImageHook
var productCardsImageAfterUpsertMu sync.Mutex
var productCardsImageAfterUpsertHooks []ProductCardsImageHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductCardsImage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductCardsImage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductCardsImage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductCardsImage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductCardsImage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductCardsImage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductCardsImage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductCardsImage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductCardsImage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardsImageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductCardsImageHook registers your hook function for all future operations.
func AddProductCardsImageHook(hookPoint boil.HookPoint, productCardsImageHook ProductCardsImageHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productCardsImageAfterSelectMu.Lock()
		productCardsImageAfterSelectHooks = append(productCardsImageAfterSelectHooks, productCardsImageHook)
		productCardsImageAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productCardsImageBeforeInsertMu.Lock()
		productCardsImageBeforeInsertHooks = append(productCardsImageBeforeInsertHooks, productCardsImageHook)
		productCardsImageBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productCardsImageAfterInsertMu.Lock()
		productCardsImageAfterInsertHooks = append(productCardsImageAfterInsertHooks, productCardsImageHook)
		productCardsImageAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productCardsImageBeforeUpdateMu.Lock()
		productCardsImageBeforeUpdateHooks = append(productCardsImageBeforeUpdateHooks, productCardsImageHook)
		productCardsImageBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productCardsImageAfterUpdateMu.Lock()
		productCardsImageAfterUpdateHooks = append(productCardsImageAfterUpdateHooks, productCardsImageHook)
		productCardsImageAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productCardsImageBeforeDeleteMu.Lock()
		productCardsImageBeforeDeleteHooks = append(productCardsImageBeforeDeleteHooks, productCardsImageHook)
		productCardsImageBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productCardsImageAfterDeleteMu.Lock()
		productCardsImageAfterDeleteHooks = append(productCardsImageAfterDeleteHooks, productCardsImageHook)
		productCardsImageAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productCardsImageBeforeUpsertMu.Lock()
		productCardsImageBeforeUpsertHooks = append(productCardsImageBeforeUpsertHooks, productCardsImageHook)
		productCardsImageBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productCardsImageAfterUpsertMu.Lock()
		productCardsImageAfterUpsertHooks = append(productCardsImageAfterUpsertHooks, productCardsImageHook)
		productCardsImageAfterUpsertMu.Unlock()
	}
}

// OneG returns a single productCardsImage record from the query using the global executor.
func (q productCardsImageQuery) OneG(ctx context.Context) (*ProductCardsImage, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productCardsImage record from the query.
func (q productCardsImageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductCardsImage, error) {
	o := &ProductCardsImage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: failed to execute a one query for product_cards_images")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductCardsImage records from the query using the global executor.
func (q productCardsImageQuery) AllG(ctx context.Context) (ProductCardsImageSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductCardsImage records from the query.
func (q productCardsImageQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductCardsImageSlice, error) {
	var o []*ProductCardsImage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "znakboil: failed to assign all query results to ProductCardsImage slice")
	}

	if len(productCardsImageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductCardsImage records in the query using the global executor
func (q productCardsImageQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductCardsImage records in the query.
func (q productCardsImageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to count product_cards_images rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productCardsImageQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productCardsImageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: failed to check if product_cards_images exists")
	}

	return count > 0, nil
}

// ProductCardsImages retrieves all the records using an executor.
func ProductCardsImages(mods ...qm.QueryMod) productCardsImageQuery {
	mods = append(mods, qm.From("[dbo].[product_cards_images]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[product_cards_images].*"})
	}

	return productCardsImageQuery{q}
}

// FindProductCardsImageG retrieves a single record by ID.
func FindProductCardsImageG(ctx context.Context, iD int, selectCols ...string) (*ProductCardsImage, error) {
	return FindProductCardsImage(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductCardsImage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductCardsImage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProductCardsImage, error) {
	productCardsImageObj := &ProductCardsImage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[product_cards_images] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productCardsImageObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: unable to select from product_cards_images")
	}

	if err = productCardsImageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productCardsImageObj, err
	}

	return productCardsImageObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductCardsImage) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductCardsImage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no product_cards_images provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productCardsImageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productCardsImageInsertCacheMut.RLock()
	cache, cached := productCardsImageInsertCache[key]
	productCardsImageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productCardsImageAllColumns,
			productCardsImageColumnsWithDefault,
			productCardsImageColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productCardsImageGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productCardsImageType, productCardsImageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productCardsImageType, productCardsImageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[product_cards_images] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[product_cards_images] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "znakboil: unable to insert into product_cards_images")
	}

	if !cached {
		productCardsImageInsertCacheMut.Lock()
		productCardsImageInsertCache[key] = cache
		productCardsImageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductCardsImage record using the global executor.
// See Update for more documentation.
func (o *ProductCardsImage) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductCardsImage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductCardsImage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productCardsImageUpdateCacheMut.RLock()
	cache, cached := productCardsImageUpdateCache[key]
	productCardsImageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productCardsImageAllColumns,
			productCardsImagePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productCardsImageGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("znakboil: unable to update product_cards_images, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[product_cards_images] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, productCardsImagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productCardsImageType, productCardsImageMapping, append(wl, productCardsImagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "znakboil: unable to update product_cards_images row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by update for product_cards_images")
	}

	if !cached {
		productCardsImageUpdateCacheMut.Lock()
		productCardsImageUpdateCache[key] = cache
		productCardsImageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productCardsImageQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productCardsImageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all for product_cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected for product_cards_images")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductCardsImageSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductCardsImageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[product_cards_images] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productCardsImagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all in productCardsImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected all in update all productCardsImage")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductCardsImage) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductCardsImage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no product_cards_images provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productCardsImageColumnsWithDefault, o)

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

	productCardsImageUpsertCacheMut.RLock()
	cache, cached := productCardsImageUpsertCache[key]
	productCardsImageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productCardsImageAllColumns,
			productCardsImageColumnsWithDefault,
			productCardsImageColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, productCardsImageGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(productCardsImagePrimaryKeyColumns, v) && strmangle.ContainsAny(productCardsImageColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("znakboil: unable to upsert product_cards_images, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			productCardsImageAllColumns,
			productCardsImagePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, productCardsImageGeneratedColumns)

		ret := strmangle.SetComplement(productCardsImageAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("znakboil: unable to upsert product_cards_images, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[product_cards_images]", productCardsImagePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(productCardsImagePrimaryKeyColumns))
		copy(whitelist, productCardsImagePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(productCardsImageType, productCardsImageMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productCardsImageType, productCardsImageMapping, ret)
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
		return errors.Wrap(err, "znakboil: unable to upsert product_cards_images")
	}

	if !cached {
		productCardsImageUpsertCacheMut.Lock()
		productCardsImageUpsertCache[key] = cache
		productCardsImageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductCardsImage record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductCardsImage) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductCardsImage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductCardsImage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("znakboil: no ProductCardsImage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productCardsImagePrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[product_cards_images] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete from product_cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by delete for product_cards_images")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productCardsImageQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productCardsImageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("znakboil: no productCardsImageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from product_cards_images")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for product_cards_images")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductCardsImageSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductCardsImageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productCardsImageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[product_cards_images] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productCardsImagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from productCardsImage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for product_cards_images")
	}

	if len(productCardsImageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductCardsImage) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: no ProductCardsImage provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductCardsImage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductCardsImage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductCardsImageSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: empty ProductCardsImageSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductCardsImageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductCardsImageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardsImagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[product_cards_images].* FROM [dbo].[product_cards_images] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productCardsImagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "znakboil: unable to reload all in ProductCardsImageSlice")
	}

	*o = slice

	return nil
}

// ProductCardsImageExistsG checks if the ProductCardsImage row exists.
func ProductCardsImageExistsG(ctx context.Context, iD int) (bool, error) {
	return ProductCardsImageExists(ctx, boil.GetContextDB(), iD)
}

// ProductCardsImageExists checks if the ProductCardsImage row exists.
func ProductCardsImageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[product_cards_images] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: unable to check if product_cards_images exists")
	}

	return exists, nil
}

// Exists checks if the ProductCardsImage row exists.
func (o *ProductCardsImage) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductCardsImageExists(ctx, exec, o.ID)
}