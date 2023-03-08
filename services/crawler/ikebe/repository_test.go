package ikebe

import (
	"context"
	"crawler/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/null/v8"

	"crawler/config"
)

func IkebeProductTableFactory(conn boil.ContextExecutor) error {
	_, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS ikebe_product (
        name VARCHAR, 
        jan VARCHAR, 
        price BIGINT, 
        shop_code VARCHAR NOT NULL, 
        product_code VARCHAR NOT NULL, 
        url VARCHAR, 
        PRIMARY KEY (shop_code, product_code));`)
	if err != nil {
		return err
	}
	return err
}

func TestGetIkebeProductsByProductCode(t *testing.T) {
	ctx := context.Background()
	conf, _ := config.NewConfig("../sqlboiler.toml")
	conf.Psql.DBname = "test"
	conn, err := NewDBconnection(conf.Dsn())
	if err != nil {
		fmt.Println(err)
	}
	err = IkebeProductTableFactory(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	models.IkebeProducts().DeleteAll(ctx, conn)
	p := NewIkebeProduct("test", "test_code", "https://test.com", "", 1111)
	if err := p.Insert(ctx, conn, boil.Infer()); err != nil {
		logger.Error("insert error", err)
	}

	t.Run("get products", func(t *testing.T) {
		r := IkebeProductRepository{}

		products, err := r.getByProductCodes(ctx, conn, "test_code")

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(products))
		assert.Equal(t, p, products[0])
	})
}

func TestBulkUpsertIkebeProducts(t *testing.T) {
	conf, _ := config.NewConfig("../sqlboiler.toml")
	conf.Psql.DBname = "test"
	conn, err := NewDBconnection(conf.Dsn())
	if err != nil {
		fmt.Println(err)
		return
	}
	err = IkebeProductTableFactory(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	models.IkebeProducts().DeleteAll(ctx, conn)

	t.Run("upsert ikebe products", func(t *testing.T) {
		p := IkebeProducts{
			NewIkebeProduct("test", "test", "https://test.jp", "1111", 9000),
			NewIkebeProduct("test", "test1", "https://test.jp", "1111", 9000),
			NewIkebeProduct("test", "test2", "https://test.jp", "1111", 9000),
		}

		err = p.bulkUpsert(conn)

		assert.Equal(t, nil, err)
		i, _ := models.FindIkebeProduct(ctx, conn, "ikebe", "test1")
		assert.Equal(t, *p[1], IkebeProduct{*i})
	})
}

func TestMappingIkebeProducts(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		p := Products{
			NewIkebeProduct("test", "test", "http://test.jp", "", 1111),
			NewIkebeProduct("test1", "test1", "http://test.jp", "", 1111),
			NewIkebeProduct("test2", "test2", "http://test.jp", "", 1111),
		}

		dbp := Products{
			NewIkebeProduct("test", "test", "test", "4444", 4000),
			NewIkebeProduct("test", "test1", "test1", "555", 4000),
			NewIkebeProduct("test", "test2", "test2", "7777", 4000),
		}

		result := p.mapProducts(dbp)

		assert.Equal(t, 3, len(result))
		assert.Equal(t, NewIkebeProduct("test", "test", "http://test.jp", "4444", 1111), result[0])
		assert.Equal(t, NewIkebeProduct("test1", "test1", "http://test.jp", "555", 1111), result[1])
		assert.Equal(t, NewIkebeProduct("test2", "test2", "http://test.jp", "7777", 1111), result[2])
	})

	t.Run("product is empty", func(t *testing.T) {
		p := Products{}
		dbp := Products{
			NewIkebeProduct("test", "test", "test", "11111", 4000),
			NewIkebeProduct("test", "test", "test1", "55555", 4000),
		}

		result := p.mapProducts(dbp)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, p, result)
	})

	t.Run("db product is empty", func(t *testing.T) {
		p := Products{
			NewIkebeProduct("test", "test", "http://test.jp", "", 1111),
			NewIkebeProduct("test1", "test1", "http://test.jp", "", 1111),
			NewIkebeProduct("test2", "test2", "http://test.jp", "", 1111),
		}
		db := Products{}

		result := p.mapProducts(db)

		assert.Equal(t, 3, len(result))
		assert.Equal(t, p, result)
	})
}

func TestGenerateMessage(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		p := NewIkebeProduct("test", "test", "https://test.com", "", 6000)
		p.Jan = null.StringFrom("4444")
		f := "ikebe_20220301_120303"

		m, err := p.generateMessage(f)

		assert.Equal(t, nil, err)
		ex := `{"filename":"ikebe_20220301_120303","jan":"4444","cost":6000,"url":"https://test.com"}`
		assert.Equal(t, ex, string(m))
	})

	t.Run("Jan code isn't Valid", func(t *testing.T) {
		p := NewIkebeProduct("TEST", "test", "https://test.com", "", 5000)
		f := "ikebe_20220202_020222"

		m, err := p.generateMessage(f)

		assert.Error(t, err)
		assert.Equal(t, []byte(nil), m)
	})

	t.Run("Price isn't valid", func(t *testing.T) {
		p := NewIkebeProduct("TEST", "test", "https://test.com", "", 5000)
		p.Price = null.Int64FromPtr(nil)
		f := "ikebe_20220202_020222"

		m, err := p.generateMessage(f)

		assert.Error(t, err)
		assert.Equal(t, []byte(nil), m)
	})

	t.Run("URL isn't valid", func(t *testing.T) {
		p := NewIkebeProduct("TEST", "test", "https://test.com", "", 5000)
		p.URL = null.StringFromPtr(nil)
		f := "ikebe_20220202_020222"

		m, err := p.generateMessage(f)

		assert.Error(t, err)
		assert.Equal(t, []byte(nil), m)
	})
}