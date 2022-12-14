package entity

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Result struct {
	Message string    `json:"message"`
	Result  []Product `json:"result"`
}
