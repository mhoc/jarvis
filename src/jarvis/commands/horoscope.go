package commands

import (
	"jarvis/util"
	"jarvis/ws"
	"math/rand"
)

type OnionHoroscope struct{}

func NewHoroscope() OnionHoroscope {
	return OnionHoroscope{}
}

func (h OnionHoroscope) Name() string {
	return "horoscope"
}

func (h OnionHoroscope) Description() string {
	return "gives you a random horoscope for this week"
}

func (h OnionHoroscope) Examples() []string {
	return []string{"jarvis horoscope"}
}

func (h OnionHoroscope) OtherDocs() []util.HelpTopic {
	return []util.HelpTopic{}
}

func (h OnionHoroscope) SubCommands() []util.SubCommand {
	return []util.SubCommand{
		util.NewSubCommand("^jarvis horoscope$", h.RandomThisWeek),
	}
}

func (h OnionHoroscope) RandomThisWeek(m util.IncomingSlackMessage, r util.Regex) {
	// Thanks for the API anonymous internet user!
	url := "http://a.knrz.co/horoscope-api/current"
	data, err := util.HttpGetArr(url)
	if err != nil {
		ws.SendMessage("I've encountered an error contacting the API I use for horoscopes. Apologies.", m.Channel)
		return
	}
	horoscope := data[rand.Intn(len(data))]
	message := horoscope["prediction"].(string)
	ws.SendMessage(message, m.Channel)
}
