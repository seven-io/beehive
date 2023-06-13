package sevenbee

import (
	"github.com/muesli/beehive/bees"
	seven "github.com/seven-io/go-client/sms77api"
)

// SevenBee is a Bee that is able to send SMS messages.
type SevenBee struct {
	bees.Bee

	client *seven.Sms77API
	apiKey string
	from   string
	to     string
}

// Action triggers the action passed to it.
func (mod *SevenBee) Action(action bees.Action) []bees.Placeholder {
	var outs []bees.Placeholder

	switch action.Name {
	case "sms":
		text := ""
		action.Options.Bind("text", &text)
		_, err := mod.client.Sms.Json(seven.SmsBaseParams{
			From: mod.from,
			Text: text,
			To:   mod.to,
		})
		if err != nil {
			mod.LogErrorf("Error sending seven SMS: %s", err)
		}

	default:
		panic("Unknown action triggered in " + mod.Name() + ": " + action.Name)
	}

	return outs
}

// Run executes the Bee's event loop.
func (mod *SevenBee) Run(eventChan chan bees.Event) {
	mod.client = seven.New(seven.Options{ApiKey: mod.apiKey, SentWith: "BeeHive"})
}

// ReloadOptions parses the config options and initializes the Bee.
func (mod *SevenBee) ReloadOptions(options bees.BeeOptions) {
	mod.SetOptions(options)

	options.Bind("api_key", &mod.apiKey)
	options.Bind("from", &mod.from)
	options.Bind("to", &mod.to)
}
