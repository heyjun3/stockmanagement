// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProductMasters(t *testing.T) {
	t.Parallel()

	query := ProductMasters()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProductMastersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductMastersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProductMasters().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductMastersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProductMasterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProductMastersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProductMasterExists(ctx, tx, o.Sku)
	if err != nil {
		t.Errorf("Unable to check if ProductMaster exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProductMasterExists to return true, but got false.")
	}
}

func testProductMastersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	productMasterFound, err := FindProductMaster(ctx, tx, o.Sku)
	if err != nil {
		t.Error(err)
	}

	if productMasterFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProductMastersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProductMasters().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProductMastersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProductMasters().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProductMastersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	productMasterOne := &ProductMaster{}
	productMasterTwo := &ProductMaster{}
	if err = randomize.Struct(seed, productMasterOne, productMasterDBTypes, false, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}
	if err = randomize.Struct(seed, productMasterTwo, productMasterDBTypes, false, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = productMasterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = productMasterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProductMasters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProductMastersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	productMasterOne := &ProductMaster{}
	productMasterTwo := &ProductMaster{}
	if err = randomize.Struct(seed, productMasterOne, productMasterDBTypes, false, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}
	if err = randomize.Struct(seed, productMasterTwo, productMasterDBTypes, false, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = productMasterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = productMasterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func productMasterBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func productMasterAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProductMaster) error {
	*o = ProductMaster{}
	return nil
}

func testProductMastersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ProductMaster{}
	o := &ProductMaster{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, productMasterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ProductMaster object: %s", err)
	}

	AddProductMasterHook(boil.BeforeInsertHook, productMasterBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	productMasterBeforeInsertHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.AfterInsertHook, productMasterAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	productMasterAfterInsertHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.AfterSelectHook, productMasterAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	productMasterAfterSelectHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.BeforeUpdateHook, productMasterBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	productMasterBeforeUpdateHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.AfterUpdateHook, productMasterAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	productMasterAfterUpdateHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.BeforeDeleteHook, productMasterBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	productMasterBeforeDeleteHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.AfterDeleteHook, productMasterAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	productMasterAfterDeleteHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.BeforeUpsertHook, productMasterBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	productMasterBeforeUpsertHooks = []ProductMasterHook{}

	AddProductMasterHook(boil.AfterUpsertHook, productMasterAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	productMasterAfterUpsertHooks = []ProductMasterHook{}
}

func testProductMastersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProductMastersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(productMasterColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProductMastersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProductMastersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProductMasterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProductMastersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProductMasters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	productMasterDBTypes = map[string]string{`Date`: `integer`, `Name`: `character varying`, `Asin`: `character varying`, `Jan`: `character varying`, `Sku`: `character varying`, `Fnsku`: `character varying`, `DangerClass`: `character varying`, `SellPrice`: `integer`, `CostPrice`: `integer`}
	_                    = bytes.MinRead
)

func testProductMastersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(productMasterPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(productMasterAllColumns) == len(productMasterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProductMastersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(productMasterAllColumns) == len(productMasterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProductMaster{}
	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, productMasterDBTypes, true, productMasterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(productMasterAllColumns, productMasterPrimaryKeyColumns) {
		fields = productMasterAllColumns
	} else {
		fields = strmangle.SetComplement(
			productMasterAllColumns,
			productMasterPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProductMasterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProductMastersUpsert(t *testing.T) {
	t.Parallel()

	if len(productMasterAllColumns) == len(productMasterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ProductMaster{}
	if err = randomize.Struct(seed, &o, productMasterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProductMaster: %s", err)
	}

	count, err := ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, productMasterDBTypes, false, productMasterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProductMaster struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProductMaster: %s", err)
	}

	count, err = ProductMasters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
