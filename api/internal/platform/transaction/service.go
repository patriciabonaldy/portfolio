package transaction

import (
	"personal/fintual/api/internal/fintual"
	"personal/fintual/api/internal/platform/storage"
)

// service struct.
type service struct {
	repository      storage.Repository
}

type Service interface {
	CreateStock(stock fintual.Stock)
	CreatePortfolio(portfolio fintual.Portfolio)
	AddStock(portfolioId int, stock fintual.Stock)
}