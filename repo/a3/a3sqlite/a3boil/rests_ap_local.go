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

// RestsApLocal is an object representing the database table.
type RestsApLocal struct {
	ID                   int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductFullName      null.String `boil:"product_full_name" json:"product_full_name,omitempty" toml:"product_full_name" yaml:"product_full_name,omitempty"`
	ProductCapacity      null.String `boil:"product_capacity" json:"product_capacity,omitempty" toml:"product_capacity" yaml:"product_capacity,omitempty"`
	ProductAlcVolume     null.String `boil:"product_alc_volume" json:"product_alc_volume,omitempty" toml:"product_alc_volume" yaml:"product_alc_volume,omitempty"`
	ProductAlcCode       null.String `boil:"product_alc_code" json:"product_alc_code,omitempty" toml:"product_alc_code" yaml:"product_alc_code,omitempty"`
	ProductCode          null.String `boil:"product_code" json:"product_code,omitempty" toml:"product_code" yaml:"product_code,omitempty"`
	ProductUnitType      null.String `boil:"product_unit_type" json:"product_unit_type,omitempty" toml:"product_unit_type" yaml:"product_unit_type,omitempty"`
	ProductQuantity      null.String `boil:"product_quantity" json:"product_quantity,omitempty" toml:"product_quantity" yaml:"product_quantity,omitempty"`
	ProductInformF1RegID null.String `boil:"product_inform_f1_reg_id" json:"product_inform_f1_reg_id,omitempty" toml:"product_inform_f1_reg_id" yaml:"product_inform_f1_reg_id,omitempty"`
	ProductInformF2RegID null.String `boil:"product_inform_f2_reg_id" json:"product_inform_f2_reg_id,omitempty" toml:"product_inform_f2_reg_id" yaml:"product_inform_f2_reg_id,omitempty"`
	ProducerType         null.String `boil:"producer_type" json:"producer_type,omitempty" toml:"producer_type" yaml:"producer_type,omitempty"`
	ProducerClientRegID  null.String `boil:"producer_client_reg_id" json:"producer_client_reg_id,omitempty" toml:"producer_client_reg_id" yaml:"producer_client_reg_id,omitempty"`
	ProducerInn          null.String `boil:"producer_inn" json:"producer_inn,omitempty" toml:"producer_inn" yaml:"producer_inn,omitempty"`
	ProducerKPP          null.String `boil:"producer_kpp" json:"producer_kpp,omitempty" toml:"producer_kpp" yaml:"producer_kpp,omitempty"`
	ProducerFullName     null.String `boil:"producer_full_name" json:"producer_full_name,omitempty" toml:"producer_full_name" yaml:"producer_full_name,omitempty"`
	ProducerShortName    null.String `boil:"producer_short_name" json:"producer_short_name,omitempty" toml:"producer_short_name" yaml:"producer_short_name,omitempty"`
	ProducerCountryCode  null.String `boil:"producer_country_code" json:"producer_country_code,omitempty" toml:"producer_country_code" yaml:"producer_country_code,omitempty"`
	ProducerRegionCode   null.String `boil:"producer_region_code" json:"producer_region_code,omitempty" toml:"producer_region_code" yaml:"producer_region_code,omitempty"`
	ProducerDescription  null.String `boil:"producer_description" json:"producer_description,omitempty" toml:"producer_description" yaml:"producer_description,omitempty"`

	R *restsApLocalR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L restsApLocalL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RestsApLocalColumns = struct {
	ID                   string
	ProductFullName      string
	ProductCapacity      string
	ProductAlcVolume     string
	ProductAlcCode       string
	ProductCode          string
	ProductUnitType      string
	ProductQuantity      string
	ProductInformF1RegID string
	ProductInformF2RegID string
	ProducerType         string
	ProducerClientRegID  string
	ProducerInn          string
	ProducerKPP          string
	ProducerFullName     string
	ProducerShortName    string
	ProducerCountryCode  string
	ProducerRegionCode   string
	ProducerDescription  string
}{
	ID:                   "id",
	ProductFullName:      "product_full_name",
	ProductCapacity:      "product_capacity",
	ProductAlcVolume:     "product_alc_volume",
	ProductAlcCode:       "product_alc_code",
	ProductCode:          "product_code",
	ProductUnitType:      "product_unit_type",
	ProductQuantity:      "product_quantity",
	ProductInformF1RegID: "product_inform_f1_reg_id",
	ProductInformF2RegID: "product_inform_f2_reg_id",
	ProducerType:         "producer_type",
	ProducerClientRegID:  "producer_client_reg_id",
	ProducerInn:          "producer_inn",
	ProducerKPP:          "producer_kpp",
	ProducerFullName:     "producer_full_name",
	ProducerShortName:    "producer_short_name",
	ProducerCountryCode:  "producer_country_code",
	ProducerRegionCode:   "producer_region_code",
	ProducerDescription:  "producer_description",
}

var RestsApLocalTableColumns = struct {
	ID                   string
	ProductFullName      string
	ProductCapacity      string
	ProductAlcVolume     string
	ProductAlcCode       string
	ProductCode          string
	ProductUnitType      string
	ProductQuantity      string
	ProductInformF1RegID string
	ProductInformF2RegID string
	ProducerType         string
	ProducerClientRegID  string
	ProducerInn          string
	ProducerKPP          string
	ProducerFullName     string
	ProducerShortName    string
	ProducerCountryCode  string
	ProducerRegionCode   string
	ProducerDescription  string
}{
	ID:                   "rests_ap_local.id",
	ProductFullName:      "rests_ap_local.product_full_name",
	ProductCapacity:      "rests_ap_local.product_capacity",
	ProductAlcVolume:     "rests_ap_local.product_alc_volume",
	ProductAlcCode:       "rests_ap_local.product_alc_code",
	ProductCode:          "rests_ap_local.product_code",
	ProductUnitType:      "rests_ap_local.product_unit_type",
	ProductQuantity:      "rests_ap_local.product_quantity",
	ProductInformF1RegID: "rests_ap_local.product_inform_f1_reg_id",
	ProductInformF2RegID: "rests_ap_local.product_inform_f2_reg_id",
	ProducerType:         "rests_ap_local.producer_type",
	ProducerClientRegID:  "rests_ap_local.producer_client_reg_id",
	ProducerInn:          "rests_ap_local.producer_inn",
	ProducerKPP:          "rests_ap_local.producer_kpp",
	ProducerFullName:     "rests_ap_local.producer_full_name",
	ProducerShortName:    "rests_ap_local.producer_short_name",
	ProducerCountryCode:  "rests_ap_local.producer_country_code",
	ProducerRegionCode:   "rests_ap_local.producer_region_code",
	ProducerDescription:  "rests_ap_local.producer_description",
}

// Generated where

var RestsApLocalWhere = struct {
	ID                   whereHelperint64
	ProductFullName      whereHelpernull_String
	ProductCapacity      whereHelpernull_String
	ProductAlcVolume     whereHelpernull_String
	ProductAlcCode       whereHelpernull_String
	ProductCode          whereHelpernull_String
	ProductUnitType      whereHelpernull_String
	ProductQuantity      whereHelpernull_String
	ProductInformF1RegID whereHelpernull_String
	ProductInformF2RegID whereHelpernull_String
	ProducerType         whereHelpernull_String
	ProducerClientRegID  whereHelpernull_String
	ProducerInn          whereHelpernull_String
	ProducerKPP          whereHelpernull_String
	ProducerFullName     whereHelpernull_String
	ProducerShortName    whereHelpernull_String
	ProducerCountryCode  whereHelpernull_String
	ProducerRegionCode   whereHelpernull_String
	ProducerDescription  whereHelpernull_String
}{
	ID:                   whereHelperint64{field: "\"rests_ap_local\".\"id\""},
	ProductFullName:      whereHelpernull_String{field: "\"rests_ap_local\".\"product_full_name\""},
	ProductCapacity:      whereHelpernull_String{field: "\"rests_ap_local\".\"product_capacity\""},
	ProductAlcVolume:     whereHelpernull_String{field: "\"rests_ap_local\".\"product_alc_volume\""},
	ProductAlcCode:       whereHelpernull_String{field: "\"rests_ap_local\".\"product_alc_code\""},
	ProductCode:          whereHelpernull_String{field: "\"rests_ap_local\".\"product_code\""},
	ProductUnitType:      whereHelpernull_String{field: "\"rests_ap_local\".\"product_unit_type\""},
	ProductQuantity:      whereHelpernull_String{field: "\"rests_ap_local\".\"product_quantity\""},
	ProductInformF1RegID: whereHelpernull_String{field: "\"rests_ap_local\".\"product_inform_f1_reg_id\""},
	ProductInformF2RegID: whereHelpernull_String{field: "\"rests_ap_local\".\"product_inform_f2_reg_id\""},
	ProducerType:         whereHelpernull_String{field: "\"rests_ap_local\".\"producer_type\""},
	ProducerClientRegID:  whereHelpernull_String{field: "\"rests_ap_local\".\"producer_client_reg_id\""},
	ProducerInn:          whereHelpernull_String{field: "\"rests_ap_local\".\"producer_inn\""},
	ProducerKPP:          whereHelpernull_String{field: "\"rests_ap_local\".\"producer_kpp\""},
	ProducerFullName:     whereHelpernull_String{field: "\"rests_ap_local\".\"producer_full_name\""},
	ProducerShortName:    whereHelpernull_String{field: "\"rests_ap_local\".\"producer_short_name\""},
	ProducerCountryCode:  whereHelpernull_String{field: "\"rests_ap_local\".\"producer_country_code\""},
	ProducerRegionCode:   whereHelpernull_String{field: "\"rests_ap_local\".\"producer_region_code\""},
	ProducerDescription:  whereHelpernull_String{field: "\"rests_ap_local\".\"producer_description\""},
}

// RestsApLocalRels is where relationship names are stored.
var RestsApLocalRels = struct {
}{}

// restsApLocalR is where relationships are stored.
type restsApLocalR struct {
}

// NewStruct creates a new relationship struct
func (*restsApLocalR) NewStruct() *restsApLocalR {
	return &restsApLocalR{}
}

// restsApLocalL is where Load methods for each relationship are stored.
type restsApLocalL struct{}

var (
	restsApLocalAllColumns            = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type", "product_quantity", "product_inform_f1_reg_id", "product_inform_f2_reg_id", "producer_type", "producer_client_reg_id", "producer_inn", "producer_kpp", "producer_full_name", "producer_short_name", "producer_country_code", "producer_region_code", "producer_description"}
	restsApLocalColumnsWithoutDefault = []string{}
	restsApLocalColumnsWithDefault    = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type", "product_quantity", "product_inform_f1_reg_id", "product_inform_f2_reg_id", "producer_type", "producer_client_reg_id", "producer_inn", "producer_kpp", "producer_full_name", "producer_short_name", "producer_country_code", "producer_region_code", "producer_description"}
	restsApLocalPrimaryKeyColumns     = []string{"id"}
	restsApLocalGeneratedColumns      = []string{"id"}
)

type (
	// RestsApLocalSlice is an alias for a slice of pointers to RestsApLocal.
	// This should almost always be used instead of []RestsApLocal.
	RestsApLocalSlice []*RestsApLocal
	// RestsApLocalHook is the signature for custom RestsApLocal hook methods
	RestsApLocalHook func(context.Context, boil.ContextExecutor, *RestsApLocal) error

	restsApLocalQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	restsApLocalType                 = reflect.TypeOf(&RestsApLocal{})
	restsApLocalMapping              = queries.MakeStructMapping(restsApLocalType)
	restsApLocalPrimaryKeyMapping, _ = queries.BindMapping(restsApLocalType, restsApLocalMapping, restsApLocalPrimaryKeyColumns)
	restsApLocalInsertCacheMut       sync.RWMutex
	restsApLocalInsertCache          = make(map[string]insertCache)
	restsApLocalUpdateCacheMut       sync.RWMutex
	restsApLocalUpdateCache          = make(map[string]updateCache)
	restsApLocalUpsertCacheMut       sync.RWMutex
	restsApLocalUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var restsApLocalAfterSelectMu sync.Mutex
var restsApLocalAfterSelectHooks []RestsApLocalHook

var restsApLocalBeforeInsertMu sync.Mutex
var restsApLocalBeforeInsertHooks []RestsApLocalHook
var restsApLocalAfterInsertMu sync.Mutex
var restsApLocalAfterInsertHooks []RestsApLocalHook

var restsApLocalBeforeUpdateMu sync.Mutex
var restsApLocalBeforeUpdateHooks []RestsApLocalHook
var restsApLocalAfterUpdateMu sync.Mutex
var restsApLocalAfterUpdateHooks []RestsApLocalHook

var restsApLocalBeforeDeleteMu sync.Mutex
var restsApLocalBeforeDeleteHooks []RestsApLocalHook
var restsApLocalAfterDeleteMu sync.Mutex
var restsApLocalAfterDeleteHooks []RestsApLocalHook

var restsApLocalBeforeUpsertMu sync.Mutex
var restsApLocalBeforeUpsertHooks []RestsApLocalHook
var restsApLocalAfterUpsertMu sync.Mutex
var restsApLocalAfterUpsertHooks []RestsApLocalHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RestsApLocal) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RestsApLocal) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RestsApLocal) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RestsApLocal) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RestsApLocal) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RestsApLocal) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RestsApLocal) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RestsApLocal) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RestsApLocal) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range restsApLocalAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRestsApLocalHook registers your hook function for all future operations.
func AddRestsApLocalHook(hookPoint boil.HookPoint, restsApLocalHook RestsApLocalHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		restsApLocalAfterSelectMu.Lock()
		restsApLocalAfterSelectHooks = append(restsApLocalAfterSelectHooks, restsApLocalHook)
		restsApLocalAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		restsApLocalBeforeInsertMu.Lock()
		restsApLocalBeforeInsertHooks = append(restsApLocalBeforeInsertHooks, restsApLocalHook)
		restsApLocalBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		restsApLocalAfterInsertMu.Lock()
		restsApLocalAfterInsertHooks = append(restsApLocalAfterInsertHooks, restsApLocalHook)
		restsApLocalAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		restsApLocalBeforeUpdateMu.Lock()
		restsApLocalBeforeUpdateHooks = append(restsApLocalBeforeUpdateHooks, restsApLocalHook)
		restsApLocalBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		restsApLocalAfterUpdateMu.Lock()
		restsApLocalAfterUpdateHooks = append(restsApLocalAfterUpdateHooks, restsApLocalHook)
		restsApLocalAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		restsApLocalBeforeDeleteMu.Lock()
		restsApLocalBeforeDeleteHooks = append(restsApLocalBeforeDeleteHooks, restsApLocalHook)
		restsApLocalBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		restsApLocalAfterDeleteMu.Lock()
		restsApLocalAfterDeleteHooks = append(restsApLocalAfterDeleteHooks, restsApLocalHook)
		restsApLocalAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		restsApLocalBeforeUpsertMu.Lock()
		restsApLocalBeforeUpsertHooks = append(restsApLocalBeforeUpsertHooks, restsApLocalHook)
		restsApLocalBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		restsApLocalAfterUpsertMu.Lock()
		restsApLocalAfterUpsertHooks = append(restsApLocalAfterUpsertHooks, restsApLocalHook)
		restsApLocalAfterUpsertMu.Unlock()
	}
}

