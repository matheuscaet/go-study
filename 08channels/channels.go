package channels

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_BITCOIN_VALUE = 5
const MAX_ETHEREUM_VALUE = 10

func Channels() {
	fmt.Println("-------------------------------- 08 channels starts --------------------------------")

	ch := make(chan int)

	go process(ch)

	for v := range ch {
		fmt.Println(v)
	}

	var platforms = []string{"binance", "coinbase", "kraken", "bitstamp", "bitfinex"}
	var bitcoinChannel = make(chan string)
	var ethereumChannel = make(chan string)

	for _, platform := range platforms {
		go checkBitcoinPrice(platform, bitcoinChannel)
		go checkEthereumPrice(platform, ethereumChannel)
	}

	sendPrices(bitcoinChannel, ethereumChannel)

	fmt.Println("-------------------------------- 08 channels ends --------------------------------")
}

func process(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func sendPrices(bitcoinChannel chan string, ethereumChannel chan string) {
	select {
	case price := <-bitcoinChannel:
		fmt.Printf("Price found in bitcoin in %s\n", price)
	case price := <-ethereumChannel:
		fmt.Printf("Price found in ethereum in %s\n", price)
	}
}

func checkBitcoinPrice(platform string, bitcoinChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var bitcoinPrice = rand.Int() * 20
		if bitcoinPrice > MAX_BITCOIN_VALUE {
			bitcoinChannel <- platform
			break
		}
	}
}

func checkEthereumPrice(platform string, ethereumChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var ethereumPrice = rand.Int() * 20
		if ethereumPrice > MAX_ETHEREUM_VALUE {
			ethereumChannel <- platform
			break
		}
	}
}
