# STOVE 인증 설정 가이드

## 개요

이 문서는 Witch's Lounge 백엔드에서 STOVE 플랫폼 인증을 설정하고 사용하는 방법을 설명합니다.

## STOVE 인증 플로우

STOVE는 **2단계 토큰 인증** 방식을 사용합니다:

### 1. 서버 초기화 단계
서버 시작 시 또는 최초 요청 시:
- 서버가 **API Access Token**을 STOVE API로부터 발급받습니다
- `client_id`, `client_secret`, `service_id`를 사용하여 발급
- 유효기간: 30일 (자동 갱신)

### 2. 사용자 인증 단계
클라이언트 로그인 시:
1. **클라이언트**: STOVE PC SDK의 `GetToken()`을 통해 **Game User Access Token** 획득
2. **클라이언트**: Game User Access Token을 게임 서버로 전송
3. **서버**: API Access Token을 헤더에 포함하여 STOVE API로 Game User Access Token 검증
4. **서버**: 검증 성공 시 사용자 정보 추출 및 세션 생성

### 토큰 종류

| 토큰 유형 | 발급 주체 | 사용 목적 | 유효 기간 |
|---------|---------|---------|---------|
| API Access Token | 서버 | 서버가 STOVE API와 통신하기 위한 인증 토큰 | 30일 |
| Game User Access Token | 클라이언트 | 사용자 인증 토큰 (클라이언트가 서버로 전송) | 6시간 |

## 참고 문서

