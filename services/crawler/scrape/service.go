package scrape

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/uptrace/bun"

	"crawler/config"
)

var logger = config.Logger

type Service struct{
	Parser Parser
	FetchProductByProductCodes func(*bun.DB, context.Context, ...string)(Products, error)
}

type Parser interface {
	ProductList(io.ReadCloser) (Products, string)
	Product(io.ReadCloser) (string, error)
}

func (s Service) StartScrape(url, shopName string) {
	client := NewClient()
	mqClient := NewMQClient(config.MQDsn, "mws")
	wg := sync.WaitGroup{}
	wg.Add(1)

	c1 := s.ScrapeProductsList(client, url)
	c2 := s.GetProducts(c1, config.DBDsn)
	c3 := s.ScrapeProduct(c2, client)
	c4 := s.SaveProduct(c3, config.DBDsn)
	s.SendMessage(c4, mqClient, shopName, &wg)

	wg.Wait()
}

func (s Service) ScrapeProductsList(
	client httpClient, url string) chan Products {
	c := make(chan Products, 100)
	go func() {
		defer close(c)
		for url != "" {
			logger.Info("product list request url", "url", url)
			res, err := client.Request("GET", url, nil)
			if err != nil {
				logger.Error("http request error", err)
				break
			}
			var products Products
			products, url = s.Parser.ProductList(res.Body)
			res.Body.Close()

			c <- products
		}
	}()
	return c
}

func (s Service) GetProducts(c chan Products, dsn string) chan Products {
	send := make(chan Products, 100)
	go func() {
		defer close(send)
		ctx := context.Background()
		conn := CreateDBConnection(dsn)

		for p := range c {
			dbProduct, err := s.FetchProductByProductCodes(conn, ctx, p.getProductCodes()...)
			if err != nil {
				logger.Error("db get product error", err)
				continue
			}
			products := p.MapProducts(dbProduct)
			send <- products
		}
	}()
	return send
}

func (s Service) ScrapeProduct(
	ch chan Products, client httpClient) chan Product {

	send := make(chan Product, 100)
	go func() {
		defer close(send)
		for products := range ch {
			for _, product := range products {
				if product.IsValidJan() {
					send <- product
					continue
				}

				logger.Info("product request url", "url", product.GetURL())
				res, err := client.Request("GET", product.GetURL(), nil)
				if err != nil {
					logger.Error("http request error", err, "action", "scrapeProduct")
					continue
				}
				jan, _ := s.Parser.Product(res.Body)
				res.Body.Close()
				product.SetJan(jan)
				send <- product
			}
		}
	}()
	return send
}

func (s Service) SaveProduct(ch chan Product, dsn string) chan Product {

	send := make(chan Product, 100)
	go func() {
		defer close(send)
		ctx := context.Background()
		conn := CreateDBConnection(dsn)
		for p := range ch {
			err := p.Upsert(conn, ctx)
			if err != nil {
				logger.Error("product upsert error", err)
				continue
			}
			send <- p
		}
	}()
	return send
}

func (s Service) SendMessage(
	ch chan Product, client RabbitMQClient,
	shop_name string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		filename := shop_name + "_" + timeToStr(time.Now())
		for p := range ch {
			m, err := p.GenerateMessage(filename)
			if err != nil {
				logger.Error("generate message error", err)
				continue
			}

			err = client.Publish(m)
			if err != nil {
				logger.Error("message publish error", err)
			}
		}
	}()
}
