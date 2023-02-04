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

// SuperProduct is an object representing the database table.
type SuperProduct struct {
	Name        null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	ProductCode string      `boil:"product_code" json:"product_code" toml:"product_code" yaml:"product_code"`
	Price       null.Int64  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
	ShopCode    null.String `boil:"shop_code" json:"shop_code,omitempty" toml:"shop_code" yaml:"shop_code,omitempty"`
	URL         null.String `boil:"url" json:"url,omitempty" toml:"url" yaml:"url,omitempty"`

	R *superProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L superProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SuperProductColumns = struct {
	Name        string
	ProductCode string
	Price       string
	ShopCode    string
	URL         string
}{
	Name:        "name",
	ProductCode: "product_code",
	Price:       "price",
	ShopCode:    "shop_code",
	URL:         "url",
}

var SuperProductTableColumns = struct {
	Name        string
	ProductCode string
	Price       string
	ShopCode    string
	URL         string
}{
	Name:        "super_products.name",
	ProductCode: "super_products.product_code",
	Price:       "super_products.price",
	ShopCode:    "super_products.shop_code",
	URL:         "super_products.url",
}

// Generated where

var SuperProductWhere = struct {
	Name        whereHelpernull_String
	ProductCode whereHelperstring
	Price       whereHelpernull_Int64
	ShopCode    whereHelpernull_String
	URL         whereHelpernull_String
}{
	Name:        whereHelpernull_String{field: "\"super_products\".\"name\""},
	ProductCode: whereHelperstring{field: "\"super_products\".\"product_code\""},
	Price:       whereHelpernull_Int64{field: "\"super_products\".\"price\""},
	ShopCode:    whereHelpernull_String{field: "\"super_products\".\"shop_code\""},
	URL:         whereHelpernull_String{field: "\"super_products\".\"url\""},
}

// SuperProductRels is where relationship names are stored.
var SuperProductRels = struct {
	ProductCodeSuperProductDetails string
}{
	ProductCodeSuperProductDetails: "ProductCodeSuperProductDetails",
}

// superProductR is where relationships are stored.
type superProductR struct {
	ProductCodeSuperProductDetails SuperProductDetailSlice `boil:"ProductCodeSuperProductDetails" json:"ProductCodeSuperProductDetails" toml:"ProductCodeSuperProductDetails" yaml:"ProductCodeSuperProductDetails"`
}

// NewStruct creates a new relationship struct
func (*superProductR) NewStruct() *superProductR {
	return &superProductR{}
}

func (r *superProductR) GetProductCodeSuperProductDetails() SuperProductDetailSlice {
	if r == nil {
		return nil
	}
	return r.ProductCodeSuperProductDetails
}

// superProductL is where Load methods for each relationship are stored.
type superProductL struct{}

var (
	superProductAllColumns            = []string{"name", "product_code", "price", "shop_code", "url"}
	superProductColumnsWithoutDefault = []string{"product_code"}
	superProductColumnsWithDefault    = []string{"name", "price", "shop_code", "url"}
	superProductPrimaryKeyColumns     = []string{"product_code"}
	superProductGeneratedColumns      = []string{}
)

type (
	// SuperProductSlice is an alias for a slice of pointers to SuperProduct.
	// This should almost always be used instead of []SuperProduct.
	SuperProductSlice []*SuperProduct
	// SuperProductHook is the signature for custom SuperProduct hook methods
	SuperProductHook func(context.Context, boil.ContextExecutor, *SuperProduct) error

	superProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	superProductType                 = reflect.TypeOf(&SuperProduct{})
	superProductMapping              = queries.MakeStructMapping(superProductType)
	superProductPrimaryKeyMapping, _ = queries.BindMapping(superProductType, superProductMapping, superProductPrimaryKeyColumns)
	superProductInsertCacheMut       sync.RWMutex
	superProductInsertCache          = make(map[string]insertCache)
	superProductUpdateCacheMut       sync.RWMutex
	superProductUpdateCache          = make(map[string]updateCache)
	superProductUpsertCacheMut       sync.RWMutex
	superProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var superProductAfterSelectHooks []SuperProductHook

var superProductBeforeInsertHooks []SuperProductHook
var superProductAfterInsertHooks []SuperProductHook

var superProductBeforeUpdateHooks []SuperProductHook
var superProductAfterUpdateHooks []SuperProductHook

var superProductBeforeDeleteHooks []SuperProductHook
var superProductAfterDeleteHooks []SuperProductHook

var superProductBeforeUpsertHooks []SuperProductHook
var superProductAfterUpsertHooks []SuperProductHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SuperProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SuperProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SuperProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SuperProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SuperProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SuperProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SuperProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SuperProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SuperProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range superProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSuperProductHook registers your hook function for all future operations.
func AddSuperProductHook(hookPoint boil.HookPoint, superProductHook SuperProductHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		superProductAfterSelectHooks = append(superProductAfterSelectHooks, superProductHook)
	case boil.BeforeInsertHook:
		superProductBeforeInsertHooks = append(superProductBeforeInsertHooks, superProductHook)
	case boil.AfterInsertHook:
		superProductAfterInsertHooks = append(superProductAfterInsertHooks, superProductHook)
	case boil.BeforeUpdateHook:
		superProductBeforeUpdateHooks = append(superProductBeforeUpdateHooks, superProductHook)
	case boil.AfterUpdateHook:
		superProductAfterUpdateHooks = append(superProductAfterUpdateHooks, superProductHook)
	case boil.BeforeDeleteHook:
		superProductBeforeDeleteHooks = append(superProductBeforeDeleteHooks, superProductHook)
	case boil.AfterDeleteHook:
		superProductAfterDeleteHooks = append(superProductAfterDeleteHooks, superProductHook)
	case boil.BeforeUpsertHook:
		superProductBeforeUpsertHooks = append(superProductBeforeUpsertHooks, superProductHook)
	case boil.AfterUpsertHook:
		superProductAfterUpsertHooks = append(superProductAfterUpsertHooks, superProductHook)
	}
}

// One returns a single superProduct record from the query.
func (q superProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SuperProduct, error) {
	o := &SuperProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for super_products")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SuperProduct records from the query.
func (q superProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (SuperProductSlice, error) {
	var o []*SuperProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SuperProduct slice")
	}

	if len(superProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SuperProduct records in the query.
func (q superProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count super_products rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q superProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if super_products exists")
	}

	return count > 0, nil
}

// ProductCodeSuperProductDetails retrieves all the super_product_detail's SuperProductDetails with an executor via product_code column.
func (o *SuperProduct) ProductCodeSuperProductDetails(mods ...qm.QueryMod) superProductDetailQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"super_product_details\".\"product_code\"=?", o.ProductCode),
	)

	return SuperProductDetails(queryMods...)
}

