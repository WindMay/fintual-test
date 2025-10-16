package main

import (
	"encoding/json"
	"fmt"

	"main.go/portfolio"
	"main.go/services"
)

/** Fintual test */

//Construct a simple Portfolio class that has a collection of Stocks. Assume each Stock has a “Current Price” method
//that receives the last available price. Also, the Portfolio class has a collection of “allocated” Stocks that
//represents the distribution of the Stocks the Portfolio is aiming (i.e. 40% META, 60% APPL)
//Provide a portfolio rebalance method to know which Stocks should be sold and which ones should be bought to have a
//balanced Portfolio based on the portfolio’s allocation.
//Add documentation/comments to understand your thinking process and solution

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	fmt.Println("FINTUAL Job application")
}

// the main func is the entry point for the program.
func main() {
	//
	stockSrv := services.StockMockService{}
	factory := func(code string) portfolio.ExchangeStock {
		return portfolio.NewStock(code, stockSrv)
	}

	initialStocks := []portfolio.InitialStock{
		{Code: "APPL", Quantity: 10, Goal: 40},
		{Code: "META", Quantity: 5, Goal: 60},
	}

	newPortfolio, _ := portfolio.CreatePortfolio("User's portfolio", initialStocks, factory)
	newPortfolio.Rebalance()

}
