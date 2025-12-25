package mcp

import (
	"github.com/mark3labs/mcp-go/server"
)

// NewServer initializes and returns a new MCP server.
func NewServer() *server.MCPServer {
	s := server.NewMCPServer(
		"Prometheus MCP Server Based Go",
		"0.0.1",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	return s
}
