package client

const (
	NotificationType                 = "notification"
	NotificationFieldMessage         = "message"
	NotificationFieldPagerdutyConfig = "pagerdutyConfig"
	NotificationFieldSMTPConfig      = "smtpConfig"
	NotificationFieldSlackConfig     = "slackConfig"
	NotificationFieldTelegramConfig  = "telegramConfig"
	NotificationFieldWebhookConfig   = "webhookConfig"
	NotificationFieldWechatConfig    = "wechatConfig"
)

type Notification struct {
	Message         string           `json:"message,omitempty" yaml:"message,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty" yaml:"pagerdutyConfig,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty" yaml:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty" yaml:"slackConfig,omitempty"`
	TelegramConfig  *TelegramConfig  `json:"telegramConfig,omitempty" yaml:"telegramConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty" yaml:"webhookConfig,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty" yaml:"wechatConfig,omitempty"`
}
