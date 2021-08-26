package storage

// Repository interface.
type Repository interface {
	GetStock(stockID int) stockRepo
	GetPortfolio() []potfolioRepo
	UpdateStock(stockID int, price string, date string)
	AddStock(price string, date string)
	AddStockHistory(stockID int, price string, date string)
	AddPortfolio(stockID int, date string)
}

// NewRepository constructor.
func NewRepository() Repository {
	return &Storage{}
}