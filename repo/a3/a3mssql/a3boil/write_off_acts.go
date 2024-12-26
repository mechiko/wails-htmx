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

// WriteOffAct is an object representing the database table.
type WriteOffAct struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreateDate  null.String `boil:"create_date" json:"create_date,omitempty" toml:"create_date" yaml:"create_date,omitempty"`
	DocIddomain null.String `boil:"doc_iddomain" json:"doc_iddomain,omitempty" toml:"doc_iddomain" yaml:"doc_iddomain,omitempty"`
	DocType     null.String `boil:"doc_type" json:"doc_type,omitempty" toml:"doc_type" yaml:"doc_type,omitempty"`
	DocNumber   null.String `boil:"doc_number" json:"doc_number,omitempty" toml:"doc_number" yaml:"doc_number,omitempty"`
	DocDate     null.String `boil:"doc_date" json:"doc_date,omitempty" toml:"doc_date" yaml:"doc_date,omitempty"`
	DocComment  null.String `boil:"doc_comment" json:"doc_comment,omitempty" toml:"doc_comment" yaml:"doc_comment,omitempty"`
	Version     null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	State       null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Status      null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID     null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`
	Archive     null.Int    `boil:"archive" json:"archive,omitempty" toml:"archive" yaml:"archive,omitempty"`
	XML         null.String `boil:"xml" json:"xml,omitempty" toml:"xml" yaml:"xml,omitempty"`

	R *writeOffActR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L writeOffActL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WriteOffActColumns = struct {
	ID          string
	CreateDate  string
	DocIddomain string
	DocType     string
	DocNumber   string
	DocDate     string
	DocComment  string
	Version     string
	State       string
	Status      string
	ReplyID     string
	Archive     string
	XML         string
}{
	ID:          "id",
	CreateDate:  "create_date",
	DocIddomain: "doc_iddomain",
	DocType:     "doc_type",
	DocNumber:   "doc_number",
	DocDate:     "doc_date",
	DocComment:  "doc_comment",
	Version:     "version",
	State:       "state",
	Status:      "status",
	ReplyID:     "reply_id",
	Archive:     "archive",
	XML:         "xml",
}

var WriteOffActTableColumns = struct {
	ID          string
	CreateDate  string
	DocIddomain string
	DocType     string
	DocNumber   string
	DocDate     string
	DocComment  string
	Version     string
	State       string
	Status      string
	ReplyID     string
	Archive     string
	XML         string
}{
	ID:          "write_off_acts.id",
	CreateDate:  "write_off_acts.create_date",
	DocIddomain: "write_off_acts.doc_iddomain",
	DocType:     "write_off_acts.doc_type",
	DocNumber:   "write_off_acts.doc_number",
	DocDate:     "write_off_acts.doc_date",
	DocComment:  "write_off_acts.doc_comment",
	Version:     "write_off_acts.version",
	State:       "write_off_acts.state",
	Status:      "write_off_acts.status",
	ReplyID:     "write_off_acts.reply_id",
	Archive:     "write_off_acts.archive",
	XML:         "write_off_acts.xml",
}

// Generated where

var WriteOffActWhere = struct {
	ID          whereHelperint
	CreateDate  whereHelpernull_String
	DocIddomain whereHelpernull_String
	DocType     whereHelpernull_String
	DocNumber   whereHelpernull_String
	DocDate     whereHelpernull_String
	DocComment  whereHelpernull_String
	Version     whereHelpernull_String
	State       whereHelpernull_String
	Status      whereHelpernull_String
	ReplyID     whereHelpernull_String
	Archive     whereHelpernull_Int
	XML         whereHelpernull_String
}{
	ID:          whereHelperint{field: "[dbo].[write_off_acts].[id]"},
	CreateDate:  whereHelpernull_String{field: "[dbo].[write_off_acts].[create_date]"},
	DocIddomain: whereHelpernull_String{field: "[dbo].[write_off_acts].[doc_iddomain]"},
	DocType:     whereHelpernull_String{field: "[dbo].[write_off_acts].[doc_type]"},
	DocNumber:   whereHelpernull_String{field: "[dbo].[write_off_acts].[doc_number]"},
	DocDate:     whereHelpernull_String{field: "[dbo].[write_off_acts].[doc_date]"},
	DocComment:  whereHelpernull_String{field: "[dbo].[write_off_acts].[doc_comment]"},
	Version:     whereHelpernull_String{field: "[dbo].[write_off_acts].[version]"},
	State:       whereHelpernull_String{field: "[dbo].[write_off_acts].[state]"},
	Status:      whereHelpernull_String{field: "[dbo].[write_off_acts].[status]"},
	ReplyID:     whereHelpernull_String{field: "[dbo].[write_off_acts].[reply_id]"},
	Archive:     whereHelpernull_Int{field: "[dbo].[write_off_acts].[archive]"},
	XML:         whereHelpernull_String{field: "[dbo].[write_off_acts].[xml]"},
}

// WriteOffActRels is where relationship names are stored.
var WriteOffActRels = struct {
}{}

// writeOffActR is where relationships are stored.
type writeOffActR struct {
}

// NewStruct creates a new relationship struct
func (*writeOffActR) NewStruct() *writeOffActR {
	return &writeOffActR{}
}

// writeOffActL is where Load methods for each relationship are stored.
type writeOffActL struct{}

var (
	writeOffActAllColumns            = []string{"id", "create_date", "doc_iddomain", "doc_type", "doc_number", "doc_date", "doc_comment", "version", "state", "status", "reply_id", "archive", "xml"}
	writeOffActColumnsWithoutDefault = []string{"create_date", "doc_iddomain", "doc_type", "doc_number", "doc_date", "doc_comment", "version", "state", "status", "reply_id", "archive", "xml"}
	writeOffActColumnsWithDefault    = []string{"id"}
	writeOffActPrimaryKeyColumns     = []string{"id"}
	writeOffActGeneratedColumns      = []string{"id"}
)

type (
	// WriteOffActSlice is an alias for a slice of pointers to WriteOffAct.
	// This should almost always be used instead of []WriteOffAct.
	WriteOffActSlice []*WriteOffAct
	// WriteOffActHook is the signature for custom WriteOffAct hook methods
	WriteOffActHook func(context.Context, boil.ContextExecutor, *WriteOffAct) error

	writeOffActQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	writeOffActType                 = reflect.TypeOf(&WriteOffAct{})
	writeOffActMapping              = queries.MakeStructMapping(writeOffActType)
	writeOffActPrimaryKeyMapping, _ = queries.BindMapping(writeOffActType, writeOffActMapping, writeOffActPrimaryKeyColumns)
	writeOffActInsertCacheMut       sync.RWMutex
	writeOffActInsertCache          = make(map[string]insertCache)
	writeOffActUpdateCacheMut       sync.RWMutex
	writeOffActUpdateCache          = make(map[string]updateCache)
	writeOffActUpsertCacheMut       sync.RWMutex
	writeOffActUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var writeOffActAfterSelectMu sync.Mutex
var writeOffActAfterSelectHooks []WriteOffActHook

var writeOffActBeforeInsertMu sync.Mutex
var writeOffActBeforeInsertHooks []WriteOffActHook
var writeOffActAfterInsertMu sync.Mutex
var writeOffActAfterInsertHooks []WriteOffActHook

var writeOffActBeforeUpdateMu sync.Mutex
var writeOffActBeforeUpdateHooks []WriteOffActHook
var writeOffActAfterUpdateMu sync.Mutex
var writeOffActAfterUpdateHooks []WriteOffActHook

var writeOffActBeforeDeleteMu sync.Mutex
var writeOffActBeforeDeleteHooks []WriteOffActHook
var writeOffActAfterDeleteMu sync.Mutex
var writeOffActAfterDeleteHooks []WriteOffActHook

var writeOffActBeforeUpsertMu sync.Mutex
var writeOffActBeforeUpsertHooks []WriteOffActHook
var writeOffActAfterUpsertMu sync.Mutex
var writeOffActAfterUpsertHooks []WriteOffActHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *WriteOffAct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *WriteOffAct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *WriteOffAct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *WriteOffAct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *WriteOffAct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *WriteOffAct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *WriteOffAct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *WriteOffAct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *WriteOffAct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range writeOffActAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWriteOffActHook registers your hook function for all future operations.
func AddWriteOffActHook(hookPoint boil.HookPoint, writeOffActHook WriteOffActHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		writeOffActAfterSelectMu.Lock()
		writeOffActAfterSelectHooks = append(writeOffActAfterSelectHooks, writeOffActHook)
		writeOffActAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		writeOffActBeforeInsertMu.Lock()
		writeOffActBeforeInsertHooks = append(writeOffActBeforeInsertHooks, writeOffActHook)
		writeOffActBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		writeOffActAfterInsertMu.Lock()
		writeOffActAfterInsertHooks = append(writeOffActAfterInsertHooks, writeOffActHook)
		writeOffActAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		writeOffActBeforeUpdateMu.Lock()
		writeOffActBeforeUpdateHooks = append(writeOffActBeforeUpdateHooks, writeOffActHook)
		writeOffActBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		writeOffActAfterUpdateMu.Lock()
		writeOffActAfterUpdateHooks = append(writeOffActAfterUpdateHooks, writeOffActHook)
		writeOffActAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		writeOffActBeforeDeleteMu.Lock()
		writeOffActBeforeDeleteHooks = append(writeOffActBeforeDeleteHooks, writeOffActHook)
		writeOffActBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		writeOffActAfterDeleteMu.Lock()
		writeOffActAfterDeleteHooks = append(writeOffActAfterDeleteHooks, writeOffActHook)
		writeOffActAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		writeOffActBeforeUpsertMu.Lock()
		writeOffActBeforeUpsertHooks = append(writeOffActBeforeUpsertHooks, writeOffActHook)
		writeOffActBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		writeOffActAfterUpsertMu.Lock()
		writeOffActAfterUpsertHooks = append(writeOffActAfterUpsertHooks, writeOffActHook)
		writeOffActAfterUpsertMu.Unlock()
	}
}

// OneG returns a single writeOffAct record from the query using the global executor.
func (q writeOffActQuery) OneG(ctx context.Context) (*WriteOffAct, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single writeOffAct record from the query.
func (q writeOffActQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WriteOffAct, error) {
	o := &WriteOffAct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for write_off_acts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all WriteOffAct records from the query using the global executor.
func (q writeOffActQuery) AllG(ctx context.Context) (WriteOffActSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all WriteOffAct records from the query.
func (q writeOffActQuery) All(ctx context.Context, exec boil.ContextExecutor) (WriteOffActSlice, error) {
	var o []*WriteOffAct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to WriteOffAct slice")
	}

	if len(writeOffActAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all WriteOffAct records in the query using the global executor
func (q writeOffActQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all WriteOffAct records in the query.
func (q writeOffActQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count write_off_acts rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q writeOffActQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q writeOffActQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if write_off_acts exists")
	}

	return count > 0, nil
}

// WriteOffActs retrieves all the records using an executor.
func WriteOffActs(mods ...qm.QueryMod) writeOffActQuery {
	mods = append(mods, qm.From("[dbo].[write_off_acts]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[write_off_acts].*"})
	}

	return writeOffActQuery{q}
}

// FindWriteOffActG retrieves a single record by ID.
func FindWriteOffActG(ctx context.Context, iD int, selectCols ...string) (*WriteOffAct, error) {
	return FindWriteOffAct(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindWriteOffAct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWriteOffAct(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*WriteOffAct, error) {
	writeOffActObj := &WriteOffAct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[write_off_acts] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, writeOffActObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from write_off_acts")
	}

	if err = writeOffActObj.doAfterSelectHooks(ctx, exec); err != nil {
		return writeOffActObj, err
	}

	return writeOffActObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *WriteOffAct) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WriteOffAct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no write_off_acts provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(writeOffActColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	writeOffActInsertCacheMut.RLock()
	cache, cached := writeOffActInsertCache[key]
	writeOffActInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			writeOffActAllColumns,
			writeOffActColumnsWithDefault,
			writeOffActColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, writeOffActGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(writeOffActType, writeOffActMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(writeOffActType, writeOffActMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[write_off_acts] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[write_off_acts] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into write_off_acts")
	}

	if !cached {
		writeOffActInsertCacheMut.Lock()
		writeOffActInsertCache[key] = cache
		writeOffActInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single WriteOffAct record using the global executor.
// See Update for more documentation.
func (o *WriteOffAct) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the WriteOffAct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WriteOffAct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	writeOffActUpdateCacheMut.RLock()
	cache, cached := writeOffActUpdateCache[key]
	writeOffActUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			writeOffActAllColumns,
			writeOffActPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, writeOffActGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update write_off_acts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[write_off_acts] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, writeOffActPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(writeOffActType, writeOffActMapping, append(wl, writeOffActPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update write_off_acts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for write_off_acts")
	}

	if !cached {
		writeOffActUpdateCacheMut.Lock()
		writeOffActUpdateCache[key] = cache
		writeOffActUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q writeOffActQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q writeOffActQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for write_off_acts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for write_off_acts")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o WriteOffActSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WriteOffActSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), writeOffActPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[write_off_acts] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, writeOffActPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in writeOffAct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all writeOffAct")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *WriteOffAct) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *WriteOffAct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no write_off_acts provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(writeOffActColumnsWithDefault, o)

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

	writeOffActUpsertCacheMut.RLock()
	cache, cached := writeOffActUpsertCache[key]
	writeOffActUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			writeOffActAllColumns,
			writeOffActColumnsWithDefault,
			writeOffActColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, writeOffActGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(writeOffActPrimaryKeyColumns, v) && strmangle.ContainsAny(writeOffActColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert write_off_acts, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			writeOffActAllColumns,
			writeOffActPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, writeOffActGeneratedColumns)

		ret := strmangle.SetComplement(writeOffActAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert write_off_acts, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[write_off_acts]", writeOffActPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(writeOffActPrimaryKeyColumns))
		copy(whitelist, writeOffActPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(writeOffActType, writeOffActMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(writeOffActType, writeOffActMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert write_off_acts")
	}

	if !cached {
		writeOffActUpsertCacheMut.Lock()
		writeOffActUpsertCache[key] = cache
		writeOffActUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single WriteOffAct record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *WriteOffAct) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single WriteOffAct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WriteOffAct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no WriteOffAct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), writeOffActPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[write_off_acts] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from write_off_acts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for write_off_acts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q writeOffActQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q writeOffActQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no writeOffActQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from write_off_acts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for write_off_acts")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o WriteOffActSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WriteOffActSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(writeOffActBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), writeOffActPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[write_off_acts] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, writeOffActPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from writeOffAct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for write_off_acts")
	}

	if len(writeOffActAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *WriteOffAct) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no WriteOffAct provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WriteOffAct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWriteOffAct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WriteOffActSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty WriteOffActSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WriteOffActSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WriteOffActSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), writeOffActPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[write_off_acts].* FROM [dbo].[write_off_acts] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, writeOffActPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in WriteOffActSlice")
	}

	*o = slice

	return nil
}

// WriteOffActExistsG checks if the WriteOffAct row exists.
func WriteOffActExistsG(ctx context.Context, iD int) (bool, error) {
	return WriteOffActExists(ctx, boil.GetContextDB(), iD)
}

// WriteOffActExists checks if the WriteOffAct row exists.
func WriteOffActExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[write_off_acts] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if write_off_acts exists")
	}

	return exists, nil
}

// Exists checks if the WriteOffAct row exists.
func (o *WriteOffAct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return WriteOffActExists(ctx, exec, o.ID)
}
