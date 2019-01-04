package main

import (
	"fmt"
	"runtime"
	"telegram-bot-bc/binance"
	"telegram-bot-bc/webhook"
	"time"
)

func main()  {
	/*tlgApi, err := webhook.InitTelegram()
	if err != nil {
		fmt.Println("#Err ", err)
		return
	}

	webhook.HandleTelegramMess(tlgApi)*/

	bitCoin := binance.NewBitCoin()
	go bitCoin.GetCoinData()

	tlgApi, err := webhook.InitTelegram()
	if err != nil {
		fmt.Println("#Err ", err)
		return
	}
	go func() {
		for {
			webhook.HandleTelegramMess(tlgApi, bitCoin)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	//terminate routine call system exit
	runtime.Goexit()
}
