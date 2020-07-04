package product

// Product Struct
type Product struct {
	ID           string `json:"id"`
	ProductName  string `json:"productname"`
	CatID        string `json:"catid"`
	CategoryName string `json:"categoryname"`
	Harga        string `json:"harga"`
	Created      string `json:"created"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
}
