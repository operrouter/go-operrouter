package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/operrouter/go-operrouter/operrouter"
)

func main() {
	client := operrouter.NewHTTP("http://localhost:8080")
	// brief wait in case server is starting separately
	time.Sleep(200 * time.Millisecond)
	res, err := client.Ping(context.Background())
	if err != nil {
		log.Fatalf("ping failed: %v", err)
	}
	fmt.Printf("ping: %#v\n", res)
}
