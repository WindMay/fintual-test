package services

// StockExchangeSrv defines an interface for retrieving stock prices based on a given stock code.
type StockExchangeSrv interface {
	GetPrice(stockCode string) (int, error)
}
