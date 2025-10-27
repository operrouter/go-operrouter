package main

import (
	"context"
	"fmt"
	"log"

	"github.com/operrouter/go-operrouter/operrouter"
)

func main() {
	// Create FFI client (adjust path to your liboperrouter_core_ffi.so)
	client, err := operrouter.NewFFI("/usr/local/lib/liboperrouter_core_ffi.so")
	if err != nil {
		log.Fatalf("Failed to create FFI client: %v", err)
	}
	ctx := context.Background()

	// 1. Create an LLM
	config := map[string]interface{}{
		"provider": "openai",
		"model":    "gpt-3.5-turbo",
		"api_key":  "sk-your-api-key-here",
	}

	createResp, err := client.CreateLLM(ctx, "my_llm", config)
	if err != nil {
		log.Fatalf("CreateLLM error: %v", err)
	}
	fmt.Printf("CreateLLM: Success=%v, Message=%s\n", createResp.Success, createResp.Message)

	// 2. Generate text
	generateResp, err := client.GenerateLLM(ctx, "my_llm", "Tell me a joke about programming")
	if err != nil {
		log.Fatalf("GenerateLLM error: %v", err)
	}
	fmt.Printf("GenerateLLM: Success=%v\n", generateResp.Success)
	fmt.Printf("Generated text: %s\n", generateResp.Text)

	// 3. Chat conversation
	messages := []map[string]interface{}{
		{"role": "system", "content": "You are a helpful assistant."},
		{"role": "user", "content": "What is the capital of France?"},
	}
	chatResp, err := client.ChatLLM(ctx, "my_llm", messages)
	if err != nil {
		log.Fatalf("ChatLLM error: %v", err)
	}
	fmt.Printf("ChatLLM: Success=%v\n", chatResp.Success)
	fmt.Printf("Assistant response: %s\n", chatResp.Text)

	// 4. Generate embeddings
	embeddingResp, err := client.EmbeddingLLM(ctx, "my_llm", "Hello world")
	if err != nil {
		log.Fatalf("EmbeddingLLM error: %v", err)
	}
	fmt.Printf("EmbeddingLLM: Success=%v\n", embeddingResp.Success)
	fmt.Printf("Embedding (first 5 values): %v\n", embeddingResp.Embedding[:min(5, len(embeddingResp.Embedding))])

	// 5. Ping to check health
	pingResp, err := client.PingLLM(ctx, "my_llm")
	if err != nil {
		log.Fatalf("PingLLM error: %v", err)
	}
	fmt.Printf("PingLLM: Success=%v, Message=%s\n", pingResp.Success, pingResp.Message)

	// 6. Close the LLM
	closeResp, err := client.CloseLLM(ctx, "my_llm")
	if err != nil {
		log.Fatalf("CloseLLM error: %v", err)
	}
	fmt.Printf("CloseLLM: Success=%v, Message=%s\n", closeResp.Success, closeResp.Message)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
