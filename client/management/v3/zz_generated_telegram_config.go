package client

const (
	TelegramConfigType                       = "telegramConfig"
	TelegramConfigFieldBotToken              = "botToken"
	TelegramConfigFieldDefaultRecipient      = "defaultRecipient"
	TelegramConfigFieldDisableNotification   = "disableNotification"
	TelegramConfigFieldDisableWebPagePreview = "disableWebPagePreview"
	TelegramConfigFieldProxyURL              = "proxyUrl"
)

type TelegramConfig struct {
	BotToken              string `json:"botToken,omitempty" yaml:"botToken,omitempty"`
	DefaultRecipient      string `json:"defaultRecipient,omitempty" yaml:"defaultRecipient,omitempty"`
	DisableNotification   bool   `json:"disableNotification,omitempty" yaml:"disableNotification,omitempty"`
	DisableWebPagePreview bool   `json:"disableWebPagePreview,omitempty" yaml:"disableWebPagePreview,omitempty"`
	ProxyURL              string `json:"proxyUrl,omitempty" yaml:"proxyUrl,omitempty"`
}
