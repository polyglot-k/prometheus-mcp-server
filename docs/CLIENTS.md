# 다양한 클라이언트 설정 가이드

이 문서는 다양한 AI 클라이언트(Claude Desktop, Claude Code, VSCode 등)에서 Prometheus MCP 서버를 설정하고 사용하는 방법을 안내합니다.

## 1. Claude Desktop (Mac/Windows)

Claude Desktop 앱에서 이 서버를 사용하려면 설정 파일을 수정해야 합니다.

**설정 파일 위치:**
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`
- **Mac**: `~/Library/Application Support/Claude/claude_desktop_config.json`

**설정 예시 (Docker 사용 시):**
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

**설정 예시 (로컬 바이너리 사용 시):**
```json
{
  "mcpServers": {
    "prometheus": {
      "command": "c:\\절대경로\\prometheus-mcp.exe",
      "args": [],
      "env": {
        "PROMETHEUS_URL": "http://localhost:9090"
      }
    }
  }
}
```

## 2. Claude Code (CLI)

터미널 기반의 Claude Code 도구에서 사용하려면 프로젝트 루트 또는 홈 디렉토리의 설정 파일을 수정합니다.

**설정 파일 위치:**
- `mcp_config.json` (프로젝트 루트 또는 `~/.gemini/antigravity/` 등, 사용 환경에 따라 다름)

**설정 내용:**
위의 Claude Desktop 설정과 `mcpServers` 구조는 동일합니다.

## 3. VSCode (Roombot / Autogen 등 MCP 지원 확장)

VSCode에서 MCP를 지원하는 확장을 사용할 경우, 해당 확장의 설정 방식에 따릅니다. 보통 `.vscode/settings.json`이나 별도의 config 파일을 사용합니다.

**일반적인 설정 구조:**
```json
"mcp.servers": {
    "prometheus": {
        "command": "c:\\절대경로\\prometheus-mcp.exe",
        "env": {
            "PROMETHEUS_URL": "http://localhost:9090"
        }
    }
}
```

## 4. Cursor / Windsurf

현재 Cursor나 Windsurf는 MCP를 직접 지원하지 않을 수 있습니다. 
하지만 MCP 서버를 표준 입출력(Stdio)으로 실행하는 스크립트를 만들어 "External Tool" 등으로 등록하여 활용하는 우회 방법을 사용할 수 있습니다.

---

### 문제 해결 (Troubleshooting)

- **연결 실패**: `PROMETHEUS_URL`이 올바른지 확인하세요. Docker 사용 시 `localhost` 대신 `host.docker.internal`을 사용해야 할 수 있습니다.
- **권한 오류**: 윈도우에서 실행 파일 경로에 공백이나 권한 문제가 없는지 확인하세요.
- **로그 확인**: 실행 시 오류가 발생하면 로그 파일이 생성되도록 래퍼 스크립트(.bat / .cmd)를 만들어 디버깅하는 것이 좋습니다.
