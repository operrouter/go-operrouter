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

	// Example 1: Create a PostgreSQL DataSource
	fmt.Println("=== Creating PostgreSQL DataSource ===")
	createResp, err := client.CreateDataSource(ctx, "my_postgres", map[string]interface{}{
		"driver":   "postgres",
		"host":     "localhost",
		"port":     5432,
		"database": "testdb",
		"username": "postgres",
		"password": "password",
	})
	if err != nil {
		log.Fatalf("Failed to create datasource: %v", err)
	}
	fmt.Printf("✅ Create Success: %v, Message: %s\n", createResp.Success, createResp.Message)

	// Example 2: Query the DataSource
	fmt.Println("\n=== Querying DataSource ===")
	queryResp, err := client.QueryDataSource(ctx, "my_postgres", "SELECT 1 as id, 'test' as name")
	if err != nil {
		log.Fatalf("Failed to query datasource: %v", err)
	}
	fmt.Printf("✅ Query Success: %v\n", queryResp.Success)
	fmt.Printf("   Rows returned: %d\n", len(queryResp.Rows))
	for i, row := range queryResp.Rows {
		fmt.Printf("   Row %d: %v\n", i+1, row)
	}

	// Example 3: Execute a write operation
	fmt.Println("\n=== Executing Write Operation ===")
	execResp, err := client.ExecuteDataSource(ctx, "my_postgres", "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatalf("Failed to execute: %v", err)
	}
	fmt.Printf("✅ Execute Success: %v, Message: %s\n", execResp.Success, execResp.Message)

	// Example 4: Insert data
	fmt.Println("\n=== Inserting Data ===")
	insertResp, err := client.InsertDataSource(ctx, "my_postgres", map[string]interface{}{
		"table": "users",
		"data": map[string]interface{}{
			"name": "Alice",
		},
	})
	if err != nil {
		log.Fatalf("Failed to insert: %v", err)
	}
	fmt.Printf("✅ Insert Success: %v, Message: %s\n", insertResp.Success, insertResp.Message)

	// Example 5: Ping the DataSource
	fmt.Println("\n=== Pinging DataSource ===")
	pingResp, err := client.PingDataSource(ctx, "my_postgres")
	if err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}
	fmt.Printf("✅ Ping Success: %v, Message: %s\n", pingResp.Success, pingResp.Message)

	// Example 6: Close the DataSource
	fmt.Println("\n=== Closing DataSource ===")
	closeResp, err := client.CloseDataSource(ctx, "my_postgres")
	if err != nil {
		log.Fatalf("Failed to close: %v", err)
	}
	fmt.Printf("✅ Close Success: %v, Message: %s\n", closeResp.Success, closeResp.Message)

	fmt.Println("\n🎉 All DataSource operations completed successfully!")
}
