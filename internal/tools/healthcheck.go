package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alaroc/edc-mcp/internal/runtime"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterHealthcheck exposes a stateful liveness probe for agents.
func RegisterHealthcheck(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "healthcheck",
		Description: "Verify the EDC MCP server is running live tool logic (not a static stub). " +
			"Returns process metadata and a monotonically increasing healthcheck_call_count " +
			"so repeated calls prove in-process state is updating.",
	}, healthcheck)
}

type healthcheckArgs struct {
	Echo string `json:"echo,omitempty" jsonschema:"optional string echoed back to confirm round-trip handling"`
}

func healthcheck(_ context.Context, _ *mcp.CallToolRequest, args healthcheckArgs) (*mcp.CallToolResult, any, error) {
	snap := runtime.Health()
	out := map[string]any{
		"health": snap,
	}
	if args.Echo != "" {
		out["echo"] = args.Echo
	}

	raw, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("marshal healthcheck: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(raw)},
		},
	}, out, nil
}