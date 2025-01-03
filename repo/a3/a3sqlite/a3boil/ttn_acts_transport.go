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

// TTNActsTransport is an object representing the database table.
type TTNActsTransport struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTTNActs null.Int64  `boil:"id_ttn_acts" json:"id_ttn_acts,omitempty" toml:"id_ttn_acts" yaml:"id_ttn_acts,omitempty"`
	Ownership null.String `boil:"ownership" json:"ownership,omitempty" toml:"ownership" yaml:"ownership,omitempty"`

	R *ttnActsTransportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ttnActsTransportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TTNActsTransportColumns = struct {
	ID        string
	IDTTNActs string
	Ownership string
}{
	ID:        "id",
	IDTTNActs: "id_ttn_acts",
	Ownership: "ownership",
}

var TTNActsTransportTableColumns = struct {
	ID        string
	IDTTNActs string
	Ownership string
}{
	ID:        "ttn_acts_transport.id",
	IDTTNActs: "ttn_acts_transport.id_ttn_acts",
	Ownership: "ttn_acts_transport.ownership",
}

// Generated where

var TTNActsTransportWhere = struct {
	ID        whereHelperint64
	IDTTNActs whereHelpernull_Int64
	Ownership whereHelpernull_String
}{
	ID:        whereHelperint64{field: "\"ttn_acts_transport\".\"id\""},
	IDTTNActs: whereHelpernull_Int64{field: "\"ttn_acts_transport\".\"id_ttn_acts\""},
	Ownership: whereHelpernull_String{field: "\"ttn_acts_transport\".\"ownership\""},
}

// TTNActsTransportRels is where relationship names are stored.
var TTNActsTransportRels = struct {
}{}

// ttnActsTransportR is where relationships are stored.
type ttnActsTransportR struct {
}

// NewStruct creates a new relationship struct
func (*ttnActsTransportR) NewStruct() *ttnActsTransportR {
	return &ttnActsTransportR{}
}

// ttnActsTransportL is where Load methods for each relationship are stored.
type ttnActsTransportL struct{}

var (
	ttnActsTransportAllColumns            = []string{"id", "id_ttn_acts", "ownership"}
	ttnActsTransportColumnsWithoutDefault = []string{}
	ttnActsTransportColumnsWithDefault    = []string{"id", "id_ttn_acts", "ownership"}
	ttnActsTransportPrimaryKeyColumns     = []string{"id"}
	ttnActsTransportGeneratedColumns      = []string{"id"}
)

type (
	// TTNActsTransportSlice is an alias for a slice of pointers to TTNActsTransport.
	// This should almost always be used instead of []TTNActsTransport.
	TTNActsTransportSlice []*TTNActsTransport
	// TTNActsTransportHook is the signature for custom TTNActsTransport hook methods
	TTNActsTransportHook func(context.Context, boil.ContextExecutor, *TTNActsTransport) error

	ttnActsTransportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ttnActsTransportType                 = reflect.TypeOf(&TTNActsTransport{})
	ttnActsTransportMapping              = queries.MakeStructMapping(ttnActsTransportType)
	ttnActsTransportPrimaryKeyMapping, _ = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, ttnActsTransportPrimaryKeyColumns)
	ttnActsTransportInsertCacheMut       sync.RWMutex
	ttnActsTransportInsertCache          = make(map[string]insertCache)
	ttnActsTransportUpdateCacheMut       sync.RWMutex
	ttnActsTransportUpdateCache          = make(map[string]updateCache)
	ttnActsTransportUpsertCacheMut       sync.RWMutex
	ttnActsTransportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ttnActsTransportAfterSelectMu sync.Mutex
var ttnActsTransportAfterSelectHooks []TTNActsTransportHook

var ttnActsTransportBeforeInsertMu sync.Mutex
var ttnActsTransportBeforeInsertHooks []TTNActsTransportHook
var ttnActsTransportAfterInsertMu sync.Mutex
var ttnActsTransportAfterInsertHooks []TTNActsTransportHook

var ttnActsTransportBeforeUpdateMu sync.Mutex
var ttnActsTransportBeforeUpdateHooks []TTNActsTransportHook
var ttnActsTransportAfterUpdateMu sync.Mutex
var ttnActsTransportAfterUpdateHooks []TTNActsTransportHook

var ttnActsTransportBeforeDeleteMu sync.Mutex
var ttnActsTransportBeforeDeleteHooks []TTNActsTransportHook
var ttnActsTransportAfterDeleteMu sync.Mutex
var ttnActsTransportAfterDeleteHooks []TTNActsTransportHook

