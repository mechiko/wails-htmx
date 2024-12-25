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

// ProductCard is an object representing the database table.
type ProductCard struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Gtin       null.String `boil:"gtin" json:"gtin,omitempty" toml:"gtin" yaml:"gtin,omitempty"`
	GoodID     null.Int    `boil:"good_id" json:"good_id,omitempty" toml:"good_id" yaml:"good_id,omitempty"`
	GoodName   null.String `boil:"good_name" json:"good_name,omitempty" toml:"good_name" yaml:"good_name,omitempty"`
	IsKit      null.String `boil:"is_kit" json:"is_kit,omitempty" toml:"is_kit" yaml:"is_kit,omitempty"`
	IsSet      null.String `boil:"is_set" json:"is_set,omitempty" toml:"is_set" yaml:"is_set,omitempty"`
	GoodURL    null.String `boil:"good_url" json:"good_url,omitempty" toml:"good_url" yaml:"good_url,omitempty"`
	GoodImg    null.String `boil:"good_img" json:"good_img,omitempty" toml:"good_img" yaml:"good_img,omitempty"`
	GoodStatus null.String `boil:"good_status" json:"good_status,omitempty" toml:"good_status" yaml:"good_status,omitempty"`
	CreateDate null.String `boil:"create_date" json:"create_date,omitempty" toml:"create_date" yaml:"create_date,omitempty"`
	UpdateDate null.String `boil:"update_date" json:"update_date,omitempty" toml:"update_date" yaml:"update_date,omitempty"`
	BrandID    null.Int    `boil:"brand_id" json:"brand_id,omitempty" toml:"brand_id" yaml:"brand_id,omitempty"`
	BrandName  null.String `boil:"brand_name" json:"brand_name,omitempty" toml:"brand_name" yaml:"brand_name,omitempty"`
	GoodRating null.Int    `boil:"good_rating" json:"good_rating,omitempty" toml:"good_rating" yaml:"good_rating,omitempty"`

	R *productCardR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productCardL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductCardColumns = struct {
	ID         string
	Gtin       string
	GoodID     string
	GoodName   string
	IsKit      string
	IsSet      string
	GoodURL    string
	GoodImg    string
	GoodStatus string
	CreateDate string
	UpdateDate string
	BrandID    string
	BrandName  string
	GoodRating string
}{
	ID:         "id",
	Gtin:       "gtin",
	GoodID:     "good_id",
	GoodName:   "good_name",
	IsKit:      "is_kit",
	IsSet:      "is_set",
	GoodURL:    "good_url",
	GoodImg:    "good_img",
	GoodStatus: "good_status",
	CreateDate: "create_date",
	UpdateDate: "update_date",
	BrandID:    "brand_id",
	BrandName:  "brand_name",
	GoodRating: "good_rating",
}

var ProductCardTableColumns = struct {
	ID         string
	Gtin       string
	GoodID     string
	GoodName   string
	IsKit      string
	IsSet      string
	GoodURL    string
	GoodImg    string
	GoodStatus string
	CreateDate string
	UpdateDate string
	BrandID    string
	BrandName  string
	GoodRating string
}{
	ID:         "product_cards.id",
	Gtin:       "product_cards.gtin",
	GoodID:     "product_cards.good_id",
	GoodName:   "product_cards.good_name",
	IsKit:      "product_cards.is_kit",
	IsSet:      "product_cards.is_set",
	GoodURL:    "product_cards.good_url",
	GoodImg:    "product_cards.good_img",
	GoodStatus: "product_cards.good_status",
	CreateDate: "product_cards.create_date",
	UpdateDate: "product_cards.update_date",
	BrandID:    "product_cards.brand_id",
	BrandName:  "product_cards.brand_name",
	GoodRating: "product_cards.good_rating",
}

// Generated where

