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

// ApLocal is an object representing the database table.
type ApLocal struct {
	ID                  int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductFullName     null.String `boil:"product_full_name" json:"product_full_name,omitempty" toml:"product_full_name" yaml:"product_full_name,omitempty"`
	ProductCapacity     null.String `boil:"product_capacity" json:"product_capacity,omitempty" toml:"product_capacity" yaml:"product_capacity,omitempty"`
	ProductAlcVolume    null.String `boil:"product_alc_volume" json:"product_alc_volume,omitempty" toml:"product_alc_volume" yaml:"product_alc_volume,omitempty"`
	ProductAlcCode      null.String `boil:"product_alc_code" json:"product_alc_code,omitempty" toml:"product_alc_code" yaml:"product_alc_code,omitempty"`
	ProductCode         null.String `boil:"product_code" json:"product_code,omitempty" toml:"product_code" yaml:"product_code,omitempty"`
	ProductUnitType     null.String `boil:"product_unit_type" json:"product_unit_type,omitempty" toml:"product_unit_type" yaml:"product_unit_type,omitempty"`
	ProducerType        null.String `boil:"producer_type" json:"producer_type,omitempty" toml:"producer_type" yaml:"producer_type,omitempty"`
	ProducerClientRegID null.String `boil:"producer_client_reg_id" json:"producer_client_reg_id,omitempty" toml:"producer_client_reg_id" yaml:"producer_client_reg_id,omitempty"`
	ProducerInn         null.String `boil:"producer_inn" json:"producer_inn,omitempty" toml:"producer_inn" yaml:"producer_inn,omitempty"`
	ProducerKPP         null.String `boil:"producer_kpp" json:"producer_kpp,omitempty" toml:"producer_kpp" yaml:"producer_kpp,omitempty"`
	ProducerFullName    null.String `boil:"producer_full_name" json:"producer_full_name,omitempty" toml:"producer_full_name" yaml:"producer_full_name,omitempty"`
	ProducerShortName   null.String `boil:"producer_short_name" json:"producer_short_name,omitempty" toml:"producer_short_name" yaml:"producer_short_name,omitempty"`
	ProducerCountryCode null.String `boil:"producer_country_code" json:"producer_country_code,omitempty" toml:"producer_country_code" yaml:"producer_country_code,omitempty"`
	ProducerRegionCode  null.String `boil:"producer_region_code" json:"producer_region_code,omitempty" toml:"producer_region_code" yaml:"producer_region_code,omitempty"`
	ProducerDescription null.String `boil:"producer_description" json:"producer_description,omitempty" toml:"producer_description" yaml:"producer_description,omitempty"`

	R *apLocalR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L apLocalL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ApLocalColumns = struct {
	ID                  string
	ProductFullName     string
	ProductCapacity     string
	ProductAlcVolume    string
	ProductAlcCode      string
	ProductCode         string
	ProductUnitType     string
	ProducerType        string
	ProducerClientRegID string
	ProducerInn         string
	ProducerKPP         string
	ProducerFullName    string
	ProducerShortName   string
	ProducerCountryCode string
	ProducerRegionCode  string
	ProducerDescription string
}{
	ID:                  "id",
	ProductFullName:     "product_full_name",
	ProductCapacity:     "product_capacity",
	ProductAlcVolume:    "product_alc_volume",
	ProductAlcCode:      "product_alc_code",
	ProductCode:         "product_code",
	ProductUnitType:     "product_unit_type",
	ProducerType:        "producer_type",
	ProducerClientRegID: "producer_client_reg_id",
	ProducerInn:         "producer_inn",
	ProducerKPP:         "producer_kpp",
	ProducerFullName:    "producer_full_name",
	ProducerShortName:   "producer_short_name",
	ProducerCountryCode: "producer_country_code",
	ProducerRegionCode:  "producer_region_code",
	ProducerDescription: "producer_description",
}

var ApLocalTableColumns = struct {
	ID                  string
	ProductFullName     string
	ProductCapacity     string
	ProductAlcVolume    string
	ProductAlcCode      string
	ProductCode         string
	ProductUnitType     string
	ProducerType        string
	ProducerClientRegID string
	ProducerInn         string
	ProducerKPP         string
	ProducerFullName    string
	ProducerShortName   string
	ProducerCountryCode string
	ProducerRegionCode  string
	ProducerDescription string
}{
	ID:                  "ap_local.id",
	ProductFullName:     "ap_local.product_full_name",
	ProductCapacity:     "ap_local.product_capacity",
	ProductAlcVolume:    "ap_local.product_alc_volume",
	ProductAlcCode:      "ap_local.product_alc_code",
	ProductCode:         "ap_local.product_code",
	ProductUnitType:     "ap_local.product_unit_type",
	ProducerType:        "ap_local.producer_type",
	ProducerClientRegID: "ap_local.producer_client_reg_id",
	ProducerInn:         "ap_local.producer_inn",
	ProducerKPP:         "ap_local.producer_kpp",
	ProducerFullName:    "ap_local.producer_full_name",
	ProducerShortName:   "ap_local.producer_short_name",
	ProducerCountryCode: "ap_local.producer_country_code",
	ProducerRegionCode:  "ap_local.producer_region_code",
	ProducerDescription: "ap_local.producer_description",
}

// Generated where

var ApLocalWhere = struct {
	ID                  whereHelperint64
	ProductFullName     whereHelpernull_String
	ProductCapacity     whereHelpernull_String
	ProductAlcVolume    whereHelpernull_String
	ProductAlcCode      whereHelpernull_String
	ProductCode         whereHelpernull_String
	ProductUnitType     whereHelpernull_String
	ProducerType        whereHelpernull_String
	ProducerClientRegID whereHelpernull_String
	ProducerInn         whereHelpernull_String
	ProducerKPP         whereHelpernull_String
	ProducerFullName    whereHelpernull_String
	ProducerShortName   whereHelpernull_String
	ProducerCountryCode whereHelpernull_String
	ProducerRegionCode  whereHelpernull_String
	ProducerDescription whereHelpernull_String
}{
	ID:                  whereHelperint64{field: "\"ap_local\".\"id\""},
	ProductFullName:     whereHelpernull_String{field: "\"ap_local\".\"product_full_name\""},
	ProductCapacity:     whereHelpernull_String{field: "\"ap_local\".\"product_capacity\""},
	ProductAlcVolume:    whereHelpernull_String{field: "\"ap_local\".\"product_alc_volume\""},
	ProductAlcCode:      whereHelpernull_String{field: "\"ap_local\".\"product_alc_code\""},
	ProductCode:         whereHelpernull_String{field: "\"ap_local\".\"product_code\""},
	ProductUnitType:     whereHelpernull_String{field: "\"ap_local\".\"product_unit_type\""},
	ProducerType:        whereHelpernull_String{field: "\"ap_local\".\"producer_type\""},
	ProducerClientRegID: whereHelpernull_String{field: "\"ap_local\".\"producer_client_reg_id\""},
	ProducerInn:         whereHelpernull_String{field: "\"ap_local\".\"producer_inn\""},
	ProducerKPP:         whereHelpernull_String{field: "\"ap_local\".\"producer_kpp\""},
	ProducerFullName:    whereHelpernull_String{field: "\"ap_local\".\"producer_full_name\""},
	ProducerShortName:   whereHelpernull_String{field: "\"ap_local\".\"producer_short_name\""},
	ProducerCountryCode: whereHelpernull_String{field: "\"ap_local\".\"producer_country_code\""},
	ProducerRegionCode:  whereHelpernull_String{field: "\"ap_local\".\"producer_region_code\""},
	ProducerDescription: whereHelpernull_String{field: "\"ap_local\".\"producer_description\""},
}

// ApLocalRels is where relationship names are stored.
var ApLocalRels = struct {
}{}

// apLocalR is where relationships are stored.
type apLocalR struct {
}

// NewStruct creates a new relationship struct
func (*apLocalR) NewStruct() *apLocalR {
	return &apLocalR{}
}

// apLocalL is where Load methods for each relationship are stored.
type apLocalL struct{}

var (
	apLocalAllColumns            = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type", "producer_type", "producer_client_reg_id", "producer_inn", "producer_kpp", "producer_full_name", "producer_short_name", "producer_country_code", "producer_region_code", "producer_description"}
	apLocalColumnsWithoutDefault = []string{}
	apLocalColumnsWithDefault    = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type", "producer_type", "producer_client_reg_id", "producer_inn", "producer_kpp", "producer_full_name", "producer_short_name", "producer_country_code", "producer_region_code", "producer_description"}
	apLocalPrimaryKeyColumns     = []string{"id"}
	apLocalGeneratedColumns      = []string{"id"}
)

type (
	// ApLocalSlice is an alias for a slice of pointers to ApLocal.
	// This should almost always be used instead of []ApLocal.
	ApLocalSlice []*ApLocal
	// ApLocalHook is the signature for custom ApLocal hook methods
	ApLocalHook func(context.Context, boil.ContextExecutor, *ApLocal) error

	apLocalQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	apLocalType                 = reflect.TypeOf(&ApLocal{})
	apLocalMapping              = queries.MakeStructMapping(apLocalType)
	apLocalPrimaryKeyMapping, _ = queries.BindMapping(apLocalType, apLocalMapping, apLocalPrimaryKeyColumns)
	apLocalInsertCacheMut       sync.RWMutex
	apLocalInsertCache          = make(map[string]insertCache)
	apLocalUpdateCacheMut       sync.RWMutex
	apLocalUpdateCache          = make(map[string]updateCache)
	apLocalUpsertCacheMut       sync.RWMutex
	apLocalUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var apLocalAfterSelectMu sync.Mutex
var apLocalAfterSelectHooks []ApLocalHook

var apLocalBeforeInsertMu sync.Mutex
var apLocalBeforeInsertHooks []ApLocalHook
var apLocalAfterInsertMu sync.Mutex
var apLocalAfterInsertHooks []ApLocalHook

var apLocalBeforeUpdateMu sync.Mutex
var apLocalBeforeUpdateHooks []ApLocalHook
var apLocalAfterUpdateMu sync.Mutex
var apLocalAfterUpdateHooks []ApLocalHook

var apLocalBeforeDeleteMu sync.Mutex
var apLocalBeforeDeleteHooks []ApLocalHook
var apLocalAfterDeleteMu sync.Mutex
var apLocalAfterDeleteHooks []ApLocalHook

var apLocalBeforeUpsertMu sync.Mutex
var apLocalBeforeUpsertHooks []ApLocalHook
var apLocalAfterUpsertMu sync.Mutex
var apLocalAfterUpsertHooks []ApLocalHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ApLocal) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ApLocal) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ApLocal) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ApLocal) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ApLocal) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ApLocal) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ApLocal) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ApLocal) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ApLocal) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apLocalAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddApLocalHook registers your hook function for all future operations.
func AddApLocalHook(hookPoint boil.HookPoint, apLocalHook ApLocalHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		apLocalAfterSelectMu.Lock()
		apLocalAfterSelectHooks = append(apLocalAfterSelectHooks, apLocalHook)
		apLocalAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		apLocalBeforeInsertMu.Lock()
		apLocalBeforeInsertHooks = append(apLocalBeforeInsertHooks, apLocalHook)
		apLocalBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		apLocalAfterInsertMu.Lock()
		apLocalAfterInsertHooks = append(apLocalAfterInsertHooks, apLocalHook)
		apLocalAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		apLocalBeforeUpdateMu.Lock()
		apLocalBeforeUpdateHooks = append(apLocalBeforeUpdateHooks, apLocalHook)
		apLocalBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		apLocalAfterUpdateMu.Lock()
		apLocalAfterUpdateHooks = append(apLocalAfterUpdateHooks, apLocalHook)
		apLocalAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		apLocalBeforeDeleteMu.Lock()
		apLocalBeforeDeleteHooks = append(apLocalBeforeDeleteHooks, apLocalHook)
		apLocalBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		apLocalAfterDeleteMu.Lock()
		apLocalAfterDeleteHooks = append(apLocalAfterDeleteHooks, apLocalHook)
		apLocalAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		apLocalBeforeUpsertMu.Lock()
		apLocalBeforeUpsertHooks = append(apLocalBeforeUpsertHooks, apLocalHook)
		apLocalBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		apLocalAfterUpsertMu.Lock()
		apLocalAfterUpsertHooks = append(apLocalAfterUpsertHooks, apLocalHook)
		apLocalAfterUpsertMu.Unlock()
	}
}

