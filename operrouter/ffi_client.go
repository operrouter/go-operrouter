//go:build cgo
// +build cgo

package operrouter

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include <stdint.h>
#include <dlfcn.h>

typedef struct {
    uint8_t* data;
    size_t len;
} ProtoBuffer;

typedef ProtoBuffer (*ping_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*validate_config_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*load_config_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*get_metadata_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_create_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_query_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_execute_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_insert_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_ping_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*datasource_close_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_create_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_generate_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_chat_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_embedding_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_ping_proto_fn)(const uint8_t*, size_t);
typedef ProtoBuffer (*llm_close_proto_fn)(const uint8_t*, size_t);
typedef void (*proto_buffer_free_fn)(ProtoBuffer);

// Helper functions to call FFI with dynamic loading
static ProtoBuffer call_ping_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    ping_proto_fn fn = (ping_proto_fn)dlsym(handle, "ping_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_validate_config_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    validate_config_proto_fn fn = (validate_config_proto_fn)dlsym(handle, "validate_config_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_load_config_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    load_config_proto_fn fn = (load_config_proto_fn)dlsym(handle, "load_config_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_get_metadata_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    get_metadata_proto_fn fn = (get_metadata_proto_fn)dlsym(handle, "get_metadata_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static void call_proto_buffer_free(void* handle, ProtoBuffer buf) {
    proto_buffer_free_fn fn = (proto_buffer_free_fn)dlsym(handle, "proto_buffer_free");
    if (fn) fn(buf);
}

// DataSource helper functions
static ProtoBuffer call_datasource_create_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_create_proto_fn fn = (datasource_create_proto_fn)dlsym(handle, "datasource_create_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_datasource_query_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_query_proto_fn fn = (datasource_query_proto_fn)dlsym(handle, "datasource_query_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_datasource_execute_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_execute_proto_fn fn = (datasource_execute_proto_fn)dlsym(handle, "datasource_execute_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_datasource_insert_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_insert_proto_fn fn = (datasource_insert_proto_fn)dlsym(handle, "datasource_insert_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_datasource_ping_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_ping_proto_fn fn = (datasource_ping_proto_fn)dlsym(handle, "datasource_ping_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_datasource_close_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    datasource_close_proto_fn fn = (datasource_close_proto_fn)dlsym(handle, "datasource_close_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

// LLM helper functions
static ProtoBuffer call_llm_create_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_create_proto_fn fn = (llm_create_proto_fn)dlsym(handle, "llm_create_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_llm_generate_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_generate_proto_fn fn = (llm_generate_proto_fn)dlsym(handle, "llm_generate_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_llm_chat_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_chat_proto_fn fn = (llm_chat_proto_fn)dlsym(handle, "llm_chat_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_llm_embedding_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_embedding_proto_fn fn = (llm_embedding_proto_fn)dlsym(handle, "llm_embedding_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_llm_ping_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_ping_proto_fn fn = (llm_ping_proto_fn)dlsym(handle, "llm_ping_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}

static ProtoBuffer call_llm_close_proto(void* handle, const uint8_t* input_ptr, size_t input_len) {
    llm_close_proto_fn fn = (llm_close_proto_fn)dlsym(handle, "llm_close_proto");
    if (!fn) return (ProtoBuffer){NULL, 0};
    return fn(input_ptr, input_len);
}
*/
import "C"
import (
	"context"
	"fmt"
	"unsafe"

	pb "github.com/operrouter/go-operrouter/gen/proto"
	"google.golang.org/protobuf/proto"
)

// FFIClient implements the Client interface using FFI (cgo)
type FFIClient struct {
	handle unsafe.Pointer
	path   string
}

// NewFFI creates a new FFI client by loading the shared library
// Example: client, err := operrouter.NewFFI("/path/to/liboperrouter_core_ffi.so")
func NewFFI(libraryPath string, opts ...ClientOption) (*FFIClient, error) {
	cPath := C.CString(libraryPath)
	defer C.free(unsafe.Pointer(cPath))

	// Load the shared library with RTLD_LAZY | RTLD_LOCAL
	handle := C.dlopen(cPath, C.RTLD_LAZY)
	if handle == nil {
		errStr := C.GoString(C.dlerror())
		return nil, fmt.Errorf("failed to load FFI library %s: %s", libraryPath, errStr)
	}

	client := &FFIClient{
		handle: handle,
		path:   libraryPath,
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// callFFI is a helper to call FFI functions with protobuf marshaling/unmarshaling
func (c *FFIClient) callFFI(
	callFunc func(unsafe.Pointer, *C.uint8_t, C.size_t) C.ProtoBuffer,
	req proto.Message,
	resp proto.Message,
) error {
	// Marshal request
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	var inputPtr *C.uint8_t
	var inputLen C.size_t

	if len(reqBytes) > 0 {
		inputPtr = (*C.uint8_t)(unsafe.Pointer(&reqBytes[0]))
		inputLen = C.size_t(len(reqBytes))
	} else {
		// For empty messages, pass NULL pointer with 0 length
		inputPtr = nil
		inputLen = 0
	}

	// Call FFI function
	outputBuf := callFunc(c.handle, inputPtr, inputLen)

	// Check for null response
	if outputBuf.data == nil && outputBuf.len > 0 {
		return fmt.Errorf("FFI call returned null pointer")
	}

	// Copy response data before freeing
	var respBytes []byte
	if outputBuf.len > 0 && outputBuf.data != nil {
		respBytes = C.GoBytes(unsafe.Pointer(outputBuf.data), C.int(outputBuf.len))
	}

	// Free the buffer
	C.call_proto_buffer_free(c.handle, outputBuf)

	// Unmarshal response
	if len(respBytes) > 0 {
		if err := proto.Unmarshal(respBytes, resp); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

// Ping checks the service health
func (c *FFIClient) Ping(ctx context.Context) (*PingResponse, error) {
	req := &pb.PingRequest{}
	resp := &pb.PingResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_ping_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &PingResponse{
		Status:  resp.Status,
		Version: resp.Version,
	}, nil
}

// ValidateConfig validates operator configuration
func (c *FFIClient) ValidateConfig(ctx context.Context, tomlContent string) (*ValidateConfigResponse, error) {
	req := &pb.ValidateConfigRequest{
		TomlContent: tomlContent,
	}
	resp := &pb.ValidateConfigResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_validate_config_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("validate config failed: %w", err)
	}

	return &ValidateConfigResponse{
		Valid:  resp.Valid,
		Errors: resp.Errors,
	}, nil
}

// LoadConfig loads operator configuration from file
func (c *FFIClient) LoadConfig(ctx context.Context, configPath string) (*LoadConfigResponse, error) {
	req := &pb.LoadConfigRequest{
		ConfigPath: configPath,
	}
	resp := &pb.LoadConfigResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_load_config_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("load config failed: %w", err)
	}

	return &LoadConfigResponse{
		Success:      resp.Success,
		OperatorName: resp.OperatorName,
		Error:        resp.Error,
	}, nil
}

// GetMetadata retrieves operator metadata
func (c *FFIClient) GetMetadata(ctx context.Context) (*MetadataResponse, error) {
	req := &pb.GetMetadataRequest{}
	resp := &pb.GetMetadataResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_get_metadata_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("get metadata failed: %w", err)
	}

	if resp.Metadata == nil {
		return nil, fmt.Errorf("metadata is nil in response")
	}

	return &MetadataResponse{
		Name:        resp.Metadata.Name,
		Version:     resp.Metadata.Version,
		Description: resp.Metadata.Description,
	}, nil
}

// Close closes the FFI library handle
func (c *FFIClient) Close() error {
	if c.handle != nil {
		C.dlclose(c.handle)
		c.handle = nil
	}
	return nil
}

// DataSource operations

// CreateDataSource creates a new DataSource connection
func (c *FFIClient) CreateDataSource(ctx context.Context, name string, config map[string]interface{}) (*DataSourceResponse, error) {
	// Build connection URL from config
	url := ""
	extra := make(map[string]string)

	// Extract driver type
	driver := ""
	if d, ok := config["driver"].(string); ok {
		driver = d
		// Build URL based on driver
		switch driver {
		case "postgres", "postgresql":
			host := config["host"].(string)
			port := int(config["port"].(float64))
			database := config["database"].(string)
			url = fmt.Sprintf("postgres://%s:%d/%s", host, port, database)
			if username, ok := config["username"].(string); ok {
				extra["username"] = username
			}
			if password, ok := config["password"].(string); ok {
				extra["password"] = password
			}
		case "mysql":
			host := config["host"].(string)
			port := int(config["port"].(float64))
			database := config["database"].(string)
			url = fmt.Sprintf("mysql://%s:%d/%s", host, port, database)
		case "redis":
			host := config["host"].(string)
			port := int(config["port"].(float64))
			url = fmt.Sprintf("redis://%s:%d", host, port)
		case "mongodb", "mongo":
			host := config["host"].(string)
			port := int(config["port"].(float64))
			url = fmt.Sprintf("mongodb://%s:%d", host, port)
		case "kafka":
			if brokers, ok := config["brokers"].(string); ok {
				url = brokers
			}
		}
	}

	// Copy remaining config to extra
	for k, v := range config {
		if k != "driver" && k != "host" && k != "port" && k != "database" && k != "username" && k != "password" && k != "brokers" {
			extra[k] = fmt.Sprintf("%v", v)
		}
	}

	req := &pb.CreateDataSourceRequest{
		Name: name,
		Config: &pb.DataSourceConfig{
			Type:  mapDataSourceType(driver),
			Url:   url,
			Extra: extra,
		},
	}
	resp := &pb.CreateDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_create_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource create failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// QueryDataSource executes a read query on a DataSource
func (c *FFIClient) QueryDataSource(ctx context.Context, name string, query string) (*DataSourceQueryResponse, error) {
	req := &pb.QueryDataSourceRequest{
		Name:  name,
		Query: query,
	}
	resp := &pb.QueryDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_query_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource query failed: %w", err)
	}

	// Simplified: Convert proto rows to map format (implement proper parsing based on your needs)
	rows := make([]map[string]interface{}, 0)

	return &DataSourceQueryResponse{
		Success: resp.Success,
		Rows:    rows,
		Message: resp.Error,
	}, nil
}

// ExecuteDataSource executes a write operation on a DataSource
func (c *FFIClient) ExecuteDataSource(ctx context.Context, name string, query string) (*DataSourceResponse, error) {
	req := &pb.ExecuteDataSourceRequest{
		Name:  name,
		Query: query,
	}
	resp := &pb.ExecuteDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_execute_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource execute failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// InsertDataSource inserts data into a DataSource
func (c *FFIClient) InsertDataSource(ctx context.Context, name string, data map[string]interface{}) (*DataSourceResponse, error) {
	// Convert map to proto Row
	columns := make(map[string]*pb.Value)
	for key, val := range data {
		// Simplified: convert to string value
		columns[key] = &pb.Value{
			Value: &pb.Value_StringValue{StringValue: fmt.Sprintf("%v", val)},
		}
	}

	req := &pb.InsertDataSourceRequest{
		Name:  name,
		Table: "", // Table name could be extracted from data or passed separately
		Data: &pb.Row{
			Columns: columns,
		},
	}
	resp := &pb.InsertDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_insert_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource insert failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// PingDataSource checks if a DataSource is alive
func (c *FFIClient) PingDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	req := &pb.PingDataSourceRequest{
		Name: name,
	}
	resp := &pb.PingDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_ping_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource ping failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Healthy, // PingDataSourceResponse uses Healthy not Success
		Message: resp.Error,
	}, nil
}

// CloseDataSource closes a DataSource connection
func (c *FFIClient) CloseDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	req := &pb.CloseDataSourceRequest{
		Name: name,
	}
	resp := &pb.CloseDataSourceResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_datasource_close_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("datasource close failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// LLM operations

// CreateLLM creates a new LLM client
func (c *FFIClient) CreateLLM(ctx context.Context, name string, config map[string]interface{}) (*LLMResponse, error) {
	provider := ""
	if p, ok := config["provider"].(string); ok {
		provider = p
	}

	configProto := &pb.LLMConfig{
		Provider: mapLLMProvider(provider),
	}
	if model, ok := config["model"].(string); ok {
		configProto.Model = model
	}
	if apiKey, ok := config["api_key"].(string); ok {
		configProto.ApiKey = &apiKey
	}

	req := &pb.CreateLLMRequest{
		Name:   name,
		Config: configProto,
	}
	resp := &pb.CreateLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_create_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm create failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// GenerateLLM generates text from a prompt
func (c *FFIClient) GenerateLLM(ctx context.Context, name string, prompt string) (*LLMGenerateResponse, error) {
	req := &pb.GenerateLLMRequest{
		Name:   name,
		Prompt: prompt,
	}
	resp := &pb.GenerateLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_generate_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm generate failed: %w", err)
	}

	return &LLMGenerateResponse{
		Success: resp.Success,
		Text:    resp.Text,
		Message: resp.Error,
	}, nil
}

// ChatLLM performs a chat conversation with message history
func (c *FFIClient) ChatLLM(ctx context.Context, name string, messages []map[string]interface{}) (*LLMGenerateResponse, error) {
	protoMessages := make([]*pb.LLMMessage, len(messages))
	for i, msg := range messages {
		role := ""
		if r, ok := msg["role"].(string); ok {
			role = r
		}
		content := ""
		if c, ok := msg["content"].(string); ok {
			content = c
		}
		protoMessages[i] = &pb.LLMMessage{
			Role:    mapMessageRole(role),
			Content: content,
		}
	}

	req := &pb.ChatLLMRequest{
		Name:     name,
		Messages: protoMessages,
	}
	resp := &pb.ChatLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_chat_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm chat failed: %w", err)
	}

	return &LLMGenerateResponse{
		Success: resp.Success,
		Text:    resp.Text,
		Message: resp.Error,
	}, nil
}

// EmbeddingLLM generates embeddings for text
func (c *FFIClient) EmbeddingLLM(ctx context.Context, name string, text string) (*LLMEmbeddingResponse, error) {
	req := &pb.EmbeddingLLMRequest{
		Name: name,
		Text: text,
	}
	resp := &pb.EmbeddingLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_embedding_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm embedding failed: %w", err)
	}

	// Convert []float32 to []float64
	embedding := make([]float64, len(resp.Embedding))
	for i, v := range resp.Embedding {
		embedding[i] = float64(v)
	}

	return &LLMEmbeddingResponse{
		Success:   resp.Success,
		Embedding: embedding,
		Message:   resp.Error,
	}, nil
}

// PingLLM checks if an LLM client is alive
func (c *FFIClient) PingLLM(ctx context.Context, name string) (*LLMResponse, error) {
	req := &pb.PingLLMRequest{
		Name: name,
	}
	resp := &pb.PingLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_ping_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm ping failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Healthy, // PingLLMResponse uses Healthy not Success
		Message: resp.Error,
	}, nil
}

// CloseLLM closes an LLM client
func (c *FFIClient) CloseLLM(ctx context.Context, name string) (*LLMResponse, error) {
	req := &pb.CloseLLMRequest{
		Name: name,
	}
	resp := &pb.CloseLLMResponse{}

	if err := c.callFFI(func(h unsafe.Pointer, ptr *C.uint8_t, len C.size_t) C.ProtoBuffer {
		return C.call_llm_close_proto(h, ptr, len)
	}, req, resp); err != nil {
		return nil, fmt.Errorf("llm close failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// Ensure FFIClient implements Client interface
var _ Client = (*FFIClient)(nil)