var ProductCardWhere = struct {
	ID         whereHelperint
	Gtin       whereHelpernull_String
	GoodID     whereHelpernull_Int
	GoodName   whereHelpernull_String
	IsKit      whereHelpernull_String
	IsSet      whereHelpernull_String
	GoodURL    whereHelpernull_String
	GoodImg    whereHelpernull_String
	GoodStatus whereHelpernull_String
	CreateDate whereHelpernull_String
	UpdateDate whereHelpernull_String
	BrandID    whereHelpernull_Int
	BrandName  whereHelpernull_String
	GoodRating whereHelpernull_Int
}{
	ID:         whereHelperint{field: "[dbo].[product_cards].[id]"},
	Gtin:       whereHelpernull_String{field: "[dbo].[product_cards].[gtin]"},
	GoodID:     whereHelpernull_Int{field: "[dbo].[product_cards].[good_id]"},
	GoodName:   whereHelpernull_String{field: "[dbo].[product_cards].[good_name]"},
	IsKit:      whereHelpernull_String{field: "[dbo].[product_cards].[is_kit]"},
	IsSet:      whereHelpernull_String{field: "[dbo].[product_cards].[is_set]"},
	GoodURL:    whereHelpernull_String{field: "[dbo].[product_cards].[good_url]"},
	GoodImg:    whereHelpernull_String{field: "[dbo].[product_cards].[good_img]"},
	GoodStatus: whereHelpernull_String{field: "[dbo].[product_cards].[good_status]"},
	CreateDate: whereHelpernull_String{field: "[dbo].[product_cards].[create_date]"},
	UpdateDate: whereHelpernull_String{field: "[dbo].[product_cards].[update_date]"},
	BrandID:    whereHelpernull_Int{field: "[dbo].[product_cards].[brand_id]"},
	BrandName:  whereHelpernull_String{field: "[dbo].[product_cards].[brand_name]"},
	GoodRating: whereHelpernull_Int{field: "[dbo].[product_cards].[good_rating]"},
}

// ProductCardRels is where relationship names are stored.
var ProductCardRels = struct {
}{}

// productCardR is where relationships are stored.
type productCardR struct {
}

// NewStruct creates a new relationship struct
func (*productCardR) NewStruct() *productCardR {
	return &productCardR{}
}

// productCardL is where Load methods for each relationship are stored.
type productCardL struct{}

var (
	productCardAllColumns            = []string{"id", "gtin", "good_id", "good_name", "is_kit", "is_set", "good_url", "good_img", "good_status", "create_date", "update_date", "brand_id", "brand_name", "good_rating"}
	productCardColumnsWithoutDefault = []string{"gtin", "good_id", "good_name", "is_kit", "is_set", "good_url", "good_img", "good_status", "create_date", "update_date", "brand_id", "brand_name", "good_rating"}
	productCardColumnsWithDefault    = []string{"id"}
	productCardPrimaryKeyColumns     = []string{"id"}
	productCardGeneratedColumns      = []string{"id"}
)

