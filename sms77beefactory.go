package sms77bee

import (
	"github.com/muesli/beehive/bees"
)

type Sms77BeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *Sms77BeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := Sms77Bee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *Sms77BeeFactory) ID() string {
	return "sms77bee"
}

// Name returns the name of this Bee.
func (factory *Sms77BeeFactory) Name() string {
	return "Sms77"
}

// Description returns the description of this Bee.
func (factory *Sms77BeeFactory) Description() string {
	return "Sends SMS messages"
}

// Image returns the filename of an image for this Bee.
func (factory *Sms77BeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *Sms77BeeFactory) LogoColor() string {
	return "#18D46A"
}

// Options returns the options available to configure this Bee.
func (factory *Sms77BeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "api_key",
			Description: "Sms77 API key",
			Type:        "string",
			Mandatory:   true,
		},
		{
			Name:        "to",
			Description: "Phone number(s) to send SMS messages to (ex: \"+490123456789,+456789012345\")",
			Type:        "string",
			Mandatory:   true,
		},
		{
			Name:        "from",
			Description: "Phone number to send SMS messages from (ex: \"+490123456789\")",
			Type:        "string",
			Mandatory:   false,
		},
	}
	return opts
}

// Actions describes the available actions provided by this Bee.
func (factory *Sms77BeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "sms",
			Description: "Sends SMS message(s)",
			Options: []bees.PlaceholderDescriptor{
				{
					Description: "Message text",
					Mandatory:   true,
					Name:        "text",
					Type:        "string",
				},
			},
		},
	}
	return actions
}

func init() {
	f := Sms77BeeFactory{}
	bees.RegisterFactory(&f)
}
