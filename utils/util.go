package utils

import (
	"net/http"
	"io"
	"fmt"
	"io/ioutil"
	"time"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"github.com/rakeshkumargupt/scrape-go/model"
	"strings"
	"strconv"
)

func CallHttp(url, method string, body io.Reader, m map[string]string) (*http.Response, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for i, v := range m{
		req.Header.Add(i, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetDocument(resp *http.Response) (*goquery.Document, error) {

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {

		if err == io.EOF {

			fmt.Println("Error While Requesting : EOF")
			fmt.Println(err)

			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))

			return nil, err

		}

		fmt.Println(time.Now().String() + ":Error in Converting Request to Doc :" + err.Error())
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil, err

	}

	return doc, nil

}


func CreateTrendingProductListItem(s *goquery.Selection) model.ProductList {
	var listItem model.ProductList
	var id string
	selectHref := s.Find("a[class*='a-link-normal']").First()
	href, ex := selectHref.Attr("href")
	if !ex {
		fmt.Println("List Item : attribute doesnt exist")
		return listItem
	}
	if href != "" {
		if !strings.Contains(href, "picassoRedirect.html") {
			tokenized := strings.Split(href, "/")
			id = tokenized[3]
		} else {
			pURL, _ := url.Parse(href)
			lUrl := pURL.Query().Get("url")
			if lUrl != "" {
				tokenized := strings.Split(lUrl, "/")
				id = tokenized[3]
			}
		}

		name, _ := s.Find("img").Attr("alt")
		name = removeExtraCharacters(name)
		imageURL, exists := s.Find("img").First().Attr("src")
		if exists {
			imageURL = removeExtraCharacters(imageURL)
		}
		//getting price
		priceSelecotor := s.Find("span[class*='a-size-base'] span").First()
		if priceSelecotor != nil {
			wholePrice := strings.TrimSpace(priceSelecotor.Text())
			if wholePrice != "" {
				mrp, err := strconv.ParseFloat(strings.Replace(wholePrice, "$", "", -1), 64)
				if err != nil {
					fmt.Println("Issue While converting mrp string to float : getMRP : ", id)
					listItem.Price = model.PriceType{Amount: -1, Currency: "Dollar"}
				} else {
					listItem.Price = model.PriceType{Amount: mrp, Currency: "Dollar"}
				}
			}

		}

		//getting rating
		ratingText := s.Find("i[class*='a-icon a-icon-star'] span").First().Text()
		if ratingText != "" {
			ratingText = strings.TrimSpace(strings.Replace(strings.Replace(strings.Replace(strings.ToLower(ratingText), "out", "", -1), "of", "", -1), "stars", "", -1))
			s := strings.Split(ratingText, " ")
			val, err1 := strconv.ParseFloat(s[0], 64)
			outOff, err2 := strconv.ParseFloat(s[len(s) - 1], 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Issue While converting rating string to float : getRatings : ", id)

			} else {
				listItem.Rating = (val / outOff) * 100
			}
		}
		listItem.ProductID = id
		listItem.ProductName = name
		listItem.URL = href
		listItem.Marketplace = "Amazon"
		listItem.ImageURL = imageURL
	}
	return listItem
}

func removeExtraCharacters(str string) string {

	var removeStringList []string

	removeStringList = append(removeStringList, "\n")
	removeStringList = append(removeStringList, "\t")

	for _, val := range removeStringList {
		str = strings.Replace(str, val, "", -1)
	}

	str = strings.TrimSpace(str)

	return str

}
