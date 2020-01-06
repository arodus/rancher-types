package client

const (
	TelegramConfigType                  = "telegramConfig"
	TelegramConfigFieldBotToken         = "botToken"
	TelegramConfigFieldDefaultRecipient = "defaultRecipient"
	TelegramConfigFieldProxyURL         = "proxyUrl"
)

type TelegramConfig struct {
	BotToken         string `json:"botToken,omitempty" yaml:"botToken,omitempty"`
	DefaultRecipient string `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	ProxyURL         string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
}
