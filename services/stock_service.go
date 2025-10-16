package services

type StockExchangeSrv interface {
	GetPrice(stockCode string) (int, error)
}
