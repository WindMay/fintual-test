package main

import (
	"fmt"

	"main.go/portfolio"
	"main.go/services"
)

func init() {
	fmt.Println("FINTUAL Job application")
}

// the main func is the entry point for the program.
func main() {
	// We use a test mock service, we define it as a service since stock prices are not static
	// We use factory service/injection approach since its easier to test
	stockSrv := services.StockMockService{}
	factory := func(code string) portfolio.ExchangeStock {
		return portfolio.NewStock(code, stockSrv)
	}

	// Initial literal for the suggested example
	initialStocks := []portfolio.InitialStock{
		{Code: "APPL", Quantity: 10, Goal: 40},
		{Code: "META", Quantity: 5, Goal: 60},
	}

	// Create a new portfolio with initial allocated exchange stocks
	newPortfolio, _ := portfolio.CreatePortfolio("User's portfolio", initialStocks, factory)
	newPortfolio.Rebalance()
}