var ttnActsTransportBeforeUpsertMu sync.Mutex
var ttnActsTransportBeforeUpsertHooks []TTNActsTransportHook
var ttnActsTransportAfterUpsertMu sync.Mutex
var ttnActsTransportAfterUpsertHooks []TTNActsTransportHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TTNActsTransport) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TTNActsTransport) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TTNActsTransport) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TTNActsTransport) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TTNActsTransport) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TTNActsTransport) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TTNActsTransport) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TTNActsTransport) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TTNActsTransport) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnActsTransportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTTNActsTransportHook registers your hook function for all future operations.
func AddTTNActsTransportHook(hookPoint boil.HookPoint, ttnActsTransportHook TTNActsTransportHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ttnActsTransportAfterSelectMu.Lock()
		ttnActsTransportAfterSelectHooks = append(ttnActsTransportAfterSelectHooks, ttnActsTransportHook)
		ttnActsTransportAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		ttnActsTransportBeforeInsertMu.Lock()
		ttnActsTransportBeforeInsertHooks = append(ttnActsTransportBeforeInsertHooks, ttnActsTransportHook)
		ttnActsTransportBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		ttnActsTransportAfterInsertMu.Lock()
		ttnActsTransportAfterInsertHooks = append(ttnActsTransportAfterInsertHooks, ttnActsTransportHook)
		ttnActsTransportAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		ttnActsTransportBeforeUpdateMu.Lock()
		ttnActsTransportBeforeUpdateHooks = append(ttnActsTransportBeforeUpdateHooks, ttnActsTransportHook)
		ttnActsTransportBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		ttnActsTransportAfterUpdateMu.Lock()
		ttnActsTransportAfterUpdateHooks = append(ttnActsTransportAfterUpdateHooks, ttnActsTransportHook)
		ttnActsTransportAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		ttnActsTransportBeforeDeleteMu.Lock()
		ttnActsTransportBeforeDeleteHooks = append(ttnActsTransportBeforeDeleteHooks, ttnActsTransportHook)
		ttnActsTransportBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		ttnActsTransportAfterDeleteMu.Lock()
		ttnActsTransportAfterDeleteHooks = append(ttnActsTransportAfterDeleteHooks, ttnActsTransportHook)
		ttnActsTransportAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		ttnActsTransportBeforeUpsertMu.Lock()
		ttnActsTransportBeforeUpsertHooks = append(ttnActsTransportBeforeUpsertHooks, ttnActsTransportHook)
		ttnActsTransportBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		ttnActsTransportAfterUpsertMu.Lock()
		ttnActsTransportAfterUpsertHooks = append(ttnActsTransportAfterUpsertHooks, ttnActsTransportHook)
		ttnActsTransportAfterUpsertMu.Unlock()
	}
}

