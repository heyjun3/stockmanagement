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

func testSpapiFees(t *testing.T) {
	t.Parallel()

	query := SpapiFees()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpapiFeesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
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

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpapiFeesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpapiFees().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpapiFeesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpapiFeeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpapiFeesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpapiFeeExists(ctx, tx, o.Asin)
	if err != nil {
		t.Errorf("Unable to check if SpapiFee exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpapiFeeExists to return true, but got false.")
	}
}

func testSpapiFeesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spapiFeeFound, err := FindSpapiFee(ctx, tx, o.Asin)
	if err != nil {
		t.Error(err)
	}

	if spapiFeeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpapiFeesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpapiFees().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpapiFeesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpapiFees().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpapiFeesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spapiFeeOne := &SpapiFee{}
	spapiFeeTwo := &SpapiFee{}
	if err = randomize.Struct(seed, spapiFeeOne, spapiFeeDBTypes, false, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}
	if err = randomize.Struct(seed, spapiFeeTwo, spapiFeeDBTypes, false, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spapiFeeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spapiFeeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpapiFees().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpapiFeesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spapiFeeOne := &SpapiFee{}
	spapiFeeTwo := &SpapiFee{}
	if err = randomize.Struct(seed, spapiFeeOne, spapiFeeDBTypes, false, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}
	if err = randomize.Struct(seed, spapiFeeTwo, spapiFeeDBTypes, false, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spapiFeeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spapiFeeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spapiFeeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func spapiFeeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
	*o = SpapiFee{}
	return nil
}

func testSpapiFeesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpapiFee{}
	o := &SpapiFee{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpapiFee object: %s", err)
	}

	AddSpapiFeeHook(boil.BeforeInsertHook, spapiFeeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spapiFeeBeforeInsertHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.AfterInsertHook, spapiFeeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spapiFeeAfterInsertHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.AfterSelectHook, spapiFeeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spapiFeeAfterSelectHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.BeforeUpdateHook, spapiFeeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spapiFeeBeforeUpdateHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.AfterUpdateHook, spapiFeeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spapiFeeAfterUpdateHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.BeforeDeleteHook, spapiFeeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spapiFeeBeforeDeleteHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.AfterDeleteHook, spapiFeeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spapiFeeAfterDeleteHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.BeforeUpsertHook, spapiFeeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spapiFeeBeforeUpsertHooks = []SpapiFeeHook{}

	AddSpapiFeeHook(boil.AfterUpsertHook, spapiFeeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spapiFeeAfterUpsertHooks = []SpapiFeeHook{}
}

func testSpapiFeesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpapiFeesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spapiFeeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpapiFeeToOneAsinsInfoUsingAsinAsinsInfo(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpapiFee
	var foreign AsinsInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spapiFeeDBTypes, false, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, asinsInfoDBTypes, false, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.Asin = foreign.Asin
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AsinAsinsInfo().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Asin != foreign.Asin {
		t.Errorf("want: %v, got %v", foreign.Asin, check.Asin)
	}

	ranAfterSelectHook := false
	AddAsinsInfoHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := SpapiFeeSlice{&local}
	if err = local.L.LoadAsinAsinsInfo(ctx, tx, false, (*[]*SpapiFee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinAsinsInfo == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AsinAsinsInfo = nil
	if err = local.L.LoadAsinAsinsInfo(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinAsinsInfo == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testSpapiFeeToOneSetOpAsinsInfoUsingAsinAsinsInfo(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpapiFee
	var b, c AsinsInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spapiFeeDBTypes, false, strmangle.SetComplement(spapiFeePrimaryKeyColumns, spapiFeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, asinsInfoDBTypes, false, strmangle.SetComplement(asinsInfoPrimaryKeyColumns, asinsInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, asinsInfoDBTypes, false, strmangle.SetComplement(asinsInfoPrimaryKeyColumns, asinsInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AsinsInfo{&b, &c} {
		err = a.SetAsinAsinsInfo(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AsinAsinsInfo != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AsinSpapiFee != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Asin != x.Asin {
			t.Error("foreign key was wrong value", a.Asin)
		}

		if exists, err := SpapiFeeExists(ctx, tx, a.Asin); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testSpapiFeesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
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

func testSpapiFeesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpapiFeeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpapiFeesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpapiFees().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spapiFeeDBTypes = map[string]string{`Asin`: `character varying`, `FeeRate`: `double precision`, `ShipFee`: `bigint`, `Modified`: `date`}
	_               = bytes.MinRead
)

func testSpapiFeesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spapiFeePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spapiFeeAllColumns) == len(spapiFeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpapiFeesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spapiFeeAllColumns) == len(spapiFeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpapiFee{}
	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spapiFeeDBTypes, true, spapiFeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spapiFeeAllColumns, spapiFeePrimaryKeyColumns) {
		fields = spapiFeeAllColumns
	} else {
		fields = strmangle.SetComplement(
			spapiFeeAllColumns,
			spapiFeePrimaryKeyColumns,
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

	slice := SpapiFeeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpapiFeesUpsert(t *testing.T) {
	t.Parallel()

	if len(spapiFeeAllColumns) == len(spapiFeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpapiFee{}
	if err = randomize.Struct(seed, &o, spapiFeeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpapiFee: %s", err)
	}

	count, err := SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spapiFeeDBTypes, false, spapiFeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpapiFee: %s", err)
	}

	count, err = SpapiFees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
