# scrape-go
A simple and fast web scrapping in go.                                                                                              
This package is a place to put some simple tools which build on top of the Go HTML parsing library and GoQuery.                     
Scrape defines traversal functions like Callhttp with input request and GetProductDetails, while attempting to be generic.         
It also uses convenience functions such as find and selection and more from GoQuery.

# Sample

        package main

        import (

        "encoding/json"
        "fmt"
        "github.com/PuerkitoBio/goquery"
        "github.com/rakeshkumargupt/scrape-go/utils"
        "io"
        "os"
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

        // URL to call for scrapping
        // Example : choose here different category for getting data about product
        url := "http://www.amazon.com/gp/bestsellers/beauty#1"

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
        }
