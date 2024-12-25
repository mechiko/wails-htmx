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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TaskName is an object representing the database table.
type TaskName struct {
	Name        string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Timeout     int64  `boil:"timeout" json:"timeout" toml:"timeout" yaml:"timeout"`
	RuleRequest string `boil:"rule_request" json:"rule_request" toml:"rule_request" yaml:"rule_request"`
	Rem         string `boil:"rem" json:"rem" toml:"rem" yaml:"rem"`

	R *taskNameR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taskNameL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaskNameColumns = struct {
	Name        string
	Timeout     string
	RuleRequest string
	Rem         string
}{
	Name:        "name",
	Timeout:     "timeout",
	RuleRequest: "rule_request",
	Rem:         "rem",
}

var TaskNameTableColumns = struct {
	Name        string
	Timeout     string
	RuleRequest string
	Rem         string
}{
	Name:        "task_name.name",
	Timeout:     "task_name.timeout",
	RuleRequest: "task_name.rule_request",
	Rem:         "task_name.rem",
}

// Generated where

var TaskNameWhere = struct {
	Name        whereHelperstring
	Timeout     whereHelperint64
	RuleRequest whereHelperstring
	Rem         whereHelperstring
}{
	Name:        whereHelperstring{field: "\"task_name\".\"name\""},
	Timeout:     whereHelperint64{field: "\"task_name\".\"timeout\""},
	RuleRequest: whereHelperstring{field: "\"task_name\".\"rule_request\""},
	Rem:         whereHelperstring{field: "\"task_name\".\"rem\""},
}

// TaskNameRels is where relationship names are stored.
var TaskNameRels = struct {
}{}

// taskNameR is where relationships are stored.
type taskNameR struct {
}

// NewStruct creates a new relationship struct
func (*taskNameR) NewStruct() *taskNameR {
	return &taskNameR{}
}

// taskNameL is where Load methods for each relationship are stored.
type taskNameL struct{}

var (
	taskNameAllColumns            = []string{"name", "timeout", "rule_request", "rem"}
	taskNameColumnsWithoutDefault = []string{"name"}
	taskNameColumnsWithDefault    = []string{"timeout", "rule_request", "rem"}
	taskNamePrimaryKeyColumns     = []string{"name"}
	taskNameGeneratedColumns      = []string{}
)

type (
	// TaskNameSlice is an alias for a slice of pointers to TaskName.
	// This should almost always be used instead of []TaskName.
	TaskNameSlice []*TaskName
	// TaskNameHook is the signature for custom TaskName hook methods
	TaskNameHook func(context.Context, boil.ContextExecutor, *TaskName) error

	taskNameQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taskNameType                 = reflect.TypeOf(&TaskName{})
	taskNameMapping              = queries.MakeStructMapping(taskNameType)
	taskNamePrimaryKeyMapping, _ = queries.BindMapping(taskNameType, taskNameMapping, taskNamePrimaryKeyColumns)
	taskNameInsertCacheMut       sync.RWMutex
	taskNameInsertCache          = make(map[string]insertCache)
	taskNameUpdateCacheMut       sync.RWMutex
	taskNameUpdateCache          = make(map[string]updateCache)
	taskNameUpsertCacheMut       sync.RWMutex
	taskNameUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taskNameAfterSelectMu sync.Mutex
var taskNameAfterSelectHooks []TaskNameHook

var taskNameBeforeInsertMu sync.Mutex
var taskNameBeforeInsertHooks []TaskNameHook
var taskNameAfterInsertMu sync.Mutex
var taskNameAfterInsertHooks []TaskNameHook

var taskNameBeforeUpdateMu sync.Mutex
var taskNameBeforeUpdateHooks []TaskNameHook
var taskNameAfterUpdateMu sync.Mutex
var taskNameAfterUpdateHooks []TaskNameHook

var taskNameBeforeDeleteMu sync.Mutex
var taskNameBeforeDeleteHooks []TaskNameHook
var taskNameAfterDeleteMu sync.Mutex
var taskNameAfterDeleteHooks []TaskNameHook

var taskNameBeforeUpsertMu sync.Mutex
var taskNameBeforeUpsertHooks []TaskNameHook
var taskNameAfterUpsertMu sync.Mutex
var taskNameAfterUpsertHooks []TaskNameHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TaskName) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TaskName) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TaskName) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TaskName) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TaskName) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TaskName) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TaskName) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TaskName) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TaskName) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskNameAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaskNameHook registers your hook function for all future operations.
func AddTaskNameHook(hookPoint boil.HookPoint, taskNameHook TaskNameHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taskNameAfterSelectMu.Lock()
		taskNameAfterSelectHooks = append(taskNameAfterSelectHooks, taskNameHook)
		taskNameAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		taskNameBeforeInsertMu.Lock()
		taskNameBeforeInsertHooks = append(taskNameBeforeInsertHooks, taskNameHook)
		taskNameBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		taskNameAfterInsertMu.Lock()
		taskNameAfterInsertHooks = append(taskNameAfterInsertHooks, taskNameHook)
		taskNameAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		taskNameBeforeUpdateMu.Lock()
		taskNameBeforeUpdateHooks = append(taskNameBeforeUpdateHooks, taskNameHook)
		taskNameBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		taskNameAfterUpdateMu.Lock()
		taskNameAfterUpdateHooks = append(taskNameAfterUpdateHooks, taskNameHook)
		taskNameAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		taskNameBeforeDeleteMu.Lock()
		taskNameBeforeDeleteHooks = append(taskNameBeforeDeleteHooks, taskNameHook)
		taskNameBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		taskNameAfterDeleteMu.Lock()
		taskNameAfterDeleteHooks = append(taskNameAfterDeleteHooks, taskNameHook)
		taskNameAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		taskNameBeforeUpsertMu.Lock()
		taskNameBeforeUpsertHooks = append(taskNameBeforeUpsertHooks, taskNameHook)
		taskNameBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		taskNameAfterUpsertMu.Lock()
		taskNameAfterUpsertHooks = append(taskNameAfterUpsertHooks, taskNameHook)
		taskNameAfterUpsertMu.Unlock()
	}
}

