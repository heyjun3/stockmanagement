package handler

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"api-server/product"
)

type Res struct {
	Message string `json:"status"`
}

type FileListRes struct {
	List []string `json:"list"`
}

type StatusRes struct {
	Keepa map[string]int `json:"keepa"`
	Mws   map[string]int `json:"mws"`
}

type ProductWithChart struct {
	product.Product
	product.ChartData
}

type ChartRes struct {
	Charts      []ProductWithChart `json:"chart_data"`
	CurrentPage int                `json:"current_page"`
	MaxPage     int                `json:"max_page"`
}

var db *bun.DB
var productRepo product.ProductRepository
var keepaRepo product.KeepaRepository

func init() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("db dsn is null")
	}
	db = product.OpenDB(dsn)
	productRepo = product.ProductRepository{DB: db}
	keepaRepo = product.KeepaRepository{DB: db}
}

func GetCharts(c echo.Context) error {
	filename := c.Param("filename")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if filename == "" || err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, Res{"error"})
	}
	limit := 100
	ctx := context.Background()
	charts, total, err := productRepo.GetProductWithChart(ctx, filename, page, limit)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Res{"error"})
	}

	maxPage := int(math.Ceil((float64(total) / float64(limit))))
	if page > maxPage {
		fmt.Println("page over max pages")
		return c.JSON(http.StatusNotFound, Res{"error"})
	}

	products := make([]ProductWithChart, len(charts))
	for i := 0; i < len(charts); i++ {
		products[i] = ProductWithChart{charts[i].Product, charts[i].Chart}
	}
	return c.JSON(http.StatusOK, ChartRes{products, page, maxPage})
}

func GetFilenames(c echo.Context) error {
	ctx := context.Background()
	filenames, err := productRepo.GetFilenames(ctx)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Res{"error"})
	}
	return c.JSON(http.StatusOK, FileListRes{filenames})
}

func GetStatusCounts(c echo.Context) error {
	ctx := context.Background()
	keepa, err := keepaRepo.GetCounts(ctx)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Res{"error"})
	}
	mws, err := productRepo.GetCounts(ctx)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Res{"error"})
	}
	return c.JSON(http.StatusOK, StatusRes{keepa, mws})
}

func DeleteProducts(c echo.Context) error {
	ctx := context.Background()
	err := productRepo.DeleteByFilename(ctx, c.Param("filename"))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Res{"error"})
	}
	return c.JSON(http.StatusOK, Res{"ok"})
}