type (
	// ProductCardSlice is an alias for a slice of pointers to ProductCard.
	// This should almost always be used instead of []ProductCard.
	ProductCardSlice []*ProductCard
	// ProductCardHook is the signature for custom ProductCard hook methods
	ProductCardHook func(context.Context, boil.ContextExecutor, *ProductCard) error

	productCardQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productCardType                 = reflect.TypeOf(&ProductCard{})
	productCardMapping              = queries.MakeStructMapping(productCardType)
	productCardPrimaryKeyMapping, _ = queries.BindMapping(productCardType, productCardMapping, productCardPrimaryKeyColumns)
	productCardInsertCacheMut       sync.RWMutex
	productCardInsertCache          = make(map[string]insertCache)
	productCardUpdateCacheMut       sync.RWMutex
	productCardUpdateCache          = make(map[string]updateCache)
	productCardUpsertCacheMut       sync.RWMutex
	productCardUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productCardAfterSelectMu sync.Mutex
var productCardAfterSelectHooks []ProductCardHook

var productCardBeforeInsertMu sync.Mutex
var productCardBeforeInsertHooks []ProductCardHook
var productCardAfterInsertMu sync.Mutex
var productCardAfterInsertHooks []ProductCardHook

var productCardBeforeUpdateMu sync.Mutex
var productCardBeforeUpdateHooks []ProductCardHook
var productCardAfterUpdateMu sync.Mutex
var productCardAfterUpdateHooks []ProductCardHook

var productCardBeforeDeleteMu sync.Mutex
var productCardBeforeDeleteHooks []ProductCardHook
var productCardAfterDeleteMu sync.Mutex
var productCardAfterDeleteHooks []ProductCardHook

var productCardBeforeUpsertMu sync.Mutex
var productCardBeforeUpsertHooks []ProductCardHook
var productCardAfterUpsertMu sync.Mutex
var productCardAfterUpsertHooks []ProductCardHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductCard) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductCard) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductCard) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductCard) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductCard) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductCard) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductCard) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductCard) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductCard) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productCardAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductCardHook registers your hook function for all future operations.
func AddProductCardHook(hookPoint boil.HookPoint, productCardHook ProductCardHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productCardAfterSelectMu.Lock()
		productCardAfterSelectHooks = append(productCardAfterSelectHooks, productCardHook)
		productCardAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productCardBeforeInsertMu.Lock()
		productCardBeforeInsertHooks = append(productCardBeforeInsertHooks, productCardHook)
		productCardBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productCardAfterInsertMu.Lock()
		productCardAfterInsertHooks = append(productCardAfterInsertHooks, productCardHook)
		productCardAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productCardBeforeUpdateMu.Lock()
		productCardBeforeUpdateHooks = append(productCardBeforeUpdateHooks, productCardHook)
		productCardBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productCardAfterUpdateMu.Lock()
		productCardAfterUpdateHooks = append(productCardAfterUpdateHooks, productCardHook)
		productCardAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productCardBeforeDeleteMu.Lock()
		productCardBeforeDeleteHooks = append(productCardBeforeDeleteHooks, productCardHook)
		productCardBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productCardAfterDeleteMu.Lock()
		productCardAfterDeleteHooks = append(productCardAfterDeleteHooks, productCardHook)
		productCardAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productCardBeforeUpsertMu.Lock()
		productCardBeforeUpsertHooks = append(productCardBeforeUpsertHooks, productCardHook)
		productCardBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productCardAfterUpsertMu.Lock()
		productCardAfterUpsertHooks = append(productCardAfterUpsertHooks, productCardHook)
		productCardAfterUpsertMu.Unlock()
	}
}

