package inventory

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"golang.org/x/exp/slog"

	"api-server/database"
	"api-server/spapi"
	"api-server/spapi/price"
)

var SpapiServiceURL string
var db *bun.DB
var inventoryRepository InventoryRepository
var priceRepository PriceRepository[*CurrentPrice]
var lowestPriceRepository PriceRepository[*LowestPrice]
var desiredPriceRepository PriceRepository[*DesiredPrice]
var spapiClient *spapi.SpapiClient

func init() {
	SpapiServiceURL = os.Getenv("SPAPI_SERVICE_URL")
	if SpapiServiceURL == "" {
		panic(errors.New("don't set SPAPI_SERVICE_URL"))
	}
	var err error
	spapiClient, err = spapi.NewSpapiClient(SpapiServiceURL)
	if err != nil {
		panic(err)
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic(errors.New("don't set DB_DSN"))
	}
	db = database.OpenDB(dsn)
	inventoryRepository = InventoryRepository{}
	priceRepository = PriceRepository[*CurrentPrice]{}
	lowestPriceRepository = PriceRepository[*LowestPrice]{}
}

func RefreshInventory(c echo.Context) error {
	var nextToken string
	for {
		res, err := spapiClient.InventorySummaries(nextToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		inventories := Inventories{}
		for _, inventory := range res.Payload.InventorySummaries {
			iv, err := NewInventoryFromInventory(inventory)
			if err != nil {
				slog.Warn(err.Error(), "struct", *inventory, "sku", inventory.SellerSku)
				continue
			}
			inventories = append(inventories, iv)
		}
		if err := inventoryRepository.Save(context.Background(), db, inventories); err != nil {
			slog.Error("failed save inventories", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		slog.Info("logging next token", "nextToken", res.Pagination.NextToken)
		nextToken = res.Pagination.NextToken
		if nextToken == "" {
			slog.Info("break loop")
			break
		}
	}
	return c.JSON(http.StatusOK, "success")
}

func RefreshPricing(c echo.Context) error {
	var inventories Inventories
	var cursor Cursor
	var err error
	for {
		inventories, cursor, err = inventoryRepository.GetNextPage(context.Background(), db, cursor.End, 20)
		if err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		if len(inventories) == 0 || inventories == nil {
			return c.JSON(http.StatusOK, "success")
		}

		res, err := spapiClient.GetPricing(inventories.Skus(), price.Sku)
		if err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		prices := make(CurrentPrices, 0, len(inventories))
		for _, payload := range res.Payload {
			offers := payload.Product.Offers
			if len(offers) == 0 {
				continue
			}
			sku := payload.SellerSKU
			amount := offers[0].BuyingPrice.ListingPrice.Amount
			points := offers[0].BuyingPrice.Points.PointsNumber
			if amount == nil || points == nil {
				slog.Warn("amount or points is nil. must be not nil", "sku", sku)
				continue
			}
			price, err := NewCurrentPrice(sku, Ptr(int(*amount)), Ptr(int(*points)))
			if err != nil {
				slog.Error(err.Error(), "sku", sku)
				continue
			}
			prices = append(prices, price)
		}
		if err := priceRepository.Save(context.Background(), db, prices); err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
}

func RefreshLowestPricing(c echo.Context) error {
	var inventories Inventories
	var cursor Cursor
	var err error
	for {
		inventories, cursor, err = inventoryRepository.GetNextPage(context.Background(), db, cursor.End, 20)
		if err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		if len(inventories) == 0 || inventories == nil {
			return c.JSON(http.StatusOK, "success")
		}

		res, err := spapiClient.GetLowestPricing(inventories.Skus())
		if err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		prices := make(LowestPrices, 0, len(inventories))
		for _, response := range res.Responses {
			offers := response.Body.Payload.Offers
			if len(offers) == 0 {
				continue
			}
			sku := response.Body.Payload.SKU
			amount := offers[0].Price.Amount
			point := offers[0].Points.PointsNumber
			if amount == nil || point == nil {
				slog.Warn("amount or point is nil. must be not nil", "sku", sku, "amount", amount, "point", point)
				continue
			}
			price, err := NewLowestPrice(sku, Ptr(int(*amount)), Ptr(int(*point)))
			if err != nil {
				slog.Error(err.Error(), "sku", sku)
				continue
			}
			prices = append(prices, price)
		}
		if err := lowestPriceRepository.Save(context.Background(), db, prices); err != nil {
			slog.Error("error", "detail", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
}

func GetInventories(c echo.Context) error {
	ctx := context.Background()
	condition := Condition{Quantity: Ptr(0), IsNotOnlyLowestPrice: true}
	inventories, err := inventoryRepository.GetByCondition(ctx, db, condition)
	if err != nil {
		slog.Error("get inventories error", "detail", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, inventories)
}

type UpdatePricingDTO struct {
	Price        int `json:"price"`
	PercnetPoint int `json:"percentPoint"`
}

func UpdatePricing(c echo.Context) error {
	sku := c.Param("sku")
	if sku == "" {
		return c.JSON(http.StatusBadRequest, errors.New("sku is not empty string"))
	}
	dto := new(UpdatePricingDTO)
	if err := c.Bind(dto); err != nil {
		slog.Error("un bind body", err)
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	price, err := NewDesiredPrice(&sku, &dto.Price, &dto.PercnetPoint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if err := desiredPriceRepository.Save(context.Background(), db, []*DesiredPrice{price}); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, dto)
}
