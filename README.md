# Gemini-to-OpenAI Proxy for Janitor AI

Это легковесный, быстрый бэкенд на чистом Go, который конвертирует запросы из формата OpenAI (который использует Janitor AI, SillyTavern и т.д.) в формат Google Gemini API и обратно.

## Особенности
- Полная поддержка стриминга (SSE).
- Поддержка System Prompts (через `systemInstruction` в Gemini).
- Минимум зависимостей (только `godotenv`).
- Работает напрямую с REST API Gemini (без SDK).

## Как запустить

1. Установите [Go 1.22+](https://go.dev/dl/).
2. Склонируйте или скачайте проект.
3. Скопируйте файл `.env.example` в `.env`:
   ```bash
   cp .env.example .env