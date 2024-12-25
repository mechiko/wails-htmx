// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package liteboil

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// VidAp is an object representing the database table.
type VidAp struct {
	Vcode   string `boil:"vcode" json:"vcode" toml:"vcode" yaml:"vcode"`
	Deleted int64  `boil:"deleted" json:"deleted" toml:"deleted" yaml:"deleted"`
	Name    string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Typ     string `boil:"typ" json:"typ" toml:"typ" yaml:"typ"`
	Okved2  string `boil:"okved2" json:"okved2" toml:"okved2" yaml:"okved2"`
	Other   string `boil:"other" json:"other" toml:"other" yaml:"other"`

	R *vidApR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L vidApL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VidApColumns = struct {
	Vcode   string
	Deleted string
	Name    string
	Typ     string
	Okved2  string
	Other   string
}{
	Vcode:   "vcode",
	Deleted: "deleted",
	Name:    "name",
	Typ:     "typ",
	Okved2:  "okved2",
	Other:   "other",
}

var VidApTableColumns = struct {
	Vcode   string
	Deleted string
	Name    string
	Typ     string
	Okved2  string
	Other   string
}{
	Vcode:   "vid_ap.vcode",
	Deleted: "vid_ap.deleted",
	Name:    "vid_ap.name",
	Typ:     "vid_ap.typ",
	Okved2:  "vid_ap.okved2",
	Other:   "vid_ap.other",
}

// Generated where

var VidApWhere = struct {
	Vcode   whereHelperstring
	Deleted whereHelperint64
	Name    whereHelperstring
	Typ     whereHelperstring
	Okved2  whereHelperstring
	Other   whereHelperstring
}{
	Vcode:   whereHelperstring{field: "\"vid_ap\".\"vcode\""},
	Deleted: whereHelperint64{field: "\"vid_ap\".\"deleted\""},
	Name:    whereHelperstring{field: "\"vid_ap\".\"name\""},
	Typ:     whereHelperstring{field: "\"vid_ap\".\"typ\""},
	Okved2:  whereHelperstring{field: "\"vid_ap\".\"okved2\""},
	Other:   whereHelperstring{field: "\"vid_ap\".\"other\""},
}

// VidApRels is where relationship names are stored.
var VidApRels = struct {
}{}

// vidApR is where relationships are stored.
type vidApR struct {
}

// NewStruct creates a new relationship struct
func (*vidApR) NewStruct() *vidApR {
	return &vidApR{}
}

// vidApL is where Load methods for each relationship are stored.
type vidApL struct{}

var (
	vidApAllColumns            = []string{"vcode", "deleted", "name", "typ", "okved2", "other"}
	vidApColumnsWithoutDefault = []string{"vcode"}
	vidApColumnsWithDefault    = []string{"deleted", "name", "typ", "okved2", "other"}
	vidApPrimaryKeyColumns     = []string{"vcode"}
	vidApGeneratedColumns      = []string{}
)

type (
	// VidApSlice is an alias for a slice of pointers to VidAp.
	// This should almost always be used instead of []VidAp.
	VidApSlice []*VidAp
	// VidApHook is the signature for custom VidAp hook methods
	VidApHook func(context.Context, boil.ContextExecutor, *VidAp) error

	vidApQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	vidApType                 = reflect.TypeOf(&VidAp{})
	vidApMapping              = queries.MakeStructMapping(vidApType)
	vidApPrimaryKeyMapping, _ = queries.BindMapping(vidApType, vidApMapping, vidApPrimaryKeyColumns)
	vidApInsertCacheMut       sync.RWMutex
	vidApInsertCache          = make(map[string]insertCache)
	vidApUpdateCacheMut       sync.RWMutex
	vidApUpdateCache          = make(map[string]updateCache)
	vidApUpsertCacheMut       sync.RWMutex
	vidApUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var vidApAfterSelectMu sync.Mutex
var vidApAfterSelectHooks []VidApHook

var vidApBeforeInsertMu sync.Mutex
var vidApBeforeInsertHooks []VidApHook
var vidApAfterInsertMu sync.Mutex
var vidApAfterInsertHooks []VidApHook

var vidApBeforeUpdateMu sync.Mutex
var vidApBeforeUpdateHooks []VidApHook
var vidApAfterUpdateMu sync.Mutex
var vidApAfterUpdateHooks []VidApHook

var vidApBeforeDeleteMu sync.Mutex
var vidApBeforeDeleteHooks []VidApHook
var vidApAfterDeleteMu sync.Mutex
var vidApAfterDeleteHooks []VidApHook

var vidApBeforeUpsertMu sync.Mutex
var vidApBeforeUpsertHooks []VidApHook
var vidApAfterUpsertMu sync.Mutex
var vidApAfterUpsertHooks []VidApHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VidAp) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VidAp) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VidAp) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VidAp) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VidAp) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VidAp) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VidAp) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VidAp) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VidAp) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vidApAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVidApHook registers your hook function for all future operations.
func AddVidApHook(hookPoint boil.HookPoint, vidApHook VidApHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		vidApAfterSelectMu.Lock()
		vidApAfterSelectHooks = append(vidApAfterSelectHooks, vidApHook)
		vidApAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		vidApBeforeInsertMu.Lock()
		vidApBeforeInsertHooks = append(vidApBeforeInsertHooks, vidApHook)
		vidApBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		vidApAfterInsertMu.Lock()
		vidApAfterInsertHooks = append(vidApAfterInsertHooks, vidApHook)
		vidApAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		vidApBeforeUpdateMu.Lock()
		vidApBeforeUpdateHooks = append(vidApBeforeUpdateHooks, vidApHook)
		vidApBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		vidApAfterUpdateMu.Lock()
		vidApAfterUpdateHooks = append(vidApAfterUpdateHooks, vidApHook)
		vidApAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		vidApBeforeDeleteMu.Lock()
		vidApBeforeDeleteHooks = append(vidApBeforeDeleteHooks, vidApHook)
		vidApBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		vidApAfterDeleteMu.Lock()
		vidApAfterDeleteHooks = append(vidApAfterDeleteHooks, vidApHook)
		vidApAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		vidApBeforeUpsertMu.Lock()
		vidApBeforeUpsertHooks = append(vidApBeforeUpsertHooks, vidApHook)
		vidApBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		vidApAfterUpsertMu.Lock()
		vidApAfterUpsertHooks = append(vidApAfterUpsertHooks, vidApHook)
		vidApAfterUpsertMu.Unlock()
	}
}

