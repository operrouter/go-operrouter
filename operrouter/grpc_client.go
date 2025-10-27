package operrouter

import (
	"context"
	"fmt"
	"time"

	pb "github.com/operrouter/go-operrouter/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient implements the Client interface using gRPC
type GRPCClient struct {
	conn    *grpc.ClientConn
	service pb.OperRouterClient // Changed from OperRouterServiceClient
	timeout time.Duration
}

// NewGRPC creates a new gRPC client
// Example: client, err := operrouter.NewGRPC("localhost:50051")
func NewGRPC(address string, opts ...ClientOption) (*GRPCClient, error) {
	return NewGRPCWithOptions(address, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, opts...)
}

// NewGRPCWithOptions creates a new gRPC client with custom dial options
func NewGRPCWithOptions(address string, dialOpts []grpc.DialOption, clientOpts ...ClientOption) (*GRPCClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	client := &GRPCClient{
		conn:    conn,
		service: pb.NewOperRouterClient(conn), // Changed from NewOperRouterServiceClient
		timeout: 5 * time.Second,
	}

	// Apply options
	for _, opt := range clientOpts {
		opt(client)
	}

	return client, nil
}

// Ping checks the service health
func (c *GRPCClient) Ping(ctx context.Context) (*PingResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.PingRequest{}
	resp, err := c.service.Ping(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &PingResponse{
		Status:  resp.Status,
		Version: resp.Version,
	}, nil
}

// ValidateConfig validates operator configuration
func (c *GRPCClient) ValidateConfig(ctx context.Context, tomlContent string) (*ValidateConfigResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.ValidateConfigRequest{
		TomlContent: tomlContent,
	}
	resp, err := c.service.ValidateConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("validate config failed: %w", err)
	}

	return &ValidateConfigResponse{
		Valid:  resp.Valid,
		Errors: resp.Errors,
	}, nil
}

// LoadConfig loads operator configuration from file
func (c *GRPCClient) LoadConfig(ctx context.Context, configPath string) (*LoadConfigResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.LoadConfigRequest{
		ConfigPath: configPath,
	}
	resp, err := c.service.LoadConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("load config failed: %w", err)
	}

	return &LoadConfigResponse{
		Success:      resp.Success,
		OperatorName: resp.OperatorName,
		Error:        resp.Error,
	}, nil
}

