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
	if err != nil{
		log.Fatal(err)
	}

	b.Handle("/poem", func(m *tb.Message) {
		b.Send(m.Sender, 
			`The lazy artist-boor is blacking
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
			Of virgin days I left behind.`)
		b.Send(m.Sender, `By Aleksandr Pushkin`)
	})

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Greetings, soul")
	})

	b.Handle("", func(m *tb.Message) {
		b.Send(m.Sender, "You entered "+m.Text)
	})
	  
	b.Start()
}