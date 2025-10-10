# Steam 인증 설정 가이드

## 개요

이 문서는 Witch's Lounge 백엔드에서 Steam 플랫폼 인증을 설정하고 사용하는 방법을 설명합니다.

## Steam 인증 플로우

1. **클라이언트**: Steamworks SDK의 `GetAuthSessionTicket()` 함수를 호출하여 인증 티켓 획득
2. **클라이언트**: 획득한 티켓을 hex 문자열로 변환하여 게임 서버로 전송
3. **서버**: Steam Web API의 `ISteamUserAuth/AuthenticateUserTicket`을 호출하여 티켓 검증
4. **서버**: 검증 성공 시 `ISteamUser/GetPlayerSummaries`를 호출하여 사용자 정보 조회
5. **서버**: 사용자 정보 저장 및 세션 생성

## Steam Web API 키 발급

### 1. Steam 계정 필요
- Steam 계정으로 로그인 필요
- 개발자 계정이 아니어도 가능

### 2. API 키 발급
1. [Steam Web API 키 등록 페이지](https://steamcommunity.com/dev/apikey) 접속
2. 도메인 이름 입력 (예: `localhost` 또는 실제 도메인)
3. 약관 동의 후 키 발급
4. 발급받은 키를 안전하게 보관

### 3. 환경 변수 설정

```bash
# .env 파일 또는 환경 변수
export STEAM_WEB_API_KEY="your_steam_web_api_key_here"
```

## API 엔드포인트

### 로그인 API

**Endpoint**: `POST /api/v1/users/signin`

**Request Body**:
```json
{
  "platform_type": "steam",
  "app_id": "480",
  "ticket": "14000000b200000001000000020000004a4e4..."
}
```

**Response** (성공 시):
```json
{
  "session_id": "generated-session-id",
  "user": {
    "id": "user-uuid",
    "platform_type": "steam",
    "platform_user_id": "76561198012345678",
    "platform_display_name": "PlayerName",
    "platform_avatar_url": "https://avatars.steamstatic.com/...",
    "nickname": "PlayerName",
    "language": "ko",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## Steam Web API

### 1. AuthenticateUserTicket

티켓 검증 API

- **URL**: `https://api.steampowered.com/ISteamUserAuth/AuthenticateUserTicket/v1/`
- **Method**: `GET`
- **Parameters**:
  - `key`: Steam Web API 키
  - `appid`: 게임 App ID
  - `ticket`: 클라이언트에서 받은 인증 티켓

**응답 예시**:
```json
{
  "response": {
    "params": {
      "result": "OK",
      "steamid": "76561198012345678",
      "ownersteamid": "76561198012345678",
      "vacbanned": false,
      "publisherbanned": false
    }
  }
}
```

### 2. GetPlayerSummaries

사용자 프로필 정보 조회 API

- **URL**: `https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/`
- **Method**: `GET`
- **Parameters**:
  - `key`: Steam Web API 키
  - `steamids`: 조회할 Steam ID (쉼표로 구분하여 여러 개 가능)

**응답 예시**:
```json
{
  "response": {
    "players": [
      {
        "steamid": "76561198012345678",
        "communityvisibilitystate": 3,
        "profilestate": 1,
        "personaname": "PlayerName",
        "profileurl": "https://steamcommunity.com/id/playername/",
        "avatar": "https://avatars.steamstatic.com/.../avatar.jpg",
        "avatarmedium": "https://avatars.steamstatic.com/.../avatar_medium.jpg",
        "avatarfull": "https://avatars.steamstatic.com/.../avatar_full.jpg",
        "personastate": 1,
        "realname": "Real Name",
        "timecreated": 1234567890,
        "loccountrycode": "KR"
      }
    ]
  }
}
```

## 개발 모드

API 키가 설정되지 않은 경우 **개발 모드**로 동작합니다:

- 티켓을 SteamID로 간주
- 실제 Steam API 호출 없이 기본 사용자 정보 생성
- 테스트 및 개발 환경에서 유용

```bash
# API 키 없이 실행
# 자동으로 개발 모드로 전환
go run app/main.go
```

**개발 모드 로그인 예시**:
```json
{
  "platform_type": "steam",
  "app_id": "480",
  "ticket": "76561198012345678"
}
```

## 클라이언트 구현 예시

### C++ (Steamworks SDK)

```cpp
#include "steam/steam_api.h"

// 인증 티켓 획득
HAuthTicket hAuthTicket;
uint8 rgubTicket[1024];
uint32 unTicketSize;

hAuthTicket = SteamUser()->GetAuthSessionTicket(
    rgubTicket, 
    sizeof(rgubTicket), 
    &unTicketSize
);

// 티켓을 hex 문자열로 변환
std::string ticketHex;
for (uint32 i = 0; i < unTicketSize; i++) {
    char hex[3];
    sprintf(hex, "%02X", rgubTicket[i]);
    ticketHex += hex;
}

// 서버로 전송
SendTicketToServer(ticketHex);
```

### Unity (Steamworks.NET)

```csharp
using Steamworks;

// 인증 티켓 획득
byte[] ticket = new byte[1024];
uint ticketSize;
HAuthTicket authTicket = SteamUser.GetAuthSessionTicket(
    ticket, 
    1024, 
    out ticketSize
);

// 티켓을 hex 문자열로 변환
string ticketHex = BitConverter.ToString(ticket, 0, (int)ticketSize)
    .Replace("-", "");

// 서버로 전송
SendTicketToServer(ticketHex);
```

## 구현 세부사항

### 파일 구조

- **인터페이스**: `internal/domain/strategy/platform_auth.go`
- **Steam 구현체**: `internal/strategy/steam_auth.go`
- **팩토리**: `internal/strategy/platform_factory.go`
- **핸들러**: `internal/delivery/http/handler/user.go`
- **UseCase**: `internal/usecase/user.go`

### 주요 기능

1. **실제 Steam Web API 호출**
   - 티켓 검증: `ISteamUserAuth/AuthenticateUserTicket`
   - 사용자 정보: `ISteamUser/GetPlayerSummaries`

2. **풍부한 사용자 정보**
   - Steam ID
   - 닉네임 (Persona Name)
   - 프로필 URL
   - 아바타 이미지 (3가지 크기)
   - 실명 (공개된 경우)
   - 국가 코드
   - 계정 생성 시간

3. **에러 핸들링**
   - API 키 없을 때: 개발 모드로 전환
   - 티켓 검증 실패: 명확한 오류 메시지
   - 사용자 정보 조회 실패: 기본 정보로 계속 진행

4. **국가 코드 → 언어 매핑**
   - KR → ko
   - US, GB → en
   - JP → ja
   - CN, TW → zh

## 테스트

### 단위 테스트

```bash
go test ./internal/delivery/http/handler -v -run TestUserHandler_SignIn
```

### 통합 테스트

```bash
# API 키 설정
export STEAM_WEB_API_KEY="your_key"

# 서버 실행
go run app/main.go

# 로그인 테스트
curl -X POST http://localhost:8080/api/v1/users/signin \
  -H "Content-Type: application/json" \
  -d '{
    "platform_type": "steam",
    "app_id": "480",
    "ticket": "your_actual_ticket_from_client"
  }'
```

## 보안 고려사항

1. **API 키 관리**
   - 환경 변수로만 관리
   - Git에 커밋하지 않기 (.env 파일을 .gitignore에 추가)
   - 프로덕션 환경에서는 시크릿 관리 도구 사용

2. **HTTPS 사용**
   - 프로덕션 환경에서는 반드시 HTTPS 사용
   - 티켓이 중간에 가로채지지 않도록 보호

3. **티켓 재사용 방지**
   - 클라이언트는 로그인 시마다 새 티켓 생성
   - 서버는 한 번 사용된 티켓을 기록하여 재사용 방지 (선택사항)

4. **타임아웃 설정**
   - HTTP 클라이언트 타임아웃: 10초
   - 과도한 대기 방지

## 문제 해결

### API 키 오류

```
⚠️  경고: STEAM_WEB_API_KEY 환경 변수가 설정되지 않았습니다
🔧 개발 모드: Steam API 키가 없어 티켓을 SteamID로 간주합니다
```

**해결**: 환경 변수 `STEAM_WEB_API_KEY` 설정

### 티켓 검증 실패

```
Steam API 오류: Invalid ticket (코드: 102)
```

**원인**:
- 만료된 티켓
- 잘못된 App ID
- 잘못된 티켓 형식

**해결**:
- 클라이언트에서 새 티켓 생성
- 올바른 App ID 사용
- 티켓을 hex 문자열로 올바르게 변환

### 사용자 정보 조회 실패

```
⚠️  Steam 사용자 정보 조회 실패: 사용자 정보를 찾을 수 없습니다
```

**원인**:
- 비공개 프로필
- 잘못된 Steam ID

**해결**:
- 기본 정보로 계속 진행 (자동 처리됨)
- 사용자에게 프로필 공개 요청 (선택사항)

## 참고 자료

- [Steamworks Web API 문서](https://partner.steamgames.com/doc/webapi_overview?l=koreana)
- [Steam Web API 키 발급](https://steamcommunity.com/dev/apikey)
- [ISteamUserAuth 인터페이스](https://partner.steamgames.com/doc/webapi/ISteamUserAuth)
- [ISteamUser 인터페이스](https://partner.steamgames.com/doc/webapi/ISteamUser)
- [Steamworks SDK 다운로드](https://partner.steamgames.com/downloads/steamworks_sdk.zip)

## 다음 개선 사항

1. 티켓 캐싱 및 재사용 방지
2. Rate limiting 구현
3. Retry 로직 추가 (네트워크 오류 시)
4. 더 많은 Steam 사용자 정보 활용
5. VAC Ban 체크 및 처리
6. 게임 소유권 검증 추가

