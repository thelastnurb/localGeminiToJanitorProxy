# Gemini-OpenAI Proxy (Optimized for Render) BY NURB

Этот сервер - прослойка (proxy), которая позволяет использовать мощные модели Google Gemini через интерфейс приложений, поддерживающих OpenAI (Janitor AI, SillyTavern и др.).

**Репозиторий:** `https://github.com/thelastnurb/desktop-tutorial`

##  Основные возможности
- **24/7 Работа**: Благодаря деплою на Render, прокси доступен всегда.
- **HTTPS из коробки**: Решает проблему "Failed to fetch" в браузере.
- **Стриминг (SSE)**: Текст в Janitor AI печатается постепенно, как у оригинала.
- **Поддержка любых моделей**: Gemini 2.0 Flash, 1.5 Pro и др.
- **Безопасность**: Ваши ключи API не хранятся в коде, а задаются через переменные окружения.

## 🛠 Инструкция по деплою на Render.com

1. **Создайте аккаунт** на [Render.com](https://render.com) через ваш GitHub.
2. Нажмите **New +** -> **Web Service**.
3. Подключите ваш репозиторий `thelastnurb/desktop-tutorial`.
4. В настройках укажите:
   - **Runtime**: `Go`
   - **Build Command**: `go build -o server main.go`
   - **Start Command**: `./server`
   - **Instance Type**: `Free`
5. Нажмите кнопку **Advanced** и добавьте **Environment Variables**:
   - `GEMINI_API_KEY`: Ваш ключ из [Google AI Studio](https://aistudio.google.com/app/apikey).
   - `PROXY_API_KEY`: Ваш секретный пароль (придумайте любой), чтобы никто другой не мог пользоваться вашим прокси.
6. Нажмите **Create Web Service**.

После завершения деплоя вы получите ссылку вида: `https://desktop-tutorial.onrender.com`.

## ⚙️ Настройка в Janitor AI

1. Откройте **API Settings** -> выберите **OpenAI**.
2. В поле **Reverse Proxy URL** вставьте: `https://ваша-ссылка.onrender.com/v1` (обязательно с `/v1`).
3. В поле **OpenAI Key** введите ваш `PROXY_API_KEY`.
4. В поле **Model** впишите: `gemini-2.0-flash` (или другую модель).
5. Наслаждайтесь общением!

## 🔐 Безопасность (Важно!)
Никогда не фиксируйте изменения в файле `.env` и не загружайте его на GitHub. Все ключи должны находиться только в настройках Render. Файл `.gitignore` в этом репозитории настроен на игнорирование `.env`.

---
*Разработано для использования с Janitor AI и другими ролевыми фронтендами.
Если вы хотите чтобы я добавил фичи по типу распознования фото или количество потраченных токенов маякните мне в [telegram](https://t.me/thenurb)*