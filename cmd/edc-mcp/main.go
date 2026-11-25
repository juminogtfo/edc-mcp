// Command edc-mcp is a stdio MCP server that exposes everyday carry tools for agents.
package main

import (
	"context"
	"log"
	"os"

	"github.com/alaroc/edc-mcp/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "edc-mcp",
		Version: "0.1.0",
	}, nil)

	tools.RegisterHealthcheck(server)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("edc-mcp: %v", err)
	}
}