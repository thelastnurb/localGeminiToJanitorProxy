package converter

import "strings"

func ToGemini(req OpenAIRequest) (GeminiRequest, string) {
	var geminiReq GeminiRequest
	var systemTexts []string

	for _, msg := range req.Messages {
		if msg.Role == "system" {
			systemTexts = append(systemTexts, msg.Content)
			continue
		}

		role := "user"
		if msg.Role == "assistant" {
			role = "model"
		}

		geminiReq.Contents = append(geminiReq.Contents, GeminiContent{
			Role: role,
			Parts: []GeminiPart{
				{Text: msg.Content},
			},
		})
	}

	if len(systemTexts) > 0 {
		geminiReq.SystemInstruction = &GeminiInstruction{
			Parts: []GeminiPart{
				{Text: strings.Join(systemTexts, "\n")},
			},
		}
	}

	geminiReq.GenerationConfig = GeminiGenConfig{
		MaxOutputTokens: req.MaxTokens,
		Temperature:     req.Temperature,
	}

	// Очистка имени модели
	modelName := strings.ReplaceAll(req.Model, "models/", "")

	return geminiReq, modelName
}