- [STOVE Game User Access Token 검증 가이드](https://studio-docs.onstove.com/pc/server/auth/api_user_token_verify.html)
- [STOVE 인증 서비스 연동 가이드](https://studio-docs.onstove.com/pc/server/auth/auth.html)

## API 엔드포인트

### 로그인 API

**Endpoint**: `POST /api/v1/users/signin`

**Request Body**:
```json
{
  "platform_type": "stove",
  "app_id": "YOUR_STOVE_APP_ID",
  "ticket": "STOVE_ACCESS_TOKEN_FROM_CLIENT"
}
```

**Response** (성공 시):
```json
{
  "session_id": "generated-session-id",
  "user": {
    "id": "user-uuid",
    "steam_id": "stove_member_no",
    "nickname": "User Nickname",
    "steam_avatar_url": "",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## STOVE API 설정

### 환경 변수

서버 시작 전에 다음 환경 변수를 설정해야 합니다:

```bash
# .env 예시
STOVE_CLIENT_ID=com.stove.your.game.server      # STOVE 기술PM에게 요청
STOVE_CLIENT_SECRET=your_client_secret_key      # STOVE 기술PM에게 요청
STOVE_SERVICE_ID=YOUR_GAME_ID                   # 파트너스에서 등록한 Game ID
STOVE_API_BASE_URL=https://api.onstove.com      # LIVE: https://api.onstove.com, Sandbox: https://api.gate8.com
```

### STOVE API 엔드포인트

#### 1. API Access Token 발급
서버가 STOVE API와 통신하기 위한 토큰을 발급받습니다.

- **URL**: `POST {base_url}/auth/v5/server_token`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "client_id": "com.stove.your.game.server",
    "client_secret": "your_secret_key",
    "service_id": "YOUR_GAME_ID"
  }
  ```
- **Response**:
  ```json
  {
    "code": 0,
    "message": "success",
    "response_data": {
      "access_token": "server_access_token_string",
      "token_type": "bearer",
      "expires_in": 2591999
    }
  }
  ```

#### 2. Game User Access Token 검증
클라이언트로부터 받은 사용자 토큰을 검증합니다.

- **URL**: `POST {base_url}/member/v3.0/{game_id}/token/verify`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`
  - `Authorization: Bearer {api_access_token}` (1번에서 발급받은 토큰)
- **Request Body**:
  ```json
  {
    "access_token": "game_user_access_token_from_client"
  }
  ```
- **Response**:
  ```json
  {
    "code": 0,
    "message": "success",
    "value": {
      "member_no": 20005061986,
      "guid": 200000000397
    }
  }
  ```

## 구현 세부사항

### 전략 패턴

STOVE 인증은 전략 패턴을 사용하여 구현되었습니다:

1. **PlatformAuthStrategy**: 플랫폼별 인증 전략 인터페이스
2. **StoveAuthStrategy**: STOVE 플랫폼 구현체
   - API Access Token 자동 발급 및 캐싱
   - 토큰 만료 5분 전 자동 갱신
   - Thread-safe한 토큰 관리 (sync.RWMutex 사용)
3. **PlatformAuthFactory**: 플랫폼 타입에 따른 전략 생성

### 코드 위치

- 인터페이스: `internal/domain/strategy/platform_auth.go`
- STOVE 구현체: `internal/strategy/stove_auth.go`
- 팩토리: `internal/strategy/platform_factory.go`
- UseCase: `internal/usecase/user.go`
- 핸들러: `internal/delivery/http/handler/user.go`

### API Access Token 캐싱

서버는 API Access Token을 메모리에 캐싱하여 불필요한 API 호출을 방지합니다:

- **캐시 만료 시점**: 실제 만료 5분 전
- **자동 갱신**: 캐시된 토큰이 만료되면 자동으로 새 토큰 발급
- **동시성 제어**: `sync.RWMutex`를 사용한 thread-safe 구현
- **Double-check locking**: 여러 goroutine이 동시에 토큰을 요청해도 한 번만 발급

## 테스트

### 단위 테스트

```bash
go test ./internal/delivery/http/handler -v
```

### 통합 테스트 예시

```bash
curl -X POST http://localhost:8080/api/v1/users/signin \
  -H "Content-Type: application/json" \
  -d '{
    "platform_type": "stove",
    "app_id": "your_app_id",
    "ticket": "your_stove_access_token"
  }'
```

## 주의사항

### 서버측 (API Access Token)
1. **유효기간**: 30일 (서버가 자동으로 관리)
2. **발급 조건**: `client_id`, `client_secret`, `service_id` 필요
3. **갱신 정책**: 유효 기간의 70% 경과 시 신규 토큰 발급 (기존 토큰도 만료 전까지 사용 가능)
4. **보안**: `client_secret`은 절대 클라이언트에 노출하면 안 됨

### 클라이언트측 (Game User Access Token)
1. **유효기간**: 6시간
2. **토큰 갱신**: STOVE PC SDK의 `OnRenewToken()` 콜백을 통해 자동 갱신
3. **게임 이용약관**: 사용자가 게임 이용약관에 동의해야 토큰 발급 가능
4. **GUID**: 게임별 고유 회원 번호는 게임 약관 동의 시 생성됨

### 환경별 설정
- **LIVE 환경**: `https://api.onstove.com`
- **Sandbox 환경**: `https://api.gate8.com`
- 환경별로 별도의 `client_id`, `client_secret` 필요

## 구현 완료 사항

1. ✅ API Access Token 자동 발급 및 캐싱
2. ✅ 토큰 만료 전 자동 갱신 (만료 5분 전)
3. ✅ Thread-safe한 토큰 관리
4. ✅ Game User Access Token 검증 플로우
5. ✅ 사용자 정보 추출 및 세션 생성

## 다음 개선 사항

1. 환경 변수에서 설정값 자동 로드 (현재는 생성자에서 직접 전달 필요)
2. STOVE 회원 정보 조회 API 연동 (닉네임, 아바타 등 추가 정보)
3. API 호출 실패 시 retry 로직 추가
4. 구조화된 로깅 추가 (zerolog 등)
5. Rate limiting 추가
6. API Access Token을 Redis 등 외부 저장소에 저장하여 서버 재시작 시에도 유지

## 문제 해결

### API 호출 실패

- STOVE API 엔드포인트 URL이 올바른지 확인
- App ID가 정확한지 확인
- Access Token이 유효한지 확인

### 인증 실패

- 클라이언트에서 최신 access token을 전송하는지 확인
- 게임 이용약관 동의 여부 확인
- STOVE 개발자 콘솔에서 앱 설정 확인


