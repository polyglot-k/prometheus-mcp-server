# Client Configuration Guide

This guide explains how to configure and use the Prometheus MCP Server with various AI clients (Claude Desktop, Claude Code, VSCode, etc.).

## 1. Claude Desktop (Mac/Windows)

To use this server with the Claude Desktop app, you need to modify the configuration file.

**Config File Location:**
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Mac**: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Example Config (Using Docker):**
```json
{
  "mcpServers": {
    "prometheus": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e", "PROMETHEUS_URL=http://host.docker.internal:9090",
        "ghcr.io/polyglot-k/prometheus-mcp-server:latest"
      ]
    }
  }
}
```

**Example Config (Using Local Binary):**
```json
{
  "mcpServers": {
    "prometheus": {
      "command": "c:\\absolute\\path\\to\\prometheus-mcp.exe",
      "args": [],
      "env": {
        "PROMETHEUS_URL": "http://localhost:9090"
      }
    }
  }
}
```

## 2. Claude Code (CLI)

To use with the terminal-based Claude Code tool, modify the configuration file in your project root or home directory.

**Config File Location:**
- `mcp_config.json` (Project root or `~/.gemini/antigravity/` depending on your environment)

**Configuration Content:**
The `mcpServers` structure is identical to the Claude Desktop configuration above.

## 3. VSCode (Roombot / Autogen etc.)

If you are using a VSCode extension that supports MCP, follow the extension's configuration instructions. Usually, this involves `.vscode/settings.json` or a separate config file.

**General Config Structure:**
```json
"mcp.servers": {
    "prometheus": {
        "command": "c:\\absolute\\path\\to\\prometheus-mcp.exe",
        "env": {
            "PROMETHEUS_URL": "http://localhost:9090"
        }
    }
}
```

## 4. Cursor / Windsurf

Currently, Cursor or Windsurf might not natively support MCP.
However, you can use workarounds like creating a wrapper script that runs the MCP server via Stdio and registering it as an "External Tool".

---

### Troubleshooting

- **Connection Failed**: Verify `PROMETHEUS_URL`. When using Docker, you might need `host.docker.internal` instead of `localhost`.
- **Permission Errors**: On Windows, check for spaces in paths or permission issues with the executable.
- **Check Logs**: If errors occur, create a wrapper script (.bat / .cmd) to capture logs for debugging.
