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

// ProductionForm1 is an object representing the database table.
type ProductionForm1 struct {
	ID                  int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDProductionReports null.Int64  `boil:"id_production_reports" json:"id_production_reports,omitempty" toml:"id_production_reports" yaml:"id_production_reports,omitempty"`
	DocIdentity         null.String `boil:"doc_identity" json:"doc_identity,omitempty" toml:"doc_identity" yaml:"doc_identity,omitempty"`
	DocRegID            null.String `boil:"doc_reg_id" json:"doc_reg_id,omitempty" toml:"doc_reg_id" yaml:"doc_reg_id,omitempty"`
	ClientType          null.String `boil:"client_type" json:"client_type,omitempty" toml:"client_type" yaml:"client_type,omitempty"`
	ClientRegID         null.String `boil:"client_reg_id" json:"client_reg_id,omitempty" toml:"client_reg_id" yaml:"client_reg_id,omitempty"`
	ClientInn           null.String `boil:"client_inn" json:"client_inn,omitempty" toml:"client_inn" yaml:"client_inn,omitempty"`
	ClientKPP           null.String `boil:"client_kpp" json:"client_kpp,omitempty" toml:"client_kpp" yaml:"client_kpp,omitempty"`
	ClientFullName      null.String `boil:"client_full_name" json:"client_full_name,omitempty" toml:"client_full_name" yaml:"client_full_name,omitempty"`
	ClientShortName     null.String `boil:"client_short_name" json:"client_short_name,omitempty" toml:"client_short_name" yaml:"client_short_name,omitempty"`
	ClientCountryCode   null.String `boil:"client_country_code" json:"client_country_code,omitempty" toml:"client_country_code" yaml:"client_country_code,omitempty"`
	ClientRegionCode    null.String `boil:"client_region_code" json:"client_region_code,omitempty" toml:"client_region_code" yaml:"client_region_code,omitempty"`
	ClientDescription   null.String `boil:"client_description" json:"client_description,omitempty" toml:"client_description" yaml:"client_description,omitempty"`
	XML                 null.String `boil:"xml" json:"xml,omitempty" toml:"xml" yaml:"xml,omitempty"`

	R *productionForm1R `boil:"-" json:"-" toml:"-" yaml:"-"`
	L productionForm1L  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProductionForm1Columns = struct {
	ID                  string
	IDProductionReports string
	DocIdentity         string
	DocRegID            string
	ClientType          string
	ClientRegID         string
	ClientInn           string
	ClientKPP           string
	ClientFullName      string
	ClientShortName     string
	ClientCountryCode   string
	ClientRegionCode    string
	ClientDescription   string
	XML                 string
}{
	ID:                  "id",
	IDProductionReports: "id_production_reports",
	DocIdentity:         "doc_identity",
	DocRegID:            "doc_reg_id",
	ClientType:          "client_type",
	ClientRegID:         "client_reg_id",
	ClientInn:           "client_inn",
	ClientKPP:           "client_kpp",
	ClientFullName:      "client_full_name",
	ClientShortName:     "client_short_name",
	ClientCountryCode:   "client_country_code",
	ClientRegionCode:    "client_region_code",
	ClientDescription:   "client_description",
	XML:                 "xml",
}

var ProductionForm1TableColumns = struct {
	ID                  string
	IDProductionReports string
	DocIdentity         string
	DocRegID            string
	ClientType          string
	ClientRegID         string
	ClientInn           string
	ClientKPP           string
	ClientFullName      string
	ClientShortName     string
	ClientCountryCode   string
	ClientRegionCode    string
	ClientDescription   string
	XML                 string
}{
	ID:                  "production_form1.id",
	IDProductionReports: "production_form1.id_production_reports",
	DocIdentity:         "production_form1.doc_identity",
	DocRegID:            "production_form1.doc_reg_id",
	ClientType:          "production_form1.client_type",
	ClientRegID:         "production_form1.client_reg_id",
	ClientInn:           "production_form1.client_inn",
	ClientKPP:           "production_form1.client_kpp",
	ClientFullName:      "production_form1.client_full_name",
	ClientShortName:     "production_form1.client_short_name",
	ClientCountryCode:   "production_form1.client_country_code",
	ClientRegionCode:    "production_form1.client_region_code",
	ClientDescription:   "production_form1.client_description",
	XML:                 "production_form1.xml",
}

// Generated where

var ProductionForm1Where = struct {
	ID                  whereHelperint64
	IDProductionReports whereHelpernull_Int64
	DocIdentity         whereHelpernull_String
	DocRegID            whereHelpernull_String
	ClientType          whereHelpernull_String
	ClientRegID         whereHelpernull_String
	ClientInn           whereHelpernull_String
	ClientKPP           whereHelpernull_String
	ClientFullName      whereHelpernull_String
	ClientShortName     whereHelpernull_String
	ClientCountryCode   whereHelpernull_String
	ClientRegionCode    whereHelpernull_String
	ClientDescription   whereHelpernull_String
	XML                 whereHelpernull_String
}{
	ID:                  whereHelperint64{field: "\"production_form1\".\"id\""},
	IDProductionReports: whereHelpernull_Int64{field: "\"production_form1\".\"id_production_reports\""},
	DocIdentity:         whereHelpernull_String{field: "\"production_form1\".\"doc_identity\""},
	DocRegID:            whereHelpernull_String{field: "\"production_form1\".\"doc_reg_id\""},
	ClientType:          whereHelpernull_String{field: "\"production_form1\".\"client_type\""},
	ClientRegID:         whereHelpernull_String{field: "\"production_form1\".\"client_reg_id\""},
	ClientInn:           whereHelpernull_String{field: "\"production_form1\".\"client_inn\""},
	ClientKPP:           whereHelpernull_String{field: "\"production_form1\".\"client_kpp\""},
	ClientFullName:      whereHelpernull_String{field: "\"production_form1\".\"client_full_name\""},
	ClientShortName:     whereHelpernull_String{field: "\"production_form1\".\"client_short_name\""},
	ClientCountryCode:   whereHelpernull_String{field: "\"production_form1\".\"client_country_code\""},
	ClientRegionCode:    whereHelpernull_String{field: "\"production_form1\".\"client_region_code\""},
	ClientDescription:   whereHelpernull_String{field: "\"production_form1\".\"client_description\""},
	XML:                 whereHelpernull_String{field: "\"production_form1\".\"xml\""},
}

// ProductionForm1Rels is where relationship names are stored.
var ProductionForm1Rels = struct {
}{}

// productionForm1R is where relationships are stored.
type productionForm1R struct {
}

// NewStruct creates a new relationship struct
func (*productionForm1R) NewStruct() *productionForm1R {
	return &productionForm1R{}
}

// productionForm1L is where Load methods for each relationship are stored.
type productionForm1L struct{}

var (
	productionForm1AllColumns            = []string{"id", "id_production_reports", "doc_identity", "doc_reg_id", "client_type", "client_reg_id", "client_inn", "client_kpp", "client_full_name", "client_short_name", "client_country_code", "client_region_code", "client_description", "xml"}
	productionForm1ColumnsWithoutDefault = []string{}
	productionForm1ColumnsWithDefault    = []string{"id", "id_production_reports", "doc_identity", "doc_reg_id", "client_type", "client_reg_id", "client_inn", "client_kpp", "client_full_name", "client_short_name", "client_country_code", "client_region_code", "client_description", "xml"}
	productionForm1PrimaryKeyColumns     = []string{"id"}
	productionForm1GeneratedColumns      = []string{"id"}
)

type (
	// ProductionForm1Slice is an alias for a slice of pointers to ProductionForm1.
	// This should almost always be used instead of []ProductionForm1.
	ProductionForm1Slice []*ProductionForm1
	// ProductionForm1Hook is the signature for custom ProductionForm1 hook methods
	ProductionForm1Hook func(context.Context, boil.ContextExecutor, *ProductionForm1) error

	productionForm1Query struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	productionForm1Type                 = reflect.TypeOf(&ProductionForm1{})
	productionForm1Mapping              = queries.MakeStructMapping(productionForm1Type)
	productionForm1PrimaryKeyMapping, _ = queries.BindMapping(productionForm1Type, productionForm1Mapping, productionForm1PrimaryKeyColumns)
	productionForm1InsertCacheMut       sync.RWMutex
	productionForm1InsertCache          = make(map[string]insertCache)
	productionForm1UpdateCacheMut       sync.RWMutex
	productionForm1UpdateCache          = make(map[string]updateCache)
	productionForm1UpsertCacheMut       sync.RWMutex
	productionForm1UpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var productionForm1AfterSelectMu sync.Mutex
var productionForm1AfterSelectHooks []ProductionForm1Hook

var productionForm1BeforeInsertMu sync.Mutex
var productionForm1BeforeInsertHooks []ProductionForm1Hook
var productionForm1AfterInsertMu sync.Mutex
var productionForm1AfterInsertHooks []ProductionForm1Hook

var productionForm1BeforeUpdateMu sync.Mutex
var productionForm1BeforeUpdateHooks []ProductionForm1Hook
var productionForm1AfterUpdateMu sync.Mutex
var productionForm1AfterUpdateHooks []ProductionForm1Hook

var productionForm1BeforeDeleteMu sync.Mutex
var productionForm1BeforeDeleteHooks []ProductionForm1Hook
var productionForm1AfterDeleteMu sync.Mutex
var productionForm1AfterDeleteHooks []ProductionForm1Hook

var productionForm1BeforeUpsertMu sync.Mutex
var productionForm1BeforeUpsertHooks []ProductionForm1Hook
var productionForm1AfterUpsertMu sync.Mutex
var productionForm1AfterUpsertHooks []ProductionForm1Hook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProductionForm1) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1AfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProductionForm1) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1BeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProductionForm1) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1AfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProductionForm1) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1BeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProductionForm1) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1AfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProductionForm1) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1BeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProductionForm1) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1AfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProductionForm1) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1BeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProductionForm1) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range productionForm1AfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProductionForm1Hook registers your hook function for all future operations.
func AddProductionForm1Hook(hookPoint boil.HookPoint, productionForm1Hook ProductionForm1Hook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		productionForm1AfterSelectMu.Lock()
		productionForm1AfterSelectHooks = append(productionForm1AfterSelectHooks, productionForm1Hook)
		productionForm1AfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		productionForm1BeforeInsertMu.Lock()
		productionForm1BeforeInsertHooks = append(productionForm1BeforeInsertHooks, productionForm1Hook)
		productionForm1BeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		productionForm1AfterInsertMu.Lock()
		productionForm1AfterInsertHooks = append(productionForm1AfterInsertHooks, productionForm1Hook)
		productionForm1AfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		productionForm1BeforeUpdateMu.Lock()
		productionForm1BeforeUpdateHooks = append(productionForm1BeforeUpdateHooks, productionForm1Hook)
		productionForm1BeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		productionForm1AfterUpdateMu.Lock()
		productionForm1AfterUpdateHooks = append(productionForm1AfterUpdateHooks, productionForm1Hook)
		productionForm1AfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		productionForm1BeforeDeleteMu.Lock()
		productionForm1BeforeDeleteHooks = append(productionForm1BeforeDeleteHooks, productionForm1Hook)
		productionForm1BeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		productionForm1AfterDeleteMu.Lock()
		productionForm1AfterDeleteHooks = append(productionForm1AfterDeleteHooks, productionForm1Hook)
		productionForm1AfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		productionForm1BeforeUpsertMu.Lock()
		productionForm1BeforeUpsertHooks = append(productionForm1BeforeUpsertHooks, productionForm1Hook)
		productionForm1BeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		productionForm1AfterUpsertMu.Lock()
		productionForm1AfterUpsertHooks = append(productionForm1AfterUpsertHooks, productionForm1Hook)
		productionForm1AfterUpsertMu.Unlock()
	}
}

