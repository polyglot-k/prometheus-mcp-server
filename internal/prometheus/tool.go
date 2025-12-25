package prometheus

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all Prometheus-related tools to the MCP server.
func RegisterTools(s *server.MCPServer) {
	// 1. Initialize shared client
	client := NewClient()

	// 2. Register each tool
	NewHealthTool(client).Register(s)
	NewQueryTool(client).Register(s)
	NewMetaTool(client).Register(s)
	NewAdminTool(client).Register(s)
}
