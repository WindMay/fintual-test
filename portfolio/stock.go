package portfolio

import "main.go/services"

type Stock interface {
	Code() string
	CurrentPrice() (int, error)
}

type ExchangeStock struct {
	stockCode   string
	exchangeSrv services.StockExchangeSrv
}

func NewStock(stockCode string, srv services.StockExchangeSrv) ExchangeStock {
	return ExchangeStock{
		stockCode:   stockCode,
		exchangeSrv: srv,
	}
}

func (exStock *ExchangeStock) CurrentPrice() (int, error) {
	price, err := exStock.exchangeSrv.GetPrice(exStock.stockCode)
	if err != nil {
		return -1, err
	}
	return price, nil
}

func (exStock *ExchangeStock) GetStockCode() string {
	return exStock.stockCode
}
