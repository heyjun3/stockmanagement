// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// SpapiFee is an object representing the database table.
type SpapiFee struct {
	Asin     string       `boil:"asin" json:"asin" toml:"asin" yaml:"asin"`
	FeeRate  null.Float64 `boil:"fee_rate" json:"fee_rate,omitempty" toml:"fee_rate" yaml:"fee_rate,omitempty"`
	ShipFee  null.Int64   `boil:"ship_fee" json:"ship_fee,omitempty" toml:"ship_fee" yaml:"ship_fee,omitempty"`
	Modified null.Time    `boil:"modified" json:"modified,omitempty" toml:"modified" yaml:"modified,omitempty"`

	R *spapiFeeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spapiFeeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpapiFeeColumns = struct {
	Asin     string
	FeeRate  string
	ShipFee  string
	Modified string
}{
	Asin:     "asin",
	FeeRate:  "fee_rate",
	ShipFee:  "ship_fee",
	Modified: "modified",
}

var SpapiFeeTableColumns = struct {
	Asin     string
	FeeRate  string
	ShipFee  string
	Modified string
}{
	Asin:     "spapi_fees.asin",
	FeeRate:  "spapi_fees.fee_rate",
	ShipFee:  "spapi_fees.ship_fee",
	Modified: "spapi_fees.modified",
}

// Generated where

var SpapiFeeWhere = struct {
	Asin     whereHelperstring
	FeeRate  whereHelpernull_Float64
	ShipFee  whereHelpernull_Int64
	Modified whereHelpernull_Time
}{
	Asin:     whereHelperstring{field: "\"spapi_fees\".\"asin\""},
	FeeRate:  whereHelpernull_Float64{field: "\"spapi_fees\".\"fee_rate\""},
	ShipFee:  whereHelpernull_Int64{field: "\"spapi_fees\".\"ship_fee\""},
	Modified: whereHelpernull_Time{field: "\"spapi_fees\".\"modified\""},
}

// SpapiFeeRels is where relationship names are stored.
var SpapiFeeRels = struct {
	AsinAsinsInfo string
}{
	AsinAsinsInfo: "AsinAsinsInfo",
}

// spapiFeeR is where relationships are stored.
type spapiFeeR struct {
	AsinAsinsInfo *AsinsInfo `boil:"AsinAsinsInfo" json:"AsinAsinsInfo" toml:"AsinAsinsInfo" yaml:"AsinAsinsInfo"`
}

// NewStruct creates a new relationship struct
func (*spapiFeeR) NewStruct() *spapiFeeR {
	return &spapiFeeR{}
}

func (r *spapiFeeR) GetAsinAsinsInfo() *AsinsInfo {
	if r == nil {
		return nil
	}
	return r.AsinAsinsInfo
}

// spapiFeeL is where Load methods for each relationship are stored.
type spapiFeeL struct{}

var (
	spapiFeeAllColumns            = []string{"asin", "fee_rate", "ship_fee", "modified"}
	spapiFeeColumnsWithoutDefault = []string{"asin"}
	spapiFeeColumnsWithDefault    = []string{"fee_rate", "ship_fee", "modified"}
	spapiFeePrimaryKeyColumns     = []string{"asin"}
	spapiFeeGeneratedColumns      = []string{}
)

type (
	// SpapiFeeSlice is an alias for a slice of pointers to SpapiFee.
	// This should almost always be used instead of []SpapiFee.
	SpapiFeeSlice []*SpapiFee
	// SpapiFeeHook is the signature for custom SpapiFee hook methods
	SpapiFeeHook func(context.Context, boil.ContextExecutor, *SpapiFee) error

	spapiFeeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spapiFeeType                 = reflect.TypeOf(&SpapiFee{})
	spapiFeeMapping              = queries.MakeStructMapping(spapiFeeType)
	spapiFeePrimaryKeyMapping, _ = queries.BindMapping(spapiFeeType, spapiFeeMapping, spapiFeePrimaryKeyColumns)
	spapiFeeInsertCacheMut       sync.RWMutex
	spapiFeeInsertCache          = make(map[string]insertCache)
	spapiFeeUpdateCacheMut       sync.RWMutex
	spapiFeeUpdateCache          = make(map[string]updateCache)
	spapiFeeUpsertCacheMut       sync.RWMutex
	spapiFeeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var spapiFeeAfterSelectHooks []SpapiFeeHook

var spapiFeeBeforeInsertHooks []SpapiFeeHook
var spapiFeeAfterInsertHooks []SpapiFeeHook

var spapiFeeBeforeUpdateHooks []SpapiFeeHook
var spapiFeeAfterUpdateHooks []SpapiFeeHook

var spapiFeeBeforeDeleteHooks []SpapiFeeHook
var spapiFeeAfterDeleteHooks []SpapiFeeHook

var spapiFeeBeforeUpsertHooks []SpapiFeeHook
var spapiFeeAfterUpsertHooks []SpapiFeeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SpapiFee) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SpapiFee) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SpapiFee) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SpapiFee) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SpapiFee) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SpapiFee) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SpapiFee) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SpapiFee) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SpapiFee) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spapiFeeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSpapiFeeHook registers your hook function for all future operations.
func AddSpapiFeeHook(hookPoint boil.HookPoint, spapiFeeHook SpapiFeeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		spapiFeeAfterSelectHooks = append(spapiFeeAfterSelectHooks, spapiFeeHook)
	case boil.BeforeInsertHook:
		spapiFeeBeforeInsertHooks = append(spapiFeeBeforeInsertHooks, spapiFeeHook)
	case boil.AfterInsertHook:
		spapiFeeAfterInsertHooks = append(spapiFeeAfterInsertHooks, spapiFeeHook)
	case boil.BeforeUpdateHook:
		spapiFeeBeforeUpdateHooks = append(spapiFeeBeforeUpdateHooks, spapiFeeHook)
	case boil.AfterUpdateHook:
		spapiFeeAfterUpdateHooks = append(spapiFeeAfterUpdateHooks, spapiFeeHook)
	case boil.BeforeDeleteHook:
		spapiFeeBeforeDeleteHooks = append(spapiFeeBeforeDeleteHooks, spapiFeeHook)
	case boil.AfterDeleteHook:
		spapiFeeAfterDeleteHooks = append(spapiFeeAfterDeleteHooks, spapiFeeHook)
	case boil.BeforeUpsertHook:
		spapiFeeBeforeUpsertHooks = append(spapiFeeBeforeUpsertHooks, spapiFeeHook)
	case boil.AfterUpsertHook:
		spapiFeeAfterUpsertHooks = append(spapiFeeAfterUpsertHooks, spapiFeeHook)
	}
}

