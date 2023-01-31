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

func testAsinsInfos(t *testing.T) {
	t.Parallel()

	query := AsinsInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAsinsInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
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

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAsinsInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AsinsInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAsinsInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AsinsInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAsinsInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AsinsInfoExists(ctx, tx, o.Asin)
	if err != nil {
		t.Errorf("Unable to check if AsinsInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AsinsInfoExists to return true, but got false.")
	}
}

func testAsinsInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	asinsInfoFound, err := FindAsinsInfo(ctx, tx, o.Asin)
	if err != nil {
		t.Error(err)
	}

	if asinsInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAsinsInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AsinsInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAsinsInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AsinsInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAsinsInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	asinsInfoOne := &AsinsInfo{}
	asinsInfoTwo := &AsinsInfo{}
	if err = randomize.Struct(seed, asinsInfoOne, asinsInfoDBTypes, false, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, asinsInfoTwo, asinsInfoDBTypes, false, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = asinsInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = asinsInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AsinsInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAsinsInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	asinsInfoOne := &AsinsInfo{}
	asinsInfoTwo := &AsinsInfo{}
	if err = randomize.Struct(seed, asinsInfoOne, asinsInfoDBTypes, false, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, asinsInfoTwo, asinsInfoDBTypes, false, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = asinsInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = asinsInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func asinsInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func asinsInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AsinsInfo) error {
	*o = AsinsInfo{}
	return nil
}

func testAsinsInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AsinsInfo{}
	o := &AsinsInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AsinsInfo object: %s", err)
	}

	AddAsinsInfoHook(boil.BeforeInsertHook, asinsInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	asinsInfoBeforeInsertHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.AfterInsertHook, asinsInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	asinsInfoAfterInsertHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.AfterSelectHook, asinsInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	asinsInfoAfterSelectHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.BeforeUpdateHook, asinsInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	asinsInfoBeforeUpdateHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.AfterUpdateHook, asinsInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	asinsInfoAfterUpdateHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.BeforeDeleteHook, asinsInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	asinsInfoBeforeDeleteHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.AfterDeleteHook, asinsInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	asinsInfoAfterDeleteHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.BeforeUpsertHook, asinsInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	asinsInfoBeforeUpsertHooks = []AsinsInfoHook{}

	AddAsinsInfoHook(boil.AfterUpsertHook, asinsInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	asinsInfoAfterUpsertHooks = []AsinsInfoHook{}
}

func testAsinsInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAsinsInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(asinsInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAsinsInfoOneToOneSpapiFeeUsingAsinSpapiFee(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign SpapiFee
	var local AsinsInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, spapiFeeDBTypes, true, spapiFeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiFee struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.Asin = local.Asin
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AsinSpapiFee().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Asin != foreign.Asin {
		t.Errorf("want: %v, got %v", foreign.Asin, check.Asin)
	}

	ranAfterSelectHook := false
	AddSpapiFeeHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *SpapiFee) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := AsinsInfoSlice{&local}
	if err = local.L.LoadAsinSpapiFee(ctx, tx, false, (*[]*AsinsInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinSpapiFee == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AsinSpapiFee = nil
	if err = local.L.LoadAsinSpapiFee(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinSpapiFee == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testAsinsInfoOneToOneSpapiPriceUsingAsinSpapiPrice(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign SpapiPrice
	var local AsinsInfo

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, spapiPriceDBTypes, true, spapiPriceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpapiPrice struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreign.Asin = local.Asin
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.AsinSpapiPrice().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Asin != foreign.Asin {
		t.Errorf("want: %v, got %v", foreign.Asin, check.Asin)
	}

	ranAfterSelectHook := false
	AddSpapiPriceHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *SpapiPrice) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := AsinsInfoSlice{&local}
	if err = local.L.LoadAsinSpapiPrice(ctx, tx, false, (*[]*AsinsInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinSpapiPrice == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AsinSpapiPrice = nil
	if err = local.L.LoadAsinSpapiPrice(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.AsinSpapiPrice == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testAsinsInfoOneToOneSetOpSpapiFeeUsingAsinSpapiFee(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AsinsInfo
	var b, c SpapiFee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, asinsInfoDBTypes, false, strmangle.SetComplement(asinsInfoPrimaryKeyColumns, asinsInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, spapiFeeDBTypes, false, strmangle.SetComplement(spapiFeePrimaryKeyColumns, spapiFeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, spapiFeeDBTypes, false, strmangle.SetComplement(spapiFeePrimaryKeyColumns, spapiFeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SpapiFee{&b, &c} {
		err = a.SetAsinSpapiFee(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AsinSpapiFee != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.AsinAsinsInfo != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.Asin != x.Asin {
			t.Error("foreign key was wrong value", a.Asin)
		}

		if exists, err := SpapiFeeExists(ctx, tx, x.Asin); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.Asin != x.Asin {
			t.Error("foreign key was wrong value", a.Asin, x.Asin)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}
func testAsinsInfoOneToOneSetOpSpapiPriceUsingAsinSpapiPrice(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AsinsInfo
	var b, c SpapiPrice

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, asinsInfoDBTypes, false, strmangle.SetComplement(asinsInfoPrimaryKeyColumns, asinsInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, spapiPriceDBTypes, false, strmangle.SetComplement(spapiPricePrimaryKeyColumns, spapiPriceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, spapiPriceDBTypes, false, strmangle.SetComplement(spapiPricePrimaryKeyColumns, spapiPriceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SpapiPrice{&b, &c} {
		err = a.SetAsinSpapiPrice(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AsinSpapiPrice != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.AsinAsinsInfo != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if a.Asin != x.Asin {
			t.Error("foreign key was wrong value", a.Asin)
		}

		if exists, err := SpapiPriceExists(ctx, tx, x.Asin); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'x' to exist")
		}

		if a.Asin != x.Asin {
			t.Error("foreign key was wrong value", a.Asin, x.Asin)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testAsinsInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
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

func testAsinsInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AsinsInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAsinsInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AsinsInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	asinsInfoDBTypes = map[string]string{`Asin`: `character varying`, `Jan`: `character varying`, `Title`: `character varying`, `Quantity`: `bigint`, `Modified`: `date`}
	_                = bytes.MinRead
)

func testAsinsInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(asinsInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(asinsInfoAllColumns) == len(asinsInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAsinsInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(asinsInfoAllColumns) == len(asinsInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AsinsInfo{}
	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, asinsInfoDBTypes, true, asinsInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(asinsInfoAllColumns, asinsInfoPrimaryKeyColumns) {
		fields = asinsInfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			asinsInfoAllColumns,
			asinsInfoPrimaryKeyColumns,
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

	slice := AsinsInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAsinsInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(asinsInfoAllColumns) == len(asinsInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AsinsInfo{}
	if err = randomize.Struct(seed, &o, asinsInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AsinsInfo: %s", err)
	}

	count, err := AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, asinsInfoDBTypes, false, asinsInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AsinsInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AsinsInfo: %s", err)
	}

	count, err = AsinsInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
