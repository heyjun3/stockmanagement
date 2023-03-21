package ikebe

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"

	"crawler/config"
	"crawler/scrape"
)

func IkebeDatabaseFactory() (*bun.DB, context.Context, error) {
	ctx := context.Background()
	conf, _ := config.NewConfig("../sqlboiler.toml")
	conf.Psql.DBname = "test"
	conn := scrape.CreateDBConnection(conf.Dsn())
	_, err := conn.NewCreateTable().Model((*IkebeProduct)(nil)).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	conn.NewDelete().Model((*IkebeProduct)(nil)).Exec(ctx)
	return conn, ctx, nil
}

func TestGetIkebeProductsByProductCode(t *testing.T) {
	conn, ctx, err := IkebeDatabaseFactory()
	if err != nil {
		return
	}
	p := NewIkebeProduct("test", "test_code", "https://test.com", "", 1111)
	repo := IkebeProductRepository{}
	if err := repo.Upsert(conn, ctx, p); err != nil {
		logger.Error("insert error", err)
	}

	t.Run("get products", func(t *testing.T) {

		products, err := repo.GetByProductCodes(conn, ctx, "test_code")

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(products))
		assert.Equal(t, p, products[0])
	})
}

func TestUpsert(t *testing.T) {
	conn, ctx, err := IkebeDatabaseFactory()
	if err != nil {
		return
	}
	repo := IkebeProductRepository{}

	t.Run("upsert ikebe product", func(t *testing.T) {
		p := NewIkebeProduct("test", "test", "test url", "1111", 9000)

		err := repo.Upsert(conn, ctx, p)

		assert.Equal(t, nil, err)
		expectd, _ := repo.GetByProductCodes(conn, ctx, "test")
		assert.Equal(t, (expectd[0]).(*IkebeProduct), p)
	})
}
