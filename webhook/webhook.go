package webhook

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"telegram-bot-bc/binance"
	"time"
)

const (
	TELEMGRAM_BOT_USERNAME = "testbitcoinprice_bot"
	TELEGRAME_BOT_LINK = "t.me/testbitcoinprice_bot."
	TELEGRAM_BOT_TOKEN = "631612541:AAFZiNDk4Ui9Sudl0xh_-r6Lx3W7PL8Ibmw"
	SLEEP_TIME = 500
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)


func InitTelegram() (api *tgbotapi.BotAPI, err error) {
	botTgr, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_TOKEN)
	if err != nil {
		return nil, err
	}

	//u := tgbotapi.NewUpdate(0)
	return botTgr, nil
}

func HandleTelegramMess(api *tgbotapi.BotAPI, b *binance.BitCoin) {
	fmt.Println("handling....")

	var messID int = 0

	for {

		time.Sleep(500 * time.Millisecond)

		updateMesses := tgbotapi.NewUpdate(messID)
		updateMesses.Timeout = 60

		updates, err := api.GetUpdatesChan(updateMesses)

		if err != nil {
			continue
		}

		for updateMess := range updates {
			if updateMess.Message == nil {
				continue
			}

			command := strings.ToUpper(updateMess.Message.Command())
			messResponse := ""

			for key, price := range b.Prices {
				if command == key {
					messResponse = convertPriceToString(price, key)
				}
			}

			if messResponse == "" {
				messResponse = convertAllPricesToString(b)
			}

			msg := tgbotapi.NewMessage(updateMess.Message.Chat.ID, messResponse)
			msg.ReplyToMessageID = updateMess.Message.MessageID
			msg.ParseMode = tgbotapi.ModeHTML

			_, err = api.Send(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

/*func convertAllPricesToString(b *binance.BitCoin) string {
	str := ""

	for key, price := range b.Prices {
		if price.OpenPrice == 0 || price.ClosePrice == 0 {
			continue
		}

		line := fmt.Sprintf("name: %s, open-price: %f, close-price: %f \n", key, price.OpenPrice, price.OpenPrice)
		str += line
	}

	return str
}*/

func convertPriceToString(price binance.Prices, code string) string {
	if price.OpenPrice == 0 || price.ClosePrice == 0 {
		return ""
	}

	return fmt.Sprintf("%s - open price: %f - close price : %f", code, price.OpenPrice, price.ClosePrice)
}

func convertAllPricesToString(b *binance.BitCoin) string {
	str := "Hi, \n I am bot chat who will give you information about coin/'s price. You should type some commands above "
	isAppend := false

	for key, price := range b.Prices {
		isAppend = true
		if price.OpenPrice != 0 && price.ClosePrice != 0 {
			str += fmt.Sprintf("\n /%s", key)
		}
	}

	if !isAppend {
		return ""
	} else {
		return str
	}

}
