# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

### ✨ Added
- CLI: `scrapeless init` command to scaffold a new actor.
- Support for Anthropic Claude backend.
- Timeout fallback mechanism for long-running scraping jobs.

### 🐛 Fixed
- Fix actor context not propagating user API key metadata.
- Resolve panic in queue actor on empty message payload.

### 🚀 Changed
- Improve Redis connection pooling logic.
- Use `zap` logger in all internal modules.

### 🧹 Removed
- Deprecated `UniversalService.LegacyCall` method removed.

---

## [1.0.0] - 2025-05-26

### ✨ Added
- Initial release with complete actor execution engine.
- Support for OpenAI, Claude, and Local model backends.
- Multi-user API key authentication system.
- Redis-based rate limiter (token bucket strategy).
- Logging integration with Graylog and database output.
- Modular client interface for all remote services.
- Example modules: `scraping`, `captcha`, `proxy`, `router`, etc.
- Built-in development HTTP server (`httpserver`) with health check.

### 🐛 Fixed
- Correct error message formatting in `proxy` service.
- Normalize environment variable loading across packages.

### 🚀 Changed
- All services now follow unified `ActorService` interface.
- Split `storage` service into `dataset`, `queue`, `kv`, and `object`.

---

<!-- Historical version example, uncomment if needed
## [0.1.0] - 2025-04-01

### ✨ Added
- Initial proof of concept with basic actor runner and CLI.
-->

---

## Legend

- ✨ `Added`: New features
- 🐛 `Fixed`: Bug fixes
- 🚀 `Changed`: Enhancements or modifications
- 🧹 `Removed`: Deprecated or removed functionality