// OneG returns a single vidAp record from the query using the global executor.
func (q vidApQuery) OneG(ctx context.Context) (*VidAp, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single vidAp record from the query.
func (q vidApQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VidAp, error) {
	o := &VidAp{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "liteboil: failed to execute a one query for vid_ap")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all VidAp records from the query using the global executor.
func (q vidApQuery) AllG(ctx context.Context) (VidApSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all VidAp records from the query.
func (q vidApQuery) All(ctx context.Context, exec boil.ContextExecutor) (VidApSlice, error) {
	var o []*VidAp

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "liteboil: failed to assign all query results to VidAp slice")
	}

	if len(vidApAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all VidAp records in the query using the global executor
func (q vidApQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all VidAp records in the query.
func (q vidApQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: failed to count vid_ap rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q vidApQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q vidApQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "liteboil: failed to check if vid_ap exists")
	}

	return count > 0, nil
}

// VidAps retrieves all the records using an executor.
func VidAps(mods ...qm.QueryMod) vidApQuery {
	mods = append(mods, qm.From("\"vid_ap\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"vid_ap\".*"})
	}

	return vidApQuery{q}
}

// FindVidApG retrieves a single record by ID.
func FindVidApG(ctx context.Context, vcode string, selectCols ...string) (*VidAp, error) {
	return FindVidAp(ctx, boil.GetContextDB(), vcode, selectCols...)
}

// FindVidAp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVidAp(ctx context.Context, exec boil.ContextExecutor, vcode string, selectCols ...string) (*VidAp, error) {
	vidApObj := &VidAp{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"vid_ap\" where \"vcode\"=?", sel,
	)

	q := queries.Raw(query, vcode)

	err := q.Bind(ctx, exec, vidApObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "liteboil: unable to select from vid_ap")
	}

	if err = vidApObj.doAfterSelectHooks(ctx, exec); err != nil {
		return vidApObj, err
	}

	return vidApObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VidAp) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VidAp) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("liteboil: no vid_ap provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(vidApColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	vidApInsertCacheMut.RLock()
	cache, cached := vidApInsertCache[key]
	vidApInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			vidApAllColumns,
			vidApColumnsWithDefault,
			vidApColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(vidApType, vidApMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(vidApType, vidApMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"vid_ap\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"vid_ap\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "liteboil: unable to insert into vid_ap")
	}

	if !cached {
		vidApInsertCacheMut.Lock()
		vidApInsertCache[key] = cache
		vidApInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single VidAp record using the global executor.
// See Update for more documentation.
func (o *VidAp) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the VidAp.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VidAp) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	vidApUpdateCacheMut.RLock()
	cache, cached := vidApUpdateCache[key]
	vidApUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			vidApAllColumns,
			vidApPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("liteboil: unable to update vid_ap, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"vid_ap\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, vidApPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(vidApType, vidApMapping, append(wl, vidApPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "liteboil: unable to update vid_ap row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: failed to get rows affected by update for vid_ap")
	}

	if !cached {
		vidApUpdateCacheMut.Lock()
		vidApUpdateCache[key] = cache
		vidApUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q vidApQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q vidApQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to update all for vid_ap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to retrieve rows affected for vid_ap")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VidApSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VidApSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("liteboil: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vidApPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"vid_ap\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vidApPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to update all in vidAp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to retrieve rows affected all in update all vidAp")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VidAp) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VidAp) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("liteboil: no vid_ap provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(vidApColumnsWithDefault, o)

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

	vidApUpsertCacheMut.RLock()
	cache, cached := vidApUpsertCache[key]
	vidApUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			vidApAllColumns,
			vidApColumnsWithDefault,
			vidApColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			vidApAllColumns,
			vidApPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("liteboil: unable to upsert vid_ap, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(vidApPrimaryKeyColumns))
			copy(conflict, vidApPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"vid_ap\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(vidApType, vidApMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(vidApType, vidApMapping, ret)
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
		return errors.Wrap(err, "liteboil: unable to upsert vid_ap")
	}

	if !cached {
		vidApUpsertCacheMut.Lock()
		vidApUpsertCache[key] = cache
		vidApUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single VidAp record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VidAp) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single VidAp record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VidAp) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("liteboil: no VidAp provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), vidApPrimaryKeyMapping)
	sql := "DELETE FROM \"vid_ap\" WHERE \"vcode\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to delete from vid_ap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: failed to get rows affected by delete for vid_ap")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q vidApQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q vidApQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("liteboil: no vidApQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to delete all from vid_ap")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: failed to get rows affected by deleteall for vid_ap")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o VidApSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VidApSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(vidApBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vidApPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"vid_ap\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vidApPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: unable to delete all from vidAp slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "liteboil: failed to get rows affected by deleteall for vid_ap")
	}

	if len(vidApAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VidAp) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("liteboil: no VidAp provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VidAp) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVidAp(ctx, exec, o.Vcode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VidApSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("liteboil: empty VidApSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VidApSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VidApSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vidApPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"vid_ap\".* FROM \"vid_ap\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vidApPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "liteboil: unable to reload all in VidApSlice")
	}

	*o = slice

	return nil
}

// VidApExistsG checks if the VidAp row exists.
func VidApExistsG(ctx context.Context, vcode string) (bool, error) {
	return VidApExists(ctx, boil.GetContextDB(), vcode)
}

// VidApExists checks if the VidAp row exists.
func VidApExists(ctx context.Context, exec boil.ContextExecutor, vcode string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"vid_ap\" where \"vcode\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, vcode)
	}
	row := exec.QueryRowContext(ctx, sql, vcode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "liteboil: unable to check if vid_ap exists")
	}

	return exists, nil
}

// Exists checks if the VidAp row exists.
func (o *VidAp) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VidApExists(ctx, exec, o.Vcode)
}