package main

import (
	"fmt"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	getCurrencyData("BTC")
	getCurrencyData("ETH")
	getCurrencyData("BCH")
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
