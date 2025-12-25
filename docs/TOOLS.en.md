# Available Tools

List of tools available for AI models and how to use them.

## 1. Query

### `execute_query`
Executes a PromQL query to retrieve metric values at the current time (or a specific time).
- **Arguments**:
  - `query` (Required): PromQL query string (e.g., `up`, `node_cpu_seconds_total`)
  - `time` (Optional): Time for evaluation (Unix timestamp or RFC3339)

### `execute_range_query`
Retrieves metric data over a specific range of time. Useful for drawing graphs.
- **Arguments**:
  - `query` (Required): PromQL query string
  - `start` (Required): Start time
  - `end` (Required): End time
  - `step` (Required): Query resolution step width (e.g., `15s`, `1m`)

## 2. Metadata & Exploration

### `list_metrics`
Lists all metric names currently collected by Prometheus.
- **Arguments**:
  - `limit`: Maximum number of metrics to return
  - `offset`: Number of metrics to skip
  - `filter_pattern`: Substring filter for metric names

### `get_metric_metadata`
Retrieves metadata for a specific metric, including its help text and type (Gauge/Counter).
- **Arguments**:
  - `metric`: Metric name

### `get_targets`
Checks the status of scrape targets. Use this to verify if targets are `UP`.

## 3. Administration & Diagnostics (Admin)

### `get_alerts`
Lists all active (Firing) or pending alerts.
- **Usage Example**: Used when asking "What's wrong with the server right now?"

### `get_config`
Retrieves the current Prometheus configuration (`prometheus.yml`).

### `get_tsdb_stats`
Retrieves statistics about the Time Series Database (TSDB).
- **Usage**: Useful for finding metrics consuming high memory (Top 10 Series).

## 4. Health Check

### `health_check`
Verifies if the MCP server is running correctly and connected to Prometheus.