// OneG returns a single productCard record from the query using the global executor.
func (q productCardQuery) OneG(ctx context.Context) (*ProductCard, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productCard record from the query.
func (q productCardQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProductCard, error) {
	o := &ProductCard{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: failed to execute a one query for product_cards")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductCard records from the query using the global executor.
func (q productCardQuery) AllG(ctx context.Context) (ProductCardSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductCard records from the query.
func (q productCardQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProductCardSlice, error) {
	var o []*ProductCard

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "znakboil: failed to assign all query results to ProductCard slice")
	}

	if len(productCardAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductCard records in the query using the global executor
func (q productCardQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductCard records in the query.
func (q productCardQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to count product_cards rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productCardQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productCardQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: failed to check if product_cards exists")
	}

	return count > 0, nil
}

// ProductCards retrieves all the records using an executor.
func ProductCards(mods ...qm.QueryMod) productCardQuery {
	mods = append(mods, qm.From("[dbo].[product_cards]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[product_cards].*"})
	}

	return productCardQuery{q}
}

// FindProductCardG retrieves a single record by ID.
func FindProductCardG(ctx context.Context, iD int, selectCols ...string) (*ProductCard, error) {
	return FindProductCard(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductCard retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductCard(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProductCard, error) {
	productCardObj := &ProductCard{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[product_cards] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productCardObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "znakboil: unable to select from product_cards")
	}

	if err = productCardObj.doAfterSelectHooks(ctx, exec); err != nil {
		return productCardObj, err
	}

	return productCardObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductCard) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductCard) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no product_cards provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productCardColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productCardInsertCacheMut.RLock()
	cache, cached := productCardInsertCache[key]
	productCardInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productCardAllColumns,
			productCardColumnsWithDefault,
			productCardColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productCardGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productCardType, productCardMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productCardType, productCardMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[product_cards] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[product_cards] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "znakboil: unable to insert into product_cards")
	}

	if !cached {
		productCardInsertCacheMut.Lock()
		productCardInsertCache[key] = cache
		productCardInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductCard record using the global executor.
// See Update for more documentation.
func (o *ProductCard) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductCard.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductCard) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productCardUpdateCacheMut.RLock()
	cache, cached := productCardUpdateCache[key]
	productCardUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productCardAllColumns,
			productCardPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productCardGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("znakboil: unable to update product_cards, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[product_cards] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, productCardPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productCardType, productCardMapping, append(wl, productCardPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "znakboil: unable to update product_cards row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by update for product_cards")
	}

	if !cached {
		productCardUpdateCacheMut.Lock()
		productCardUpdateCache[key] = cache
		productCardUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productCardQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productCardQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all for product_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected for product_cards")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductCardSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductCardSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[product_cards] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, productCardPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to update all in productCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to retrieve rows affected all in update all productCard")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductCard) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ProductCard) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("znakboil: no product_cards provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productCardColumnsWithDefault, o)

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

	productCardUpsertCacheMut.RLock()
	cache, cached := productCardUpsertCache[key]
	productCardUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			productCardAllColumns,
			productCardColumnsWithDefault,
			productCardColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, productCardGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(productCardPrimaryKeyColumns, v) && strmangle.ContainsAny(productCardColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("znakboil: unable to upsert product_cards, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			productCardAllColumns,
			productCardPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, productCardGeneratedColumns)

		ret := strmangle.SetComplement(productCardAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("znakboil: unable to upsert product_cards, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[product_cards]", productCardPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(productCardPrimaryKeyColumns))
		copy(whitelist, productCardPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(productCardType, productCardMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productCardType, productCardMapping, ret)
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
		return errors.Wrap(err, "znakboil: unable to upsert product_cards")
	}

	if !cached {
		productCardUpsertCacheMut.Lock()
		productCardUpsertCache[key] = cache
		productCardUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductCard record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductCard) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductCard record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductCard) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("znakboil: no ProductCard provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productCardPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[product_cards] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete from product_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by delete for product_cards")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productCardQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productCardQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("znakboil: no productCardQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from product_cards")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for product_cards")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductCardSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductCardSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productCardBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[product_cards] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productCardPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: unable to delete all from productCard slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "znakboil: failed to get rows affected by deleteall for product_cards")
	}

	if len(productCardAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductCard) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: no ProductCard provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductCard) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductCard(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductCardSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("znakboil: empty ProductCardSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductCardSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductCardSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productCardPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[product_cards].* FROM [dbo].[product_cards] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, productCardPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "znakboil: unable to reload all in ProductCardSlice")
	}

	*o = slice

	return nil
}

// ProductCardExistsG checks if the ProductCard row exists.
func ProductCardExistsG(ctx context.Context, iD int) (bool, error) {
	return ProductCardExists(ctx, boil.GetContextDB(), iD)
}

// ProductCardExists checks if the ProductCard row exists.
func ProductCardExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[product_cards] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "znakboil: unable to check if product_cards exists")
	}

	return exists, nil
}

// Exists checks if the ProductCard row exists.
func (o *ProductCard) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductCardExists(ctx, exec, o.ID)
}
