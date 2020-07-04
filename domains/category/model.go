package category

// Category Struct
type Category struct {
	ID           string   `json:"id"`
	CategoryName string   `json:"categoryname"`
	Produk       []string `json:"produk"`
	Created      string   `json:"created"`
	Updated      string   `json:"updated"`
	Status       string   `json:"status"`
}
