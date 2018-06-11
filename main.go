package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/piquette/finance-go/equity"
)

func main() {
	if len(os.Args) < 2 {
		exit("please provide a ticker to fetch")
	}

	ticker := flag.String("ticker", "", "ticker to fetch quote for")
	flag.Parse()

	quote, err := equity.Get(*ticker)
	exitOn(err)

	fmt.Println(quote.RegularMarketPrice)
}

func exit(reason string) {
	exitOn(errors.New(reason))
}

func exitOn(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
