package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/piquette/finance-go/equity"
)

type fetchArgs struct {
	Ticker string `json:"ticker,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		exit("please provide an action")
	}

	switch os.Args[1] {
	case "lifecycle":
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
	case "fetch":
		if len(os.Args) < 3 {
			exit("please provide args to fetch")
		}

		fetchArgs := &fetchArgs{}
		if err := json.Unmarshal([]byte(os.Args[2]), &fetchArgs); err != nil {
			exit("please provide valid args json")
		}

		if fetchArgs.Ticker == "" {
			exit("please provide a ticker in the args json")
		}

		quote, err := equity.Get(fetchArgs.Ticker)
		exitOn(err)

		fmt.Println(quote.RegularMarketPrice)
	default:
		exit("please provide 'fetch' as the action")
	}

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
