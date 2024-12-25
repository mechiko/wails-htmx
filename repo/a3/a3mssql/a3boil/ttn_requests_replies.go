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

// TTNRequestsReply is an object representing the database table.
type TTNRequestsReply struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IDTTNRequests null.Int    `boil:"id_ttn_requests" json:"id_ttn_requests,omitempty" toml:"id_ttn_requests" yaml:"id_ttn_requests,omitempty"`
	ReplyType     null.String `boil:"reply_type" json:"reply_type,omitempty" toml:"reply_type" yaml:"reply_type,omitempty"`
	ReplyNumber   null.String `boil:"reply_number" json:"reply_number,omitempty" toml:"reply_number" yaml:"reply_number,omitempty"`
	ReplyDate     null.String `boil:"reply_date" json:"reply_date,omitempty" toml:"reply_date" yaml:"reply_date,omitempty"`
	ReplyRegID    null.String `boil:"reply_reg_id" json:"reply_reg_id,omitempty" toml:"reply_reg_id" yaml:"reply_reg_id,omitempty"`
	ReplyComment  null.String `boil:"reply_comment" json:"reply_comment,omitempty" toml:"reply_comment" yaml:"reply_comment,omitempty"`
	Status        null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	ReplyID       null.String `boil:"reply_id" json:"reply_id,omitempty" toml:"reply_id" yaml:"reply_id,omitempty"`

	R *ttnRequestsReplyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ttnRequestsReplyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TTNRequestsReplyColumns = struct {
	ID            string
	IDTTNRequests string
	ReplyType     string
	ReplyNumber   string
	ReplyDate     string
	ReplyRegID    string
	ReplyComment  string
	Status        string
	ReplyID       string
}{
	ID:            "id",
	IDTTNRequests: "id_ttn_requests",
	ReplyType:     "reply_type",
	ReplyNumber:   "reply_number",
	ReplyDate:     "reply_date",
	ReplyRegID:    "reply_reg_id",
	ReplyComment:  "reply_comment",
	Status:        "status",
	ReplyID:       "reply_id",
}

var TTNRequestsReplyTableColumns = struct {
	ID            string
	IDTTNRequests string
	ReplyType     string
	ReplyNumber   string
	ReplyDate     string
	ReplyRegID    string
	ReplyComment  string
	Status        string
	ReplyID       string
}{
	ID:            "ttn_requests_replies.id",
	IDTTNRequests: "ttn_requests_replies.id_ttn_requests",
	ReplyType:     "ttn_requests_replies.reply_type",
	ReplyNumber:   "ttn_requests_replies.reply_number",
	ReplyDate:     "ttn_requests_replies.reply_date",
	ReplyRegID:    "ttn_requests_replies.reply_reg_id",
	ReplyComment:  "ttn_requests_replies.reply_comment",
	Status:        "ttn_requests_replies.status",
	ReplyID:       "ttn_requests_replies.reply_id",
}

// Generated where

var TTNRequestsReplyWhere = struct {
	ID            whereHelperint
	IDTTNRequests whereHelpernull_Int
	ReplyType     whereHelpernull_String
	ReplyNumber   whereHelpernull_String
	ReplyDate     whereHelpernull_String
	ReplyRegID    whereHelpernull_String
	ReplyComment  whereHelpernull_String
	Status        whereHelpernull_String
	ReplyID       whereHelpernull_String
}{
	ID:            whereHelperint{field: "[dbo].[ttn_requests_replies].[id]"},
	IDTTNRequests: whereHelpernull_Int{field: "[dbo].[ttn_requests_replies].[id_ttn_requests]"},
	ReplyType:     whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_type]"},
	ReplyNumber:   whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_number]"},
	ReplyDate:     whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_date]"},
	ReplyRegID:    whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_reg_id]"},
	ReplyComment:  whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_comment]"},
	Status:        whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[status]"},
	ReplyID:       whereHelpernull_String{field: "[dbo].[ttn_requests_replies].[reply_id]"},
}

// TTNRequestsReplyRels is where relationship names are stored.
var TTNRequestsReplyRels = struct {
}{}

// ttnRequestsReplyR is where relationships are stored.
type ttnRequestsReplyR struct {
}

// NewStruct creates a new relationship struct
func (*ttnRequestsReplyR) NewStruct() *ttnRequestsReplyR {
	return &ttnRequestsReplyR{}
}

// ttnRequestsReplyL is where Load methods for each relationship are stored.
type ttnRequestsReplyL struct{}

