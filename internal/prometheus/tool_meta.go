package prometheus

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// MetaTool handles metadata related requests
type MetaTool struct {
	client *Client
}

func NewMetaTool(c *Client) *MetaTool {
	return &MetaTool{client: c}
}

func (t *MetaTool) Register(s *server.MCPServer) {
	// List Metrics
	listTool := mcp.NewTool("list_metrics",
		mcp.WithDescription("List all available metrics in Prometheus"),
		mcp.WithNumber("limit", mcp.Description("Maximum number of metrics to return")),
		mcp.WithNumber("offset", mcp.Description("Number of metrics to skip")),
		mcp.WithString("filter_pattern", mcp.Description("Optional substring to filter metric names")),
	)
	s.AddTool(listTool, t.ListMetrics)

	// Get Metadata
	metaTool := mcp.NewTool("get_metric_metadata",
		mcp.WithDescription("Get metadata for a specific metric"),
		mcp.WithString("metric", mcp.Required(), mcp.Description("The name of the metric")),
	)
	s.AddTool(metaTool, t.GetMetricMetadata)

	// Get Targets
	targetsTool := mcp.NewTool("get_targets",
		mcp.WithDescription("Get information about all scrape targets"),
	)
	s.AddTool(targetsTool, t.GetTargets)
}

func (t *MetaTool) ListMetrics(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := t.client.Request("label/__name__/values", nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	metrics, ok := result.([]interface{})
	if !ok {
		return mcp.NewToolResultError("unexpected response format"), nil
	}

	// Filter
	var filtered []string
	var filterPattern string

	args, _ := request.Params.Arguments.(map[string]interface{})
	if args != nil {
		if fp, ok := args["filter_pattern"].(string); ok {
			filterPattern = fp
		}
	}

	for _, m := range metrics {
		if s, ok := m.(string); ok {
			if filterPattern == "" || strings.Contains(strings.ToLower(s), strings.ToLower(filterPattern)) {
				filtered = append(filtered, s)
			}
		}
	}

	// Pagination
	limit := 0
	if args != nil {
		if l, ok := args["limit"].(float64); ok {
			limit = int(l)
		}
	}

	offset := 0
	if args != nil {
		if o, ok := args["offset"].(float64); ok {
			offset = int(o)
		}
	}

	totalCount := len(filtered)
	end := totalCount
	if limit > 0 {
		end = offset + limit
		if end > totalCount {
			end = totalCount
		}
	}

	if offset >= totalCount {
		offset = totalCount
	}

	paginated := filtered[offset:end]

	response := map[string]interface{}{
		"metrics":        paginated,
		"total_count":    totalCount,
		"returned_count": len(paginated),
		"offset":         offset,
		"has_more":       end < totalCount,
	}

	jsonResp, _ := json.Marshal(response)
	return mcp.NewToolResultText(string(jsonResp)), nil
}

func (t *MetaTool) GetMetricMetadata(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	metric, err := request.RequireString("metric")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := map[string]string{"metric": metric}
	result, err := t.client.Request("metadata", params)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}

func (t *MetaTool) GetTargets(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	result, err := t.client.Request("targets", nil)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonResult, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonResult)), nil
}
