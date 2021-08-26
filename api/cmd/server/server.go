package server

import (
	"bufio"
	"fmt"
	"os"
	"personal/fintual/api/internal/fintual"
	"time"
)

func Annualized(service fintual.PortfolioService) {
	profit, err := service.AnnualizedProfit()
	if err != nil {
		fmt.Println(err)
		fmt.Println("You can add new random stock, text 's' ")
		return
	}

	fmt.Printf("Annualized Profit is %f\n", profit)
}

func AddStock(service fintual.PortfolioService) {
	stock := fintual.NewStock()
	service.AddStock(stock)
	fmt.Printf("Stock added, price: %f\n", stock.Value)
}

func Profit(service fintual.PortfolioService) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter start date, format yyyy-mm-dd: ")
	scanner.Scan()
	start := scanner.Text()
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		fmt.Println("error: invalid start date")
		return
	}

	fmt.Println("Enter end date, format yyyy-mm-dd: ")
	scanner.Scan()
	end := scanner.Text()
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		fmt.Println("error: invalid end date")
		return
	}

	profit, err := service.Profit(startDate, endDate)
	if err != nil {
		fmt.Println(err)
		fmt.Println("You can add new random stock, text 's' ")
		return
	}

	fmt.Printf("Annualized Profit is %f \n", profit)
}
