# 사용 가능한 도구 (Tools)

AI 모델이 사용할 수 있는 도구 목록과 사용법입니다.

## 1. 기본 조회 (Query)

### `execute_query`
PromQL 쿼리를 실행하여 현재 시점(또는 특정 시점)의 메트릭 값을 가져옵니다.
- **인자**:
  - `query` (필수): PromQL 쿼리문 (예: `up`, `node_cpu_seconds_total`)
  - `time` (옵션): 조회할 시간 (Unix timestamp 또는 RFC3339)

### `execute_range_query`
특정 기간 동안의 메트릭 변화 추이를 조회합니다. 그래프를 그릴 때 유용합니다.
- **인자**:
  - `query` (필수): PromQL 쿼리문
  - `start` (필수): 시작 시간
  - `end` (필수): 종료 시간
  - `step` (필수): 데이터 간격 (예: `15s`, `1m`)

## 2. 메타데이터 및 탐색

### `list_metrics`
현재 프로메테우스에 수집되고 있는 모든 메트릭의 이름을 조회합니다.
- **인자**:
  - `limit`: 가져올 개수 제한
  - `offset`: 건너뛸 개수
  - `filter_pattern`: 이름 검색 필터

### `get_metric_metadata`
특정 메트릭이 어떤 의미인지(설명), 타입(Gauge/Counter)은 무엇인지 확인합니다.
- **인자**:
  - `metric`: 메트릭 이름

### `get_targets`
프로메테우스가 어디서 데이터를 수집하고 있는지(타겟) 상태를 확인합니다. `UP` 상태인지 확인할 때 씁니다.

## 3. 관리 및 진단 (Admin)

### `get_alerts`
현재 발생 중(Firing)이거나 대기 중(Pending)인 경고(Alert) 목록을 봅니다.
- **사용 예**: "현재 서버에 무슨 문제가 있어?"라고 물어볼 때 사용됨.

### `get_config`
프로메테우스 서버의 설정 파일(`prometheus.yml`) 내용을 조회합니다.

### `get_tsdb_stats`
시계열 데이터베이스(TSDB)의 통계 정보를 확인합니다.
- **용도**: 메모리를 많이 잡아먹는 메트릭(Top 10 Series)을 찾을 때 유용합니다.

## 4. 헬스 체크

### `health_check`
이 MCP 서버가 정상 작동 중인지, 프로메테우스와 연결은 잘 되어 있는지 확인합니다.
