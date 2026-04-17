package handler

import (
	"encoding/json"
	"net/http"

	"proxy/converter"
	"proxy/gemini"
)

func CompletionsHandler(w http.ResponseWriter, r *http.Request) {
	var req converter.OpenAIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendOpenAIError(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	geminiReq, modelName := converter.ToGemini(req)

	if req.Stream {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		err := gemini.StreamContent(r.Context(), modelName, geminiReq, func(resp converter.GeminiResponse) {
			chunk := converter.ToOpenAIStream(resp, req.Model)
			chunkBytes, _ := json.Marshal(chunk)
			w.Write([]byte("data: "))
			w.Write(chunkBytes)
			w.Write([]byte("\n\n"))
			
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		})

		if err != nil {
			// Если стрим еще не начался, можно отправить JSON ошибку
			sendOpenAIError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Завершающий чанк
		doneChunk := converter.ToOpenAIStreamDone(req.Model)
		doneBytes, _ := json.Marshal(doneChunk)
		w.Write([]byte("data: "))
		w.Write(doneBytes)
		w.Write([]byte("\n\n"))
		w.Write([]byte("data: [DONE]\n\n"))
		return
	}

	// Non-stream
	resp, err := gemini.GenerateContent(r.Context(), modelName, geminiReq)
	if err != nil {
		sendOpenAIError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	openAIResp := converter.ToOpenAI(resp, req.Model)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openAIResp)
}

func sendOpenAIError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]any{
			"message": message,
			"type":    "server_error",
			"code":    code,
		},
	})
}