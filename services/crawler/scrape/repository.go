package scrape

import (
	"context"
	"reflect"

	"github.com/uptrace/bun"
)

type ProductRepository[T IProduct] struct {
	product  T
	products []T
}

func NewProductRepository[T IProduct](p T, ps []T) ProductRepository[T] {
	return ProductRepository[T]{
		product:  p,
		products: ps,
	}
}

func (p ProductRepository[T]) GetProduct(ctx context.Context,
	db *bun.DB, productCode, shopCode string) (T, error) {

	product := reflect.New(reflect.ValueOf(p.product).Elem().Type()).Interface().(T)
	err := db.NewSelect().
		Model(product).
		Where("product_code = ?", productCode).
		Where("shop_code = ?", shopCode).
		Scan(ctx, product)

	return product, err
}

func (p ProductRepository[T]) GetByProductCodes(ctx context.Context, db *bun.DB, codes ...string) (Products, error) {
	products := p.products

	err := db.NewSelect().
		Model(&products).
		Where("product_code IN (?)", bun.In(codes)).
		Order("product_code ASC").
		Scan(ctx, &products)

	var result Products
	for _, p := range products {
		result = append(result, p)
	}

	return result, err
}

func BulkUpsert[T IProduct](ctx context.Context, db *bun.DB, ps []T) error {
	mapProduct := map[string]IProduct{}
	for _, v := range ps {
		mapProduct[v.GetProductCode()] = v
	}
	var products Products
	for _, v := range mapProduct {
		products = append(products, v)
	}

	_, err := db.NewInsert().
		Model(&products).
		On("CONFLICT (shop_code, product_code) DO UPDATE").
		Set(`
			name = EXCLUDED.name,
			jan = EXCLUDED.jan,
			price = EXCLUDED.price,
			url = EXCLUDED.url
		`).
		Returning("NULL").
		Exec(ctx)

	return err
}
