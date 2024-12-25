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

// SPLocal is an object representing the database table.
type SPLocal struct {
	ID               int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductFullName  null.String `boil:"product_full_name" json:"product_full_name,omitempty" toml:"product_full_name" yaml:"product_full_name,omitempty"`
	ProductCapacity  null.String `boil:"product_capacity" json:"product_capacity,omitempty" toml:"product_capacity" yaml:"product_capacity,omitempty"`
	ProductAlcVolume null.String `boil:"product_alc_volume" json:"product_alc_volume,omitempty" toml:"product_alc_volume" yaml:"product_alc_volume,omitempty"`
	ProductAlcCode   null.String `boil:"product_alc_code" json:"product_alc_code,omitempty" toml:"product_alc_code" yaml:"product_alc_code,omitempty"`
	ProductCode      null.String `boil:"product_code" json:"product_code,omitempty" toml:"product_code" yaml:"product_code,omitempty"`
	ProductUnitType  null.String `boil:"product_unit_type" json:"product_unit_type,omitempty" toml:"product_unit_type" yaml:"product_unit_type,omitempty"`

	R *spLocalR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spLocalL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SPLocalColumns = struct {
	ID               string
	ProductFullName  string
	ProductCapacity  string
	ProductAlcVolume string
	ProductAlcCode   string
	ProductCode      string
	ProductUnitType  string
}{
	ID:               "id",
	ProductFullName:  "product_full_name",
	ProductCapacity:  "product_capacity",
	ProductAlcVolume: "product_alc_volume",
	ProductAlcCode:   "product_alc_code",
	ProductCode:      "product_code",
	ProductUnitType:  "product_unit_type",
}

var SPLocalTableColumns = struct {
	ID               string
	ProductFullName  string
	ProductCapacity  string
	ProductAlcVolume string
	ProductAlcCode   string
	ProductCode      string
	ProductUnitType  string
}{
	ID:               "sp_local.id",
	ProductFullName:  "sp_local.product_full_name",
	ProductCapacity:  "sp_local.product_capacity",
	ProductAlcVolume: "sp_local.product_alc_volume",
	ProductAlcCode:   "sp_local.product_alc_code",
	ProductCode:      "sp_local.product_code",
	ProductUnitType:  "sp_local.product_unit_type",
}

// Generated where

var SPLocalWhere = struct {
	ID               whereHelperint64
	ProductFullName  whereHelpernull_String
	ProductCapacity  whereHelpernull_String
	ProductAlcVolume whereHelpernull_String
	ProductAlcCode   whereHelpernull_String
	ProductCode      whereHelpernull_String
	ProductUnitType  whereHelpernull_String
}{
	ID:               whereHelperint64{field: "\"sp_local\".\"id\""},
	ProductFullName:  whereHelpernull_String{field: "\"sp_local\".\"product_full_name\""},
	ProductCapacity:  whereHelpernull_String{field: "\"sp_local\".\"product_capacity\""},
	ProductAlcVolume: whereHelpernull_String{field: "\"sp_local\".\"product_alc_volume\""},
	ProductAlcCode:   whereHelpernull_String{field: "\"sp_local\".\"product_alc_code\""},
	ProductCode:      whereHelpernull_String{field: "\"sp_local\".\"product_code\""},
	ProductUnitType:  whereHelpernull_String{field: "\"sp_local\".\"product_unit_type\""},
}

// SPLocalRels is where relationship names are stored.
var SPLocalRels = struct {
}{}

// spLocalR is where relationships are stored.
type spLocalR struct {
}

// NewStruct creates a new relationship struct
func (*spLocalR) NewStruct() *spLocalR {
	return &spLocalR{}
}

// spLocalL is where Load methods for each relationship are stored.
type spLocalL struct{}

var (
	spLocalAllColumns            = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type"}
	spLocalColumnsWithoutDefault = []string{}
	spLocalColumnsWithDefault    = []string{"id", "product_full_name", "product_capacity", "product_alc_volume", "product_alc_code", "product_code", "product_unit_type"}
	spLocalPrimaryKeyColumns     = []string{"id"}
	spLocalGeneratedColumns      = []string{"id"}
)

type (
	// SPLocalSlice is an alias for a slice of pointers to SPLocal.
	// This should almost always be used instead of []SPLocal.
	SPLocalSlice []*SPLocal
	// SPLocalHook is the signature for custom SPLocal hook methods
	SPLocalHook func(context.Context, boil.ContextExecutor, *SPLocal) error

	spLocalQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spLocalType                 = reflect.TypeOf(&SPLocal{})
	spLocalMapping              = queries.MakeStructMapping(spLocalType)
	spLocalPrimaryKeyMapping, _ = queries.BindMapping(spLocalType, spLocalMapping, spLocalPrimaryKeyColumns)
	spLocalInsertCacheMut       sync.RWMutex
	spLocalInsertCache          = make(map[string]insertCache)
	spLocalUpdateCacheMut       sync.RWMutex
	spLocalUpdateCache          = make(map[string]updateCache)
	spLocalUpsertCacheMut       sync.RWMutex
	spLocalUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var spLocalAfterSelectMu sync.Mutex
var spLocalAfterSelectHooks []SPLocalHook

var spLocalBeforeInsertMu sync.Mutex
var spLocalBeforeInsertHooks []SPLocalHook
var spLocalAfterInsertMu sync.Mutex
var spLocalAfterInsertHooks []SPLocalHook

var spLocalBeforeUpdateMu sync.Mutex
var spLocalBeforeUpdateHooks []SPLocalHook
var spLocalAfterUpdateMu sync.Mutex
var spLocalAfterUpdateHooks []SPLocalHook

var spLocalBeforeDeleteMu sync.Mutex
var spLocalBeforeDeleteHooks []SPLocalHook
var spLocalAfterDeleteMu sync.Mutex
var spLocalAfterDeleteHooks []SPLocalHook

var spLocalBeforeUpsertMu sync.Mutex
var spLocalBeforeUpsertHooks []SPLocalHook
var spLocalAfterUpsertMu sync.Mutex
var spLocalAfterUpsertHooks []SPLocalHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SPLocal) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SPLocal) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SPLocal) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SPLocal) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SPLocal) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SPLocal) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SPLocal) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SPLocal) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SPLocal) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spLocalAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSPLocalHook registers your hook function for all future operations.
func AddSPLocalHook(hookPoint boil.HookPoint, spLocalHook SPLocalHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		spLocalAfterSelectMu.Lock()
		spLocalAfterSelectHooks = append(spLocalAfterSelectHooks, spLocalHook)
		spLocalAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		spLocalBeforeInsertMu.Lock()
		spLocalBeforeInsertHooks = append(spLocalBeforeInsertHooks, spLocalHook)
		spLocalBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		spLocalAfterInsertMu.Lock()
		spLocalAfterInsertHooks = append(spLocalAfterInsertHooks, spLocalHook)
		spLocalAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		spLocalBeforeUpdateMu.Lock()
		spLocalBeforeUpdateHooks = append(spLocalBeforeUpdateHooks, spLocalHook)
		spLocalBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		spLocalAfterUpdateMu.Lock()
		spLocalAfterUpdateHooks = append(spLocalAfterUpdateHooks, spLocalHook)
		spLocalAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		spLocalBeforeDeleteMu.Lock()
		spLocalBeforeDeleteHooks = append(spLocalBeforeDeleteHooks, spLocalHook)
		spLocalBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		spLocalAfterDeleteMu.Lock()
		spLocalAfterDeleteHooks = append(spLocalAfterDeleteHooks, spLocalHook)
		spLocalAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		spLocalBeforeUpsertMu.Lock()
		spLocalBeforeUpsertHooks = append(spLocalBeforeUpsertHooks, spLocalHook)
		spLocalBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		spLocalAfterUpsertMu.Lock()
		spLocalAfterUpsertHooks = append(spLocalAfterUpsertHooks, spLocalHook)
		spLocalAfterUpsertMu.Unlock()
	}
}

