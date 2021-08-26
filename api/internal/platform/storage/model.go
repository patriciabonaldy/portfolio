package storage

type stockRepo struct {
	ID    int    `json:"id"`
	Price string `json:"price"`
	DateCreated  string `json:"date_created"`
}

type stockHistoryRepo struct {
	ID      int    `json:"id"`
	StockID int    `json:"stock_id"`
	Price   string `json:"price"`
	DateHistory  string `json:"date_history"`
}

type potfolioRepo struct {
	ID          int    `json:"id"`
	stockID int    `json:"stock_id"`
	DateCreated  string `json:"date_created"`
}
