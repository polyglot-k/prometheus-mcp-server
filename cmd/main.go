package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/polyglot-k/prometheus-mcp-server/internal/mcp"
	"github.com/polyglot-k/prometheus-mcp-server/internal/prometheus"
)

func main() {
	// Initialize the server
	s := mcp.NewServer()

	// Register tools
	prometheus.RegisterTools(s)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
