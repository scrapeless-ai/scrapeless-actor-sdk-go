package config

import (
	"github.com/scrapeless-ai/scrapeless-actor-sdk-go/internal/helper"
)

var (
	LogLevel                = helper.GetString("SCRAPELESS_LOG_LEVEL", "debug") // 日志级别
	GatewayServiceHost      = helper.GetString("SCRAPELESS_GATEWAY_SERVICE_HOST", "")
	StorageServiceHost      = helper.GetString("SCRAPELESS_STORAGE_SERVICE_HOST", "")
	ScrapingBrowserApiHost  = helper.GetString("SCRAPING_BROWSER_API_HOST", "")
	Token                   = helper.GetString("SCRAPELESS_API_KEY", "")
	ProxyCountry            = helper.GetString("PROXY_COUNTRY", "ANY")
	ProxySessionDurationMax = helper.GetInt64("PROXY_SESSION_DURATION_MAX", 120)
	ProxyGatewayHost        = helper.GetString("PROXY_GATEWAY_HOST", "gw-us.scrapeless.io:8789")
	ScrapelessApiHost       = helper.GetString("SCRAPELESS_API_HOST", "")
)
