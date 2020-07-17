package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	poem := `The lazy artist-boor is blacking
	The genius's picture with his stuff,
	Without any sense a-making
	His low drawing above.
		
	But alien paints, in stride of years,
	Are falling down as a dust,
	The genius's masterpiece appears
	With former brilliance to us.
		
	Like this, the darkly apparitions
	Are leaving off my tortured heart,
	And it again revives the visions
	Of virgin days I left behind.`

	inlineBtn1 := tb.InlineButton{
		Unique: "moon",
		Text:   "Moon 🌚",
	}

	inlineBtn2 := tb.InlineButton{
		Unique: "sun",
		Text:   "Sun 🌞",
	}

	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn1, inlineBtn2},
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, `I greet you human!
		Here is the list of commands that I know:
		- /hello
		- /poem
		- /pick_time`)
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Greetings, soul")
	})

	b.Handle("/poem", func(m *tb.Message) {
		b.Send(m.Sender, poem)
		b.Send(m.Sender, `By Aleksandr Pushkin`)
	})

	b.Handle(&inlineBtn1, func(c *tb.Callback) {
		// Required for proper work
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		// Send messages here
		b.Send(c.Sender, "Moon says 'Hi'!")
	})

	b.Handle(&inlineBtn2, func(c *tb.Callback) {
		b.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		b.Send(c.Sender, "Sun says 'Hi'!")
	})

	b.Handle("/pick_time", func(m *tb.Message) {
		b.Send(
			m.Sender,
			"Day or night, you choose",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
	})

	b.Start()
}
