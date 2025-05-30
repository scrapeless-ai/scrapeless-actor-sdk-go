# 更新日志

所有重要的变更都会记录在此文件中。

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/) 格式，并遵循 [语义化版本规范（Semantic Versioning）](https://semver.org/lang/zh-CN/)

---

## [Unreleased] — 尚未发布

### ✨ 新增
- CLI：添加 `scrapeless init` 命令用于初始化新 actor 项目。
- 新增对 Anthropic Claude 模型后端的支持。
- scraping 任务添加超时降级策略。

### 🐛 修复
- 修复 actor context 无法正确传递用户 API key 的问题。
- 修复 queue actor 在消息为空时报错的问题。

### 🚀 变更
- 优化 Redis 连接池逻辑。
- 内部模块统一使用 `zap` 日志库。

### 🧹 移除
- 移除已弃用的 `UniversalService.LegacyCall` 方法。

---

## [1.0.0] - 2025-05-26

### ✨ 新增
- 正式发布初版：完整的 actor 执行引擎。
- 支持 OpenAI、Claude 与本地模型后端。
- 多用户 API Key 鉴权系统。
- Redis 限流功能（令牌桶策略）。
- 日志记录系统（支持 Graylog 与数据库）。
- 模块化远程服务客户端接口。
- 示例模块：`scraping`、`captcha`、`proxy`、`router` 等。
- 内建开发用 HTTP Server，支持健康检查。

### 🐛 修复
- 修复 `proxy` 模块错误信息显示不准确的问题。
- 修复部分包加载环境变量不一致的问题。

### 🚀 变更
- 所有服务统一实现 `ActorService` 接口。
- 将 `storage` 模块拆分为 `dataset`、`queue`、`kv` 与 `object` 子模块。

---

<!-- 历史版本举例，如需可启用 -->
<!--
## [0.1.0] - 2025-04-01

### ✨ 新增
- 初始版本，包含基础的 actor 执行器与 CLI 工具。
-->

---

## 标记说明

- ✨ `新增`：新增功能
- 🐛 `修复`：问题修复
- 🚀 `变更`：功能优化、调整
- 🧹 `移除`：移除已弃用或无效的功能
