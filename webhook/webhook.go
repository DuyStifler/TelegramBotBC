package webhook

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

const (
	TELEMGRAM_BOT_USERNAME = "testbitcoinprice_bot"
	TELEGRAME_BOT_LINK = "t.me/testbitcoinprice_bot."
	TELEGRAM_BOT_TOKEN = "631612541:AAFZiNDk4Ui9Sudl0xh_-r6Lx3W7PL8Ibmw"
	SLEEP_TIME = 500
)

func InitTelegram() (api *tgbotapi.BotAPI, err error) {
	botTgr, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_TOKEN)
	if err != nil {
		return nil, err
	}

	//u := tgbotapi.NewUpdate(0)
	return botTgr, nil
}

func HandleTelegramMess(api *tgbotapi.BotAPI) {
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

			text := updateMess.Message.
		}
	}

}
