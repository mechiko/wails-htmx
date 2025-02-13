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

// TTNProductsBox is an object representing the database table.
type TTNProductsBox struct {
	ID                   int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTTNProducts        null.Int    `boil:"id_ttn_products" json:"id_ttn_products,omitempty" toml:"id_ttn_products" yaml:"id_ttn_products,omitempty"`
	IDTTNProductsPallets null.Int    `boil:"id_ttn_products_pallets" json:"id_ttn_products_pallets,omitempty" toml:"id_ttn_products_pallets" yaml:"id_ttn_products_pallets,omitempty"`
	BoxNumber            null.String `boil:"box_number" json:"box_number,omitempty" toml:"box_number" yaml:"box_number,omitempty"`

	R *ttnProductsBoxR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ttnProductsBoxL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TTNProductsBoxColumns = struct {
	ID                   string
	IDTTNProducts        string
	IDTTNProductsPallets string
	BoxNumber            string
}{
	ID:                   "id",
	IDTTNProducts:        "id_ttn_products",
	IDTTNProductsPallets: "id_ttn_products_pallets",
	BoxNumber:            "box_number",
}

var TTNProductsBoxTableColumns = struct {
	ID                   string
	IDTTNProducts        string
	IDTTNProductsPallets string
	BoxNumber            string
}{
	ID:                   "ttn_products_boxes.id",
	IDTTNProducts:        "ttn_products_boxes.id_ttn_products",
	IDTTNProductsPallets: "ttn_products_boxes.id_ttn_products_pallets",
	BoxNumber:            "ttn_products_boxes.box_number",
}

// Generated where

var TTNProductsBoxWhere = struct {
	ID                   whereHelperint
	IDTTNProducts        whereHelpernull_Int
	IDTTNProductsPallets whereHelpernull_Int
	BoxNumber            whereHelpernull_String
}{
	ID:                   whereHelperint{field: "[dbo].[ttn_products_boxes].[id]"},
	IDTTNProducts:        whereHelpernull_Int{field: "[dbo].[ttn_products_boxes].[id_ttn_products]"},
	IDTTNProductsPallets: whereHelpernull_Int{field: "[dbo].[ttn_products_boxes].[id_ttn_products_pallets]"},
	BoxNumber:            whereHelpernull_String{field: "[dbo].[ttn_products_boxes].[box_number]"},
}

// TTNProductsBoxRels is where relationship names are stored.
var TTNProductsBoxRels = struct {
}{}

// ttnProductsBoxR is where relationships are stored.
type ttnProductsBoxR struct {
}

// NewStruct creates a new relationship struct
func (*ttnProductsBoxR) NewStruct() *ttnProductsBoxR {
	return &ttnProductsBoxR{}
}

// ttnProductsBoxL is where Load methods for each relationship are stored.
type ttnProductsBoxL struct{}

var (
	ttnProductsBoxAllColumns            = []string{"id", "id_ttn_products", "id_ttn_products_pallets", "box_number"}
	ttnProductsBoxColumnsWithoutDefault = []string{"id_ttn_products", "id_ttn_products_pallets", "box_number"}
	ttnProductsBoxColumnsWithDefault    = []string{"id"}
	ttnProductsBoxPrimaryKeyColumns     = []string{"id"}
	ttnProductsBoxGeneratedColumns      = []string{"id"}
)

type (
	// TTNProductsBoxSlice is an alias for a slice of pointers to TTNProductsBox.
	// This should almost always be used instead of []TTNProductsBox.
	TTNProductsBoxSlice []*TTNProductsBox
	// TTNProductsBoxHook is the signature for custom TTNProductsBox hook methods
	TTNProductsBoxHook func(context.Context, boil.ContextExecutor, *TTNProductsBox) error

	ttnProductsBoxQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ttnProductsBoxType                 = reflect.TypeOf(&TTNProductsBox{})
	ttnProductsBoxMapping              = queries.MakeStructMapping(ttnProductsBoxType)
	ttnProductsBoxPrimaryKeyMapping, _ = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, ttnProductsBoxPrimaryKeyColumns)
	ttnProductsBoxInsertCacheMut       sync.RWMutex
	ttnProductsBoxInsertCache          = make(map[string]insertCache)
	ttnProductsBoxUpdateCacheMut       sync.RWMutex
	ttnProductsBoxUpdateCache          = make(map[string]updateCache)
	ttnProductsBoxUpsertCacheMut       sync.RWMutex
	ttnProductsBoxUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ttnProductsBoxAfterSelectMu sync.Mutex
var ttnProductsBoxAfterSelectHooks []TTNProductsBoxHook

var ttnProductsBoxBeforeInsertMu sync.Mutex
var ttnProductsBoxBeforeInsertHooks []TTNProductsBoxHook
var ttnProductsBoxAfterInsertMu sync.Mutex
var ttnProductsBoxAfterInsertHooks []TTNProductsBoxHook

var ttnProductsBoxBeforeUpdateMu sync.Mutex
var ttnProductsBoxBeforeUpdateHooks []TTNProductsBoxHook
var ttnProductsBoxAfterUpdateMu sync.Mutex
var ttnProductsBoxAfterUpdateHooks []TTNProductsBoxHook

var ttnProductsBoxBeforeDeleteMu sync.Mutex
var ttnProductsBoxBeforeDeleteHooks []TTNProductsBoxHook
var ttnProductsBoxAfterDeleteMu sync.Mutex
var ttnProductsBoxAfterDeleteHooks []TTNProductsBoxHook

var ttnProductsBoxBeforeUpsertMu sync.Mutex
var ttnProductsBoxBeforeUpsertHooks []TTNProductsBoxHook
var ttnProductsBoxAfterUpsertMu sync.Mutex
var ttnProductsBoxAfterUpsertHooks []TTNProductsBoxHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TTNProductsBox) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TTNProductsBox) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TTNProductsBox) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TTNProductsBox) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TTNProductsBox) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TTNProductsBox) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TTNProductsBox) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TTNProductsBox) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TTNProductsBox) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnProductsBoxAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTTNProductsBoxHook registers your hook function for all future operations.
func AddTTNProductsBoxHook(hookPoint boil.HookPoint, ttnProductsBoxHook TTNProductsBoxHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ttnProductsBoxAfterSelectMu.Lock()
		ttnProductsBoxAfterSelectHooks = append(ttnProductsBoxAfterSelectHooks, ttnProductsBoxHook)
		ttnProductsBoxAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		ttnProductsBoxBeforeInsertMu.Lock()
		ttnProductsBoxBeforeInsertHooks = append(ttnProductsBoxBeforeInsertHooks, ttnProductsBoxHook)
		ttnProductsBoxBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		ttnProductsBoxAfterInsertMu.Lock()
		ttnProductsBoxAfterInsertHooks = append(ttnProductsBoxAfterInsertHooks, ttnProductsBoxHook)
		ttnProductsBoxAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		ttnProductsBoxBeforeUpdateMu.Lock()
		ttnProductsBoxBeforeUpdateHooks = append(ttnProductsBoxBeforeUpdateHooks, ttnProductsBoxHook)
		ttnProductsBoxBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		ttnProductsBoxAfterUpdateMu.Lock()
		ttnProductsBoxAfterUpdateHooks = append(ttnProductsBoxAfterUpdateHooks, ttnProductsBoxHook)
		ttnProductsBoxAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		ttnProductsBoxBeforeDeleteMu.Lock()
		ttnProductsBoxBeforeDeleteHooks = append(ttnProductsBoxBeforeDeleteHooks, ttnProductsBoxHook)
		ttnProductsBoxBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		ttnProductsBoxAfterDeleteMu.Lock()
		ttnProductsBoxAfterDeleteHooks = append(ttnProductsBoxAfterDeleteHooks, ttnProductsBoxHook)
		ttnProductsBoxAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		ttnProductsBoxBeforeUpsertMu.Lock()
		ttnProductsBoxBeforeUpsertHooks = append(ttnProductsBoxBeforeUpsertHooks, ttnProductsBoxHook)
		ttnProductsBoxBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		ttnProductsBoxAfterUpsertMu.Lock()
		ttnProductsBoxAfterUpsertHooks = append(ttnProductsBoxAfterUpsertHooks, ttnProductsBoxHook)
		ttnProductsBoxAfterUpsertMu.Unlock()
	}
}