// OneG returns a single apLocal record from the query using the global executor.
func (q apLocalQuery) OneG(ctx context.Context) (*ApLocal, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single apLocal record from the query.
func (q apLocalQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ApLocal, error) {
	o := &ApLocal{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for ap_local")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all ApLocal records from the query using the global executor.
func (q apLocalQuery) AllG(ctx context.Context) (ApLocalSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ApLocal records from the query.
func (q apLocalQuery) All(ctx context.Context, exec boil.ContextExecutor) (ApLocalSlice, error) {
	var o []*ApLocal

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to ApLocal slice")
	}

	if len(apLocalAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all ApLocal records in the query using the global executor
func (q apLocalQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ApLocal records in the query.
func (q apLocalQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count ap_local rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q apLocalQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q apLocalQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if ap_local exists")
	}

	return count > 0, nil
}

// ApLocals retrieves all the records using an executor.
func ApLocals(mods ...qm.QueryMod) apLocalQuery {
	mods = append(mods, qm.From("\"ap_local\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"ap_local\".*"})
	}

	return apLocalQuery{q}
}

// FindApLocalG retrieves a single record by ID.
func FindApLocalG(ctx context.Context, iD int64, selectCols ...string) (*ApLocal, error) {
	return FindApLocal(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindApLocal retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindApLocal(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ApLocal, error) {
	apLocalObj := &ApLocal{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ap_local\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, apLocalObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from ap_local")
	}

	if err = apLocalObj.doAfterSelectHooks(ctx, exec); err != nil {
		return apLocalObj, err
	}

	return apLocalObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ApLocal) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ApLocal) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ap_local provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(apLocalColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	apLocalInsertCacheMut.RLock()
	cache, cached := apLocalInsertCache[key]
	apLocalInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			apLocalAllColumns,
			apLocalColumnsWithDefault,
			apLocalColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, apLocalGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(apLocalType, apLocalMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(apLocalType, apLocalMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ap_local\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ap_local\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into ap_local")
	}

	if !cached {
		apLocalInsertCacheMut.Lock()
		apLocalInsertCache[key] = cache
		apLocalInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single ApLocal record using the global executor.
// See Update for more documentation.
func (o *ApLocal) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ApLocal.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ApLocal) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	apLocalUpdateCacheMut.RLock()
	cache, cached := apLocalUpdateCache[key]
	apLocalUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			apLocalAllColumns,
			apLocalPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, apLocalGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update ap_local, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ap_local\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, apLocalPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(apLocalType, apLocalMapping, append(wl, apLocalPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update ap_local row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for ap_local")
	}

	if !cached {
		apLocalUpdateCacheMut.Lock()
		apLocalUpdateCache[key] = cache
		apLocalUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q apLocalQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q apLocalQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for ap_local")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ApLocalSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ApLocalSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ap_local\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, apLocalPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in apLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all apLocal")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ApLocal) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ApLocal) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ap_local provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(apLocalColumnsWithDefault, o)

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

	apLocalUpsertCacheMut.RLock()
	cache, cached := apLocalUpsertCache[key]
	apLocalUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			apLocalAllColumns,
			apLocalColumnsWithDefault,
			apLocalColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			apLocalAllColumns,
			apLocalPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert ap_local, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(apLocalPrimaryKeyColumns))
			copy(conflict, apLocalPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"ap_local\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(apLocalType, apLocalMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(apLocalType, apLocalMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert ap_local")
	}

	if !cached {
		apLocalUpsertCacheMut.Lock()
		apLocalUpsertCache[key] = cache
		apLocalUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single ApLocal record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ApLocal) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ApLocal record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ApLocal) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no ApLocal provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), apLocalPrimaryKeyMapping)
	sql := "DELETE FROM \"ap_local\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for ap_local")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q apLocalQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q apLocalQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no apLocalQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ap_local")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ApLocalSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ApLocalSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(apLocalBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ap_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, apLocalPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from apLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ap_local")
	}

	if len(apLocalAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ApLocal) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no ApLocal provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ApLocal) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindApLocal(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ApLocalSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty ApLocalSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ApLocalSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ApLocalSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ap_local\".* FROM \"ap_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, apLocalPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in ApLocalSlice")
	}

	*o = slice

	return nil
}

// ApLocalExistsG checks if the ApLocal row exists.
func ApLocalExistsG(ctx context.Context, iD int64) (bool, error) {
	return ApLocalExists(ctx, boil.GetContextDB(), iD)
}

// ApLocalExists checks if the ApLocal row exists.
func ApLocalExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ap_local\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if ap_local exists")
	}

	return exists, nil
}

// Exists checks if the ApLocal row exists.
func (o *ApLocal) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ApLocalExists(ctx, exec, o.ID)
}
