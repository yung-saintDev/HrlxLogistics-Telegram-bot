package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {

	apiToken := os.Getenv("TELEGRAM_API_TOKEN")
       if apiToken == "" {
           log.Fatal("API Token not set")
       }

       // Initialize your bot with apiToken
       bot, err := telego.NewBot(apiToken)
       if err != nil {
           log.Fatal(err)
       }


	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик /start
		chatID := tu.ID(update.Message.Chat.ID)


		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("🛒 Рассчитать стоимость"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("💴 Узнать курс"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("🥋 Ассортимент"),
				tu.KeyboardButton("📞 Контакты"),
			),
		).WithResizeKeyboard()

		message := tu.Message(
			chatID, 
			"Привет🤝\nHrlxLogistics - лучший вариант для заказа\nс китайских площадок, таких как: Poizon,\n1688, taobao и др.\nУ нас самая выгодная цена!\nБыстрая доставка\n\nℹ️Этот бот поможет быстро рассчитать стоимость вещи.\n\n✅ @hrlxLogisticsss - пишите сюда для\nоформления заказа, либо если остались\nинтересующие вопросы. С радостью со\nвсем подскажем 😉",
		).WithReplyMarkup(keyboard)

		_, _= bot.SendMessage(message)

	}, th.CommandEqual("start"))



	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик назад
		chatID := tu.ID(update.Message.Chat.ID)


		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("🛒 Рассчитать стоимость"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("💴 Узнать курс"),
			),
			tu.KeyboardRow(
				tu.KeyboardButton("🥋 Ассортимент"),
				tu.KeyboardButton("📞 Контакты"),
			),
		).WithResizeKeyboard()

		message := tu.Message(
			chatID, 
			"Привет🤝\nHrlxLogistics - лучший вариант для заказа\nс китайских площадок, таких как: Poizon,\n1688, taobao и др.\nУ нас самая выгодная цена!\nБыстрая доставка\n\nℹ️Этот бот поможет быстро рассчитать стоимость вещи.\n\n✅ @hrlxLogisticsss - пишите сюда для\nоформления заказа, либо если остались\nинтересующие вопросы. С радостью со\nвсем подскажем 😉",
		).WithReplyMarkup(keyboard)

		_, _= bot.SendMessage(message)

	}, th.TextEqual("🔙 Назад"))
	
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик стоимости
		chatID := tu.ID(update.Message.Chat.ID)

		backButton := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("🔙 Назад"),
			),
		).WithResizeKeyboard()

		enterPrice := tu.Message(
			chatID, 
			"Введите цену товара в Юанях",
		).WithReplyMarkup(backButton)

		_, _ = bot.SendMessage(enterPrice)

	}, th.TextEqual("🛒 Рассчитать стоимость"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// обработчик расчета цены
		chatID := tu.ID(update.Message.Chat.ID)

		text := update.Message.Text
		price, err := strconv.Atoi(text)
		if err != nil {
			// handle err
			return
		}

		var result int
		if price < 1000 {
			result = int(float64(price*14) * 1.1)
		} else if price >= 1000{
			result = price*14 + 1000
		} else if price >= 2000{
			result = price*14 + 1500
		} else if price >= 4000{
			result = price *14 + 2500
		}

		backButton := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("🔙 Назад"),
			),
		).WithResizeKeyboard()


		resultMessage := tu.Message(
			chatID, 
			fmt.Sprintf("Цена с учетом доставки: %d₽ + 480₽/кг\n\nВ стоимость включено: \n- Выкуп товара с Poizon\n- Доставка по Китаю\n- Доставка из Китая до склада в Москве\n- Страхование груза\n\nДля заказа @hrlxLogisticsss", result),
		).WithReplyMarkup(backButton)


		_, _ = bot.SendMessage(resultMessage)
	}, func(update telego.Update) bool {
		//check if message is a number
		_, err := strconv.Atoi(update.Message.Text)
		return err == nil
	})

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик курса
		chatID := tu.ID(update.Message.Chat.ID)

		backButton := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("🔙 Назад"),
			),
		).WithResizeKeyboard()

		courseOfCNY := tu.Message(
			chatID,
			"1₽ = 14.0 ¥",
		).WithReplyMarkup(backButton)

		_, _ = bot.SendMessage(courseOfCNY)
	}, th.TextEqual("💴 Узнать курс"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//обработчик ассортимента
		chatID := tu.ID(update.Message.Chat.ID)

		link := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("ТУТ",).WithURL("https://t.me/HrlxLogisticss"),
			),
		)

	
		inStock := tu.Message(
			chatID,
			"Весь товар в наличии вы можете посмотреть по ссылке внизу",
		).WithReplyMarkup(link)

		_, _ = bot.SendMessage(inStock)
		

	}, th.TextEqual("🥋 Ассортимент"))




	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//Обработчик обратной связи
		chatID := tu.ID(update.Message.Chat.ID)

		contactButton := tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Написать менеджеру 👨🏻‍💼").WithURL("https://t.me/@hrlxLogisticsss"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Обменник 💱").WithURL("https://t.me/HrlxExchange"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Наш телеграм канал 📢").WithURL("https://t.me/HrlxLogisticss"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Разработчик 👨‍💻").WithURL( "https://t.me/varUserName"),
			),
		)

		contactText := tu.Message(
			chatID,
			"Наши контакты",
		).WithReplyMarkup(contactButton)
		_, _ = bot.SendMessage(contactText)
	}, th.TextEqual("📞 Контакты"))

	bh.Start()
}



