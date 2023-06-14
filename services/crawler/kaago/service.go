package kaago

import (
	"log"
	"net/http"
	"strings"

	"crawler/scrape"
)

func NewScrapeService(url, payload string) scrape.Service[*KaagoProduct] {
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Fatalln(err)
	}
	service := scrape.NewService(KaagoParser{}, &KaagoProduct{}, []*KaagoProduct{})
	service.EntryReq = req
	return service
}
