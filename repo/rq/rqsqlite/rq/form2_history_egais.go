// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package rq

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

// Form2HistoryEgais is an object representing the database table.
type Form2HistoryEgais struct {
	ID               int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	RequestID        int64       `boil:"request_id" json:"request_id" toml:"request_id" yaml:"request_id"`
	InformF2RegID    string      `boil:"inform_f2_reg_id" json:"inform_f2_reg_id" toml:"inform_f2_reg_id" yaml:"inform_f2_reg_id"`
	DocType          null.String `boil:"doc_type" json:"doc_type,omitempty" toml:"doc_type" yaml:"doc_type,omitempty"`
	DocID            null.String `boil:"doc_id" json:"doc_id,omitempty" toml:"doc_id" yaml:"doc_id,omitempty"`
	ProductQuantity  null.String `boil:"product_quantity" json:"product_quantity,omitempty" toml:"product_quantity" yaml:"product_quantity,omitempty"`
	OperationName    null.String `boil:"operation_name" json:"operation_name,omitempty" toml:"operation_name" yaml:"operation_name,omitempty"`
	OperationDate    null.String `boil:"operation_date" json:"operation_date,omitempty" toml:"operation_date" yaml:"operation_date,omitempty"`
	Form2HistoryDate null.String `boil:"form2_history_date" json:"form2_history_date,omitempty" toml:"form2_history_date" yaml:"form2_history_date,omitempty"`

	R *form2HistoryEgaisR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L form2HistoryEgaisL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var Form2HistoryEgaisColumns = struct {
	ID               string
	RequestID        string
	InformF2RegID    string
	DocType          string
	DocID            string
	ProductQuantity  string
	OperationName    string
	OperationDate    string
	Form2HistoryDate string
}{
	ID:               "id",
	RequestID:        "request_id",
	InformF2RegID:    "inform_f2_reg_id",
	DocType:          "doc_type",
	DocID:            "doc_id",
	ProductQuantity:  "product_quantity",
	OperationName:    "operation_name",
	OperationDate:    "operation_date",
	Form2HistoryDate: "form2_history_date",
}

var Form2HistoryEgaisTableColumns = struct {
	ID               string
	RequestID        string
	InformF2RegID    string
	DocType          string
	DocID            string
	ProductQuantity  string
	OperationName    string
	OperationDate    string
	Form2HistoryDate string
}{
	ID:               "form2_history_egais.id",
	RequestID:        "form2_history_egais.request_id",
	InformF2RegID:    "form2_history_egais.inform_f2_reg_id",
	DocType:          "form2_history_egais.doc_type",
	DocID:            "form2_history_egais.doc_id",
	ProductQuantity:  "form2_history_egais.product_quantity",
	OperationName:    "form2_history_egais.operation_name",
	OperationDate:    "form2_history_egais.operation_date",
	Form2HistoryDate: "form2_history_egais.form2_history_date",
}

// Generated where

var Form2HistoryEgaisWhere = struct {
	ID               whereHelperint64
	RequestID        whereHelperint64
	InformF2RegID    whereHelperstring
	DocType          whereHelpernull_String
	DocID            whereHelpernull_String
	ProductQuantity  whereHelpernull_String
	OperationName    whereHelpernull_String
	OperationDate    whereHelpernull_String
	Form2HistoryDate whereHelpernull_String
}{
	ID:               whereHelperint64{field: "\"form2_history_egais\".\"id\""},
	RequestID:        whereHelperint64{field: "\"form2_history_egais\".\"request_id\""},
	InformF2RegID:    whereHelperstring{field: "\"form2_history_egais\".\"inform_f2_reg_id\""},
	DocType:          whereHelpernull_String{field: "\"form2_history_egais\".\"doc_type\""},
	DocID:            whereHelpernull_String{field: "\"form2_history_egais\".\"doc_id\""},
	ProductQuantity:  whereHelpernull_String{field: "\"form2_history_egais\".\"product_quantity\""},
	OperationName:    whereHelpernull_String{field: "\"form2_history_egais\".\"operation_name\""},
	OperationDate:    whereHelpernull_String{field: "\"form2_history_egais\".\"operation_date\""},
	Form2HistoryDate: whereHelpernull_String{field: "\"form2_history_egais\".\"form2_history_date\""},
}

// Form2HistoryEgaisRels is where relationship names are stored.
var Form2HistoryEgaisRels = struct {
}{}

// form2HistoryEgaisR is where relationships are stored.
type form2HistoryEgaisR struct {
}

// NewStruct creates a new relationship struct
func (*form2HistoryEgaisR) NewStruct() *form2HistoryEgaisR {
	return &form2HistoryEgaisR{}
}

// form2HistoryEgaisL is where Load methods for each relationship are stored.
type form2HistoryEgaisL struct{}

var (
	form2HistoryEgaisAllColumns            = []string{"id", "request_id", "inform_f2_reg_id", "doc_type", "doc_id", "product_quantity", "operation_name", "operation_date", "form2_history_date"}
	form2HistoryEgaisColumnsWithoutDefault = []string{"inform_f2_reg_id"}
	form2HistoryEgaisColumnsWithDefault    = []string{"id", "request_id", "doc_type", "doc_id", "product_quantity", "operation_name", "operation_date", "form2_history_date"}
	form2HistoryEgaisPrimaryKeyColumns     = []string{"id"}
	form2HistoryEgaisGeneratedColumns      = []string{"id"}
)

type (
	// Form2HistoryEgaisSlice is an alias for a slice of pointers to Form2HistoryEgais.
	// This should almost always be used instead of []Form2HistoryEgais.
	Form2HistoryEgaisSlice []*Form2HistoryEgais
	// Form2HistoryEgaisHook is the signature for custom Form2HistoryEgais hook methods
	Form2HistoryEgaisHook func(context.Context, boil.ContextExecutor, *Form2HistoryEgais) error

	form2HistoryEgaisQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	form2HistoryEgaisType                 = reflect.TypeOf(&Form2HistoryEgais{})
	form2HistoryEgaisMapping              = queries.MakeStructMapping(form2HistoryEgaisType)
	form2HistoryEgaisPrimaryKeyMapping, _ = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, form2HistoryEgaisPrimaryKeyColumns)
	form2HistoryEgaisInsertCacheMut       sync.RWMutex
	form2HistoryEgaisInsertCache          = make(map[string]insertCache)
	form2HistoryEgaisUpdateCacheMut       sync.RWMutex
	form2HistoryEgaisUpdateCache          = make(map[string]updateCache)
	form2HistoryEgaisUpsertCacheMut       sync.RWMutex
	form2HistoryEgaisUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var form2HistoryEgaisAfterSelectMu sync.Mutex
var form2HistoryEgaisAfterSelectHooks []Form2HistoryEgaisHook

var form2HistoryEgaisBeforeInsertMu sync.Mutex
var form2HistoryEgaisBeforeInsertHooks []Form2HistoryEgaisHook
var form2HistoryEgaisAfterInsertMu sync.Mutex
var form2HistoryEgaisAfterInsertHooks []Form2HistoryEgaisHook

var form2HistoryEgaisBeforeUpdateMu sync.Mutex
var form2HistoryEgaisBeforeUpdateHooks []Form2HistoryEgaisHook
var form2HistoryEgaisAfterUpdateMu sync.Mutex
var form2HistoryEgaisAfterUpdateHooks []Form2HistoryEgaisHook

var form2HistoryEgaisBeforeDeleteMu sync.Mutex
var form2HistoryEgaisBeforeDeleteHooks []Form2HistoryEgaisHook
var form2HistoryEgaisAfterDeleteMu sync.Mutex
var form2HistoryEgaisAfterDeleteHooks []Form2HistoryEgaisHook

var form2HistoryEgaisBeforeUpsertMu sync.Mutex
var form2HistoryEgaisBeforeUpsertHooks []Form2HistoryEgaisHook
var form2HistoryEgaisAfterUpsertMu sync.Mutex
var form2HistoryEgaisAfterUpsertHooks []Form2HistoryEgaisHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Form2HistoryEgais) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Form2HistoryEgais) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Form2HistoryEgais) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Form2HistoryEgais) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Form2HistoryEgais) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Form2HistoryEgais) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Form2HistoryEgais) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Form2HistoryEgais) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Form2HistoryEgais) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range form2HistoryEgaisAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddForm2HistoryEgaisHook registers your hook function for all future operations.
func AddForm2HistoryEgaisHook(hookPoint boil.HookPoint, form2HistoryEgaisHook Form2HistoryEgaisHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		form2HistoryEgaisAfterSelectMu.Lock()
		form2HistoryEgaisAfterSelectHooks = append(form2HistoryEgaisAfterSelectHooks, form2HistoryEgaisHook)
		form2HistoryEgaisAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		form2HistoryEgaisBeforeInsertMu.Lock()
		form2HistoryEgaisBeforeInsertHooks = append(form2HistoryEgaisBeforeInsertHooks, form2HistoryEgaisHook)
		form2HistoryEgaisBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		form2HistoryEgaisAfterInsertMu.Lock()
		form2HistoryEgaisAfterInsertHooks = append(form2HistoryEgaisAfterInsertHooks, form2HistoryEgaisHook)
		form2HistoryEgaisAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		form2HistoryEgaisBeforeUpdateMu.Lock()
		form2HistoryEgaisBeforeUpdateHooks = append(form2HistoryEgaisBeforeUpdateHooks, form2HistoryEgaisHook)
		form2HistoryEgaisBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		form2HistoryEgaisAfterUpdateMu.Lock()
		form2HistoryEgaisAfterUpdateHooks = append(form2HistoryEgaisAfterUpdateHooks, form2HistoryEgaisHook)
		form2HistoryEgaisAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		form2HistoryEgaisBeforeDeleteMu.Lock()
		form2HistoryEgaisBeforeDeleteHooks = append(form2HistoryEgaisBeforeDeleteHooks, form2HistoryEgaisHook)
		form2HistoryEgaisBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		form2HistoryEgaisAfterDeleteMu.Lock()
		form2HistoryEgaisAfterDeleteHooks = append(form2HistoryEgaisAfterDeleteHooks, form2HistoryEgaisHook)
		form2HistoryEgaisAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		form2HistoryEgaisBeforeUpsertMu.Lock()
		form2HistoryEgaisBeforeUpsertHooks = append(form2HistoryEgaisBeforeUpsertHooks, form2HistoryEgaisHook)
		form2HistoryEgaisBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		form2HistoryEgaisAfterUpsertMu.Lock()
		form2HistoryEgaisAfterUpsertHooks = append(form2HistoryEgaisAfterUpsertHooks, form2HistoryEgaisHook)
		form2HistoryEgaisAfterUpsertMu.Unlock()
	}
}

