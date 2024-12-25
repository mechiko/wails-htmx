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

// Task is an object representing the database table.
type Task struct {
	ID                int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt         null.String `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt         null.String `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name              string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Stage             int64       `boil:"stage" json:"stage" toml:"stage" yaml:"stage"`
	TaskStateID       int64       `boil:"task_state_id" json:"task_state_id" toml:"task_state_id" yaml:"task_state_id"`
	Automatic         int64       `boil:"automatic" json:"automatic" toml:"automatic" yaml:"automatic"`
	Fsrarid           string      `boil:"fsrarid" json:"fsrarid" toml:"fsrarid" yaml:"fsrarid"`
	Started           string      `boil:"started" json:"started" toml:"started" yaml:"started"`
	Ended             string      `boil:"ended" json:"ended" toml:"ended" yaml:"ended"`
	StartOn           string      `boil:"start_on" json:"start_on" toml:"start_on" yaml:"start_on"`
	AfterOf           string      `boil:"after_of" json:"after_of" toml:"after_of" yaml:"after_of"`
	IfExists          string      `boil:"if_exists" json:"if_exists" toml:"if_exists" yaml:"if_exists"`
	Active            int64       `boil:"active" json:"active" toml:"active" yaml:"active"`
	Error             string      `boil:"error" json:"error" toml:"error" yaml:"error"`
	Items             int64       `boil:"items" json:"items" toml:"items" yaml:"items"`
	ProgressTimeout   int64       `boil:"progress_timeout" json:"progress_timeout" toml:"progress_timeout" yaml:"progress_timeout"`
	ProgressSend      int64       `boil:"progress_send" json:"progress_send" toml:"progress_send" yaml:"progress_send"`
	ProgressProcessed int64       `boil:"progress_processed" json:"progress_processed" toml:"progress_processed" yaml:"progress_processed"`
	ProgressReceived  int64       `boil:"progress_received" json:"progress_received" toml:"progress_received" yaml:"progress_received"`
	ProgressError     int64       `boil:"progress_error" json:"progress_error" toml:"progress_error" yaml:"progress_error"`
	Ticks             int64       `boil:"ticks" json:"ticks" toml:"ticks" yaml:"ticks"`
	StateTXT          string      `boil:"state_txt" json:"state_txt" toml:"state_txt" yaml:"state_txt"`
	StateJSON         string      `boil:"state_json" json:"state_json" toml:"state_json" yaml:"state_json"`
	Rem               string      `boil:"rem" json:"rem" toml:"rem" yaml:"rem"`

	R *taskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaskColumns = struct {
	ID                string
	CreatedAt         string
	UpdatedAt         string
	Name              string
	Stage             string
	TaskStateID       string
	Automatic         string
	Fsrarid           string
	Started           string
	Ended             string
	StartOn           string
	AfterOf           string
	IfExists          string
	Active            string
	Error             string
	Items             string
	ProgressTimeout   string
	ProgressSend      string
	ProgressProcessed string
	ProgressReceived  string
	ProgressError     string
	Ticks             string
	StateTXT          string
	StateJSON         string
	Rem               string
}{
	ID:                "id",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	Name:              "name",
	Stage:             "stage",
	TaskStateID:       "task_state_id",
	Automatic:         "automatic",
	Fsrarid:           "fsrarid",
	Started:           "started",
	Ended:             "ended",
	StartOn:           "start_on",
	AfterOf:           "after_of",
	IfExists:          "if_exists",
	Active:            "active",
	Error:             "error",
	Items:             "items",
	ProgressTimeout:   "progress_timeout",
	ProgressSend:      "progress_send",
	ProgressProcessed: "progress_processed",
	ProgressReceived:  "progress_received",
	ProgressError:     "progress_error",
	Ticks:             "ticks",
	StateTXT:          "state_txt",
	StateJSON:         "state_json",
	Rem:               "rem",
}

var TaskTableColumns = struct {
	ID                string
	CreatedAt         string
	UpdatedAt         string
	Name              string
	Stage             string
	TaskStateID       string
	Automatic         string
	Fsrarid           string
	Started           string
	Ended             string
	StartOn           string
	AfterOf           string
	IfExists          string
	Active            string
	Error             string
	Items             string
	ProgressTimeout   string
	ProgressSend      string
	ProgressProcessed string
	ProgressReceived  string
	ProgressError     string
	Ticks             string
	StateTXT          string
	StateJSON         string
	Rem               string
}{
	ID:                "task.id",
	CreatedAt:         "task.created_at",
	UpdatedAt:         "task.updated_at",
	Name:              "task.name",
	Stage:             "task.stage",
	TaskStateID:       "task.task_state_id",
	Automatic:         "task.automatic",
	Fsrarid:           "task.fsrarid",
	Started:           "task.started",
	Ended:             "task.ended",
	StartOn:           "task.start_on",
	AfterOf:           "task.after_of",
	IfExists:          "task.if_exists",
	Active:            "task.active",
	Error:             "task.error",
	Items:             "task.items",
	ProgressTimeout:   "task.progress_timeout",
	ProgressSend:      "task.progress_send",
	ProgressProcessed: "task.progress_processed",
	ProgressReceived:  "task.progress_received",
	ProgressError:     "task.progress_error",
	Ticks:             "task.ticks",
	StateTXT:          "task.state_txt",
	StateJSON:         "task.state_json",
	Rem:               "task.rem",
}

// Generated where

var TaskWhere = struct {
	ID                whereHelperint64
	CreatedAt         whereHelpernull_String
	UpdatedAt         whereHelpernull_String
	Name              whereHelperstring
	Stage             whereHelperint64
	TaskStateID       whereHelperint64
	Automatic         whereHelperint64
	Fsrarid           whereHelperstring
	Started           whereHelperstring
	Ended             whereHelperstring
	StartOn           whereHelperstring
	AfterOf           whereHelperstring
	IfExists          whereHelperstring
	Active            whereHelperint64
	Error             whereHelperstring
	Items             whereHelperint64
	ProgressTimeout   whereHelperint64
	ProgressSend      whereHelperint64
	ProgressProcessed whereHelperint64
	ProgressReceived  whereHelperint64
	ProgressError     whereHelperint64
	Ticks             whereHelperint64
	StateTXT          whereHelperstring
	StateJSON         whereHelperstring
	Rem               whereHelperstring
}{
	ID:                whereHelperint64{field: "\"task\".\"id\""},
	CreatedAt:         whereHelpernull_String{field: "\"task\".\"created_at\""},
	UpdatedAt:         whereHelpernull_String{field: "\"task\".\"updated_at\""},
	Name:              whereHelperstring{field: "\"task\".\"name\""},
	Stage:             whereHelperint64{field: "\"task\".\"stage\""},
	TaskStateID:       whereHelperint64{field: "\"task\".\"task_state_id\""},
	Automatic:         whereHelperint64{field: "\"task\".\"automatic\""},
	Fsrarid:           whereHelperstring{field: "\"task\".\"fsrarid\""},
	Started:           whereHelperstring{field: "\"task\".\"started\""},
	Ended:             whereHelperstring{field: "\"task\".\"ended\""},
	StartOn:           whereHelperstring{field: "\"task\".\"start_on\""},
	AfterOf:           whereHelperstring{field: "\"task\".\"after_of\""},
	IfExists:          whereHelperstring{field: "\"task\".\"if_exists\""},
	Active:            whereHelperint64{field: "\"task\".\"active\""},
	Error:             whereHelperstring{field: "\"task\".\"error\""},
	Items:             whereHelperint64{field: "\"task\".\"items\""},
	ProgressTimeout:   whereHelperint64{field: "\"task\".\"progress_timeout\""},
	ProgressSend:      whereHelperint64{field: "\"task\".\"progress_send\""},
	ProgressProcessed: whereHelperint64{field: "\"task\".\"progress_processed\""},
	ProgressReceived:  whereHelperint64{field: "\"task\".\"progress_received\""},
	ProgressError:     whereHelperint64{field: "\"task\".\"progress_error\""},
	Ticks:             whereHelperint64{field: "\"task\".\"ticks\""},
	StateTXT:          whereHelperstring{field: "\"task\".\"state_txt\""},
	StateJSON:         whereHelperstring{field: "\"task\".\"state_json\""},
	Rem:               whereHelperstring{field: "\"task\".\"rem\""},
}

// TaskRels is where relationship names are stored.
var TaskRels = struct {
}{}

// taskR is where relationships are stored.
type taskR struct {
}

// NewStruct creates a new relationship struct
func (*taskR) NewStruct() *taskR {
	return &taskR{}
}

// taskL is where Load methods for each relationship are stored.
type taskL struct{}

var (
	taskAllColumns            = []string{"id", "created_at", "updated_at", "name", "stage", "task_state_id", "automatic", "fsrarid", "started", "ended", "start_on", "after_of", "if_exists", "active", "error", "items", "progress_timeout", "progress_send", "progress_processed", "progress_received", "progress_error", "ticks", "state_txt", "state_json", "rem"}
	taskColumnsWithoutDefault = []string{}
	taskColumnsWithDefault    = []string{"id", "created_at", "updated_at", "name", "stage", "task_state_id", "automatic", "fsrarid", "started", "ended", "start_on", "after_of", "if_exists", "active", "error", "items", "progress_timeout", "progress_send", "progress_processed", "progress_received", "progress_error", "ticks", "state_txt", "state_json", "rem"}
	taskPrimaryKeyColumns     = []string{"id"}
	taskGeneratedColumns      = []string{"id"}
)

type (
	// TaskSlice is an alias for a slice of pointers to Task.
	// This should almost always be used instead of []Task.
	TaskSlice []*Task
	// TaskHook is the signature for custom Task hook methods
	TaskHook func(context.Context, boil.ContextExecutor, *Task) error

	taskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taskType                 = reflect.TypeOf(&Task{})
	taskMapping              = queries.MakeStructMapping(taskType)
	taskPrimaryKeyMapping, _ = queries.BindMapping(taskType, taskMapping, taskPrimaryKeyColumns)
	taskInsertCacheMut       sync.RWMutex
	taskInsertCache          = make(map[string]insertCache)
	taskUpdateCacheMut       sync.RWMutex
	taskUpdateCache          = make(map[string]updateCache)
	taskUpsertCacheMut       sync.RWMutex
	taskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var taskAfterSelectMu sync.Mutex
var taskAfterSelectHooks []TaskHook

var taskBeforeInsertMu sync.Mutex
var taskBeforeInsertHooks []TaskHook
var taskAfterInsertMu sync.Mutex
var taskAfterInsertHooks []TaskHook

var taskBeforeUpdateMu sync.Mutex
var taskBeforeUpdateHooks []TaskHook
var taskAfterUpdateMu sync.Mutex
var taskAfterUpdateHooks []TaskHook

var taskBeforeDeleteMu sync.Mutex
var taskBeforeDeleteHooks []TaskHook
var taskAfterDeleteMu sync.Mutex
var taskAfterDeleteHooks []TaskHook

var taskBeforeUpsertMu sync.Mutex
var taskBeforeUpsertHooks []TaskHook
var taskAfterUpsertMu sync.Mutex
var taskAfterUpsertHooks []TaskHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Task) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Task) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Task) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Task) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Task) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Task) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Task) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Task) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Task) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range taskAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTaskHook registers your hook function for all future operations.
func AddTaskHook(hookPoint boil.HookPoint, taskHook TaskHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		taskAfterSelectMu.Lock()
		taskAfterSelectHooks = append(taskAfterSelectHooks, taskHook)
		taskAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		taskBeforeInsertMu.Lock()
		taskBeforeInsertHooks = append(taskBeforeInsertHooks, taskHook)
		taskBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		taskAfterInsertMu.Lock()
		taskAfterInsertHooks = append(taskAfterInsertHooks, taskHook)
		taskAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		taskBeforeUpdateMu.Lock()
		taskBeforeUpdateHooks = append(taskBeforeUpdateHooks, taskHook)
		taskBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		taskAfterUpdateMu.Lock()
		taskAfterUpdateHooks = append(taskAfterUpdateHooks, taskHook)
		taskAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		taskBeforeDeleteMu.Lock()
		taskBeforeDeleteHooks = append(taskBeforeDeleteHooks, taskHook)
		taskBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		taskAfterDeleteMu.Lock()
		taskAfterDeleteHooks = append(taskAfterDeleteHooks, taskHook)
		taskAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		taskBeforeUpsertMu.Lock()
		taskBeforeUpsertHooks = append(taskBeforeUpsertHooks, taskHook)
		taskBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		taskAfterUpsertMu.Lock()
		taskAfterUpsertHooks = append(taskAfterUpsertHooks, taskHook)
		taskAfterUpsertMu.Unlock()
	}
}

