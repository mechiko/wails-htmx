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

// Client is an object representing the database table.
type Client struct {
	Clientregid string      `boil:"clientregid" json:"clientregid" toml:"clientregid" yaml:"clientregid"`
	RequestID   null.Int64  `boil:"request_id" json:"request_id,omitempty" toml:"request_id" yaml:"request_id,omitempty"`
	CreatedAt   null.String `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.String `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Clienttype  null.String `boil:"clienttype" json:"clienttype,omitempty" toml:"clienttype" yaml:"clienttype,omitempty"`
	Inn         null.String `boil:"inn" json:"inn,omitempty" toml:"inn" yaml:"inn,omitempty"`
	KPP         null.String `boil:"kpp" json:"kpp,omitempty" toml:"kpp" yaml:"kpp,omitempty"`
	Fullname    null.String `boil:"fullname" json:"fullname,omitempty" toml:"fullname" yaml:"fullname,omitempty"`
	Shortname   null.String `boil:"shortname" json:"shortname,omitempty" toml:"shortname" yaml:"shortname,omitempty"`
	Country     null.String `boil:"country" json:"country,omitempty" toml:"country" yaml:"country,omitempty"`
	Regioncode  null.String `boil:"regioncode" json:"regioncode,omitempty" toml:"regioncode" yaml:"regioncode,omitempty"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Tsnum       null.String `boil:"tsnum" json:"tsnum,omitempty" toml:"tsnum" yaml:"tsnum,omitempty"`
	State       null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Versionwb   null.String `boil:"versionwb" json:"versionwb,omitempty" toml:"versionwb" yaml:"versionwb,omitempty"`
	Rem         null.String `boil:"rem" json:"rem,omitempty" toml:"rem" yaml:"rem,omitempty"`

	R *clientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClientColumns = struct {
	Clientregid string
	RequestID   string
	CreatedAt   string
	UpdatedAt   string
	Clienttype  string
	Inn         string
	KPP         string
	Fullname    string
	Shortname   string
	Country     string
	Regioncode  string
	Description string
	Tsnum       string
	State       string
	Versionwb   string
	Rem         string
}{
	Clientregid: "clientregid",
	RequestID:   "request_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Clienttype:  "clienttype",
	Inn:         "inn",
	KPP:         "kpp",
	Fullname:    "fullname",
	Shortname:   "shortname",
	Country:     "country",
	Regioncode:  "regioncode",
	Description: "description",
	Tsnum:       "tsnum",
	State:       "state",
	Versionwb:   "versionwb",
	Rem:         "rem",
}

var ClientTableColumns = struct {
	Clientregid string
	RequestID   string
	CreatedAt   string
	UpdatedAt   string
	Clienttype  string
	Inn         string
	KPP         string
	Fullname    string
	Shortname   string
	Country     string
	Regioncode  string
	Description string
	Tsnum       string
	State       string
	Versionwb   string
	Rem         string
}{
	Clientregid: "client.clientregid",
	RequestID:   "client.request_id",
	CreatedAt:   "client.created_at",
	UpdatedAt:   "client.updated_at",
	Clienttype:  "client.clienttype",
	Inn:         "client.inn",
	KPP:         "client.kpp",
	Fullname:    "client.fullname",
	Shortname:   "client.shortname",
	Country:     "client.country",
	Regioncode:  "client.regioncode",
	Description: "client.description",
	Tsnum:       "client.tsnum",
	State:       "client.state",
	Versionwb:   "client.versionwb",
	Rem:         "client.rem",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ClientWhere = struct {
	Clientregid whereHelperstring
	RequestID   whereHelpernull_Int64
	CreatedAt   whereHelpernull_String
	UpdatedAt   whereHelpernull_String
	Clienttype  whereHelpernull_String
	Inn         whereHelpernull_String
	KPP         whereHelpernull_String
	Fullname    whereHelpernull_String
	Shortname   whereHelpernull_String
	Country     whereHelpernull_String
	Regioncode  whereHelpernull_String
	Description whereHelpernull_String
	Tsnum       whereHelpernull_String
	State       whereHelpernull_String
	Versionwb   whereHelpernull_String
	Rem         whereHelpernull_String
}{
	Clientregid: whereHelperstring{field: "\"client\".\"clientregid\""},
	RequestID:   whereHelpernull_Int64{field: "\"client\".\"request_id\""},
	CreatedAt:   whereHelpernull_String{field: "\"client\".\"created_at\""},
	UpdatedAt:   whereHelpernull_String{field: "\"client\".\"updated_at\""},
	Clienttype:  whereHelpernull_String{field: "\"client\".\"clienttype\""},
	Inn:         whereHelpernull_String{field: "\"client\".\"inn\""},
	KPP:         whereHelpernull_String{field: "\"client\".\"kpp\""},
	Fullname:    whereHelpernull_String{field: "\"client\".\"fullname\""},
	Shortname:   whereHelpernull_String{field: "\"client\".\"shortname\""},
	Country:     whereHelpernull_String{field: "\"client\".\"country\""},
	Regioncode:  whereHelpernull_String{field: "\"client\".\"regioncode\""},
	Description: whereHelpernull_String{field: "\"client\".\"description\""},
	Tsnum:       whereHelpernull_String{field: "\"client\".\"tsnum\""},
	State:       whereHelpernull_String{field: "\"client\".\"state\""},
	Versionwb:   whereHelpernull_String{field: "\"client\".\"versionwb\""},
	Rem:         whereHelpernull_String{field: "\"client\".\"rem\""},
}

// ClientRels is where relationship names are stored.
var ClientRels = struct {
}{}

// clientR is where relationships are stored.
type clientR struct {
}

// NewStruct creates a new relationship struct
func (*clientR) NewStruct() *clientR {
	return &clientR{}
}

// clientL is where Load methods for each relationship are stored.
type clientL struct{}

var (
	clientAllColumns            = []string{"clientregid", "request_id", "created_at", "updated_at", "clienttype", "inn", "kpp", "fullname", "shortname", "country", "regioncode", "description", "tsnum", "state", "versionwb", "rem"}
	clientColumnsWithoutDefault = []string{"clientregid"}
	clientColumnsWithDefault    = []string{"request_id", "created_at", "updated_at", "clienttype", "inn", "kpp", "fullname", "shortname", "country", "regioncode", "description", "tsnum", "state", "versionwb", "rem"}
	clientPrimaryKeyColumns     = []string{"clientregid"}
	clientGeneratedColumns      = []string{}
)

type (
	// ClientSlice is an alias for a slice of pointers to Client.
	// This should almost always be used instead of []Client.
	ClientSlice []*Client
	// ClientHook is the signature for custom Client hook methods
	ClientHook func(context.Context, boil.ContextExecutor, *Client) error

	clientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clientType                 = reflect.TypeOf(&Client{})
	clientMapping              = queries.MakeStructMapping(clientType)
	clientPrimaryKeyMapping, _ = queries.BindMapping(clientType, clientMapping, clientPrimaryKeyColumns)
	clientInsertCacheMut       sync.RWMutex
	clientInsertCache          = make(map[string]insertCache)
	clientUpdateCacheMut       sync.RWMutex
	clientUpdateCache          = make(map[string]updateCache)
	clientUpsertCacheMut       sync.RWMutex
	clientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var clientAfterSelectMu sync.Mutex
var clientAfterSelectHooks []ClientHook

var clientBeforeInsertMu sync.Mutex
var clientBeforeInsertHooks []ClientHook
var clientAfterInsertMu sync.Mutex
var clientAfterInsertHooks []ClientHook

var clientBeforeUpdateMu sync.Mutex
var clientBeforeUpdateHooks []ClientHook
var clientAfterUpdateMu sync.Mutex
var clientAfterUpdateHooks []ClientHook

var clientBeforeDeleteMu sync.Mutex
var clientBeforeDeleteHooks []ClientHook
var clientAfterDeleteMu sync.Mutex
var clientAfterDeleteHooks []ClientHook

var clientBeforeUpsertMu sync.Mutex
var clientBeforeUpsertHooks []ClientHook
var clientAfterUpsertMu sync.Mutex
var clientAfterUpsertHooks []ClientHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Client) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Client) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Client) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Client) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Client) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Client) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Client) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Client) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Client) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range clientAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddClientHook registers your hook function for all future operations.
func AddClientHook(hookPoint boil.HookPoint, clientHook ClientHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		clientAfterSelectMu.Lock()
		clientAfterSelectHooks = append(clientAfterSelectHooks, clientHook)
		clientAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		clientBeforeInsertMu.Lock()
		clientBeforeInsertHooks = append(clientBeforeInsertHooks, clientHook)
		clientBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		clientAfterInsertMu.Lock()
		clientAfterInsertHooks = append(clientAfterInsertHooks, clientHook)
		clientAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		clientBeforeUpdateMu.Lock()
		clientBeforeUpdateHooks = append(clientBeforeUpdateHooks, clientHook)
		clientBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		clientAfterUpdateMu.Lock()
		clientAfterUpdateHooks = append(clientAfterUpdateHooks, clientHook)
		clientAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		clientBeforeDeleteMu.Lock()
		clientBeforeDeleteHooks = append(clientBeforeDeleteHooks, clientHook)
		clientBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		clientAfterDeleteMu.Lock()
		clientAfterDeleteHooks = append(clientAfterDeleteHooks, clientHook)
		clientAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		clientBeforeUpsertMu.Lock()
		clientBeforeUpsertHooks = append(clientBeforeUpsertHooks, clientHook)
		clientBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		clientAfterUpsertMu.Lock()
		clientAfterUpsertHooks = append(clientAfterUpsertHooks, clientHook)
		clientAfterUpsertMu.Unlock()
	}
}

// OneG returns a single client record from the query using the global executor.
func (q clientQuery) OneG(ctx context.Context) (*Client, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single client record from the query.
func (q clientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Client, error) {
	o := &Client{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: failed to execute a one query for client")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Client records from the query using the global executor.
func (q clientQuery) AllG(ctx context.Context) (ClientSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Client records from the query.
func (q clientQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClientSlice, error) {
	var o []*Client

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "rq: failed to assign all query results to Client slice")
	}

	if len(clientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Client records in the query using the global executor
func (q clientQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Client records in the query.
func (q clientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to count client rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q clientQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q clientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "rq: failed to check if client exists")
	}

	return count > 0, nil
}

// Clients retrieves all the records using an executor.
func Clients(mods ...qm.QueryMod) clientQuery {
	mods = append(mods, qm.From("\"client\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"client\".*"})
	}

	return clientQuery{q}
}

// FindClientG retrieves a single record by ID.
func FindClientG(ctx context.Context, clientregid string, selectCols ...string) (*Client, error) {
	return FindClient(ctx, boil.GetContextDB(), clientregid, selectCols...)
}

// FindClient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClient(ctx context.Context, exec boil.ContextExecutor, clientregid string, selectCols ...string) (*Client, error) {
	clientObj := &Client{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"client\" where \"clientregid\"=?", sel,
	)

	q := queries.Raw(query, clientregid)

	err := q.Bind(ctx, exec, clientObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "rq: unable to select from client")
	}

	if err = clientObj.doAfterSelectHooks(ctx, exec); err != nil {
		return clientObj, err
	}

	return clientObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Client) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Client) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no client provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(clientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clientInsertCacheMut.RLock()
	cache, cached := clientInsertCache[key]
	clientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clientAllColumns,
			clientColumnsWithDefault,
			clientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clientType, clientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"client\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"client\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "rq: unable to insert into client")
	}

	if !cached {
		clientInsertCacheMut.Lock()
		clientInsertCache[key] = cache
		clientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Client record using the global executor.
// See Update for more documentation.
func (o *Client) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Client.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Client) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	clientUpdateCacheMut.RLock()
	cache, cached := clientUpdateCache[key]
	clientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clientAllColumns,
			clientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("rq: unable to update client, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"client\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, clientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, append(wl, clientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "rq: unable to update client row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by update for client")
	}

	if !cached {
		clientUpdateCacheMut.Lock()
		clientUpdateCache[key] = cache
		clientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q clientQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q clientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all for client")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected for client")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ClientSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ClientSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"client\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to update all in client slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to retrieve rows affected all in update all client")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Client) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Client) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("rq: no client provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(clientColumnsWithDefault, o)

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

	clientUpsertCacheMut.RLock()
	cache, cached := clientUpsertCache[key]
	clientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clientAllColumns,
			clientColumnsWithDefault,
			clientColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			clientAllColumns,
			clientPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("rq: unable to upsert client, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(clientPrimaryKeyColumns))
			copy(conflict, clientPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"client\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(clientType, clientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clientType, clientMapping, ret)
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
		return errors.Wrap(err, "rq: unable to upsert client")
	}

	if !cached {
		clientUpsertCacheMut.Lock()
		clientUpsertCache[key] = cache
		clientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Client record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Client) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Client record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Client) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("rq: no Client provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clientPrimaryKeyMapping)
	sql := "DELETE FROM \"client\" WHERE \"clientregid\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete from client")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by delete for client")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q clientQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q clientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("rq: no clientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from client")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for client")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ClientSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ClientSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(clientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"client\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "rq: unable to delete all from client slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "rq: failed to get rows affected by deleteall for client")
	}

	if len(clientAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Client) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: no Client provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Client) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClient(ctx, exec, o.Clientregid)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("rq: empty ClientSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ClientSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"client\".* FROM \"client\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "rq: unable to reload all in ClientSlice")
	}

	*o = slice

	return nil
}

// ClientExistsG checks if the Client row exists.
func ClientExistsG(ctx context.Context, clientregid string) (bool, error) {
	return ClientExists(ctx, boil.GetContextDB(), clientregid)
}

// ClientExists checks if the Client row exists.
func ClientExists(ctx context.Context, exec boil.ContextExecutor, clientregid string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"client\" where \"clientregid\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, clientregid)
	}
	row := exec.QueryRowContext(ctx, sql, clientregid)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "rq: unable to check if client exists")
	}

	return exists, nil
}

// Exists checks if the Client row exists.
func (o *Client) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ClientExists(ctx, exec, o.Clientregid)
}