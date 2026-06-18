package ports

type OllamaPort interface {
	Analyze(prompt string, imagesB64 []string) (string, error)
}
