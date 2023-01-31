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

// MWSProduct is an object representing the database table.
type MWSProduct struct {
	Asin        string       `boil:"asin" json:"asin" toml:"asin" yaml:"asin"`
	Filename    string       `boil:"filename" json:"filename" toml:"filename" yaml:"filename"`
	Title       null.String  `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Jan         null.String  `boil:"jan" json:"jan,omitempty" toml:"jan" yaml:"jan,omitempty"`
	Unit        null.Int64   `boil:"unit" json:"unit,omitempty" toml:"unit" yaml:"unit,omitempty"`
	Price       null.Int64   `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
	Cost        null.Int64   `boil:"cost" json:"cost,omitempty" toml:"cost" yaml:"cost,omitempty"`
	FeeRate     null.Float64 `boil:"fee_rate" json:"fee_rate,omitempty" toml:"fee_rate" yaml:"fee_rate,omitempty"`
	ShippingFee null.Int64   `boil:"shipping_fee" json:"shipping_fee,omitempty" toml:"shipping_fee" yaml:"shipping_fee,omitempty"`
	Profit      null.Int64   `boil:"profit" json:"profit,omitempty" toml:"profit" yaml:"profit,omitempty"`
	ProfitRate  null.Float64 `boil:"profit_rate" json:"profit_rate,omitempty" toml:"profit_rate" yaml:"profit_rate,omitempty"`
	CreatedAt   null.Time    `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	URL         null.String  `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`

	R *mwsProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mwsProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MWSProductColumns = struct {
	Asin        string
	Filename    string
	Title       string
	Jan         string
	Unit        string
	Price       string
	Cost        string
	FeeRate     string
	ShippingFee string
	Profit      string
	ProfitRate  string
	CreatedAt   string
	URL         string
}{
	Asin:        "asin",
	Filename:    "filename",
	Title:       "title",
	Jan:         "jan",
	Unit:        "unit",
	Price:       "price",
	Cost:        "cost",
	FeeRate:     "fee_rate",
	ShippingFee: "shipping_fee",
	Profit:      "profit",
	ProfitRate:  "profit_rate",
	CreatedAt:   "created_at",
	URL:         "url",
}

var MWSProductTableColumns = struct {
	Asin        string
	Filename    string
	Title       string
	Jan         string
	Unit        string
	Price       string
	Cost        string
	FeeRate     string
	ShippingFee string
	Profit      string
	ProfitRate  string
	CreatedAt   string
	URL         string
}{
	Asin:        "mws_products.asin",
	Filename:    "mws_products.filename",
	Title:       "mws_products.title",
	Jan:         "mws_products.jan",
	Unit:        "mws_products.unit",
	Price:       "mws_products.price",
	Cost:        "mws_products.cost",
	FeeRate:     "mws_products.fee_rate",
	ShippingFee: "mws_products.shipping_fee",
	Profit:      "mws_products.profit",
	ProfitRate:  "mws_products.profit_rate",
	CreatedAt:   "mws_products.created_at",
	URL:         "mws_products.url",
}

// Generated where

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Float64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Float64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var MWSProductWhere = struct {
	Asin        whereHelperstring
	Filename    whereHelperstring
	Title       whereHelpernull_String
	Jan         whereHelpernull_String
	Unit        whereHelpernull_Int64
	Price       whereHelpernull_Int64
	Cost        whereHelpernull_Int64
	FeeRate     whereHelpernull_Float64
	ShippingFee whereHelpernull_Int64
	Profit      whereHelpernull_Int64
	ProfitRate  whereHelpernull_Float64
	CreatedAt   whereHelpernull_Time
	URL         whereHelpernull_String
}{
	Asin:        whereHelperstring{field: "\"mws_products\".\"asin\""},
	Filename:    whereHelperstring{field: "\"mws_products\".\"filename\""},
	Title:       whereHelpernull_String{field: "\"mws_products\".\"title\""},
	Jan:         whereHelpernull_String{field: "\"mws_products\".\"jan\""},
	Unit:        whereHelpernull_Int64{field: "\"mws_products\".\"unit\""},
	Price:       whereHelpernull_Int64{field: "\"mws_products\".\"price\""},
	Cost:        whereHelpernull_Int64{field: "\"mws_products\".\"cost\""},
	FeeRate:     whereHelpernull_Float64{field: "\"mws_products\".\"fee_rate\""},
	ShippingFee: whereHelpernull_Int64{field: "\"mws_products\".\"shipping_fee\""},
	Profit:      whereHelpernull_Int64{field: "\"mws_products\".\"profit\""},
	ProfitRate:  whereHelpernull_Float64{field: "\"mws_products\".\"profit_rate\""},
	CreatedAt:   whereHelpernull_Time{field: "\"mws_products\".\"created_at\""},
	URL:         whereHelpernull_String{field: "\"mws_products\".\"url\""},
}

// MWSProductRels is where relationship names are stored.
var MWSProductRels = struct {
}{}

// mwsProductR is where relationships are stored.
type mwsProductR struct {
}

// NewStruct creates a new relationship struct
func (*mwsProductR) NewStruct() *mwsProductR {
	return &mwsProductR{}
}

// mwsProductL is where Load methods for each relationship are stored.
type mwsProductL struct{}

var (
	mwsProductAllColumns            = []string{"asin", "filename", "title", "jan", "unit", "price", "cost", "fee_rate", "shipping_fee", "profit", "profit_rate", "created_at", "url"}
	mwsProductColumnsWithoutDefault = []string{"asin", "filename"}
	mwsProductColumnsWithDefault    = []string{"title", "jan", "unit", "price", "cost", "fee_rate", "shipping_fee", "profit", "profit_rate", "created_at", "url"}
	mwsProductPrimaryKeyColumns     = []string{"asin", "filename"}
	mwsProductGeneratedColumns      = []string{"profit", "profit_rate"}
)

type (
	// MWSProductSlice is an alias for a slice of pointers to MWSProduct.
	// This should almost always be used instead of []MWSProduct.
	MWSProductSlice []*MWSProduct
	// MWSProductHook is the signature for custom MWSProduct hook methods
	MWSProductHook func(context.Context, boil.ContextExecutor, *MWSProduct) error

	mwsProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mwsProductType                 = reflect.TypeOf(&MWSProduct{})
	mwsProductMapping              = queries.MakeStructMapping(mwsProductType)
	mwsProductPrimaryKeyMapping, _ = queries.BindMapping(mwsProductType, mwsProductMapping, mwsProductPrimaryKeyColumns)
	mwsProductInsertCacheMut       sync.RWMutex
	mwsProductInsertCache          = make(map[string]insertCache)
	mwsProductUpdateCacheMut       sync.RWMutex
	mwsProductUpdateCache          = make(map[string]updateCache)
	mwsProductUpsertCacheMut       sync.RWMutex
	mwsProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mwsProductAfterSelectHooks []MWSProductHook

var mwsProductBeforeInsertHooks []MWSProductHook
var mwsProductAfterInsertHooks []MWSProductHook

var mwsProductBeforeUpdateHooks []MWSProductHook
var mwsProductAfterUpdateHooks []MWSProductHook

var mwsProductBeforeDeleteHooks []MWSProductHook
var mwsProductAfterDeleteHooks []MWSProductHook

var mwsProductBeforeUpsertHooks []MWSProductHook
var mwsProductAfterUpsertHooks []MWSProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MWSProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MWSProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MWSProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MWSProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MWSProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MWSProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MWSProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MWSProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MWSProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mwsProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMWSProductHook registers your hook function for all future operations.
func AddMWSProductHook(hookPoint boil.HookPoint, mwsProductHook MWSProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		mwsProductAfterSelectHooks = append(mwsProductAfterSelectHooks, mwsProductHook)
	case boil.BeforeInsertHook:
		mwsProductBeforeInsertHooks = append(mwsProductBeforeInsertHooks, mwsProductHook)
	case boil.AfterInsertHook:
		mwsProductAfterInsertHooks = append(mwsProductAfterInsertHooks, mwsProductHook)
	case boil.BeforeUpdateHook:
		mwsProductBeforeUpdateHooks = append(mwsProductBeforeUpdateHooks, mwsProductHook)
	case boil.AfterUpdateHook:
		mwsProductAfterUpdateHooks = append(mwsProductAfterUpdateHooks, mwsProductHook)
	case boil.BeforeDeleteHook:
		mwsProductBeforeDeleteHooks = append(mwsProductBeforeDeleteHooks, mwsProductHook)
	case boil.AfterDeleteHook:
		mwsProductAfterDeleteHooks = append(mwsProductAfterDeleteHooks, mwsProductHook)
	case boil.BeforeUpsertHook:
		mwsProductBeforeUpsertHooks = append(mwsProductBeforeUpsertHooks, mwsProductHook)
	case boil.AfterUpsertHook:
		mwsProductAfterUpsertHooks = append(mwsProductAfterUpsertHooks, mwsProductHook)
	}
}

// One returns a single mwsProduct record from the query.
func (q mwsProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MWSProduct, error) {
	o := &MWSProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for mws_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MWSProduct records from the query.
func (q mwsProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (MWSProductSlice, error) {
	var o []*MWSProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MWSProduct slice")
	}

	if len(mwsProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MWSProduct records in the query.
func (q mwsProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count mws_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mwsProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if mws_products exists")
	}

	return count > 0, nil
}

// MWSProducts retrieves all the records using an executor.
func MWSProducts(mods ...qm.QueryMod) mwsProductQuery {
	mods = append(mods, qm.From("\"mws_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"mws_products\".*"})
	}

	return mwsProductQuery{q}
}

// FindMWSProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMWSProduct(ctx context.Context, exec boil.ContextExecutor, asin string, filename string, selectCols ...string) (*MWSProduct, error) {
	mwsProductObj := &MWSProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"mws_products\" where \"asin\"=$1 AND \"filename\"=$2", sel,
	)

	q := queries.Raw(query, asin, filename)

	err := q.Bind(ctx, exec, mwsProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from mws_products")
	}

	if err = mwsProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return mwsProductObj, err
	}

	return mwsProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MWSProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mws_products provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mwsProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mwsProductInsertCacheMut.RLock()
	cache, cached := mwsProductInsertCache[key]
	mwsProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mwsProductAllColumns,
			mwsProductColumnsWithDefault,
			mwsProductColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, mwsProductGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(mwsProductType, mwsProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mwsProductType, mwsProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"mws_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"mws_products\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into mws_products")
	}

	if !cached {
		mwsProductInsertCacheMut.Lock()
		mwsProductInsertCache[key] = cache
		mwsProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MWSProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MWSProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mwsProductUpdateCacheMut.RLock()
	cache, cached := mwsProductUpdateCache[key]
	mwsProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mwsProductAllColumns,
			mwsProductPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, mwsProductGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update mws_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"mws_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, mwsProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mwsProductType, mwsProductMapping, append(wl, mwsProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update mws_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for mws_products")
	}

	if !cached {
		mwsProductUpdateCacheMut.Lock()
		mwsProductUpdateCache[key] = cache
		mwsProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mwsProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for mws_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for mws_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MWSProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mwsProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"mws_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, mwsProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mwsProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mwsProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MWSProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no mws_products provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mwsProductColumnsWithDefault, o)

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

	mwsProductUpsertCacheMut.RLock()
	cache, cached := mwsProductUpsertCache[key]
	mwsProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mwsProductAllColumns,
			mwsProductColumnsWithDefault,
			mwsProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			mwsProductAllColumns,
			mwsProductPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, mwsProductGeneratedColumns)
		update = strmangle.SetComplement(update, mwsProductGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert mws_products, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(mwsProductPrimaryKeyColumns))
			copy(conflict, mwsProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"mws_products\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(mwsProductType, mwsProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mwsProductType, mwsProductMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert mws_products")
	}

	if !cached {
		mwsProductUpsertCacheMut.Lock()
		mwsProductUpsertCache[key] = cache
		mwsProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MWSProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MWSProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MWSProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mwsProductPrimaryKeyMapping)
	sql := "DELETE FROM \"mws_products\" WHERE \"asin\"=$1 AND \"filename\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from mws_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for mws_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mwsProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mwsProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mws_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mws_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MWSProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(mwsProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mwsProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"mws_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mwsProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mwsProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for mws_products")
	}

	if len(mwsProductAfterDeleteHooks) != 0 {
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
func (o *MWSProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMWSProduct(ctx, exec, o.Asin, o.Filename)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MWSProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MWSProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mwsProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"mws_products\".* FROM \"mws_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, mwsProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MWSProductSlice")
	}

	*o = slice

	return nil
}

// MWSProductExists checks if the MWSProduct row exists.
func MWSProductExists(ctx context.Context, exec boil.ContextExecutor, asin string, filename string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"mws_products\" where \"asin\"=$1 AND \"filename\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, asin, filename)
	}
	row := exec.QueryRowContext(ctx, sql, asin, filename)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if mws_products exists")
	}

	return exists, nil
}

// Exists checks if the MWSProduct row exists.
func (o *MWSProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MWSProductExists(ctx, exec, o.Asin, o.Filename)
}