// OneG returns a single ttnProductsBox record from the query using the global executor.
func (q ttnProductsBoxQuery) OneG(ctx context.Context) (*TTNProductsBox, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ttnProductsBox record from the query.
func (q ttnProductsBoxQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TTNProductsBox, error) {
	o := &TTNProductsBox{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for ttn_products_boxes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TTNProductsBox records from the query using the global executor.
func (q ttnProductsBoxQuery) AllG(ctx context.Context) (TTNProductsBoxSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TTNProductsBox records from the query.
func (q ttnProductsBoxQuery) All(ctx context.Context, exec boil.ContextExecutor) (TTNProductsBoxSlice, error) {
	var o []*TTNProductsBox

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TTNProductsBox slice")
	}

	if len(ttnProductsBoxAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TTNProductsBox records in the query using the global executor
func (q ttnProductsBoxQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TTNProductsBox records in the query.
func (q ttnProductsBoxQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count ttn_products_boxes rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q ttnProductsBoxQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ttnProductsBoxQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if ttn_products_boxes exists")
	}

	return count > 0, nil
}

// TTNProductsBoxes retrieves all the records using an executor.
func TTNProductsBoxes(mods ...qm.QueryMod) ttnProductsBoxQuery {
	mods = append(mods, qm.From("[dbo].[ttn_products_boxes]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[ttn_products_boxes].*"})
	}

	return ttnProductsBoxQuery{q}
}

// FindTTNProductsBoxG retrieves a single record by ID.
func FindTTNProductsBoxG(ctx context.Context, iD int, selectCols ...string) (*TTNProductsBox, error) {
	return FindTTNProductsBox(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTTNProductsBox retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTTNProductsBox(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TTNProductsBox, error) {
	ttnProductsBoxObj := &TTNProductsBox{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[ttn_products_boxes] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ttnProductsBoxObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from ttn_products_boxes")
	}

	if err = ttnProductsBoxObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ttnProductsBoxObj, err
	}

	return ttnProductsBoxObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TTNProductsBox) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TTNProductsBox) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_products_boxes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnProductsBoxColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ttnProductsBoxInsertCacheMut.RLock()
	cache, cached := ttnProductsBoxInsertCache[key]
	ttnProductsBoxInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ttnProductsBoxAllColumns,
			ttnProductsBoxColumnsWithDefault,
			ttnProductsBoxColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, ttnProductsBoxGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[ttn_products_boxes] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[ttn_products_boxes] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into ttn_products_boxes")
	}

	if !cached {
		ttnProductsBoxInsertCacheMut.Lock()
		ttnProductsBoxInsertCache[key] = cache
		ttnProductsBoxInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TTNProductsBox record using the global executor.
// See Update for more documentation.
func (o *TTNProductsBox) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TTNProductsBox.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TTNProductsBox) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ttnProductsBoxUpdateCacheMut.RLock()
	cache, cached := ttnProductsBoxUpdateCache[key]
	ttnProductsBoxUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ttnProductsBoxAllColumns,
			ttnProductsBoxPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, ttnProductsBoxGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update ttn_products_boxes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[ttn_products_boxes] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, ttnProductsBoxPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, append(wl, ttnProductsBoxPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update ttn_products_boxes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for ttn_products_boxes")
	}

	if !cached {
		ttnProductsBoxUpdateCacheMut.Lock()
		ttnProductsBoxUpdateCache[key] = cache
		ttnProductsBoxUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q ttnProductsBoxQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ttnProductsBoxQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for ttn_products_boxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for ttn_products_boxes")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TTNProductsBoxSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TTNProductsBoxSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnProductsBoxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[ttn_products_boxes] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ttnProductsBoxPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in ttnProductsBox slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all ttnProductsBox")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TTNProductsBox) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TTNProductsBox) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_products_boxes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnProductsBoxColumnsWithDefault, o)

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

	ttnProductsBoxUpsertCacheMut.RLock()
	cache, cached := ttnProductsBoxUpsertCache[key]
	ttnProductsBoxUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			ttnProductsBoxAllColumns,
			ttnProductsBoxColumnsWithDefault,
			ttnProductsBoxColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, ttnProductsBoxGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(ttnProductsBoxPrimaryKeyColumns, v) && strmangle.ContainsAny(ttnProductsBoxColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert ttn_products_boxes, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			ttnProductsBoxAllColumns,
			ttnProductsBoxPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, ttnProductsBoxGeneratedColumns)

		ret := strmangle.SetComplement(ttnProductsBoxAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert ttn_products_boxes, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[ttn_products_boxes]", ttnProductsBoxPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(ttnProductsBoxPrimaryKeyColumns))
		copy(whitelist, ttnProductsBoxPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ttnProductsBoxType, ttnProductsBoxMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert ttn_products_boxes")
	}

	if !cached {
		ttnProductsBoxUpsertCacheMut.Lock()
		ttnProductsBoxUpsertCache[key] = cache
		ttnProductsBoxUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TTNProductsBox record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TTNProductsBox) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TTNProductsBox record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TTNProductsBox) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TTNProductsBox provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ttnProductsBoxPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[ttn_products_boxes] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from ttn_products_boxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for ttn_products_boxes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q ttnProductsBoxQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q ttnProductsBoxQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no ttnProductsBoxQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttn_products_boxes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_products_boxes")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TTNProductsBoxSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TTNProductsBoxSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ttnProductsBoxBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnProductsBoxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[ttn_products_boxes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ttnProductsBoxPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttnProductsBox slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_products_boxes")
	}

	if len(ttnProductsBoxAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TTNProductsBox) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TTNProductsBox provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TTNProductsBox) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTTNProductsBox(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNProductsBoxSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TTNProductsBoxSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNProductsBoxSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TTNProductsBoxSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnProductsBoxPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[ttn_products_boxes].* FROM [dbo].[ttn_products_boxes] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ttnProductsBoxPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TTNProductsBoxSlice")
	}

	*o = slice

	return nil
}

// TTNProductsBoxExistsG checks if the TTNProductsBox row exists.
func TTNProductsBoxExistsG(ctx context.Context, iD int) (bool, error) {
	return TTNProductsBoxExists(ctx, boil.GetContextDB(), iD)
}

// TTNProductsBoxExists checks if the TTNProductsBox row exists.
func TTNProductsBoxExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[ttn_products_boxes] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if ttn_products_boxes exists")
	}

	return exists, nil
}

// Exists checks if the TTNProductsBox row exists.
func (o *TTNProductsBox) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TTNProductsBoxExists(ctx, exec, o.ID)
}
