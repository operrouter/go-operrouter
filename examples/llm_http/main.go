package main

import (
	"context"
	"fmt"
	"log"

	"github.com/operrouter/go-operrouter/operrouter"
)

func main() {
	// Create HTTP client
	client := operrouter.NewHTTP("http://localhost:8080")
	defer client.Close()

	ctx := context.Background()

	// Example 1: Create an OpenAI LLM client
	fmt.Println("=== Creating OpenAI LLM Client ===")
	createResp, err := client.CreateLLM(ctx, "my_llm", map[string]interface{}{
		"provider":    "openai",
		"api_key":     "sk-your-api-key-here",
		"model":       "gpt-3.5-turbo",
		"temperature": 0.7,
		"max_tokens":  1000,
	})
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
	fmt.Printf("✅ Create Success: %v, Message: %s\n", createResp.Success, createResp.Message)

	// Example 2: Generate text from a prompt
	fmt.Println("\n=== Generating Text ===")
	genResp, err := client.GenerateLLM(ctx, "my_llm", "Explain quantum computing in one sentence.")
	if err != nil {
		log.Fatalf("Failed to generate: %v", err)
	}
	fmt.Printf("✅ Generate Success: %v\n", genResp.Success)
	fmt.Printf("   Text: %s\n", genResp.Text)

	// Example 3: Chat with message history
	fmt.Println("\n=== Chat Conversation ===")
	messages := []map[string]interface{}{
		{
			"role":    "system",
			"content": "You are a helpful assistant.",
		},
		{
			"role":    "user",
			"content": "What is the capital of France?",
		},
	}
	chatResp, err := client.ChatLLM(ctx, "my_llm", messages)
	if err != nil {
		log.Fatalf("Failed to chat: %v", err)
	}
	fmt.Printf("✅ Chat Success: %v\n", chatResp.Success)
	fmt.Printf("   Response: %s\n", chatResp.Text)

	// Example 4: Generate embeddings
	fmt.Println("\n=== Generating Embeddings ===")
	embResp, err := client.EmbeddingLLM(ctx, "my_llm", "Hello, world!")
	if err != nil {
		log.Fatalf("Failed to generate embedding: %v", err)
	}
	fmt.Printf("✅ Embedding Success: %v\n", embResp.Success)
	fmt.Printf("   Embedding dimensions: %d\n", len(embResp.Embedding))
	if len(embResp.Embedding) > 0 {
		fmt.Printf("   First 5 values: %v\n", embResp.Embedding[:min(5, len(embResp.Embedding))])
	}

	// Example 5: Ping the LLM client
	fmt.Println("\n=== Pinging LLM Client ===")
	pingResp, err := client.PingLLM(ctx, "my_llm")
	if err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}
	fmt.Printf("✅ Ping Success: %v, Message: %s\n", pingResp.Success, pingResp.Message)

	// Example 6: Close the LLM client
	fmt.Println("\n=== Closing LLM Client ===")
	closeResp, err := client.CloseLLM(ctx, "my_llm")
	if err != nil {
		log.Fatalf("Failed to close: %v", err)
	}
	fmt.Printf("✅ Close Success: %v, Message: %s\n", closeResp.Success, closeResp.Message)

	fmt.Println("\n🎉 All LLM operations completed successfully!")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
