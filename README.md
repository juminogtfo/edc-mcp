# edc-mcp

Go stdio MCP server for **EDC** (everyday carry) tools agents can call from chat—timezone, directory listings, and similar utilities as the catalog grows.

## Bootstrap

Today the only tool is `healthcheck`. It returns process metadata and a monotonic `healthcheck_call_count` so agents can confirm Cursor is talking to live MCP logic, not a dead or static HTTP façade.

## Build

```powershell
cd c:\Users\alaroc\.cursor\edc-mcp
go build -o bin\edc-mcp.exe .\cmd\edc-mcp
```

Cursor is configured in `mcp.json` to launch `bin\edc-mcp.exe`. Built binaries are listed in `.cursorignore` so agents do not read them directly.

## Add tools later

Register new handlers under `internal/tools/` and call them from `cmd/edc-mcp/main.go`.