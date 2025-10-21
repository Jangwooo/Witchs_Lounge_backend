# 프로젝트 아키텍처

## 개요

Witch's Lounge 백엔드는 클린 아키텍처 원칙을 따르며, 모듈화된 구조로 설계되었습니다.

## 디렉토리 구조

```
app/
  └── main.go                          # 진입점 (최소한의 코드, 설정 함수 호출만)

internal/
  ├── delivery/                        # 전송 계층 (HTTP, gRPC 등)
  │   └── http/
  │       ├── handler/                 # HTTP 핸들러
  │       │   └── stove.go            # Stove 인증 핸들러
  │       ├── middleware/             # 미들웨어
  │       └── router/
  │           └── v1/                 # API v1 라우터
  │               ├── router.go       # 통합 라우터 설정
  │               └── stove.go        # Stove 라우터
  │
  ├── domain/                          # 도메인 계층
  │   ├── entity/                     # 도메인 엔티티
  │   └── repository/                 # 리포지토리 인터페이스
  │
  ├── usecase/                         # 비즈니스 로직 계층
  │   └── stove.go                    # Stove 인증 UseCase
  │
  ├── repository/                      # 리포지토리 구현
  │   └── user.go                     # 사용자 리포지토리
  │
  └── infrastructure/                  # 인프라 계층
      ├── database/                   # 데이터베이스 설정
      ├── redis/                      # Redis 클라이언트
      ├── session/                    # 세션 관리
      └── bootstrap/                  # 초기화 로직
          ├── app.go                  # 앱 의존성 초기화
          ├── database.go             # DB 초기화
          └── redis.go                # Redis 초기화
```

## 핵심 설계 원칙

### 1. 관심사의 분리 (Separation of Concerns)

각 계층은 명확한 책임을 가집니다:

- **Delivery**: HTTP 요청/응답 처리
- **UseCase**: 비즈니스 로직 실행
- **Repository**: 데이터 접근
- **Infrastructure**: 외부 시스템 연동

### 2. 의존성 역전 (Dependency Inversion)

상위 계층은 하위 계층의 인터페이스에 의존하며, 구현체는 주입됩니다.

### 3. 모듈화 (Modularity)

#### Bootstrap 패턴

`internal/infrastructure/bootstrap/` 패키지는 애플리케이션 초기화를 담당합니다:

- **database.go**: 데이터베이스 연결 설정
- **redis.go**: Redis 연결 및 세션 스토어 초기화
- **app.go**: 전체 애플리케이션 의존성 조립

#### Router 패턴

`internal/delivery/http/router/v1/` 패키지는 라우팅을 담당합니다:

- **router.go**: `SetupRoutes()` 함수로 모든 라우터를 통합 관리
- **stove.go**: Stove 관련 엔드포인트 정의

**장점:**
- 새로운 라우터 추가 시 `router.go`만 수정
- 라우터 변경이 `main.go`에 영향을 주지 않음

## main.go 구조

```go
func main() {
    // 1. 플래그 파싱
    // 2. Database 초기화
    // 3. Redis 초기화
    // 4. 애플리케이션 의존성 초기화
    // 5. Fiber 앱 생성
    // 6. 미들웨어 설정
    // 7. Swagger 문서
    // 8. Health check
    // 9. 라우터 설정 (SetupRoutes)
    // 10. 서버 시작
}
```

**특징:**
- 설정 함수만 호출
- 비즈니스 로직 없음
- 읽기 쉬운 선언적 구조

## 라우터 추가 방법

### 1. 핸들러 생성

```go
// internal/delivery/http/handler/new_feature.go
type NewFeatureHandler struct {
    useCase usecase.NewFeatureUseCase
}
```

### 2. 라우터 파일 생성

```go
// internal/delivery/http/router/v1/new_feature.go
func NewFeatureRouter(app *fiber.App, handler *handler.NewFeatureHandler) {
    feature := app.Group("/api/v1/feature")
    feature.Get("/", handler.List)
}
```

### 3. router.go에 추가

```go
// internal/delivery/http/router/v1/router.go
type RouterConfig struct {
    StoveHandler      *handler.StoveHandler
    NewFeatureHandler *handler.NewFeatureHandler  // 추가
    SessionStore      session.SessionStore
}

func SetupRoutes(app *fiber.App, config *RouterConfig) {
    NewStoveRouter(app, config.StoveHandler, config.SessionStore)
    NewFeatureRouter(app, config.NewFeatureHandler)  // 추가
}
```

### 4. bootstrap/app.go에 의존성 추가

```go
// internal/infrastructure/bootstrap/app.go
type AppDependencies struct {
    StoveHandler      *handler.StoveHandler
    NewFeatureHandler *handler.NewFeatureHandler  // 추가
    SessionStore      session.SessionStore
}

func SetupAppDependencies(dbClient *ent.Client, sessionStore session.SessionStore) *AppDependencies {
    // ...
    newFeatureUseCase := usecase.NewNewFeatureUseCase(repo)
    newFeatureHandler := handler.NewNewFeatureHandler(newFeatureUseCase)
    
    return &AppDependencies{
        StoveHandler:      stoveHandler,
        NewFeatureHandler: newFeatureHandler,  // 추가
        SessionStore:      sessionStore,
    }
}
```

**main.go는 수정 불필요!**

## API 버전 관리

현재 구조는 `/api/v1` 버전을 사용합니다.

### 새 버전 추가

```
internal/delivery/http/router/
  ├── v1/
  │   ├── router.go
  │   └── stove.go
  └── v2/              # 새 버전
      ├── router.go
      └── stove.go
```

`main.go`에서:
```go
v1.SetupRoutes(app, v1Config)
v2.SetupRoutes(app, v2Config)
```

## 환경 설정

### 서버 모드

```bash
# 프로덕션 모드 (기본)
go run app/main.go

# 개발 모드
go run app/main.go -mode=dev
```

### 환경 변수

현재는 하드코딩되어 있으나, 향후 `.env` 파일 지원 예정:

```bash
REDIS_ADDR=redis:6379
REDIS_PASSWORD=
REDIS_DB=0
DATABASE_MODE=prod
```

## 향후 개선 사항

1. 환경 변수 관리 (viper 또는 godotenv)
2. 구조화된 로깅 (zerolog)
3. Graceful shutdown
4. 메트릭 수집 (Prometheus)
5. 분산 추적 (OpenTelemetry)

