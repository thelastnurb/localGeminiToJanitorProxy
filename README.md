# Gemini-OpenAI Proxy (Optimized for Render) BY NURB

Этот сервер - прослойка (proxy), которая позволяет использовать мощные модели Google Gemini через интерфейс приложений, поддерживающих OpenAI (Janitor AI, SillyTavern и др.).

**Репозиторий:** `https://github.com/thelastnurb/localGeminiToJanitorProxy`

## Основные возможности
- **24/7 Работа**: Благодаря деплою на Render, прокси доступен всегда.
- **HTTPS из коробки**: Решает проблему "Failed to fetch" в браузере при работе с Janitor AI.
- **Стриминг (SSE)**: Текст печатается постепенно в реальном времени.
- **Поддержка любых моделей**: Gemini 3.0 Flash(3.1 preview), 2.5 Pro и другие.
- **Безопасность**: Ваши ключи API задаются через переменные окружения и не светятся в коде.

## 🛠 Пошаговая инструкция по деплою на Render.com

1. **Создайте аккаунт** на [Render.com](https://render.com) через ваш GitHub(или другую сетку).
2. Нажмите синюю кнопку **New +** -> **Web Service**.
3. Откроется меню выбирите Public Git Repository. Вставьте `https://github.com/thelastnurb/localGeminiToJanitorProxy` и нажмите **Connect**.
4. В настройках укажите следующие данные:
   - **Runtime**: `Go`
   - **Build Command**: `go build -o server main.go`
   - **Start Command**: `./server`
   - **Instance Type**: `Free`
5. Нажмите кнопку **Advanced** и добавьте переменные окружения (**Environment Variables**):
   - `GEMINI_API_KEY`: Ваш ключ, полученный в [Google AI Studio](https://aistudio.google.com/app/apikey).
   - `PROXY_API_KEY`: Любой придуманный вами пароль API (нужен для авторизации в Janitor AI).
6. Нажмите **Create Web Service** и дождитесь надписи "Your service is live".

После завершения вы получите ссылку вида: `https://имя-вашего-сервиса.onrender.com`.

## ⚙️ Настройка в Janitor AI

1. Откройте **API Settings** и выберите пункт **proxy**.
2. В поле **Reverse Proxy URL (OpenAI Base URL)** вставьте вашу ссылку от Render и добавьте `/v1` в конец.
   *Пример: `https://my-proxy.onrender.com/v1`*
3. В поле **OpenAI Key** введите ваш секретный пароль API, который вы указали в поле `PROXY_API_KEY`.
4. В поле **Model** выберите Custom и впишите: `gemini-flash-latest`(когда вы пишите latest автоматом подгружается последняя версия флешки. Например сейчас это 3.0)  (или другую модель Gemini).
5. Нажмите Save Settings и начните чат!

## 🔐 Безопасность и обновления
ЕСЛИ БУДЕТЕ ДЕЛАТЬ ФОРК МОЕГО РЕПОЗИТОРИЯ
- Никогда не загружайте файл `.env` на GitHub. Файл `.gitignore` в этом проекте настроен так, чтобы скрывать его автоматически.
- Чтобы обновить прокси, просто внесите изменения в код и сделайте `git push` в ваш репозиторий. Render сам подхватит изменения и пересоберет сервер.

---
*Разработано для использования с Janitor AI и другими ролевыми фронтендами. 
Если вы хотите чтобы я добавил фичи по типу отдельного сайта или распознавания фото или подсчет потраченных токенов - маякните мне в [telegram](https://t.me/thenurb)*
пис