// OneG returns a single spLocal record from the query using the global executor.
func (q spLocalQuery) OneG(ctx context.Context) (*SPLocal, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single spLocal record from the query.
func (q spLocalQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SPLocal, error) {
	o := &SPLocal{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for sp_local")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all SPLocal records from the query using the global executor.
func (q spLocalQuery) AllG(ctx context.Context) (SPLocalSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all SPLocal records from the query.
func (q spLocalQuery) All(ctx context.Context, exec boil.ContextExecutor) (SPLocalSlice, error) {
	var o []*SPLocal

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to SPLocal slice")
	}

	if len(spLocalAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all SPLocal records in the query using the global executor
func (q spLocalQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all SPLocal records in the query.
func (q spLocalQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count sp_local rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q spLocalQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q spLocalQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if sp_local exists")
	}

	return count > 0, nil
}

// SPLocals retrieves all the records using an executor.
func SPLocals(mods ...qm.QueryMod) spLocalQuery {
	mods = append(mods, qm.From("\"sp_local\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sp_local\".*"})
	}

	return spLocalQuery{q}
}

// FindSPLocalG retrieves a single record by ID.
func FindSPLocalG(ctx context.Context, iD int64, selectCols ...string) (*SPLocal, error) {
	return FindSPLocal(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindSPLocal retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSPLocal(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*SPLocal, error) {
	spLocalObj := &SPLocal{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"sp_local\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, spLocalObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from sp_local")
	}

	if err = spLocalObj.doAfterSelectHooks(ctx, exec); err != nil {
		return spLocalObj, err
	}

	return spLocalObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *SPLocal) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SPLocal) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no sp_local provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spLocalColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spLocalInsertCacheMut.RLock()
	cache, cached := spLocalInsertCache[key]
	spLocalInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spLocalAllColumns,
			spLocalColumnsWithDefault,
			spLocalColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, spLocalGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(spLocalType, spLocalMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spLocalType, spLocalMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sp_local\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sp_local\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into sp_local")
	}

	if !cached {
		spLocalInsertCacheMut.Lock()
		spLocalInsertCache[key] = cache
		spLocalInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single SPLocal record using the global executor.
// See Update for more documentation.
func (o *SPLocal) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the SPLocal.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SPLocal) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	spLocalUpdateCacheMut.RLock()
	cache, cached := spLocalUpdateCache[key]
	spLocalUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spLocalAllColumns,
			spLocalPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, spLocalGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update sp_local, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"sp_local\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, spLocalPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spLocalType, spLocalMapping, append(wl, spLocalPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update sp_local row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for sp_local")
	}

	if !cached {
		spLocalUpdateCacheMut.Lock()
		spLocalUpdateCache[key] = cache
		spLocalUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q spLocalQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q spLocalQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for sp_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for sp_local")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SPLocalSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SPLocalSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"sp_local\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spLocalPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in spLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all spLocal")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *SPLocal) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SPLocal) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no sp_local provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spLocalColumnsWithDefault, o)

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

	spLocalUpsertCacheMut.RLock()
	cache, cached := spLocalUpsertCache[key]
	spLocalUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spLocalAllColumns,
			spLocalColumnsWithDefault,
			spLocalColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			spLocalAllColumns,
			spLocalPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("a3boil: unable to upsert sp_local, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(spLocalPrimaryKeyColumns))
			copy(conflict, spLocalPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"sp_local\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(spLocalType, spLocalMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spLocalType, spLocalMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert sp_local")
	}

	if !cached {
		spLocalUpsertCacheMut.Lock()
		spLocalUpsertCache[key] = cache
		spLocalUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single SPLocal record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *SPLocal) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single SPLocal record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SPLocal) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no SPLocal provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spLocalPrimaryKeyMapping)
	sql := "DELETE FROM \"sp_local\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from sp_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for sp_local")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q spLocalQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q spLocalQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no spLocalQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from sp_local")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for sp_local")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SPLocalSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SPLocalSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(spLocalBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"sp_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spLocalPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from spLocal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for sp_local")
	}

	if len(spLocalAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *SPLocal) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no SPLocal provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SPLocal) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSPLocal(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SPLocalSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty SPLocalSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SPLocalSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SPLocalSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spLocalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"sp_local\".* FROM \"sp_local\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, spLocalPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in SPLocalSlice")
	}

	*o = slice

	return nil
}

// SPLocalExistsG checks if the SPLocal row exists.
func SPLocalExistsG(ctx context.Context, iD int64) (bool, error) {
	return SPLocalExists(ctx, boil.GetContextDB(), iD)
}

// SPLocalExists checks if the SPLocal row exists.
func SPLocalExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"sp_local\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if sp_local exists")
	}

	return exists, nil
}

// Exists checks if the SPLocal row exists.
func (o *SPLocal) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SPLocalExists(ctx, exec, o.ID)
}