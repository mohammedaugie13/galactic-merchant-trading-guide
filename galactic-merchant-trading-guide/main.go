package main

import (
	"bufio"
	"fmt"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/currency"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/eventprocessor"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/parser"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/roman"
	"github.com/mohammedaugie13/galactic-merchant-trading-guide/translator"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	p := parser.NewParser()
	c := currency.NewCurrencies()
	r := roman.NewRoman()
	tr := translator.NewTranslator(r)
	ep := eventprocessor.NewEventProcessor(tr, p, c)

	reader := bufio.NewReader(os.Stdin)
	isContinue := true

	closeHandler(func() {
		isContinue = false
	})

	fmt.Println(":: Intergalactic numerals conversion ::")

	for isContinue {
		stmt := readLine(reader)
		res, err := ep.ProcessStatement(stmt)
		if err != nil {
			fmt.Println(err.Error())
		}

		if res != "" {
			fmt.Println(res)
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func closeHandler(handle func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		handle()
		fmt.Println("\rExit")
		os.Exit(0)
	}()
}
