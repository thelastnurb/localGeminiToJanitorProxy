package converter

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

func ToOpenAI(resp GeminiResponse, modelName string) OpenAIResponse {
	text := ""
	finishReason := "stop"

	if len(resp.Candidates) > 0 {
		cand := resp.Candidates[0]
		if len(cand.Content.Parts) > 0 {
			text = cand.Content.Parts[0].Text
		}
		if cand.FinishReason != "" && cand.FinishReason != "STOP" {
			finishReason = strings.ToLower(cand.FinishReason)
		}
	}

	return OpenAIResponse{
		ID:      generateID(),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []OpenAIChoice{
			{
				Index: 0,
				Message: OpenAIMessage{
					Role:    "assistant",
					Content: text,
				},
				FinishReason: finishReason,
			},
		},
		Usage: OpenAIUsage{
			PromptTokens:     resp.UsageMetadata.PromptTokenCount,
			CompletionTokens: resp.UsageMetadata.CandidatesTokenCount,
			TotalTokens:      resp.UsageMetadata.PromptTokenCount + resp.UsageMetadata.CandidatesTokenCount,
		},
	}
}

func ToOpenAIStream(resp GeminiResponse, modelName string) OpenAIStreamResponse {
	text := ""
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		text = resp.Candidates[0].Content.Parts[0].Text
	}

	return OpenAIStreamResponse{
		ID:      generateID(),
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []OpenAIStreamChoice{
			{
				Index: 0,
				Delta: OpenAIStreamMsg{
					Content: text,
				},
				FinishReason: nil,
			},
		},
	}
}

func ToOpenAIStreamDone(modelName string) OpenAIStreamResponse {
	stopReason := "stop"
	return OpenAIStreamResponse{
		ID:      generateID(),
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []OpenAIStreamChoice{
			{
				Index: 0,
				Delta: OpenAIStreamMsg{},
				FinishReason: &stopReason,
			},
		},
	}
}

func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("chatcmpl-%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}