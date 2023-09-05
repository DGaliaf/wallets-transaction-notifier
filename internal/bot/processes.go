package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wallet-transaction-notification/internal/bot/keyboards/keyboard"
)

func (b Bot) processStart(update tgbotapi.Update) {
	text := fmt.Sprintf("Добро пожаловать @%s!\n\nДля отслеживания кошелька используйте команду - /track\nили нажмите на кнопку 'Отслеживать'",
		update.Message.From.UserName)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = keyboard.MainKeyboard

	b.bot.Send(msg)
}

func (b Bot) processProfile(update tgbotapi.Update) {
	text := fmt.Sprintf("Профиль @%s",
		update.Message.From.UserName)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	//msg.ReplyMarkup = keyboard.MainKeyboard

	b.bot.Send(msg)
}

func (b Bot) processTrack(update tgbotapi.Update) {

}