// One returns a single spapiFee record from the query.
func (q spapiFeeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SpapiFee, error) {
	o := &SpapiFee{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for spapi_fees")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SpapiFee records from the query.
func (q spapiFeeQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpapiFeeSlice, error) {
	var o []*SpapiFee

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SpapiFee slice")
	}

	if len(spapiFeeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SpapiFee records in the query.
func (q spapiFeeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count spapi_fees rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q spapiFeeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if spapi_fees exists")
	}

	return count > 0, nil
}

// AsinAsinsInfo pointed to by the foreign key.
func (o *SpapiFee) AsinAsinsInfo(mods ...qm.QueryMod) asinsInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"asin\" = ?", o.Asin),
	}

	queryMods = append(queryMods, mods...)

	return AsinsInfos(queryMods...)
}

// LoadAsinAsinsInfo allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spapiFeeL) LoadAsinAsinsInfo(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpapiFee interface{}, mods queries.Applicator) error {
	var slice []*SpapiFee
	var object *SpapiFee

	if singular {
		var ok bool
		object, ok = maybeSpapiFee.(*SpapiFee)
		if !ok {
			object = new(SpapiFee)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSpapiFee)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSpapiFee))
			}
		}
	} else {
		s, ok := maybeSpapiFee.(*[]*SpapiFee)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSpapiFee)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSpapiFee))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spapiFeeR{}
		}
		args = append(args, object.Asin)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spapiFeeR{}
			}

			for _, a := range args {
				if a == obj.Asin {
					continue Outer
				}
			}

			args = append(args, obj.Asin)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`asins_info`),
		qm.WhereIn(`asins_info.asin in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AsinsInfo")
	}

	var resultSlice []*AsinsInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AsinsInfo")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for asins_info")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for asins_info")
	}

	if len(asinsInfoAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.AsinAsinsInfo = foreign
		if foreign.R == nil {
			foreign.R = &asinsInfoR{}
		}
		foreign.R.AsinSpapiFee = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Asin == foreign.Asin {
				local.R.AsinAsinsInfo = foreign
				if foreign.R == nil {
					foreign.R = &asinsInfoR{}
				}
				foreign.R.AsinSpapiFee = local
				break
			}
		}
	}

	return nil
}

// SetAsinAsinsInfo of the spapiFee to the related item.
// Sets o.R.AsinAsinsInfo to related.
// Adds o to related.R.AsinSpapiFee.
func (o *SpapiFee) SetAsinAsinsInfo(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AsinsInfo) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spapi_fees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"asin"}),
		strmangle.WhereClause("\"", "\"", 2, spapiFeePrimaryKeyColumns),
	)
	values := []interface{}{related.Asin, o.Asin}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Asin = related.Asin
	if o.R == nil {
		o.R = &spapiFeeR{
			AsinAsinsInfo: related,
		}
	} else {
		o.R.AsinAsinsInfo = related
	}

	if related.R == nil {
		related.R = &asinsInfoR{
			AsinSpapiFee: o,
		}
	} else {
		related.R.AsinSpapiFee = o
	}

	return nil
}

// SpapiFees retrieves all the records using an executor.
func SpapiFees(mods ...qm.QueryMod) spapiFeeQuery {
	mods = append(mods, qm.From("\"spapi_fees\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"spapi_fees\".*"})
	}

	return spapiFeeQuery{q}
}

// FindSpapiFee retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpapiFee(ctx context.Context, exec boil.ContextExecutor, asin string, selectCols ...string) (*SpapiFee, error) {
	spapiFeeObj := &SpapiFee{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"spapi_fees\" where \"asin\"=$1", sel,
	)

	q := queries.Raw(query, asin)

	err := q.Bind(ctx, exec, spapiFeeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from spapi_fees")
	}

	if err = spapiFeeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return spapiFeeObj, err
	}

	return spapiFeeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SpapiFee) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spapi_fees provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spapiFeeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spapiFeeInsertCacheMut.RLock()
	cache, cached := spapiFeeInsertCache[key]
	spapiFeeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spapiFeeAllColumns,
			spapiFeeColumnsWithDefault,
			spapiFeeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spapiFeeType, spapiFeeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spapiFeeType, spapiFeeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"spapi_fees\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"spapi_fees\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into spapi_fees")
	}

	if !cached {
		spapiFeeInsertCacheMut.Lock()
		spapiFeeInsertCache[key] = cache
		spapiFeeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SpapiFee.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SpapiFee) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	spapiFeeUpdateCacheMut.RLock()
	cache, cached := spapiFeeUpdateCache[key]
	spapiFeeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spapiFeeAllColumns,
			spapiFeePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update spapi_fees, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"spapi_fees\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, spapiFeePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spapiFeeType, spapiFeeMapping, append(wl, spapiFeePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update spapi_fees row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for spapi_fees")
	}

	if !cached {
		spapiFeeUpdateCacheMut.Lock()
		spapiFeeUpdateCache[key] = cache
		spapiFeeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q spapiFeeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for spapi_fees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for spapi_fees")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SpapiFeeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spapiFeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"spapi_fees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, spapiFeePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in spapiFee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all spapiFee")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SpapiFee) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spapi_fees provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(spapiFeeColumnsWithDefault, o)

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

	spapiFeeUpsertCacheMut.RLock()
	cache, cached := spapiFeeUpsertCache[key]
	spapiFeeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spapiFeeAllColumns,
			spapiFeeColumnsWithDefault,
			spapiFeeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			spapiFeeAllColumns,
			spapiFeePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert spapi_fees, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(spapiFeePrimaryKeyColumns))
			copy(conflict, spapiFeePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"spapi_fees\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(spapiFeeType, spapiFeeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spapiFeeType, spapiFeeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert spapi_fees")
	}

	if !cached {
		spapiFeeUpsertCacheMut.Lock()
		spapiFeeUpsertCache[key] = cache
		spapiFeeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SpapiFee record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SpapiFee) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SpapiFee provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spapiFeePrimaryKeyMapping)
	sql := "DELETE FROM \"spapi_fees\" WHERE \"asin\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from spapi_fees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for spapi_fees")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q spapiFeeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no spapiFeeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spapi_fees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spapi_fees")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SpapiFeeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(spapiFeeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spapiFeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"spapi_fees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spapiFeePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spapiFee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spapi_fees")
	}

	if len(spapiFeeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SpapiFee) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpapiFee(ctx, exec, o.Asin)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpapiFeeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpapiFeeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spapiFeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"spapi_fees\".* FROM \"spapi_fees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spapiFeePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpapiFeeSlice")
	}

	*o = slice

	return nil
}

// SpapiFeeExists checks if the SpapiFee row exists.
func SpapiFeeExists(ctx context.Context, exec boil.ContextExecutor, asin string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"spapi_fees\" where \"asin\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, asin)
	}
	row := exec.QueryRowContext(ctx, sql, asin)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if spapi_fees exists")
	}

	return exists, nil
}

// Exists checks if the SpapiFee row exists.
func (o *SpapiFee) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SpapiFeeExists(ctx, exec, o.Asin)
}