// OneG returns a single taskName record from the query using the global executor.
func (q taskNameQuery) OneG(ctx context.Context) (*TaskName, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single taskName record from the query.
func (q taskNameQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaskName, error) {
	o := &TaskName{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: failed to execute a one query for task_name")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TaskName records from the query using the global executor.
func (q taskNameQuery) AllG(ctx context.Context) (TaskNameSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TaskName records from the query.
func (q taskNameQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaskNameSlice, error) {
	var o []*TaskName

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rq: failed to assign all query results to TaskName slice")
	}

	if len(taskNameAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TaskName records in the query using the global executor
func (q taskNameQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TaskName records in the query.
func (q taskNameQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to count task_name rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q taskNameQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q taskNameQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rq: failed to check if task_name exists")
	}

	return count > 0, nil
}

// TaskNames retrieves all the records using an executor.
func TaskNames(mods ...qm.QueryMod) taskNameQuery {
	mods = append(mods, qm.From("\"task_name\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"task_name\".*"})
	}

	return taskNameQuery{q}
}

// FindTaskNameG retrieves a single record by ID.
func FindTaskNameG(ctx context.Context, name string, selectCols ...string) (*TaskName, error) {
	return FindTaskName(ctx, boil.GetContextDB(), name, selectCols...)
}

// FindTaskName retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaskName(ctx context.Context, exec boil.ContextExecutor, name string, selectCols ...string) (*TaskName, error) {
	taskNameObj := &TaskName{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"task_name\" where \"name\"=?", sel,
	)

	q := queries.Raw(query, name)

	err := q.Bind(ctx, exec, taskNameObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: unable to select from task_name")
	}

	if err = taskNameObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taskNameObj, err
	}

	return taskNameObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TaskName) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TaskName) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no task_name provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskNameColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taskNameInsertCacheMut.RLock()
	cache, cached := taskNameInsertCache[key]
	taskNameInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taskNameAllColumns,
			taskNameColumnsWithDefault,
			taskNameColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taskNameType, taskNameMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taskNameType, taskNameMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"task_name\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"task_name\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "rq: unable to insert into task_name")
	}

	if !cached {
		taskNameInsertCacheMut.Lock()
		taskNameInsertCache[key] = cache
		taskNameInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TaskName record using the global executor.
// See Update for more documentation.
func (o *TaskName) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TaskName.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TaskName) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taskNameUpdateCacheMut.RLock()
	cache, cached := taskNameUpdateCache[key]
	taskNameUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taskNameAllColumns,
			taskNamePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("rq: unable to update task_name, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"task_name\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, taskNamePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taskNameType, taskNameMapping, append(wl, taskNamePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rq: unable to update task_name row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by update for task_name")
	}

	if !cached {
		taskNameUpdateCacheMut.Lock()
		taskNameUpdateCache[key] = cache
		taskNameUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q taskNameQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q taskNameQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all for task_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected for task_name")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TaskNameSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaskNameSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"task_name\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskNamePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all in taskName slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected all in update all taskName")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TaskName) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TaskName) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no task_name provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskNameColumnsWithDefault, o)

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

	taskNameUpsertCacheMut.RLock()
	cache, cached := taskNameUpsertCache[key]
	taskNameUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taskNameAllColumns,
			taskNameColumnsWithDefault,
			taskNameColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			taskNameAllColumns,
			taskNamePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("rq: unable to upsert task_name, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taskNamePrimaryKeyColumns))
			copy(conflict, taskNamePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"task_name\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taskNameType, taskNameMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taskNameType, taskNameMapping, ret)
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
		return errors.Wrap(err, "rq: unable to upsert task_name")
	}

	if !cached {
		taskNameUpsertCacheMut.Lock()
		taskNameUpsertCache[key] = cache
		taskNameUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TaskName record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TaskName) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TaskName record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TaskName) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rq: no TaskName provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taskNamePrimaryKeyMapping)
	sql := "DELETE FROM \"task_name\" WHERE \"name\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete from task_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by delete for task_name")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q taskNameQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q taskNameQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rq: no taskNameQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from task_name")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for task_name")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TaskNameSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaskNameSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taskNameBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"task_name\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskNamePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from taskName slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for task_name")
	}

	if len(taskNameAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TaskName) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: no TaskName provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TaskName) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaskName(ctx, exec, o.Name)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskNameSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: empty TaskNameSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskNameSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaskNameSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskNamePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"task_name\".* FROM \"task_name\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskNamePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rq: unable to reload all in TaskNameSlice")
	}

	*o = slice

	return nil
}

// TaskNameExistsG checks if the TaskName row exists.
func TaskNameExistsG(ctx context.Context, name string) (bool, error) {
	return TaskNameExists(ctx, boil.GetContextDB(), name)
}

// TaskNameExists checks if the TaskName row exists.
func TaskNameExists(ctx context.Context, exec boil.ContextExecutor, name string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"task_name\" where \"name\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, name)
	}
	row := exec.QueryRowContext(ctx, sql, name)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rq: unable to check if task_name exists")
	}

	return exists, nil
}

// Exists checks if the TaskName row exists.
func (o *TaskName) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaskNameExists(ctx, exec, o.Name)
}