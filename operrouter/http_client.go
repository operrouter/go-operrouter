package operrouter

import (
	bytes "bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HTTPClient implements the Client interface using HTTP JSON-RPC
type HTTPClient struct {
	BaseURL string
	HTTP    *http.Client
	timeout time.Duration
}

type jsonrpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id"`
}

type jsonrpcResponse struct {
	JSONRPC string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID int `json:"id"`
}

// New creates a new HTTP client (backward compatible)
// Deprecated: Use NewHTTP for clarity
func New(baseURL string) *HTTPClient {
	return NewHTTP(baseURL)
}

// NewHTTP creates a new HTTP JSON-RPC client
// Example: client := operrouter.NewHTTP("http://localhost:8080")
func NewHTTP(baseURL string, opts ...ClientOption) *HTTPClient {
	client := &HTTPClient{
		BaseURL: baseURL,
		HTTP:    &http.Client{Timeout: 10 * time.Second},
		timeout: 5 * time.Second,
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// callJSONRPC makes a JSON-RPC call
func (c *HTTPClient) callJSONRPC(method string, params interface{}, result interface{}) error {
	req := jsonrpcRequest{JSONRPC: "2.0", Method: method, Params: params, ID: 1}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.HTTP.Post(c.BaseURL+"/jsonrpc", "application/json", bytes.NewReader(reqBytes))
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var rpcResp jsonrpcResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if rpcResp.Error != nil {
		return &httpError{code: rpcResp.Error.Code, msg: rpcResp.Error.Message}
	}

	if rpcResp.Result != nil && result != nil {
		if err := json.Unmarshal(*rpcResp.Result, result); err != nil {
			return fmt.Errorf("failed to unmarshal result: %w", err)
		}
	}

	return nil
}

// Ping checks the service health
func (c *HTTPClient) Ping(ctx context.Context) (*PingResponse, error) {
	var result map[string]string
	if err := c.callJSONRPC("ping", nil, &result); err != nil {
		return nil, err
	}

	return &PingResponse{
		Status:  result["status"],
		Version: result["version"],
	}, nil
}

// ValidateConfig validates operator configuration
func (c *HTTPClient) ValidateConfig(ctx context.Context, tomlContent string) (*ValidateConfigResponse, error) {
	params := map[string]interface{}{
		"config_toml": tomlContent,
	}

	var result struct {
		Valid  bool     `json:"valid"`
		Errors []string `json:"errors"`
	}

	if err := c.callJSONRPC("validate_config", params, &result); err != nil {
		return nil, err
	}

	return &ValidateConfigResponse{
		Valid:  result.Valid,
		Errors: result.Errors,
	}, nil
}

// LoadConfig loads operator configuration from file
func (c *HTTPClient) LoadConfig(ctx context.Context, configPath string) (*LoadConfigResponse, error) {
	params := map[string]interface{}{
		"config_path": configPath,
	}

	var result struct {
		Success      bool   `json:"success"`
		Message      string `json:"message"`
		OperatorName string `json:"operator_name"`
	}

	if err := c.callJSONRPC("load_config", params, &result); err != nil {
		return nil, err
	}

	return &LoadConfigResponse{
		Success:      result.Success,
		OperatorName: result.OperatorName,
		Error:        result.Message,
	}, nil
}

// GetMetadata retrieves operator metadata
func (c *HTTPClient) GetMetadata(ctx context.Context) (*MetadataResponse, error) {
	var result struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
	}

	if err := c.callJSONRPC("get_metadata", nil, &result); err != nil {
		return nil, err
	}

	return &MetadataResponse{
		Name:        result.Name,
		Version:     result.Version,
		Description: result.Description,
	}, nil
}

// Close releases resources (no-op for HTTP client)
func (c *HTTPClient) Close() error {
	return nil
}

// DataSource operations

// CreateDataSource creates a new DataSource connection
func (c *HTTPClient) CreateDataSource(ctx context.Context, name string, config map[string]interface{}) (*DataSourceResponse, error) {
	params := map[string]interface{}{
		"name":   name,
		"config": config,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("datasource.create", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// QueryDataSource executes a read query on a DataSource
func (c *HTTPClient) QueryDataSource(ctx context.Context, name string, query string) (*DataSourceQueryResponse, error) {
	params := map[string]interface{}{
		"name":  name,
		"query": query,
	}

	var result struct {
		Success bool                     `json:"success"`
		Rows    []map[string]interface{} `json:"rows"`
		Message string                   `json:"message"`
	}

	if err := c.callJSONRPC("datasource.query", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceQueryResponse{
		Success: result.Success,
		Rows:    result.Rows,
		Message: result.Message,
	}, nil
}

// ExecuteDataSource executes a write operation on a DataSource
func (c *HTTPClient) ExecuteDataSource(ctx context.Context, name string, query string) (*DataSourceResponse, error) {
	params := map[string]interface{}{
		"name":  name,
		"query": query,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("datasource.execute", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// InsertDataSource inserts data into a DataSource
func (c *HTTPClient) InsertDataSource(ctx context.Context, name string, data map[string]interface{}) (*DataSourceResponse, error) {
	params := map[string]interface{}{
		"name": name,
		"data": data,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("datasource.insert", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// PingDataSource checks if a DataSource is alive
func (c *HTTPClient) PingDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	params := map[string]interface{}{
		"name": name,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("datasource.ping", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// CloseDataSource closes a DataSource connection
func (c *HTTPClient) CloseDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	params := map[string]interface{}{
		"name": name,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("datasource.close", params, &result); err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// LLM operations

// CreateLLM creates a new LLM client
func (c *HTTPClient) CreateLLM(ctx context.Context, name string, config map[string]interface{}) (*LLMResponse, error) {
	params := map[string]interface{}{
		"name":   name,
		"config": config,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("llm.create", params, &result); err != nil {
		return nil, err
	}

	return &LLMResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// GenerateLLM generates text from a prompt
func (c *HTTPClient) GenerateLLM(ctx context.Context, name string, prompt string) (*LLMGenerateResponse, error) {
	params := map[string]interface{}{
		"name":   name,
		"prompt": prompt,
	}

	var result struct {
		Success bool   `json:"success"`
		Text    string `json:"text"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("llm.generate", params, &result); err != nil {
		return nil, err
	}

	return &LLMGenerateResponse{
		Success: result.Success,
		Text:    result.Text,
		Message: result.Message,
	}, nil
}

// ChatLLM performs a chat conversation with message history
func (c *HTTPClient) ChatLLM(ctx context.Context, name string, messages []map[string]interface{}) (*LLMGenerateResponse, error) {
	params := map[string]interface{}{
		"name":     name,
		"messages": messages,
	}

	var result struct {
		Success bool   `json:"success"`
		Text    string `json:"text"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("llm.chat", params, &result); err != nil {
		return nil, err
	}

	return &LLMGenerateResponse{
		Success: result.Success,
		Text:    result.Text,
		Message: result.Message,
	}, nil
}

// EmbeddingLLM generates embeddings for text
func (c *HTTPClient) EmbeddingLLM(ctx context.Context, name string, text string) (*LLMEmbeddingResponse, error) {
	params := map[string]interface{}{
		"name": name,
		"text": text,
	}

	var result struct {
		Success   bool      `json:"success"`
		Embedding []float64 `json:"embedding"`
		Message   string    `json:"message"`
	}

	if err := c.callJSONRPC("llm.embedding", params, &result); err != nil {
		return nil, err
	}

	return &LLMEmbeddingResponse{
		Success:   result.Success,
		Embedding: result.Embedding,
		Message:   result.Message,
	}, nil
}

// PingLLM checks if an LLM client is alive
func (c *HTTPClient) PingLLM(ctx context.Context, name string) (*LLMResponse, error) {
	params := map[string]interface{}{
		"name": name,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("llm.ping", params, &result); err != nil {
		return nil, err
	}

	return &LLMResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

// CloseLLM closes an LLM client
func (c *HTTPClient) CloseLLM(ctx context.Context, name string) (*LLMResponse, error) {
	params := map[string]interface{}{
		"name": name,
	}

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	if err := c.callJSONRPC("llm.close", params, &result); err != nil {
		return nil, err
	}

	return &LLMResponse{
		Success: result.Success,
		Message: result.Message,
	}, nil
}

type httpError struct {
	code int
	msg  string
}

func (e *httpError) Error() string { return e.msg }

// Ensure HTTPClient implements Client interface
var _ Client = (*HTTPClient)(nil)
