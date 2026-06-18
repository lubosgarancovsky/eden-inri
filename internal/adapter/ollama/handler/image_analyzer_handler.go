package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lubosgarancovsky/eden-inri/internal/config"
)

type ImageAnalyzerHandler struct{}

func NewImageAnalyzerHandler() *ImageAnalyzerHandler {
	return &ImageAnalyzerHandler{}
}

func (h *ImageAnalyzerHandler) Analyze(prompt string, imagesB64 []string) (string, error) {
	req := map[string]any{
		"stream": false,
		"think":  false,
		"format": "json",
		"model":  config.GlobalConfig.OllamaVisionModel,
		"messages": []map[string]any{
			{
				"role":    "user",
				"content": prompt,
				"images":  imagesB64,
			},
		},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(config.GlobalConfig.OllamaUrl+"/api/chat", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}

	return out.Message.Content, nil
}
