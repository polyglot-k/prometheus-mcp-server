package prometheus

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// QueryTool handles PromQL queries
type QueryTool struct {
	client *Client
}

func NewQueryTool(c *Client) *QueryTool {
	return &QueryTool{client: c}
}

func (t *QueryTool) Register(s *server.MCPServer) {
	// Instant Query
	queryTool := mcp.NewTool("execute_query",
		mcp.WithDescription("Execute a PromQL instant query against Prometheus"),
		mcp.WithString("query", mcp.Required(), mcp.Description("PromQL query string")),
		mcp.WithString("time", mcp.Description("Optional RFC3339 or Unix timestamp")),
	)
	s.AddTool(queryTool, t.ExecuteQuery)

	// Range Query
	rangeTool := mcp.NewTool("execute_range_query",
		mcp.WithDescription("Execute a PromQL range query against Prometheus"),
		mcp.WithString("query", mcp.Required(), mcp.Description("PromQL query string")),
		mcp.WithString("start", mcp.Required(), mcp.Description("Start time as RFC3339 or Unix timestamp")),
		mcp.WithString("end", mcp.Required(), mcp.Description("End time as RFC3339 or Unix timestamp")),
		mcp.WithString("step", mcp.Required(), mcp.Description("Query resolution step width (e.g., '15s', '1m')")),
	)
	s.AddTool(rangeTool, t.ExecuteRangeQuery)
}

func (t *QueryTool) ExecuteQuery(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := map[string]string{"query": query}
	var timeStr string
	if args, ok := request.Params.Arguments.(map[string]interface{}); ok {
		if t, ok := args["time"].(string); ok {
			timeStr = t
		}
	}
	if timeStr != "" {
		params["time"] = timeStr
	}

	result, err := t.client.Request("query", params)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}

func (t *QueryTool) ExecuteRangeQuery(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := request.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	start, err := request.RequireString("start")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	end, err := request.RequireString("end")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	step, err := request.RequireString("step")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := map[string]string{
		"query": query,
		"start": start,
		"end":   end,
		"step":  step,
	}

	result, err := t.client.Request("query_range", params)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}
