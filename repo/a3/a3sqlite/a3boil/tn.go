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

// TN is an object representing the database table.
type TN struct {
	ID                   int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTTN                null.Int64  `boil:"id_ttn" json:"id_ttn,omitempty" toml:"id_ttn" yaml:"id_ttn,omitempty"`
	CreateDate           null.String `boil:"create_date" json:"create_date,omitempty" toml:"create_date" yaml:"create_date,omitempty"`
	DocNumber            null.String `boil:"doc_number" json:"doc_number,omitempty" toml:"doc_number" yaml:"doc_number,omitempty"`
	DocDate              null.String `boil:"doc_date" json:"doc_date,omitempty" toml:"doc_date" yaml:"doc_date,omitempty"`
	DocQuantity          null.String `boil:"doc_quantity" json:"doc_quantity,omitempty" toml:"doc_quantity" yaml:"doc_quantity,omitempty"`
	DocOwnership         null.Int64  `boil:"doc_ownership" json:"doc_ownership,omitempty" toml:"doc_ownership" yaml:"doc_ownership,omitempty"`
	DocRegID             null.String `boil:"doc_reg_id" json:"doc_reg_id,omitempty" toml:"doc_reg_id" yaml:"doc_reg_id,omitempty"`
	TranType             null.String `boil:"tran_type" json:"tran_type,omitempty" toml:"tran_type" yaml:"tran_type,omitempty"`
	TransportCompany     null.String `boil:"transport_company" json:"transport_company,omitempty" toml:"transport_company" yaml:"transport_company,omitempty"`
	TransportCar         null.String `boil:"transport_car" json:"transport_car,omitempty" toml:"transport_car" yaml:"transport_car,omitempty"`
	TransportCustomer    null.String `boil:"transport_customer" json:"transport_customer,omitempty" toml:"transport_customer" yaml:"transport_customer,omitempty"`
	TransportDriver      null.String `boil:"transport_driver" json:"transport_driver,omitempty" toml:"transport_driver" yaml:"transport_driver,omitempty"`
	TransportTrailer     null.String `boil:"transport_trailer" json:"transport_trailer,omitempty" toml:"transport_trailer" yaml:"transport_trailer,omitempty"`
	TransportForwarder   null.String `boil:"transport_forwarder" json:"transport_forwarder,omitempty" toml:"transport_forwarder" yaml:"transport_forwarder,omitempty"`
	TransportLoadPoint   null.String `boil:"transport_load_point" json:"transport_load_point,omitempty" toml:"transport_load_point" yaml:"transport_load_point,omitempty"`
	TransportUnloadPoint null.String `boil:"transport_unload_point" json:"transport_unload_point,omitempty" toml:"transport_unload_point" yaml:"transport_unload_point,omitempty"`
	TransportRedirect    null.String `boil:"transport_redirect" json:"transport_redirect,omitempty" toml:"transport_redirect" yaml:"transport_redirect,omitempty"`
	Version              null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	State                null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Status               null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID              null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`
	XML                  null.String `boil:"xml" json:"xml,omitempty" toml:"xml" yaml:"xml,omitempty"`

	R *tnR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tnL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TNColumns = struct {
	ID                   string
	IDTTN                string
	CreateDate           string
	DocNumber            string
	DocDate              string
	DocQuantity          string
	DocOwnership         string
	DocRegID             string
	TranType             string
	TransportCompany     string
	TransportCar         string
	TransportCustomer    string
	TransportDriver      string
	TransportTrailer     string
	TransportForwarder   string
	TransportLoadPoint   string
	TransportUnloadPoint string
	TransportRedirect    string
	Version              string
	State                string
	Status               string
	ReplyID              string
	XML                  string
}{
	ID:                   "id",
	IDTTN:                "id_ttn",
	CreateDate:           "create_date",
	DocNumber:            "doc_number",
	DocDate:              "doc_date",
	DocQuantity:          "doc_quantity",
	DocOwnership:         "doc_ownership",
	DocRegID:             "doc_reg_id",
	TranType:             "tran_type",
	TransportCompany:     "transport_company",
	TransportCar:         "transport_car",
	TransportCustomer:    "transport_customer",
	TransportDriver:      "transport_driver",
	TransportTrailer:     "transport_trailer",
	TransportForwarder:   "transport_forwarder",
	TransportLoadPoint:   "transport_load_point",
	TransportUnloadPoint: "transport_unload_point",
	TransportRedirect:    "transport_redirect",
	Version:              "version",
	State:                "state",
	Status:               "status",
	ReplyID:              "reply_id",
	XML:                  "xml",
}

var TNTableColumns = struct {
	ID                   string
	IDTTN                string
	CreateDate           string
	DocNumber            string
	DocDate              string
	DocQuantity          string
	DocOwnership         string
	DocRegID             string
	TranType             string
	TransportCompany     string
	TransportCar         string
	TransportCustomer    string
	TransportDriver      string
	TransportTrailer     string
	TransportForwarder   string
	TransportLoadPoint   string
	TransportUnloadPoint string
	TransportRedirect    string
	Version              string
	State                string
	Status               string
	ReplyID              string
	XML                  string
}{
	ID:                   "tn.id",
	IDTTN:                "tn.id_ttn",
	CreateDate:           "tn.create_date",
	DocNumber:            "tn.doc_number",
	DocDate:              "tn.doc_date",
	DocQuantity:          "tn.doc_quantity",
	DocOwnership:         "tn.doc_ownership",
	DocRegID:             "tn.doc_reg_id",
	TranType:             "tn.tran_type",
	TransportCompany:     "tn.transport_company",
	TransportCar:         "tn.transport_car",
	TransportCustomer:    "tn.transport_customer",
	TransportDriver:      "tn.transport_driver",
	TransportTrailer:     "tn.transport_trailer",
	TransportForwarder:   "tn.transport_forwarder",
	TransportLoadPoint:   "tn.transport_load_point",
	TransportUnloadPoint: "tn.transport_unload_point",
	TransportRedirect:    "tn.transport_redirect",
	Version:              "tn.version",
	State:                "tn.state",
	Status:               "tn.status",
	ReplyID:              "tn.reply_id",
	XML:                  "tn.xml",
}

// Generated where

var TNWhere = struct {
	ID                   whereHelperint64
	IDTTN                whereHelpernull_Int64
	CreateDate           whereHelpernull_String
	DocNumber            whereHelpernull_String
	DocDate              whereHelpernull_String
	DocQuantity          whereHelpernull_String
	DocOwnership         whereHelpernull_Int64
	DocRegID             whereHelpernull_String
	TranType             whereHelpernull_String
	TransportCompany     whereHelpernull_String
	TransportCar         whereHelpernull_String
	TransportCustomer    whereHelpernull_String
	TransportDriver      whereHelpernull_String
	TransportTrailer     whereHelpernull_String
	TransportForwarder   whereHelpernull_String
	TransportLoadPoint   whereHelpernull_String
	TransportUnloadPoint whereHelpernull_String
	TransportRedirect    whereHelpernull_String
	Version              whereHelpernull_String
	State                whereHelpernull_String
	Status               whereHelpernull_String
	ReplyID              whereHelpernull_String
	XML                  whereHelpernull_String
}{
	ID:                   whereHelperint64{field: "\"tn\".\"id\""},
	IDTTN:                whereHelpernull_Int64{field: "\"tn\".\"id_ttn\""},
	CreateDate:           whereHelpernull_String{field: "\"tn\".\"create_date\""},
	DocNumber:            whereHelpernull_String{field: "\"tn\".\"doc_number\""},
	DocDate:              whereHelpernull_String{field: "\"tn\".\"doc_date\""},
	DocQuantity:          whereHelpernull_String{field: "\"tn\".\"doc_quantity\""},
	DocOwnership:         whereHelpernull_Int64{field: "\"tn\".\"doc_ownership\""},
	DocRegID:             whereHelpernull_String{field: "\"tn\".\"doc_reg_id\""},
	TranType:             whereHelpernull_String{field: "\"tn\".\"tran_type\""},
	TransportCompany:     whereHelpernull_String{field: "\"tn\".\"transport_company\""},
	TransportCar:         whereHelpernull_String{field: "\"tn\".\"transport_car\""},
	TransportCustomer:    whereHelpernull_String{field: "\"tn\".\"transport_customer\""},
	TransportDriver:      whereHelpernull_String{field: "\"tn\".\"transport_driver\""},
	TransportTrailer:     whereHelpernull_String{field: "\"tn\".\"transport_trailer\""},
	TransportForwarder:   whereHelpernull_String{field: "\"tn\".\"transport_forwarder\""},
	TransportLoadPoint:   whereHelpernull_String{field: "\"tn\".\"transport_load_point\""},
	TransportUnloadPoint: whereHelpernull_String{field: "\"tn\".\"transport_unload_point\""},
	TransportRedirect:    whereHelpernull_String{field: "\"tn\".\"transport_redirect\""},
	Version:              whereHelpernull_String{field: "\"tn\".\"version\""},
	State:                whereHelpernull_String{field: "\"tn\".\"state\""},
	Status:               whereHelpernull_String{field: "\"tn\".\"status\""},
	ReplyID:              whereHelpernull_String{field: "\"tn\".\"reply_id\""},
	XML:                  whereHelpernull_String{field: "\"tn\".\"xml\""},
}

// TNRels is where relationship names are stored.
var TNRels = struct {
}{}

// tnR is where relationships are stored.
type tnR struct {
}

// NewStruct creates a new relationship struct
func (*tnR) NewStruct() *tnR {
	return &tnR{}
}

// tnL is where Load methods for each relationship are stored.
type tnL struct{}

var (
	tnAllColumns            = []string{"id", "id_ttn", "create_date", "doc_number", "doc_date", "doc_quantity", "doc_ownership", "doc_reg_id", "tran_type", "transport_company", "transport_car", "transport_customer", "transport_driver", "transport_trailer", "transport_forwarder", "transport_load_point", "transport_unload_point", "transport_redirect", "version", "state", "status", "reply_id", "xml"}
	tnColumnsWithoutDefault = []string{}
	tnColumnsWithDefault    = []string{"id", "id_ttn", "create_date", "doc_number", "doc_date", "doc_quantity", "doc_ownership", "doc_reg_id", "tran_type", "transport_company", "transport_car", "transport_customer", "transport_driver", "transport_trailer", "transport_forwarder", "transport_load_point", "transport_unload_point", "transport_redirect", "version", "state", "status", "reply_id", "xml"}
	tnPrimaryKeyColumns     = []string{"id"}
	tnGeneratedColumns      = []string{"id"}
)

type (
	// TNSlice is an alias for a slice of pointers to TN.
	// This should almost always be used instead of []TN.
	TNSlice []*TN
	// TNHook is the signature for custom TN hook methods
	TNHook func(context.Context, boil.ContextExecutor, *TN) error

	tnQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tnType                 = reflect.TypeOf(&TN{})
	tnMapping              = queries.MakeStructMapping(tnType)
	tnPrimaryKeyMapping, _ = queries.BindMapping(tnType, tnMapping, tnPrimaryKeyColumns)
	tnInsertCacheMut       sync.RWMutex
	tnInsertCache          = make(map[string]insertCache)
	tnUpdateCacheMut       sync.RWMutex
	tnUpdateCache          = make(map[string]updateCache)
	tnUpsertCacheMut       sync.RWMutex
	tnUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tnAfterSelectMu sync.Mutex
var tnAfterSelectHooks []TNHook

var tnBeforeInsertMu sync.Mutex
var tnBeforeInsertHooks []TNHook
var tnAfterInsertMu sync.Mutex
var tnAfterInsertHooks []TNHook

var tnBeforeUpdateMu sync.Mutex
var tnBeforeUpdateHooks []TNHook
var tnAfterUpdateMu sync.Mutex
var tnAfterUpdateHooks []TNHook

var tnBeforeDeleteMu sync.Mutex
var tnBeforeDeleteHooks []TNHook
var tnAfterDeleteMu sync.Mutex
var tnAfterDeleteHooks []TNHook

var tnBeforeUpsertMu sync.Mutex
var tnBeforeUpsertHooks []TNHook
var tnAfterUpsertMu sync.Mutex
var tnAfterUpsertHooks []TNHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TN) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TN) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TN) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TN) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TN) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TN) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TN) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TN) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TN) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tnAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTNHook registers your hook function for all future operations.
func AddTNHook(hookPoint boil.HookPoint, tnHook TNHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tnAfterSelectMu.Lock()
		tnAfterSelectHooks = append(tnAfterSelectHooks, tnHook)
		tnAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tnBeforeInsertMu.Lock()
		tnBeforeInsertHooks = append(tnBeforeInsertHooks, tnHook)
		tnBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tnAfterInsertMu.Lock()
		tnAfterInsertHooks = append(tnAfterInsertHooks, tnHook)
		tnAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tnBeforeUpdateMu.Lock()
		tnBeforeUpdateHooks = append(tnBeforeUpdateHooks, tnHook)
		tnBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tnAfterUpdateMu.Lock()
		tnAfterUpdateHooks = append(tnAfterUpdateHooks, tnHook)
		tnAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tnBeforeDeleteMu.Lock()
		tnBeforeDeleteHooks = append(tnBeforeDeleteHooks, tnHook)
		tnBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tnAfterDeleteMu.Lock()
		tnAfterDeleteHooks = append(tnAfterDeleteHooks, tnHook)
		tnAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tnBeforeUpsertMu.Lock()
		tnBeforeUpsertHooks = append(tnBeforeUpsertHooks, tnHook)
		tnBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tnAfterUpsertMu.Lock()
		tnAfterUpsertHooks = append(tnAfterUpsertHooks, tnHook)
		tnAfterUpsertMu.Unlock()
	}
}

// OneG returns a single tn record from the query using the global executor.
func (q tnQuery) OneG(ctx context.Context) (*TN, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single tn record from the query.
func (q tnQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TN, error) {
	o := &TN{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for tn")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TN records from the query using the global executor.
func (q tnQuery) AllG(ctx context.Context) (TNSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TN records from the query.
func (q tnQuery) All(ctx context.Context, exec boil.ContextExecutor) (TNSlice, error) {
	var o []*TN

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TN slice")
	}

	if len(tnAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TN records in the query using the global executor
func (q tnQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TN records in the query.
func (q tnQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count tn rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q tnQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q tnQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if tn exists")
	}

	return count > 0, nil
}

// TNS retrieves all the records using an executor.
func TNS(mods ...qm.QueryMod) tnQuery {
	mods = append(mods, qm.From("\"tn\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"tn\".*"})
	}

	return tnQuery{q}
}

// FindTNG retrieves a single record by ID.
func FindTNG(ctx context.Context, iD int64, selectCols ...string) (*TN, error) {
	return FindTN(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTN retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTN(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TN, error) {
	tnObj := &TN{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tn\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tnObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from tn")
	}

	if err = tnObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tnObj, err
	}

	return tnObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TN) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TN) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tnInsertCacheMut.RLock()
	cache, cached := tnInsertCache[key]
	tnInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tnAllColumns,
			tnColumnsWithDefault,
			tnColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, tnGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(tnType, tnMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tnType, tnMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tn\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tn\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into tn")
	}

	if !cached {
		tnInsertCacheMut.Lock()
		tnInsertCache[key] = cache
		tnInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TN record using the global executor.
// See Update for more documentation.
func (o *TN) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TN.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TN) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tnUpdateCacheMut.RLock()
	cache, cached := tnUpdateCache[key]
	tnUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tnAllColumns,
			tnPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, tnGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update tn, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tn\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, tnPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tnType, tnMapping, append(wl, tnPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update tn row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for tn")
	}

	if !cached {
		tnUpdateCacheMut.Lock()
		tnUpdateCache[key] = cache
		tnUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tnQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tnQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for tn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for tn")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TNSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TNSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tn\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tnPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in tn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all tn")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TN) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TN) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no tn provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tnColumnsWithDefault, o)

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

	tnUpsertCacheMut.RLock()
	cache, cached := tnUpsertCache[key]
	tnUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tnAllColumns,
			tnColumnsWithDefault,
			tnColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tnAllColumns,
			tnPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert tn, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tnPrimaryKeyColumns))
			copy(conflict, tnPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"tn\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tnType, tnMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tnType, tnMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert tn")
	}

	if !cached {
		tnUpsertCacheMut.Lock()
		tnUpsertCache[key] = cache
		tnUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TN record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TN) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TN record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TN) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TN provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tnPrimaryKeyMapping)
	sql := "DELETE FROM \"tn\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from tn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for tn")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q tnQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q tnQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no tnQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TNSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TNSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tnBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tn\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tnPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from tn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for tn")
	}

	if len(tnAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TN) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TN provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TN) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTN(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TNSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TNSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TNSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tn\".* FROM \"tn\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tnPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TNSlice")
	}

	*o = slice

	return nil
}

// TNExistsG checks if the TN row exists.
func TNExistsG(ctx context.Context, iD int64) (bool, error) {
	return TNExists(ctx, boil.GetContextDB(), iD)
}

// TNExists checks if the TN row exists.
func TNExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tn\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if tn exists")
	}

	return exists, nil
}

// Exists checks if the TN row exists.
func (o *TN) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TNExists(ctx, exec, o.ID)
}
