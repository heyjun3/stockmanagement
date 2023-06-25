package rakuten

import (
	"context"

	"github.com/uptrace/bun"

	"crawler/scrape"
)

type RakutenProduct struct {
	bun.BaseModel `bun:"table:rakuten_products"`
	scrape.Product
	point int64
}

func NewRakutenProduct(
	name, productCode, url, jan, shopCode string, price, point int64) (*RakutenProduct, error) {

	p, err := scrape.NewProduct(name, productCode, url, jan, shopCode, price)
	if err != nil {
		return nil, err
	}
	return &RakutenProduct{
		Product: *p,
		point:   point,
	}, nil
}

func (r *RakutenProduct) calcPrice() {
	r.Price = int64(float64(r.Price)*0.91) - r.point
}

func CreateTable(conn *bun.DB, ctx context.Context) error {
	_, err := conn.NewCreateTable().
		Model((*RakutenProduct)(nil)).
		IfNotExists().
		Exec(ctx)
	return err
}
