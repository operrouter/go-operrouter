package operrouter

import (
	"context"
)

// Client is the unified interface for all transport backends (HTTP, gRPC, FFI)
type Client interface {
	// Ping checks the service health and returns version information
	Ping(ctx context.Context) (*PingResponse, error)

	// ValidateConfig validates an operator configuration in TOML format
	ValidateConfig(ctx context.Context, tomlContent string) (*ValidateConfigResponse, error)

	// LoadConfig loads an operator configuration from a file path
	LoadConfig(ctx context.Context, configPath string) (*LoadConfigResponse, error)

	// GetMetadata retrieves operator metadata
	GetMetadata(ctx context.Context) (*MetadataResponse, error)

	// DataSource operations

	// CreateDataSource creates a new DataSource connection
	CreateDataSource(ctx context.Context, name string, config map[string]interface{}) (*DataSourceResponse, error)

	// QueryDataSource executes a read query on a DataSource
	QueryDataSource(ctx context.Context, name string, query string) (*DataSourceQueryResponse, error)

	// ExecuteDataSource executes a write operation on a DataSource
	ExecuteDataSource(ctx context.Context, name string, query string) (*DataSourceResponse, error)

	// InsertDataSource inserts data into a DataSource
	InsertDataSource(ctx context.Context, name string, data map[string]interface{}) (*DataSourceResponse, error)

	// PingDataSource checks if a DataSource is alive
	PingDataSource(ctx context.Context, name string) (*DataSourceResponse, error)

	// CloseDataSource closes a DataSource connection
	CloseDataSource(ctx context.Context, name string) (*DataSourceResponse, error)

	// LLM operations

	// CreateLLM creates a new LLM client
	CreateLLM(ctx context.Context, name string, config map[string]interface{}) (*LLMResponse, error)

	// GenerateLLM generates text from a prompt
	GenerateLLM(ctx context.Context, name string, prompt string) (*LLMGenerateResponse, error)

	// ChatLLM performs a chat conversation with message history
	ChatLLM(ctx context.Context, name string, messages []map[string]interface{}) (*LLMGenerateResponse, error)

	// EmbeddingLLM generates embeddings for text
	EmbeddingLLM(ctx context.Context, name string, text string) (*LLMEmbeddingResponse, error)

	// PingLLM checks if an LLM client is alive
	PingLLM(ctx context.Context, name string) (*LLMResponse, error)

	// CloseLLM closes an LLM client
	CloseLLM(ctx context.Context, name string) (*LLMResponse, error)

	// Close releases resources associated with the client
	Close() error
}

// Common response types

type PingResponse struct {
	Status  string
	Version string
}

type ValidateConfigResponse struct {
	Valid  bool
	Errors []string
}

type LoadConfigResponse struct {
	Success      bool
	OperatorName string
	Error        string
}

type MetadataResponse struct {
	Name        string
	Version     string
	Description string
}

// DataSource response types

type DataSourceResponse struct {
	Success bool
	Message string
}

type DataSourceQueryResponse struct {
	Success bool
	Rows    []map[string]interface{}
	Message string
}

// LLM response types

type LLMResponse struct {
	Success bool
	Message string
}

type LLMGenerateResponse struct {
	Success bool
	Text    string
	Message string
}

type LLMEmbeddingResponse struct {
	Success   bool
	Embedding []float64
	Message   string
}

// ClientOption configures a client
type ClientOption func(interface{})

// WithTimeout sets the operation timeout
func WithTimeout(timeout int) ClientOption {
	return func(c interface{}) {
		// Implementation varies by backend
	}
}