// OneG returns a single form2HistoryEgais record from the query using the global executor.
func (q form2HistoryEgaisQuery) OneG(ctx context.Context) (*Form2HistoryEgais, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single form2HistoryEgais record from the query.
func (q form2HistoryEgaisQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Form2HistoryEgais, error) {
	o := &Form2HistoryEgais{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: failed to execute a one query for form2_history_egais")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Form2HistoryEgais records from the query using the global executor.
func (q form2HistoryEgaisQuery) AllG(ctx context.Context) (Form2HistoryEgaisSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Form2HistoryEgais records from the query.
func (q form2HistoryEgaisQuery) All(ctx context.Context, exec boil.ContextExecutor) (Form2HistoryEgaisSlice, error) {
	var o []*Form2HistoryEgais

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rq: failed to assign all query results to Form2HistoryEgais slice")
	}

	if len(form2HistoryEgaisAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Form2HistoryEgais records in the query using the global executor
func (q form2HistoryEgaisQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Form2HistoryEgais records in the query.
func (q form2HistoryEgaisQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to count form2_history_egais rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q form2HistoryEgaisQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q form2HistoryEgaisQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rq: failed to check if form2_history_egais exists")
	}

	return count > 0, nil
}

// Form2HistoryEgaiss retrieves all the records using an executor.
func Form2HistoryEgaiss(mods ...qm.QueryMod) form2HistoryEgaisQuery {
	mods = append(mods, qm.From("\"form2_history_egais\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"form2_history_egais\".*"})
	}

	return form2HistoryEgaisQuery{q}
}

// FindForm2HistoryEgaisG retrieves a single record by ID.
func FindForm2HistoryEgaisG(ctx context.Context, iD int64, selectCols ...string) (*Form2HistoryEgais, error) {
	return FindForm2HistoryEgais(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindForm2HistoryEgais retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindForm2HistoryEgais(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Form2HistoryEgais, error) {
	form2HistoryEgaisObj := &Form2HistoryEgais{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"form2_history_egais\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, form2HistoryEgaisObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: unable to select from form2_history_egais")
	}

	if err = form2HistoryEgaisObj.doAfterSelectHooks(ctx, exec); err != nil {
		return form2HistoryEgaisObj, err
	}

	return form2HistoryEgaisObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Form2HistoryEgais) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Form2HistoryEgais) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no form2_history_egais provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(form2HistoryEgaisColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	form2HistoryEgaisInsertCacheMut.RLock()
	cache, cached := form2HistoryEgaisInsertCache[key]
	form2HistoryEgaisInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			form2HistoryEgaisAllColumns,
			form2HistoryEgaisColumnsWithDefault,
			form2HistoryEgaisColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, form2HistoryEgaisGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"form2_history_egais\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"form2_history_egais\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "rq: unable to insert into form2_history_egais")
	}

	if !cached {
		form2HistoryEgaisInsertCacheMut.Lock()
		form2HistoryEgaisInsertCache[key] = cache
		form2HistoryEgaisInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Form2HistoryEgais record using the global executor.
// See Update for more documentation.
func (o *Form2HistoryEgais) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Form2HistoryEgais.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Form2HistoryEgais) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	form2HistoryEgaisUpdateCacheMut.RLock()
	cache, cached := form2HistoryEgaisUpdateCache[key]
	form2HistoryEgaisUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			form2HistoryEgaisAllColumns,
			form2HistoryEgaisPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, form2HistoryEgaisGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("rq: unable to update form2_history_egais, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"form2_history_egais\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, form2HistoryEgaisPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, append(wl, form2HistoryEgaisPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rq: unable to update form2_history_egais row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by update for form2_history_egais")
	}

	if !cached {
		form2HistoryEgaisUpdateCacheMut.Lock()
		form2HistoryEgaisUpdateCache[key] = cache
		form2HistoryEgaisUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q form2HistoryEgaisQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q form2HistoryEgaisQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all for form2_history_egais")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected for form2_history_egais")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o Form2HistoryEgaisSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o Form2HistoryEgaisSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("rq: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), form2HistoryEgaisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"form2_history_egais\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, form2HistoryEgaisPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all in form2HistoryEgais slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected all in update all form2HistoryEgais")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Form2HistoryEgais) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Form2HistoryEgais) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no form2_history_egais provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(form2HistoryEgaisColumnsWithDefault, o)

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

	form2HistoryEgaisUpsertCacheMut.RLock()
	cache, cached := form2HistoryEgaisUpsertCache[key]
	form2HistoryEgaisUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			form2HistoryEgaisAllColumns,
			form2HistoryEgaisColumnsWithDefault,
			form2HistoryEgaisColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			form2HistoryEgaisAllColumns,
			form2HistoryEgaisPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("rq: unable to upsert form2_history_egais, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(form2HistoryEgaisPrimaryKeyColumns))
			copy(conflict, form2HistoryEgaisPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"form2_history_egais\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(form2HistoryEgaisType, form2HistoryEgaisMapping, ret)
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
		return errors.Wrap(err, "rq: unable to upsert form2_history_egais")
	}

	if !cached {
		form2HistoryEgaisUpsertCacheMut.Lock()
		form2HistoryEgaisUpsertCache[key] = cache
		form2HistoryEgaisUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Form2HistoryEgais record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Form2HistoryEgais) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Form2HistoryEgais record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Form2HistoryEgais) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rq: no Form2HistoryEgais provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), form2HistoryEgaisPrimaryKeyMapping)
	sql := "DELETE FROM \"form2_history_egais\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete from form2_history_egais")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by delete for form2_history_egais")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q form2HistoryEgaisQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q form2HistoryEgaisQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rq: no form2HistoryEgaisQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from form2_history_egais")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for form2_history_egais")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o Form2HistoryEgaisSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o Form2HistoryEgaisSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(form2HistoryEgaisBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), form2HistoryEgaisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"form2_history_egais\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, form2HistoryEgaisPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from form2HistoryEgais slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for form2_history_egais")
	}

	if len(form2HistoryEgaisAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Form2HistoryEgais) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: no Form2HistoryEgais provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Form2HistoryEgais) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindForm2HistoryEgais(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *Form2HistoryEgaisSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: empty Form2HistoryEgaisSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *Form2HistoryEgaisSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := Form2HistoryEgaisSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), form2HistoryEgaisPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"form2_history_egais\".* FROM \"form2_history_egais\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, form2HistoryEgaisPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rq: unable to reload all in Form2HistoryEgaisSlice")
	}

	*o = slice

	return nil
}

// Form2HistoryEgaisExistsG checks if the Form2HistoryEgais row exists.
func Form2HistoryEgaisExistsG(ctx context.Context, iD int64) (bool, error) {
	return Form2HistoryEgaisExists(ctx, boil.GetContextDB(), iD)
}

// Form2HistoryEgaisExists checks if the Form2HistoryEgais row exists.
func Form2HistoryEgaisExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"form2_history_egais\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rq: unable to check if form2_history_egais exists")
	}

	return exists, nil
}

// Exists checks if the Form2HistoryEgais row exists.
func (o *Form2HistoryEgais) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return Form2HistoryEgaisExists(ctx, exec, o.ID)
}