package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wallet-transaction-notification/internal/bot/keyboards/keyboard"
)

func (b Bot) processStart(ctx context.Context, update tgbotapi.Update) {
	if err := b.db.CreateUser(ctx, update.Message.From.ID); err != nil {
		log.Panicln(err)
	}

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

	b.bot.Send(msg)
}

func (b Bot) processTrack(update tgbotapi.Update) {
	*b.token = update.Message.Text
}

func (b Bot) processAddWallet(ctx context.Context, update tgbotapi.Update) {
	var text string

	if err := b.db.AddWallet(ctx, update.Message.From.ID, update.Message.Text); err != nil {
		text = err.Error()
	} else {
		text = fmt.Sprintf("Кошелек - %s был успешно добавлен!", update.Message)
	}

	msg := tgbotapi.NewMessage(update.Message.From.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}
