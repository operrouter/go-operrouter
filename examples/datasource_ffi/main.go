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

	// 1. Create a PostgreSQL datasource
	config := map[string]interface{}{
		"driver":   "postgres",
		"host":     "localhost",
		"port":     "5432",
		"database": "testdb",
		"username": "postgres",
		"password": "password",
	}

	createResp, err := client.CreateDataSource(ctx, "my_postgres", config)
	if err != nil {
		log.Fatalf("CreateDataSource error: %v", err)
	}
	fmt.Printf("CreateDataSource: Success=%v, Message=%s\n", createResp.Success, createResp.Message)

	// 2. Query the datasource
	queryResp, err := client.QueryDataSource(ctx, "my_postgres", "SELECT * FROM users LIMIT 5")
	if err != nil {
		log.Fatalf("QueryDataSource error: %v", err)
	}
	fmt.Printf("QueryDataSource: Success=%v, Rows=%d\n", queryResp.Success, len(queryResp.Rows))
	for i, row := range queryResp.Rows {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// 3. Execute DDL/DML
	execResp, err := client.ExecuteDataSource(ctx, "my_postgres", "CREATE TABLE IF NOT EXISTS test_table (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatalf("ExecuteDataSource error: %v", err)
	}
	fmt.Printf("ExecuteDataSource: Success=%v, Message=%s\n", execResp.Success, execResp.Message)

	// 4. Insert data
	insertData := map[string]interface{}{
		"id":   1,
		"name": "Test User",
	}
	insertResp, err := client.InsertDataSource(ctx, "my_postgres", insertData)
	if err != nil {
		log.Fatalf("InsertDataSource error: %v", err)
	}
	fmt.Printf("InsertDataSource: Success=%v, Message=%s\n", insertResp.Success, insertResp.Message)

	// 5. Ping to check health
	pingResp, err := client.PingDataSource(ctx, "my_postgres")
	if err != nil {
		log.Fatalf("PingDataSource error: %v", err)
	}
	fmt.Printf("PingDataSource: Success=%v, Message=%s\n", pingResp.Success, pingResp.Message)

	// 6. Close the datasource
	closeResp, err := client.CloseDataSource(ctx, "my_postgres")
	if err != nil {
		log.Fatalf("CloseDataSource error: %v", err)
	}
	fmt.Printf("CloseDataSource: Success=%v, Message=%s\n", closeResp.Success, closeResp.Message)
}