// OneG returns a single productionForm1 record from the query using the global executor.
func (q productionForm1Query) OneG(ctx context.Context) (*ProductionForm1, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single productionForm1 record from the query.
func (q productionForm1Query) One(ctx context.Context, exec boil.ContextExecutor) (*ProductionForm1, error) {
	o := &ProductionForm1{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for production_form1")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ProductionForm1 records from the query using the global executor.
func (q productionForm1Query) AllG(ctx context.Context) (ProductionForm1Slice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ProductionForm1 records from the query.
func (q productionForm1Query) All(ctx context.Context, exec boil.ContextExecutor) (ProductionForm1Slice, error) {
	var o []*ProductionForm1

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ProductionForm1 slice")
	}

	if len(productionForm1AfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ProductionForm1 records in the query using the global executor
func (q productionForm1Query) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ProductionForm1 records in the query.
func (q productionForm1Query) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count production_form1 rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q productionForm1Query) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q productionForm1Query) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if production_form1 exists")
	}

	return count > 0, nil
}

// ProductionForm1s retrieves all the records using an executor.
func ProductionForm1s(mods ...qm.QueryMod) productionForm1Query {
	mods = append(mods, qm.From("\"production_form1\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"production_form1\".*"})
	}

	return productionForm1Query{q}
}

