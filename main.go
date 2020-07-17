package main

import{
	"os"
	tb "gopkg.in/tucnak/telebot.v2"
}

func main(){
	var(
		port = os.Getenv("PORT") //Set automatically
		publicURL = os.Getenv("PUBLIC_URL") //Set at config
		token = os.Getenv("TOKEN")//Set at config
	)

	webhook := &tb.Webhook{
		Listen: ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token: token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil{
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message){
		b.Send(m.Sender, "Greetings, soul")
	})
}