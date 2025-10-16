package portfolio

import (
	"fmt"
)

type ExStockFactory func(code string) ExchangeStock
type AllocatedExStock struct {
	Stock    ExchangeStock
	Quantity int
	Goal     int
}
type InitialStock struct {
	Code     string
	Quantity int
	Goal     int
}

type Portfolio struct {
	Id              string
	NewStock        ExStockFactory
	AllocatedStocks []AllocatedExStock
}

func CreatePortfolio(id string, initialStocksCodes []InitialStock, stockFactory ExStockFactory) (*Portfolio, error) {
	// Pricing loading may be graceful since it could be a third party service, that why a custom factory for stocks
	var initialStocks []AllocatedExStock
	for _, initStock := range initialStocksCodes {
		newStock := stockFactory(initStock.Code)
		newAllocatedStock := AllocatedExStock{
			Quantity: initStock.Quantity,
			Goal:     initStock.Goal,
			Stock:    newStock,
		}
		initialStocks = append(initialStocks, newAllocatedStock)
	}
	return &Portfolio{
		Id:              id,
		AllocatedStocks: initialStocks,
	}, nil
}

func (p *Portfolio) Rebalance() {
	var totalSharesValue int
	var valuesStocks []int
	// Fetch a sum the prices to get the total portfolio value
	for _, allocated := range p.AllocatedStocks {
		currentPrice, _ := allocated.Stock.CurrentPrice()
		value := currentPrice * allocated.Quantity
		totalSharesValue += value
		valuesStocks = append(valuesStocks, value)
	}

	// We check what to do for each allocation
	for i, price := range valuesStocks {
		currentPortion := (price * 100) / totalSharesValue
		fmt.Println("Allocated Stock ", p.AllocatedStocks[i].Stock.GetStockCode())
		fmt.Println("Goal: ", p.AllocatedStocks[i].Goal, ",  Current portion: ", currentPortion)
		if currentPortion > p.AllocatedStocks[i].Goal {
			fmt.Println("We must sell these stocks to reach the goal")
		} else if currentPortion < p.AllocatedStocks[i].Goal {
			fmt.Println("We must buy these stocks to reach the goal")
		} else {
			fmt.Println("This stock its ok")
		}
	}
}
