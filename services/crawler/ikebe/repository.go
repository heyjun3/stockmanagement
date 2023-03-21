package ikebe

import (
	"context"

	"github.com/uptrace/bun"

	"crawler/scrape"
)

type IkebeProductRepository struct{
	scrape.ProductRepository
}

func (r IkebeProductRepository) GetByProductCodes(conn *bun.DB,
	ctx context.Context, codes ...string) (scrape.Products, error) {

	var ikebeProducts []IkebeProduct
	err := conn.NewSelect().
		Model(&ikebeProducts).
		Where("product_code IN (?)", bun.In(codes)).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var products scrape.Products
	for i := 0; i < len(ikebeProducts); i++ {
		products = append(products, &ikebeProducts[i])
	}
	return products, nil
}
