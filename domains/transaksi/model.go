package transaction

// Transaction Struct
type Transaction struct {
	ID        string      `json:"id"`
	Created   string      `json:"created"`
	Updated   string      `json:"updated"`
	Total     string      `json:"total"`
	Status    string      `json:"status"`
	SoldItems []SoldItems `json:"solditems"`
}

type SoldItems struct {
	ProductName string `json:"productname"`
	Category    string `json:"category"`
	Harga       string `json:"harga"`
	Quantity    string `json:"quantity"`
	Total       string `json:"total"`
}
