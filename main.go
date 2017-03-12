package main

import (
	"fmt"
	"github.com/rakeshkumargupt/scrape-go/utils"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"
	"io"
)

func main() {
	fmt.Println("Starting main..")

	//  Set header in request with key & value[OPTIONAL]
	 m := map[string]string{}

	// Method to call http
	method := "GET"

	// OPTIONAL body if method is POST or ...
	var body io.Reader
	_ = body

	// URL to call for scrapping[choose here different category for getting data about product]
	url  := "http://www.amazon.com/gp/bestsellers/beauty#1"

	// Getting http response
	resp, err := utils.CallHttp(url, method, nil, m)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Getting goQuery  document for scrapping
	doc, err := utils.GetDocument(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Getting product information of scapped webpage
	if doc != nil {
		doc.Find("div[class='zg_itemImmersion']").Each(func(i int, s *goquery.Selection) {
			prod := utils.CreateTrendingProductListItem(s)
			if prod.ProductID != "" {

				b, err := json.MarshalIndent(prod, "", "  ")
				if err != nil {
					panic(err)
				}
				b2 := append(b, '\n')
				os.Stdout.Write(b2)
			}
		})

	}

}