// GetMetadata retrieves operator metadata
func (c *GRPCClient) GetMetadata(ctx context.Context) (*MetadataResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.GetMetadataRequest{}
	resp, err := c.service.GetMetadata(ctx, req)
	if err != nil {
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

// Close closes the gRPC connection
func (c *GRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Helper function to map driver string to DataSourceType enum
func mapDataSourceType(driver string) pb.DataSourceType {
	switch driver {
	case "postgres", "postgresql":
		return pb.DataSourceType_DATA_SOURCE_TYPE_POSTGRESQL
	case "mysql":
		return pb.DataSourceType_DATA_SOURCE_TYPE_MYSQL
	case "redis":
		return pb.DataSourceType_DATA_SOURCE_TYPE_REDIS
	case "mongodb", "mongo":
		return pb.DataSourceType_DATA_SOURCE_TYPE_MONGODB
	case "kafka":
		return pb.DataSourceType_DATA_SOURCE_TYPE_KAFKA
	default:
		return pb.DataSourceType_DATA_SOURCE_TYPE_UNSPECIFIED
	}
}

// Helper function to map provider string to LLMProvider enum
func mapLLMProvider(provider string) pb.LLMProvider {
	switch provider {
	case "openai":
		return pb.LLMProvider_LLM_PROVIDER_OPENAI
	case "ollama":
		return pb.LLMProvider_LLM_PROVIDER_OLLAMA
	case "anthropic", "claude":
		return pb.LLMProvider_LLM_PROVIDER_ANTHROPIC
	case "local":
		return pb.LLMProvider_LLM_PROVIDER_LOCAL
	default:
		return pb.LLMProvider_LLM_PROVIDER_UNSPECIFIED
	}
}

// Helper function to map role string to MessageRole enum
func mapMessageRole(role string) pb.MessageRole {
	switch role {
	case "system":
		return pb.MessageRole_MESSAGE_ROLE_SYSTEM
	case "user":
		return pb.MessageRole_MESSAGE_ROLE_USER
	case "assistant":
		return pb.MessageRole_MESSAGE_ROLE_ASSISTANT
	default:
		return pb.MessageRole_MESSAGE_ROLE_UNSPECIFIED
	}
}

// DataSource operations

// CreateDataSource creates a new DataSource connection
func (c *GRPCClient) CreateDataSource(ctx context.Context, name string, config map[string]interface{}) (*DataSourceResponse, error) {
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

	resp, err := c.service.CreateDataSource(ctx, req)
	if err != nil {
		return nil, err
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// QueryDataSource executes a read query on a DataSource
func (c *GRPCClient) QueryDataSource(ctx context.Context, name string, query string) (*DataSourceQueryResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.QueryDataSourceRequest{
		Name:  name,
		Query: query,
	}
	resp, err := c.service.QueryDataSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("query datasource failed: %w", err)
	}

	// Convert proto rows to map format
	rows := make([]map[string]interface{}, 0)
	for _, protoRow := range resp.Rows {
		rowMap := make(map[string]interface{})
		for key, value := range protoRow.Columns {
			rowMap[key] = value
		}
		rows = append(rows, rowMap)
	}

	return &DataSourceQueryResponse{
		Success: resp.Success,
		Rows:    rows,
		Message: resp.Error,
	}, nil
}

// ExecuteDataSource executes a write operation on a DataSource
func (c *GRPCClient) ExecuteDataSource(ctx context.Context, name string, query string) (*DataSourceResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.ExecuteDataSourceRequest{
		Name:  name,
		Query: query,
	}
	resp, err := c.service.ExecuteDataSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("execute datasource failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// InsertDataSource inserts data into a DataSource
func (c *GRPCClient) InsertDataSource(ctx context.Context, name string, data map[string]interface{}) (*DataSourceResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

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
	resp, err := c.service.InsertDataSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("insert datasource failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// PingDataSource checks if a DataSource is alive
func (c *GRPCClient) PingDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.PingDataSourceRequest{
		Name: name,
	}
	resp, err := c.service.PingDataSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ping datasource failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Healthy, // PingDataSourceResponse uses Healthy not Success
		Message: resp.Error,
	}, nil
}

// CloseDataSource closes a DataSource connection
func (c *GRPCClient) CloseDataSource(ctx context.Context, name string) (*DataSourceResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.CloseDataSourceRequest{
		Name: name,
	}
	resp, err := c.service.CloseDataSource(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("close datasource failed: %w", err)
	}

	return &DataSourceResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// LLM operations

// CreateLLM creates a new LLM client
func (c *GRPCClient) CreateLLM(ctx context.Context, name string, config map[string]interface{}) (*LLMResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	// Convert config map to proto format
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
	resp, err := c.service.CreateLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create llm failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// GenerateLLM generates text from a prompt
func (c *GRPCClient) GenerateLLM(ctx context.Context, name string, prompt string) (*LLMGenerateResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.GenerateLLMRequest{
		Name:   name,
		Prompt: prompt,
	}
	resp, err := c.service.GenerateLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("generate llm failed: %w", err)
	}

	return &LLMGenerateResponse{
		Success: resp.Success,
		Text:    resp.Text,
		Message: resp.Error,
	}, nil
}

// ChatLLM performs a chat conversation with message history
func (c *GRPCClient) ChatLLM(ctx context.Context, name string, messages []map[string]interface{}) (*LLMGenerateResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	// Convert messages to proto format
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
	resp, err := c.service.ChatLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("chat llm failed: %w", err)
	}

	return &LLMGenerateResponse{
		Success: resp.Success,
		Text:    resp.Text,
		Message: resp.Error,
	}, nil
}

// EmbeddingLLM generates embeddings for text
func (c *GRPCClient) EmbeddingLLM(ctx context.Context, name string, text string) (*LLMEmbeddingResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.EmbeddingLLMRequest{
		Name: name,
		Text: text,
	}
	resp, err := c.service.EmbeddingLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("embedding llm failed: %w", err)
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
func (c *GRPCClient) PingLLM(ctx context.Context, name string) (*LLMResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.PingLLMRequest{
		Name: name,
	}
	resp, err := c.service.PingLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ping llm failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Healthy, // PingLLMResponse uses Healthy not Success
		Message: resp.Error,
	}, nil
}

// CloseLLM closes an LLM client
func (c *GRPCClient) CloseLLM(ctx context.Context, name string) (*LLMResponse, error) {
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), c.timeout)
		defer cancel()
	}

	req := &pb.CloseLLMRequest{
		Name: name,
	}
	resp, err := c.service.CloseLLM(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("close llm failed: %w", err)
	}

	return &LLMResponse{
		Success: resp.Success,
		Message: resp.Error,
	}, nil
}

// Ensure GRPCClient implements Client interface
var _ Client = (*GRPCClient)(nil)