// OneG returns a single ttnActsTransport record from the query using the global executor.
func (q ttnActsTransportQuery) OneG(ctx context.Context) (*TTNActsTransport, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ttnActsTransport record from the query.
func (q ttnActsTransportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TTNActsTransport, error) {
	o := &TTNActsTransport{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for ttn_acts_transport")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TTNActsTransport records from the query using the global executor.
func (q ttnActsTransportQuery) AllG(ctx context.Context) (TTNActsTransportSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TTNActsTransport records from the query.
func (q ttnActsTransportQuery) All(ctx context.Context, exec boil.ContextExecutor) (TTNActsTransportSlice, error) {
	var o []*TTNActsTransport

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TTNActsTransport slice")
	}

	if len(ttnActsTransportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TTNActsTransport records in the query using the global executor
func (q ttnActsTransportQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TTNActsTransport records in the query.
func (q ttnActsTransportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count ttn_acts_transport rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q ttnActsTransportQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ttnActsTransportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if ttn_acts_transport exists")
	}

	return count > 0, nil
}

// TTNActsTransports retrieves all the records using an executor.
func TTNActsTransports(mods ...qm.QueryMod) ttnActsTransportQuery {
	mods = append(mods, qm.From("\"ttn_acts_transport\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"ttn_acts_transport\".*"})
	}

	return ttnActsTransportQuery{q}
}

// FindTTNActsTransportG retrieves a single record by ID.
func FindTTNActsTransportG(ctx context.Context, iD int64, selectCols ...string) (*TTNActsTransport, error) {
	return FindTTNActsTransport(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTTNActsTransport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTTNActsTransport(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TTNActsTransport, error) {
	ttnActsTransportObj := &TTNActsTransport{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ttn_acts_transport\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ttnActsTransportObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from ttn_acts_transport")
	}

	if err = ttnActsTransportObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ttnActsTransportObj, err
	}

	return ttnActsTransportObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TTNActsTransport) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TTNActsTransport) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_acts_transport provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnActsTransportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ttnActsTransportInsertCacheMut.RLock()
	cache, cached := ttnActsTransportInsertCache[key]
	ttnActsTransportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ttnActsTransportAllColumns,
			ttnActsTransportColumnsWithDefault,
			ttnActsTransportColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, ttnActsTransportGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ttn_acts_transport\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ttn_acts_transport\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into ttn_acts_transport")
	}

	if !cached {
		ttnActsTransportInsertCacheMut.Lock()
		ttnActsTransportInsertCache[key] = cache
		ttnActsTransportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TTNActsTransport record using the global executor.
// See Update for more documentation.
func (o *TTNActsTransport) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TTNActsTransport.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TTNActsTransport) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ttnActsTransportUpdateCacheMut.RLock()
	cache, cached := ttnActsTransportUpdateCache[key]
	ttnActsTransportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ttnActsTransportAllColumns,
			ttnActsTransportPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, ttnActsTransportGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update ttn_acts_transport, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ttn_acts_transport\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, ttnActsTransportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, append(wl, ttnActsTransportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update ttn_acts_transport row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for ttn_acts_transport")
	}

	if !cached {
		ttnActsTransportUpdateCacheMut.Lock()
		ttnActsTransportUpdateCache[key] = cache
		ttnActsTransportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q ttnActsTransportQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ttnActsTransportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for ttn_acts_transport")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for ttn_acts_transport")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TTNActsTransportSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TTNActsTransportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnActsTransportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ttn_acts_transport\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttnActsTransportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in ttnActsTransport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all ttnActsTransport")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TTNActsTransport) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TTNActsTransport) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_acts_transport provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnActsTransportColumnsWithDefault, o)

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

	ttnActsTransportUpsertCacheMut.RLock()
	cache, cached := ttnActsTransportUpsertCache[key]
	ttnActsTransportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ttnActsTransportAllColumns,
			ttnActsTransportColumnsWithDefault,
			ttnActsTransportColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ttnActsTransportAllColumns,
			ttnActsTransportPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert ttn_acts_transport, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ttnActsTransportPrimaryKeyColumns))
			copy(conflict, ttnActsTransportPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"ttn_acts_transport\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ttnActsTransportType, ttnActsTransportMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert ttn_acts_transport")
	}

	if !cached {
		ttnActsTransportUpsertCacheMut.Lock()
		ttnActsTransportUpsertCache[key] = cache
		ttnActsTransportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TTNActsTransport record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TTNActsTransport) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TTNActsTransport record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TTNActsTransport) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TTNActsTransport provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ttnActsTransportPrimaryKeyMapping)
	sql := "DELETE FROM \"ttn_acts_transport\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from ttn_acts_transport")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for ttn_acts_transport")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q ttnActsTransportQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q ttnActsTransportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no ttnActsTransportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttn_acts_transport")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_acts_transport")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TTNActsTransportSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TTNActsTransportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ttnActsTransportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnActsTransportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ttn_acts_transport\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttnActsTransportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttnActsTransport slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_acts_transport")
	}

	if len(ttnActsTransportAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TTNActsTransport) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TTNActsTransport provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TTNActsTransport) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTTNActsTransport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNActsTransportSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TTNActsTransportSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNActsTransportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TTNActsTransportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnActsTransportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ttn_acts_transport\".* FROM \"ttn_acts_transport\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, ttnActsTransportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TTNActsTransportSlice")
	}

	*o = slice

	return nil
}

// TTNActsTransportExistsG checks if the TTNActsTransport row exists.
func TTNActsTransportExistsG(ctx context.Context, iD int64) (bool, error) {
	return TTNActsTransportExists(ctx, boil.GetContextDB(), iD)
}

// TTNActsTransportExists checks if the TTNActsTransport row exists.
func TTNActsTransportExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ttn_acts_transport\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if ttn_acts_transport exists")
	}

	return exists, nil
}

// Exists checks if the TTNActsTransport row exists.
func (o *TTNActsTransport) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TTNActsTransportExists(ctx, exec, o.ID)
}
