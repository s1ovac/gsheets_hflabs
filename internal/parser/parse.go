package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/s1ovac/gdoc/internal/config"
	"log"
	"strconv"
	"strings"
)

type ApiCodes struct {
	code        int
	description string
}

func (a ApiCodes) Code() int {
	return a.code
}

func (a ApiCodes) Description() string {
	return a.description
}

type ParseHTML struct {
	cfg config.SheetConfig
}

func NewParseHTML(cfg config.SheetConfig) ParseHTML {
	return ParseHTML{
		cfg: cfg,
	}
}

func (h ParseHTML) ParseTable(ctx context.Context) ([]ApiCodes, error) {
	data, err := parse(h.cfg.URL())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parse(url string) ([]ApiCodes, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	var data []ApiCodes
	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		cols := row.Find("td")
		if cols.Length() == 2 {
			number, err := strconv.Atoi(strings.TrimSpace(cols.Eq(0).Text()))
			if err != nil {
				log.Printf("Error parsing number in row %d: %s", i, err)
			} else {
				value := strings.TrimSpace(cols.Eq(1).Text())
				data = append(data, ApiCodes{code: number, description: value})
			}
		}
	})

	return data, nil
}