// OneG returns a single restsApLocal record from the query using the global executor.
func (q restsApLocalQuery) OneG(ctx context.Context) (*RestsApLocal, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single restsApLocal record from the query.
func (q restsApLocalQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RestsApLocal, error) {
	o := &RestsApLocal{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for rests_ap_local")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all RestsApLocal records from the query using the global executor.
func (q restsApLocalQuery) AllG(ctx context.Context) (RestsApLocalSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all RestsApLocal records from the query.
func (q restsApLocalQuery) All(ctx context.Context, exec boil.ContextExecutor) (RestsApLocalSlice, error) {
	var o []*RestsApLocal

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to RestsApLocal slice")
	}

	if len(restsApLocalAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all RestsApLocal records in the query using the global executor
func (q restsApLocalQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all RestsApLocal records in the query.
func (q restsApLocalQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count rests_ap_local rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q restsApLocalQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q restsApLocalQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if rests_ap_local exists")
	}

	return count > 0, nil
}

// RestsApLocals retrieves all the records using an executor.
func RestsApLocals(mods ...qm.QueryMod) restsApLocalQuery {
	mods = append(mods, qm.From("\"rests_ap_local\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"rests_ap_local\".*"})
	}

	return restsApLocalQuery{q}
}

// FindRestsApLocalG retrieves a single record by ID.
func FindRestsApLocalG(ctx context.Context, iD int64, selectCols ...string) (*RestsApLocal, error) {
	return FindRestsApLocal(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindRestsApLocal retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRestsApLocal(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RestsApLocal, error) {
	restsApLocalObj := &RestsApLocal{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rests_ap_local\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, restsApLocalObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from rests_ap_local")
	}

	if err = restsApLocalObj.doAfterSelectHooks(ctx, exec); err != nil {
		return restsApLocalObj, err
	}

	return restsApLocalObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *RestsApLocal) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RestsApLocal) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no rests_ap_local provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restsApLocalColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	restsApLocalInsertCacheMut.RLock()
	cache, cached := restsApLocalInsertCache[key]
	restsApLocalInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			restsApLocalAllColumns,
			restsApLocalColumnsWithDefault,
			restsApLocalColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, restsApLocalGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(restsApLocalType, restsApLocalMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(restsApLocalType, restsApLocalMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rests_ap_local\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rests_ap_local\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into rests_ap_local")
	}

	if !cached {
		restsApLocalInsertCacheMut.Lock()
		restsApLocalInsertCache[key] = cache
		restsApLocalInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single RestsApLocal record using the global executor.
// See Update for more documentation.
func (o *RestsApLocal) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the RestsApLocal.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RestsApLocal) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	restsApLocalUpdateCacheMut.RLock()
	cache, cached := restsApLocalUpdateCache[key]
	restsApLocalUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			restsApLocalAllColumns,
			restsApLocalPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, restsApLocalGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update rests_ap_local, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rests_ap_local\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, restsApLocalPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(restsApLocalType, restsApLocalMapping, append(wl, restsApLocalPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update rests_ap_local row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for rests_ap_local")
	}

	if !cached {
		restsApLocalUpdateCacheMut.Lock()
		restsApLocalUpdateCache[key] = cache
		restsApLocalUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q restsApLocalQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q restsApLocalQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for rests_ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for rests_ap_local")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o RestsApLocalSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RestsApLocalSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restsApLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rests_ap_local\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, restsApLocalPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in restsApLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all restsApLocal")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *RestsApLocal) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RestsApLocal) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no rests_ap_local provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(restsApLocalColumnsWithDefault, o)

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

	restsApLocalUpsertCacheMut.RLock()
	cache, cached := restsApLocalUpsertCache[key]
	restsApLocalUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			restsApLocalAllColumns,
			restsApLocalColumnsWithDefault,
			restsApLocalColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			restsApLocalAllColumns,
			restsApLocalPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert rests_ap_local, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(restsApLocalPrimaryKeyColumns))
			copy(conflict, restsApLocalPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"rests_ap_local\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(restsApLocalType, restsApLocalMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(restsApLocalType, restsApLocalMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert rests_ap_local")
	}

	if !cached {
		restsApLocalUpsertCacheMut.Lock()
		restsApLocalUpsertCache[key] = cache
		restsApLocalUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single RestsApLocal record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *RestsApLocal) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single RestsApLocal record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RestsApLocal) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no RestsApLocal provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), restsApLocalPrimaryKeyMapping)
	sql := "DELETE FROM \"rests_ap_local\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from rests_ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for rests_ap_local")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q restsApLocalQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q restsApLocalQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no restsApLocalQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from rests_ap_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for rests_ap_local")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o RestsApLocalSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RestsApLocalSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(restsApLocalBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restsApLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rests_ap_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, restsApLocalPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from restsApLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for rests_ap_local")
	}

	if len(restsApLocalAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *RestsApLocal) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no RestsApLocal provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RestsApLocal) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRestsApLocal(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RestsApLocalSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty RestsApLocalSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RestsApLocalSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RestsApLocalSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), restsApLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rests_ap_local\".* FROM \"rests_ap_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, restsApLocalPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in RestsApLocalSlice")
	}

	*o = slice

	return nil
}

// RestsApLocalExistsG checks if the RestsApLocal row exists.
func RestsApLocalExistsG(ctx context.Context, iD int64) (bool, error) {
	return RestsApLocalExists(ctx, boil.GetContextDB(), iD)
}

// RestsApLocalExists checks if the RestsApLocal row exists.
func RestsApLocalExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rests_ap_local\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if rests_ap_local exists")
	}

	return exists, nil
}

// Exists checks if the RestsApLocal row exists.
func (o *RestsApLocal) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RestsApLocalExists(ctx, exec, o.ID)
}
