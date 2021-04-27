package sms77bee

import (
	"github.com/muesli/beehive/bees"
	"github.com/sms77io/go-client/sms77api"
)

// Sms77Bee is a Bee that is able to send SMS messages.
type Sms77Bee struct {
	bees.Bee

	client *sms77api.Sms77API
	apiKey string
	from   string
	to     string
}

// Action triggers the action passed to it.
func (mod *Sms77Bee) Action(action bees.Action) []bees.Placeholder {
	var outs []bees.Placeholder

	switch action.Name {
	case "sms":
		text := ""
		action.Options.Bind("text", &text)
		_, err := mod.client.Sms.Json(sms77api.SmsBaseParams{
			From: mod.from,
			Text: text,
			To:   mod.to,
		})
		if err != nil {
			mod.LogErrorf("Error sending sms77 SMS: %s", err)
		}

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// Run executes the Bee's event loop.
func (mod *Sms77Bee) Run(eventChan chan bees.Event) {
	mod.client = sms77api.New(sms77api.Options{ApiKey: mod.apiKey, SentWith: "BeeHive"})
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *Sms77Bee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)

	options.Bind("api_key", &mod.apiKey)
	options.Bind("from", &mod.from)
	options.Bind("to", &mod.to)
}
