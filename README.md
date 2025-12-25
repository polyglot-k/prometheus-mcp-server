# Prometheus MCP Server

Claude나 Gemini 같은 AI 모델이 Prometheus에 접속하여 메트릭을 조회하고 분석할 수 있게 해주는 MCP(Model Context Protocol) 서버입니다.

## 기능

- **헬스 체크**: 서버 및 프로메테우스 연결 상태 확인
- **쿼리 실행**: PromQL 쿼리 실행 (Instant & Range)
- **메트릭 탐색**: 사용 가능한 메트릭 목록 및 메타데이터 조회
- **관리 기능**: 활성 타겟, 경고(Alert), 설정 확인

## 설치 및 실행

### 필수 요구사항
- Go 1.23+ (로컬 빌드 시)
- Docker (도커 실행 시)
- 실행 중인 Prometheus 서버 (연결 대상)

### 1. 로컬에서 실행하기

1. **빌드**
   ```bash
   make build
   ```
   `prometheus-mcp.exe` 파일이 생성됩니다.

2. **MCP 설정 (`mcp_config.json`)**
   ```json
   {
     "mcpServers": {
       "prometheus-mcp": {
         "command": "c:\\경로\\prometheus-mcp.exe",
         "args": [],
         "env": {
           "PROMETHEUS_URL": "http://localhost:9090"
         }
       }
     }
   }
   ```

### 2. Docker로 실행하기

GitHub Container Registry(GHCR)에서 최신 이미지를 받아 실행할 수 있습니다.

**MCP 설정 (`mcp_config.json`)**
```json
{
  "mcpServers": {
    "prometheus-mcp": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "-e", "PROMETHEUS_URL=http://host.docker.internal:9090",
        "ghcr.io/polyglot-k/prometheus-mcp:latest"
      ]
    }
  }
}
```

> **참고**: 직접 빌드하려면 `make docker-build`를 실행하세요.

## 문서 (Documentation)

더 자세한 사용법과 설정 방법은 아래 문서를 참고하세요.

- [🔨 도구 상세 사용법 (Query, Admin 등)](docs/TOOLS.md)
- [⚙️ 클라이언트별 설정 가이드 (Claude Desktop, VSCode 등)](docs/CLIENTS.md)

## 예제 환경 실행

프로메테우스 서버가 없다면 `example` 폴더를 사용하여 테스트 환경을 띄울 수 있습니다.

```bash
cd example
docker-compose up -d
```
- **Prometheus**: http://localhost:9090
- **Node Exporter**: http://localhost:9100
