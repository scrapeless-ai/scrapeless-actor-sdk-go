package config

import (
	"github.com/spf13/viper"
)

var (
	LogLevel                = viper.GetString("SCRAPELESS_LOG_LEVEL")
	ScrapingBrowserApiHost  = viper.GetString("SCRAPELESS_BROWSER_API_HOST")
	Token                   = viper.GetString("SCRAPELESS_TOKEN")
	ProxyCountry            = viper.GetString("SCRAPELESS_PROXY_COUNTRY")
	ProxySessionDurationMax = viper.GetInt64("SCRAPELESS_PROXY_SESSION_DURATION_MAX")
	ProxyGatewayHost        = viper.GetString("SCRAPELESS_PROXY_GATEWAY_HOST")
	ScrapelessApiHost       = viper.GetString("SCRAPELESS_API_HOST")     // captcha
	ScrapelessCaptchaHost   = viper.GetString("SCRAPELESS_CAPTCHA_HOST") // captcha
)
