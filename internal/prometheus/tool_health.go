package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type HealthTool struct {
	client *Client
}

func NewHealthTool(c *Client) *HealthTool {
	return &HealthTool{client: c}
}

func (t *HealthTool) Register(s *server.MCPServer) {
	tool := mcp.NewTool("health_check",
		mcp.WithDescription("Return health status of the MCP server and Prometheus connection"),
	)
	s.AddTool(tool, t.Execute)
}

func (t *HealthTool) Execute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	status := map[string]interface{}{
		"status":    "healthy",
		"service":   "prometheus-mcp-server",
		"version":   "0.0.1",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"configuration": map[string]bool{
			"prometheus_url_configured": t.client.config.URL != "",
			"authentication_configured": t.client.config.Username != "" || t.client.config.Token != "",
			"org_id_configured":         t.client.config.OrgID != "",
		},
	}

	if t.client.config.URL != "" {
		_, err := t.client.Request("query", map[string]string{
			"query": "up",
			"time":  fmt.Sprintf("%d", time.Now().Unix()),
		})

		if err == nil {
			status["prometheus_connectivity"] = "healthy"
			status["prometheus_url"] = t.client.config.URL
		} else {
			status["status"] = "degraded"
			status["prometheus_connectivity"] = "unhealthy"
			status["prometheus_error"] = err.Error()
		}
	} else {
		status["status"] = "unhealthy"
		status["error"] = "PROMETHEUS_URL not configured"
	}

	jsonStatus, _ := json.Marshal(status)
	return mcp.NewToolResultText(string(jsonStatus)), nil
}