// LoadProductCodeSuperProductDetails allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (superProductL) LoadProductCodeSuperProductDetails(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSuperProduct interface{}, mods queries.Applicator) error {
	var slice []*SuperProduct
	var object *SuperProduct

	if singular {
		var ok bool
		object, ok = maybeSuperProduct.(*SuperProduct)
		if !ok {
			object = new(SuperProduct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSuperProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSuperProduct))
			}
		}
	} else {
		s, ok := maybeSuperProduct.(*[]*SuperProduct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSuperProduct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSuperProduct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &superProductR{}
		}
		args = append(args, object.ProductCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &superProductR{}
			}

			for _, a := range args {
				if a == obj.ProductCode {
					continue Outer
				}
			}

			args = append(args, obj.ProductCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`super_product_details`),
		qm.WhereIn(`super_product_details.product_code in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load super_product_details")
	}

	var resultSlice []*SuperProductDetail
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice super_product_details")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on super_product_details")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for super_product_details")
	}

	if len(superProductDetailAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProductCodeSuperProductDetails = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &superProductDetailR{}
			}
			foreign.R.ProductCodeSuperProduct = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ProductCode == foreign.ProductCode {
				local.R.ProductCodeSuperProductDetails = append(local.R.ProductCodeSuperProductDetails, foreign)
				if foreign.R == nil {
					foreign.R = &superProductDetailR{}
				}
				foreign.R.ProductCodeSuperProduct = local
				break
			}
		}
	}

	return nil
}

// AddProductCodeSuperProductDetails adds the given related objects to the existing relationships
// of the super_product, optionally inserting them as new records.
// Appends related to o.R.ProductCodeSuperProductDetails.
// Sets related.R.ProductCodeSuperProduct appropriately.
func (o *SuperProduct) AddProductCodeSuperProductDetails(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SuperProductDetail) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProductCode = o.ProductCode
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"super_product_details\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"product_code"}),
				strmangle.WhereClause("\"", "\"", 2, superProductDetailPrimaryKeyColumns),
			)
			values := []interface{}{o.ProductCode, rel.ProductCode, rel.SetNumber}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProductCode = o.ProductCode
		}
	}

	if o.R == nil {
		o.R = &superProductR{
			ProductCodeSuperProductDetails: related,
		}
	} else {
		o.R.ProductCodeSuperProductDetails = append(o.R.ProductCodeSuperProductDetails, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &superProductDetailR{
				ProductCodeSuperProduct: o,
			}
		} else {
			rel.R.ProductCodeSuperProduct = o
		}
	}
	return nil
}

// SuperProducts retrieves all the records using an executor.
func SuperProducts(mods ...qm.QueryMod) superProductQuery {
	mods = append(mods, qm.From("\"super_products\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"super_products\".*"})
	}

	return superProductQuery{q}
}

// FindSuperProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSuperProduct(ctx context.Context, exec boil.ContextExecutor, productCode string, selectCols ...string) (*SuperProduct, error) {
	superProductObj := &SuperProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"super_products\" where \"product_code\"=$1", sel,
	)

	q := queries.Raw(query, productCode)

	err := q.Bind(ctx, exec, superProductObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from super_products")
	}

	if err = superProductObj.doAfterSelectHooks(ctx, exec); err != nil {
		return superProductObj, err
	}

	return superProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SuperProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no super_products provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(superProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	superProductInsertCacheMut.RLock()
	cache, cached := superProductInsertCache[key]
	superProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			superProductAllColumns,
			superProductColumnsWithDefault,
			superProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(superProductType, superProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(superProductType, superProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"super_products\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"super_products\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into super_products")
	}

	if !cached {
		superProductInsertCacheMut.Lock()
		superProductInsertCache[key] = cache
		superProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SuperProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SuperProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	superProductUpdateCacheMut.RLock()
	cache, cached := superProductUpdateCache[key]
	superProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			superProductAllColumns,
			superProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update super_products, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"super_products\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, superProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(superProductType, superProductMapping, append(wl, superProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update super_products row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for super_products")
	}

	if !cached {
		superProductUpdateCacheMut.Lock()
		superProductUpdateCache[key] = cache
		superProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q superProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for super_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for super_products")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SuperProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), superProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"super_products\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, superProductPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in superProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all superProduct")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SuperProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no super_products provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(superProductColumnsWithDefault, o)

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

	superProductUpsertCacheMut.RLock()
	cache, cached := superProductUpsertCache[key]
	superProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			superProductAllColumns,
			superProductColumnsWithDefault,
			superProductColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			superProductAllColumns,
			superProductPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert super_products, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(superProductPrimaryKeyColumns))
			copy(conflict, superProductPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"super_products\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(superProductType, superProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(superProductType, superProductMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert super_products")
	}

	if !cached {
		superProductUpsertCacheMut.Lock()
		superProductUpsertCache[key] = cache
		superProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SuperProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SuperProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SuperProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), superProductPrimaryKeyMapping)
	sql := "DELETE FROM \"super_products\" WHERE \"product_code\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from super_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for super_products")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q superProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no superProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from super_products")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for super_products")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SuperProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(superProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), superProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"super_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, superProductPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from superProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for super_products")
	}

	if len(superProductAfterDeleteHooks) != 0 {
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
func (o *SuperProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSuperProduct(ctx, exec, o.ProductCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SuperProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SuperProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), superProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"super_products\".* FROM \"super_products\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, superProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SuperProductSlice")
	}

	*o = slice

	return nil
}

// SuperProductExists checks if the SuperProduct row exists.
func SuperProductExists(ctx context.Context, exec boil.ContextExecutor, productCode string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"super_products\" where \"product_code\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, productCode)
	}
	row := exec.QueryRowContext(ctx, sql, productCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if super_products exists")
	}

	return exists, nil
}

// Exists checks if the SuperProduct row exists.
func (o *SuperProduct) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SuperProductExists(ctx, exec, o.ProductCode)
}
