package app

import (
	"RushBananaBet/internal/handlers"
	"RushBananaBet/pkg/logger"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotApp struct {
	botApi *tgbotapi.BotAPI
}

func (botApp *BotApp) Start(stop chan os.Signal, handlers *handler.Handlers) {
	logger.Info("Bot started", "", "", "", nil)

	botApp.StartPolling(handlers)

	<-stop
	logger.Info("Bot stoped", "", "", "", nil)

	botApp.botApi.StopReceivingUpdates()
}

func (botApp *BotApp) StartPolling(handlers *handler.Handlers) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botApp.botApi.GetUpdatesChan(u)

	for update := range updates {
		go botApp.RouteUpdate(update, handlers)
	}
}

func (botApp *BotApp) RouteUpdate(update tgbotapi.Update, handlers *handler.Handlers) {
	if update.Message != nil {
		switch {
		case update.Message.Text == "/start":
			handlers.Start()
			// msg := handlers.HandleStart(update.Message)
			// bot.Send(msg)
			// // другие команды...
		case update.Message.Text == "/create-event":
			// ds
		case update.Message.Text == "/add-result":
			//ds
		case update.Message.Text == "/finish-tournament":
			// ыв
		case update.Message.Text == "/my-predictions":
			//ds
		case strings.Contains(update.Message.Text, "/match"):
			// dsd
		}
	}
}

func NewBotApp(botToken string, handlers *handler.Handlers) BotApp {
	botApi, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Fatal("Error creating newBot", "", "", "main-main()", err)
	}
	return BotApp{
		botApi: botApi,
	}
}

// func BuildKeyboard(username string) tgbotapi.ReplyKeyboardMarkup {
// 	var rows [][]tgbotapi.KeyboardButton

// 	// Админские кнопки
// 	if user.IsAdmin(username) {
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("➕ Добавить ивент"),
// 		))
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("🎯 Добавить результат"),
// 		))
// 		rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("🏁 Завершить турнир"),
// 		))
// 	}

// 	// Пользовательская кнопка
// 	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("📄 Мои ставки"),
// 	))

// 	// Кнопки матчей
// 	rows = append(rows, tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("⚔️ Матч 1"),
// 		tgbotapi.NewKeyboardButton("⚔️ Матч 2"),
// 	))

// 	return tgbotapi.NewReplyKeyboard(rows...)
// }
