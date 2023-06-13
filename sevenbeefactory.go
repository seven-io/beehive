package sevenbee

import (
	"github.com/muesli/beehive/bees"
)

type SevenBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *SevenBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := SevenBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *SevenBeeFactory) ID() string {
	return "sevenbee"
}

// Name returns the name of this Bee.
func (factory *SevenBeeFactory) Name() string {
	return "seven"
}

// Description returns the description of this Bee.
func (factory *SevenBeeFactory) Description() string {
	return "Sends SMS messages"
}

// Image returns the filename of an image for this Bee.
func (factory *SevenBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *SevenBeeFactory) LogoColor() string {
	return "#18D46A"
}

// Options returns the options available to configure this Bee.
func (factory *SevenBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "api_key",
			Description: "seven API key",
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
func (factory *SevenBeeFactory) Actions() []bees.ActionDescriptor {
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
	f := SevenBeeFactory{}
	bees.RegisterFactory(&f)
}
