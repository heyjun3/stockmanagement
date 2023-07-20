package murauchi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"crawler/scrape"
)

const (
	host   = "www.murauchi.com"
	scheme = "https"
	path   = "MCJ-front-web/WH/front/Default.do?type=COMMODITY_LIST"
)

type MurauchiParser struct{}

func (p MurauchiParser) ProductListByReq(r io.ReadCloser, req *http.Request) (scrape.Products, *http.Request) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		logger.Error("response parse error", err)
		return nil, nil
	}

	var products scrape.Products
	doc.Find(".window_item").Each(func(i int, s *goquery.Selection) {
		nameElem := s.Find("h2 a")
		name := nameElem.Text()
		href, exist := nameElem.Attr("href")
		URL, err := url.Parse(href)
		if !exist || err != nil {
			logger.Info("Not Found url")
			return
		}
		var paths []string
		for _, path := range strings.Split(URL.Path, "/") {
			if path != "" {
				paths = append(paths, path)
			}
		}
		productId := paths[2]
		URL = &url.URL{Host: host, Scheme: scheme, Path: strings.Join(paths[:3], "/")}
		priceText := s.Find(".price_p .price").Text()
		price, err := scrape.PullOutNumber(priceText)
		if err != nil {
			logger.Info("Not Found price")
			return
		}

		sold := s.Find(".stock span").Text()
		if sold := strings.TrimSpace(sold); sold == "販売終了" || sold == "予約中" {
			logger.Info("product is sold out")
			return
		}
		product, err := NewMurauchiProduct(name, productId, URL.String(), "", price)
		if err != nil {
			logger.Info("error", err)
			return
		}
		products = append(products, product)
	})

	nextRequest, err := generateRequestFromPreviousRequest(req)
	if err != nil {
		logger.Error("error", err)
		return nil, nil
	}

	return products, nextRequest
}

func generateRequestFromPreviousRequest(pre *http.Request) (*http.Request, error) {
	category := pre.Header.Get("x-category")
	page := pre.Header.Get("x-current")
	if category == "" || page == "" {
		return nil, fmt.Errorf("category not found error")
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	return generateRequest(p+1, category)
}

func generateRequest(page int, category string) (*http.Request, error) {
	values := map[string]string{
		"mode":         "graphic",
		"pageNumber":   fmt.Sprint(page),
		"searchType":   "keyword",
		"sortOrder":    "",
		"categoryNo":   category,
		"type":         "COMMODITY_LIST",
		"keyword":      "　",
		"listCount":    "120",
		"handlingType": "0",
	}
	form := url.Values{}
	for k, v := range values {
		form.Add(k, v)
	}
	body := strings.NewReader(form.Encode())

	u := url.URL{Scheme: scheme, Host: host, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-category", category)
	req.Header.Set("x-current", fmt.Sprint(page))

	return req, nil
}