var (
	ttnRequestsReplyAllColumns            = []string{"id", "id_ttn_requests", "reply_type", "reply_number", "reply_date", "reply_reg_id", "reply_comment", "status", "reply_id"}
	ttnRequestsReplyColumnsWithoutDefault = []string{"id_ttn_requests", "reply_type", "reply_number", "reply_date", "reply_reg_id", "reply_comment", "status", "reply_id"}
	ttnRequestsReplyColumnsWithDefault    = []string{"id"}
	ttnRequestsReplyPrimaryKeyColumns     = []string{"id"}
	ttnRequestsReplyGeneratedColumns      = []string{"id"}
)

type (
	// TTNRequestsReplySlice is an alias for a slice of pointers to TTNRequestsReply.
	// This should almost always be used instead of []TTNRequestsReply.
	TTNRequestsReplySlice []*TTNRequestsReply
	// TTNRequestsReplyHook is the signature for custom TTNRequestsReply hook methods
	TTNRequestsReplyHook func(context.Context, boil.ContextExecutor, *TTNRequestsReply) error

	ttnRequestsReplyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ttnRequestsReplyType                 = reflect.TypeOf(&TTNRequestsReply{})
	ttnRequestsReplyMapping              = queries.MakeStructMapping(ttnRequestsReplyType)
	ttnRequestsReplyPrimaryKeyMapping, _ = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, ttnRequestsReplyPrimaryKeyColumns)
	ttnRequestsReplyInsertCacheMut       sync.RWMutex
	ttnRequestsReplyInsertCache          = make(map[string]insertCache)
	ttnRequestsReplyUpdateCacheMut       sync.RWMutex
	ttnRequestsReplyUpdateCache          = make(map[string]updateCache)
	ttnRequestsReplyUpsertCacheMut       sync.RWMutex
	ttnRequestsReplyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ttnRequestsReplyAfterSelectMu sync.Mutex
var ttnRequestsReplyAfterSelectHooks []TTNRequestsReplyHook

var ttnRequestsReplyBeforeInsertMu sync.Mutex
var ttnRequestsReplyBeforeInsertHooks []TTNRequestsReplyHook
var ttnRequestsReplyAfterInsertMu sync.Mutex
var ttnRequestsReplyAfterInsertHooks []TTNRequestsReplyHook

var ttnRequestsReplyBeforeUpdateMu sync.Mutex
var ttnRequestsReplyBeforeUpdateHooks []TTNRequestsReplyHook
var ttnRequestsReplyAfterUpdateMu sync.Mutex
var ttnRequestsReplyAfterUpdateHooks []TTNRequestsReplyHook

var ttnRequestsReplyBeforeDeleteMu sync.Mutex
var ttnRequestsReplyBeforeDeleteHooks []TTNRequestsReplyHook
var ttnRequestsReplyAfterDeleteMu sync.Mutex
var ttnRequestsReplyAfterDeleteHooks []TTNRequestsReplyHook

var ttnRequestsReplyBeforeUpsertMu sync.Mutex
var ttnRequestsReplyBeforeUpsertHooks []TTNRequestsReplyHook
var ttnRequestsReplyAfterUpsertMu sync.Mutex
var ttnRequestsReplyAfterUpsertHooks []TTNRequestsReplyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TTNRequestsReply) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TTNRequestsReply) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TTNRequestsReply) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TTNRequestsReply) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TTNRequestsReply) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TTNRequestsReply) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TTNRequestsReply) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TTNRequestsReply) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TTNRequestsReply) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ttnRequestsReplyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTTNRequestsReplyHook registers your hook function for all future operations.
func AddTTNRequestsReplyHook(hookPoint boil.HookPoint, ttnRequestsReplyHook TTNRequestsReplyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ttnRequestsReplyAfterSelectMu.Lock()
		ttnRequestsReplyAfterSelectHooks = append(ttnRequestsReplyAfterSelectHooks, ttnRequestsReplyHook)
		ttnRequestsReplyAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		ttnRequestsReplyBeforeInsertMu.Lock()
		ttnRequestsReplyBeforeInsertHooks = append(ttnRequestsReplyBeforeInsertHooks, ttnRequestsReplyHook)
		ttnRequestsReplyBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		ttnRequestsReplyAfterInsertMu.Lock()
		ttnRequestsReplyAfterInsertHooks = append(ttnRequestsReplyAfterInsertHooks, ttnRequestsReplyHook)
		ttnRequestsReplyAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		ttnRequestsReplyBeforeUpdateMu.Lock()
		ttnRequestsReplyBeforeUpdateHooks = append(ttnRequestsReplyBeforeUpdateHooks, ttnRequestsReplyHook)
		ttnRequestsReplyBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		ttnRequestsReplyAfterUpdateMu.Lock()
		ttnRequestsReplyAfterUpdateHooks = append(ttnRequestsReplyAfterUpdateHooks, ttnRequestsReplyHook)
		ttnRequestsReplyAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		ttnRequestsReplyBeforeDeleteMu.Lock()
		ttnRequestsReplyBeforeDeleteHooks = append(ttnRequestsReplyBeforeDeleteHooks, ttnRequestsReplyHook)
		ttnRequestsReplyBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		ttnRequestsReplyAfterDeleteMu.Lock()
		ttnRequestsReplyAfterDeleteHooks = append(ttnRequestsReplyAfterDeleteHooks, ttnRequestsReplyHook)
		ttnRequestsReplyAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		ttnRequestsReplyBeforeUpsertMu.Lock()
		ttnRequestsReplyBeforeUpsertHooks = append(ttnRequestsReplyBeforeUpsertHooks, ttnRequestsReplyHook)
		ttnRequestsReplyBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		ttnRequestsReplyAfterUpsertMu.Lock()
		ttnRequestsReplyAfterUpsertHooks = append(ttnRequestsReplyAfterUpsertHooks, ttnRequestsReplyHook)
		ttnRequestsReplyAfterUpsertMu.Unlock()
	}
}

