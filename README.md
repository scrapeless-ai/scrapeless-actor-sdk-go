# Scrapeless Actor SDK Go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**[English](README.md) | [中文文档](README-zh.md)**

The official Go SDK of [Scrapeless AI](https://scrapeless.com) - a powerful web scraping and browser automation platform that helps you extract data from any website at scale.

## 📑 Table of Contents

- [🌟 Features](#-features)
- [📦 Installation](#-installation)
- [🚀 Quick Start](#-quick-start)
- [📖 Usage Examples](#-usage-examples)
- [🔧 API Reference](#-api-reference)
- [📚 Examples](#-examples)
- [🧪 Testing](#-testing)
- [🛠️ Contribution & Development Guide](#️-contribution--development-guide)
- [📄 License](#-license)
- [📞 Support](#-support)
- [🏢 About Scrapeless](#-about-scrapeless)

## 🌟 Features

- **Browser Automation**: Supports remote browser session operations.
- **Web Scraping**: Extracts data from any website through intelligent parsing.
- **SERP Scraping**: Extracts search engine results with high accuracy.
- **Proxy Management**: Built-in proxy rotation and geolocation.
- **Actor System**: Runs custom automation scripts in the cloud.
- **Storage Solutions**: Provides persistent data storage for your scraping projects.

## 📦 Installation

Install the SDK using `go get`:

```bash
go get -u github.com/scrapeless-ai/scrapeless-actor-sdk-go
```

## 🚀 Quick Start

### Basic Setup

```go
package main

import (
	scrapeless "github.com/scrapeless-ai/scrapeless-actor-sdk-go/scrapeless/actor"
)

func main() {
	// Initialize the actor
	actor := scrapeless.New()
	defer actor.Close()
}
```

### Environment Variables

You can also configure the SDK using environment variables:

```bash
# Required
SCRAPELESS_API_KEY=your-api-key

# Optional - Custom API endpoints
SCRAPELESS_BASE_API_URL=https://api.scrapeless.com
SCRAPELESS_ACTOR_API_URL=https://actor.scrapeless.com
SCRAPELESS_STORAGE_API_URL=https://storage.scrapeless.com
SCRAPELESS_BROWSER_API_URL=https://browser.scrapeless.com
SCRAPELESS_CRAWL_API_URL=https://crawl.scrapeless.com
```

## 📖 Usage Examples

### Browser Automation

```go
// Example code will be written according to the actual browser package.
// ... To be supplemented with specific code ...
```

### Web Scraping

```go
// Example code will be written according to the actual web scraping package.
// ... To be supplemented with specific code ...
```

### SERP Scraping

```go
// Example code will be written according to the actual SERP scraping package.
// ... To be supplemented with specific code ...
```

### Actor System

```go
// Example code will be written according to the actual Actor system package.
// ... To be supplemented with specific code ...
```

## 🔧 API Reference

### Available Services

The SDK provides the following services:

- `scrapeless.browser` - Browser session management.
- `scrapeless.scraping` - Web scraping and data extraction.
- `scrapeless.deepserp` - Search engine result extraction.
- `scrapeless.universal` - Universal data extraction.
- `scrapeless.proxies` - Proxy management.
- `scrapeless.actor` - Actor system for custom automation.
- `scrapeless.storage` - Data storage solutions.
- `scrapeless.scrapingCrawl` - Website crawling.

### Error Handling

```go
// Example code will be written according to the actual error handling logic.
// ... To be supplemented with specific code ...
```

## 📚 Examples

Check the `example` directory for complete usage examples:

- [Browser Operation Example](example/browser/browser.go)
- [Captcha Recognition Example](example/captcha/captcha.go)
- [Proxy Management Example](example/proxy/proxy.go)
- [Storage Usage Example](example/storage_*)

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

The SDK includes comprehensive tests for all services and tools.

## 🛠️ Contribution & Development Guide

All forms of contributions are welcome! For detailed information on how to submit issues, PRs, code specifications, local development, etc., please refer to the [Contribution & Development Guide](./CONTRIBUTING.md).

**Quick Start**:

```bash
git clone https://github.com/your-repo-path/scrapeless-actor-sdk-go.git
cd scrapeless-actor-sdk-go
go mod tidy
go test ./...
```

Please replace `github.com/your-repo-path` with the actual repository path.

For more information on project structure, best practices, etc., please refer to [CONTRIBUTING.md](./CONTRIBUTING.md).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

- 📖 **Documentation**: [https://docs.scrapeless.com](https://docs.scrapeless.com)
- 💬 **Community**: [Join our Discord](https://backend.scrapeless.com/app/api/v1/public/links/discord)
- 🐛 **Issues**: [GitHub Issues](https://github.com/your-repo-path/scrapeless-actor-sdk-go/issues)
- 📧 **Email**: [support@scrapeless.com](mailto:support@scrapeless.com)

Please replace `github.com/your-repo-path` with the actual repository path.

## 🏢 About Scrapeless

Scrapeless is a powerful web scraping and browser automation platform that helps enterprises extract data from any website at scale. Our platform provides:

- High-performance web scraping infrastructure.
- Global proxy network.
- Browser automation capabilities.
- Enterprise-level reliability and support.

Visit [scrapeless.com](https://scrapeless.com) to learn more and get started.

---

Made with ❤️ by the Scrapeless team
```

        