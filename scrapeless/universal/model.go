package universal

type Actor string

var (
	ScraperUniversal Actor = "unlocker.webunlocker"
)

type ScrapingTaskRequest struct {

	// Input parameters for the scraper
	Input map[string]any `json:"input"`

	// Proxy configuration
	ProxyCountry string `json:"country"`
}

type ScrapingTaskResponse struct {
	// Response message
	Message string `json:"message"`

	// Task ID for tracking the request
	TaskID string `json:"taskId"`
}
