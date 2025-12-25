package prometheus

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AdminTool handles administrative and operational requests
type AdminTool struct {
	client *Client
}

func NewAdminTool(c *Client) *AdminTool {
	return &AdminTool{client: c}
}

func (t *AdminTool) Register(s *server.MCPServer) {
	// Get Alerts
	alertsTool := mcp.NewTool("get_alerts",
		mcp.WithDescription("Get a list of all active alerts"),
	)
	s.AddTool(alertsTool, t.GetAlerts)

	// Get Config
	configTool := mcp.NewTool("get_config",
		mcp.WithDescription("Get the current Prometheus configuration"),
	)
	s.AddTool(configTool, t.GetConfig)

	// Get TSDB Stats (Cardinality)
	tsdbTool := mcp.NewTool("get_tsdb_stats",
		mcp.WithDescription("Get TSDB statistics including cardinality info"),
	)
	s.AddTool(tsdbTool, t.GetTSDBStats)
}

func (t *AdminTool) GetAlerts(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := t.client.Request("alerts", nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}

func (t *AdminTool) GetConfig(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := t.client.Request("status/config", nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}

func (t *AdminTool) GetTSDBStats(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get stats regarding the head block of the TSDB
	result, err := t.client.Request("status/tsdb", nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}