// FindProductionForm1G retrieves a single record by ID.
func FindProductionForm1G(ctx context.Context, iD int64, selectCols ...string) (*ProductionForm1, error) {
	return FindProductionForm1(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindProductionForm1 retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProductionForm1(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ProductionForm1, error) {
	productionForm1Obj := &ProductionForm1{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"production_form1\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, productionForm1Obj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from production_form1")
	}

	if err = productionForm1Obj.doAfterSelectHooks(ctx, exec); err != nil {
		return productionForm1Obj, err
	}

	return productionForm1Obj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ProductionForm1) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProductionForm1) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_form1 provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionForm1ColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	productionForm1InsertCacheMut.RLock()
	cache, cached := productionForm1InsertCache[key]
	productionForm1InsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			productionForm1AllColumns,
			productionForm1ColumnsWithDefault,
			productionForm1ColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, productionForm1GeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(productionForm1Type, productionForm1Mapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(productionForm1Type, productionForm1Mapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"production_form1\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"production_form1\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into production_form1")
	}

	if !cached {
		productionForm1InsertCacheMut.Lock()
		productionForm1InsertCache[key] = cache
		productionForm1InsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ProductionForm1 record using the global executor.
// See Update for more documentation.
func (o *ProductionForm1) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ProductionForm1.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProductionForm1) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	productionForm1UpdateCacheMut.RLock()
	cache, cached := productionForm1UpdateCache[key]
	productionForm1UpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			productionForm1AllColumns,
			productionForm1PrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, productionForm1GeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update production_form1, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"production_form1\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, productionForm1PrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(productionForm1Type, productionForm1Mapping, append(wl, productionForm1PrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update production_form1 row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for production_form1")
	}

	if !cached {
		productionForm1UpdateCacheMut.Lock()
		productionForm1UpdateCache[key] = cache
		productionForm1UpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q productionForm1Query) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q productionForm1Query) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for production_form1")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for production_form1")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ProductionForm1Slice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProductionForm1Slice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionForm1PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"production_form1\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productionForm1PrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in productionForm1 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all productionForm1")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ProductionForm1) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProductionForm1) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no production_form1 provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(productionForm1ColumnsWithDefault, o)

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

	productionForm1UpsertCacheMut.RLock()
	cache, cached := productionForm1UpsertCache[key]
	productionForm1UpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			productionForm1AllColumns,
			productionForm1ColumnsWithDefault,
			productionForm1ColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			productionForm1AllColumns,
			productionForm1PrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert production_form1, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(productionForm1PrimaryKeyColumns))
			copy(conflict, productionForm1PrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"production_form1\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(productionForm1Type, productionForm1Mapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(productionForm1Type, productionForm1Mapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert production_form1")
	}

	if !cached {
		productionForm1UpsertCacheMut.Lock()
		productionForm1UpsertCache[key] = cache
		productionForm1UpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ProductionForm1 record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ProductionForm1) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ProductionForm1 record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProductionForm1) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ProductionForm1 provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), productionForm1PrimaryKeyMapping)
	sql := "DELETE FROM \"production_form1\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from production_form1")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for production_form1")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q productionForm1Query) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q productionForm1Query) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no productionForm1Query provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from production_form1")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_form1")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ProductionForm1Slice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProductionForm1Slice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(productionForm1BeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionForm1PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"production_form1\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productionForm1PrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from productionForm1 slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for production_form1")
	}

	if len(productionForm1AfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ProductionForm1) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ProductionForm1 provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProductionForm1) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProductionForm1(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionForm1Slice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ProductionForm1Slice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProductionForm1Slice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProductionForm1Slice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), productionForm1PrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"production_form1\".* FROM \"production_form1\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, productionForm1PrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ProductionForm1Slice")
	}

	*o = slice

	return nil
}

// ProductionForm1ExistsG checks if the ProductionForm1 row exists.
func ProductionForm1ExistsG(ctx context.Context, iD int64) (bool, error) {
	return ProductionForm1Exists(ctx, boil.GetContextDB(), iD)
}

// ProductionForm1Exists checks if the ProductionForm1 row exists.
func ProductionForm1Exists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"production_form1\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if production_form1 exists")
	}

	return exists, nil
}

// Exists checks if the ProductionForm1 row exists.
func (o *ProductionForm1) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ProductionForm1Exists(ctx, exec, o.ID)
}
