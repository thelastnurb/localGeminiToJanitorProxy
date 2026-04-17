package gemini

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"proxy/converter"
)

const baseURL = "https://generativelanguage.googleapis.com/v1beta/models/"

func GenerateContent(ctx context.Context, model string, req converter.GeminiRequest) (converter.GeminiResponse, error) {
	url := fmt.Sprintf("%s%s:generateContent?key=%s", baseURL, model, os.Getenv("GEMINI_API_KEY"))
	
	body, _ := json.Marshal(req)
	
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return converter.GeminiResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return converter.GeminiResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return converter.GeminiResponse{}, parseError(resp)
	}

	var geminiResp converter.GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return converter.GeminiResponse{}, err
	}

	return geminiResp, nil
}

func StreamContent(ctx context.Context, model string, req converter.GeminiRequest, onChunk func(converter.GeminiResponse)) error {
	url := fmt.Sprintf("%s%s:streamGenerateContent?alt=sse&key=%s", baseURL, model, os.Getenv("GEMINI_API_KEY"))
	
	body, _ := json.Marshal(req)
	
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return parseError(resp)
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			var geminiResp converter.GeminiResponse
			if err := json.Unmarshal([]byte(data), &geminiResp); err == nil {
				onChunk(geminiResp)
			}
		}
	}

	return scanner.Err()
}

func parseError(resp *http.Response) error {
	var errResp struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	json.NewDecoder(resp.Body).Decode(&errResp)
	if errResp.Error.Message != "" {
		return fmt.Errorf("gemini error: %s", errResp.Error.Message)
	}
	return fmt.Errorf("gemini API returned status: %d", resp.StatusCode)
}