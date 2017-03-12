package model

type ProductList struct {

	ProductID   string  `gorm:"type:varchar(100);unique_index" json:"product_id"`
	ProductName string  `json:"product_name"`
	Rating      float64 `json:"rating"`
	ImageURL    string `json:"image_url"`
	Marketplace string `json:"marketplace"`
	URL         string `json:"url" gorm:"type:varchar(1000)"`
	Status      string
	Price       PriceType `json:"price"`
	Brand       string    `json:"brand"`
	SellerID    string    `json:"seller_id"`
}

type PriceType struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}