package binance

import (
	"context"
	"github.com/binance-exchange/go-binance"
	"github.com/go-kit/kit/log"
	"os"
	"time"
)

const (
	API_SECRET = "YOUR_API_SCERET"
	API_KEY    = "YOUR_API_KEY"
)

var (
	COIN_CODES = [...]string{"ETHBTC", "EOSBTC"}
)

type BitCoin struct {
	Binance binance.Binance
	Prices  map[string]Prices
}

type Prices struct {
	OpenPrice  float64
	ClosePrice float64
}

func NewBitCoin() *BitCoin {
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "time", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	hmacSigner := &binance.HmacSigner{
		Key: []byte(API_SECRET),
	}

	ctx, _ := context.WithCancel(context.Background())

	binanceService := binance.NewAPIService(
		"https://www.binance.com",
		API_KEY,
		hmacSigner,
		logger,
		ctx,
	)

	b := binance.NewBinance(binanceService)

	mMap := make(map[string]Prices)
	for _, code := range COIN_CODES {
		mMap[code] = Prices{}
	}

	return &BitCoin{Binance: b, Prices: mMap}
}

func (b *BitCoin) GetKLine() {
	for _, code := range COIN_CODES {
		kline, err := b.Binance.Klines(binance.KlinesRequest{Symbol: code, Interval: binance.Minute})
		p := Prices{}
		if err == nil {
			line := kline[len(kline)-1]
			p = Prices{OpenPrice: line.Open, ClosePrice: line.Close}
		}

		b.Prices[code] = p
	}
}

func(b *BitCoin) GetCoinData()  {
	for {
		b.GetKLine()
		time.Sleep(500 * time.Millisecond)
	}
}
