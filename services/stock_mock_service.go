package services

import "errors"

type StockMockService struct{}

func (StockMockService) GetPrice(stockCode string) (int, error) {

	switch stockCode {
	case "APPL":
		return 12, nil
	case "META":
		return 20, nil
	default:
		return -1, errors.New("unknow stock code, this is mock service btw")
	}
}