// OneG returns a single task record from the query using the global executor.
func (q taskQuery) OneG(ctx context.Context) (*Task, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single task record from the query.
func (q taskQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Task, error) {
	o := &Task{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: failed to execute a one query for task")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Task records from the query using the global executor.
func (q taskQuery) AllG(ctx context.Context) (TaskSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Task records from the query.
func (q taskQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaskSlice, error) {
	var o []*Task

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rq: failed to assign all query results to Task slice")
	}

	if len(taskAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Task records in the query using the global executor
func (q taskQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Task records in the query.
func (q taskQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to count task rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q taskQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q taskQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rq: failed to check if task exists")
	}

	return count > 0, nil
}

// Tasks retrieves all the records using an executor.
func Tasks(mods ...qm.QueryMod) taskQuery {
	mods = append(mods, qm.From("\"task\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"task\".*"})
	}

	return taskQuery{q}
}

// FindTaskG retrieves a single record by ID.
func FindTaskG(ctx context.Context, iD int64, selectCols ...string) (*Task, error) {
	return FindTask(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTask retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTask(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Task, error) {
	taskObj := &Task{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"task\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taskObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: unable to select from task")
	}

	if err = taskObj.doAfterSelectHooks(ctx, exec); err != nil {
		return taskObj, err
	}

	return taskObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Task) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Task) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no task provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taskInsertCacheMut.RLock()
	cache, cached := taskInsertCache[key]
	taskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taskAllColumns,
			taskColumnsWithDefault,
			taskColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, taskGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(taskType, taskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taskType, taskMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"task\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"task\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "rq: unable to insert into task")
	}

	if !cached {
		taskInsertCacheMut.Lock()
		taskInsertCache[key] = cache
		taskInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Task record using the global executor.
// See Update for more documentation.
func (o *Task) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Task.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Task) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	taskUpdateCacheMut.RLock()
	cache, cached := taskUpdateCache[key]
	taskUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taskAllColumns,
			taskPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, taskGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("rq: unable to update task, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"task\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, taskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taskType, taskMapping, append(wl, taskPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rq: unable to update task row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by update for task")
	}

	if !cached {
		taskUpdateCacheMut.Lock()
		taskUpdateCache[key] = cache
		taskUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q taskQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q taskQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all for task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected for task")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TaskSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaskSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"task\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all in task slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected all in update all task")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Task) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Task) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no task provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(taskColumnsWithDefault, o)

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

	taskUpsertCacheMut.RLock()
	cache, cached := taskUpsertCache[key]
	taskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taskAllColumns,
			taskColumnsWithDefault,
			taskColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			taskAllColumns,
			taskPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("rq: unable to upsert task, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taskPrimaryKeyColumns))
			copy(conflict, taskPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"task\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taskType, taskMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taskType, taskMapping, ret)
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
		return errors.Wrap(err, "rq: unable to upsert task")
	}

	if !cached {
		taskUpsertCacheMut.Lock()
		taskUpsertCache[key] = cache
		taskUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Task record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Task) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Task record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Task) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rq: no Task provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taskPrimaryKeyMapping)
	sql := "DELETE FROM \"task\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete from task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by delete for task")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q taskQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q taskQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rq: no taskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from task")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for task")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TaskSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaskSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(taskBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"task\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from task slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for task")
	}

	if len(taskAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Task) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: no Task provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Task) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTask(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: empty TaskSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaskSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"task\".* FROM \"task\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rq: unable to reload all in TaskSlice")
	}

	*o = slice

	return nil
}

// TaskExistsG checks if the Task row exists.
func TaskExistsG(ctx context.Context, iD int64) (bool, error) {
	return TaskExists(ctx, boil.GetContextDB(), iD)
}

// TaskExists checks if the Task row exists.
func TaskExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"task\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rq: unable to check if task exists")
	}

	return exists, nil
}

// Exists checks if the Task row exists.
func (o *Task) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TaskExists(ctx, exec, o.ID)
}
