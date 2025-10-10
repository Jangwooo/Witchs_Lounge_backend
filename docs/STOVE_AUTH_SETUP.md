# STOVE 인증 설정 가이드

## 개요

이 문서는 Witch's Lounge 백엔드에서 STOVE 플랫폼 인증을 설정하고 사용하는 방법을 설명합니다.

## STOVE 인증 플로우

1. **클라이언트**: STOVE PC SDK의 `GetToken()`을 통해 `access token` 획득
2. **클라이언트**: access token을 게임 서버로 전송
3. **서버**: STOVE API를 호출하여 access token 유효성 검증
4. **서버**: 검증 성공 시 사용자 정보 추출 및 세션 생성

## 참고 문서

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

현재 STOVE API URL은 코드에 하드코딩되어 있습니다. 프로덕션 환경에서는 환경 변수로 관리하는 것을 권장합니다.

```bash
# .env 예시
STOVE_API_BASE_URL=https://api.onstove.com
STOVE_AUTH_ENDPOINT=/auth/v1/game/validate
```

### STOVE API 엔드포인트

현재 사용 중인 STOVE API:
- **URL**: `https://api.onstove.com/auth/v1/game/validate`
- **Method**: `GET`
- **Headers**:
  - `Content-Type: application/json`
  - `Authorization: Bearer {access_token}`
  - `X-Stove-App-Id: {app_id}`

### API 응답 구조

```json
{
  "result": {
    "code": 0,
    "message": "success"
  },
  "member": {
    "memberNo": "stove_member_id",
    "gameUserNo": "game_user_id"
  }
}
```

## 구현 세부사항

### 전략 패턴

STOVE 인증은 전략 패턴을 사용하여 구현되었습니다:

1. **PlatformAuthStrategy**: 플랫폼별 인증 전략 인터페이스
2. **StoveAuthStrategy**: STOVE 플랫폼 구현체
3. **PlatformAuthFactory**: 플랫폼 타입에 따른 전략 생성

### 코드 위치

- 인터페이스: `internal/domain/strategy/platform_auth.go`
- STOVE 구현체: `internal/strategy/stove_auth.go`
- 팩토리: `internal/strategy/platform_factory.go`
- 핸들러: `internal/delivery/http/handler/user.go`

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

1. **Access Token 유효기간**: STOVE access token의 유효기간은 6시간입니다.
2. **토큰 갱신**: STOVE PC SDK의 `OnRenewToken()` 콜백을 통해 토큰이 자동으로 갱신됩니다.
3. **게임 이용약관**: 사용자가 게임 이용약관에 동의해야 토큰을 받을 수 있습니다.
4. **API URL 확인**: 실제 STOVE API URL은 STOVE 개발자 센터에서 확인하고 업데이트해야 합니다.

## 다음 개선 사항

1. STOVE API URL을 환경 변수로 관리
2. STOVE API에서 추가 사용자 정보 가져오기 (닉네임, 아바타 등)
3. 토큰 갱신 로직 구현
4. 에러 핸들링 개선 및 로깅 추가
5. Rate limiting 및 retry 로직 추가

## 문제 해결

### API 호출 실패

- STOVE API 엔드포인트 URL이 올바른지 확인
- App ID가 정확한지 확인
- Access Token이 유효한지 확인

### 인증 실패

- 클라이언트에서 최신 access token을 전송하는지 확인
- 게임 이용약관 동의 여부 확인
- STOVE 개발자 콘솔에서 앱 설정 확인


