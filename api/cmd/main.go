package main

import (
	"bufio"
	"fmt"
	"os"
	"personal/fintual/api/cmd/server"
	"personal/fintual/api/internal/fintual"
)

func main() {
	portfolio:= fintual.NewPortfolio()
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	fmt.Println("Fintual App ")
	fmt.Println("Menu ")
	fmt.Println("1-  Add Stock to your Portfolio, text: 's'")
	fmt.Println("2-  Get Profit your Portfolio, text: 'p'")
	fmt.Println("3-  Get Annualized Profit your Portfolio, text: 'a'")
	fmt.Println("4-  Quit, text: 'q'")
	for text != "q" {
		fmt.Println("--------------Enter your option: ")
		scanner.Scan()
		text = scanner.Text()
		switch text {
		case "s":
			server.AddStock(portfolio)
		case "p":
			server.Profit(portfolio)
		case "a":
			server.Annualized(portfolio)
		case "q":
			fmt.Println("bye")
		}
	}
}
