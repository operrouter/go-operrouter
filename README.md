# Go OperRouter SDK

Go client library for OperRouter - operator lifecycle management framework.

## Features

- 🚀 **Three Transport Backends**: HTTP, gRPC, and FFI
- 🔄 **Unified Interface**: Same API across all backends
- 🗄️ **DataSource Support**: Connect to PostgreSQL, MySQL, Redis, MongoDB, Kafka
- 🤖 **LLM Support**: Integrate with OpenAI, Anthropic, Ollama, and local models
- ⚡ **High Performance**: FFI ~140ns, gRPC ~0.5-2ms, HTTP ~1-10ms
- 🛡️ **Type Safe**: Compile-time type checking with protobuf
- 🎯 **Context Support**: Cancellation and timeout handling

## Installation

```bash
go get github.com/operrouter/go-operrouter
```

## Quick Start

### HTTP Backend (Simplest)

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/operrouter/go-operrouter/operrouter"
)

func main() {
    client := operrouter.NewHTTP("http://localhost:8080")
    defer client.Close()
    
    resp, err := client.Ping(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Version: %s\n", resp.Version)
}
```

### DataSource Operations

```go
// Create a PostgreSQL connection
client.CreateDataSource(ctx, "my_db", map[string]interface{}{
    "driver":   "postgres",
    "host":     "localhost",
    "port":     5432,
    "database": "mydb",
})

// Query data
rows, err := client.QueryDataSource(ctx, "my_db", "SELECT * FROM users")

// Execute write operations
client.ExecuteDataSource(ctx, "my_db", "INSERT INTO users (name) VALUES ('Alice')")

// Close connection
client.CloseDataSource(ctx, "my_db")
```

### LLM Operations

```go
// Create an OpenAI LLM client
client.CreateLLM(ctx, "my_llm", map[string]interface{}{
    "provider":  "openai",
    "api_key":   "sk-...",
    "model":     "gpt-4",
})

// Generate text
resp, err := client.GenerateLLM(ctx, "my_llm", "Explain quantum computing")

// Chat conversation
messages := []map[string]interface{}{
    {"role": "system", "content": "You are a helpful assistant"},
    {"role": "user", "content": "What is Rust?"},
}
chatResp, err := client.ChatLLM(ctx, "my_llm", messages)

// Generate embeddings
embResp, err := client.EmbeddingLLM(ctx, "my_llm", "Hello world")

// Close client
client.CloseLLM(ctx, "my_llm")
```

### gRPC Backend (High Performance)

```go
client, err := operrouter.NewGRPC("localhost:50051")
if err != nil {
    log.Fatal(err)
}
defer client.Close()

resp, err := client.Ping(context.Background())
```

### FFI Backend (Fastest)

```go
libPath := "/path/to/liboperrouter_core_ffi.so"
client, err := operrouter.NewFFI(libPath)
if err != nil {
    log.Fatal(err)
}
defer client.Close()

resp, err := client.Ping(context.Background())
```

**Note**: FFI backend requires CGO

## API Reference

### Core Operations

- `Ping(ctx) (*PingResponse, error)` - Health check
- `ValidateConfig(ctx, toml) (*ValidateConfigResponse, error)` - Validate TOML config
- `LoadConfig(ctx, path) (*LoadConfigResponse, error)` - Load config from file
- `GetMetadata(ctx) (*MetadataResponse, error)` - Get operator metadata

### DataSource Operations

- `CreateDataSource(ctx, name, config) (*DataSourceResponse, error)` - Create connection
- `QueryDataSource(ctx, name, query) (*DataSourceQueryResponse, error)` - Execute SELECT query
- `ExecuteDataSource(ctx, name, query) (*DataSourceResponse, error)` - Execute INSERT/UPDATE/DELETE
- `InsertDataSource(ctx, name, data) (*DataSourceResponse, error)` - Insert data
- `PingDataSource(ctx, name) (*DataSourceResponse, error)` - Check connection
- `CloseDataSource(ctx, name) (*DataSourceResponse, error)` - Close connection

### LLM Operations

- `CreateLLM(ctx, name, config) (*LLMResponse, error)` - Create LLM client
- `GenerateLLM(ctx, name, prompt) (*LLMGenerateResponse, error)` - Generate text
- `ChatLLM(ctx, name, messages) (*LLMGenerateResponse, error)` - Chat conversation
- `EmbeddingLLM(ctx, name, text) (*LLMEmbeddingResponse, error)` - Generate embeddings
- `PingLLM(ctx, name) (*LLMResponse, error)` - Check LLM client
- `CloseLLM(ctx, name) (*LLMResponse, error)` - Close LLM client

## Backend Comparison

| Feature | HTTP | gRPC | FFI |
|---------|------|------|-----|
| **Latency** | 1-10ms | 0.5-2ms | ~140ns |
| **Network** | Required | Required | Not required |
| **Serialization** | JSON | Protobuf | Protobuf |
| **DataSource** | ✅ Full | ✅ Full | ✅ Full |
| **LLM** | ✅ Full | ✅ Full | ✅ Full |
| **Best For** | Web services | Microservices | Embedded/local |

## Supported DataSources

- **PostgreSQL** - Full SQL support
- **MySQL** - Full SQL support
- **Redis** - Key-value operations
- **MongoDB** - Document operations
- **Kafka** - Message streaming
- **HTTP** - REST API calls

## Supported LLM Providers

- **OpenAI** - GPT-3.5, GPT-4, Embeddings
- **Anthropic** - Claude models
- **Ollama** - Local models
- **Local** - Custom local models

## Examples

See `examples/` directory for complete examples:

- `datasource_http.go` - HTTP DataSource operations
- `llm_http.go` - HTTP LLM operations
- `ping/` - Basic health check examples

## Documentation

- [Complete Implementation Guide](./GO_SDK_IMPLEMENTATION.md)
- [Implementation Status](./IMPLEMENTATION_STATUS.md)

## License

MIT OR Apache-2.0
