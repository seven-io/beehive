<img src="https://www.seven.io/wp-content/uploads/Logo.svg" width="250" />

# [seven](https://www.seven.io) for [Beehive](https://github.com/muesli/beehive)

Send SMS message(s) to phone number(s).

## Installation

1. Fetch the package by running `go get github.com/seven-io/beehive`
2. Add `"github.com/seven-io/beehive"` to `import` in `hives.go`
3. Download module dependencies by running `go mod download`
4. Recompile the binary by running `make`

### Configuration

```json
{
  "Bees": [
    {
      "Name": "seven SMS Example Bee",
      "Class": "sevenbee",
      "Description": "A bee to demonstrate seven with Beehive",
      "Options": [
        {
          "Name": "api_key",
          "Value": "YOUR_SEVEN_API_KEY"
        },
        {
          "Name": "from",
          "Value": "+490123456789"
        },
        {
          "Name": "to",
          "Value": "+490123456789,+456789012345"
        }
      ]
    }
  ]
}
```

**api_key**: seven [API key](https://help.seven.io/en/api-key-access). Sign up and get one for free from [seven](https://www.seven.io/signup).

**from**: Sender number. It may contain a maximum of 11 alphanumeric or 16 numeric characters.

**to_number**: Recipient number – possible are numbers and address book entries (groups and contacts). Multiple
recipients can be specified separated by commas.

#### Actions

**sms**: Send SMS message(s). Needs the text of the message to send. You can use interpolation to send somthing from the
event received:

```json
{
  "Actions": [
    {
      "Bee": "seven SMS Example Bee",
      "Name": "sms",
      "Options": [
        {
          "Name": "text",
          "Type": "string",
          "Value": "The current time is {{.timestamp}}"
        }
      ]
    }
  ]
}
```

###### Support

Need help? Feel free to [contact us](https://www.seven.io/en/company/contact/).

[![MIT](https://img.shields.io/badge/License-MIT-teal.svg)](LICENSE)
