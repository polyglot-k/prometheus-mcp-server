# Prometheus MCP Server

An MCP (Model Context Protocol) server that allows AI models like Claude or Gemini to connect to Prometheus to query metrics and analyze data.

## Features

- **Health Check**: Verify server status and Prometheus connectivity
- **Execute Queries**: Run PromQL queries (Instant & Range)
- **Explore Metrics**: List available metrics and view metadata
- **Management**: View active targets, alerts, and configuration

## Installation & Running

### Prerequisites
- Go 1.23+ (for local build)
- Docker (for Docker execution)
- Running Prometheus server (connection target)

### 1. Run Locally

1. **Build**
   ```bash
   make build
   ```
   This generates the `prometheus-mcp.exe` file.

2. **MCP Config (`mcp_config.json`)**
   ```json
   {
     "mcpServers": {
       "prometheus-mcp": {
         "command": "c:\\path\\to\\prometheus-mcp.exe",
         "args": [],
         "env": {
           "PROMETHEUS_URL": "http://localhost:9090"
         }
       }
     }
   }
   ```

### 2. Run with Docker

You can use the latest image from GitHub Container Registry (GHCR).

**MCP Config (`mcp_config.json`)**
```json
{
  "mcpServers": {
    "prometheus-mcp": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "-e", "PROMETHEUS_URL=http://host.docker.internal:9090",
        "ghcr.io/polyglot-k/prometheus-mcp-server:latest"
      ]
    }
  }
}
```

> **Note**: To build manually, run `make docker-build`.

## Documentation

For more detailed usage and configuration guides, refer to the documents below:

- [🔨 Detailed Tool Usage (Query, Admin, etc.)](docs/TOOLS.en.md)
- [⚙️ Client Configuration Guide (Claude Desktop, VSCode, etc.)](docs/CLIENTS.en.md)

## Run Example Environment

If you don't have a Prometheus server, you can set up a test environment using the `example` folder.

```bash
cd example
docker-compose up -d
```
- **Prometheus**: http://localhost:9090
- **Node Exporter**: http://localhost:9100