// OneG returns a single ttnRequestsReply record from the query using the global executor.
func (q ttnRequestsReplyQuery) OneG(ctx context.Context) (*TTNRequestsReply, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ttnRequestsReply record from the query.
func (q ttnRequestsReplyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TTNRequestsReply, error) {
	o := &TTNRequestsReply{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: failed to execute a one query for ttn_requests_replies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TTNRequestsReply records from the query using the global executor.
func (q ttnRequestsReplyQuery) AllG(ctx context.Context) (TTNRequestsReplySlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TTNRequestsReply records from the query.
func (q ttnRequestsReplyQuery) All(ctx context.Context, exec boil.ContextExecutor) (TTNRequestsReplySlice, error) {
	var o []*TTNRequestsReply

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "a3boil: failed to assign all query results to TTNRequestsReply slice")
	}

	if len(ttnRequestsReplyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TTNRequestsReply records in the query using the global executor
func (q ttnRequestsReplyQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TTNRequestsReply records in the query.
func (q ttnRequestsReplyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to count ttn_requests_replies rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q ttnRequestsReplyQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ttnRequestsReplyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: failed to check if ttn_requests_replies exists")
	}

	return count > 0, nil
}

// TTNRequestsReplies retrieves all the records using an executor.
func TTNRequestsReplies(mods ...qm.QueryMod) ttnRequestsReplyQuery {
	mods = append(mods, qm.From("[dbo].[ttn_requests_replies]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[dbo].[ttn_requests_replies].*"})
	}

	return ttnRequestsReplyQuery{q}
}

// FindTTNRequestsReplyG retrieves a single record by ID.
func FindTTNRequestsReplyG(ctx context.Context, iD int, selectCols ...string) (*TTNRequestsReply, error) {
	return FindTTNRequestsReply(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTTNRequestsReply retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTTNRequestsReply(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TTNRequestsReply, error) {
	ttnRequestsReplyObj := &TTNRequestsReply{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [dbo].[ttn_requests_replies] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ttnRequestsReplyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "a3boil: unable to select from ttn_requests_replies")
	}

	if err = ttnRequestsReplyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ttnRequestsReplyObj, err
	}

	return ttnRequestsReplyObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TTNRequestsReply) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TTNRequestsReply) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_requests_replies provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnRequestsReplyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ttnRequestsReplyInsertCacheMut.RLock()
	cache, cached := ttnRequestsReplyInsertCache[key]
	ttnRequestsReplyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ttnRequestsReplyAllColumns,
			ttnRequestsReplyColumnsWithDefault,
			ttnRequestsReplyColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, ttnRequestsReplyGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [dbo].[ttn_requests_replies] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [dbo].[ttn_requests_replies] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "a3boil: unable to insert into ttn_requests_replies")
	}

	if !cached {
		ttnRequestsReplyInsertCacheMut.Lock()
		ttnRequestsReplyInsertCache[key] = cache
		ttnRequestsReplyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TTNRequestsReply record using the global executor.
// See Update for more documentation.
func (o *TTNRequestsReply) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TTNRequestsReply.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TTNRequestsReply) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ttnRequestsReplyUpdateCacheMut.RLock()
	cache, cached := ttnRequestsReplyUpdateCache[key]
	ttnRequestsReplyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ttnRequestsReplyAllColumns,
			ttnRequestsReplyPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, ttnRequestsReplyGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("a3boil: unable to update ttn_requests_replies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [dbo].[ttn_requests_replies] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, ttnRequestsReplyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, append(wl, ttnRequestsReplyPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "a3boil: unable to update ttn_requests_replies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by update for ttn_requests_replies")
	}

	if !cached {
		ttnRequestsReplyUpdateCacheMut.Lock()
		ttnRequestsReplyUpdateCache[key] = cache
		ttnRequestsReplyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q ttnRequestsReplyQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ttnRequestsReplyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all for ttn_requests_replies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected for ttn_requests_replies")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TTNRequestsReplySlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TTNRequestsReplySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnRequestsReplyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [dbo].[ttn_requests_replies] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ttnRequestsReplyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to update all in ttnRequestsReply slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to retrieve rows affected all in update all ttnRequestsReply")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TTNRequestsReply) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *TTNRequestsReply) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("a3boil: no ttn_requests_replies provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ttnRequestsReplyColumnsWithDefault, o)

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

	ttnRequestsReplyUpsertCacheMut.RLock()
	cache, cached := ttnRequestsReplyUpsertCache[key]
	ttnRequestsReplyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			ttnRequestsReplyAllColumns,
			ttnRequestsReplyColumnsWithDefault,
			ttnRequestsReplyColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, ttnRequestsReplyGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(ttnRequestsReplyPrimaryKeyColumns, v) && strmangle.ContainsAny(ttnRequestsReplyColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("a3boil: unable to upsert ttn_requests_replies, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			ttnRequestsReplyAllColumns,
			ttnRequestsReplyPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, ttnRequestsReplyGeneratedColumns)

		ret := strmangle.SetComplement(ttnRequestsReplyAllColumns, strmangle.SetIntersect(insert, update))

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("a3boil: unable to upsert ttn_requests_replies, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[dbo].[ttn_requests_replies]", ttnRequestsReplyPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(ttnRequestsReplyPrimaryKeyColumns))
		copy(whitelist, ttnRequestsReplyPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ttnRequestsReplyType, ttnRequestsReplyMapping, ret)
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
		return errors.Wrap(err, "a3boil: unable to upsert ttn_requests_replies")
	}

	if !cached {
		ttnRequestsReplyUpsertCacheMut.Lock()
		ttnRequestsReplyUpsertCache[key] = cache
		ttnRequestsReplyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TTNRequestsReply record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TTNRequestsReply) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TTNRequestsReply record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TTNRequestsReply) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("a3boil: no TTNRequestsReply provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ttnRequestsReplyPrimaryKeyMapping)
	sql := "DELETE FROM [dbo].[ttn_requests_replies] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete from ttn_requests_replies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by delete for ttn_requests_replies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q ttnRequestsReplyQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q ttnRequestsReplyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("a3boil: no ttnRequestsReplyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttn_requests_replies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_requests_replies")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TTNRequestsReplySlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TTNRequestsReplySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ttnRequestsReplyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnRequestsReplyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [dbo].[ttn_requests_replies] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ttnRequestsReplyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: unable to delete all from ttnRequestsReply slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "a3boil: failed to get rows affected by deleteall for ttn_requests_replies")
	}

	if len(ttnRequestsReplyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TTNRequestsReply) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: no TTNRequestsReply provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TTNRequestsReply) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTTNRequestsReply(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNRequestsReplySlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("a3boil: empty TTNRequestsReplySlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TTNRequestsReplySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TTNRequestsReplySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ttnRequestsReplyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [dbo].[ttn_requests_replies].* FROM [dbo].[ttn_requests_replies] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ttnRequestsReplyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "a3boil: unable to reload all in TTNRequestsReplySlice")
	}

	*o = slice

	return nil
}

// TTNRequestsReplyExistsG checks if the TTNRequestsReply row exists.
func TTNRequestsReplyExistsG(ctx context.Context, iD int) (bool, error) {
	return TTNRequestsReplyExists(ctx, boil.GetContextDB(), iD)
}

// TTNRequestsReplyExists checks if the TTNRequestsReply row exists.
func TTNRequestsReplyExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [dbo].[ttn_requests_replies] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "a3boil: unable to check if ttn_requests_replies exists")
	}

	return exists, nil
}

// Exists checks if the TTNRequestsReply row exists.
func (o *TTNRequestsReply) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TTNRequestsReplyExists(ctx, exec, o.ID)
}