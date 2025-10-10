# Witch's Lounge Backend API 설계서 초안

## 1. 개요

### 1.1 프로젝트 정보
- **프로젝트명**: Witch's Lounge Backend
- **버전**: 1.0
- **설명**: 리듬 게임 "Witch's Lounge"의 백엔드 API 서버
- **기술 스택**: Go, Fiber, EntGo, Redis, PostgreSQL

### 1.2 기본 정보
- **Base URL**: `http://localhost:8080/api/v1`
- **인증 방식**: Bearer Token (JWT)
- **응답 형식**: JSON

## 2. 데이터 모델

### 2.1 사용자 (User)
```go
type User struct {
    ID                   uuid.UUID `json:"id"`
    PlatformType         string    `json:"platform_type"`        // steam
    PlatformUserID       string    `json:"platform_user_id"`     
    PlatformEmail        string    `json:"platform_email"`       
    PlatformAvatarURL    string    `json:"platform_avatar_url"`  
    PlatformDisplayName  string    `json:"platform_display_name"`
    Language             string    `json:"language"`              // ko
    PlatformData         map[string]interface{} `json:"platform_data"`
    IsVerified           bool      `json:"is_verified"`
    
    Nickname             string    `json:"nickname"`              // 게임 내 닉네임
    DisplayName          string    `json:"display_name"`          
    LastLoginAt          time.Time `json:"last_login_at"`
    Level                int       `json:"level"`                 // 유저 레벨
    Exp                  int       `json:"exp"`                   // 경험치
    Coin                 int       `json:"coin"`                  // 게임 내 코인
    Gem                  int       `json:"gem"`                   // 프리미엄 재화
    Settings             map[string]interface{} `json:"settings"`
    CustomizeData        map[string]interface{} `json:"customize_data"`
    SaveData             map[string]interface{} `json:"save_data"`
    IsBanned             bool      `json:"is_banned"`
    BannedUntil          time.Time `json:"banned_until"`
    BanReason            string    `json:"ban_reason"`
    
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}
```

### 2.2 음악 (Music)
```go
type Music struct {
    ID           uuid.UUID `json:"id"`
    Name         string    `json:"name"`          // 곡 제목
    Artist       string    `json:"artist"`        // 아티스트
    Composer     string    `json:"composer"`      // 작곡가
    MusicSource  string    `json:"music_source"`  // 음악 파일 경로
    JacketSource string    `json:"jacket_source"` // 재킷 이미지 경로
    Duration     float64   `json:"duration"`      // 곡 길이(초)
    BPM          float64   `json:"bpm"`           // BPM
    Genre        string    `json:"genre"`         // 장르
    Description  string    `json:"description"`   // 곡 설명
    IsFeatured   bool      `json:"is_featured"`   // 추천곡 여부
    IsFree       bool      `json:"is_free"`       // 무료곡 여부
    UnlockLevel  int       `json:"unlock_level"`  // 해금 레벨
    ReleaseDate  time.Time `json:"release_date"`  // 출시일
    IsActive     bool      `json:"is_active"`     // 활성 여부
    
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
```

### 2.3 스테이지 (Stage)
```go
type Stage struct {
    ID             uuid.UUID `json:"id"`
    MusicID        uuid.UUID `json:"music_id"`        // 음악 ID
    LevelName      string    `json:"level_name"`      // Easy, Normal, Hard, Expert
    Difficulty     int       `json:"difficulty"`      // 난이도 수치 (1-10)
    LevelAddress   string    `json:"level_address"`   // 채보 파일 경로
    JacketAddress  string    `json:"jacket_address"`  // 난이도별 재킷 이미지 경로
    TotalNotes     int       `json:"total_notes"`     // 총 노트 수
    MaxCombo       int       `json:"max_combo"`       // 최대 콤보
    IsActive       bool      `json:"is_active"`       // 활성 여부
    
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
```

### 2.4 게임 기록 (Record)
```go
type Record struct {
    ID             uuid.UUID `json:"id"`
    UserID         uuid.UUID `json:"user_id"`         // 유저 ID
    MusicID        uuid.UUID `json:"music_id"`        // 음악 ID
    StageID        uuid.UUID `json:"stage_id"`        // 스테이지 ID
    CharacterID    uuid.UUID `json:"character_id"`    // 캐릭터 ID
    Score          int       `json:"score"`           // 점수
    PerfectCount   int       `json:"perfect_count"`   // Perfect 개수
    GoodCount      int       `json:"good_count"`      // Good 개수
    BadCount       int       `json:"bad_count"`       // Bad 개수
    MissCount      int       `json:"miss_count"`      // Miss 개수
    MaxCombo       int       `json:"max_combo"`       // 최대 콤보
    Accuracy       float64   `json:"accuracy"`        // 정확도 (%)
    Rank           string    `json:"rank"`            // F, D, C, B, A, S, SS, SSS
    IsFullCombo    bool      `json:"is_full_combo"`   // 풀콤보 여부
    IsPerfectPlay  bool      `json:"is_perfect_play"` // 퍼펙트 플레이 여부
    PlayedAt       time.Time `json:"played_at"`       // 플레이 시간
    PlayDuration   int       `json:"play_duration"`   // 플레이 소요시간(초)
    AdditionalInfo map[string]interface{} `json:"additional_info"`
    IsValid        bool      `json:"is_valid"`        // 유효한 기록 여부
    
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}
```

### 2.5 업적 (Achievement)
```go
type Achievement struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`        // 업적 이름
    Description string    `json:"description"` // 업적 설명
    IconURL     string    `json:"icon_url"`    // 업적 아이콘 URL
    Type        string    `json:"type"`        // score, combo, accuracy, play_count, special
    Conditions  map[string]interface{} `json:"conditions"` // 달성 조건
    Rewards     map[string]interface{} `json:"rewards"`    // 보상
    Points      int       `json:"points"`      // 업적 포인트
    IsHidden    bool      `json:"is_hidden"`   // 숨김 업적 여부
    IsActive    bool      `json:"is_active"`   // 활성 여부
    
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### 2.6 캐릭터 (Character)
```go
type Character struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`        // 캐릭터 이름
    Description string    `json:"description"` // 캐릭터 설명
    Source      string    `json:"source"`      // 캐릭터 이미지 경로
    
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### 2.7 아이템 (Item)
```go
type Item struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`        // 아이템 이름
    Description string    `json:"description"` // 아이템 설명
    EffectID    string    `json:"effect_id"`   // 효과 ID
    Type        string    `json:"type"`        // hat, cane, clothes
    Source      string    `json:"source"`      // 아이템 이미지 경로
    
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### 2.8 상품 (Product)
```go
type Product struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`        // 상품 이름
    Description string    `json:"description"` // 상품 설명
    Price       float64   `json:"price"`       // 가격
    Type        string    `json:"type"`        // hat, cane, clothes, character
    ItemID      uuid.UUID `json:"item_id"`     // 아이템 ID
    CharacterID uuid.UUID `json:"character_id"` // 캐릭터 ID
    
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 3. API 엔드포인트

### 3.1 인증 관련 API

#### 3.1.1 플랫폼 로그인
```http
POST /api/v1/users
```

**Request Body:**
```json
{
    "platform_type": "steam",
    "app_id": "your_app_id",
    "ticket": "steam_auth_ticket"
}
```

**Response:**
```json
{
    "access_token": "jwt_token",
    "refresh_token": "refresh_token",
    "expires_in": 3600,
    "user": {
        "id": "user_uuid",
        "nickname": "user_nickname",
        "level": 1,
        "exp": 0,
        "coin": 0,
        "gem": 0
    }
}
```

#### 3.1.2 내 정보 조회
```http
GET /api/v1/users/me
Authorization: Bearer {access_token}
```

**Response:**
```json
{
    "id": "user_uuid",
    "platform_type": "steam",
    "nickname": "user_nickname",
    "display_name": "display_name",
    "level": 1,
    "exp": 0,
    "coin": 0,
    "gem": 0,
    "settings": {},
    "customize_data": {},
    "last_login_at": "2024-01-01T00:00:00Z",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
}
```

### 3.2 음악 관련 API

#### 3.2.1 음악 목록 조회
```http
GET /api/v1/musics
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `page`: 페이지 번호 (default: 1)
- `limit`: 페이지 크기 (default: 20)
- `genre`: 장르 필터
- `is_featured`: 추천곡 필터 (true/false)
- `is_free`: 무료곡 필터 (true/false)

**Response:**
```json
{
    "data": [
        {
            "id": "music_uuid",
            "name": "곡 제목",
            "artist": "아티스트",
            "composer": "작곡가",
            "music_source": "/path/to/music.mp3",
            "jacket_source": "/path/to/jacket.jpg",
            "duration": 180.5,
            "bpm": 120.0,
            "genre": "Pop",
            "description": "곡 설명",
            "is_featured": true,
            "is_free": true,
            "unlock_level": 1,
            "stages": [
                {
                    "id": "stage_uuid",
                    "level_name": "Easy",
                    "difficulty": 3,
                    "total_notes": 200,
                    "max_combo": 200
                }
            ]
        }
    ],
    "pagination": {
        "page": 1,
        "limit": 20,
        "total": 100,
        "total_pages": 5
    }
}
```

#### 3.2.2 음악 상세 조회
```http
GET /api/v1/musics/{music_id}
Authorization: Bearer {access_token}
```

**Response:**
```json
{
    "id": "music_uuid",
    "name": "곡 제목",
    "artist": "아티스트",
    "composer": "작곡가",
    "music_source": "/path/to/music.mp3",
    "jacket_source": "/path/to/jacket.jpg",
    "duration": 180.5,
    "bpm": 120.0,
    "genre": "Pop",
    "description": "곡 설명",
    "is_featured": true,
    "is_free": true,
    "unlock_level": 1,
    "stages": [
        {
            "id": "stage_uuid",
            "level_name": "Easy",
            "difficulty": 3,
            "level_address": "/path/to/chart.json",
            "jacket_address": "/path/to/jacket_easy.jpg",
            "total_notes": 200,
            "max_combo": 200
        }
    ]
}
```

### 3.3 게임 기록 관련 API

#### 3.3.1 게임 기록 저장
```http
POST /api/v1/records
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
    "music_id": "music_uuid",
    "stage_id": "stage_uuid",
    "character_id": "character_uuid",
    "score": 950000,
    "perfect_count": 180,
    "good_count": 15,
    "bad_count": 3,
    "miss_count": 2,
    "max_combo": 195,
    "accuracy": 95.5,
    "rank": "A",
    "is_full_combo": false,
    "is_perfect_play": false,
    "play_duration": 180,
    "additional_info": {}
}
```

**Response:**
```json
{
    "id": "record_uuid",
    "score": 950000,
    "rank": "A",
    "accuracy": 95.5,
    "is_full_combo": false,
    "is_perfect_play": false,
    "played_at": "2024-01-01T00:00:00Z",
    "is_personal_best": true,
    "exp_gained": 100,
    "coin_gained": 50
}
```

#### 3.3.2 내 기록 조회
```http
GET /api/v1/records/me
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `page`: 페이지 번호 (default: 1)
- `limit`: 페이지 크기 (default: 20)
- `music_id`: 음악 ID 필터
- `stage_id`: 스테이지 ID 필터

**Response:**
```json
{
    "data": [
        {
            "id": "record_uuid",
            "music": {
                "id": "music_uuid",
                "name": "곡 제목",
                "artist": "아티스트",
                "jacket_source": "/path/to/jacket.jpg"
            },
            "stage": {
                "id": "stage_uuid",
                "level_name": "Hard",
                "difficulty": 7
            },
            "score": 950000,
            "rank": "A",
            "accuracy": 95.5,
            "max_combo": 195,
            "is_full_combo": false,
            "is_perfect_play": false,
            "played_at": "2024-01-01T00:00:00Z"
        }
    ],
    "pagination": {
        "page": 1,
        "limit": 20,
        "total": 50,
        "total_pages": 3
    }
}
```

#### 3.3.3 리더보드 조회
```http
GET /api/v1/records/leaderboard
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `music_id`: 음악 ID (required)
- `stage_id`: 스테이지 ID (required)
- `limit`: 조회할 기록 수 (default: 10, max: 100)

**Response:**
```json
{
    "data": [
        {
            "rank": 1,
            "user": {
                "id": "user_uuid",
                "nickname": "Player1",
                "level": 50
            },
            "score": 999000,
            "accuracy": 99.8,
            "max_combo": 500,
            "rank_grade": "SSS",
            "is_full_combo": true,
            "is_perfect_play": true,
            "played_at": "2024-01-01T00:00:00Z"
        }
    ]
}
```

### 3.4 업적 관련 API

#### 3.4.1 업적 목록 조회
```http
GET /api/v1/achievements
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `type`: 업적 타입 필터 (score, combo, accuracy, play_count, special)
- `achieved`: 달성 여부 필터 (true/false)

**Response:**
```json
{
    "data": [
        {
            "id": "achievement_uuid",
            "name": "첫 걸음",
            "description": "첫 번째 곡을 플레이하세요",
            "icon_url": "/path/to/icon.png",
            "type": "play_count",
            "points": 10,
            "is_hidden": false,
            "is_achieved": true,
            "achieved_at": "2024-01-01T00:00:00Z"
        }
    ]
}
```

#### 3.4.2 내 업적 조회
```http
GET /api/v1/achievements/me
Authorization: Bearer {access_token}
```

**Response:**
```json
{
    "total_achievements": 100,
    "achieved_count": 25,
    "total_points": 1500,
    "achievements": [
        {
            "id": "achievement_uuid",
            "name": "첫 걸음",
            "description": "첫 번째 곡을 플레이하세요",
            "icon_url": "/path/to/icon.png",
            "type": "play_count",
            "points": 10,
            "achieved_at": "2024-01-01T00:00:00Z"
        }
    ]
}
```

### 3.5 상점 관련 API

#### 3.5.1 상품 목록 조회
```http
GET /api/v1/products
Authorization: Bearer {access_token}
```

**Query Parameters:**
- `type`: 상품 타입 필터 (hat, cane, clothes, character)
- `page`: 페이지 번호 (default: 1)
- `limit`: 페이지 크기 (default: 20)

**Response:**
```json
{
    "data": [
        {
            "id": "product_uuid",
            "name": "마법사 모자",
            "description": "신비한 마법사 모자",
            "price": 1000.0,
            "type": "hat",
            "item": {
                "id": "item_uuid",
                "name": "마법사 모자",
                "source": "/path/to/hat.png",
                "effect_id": "magic_boost"
            },
            "is_purchased": false
        }
    ],
    "pagination": {
        "page": 1,
        "limit": 20,
        "total": 50,
        "total_pages": 3
    }
}
```

#### 3.5.2 상품 구매
```http
POST /api/v1/products/{product_id}/purchase
Authorization: Bearer {access_token}
```

**Response:**
```json
{
    "purchase_id": "purchase_uuid",
    "product": {
        "id": "product_uuid",
        "name": "마법사 모자",
        "type": "hat"
    },
    "price": 1000.0,
    "purchased_at": "2024-01-01T00:00:00Z",
    "remaining_coin": 500,
    "remaining_gem": 100
}
```

### 3.6 캐릭터 관련 API

#### 3.6.1 캐릭터 목록 조회
```http
GET /api/v1/characters
Authorization: Bearer {access_token}
```

**Response:**
```json
{
    "data": [
        {
            "id": "character_uuid",
            "name": "마녀 루나",
            "description": "달빛을 다루는 마녀",
            "source": "/path/to/character.png",
            "is_owned": true,
            "is_default": true
        }
    ]
}
```

### 3.7 사용자 관리 API

#### 3.7.1 닉네임 변경
```http
PUT /api/v1/users/me/nickname
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
    "nickname": "new_nickname"
}
```

**Response:**
```json
{
    "message": "닉네임이 성공적으로 변경되었습니다",
    "nickname": "new_nickname"
}
```

#### 3.7.2 게임 설정 업데이트
```http
PUT /api/v1/users/me/settings
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
    "settings": {
        "volume": 0.8,
        "effect_volume": 0.6,
        "auto_play": false,
        "note_speed": 5
    }
}
```

**Response:**
```json
{
    "message": "설정이 성공적으로 업데이트되었습니다",
    "settings": {
        "volume": 0.8,
        "effect_volume": 0.6,
        "auto_play": false,
        "note_speed": 5
    }
}
```

#### 3.7.3 게임 데이터 저장
```http
PUT /api/v1/users/me/save-data
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
    "save_data": {
        "tutorial_completed": true,
        "story_progress": 10,
        "unlocked_features": ["ranking", "achievements"]
    }
}
```

**Response:**
```json
{
    "message": "게임 데이터가 성공적으로 저장되었습니다"
}
```

## 4. 에러 코드

### 4.1 HTTP 상태 코드
- `200 OK`: 요청 성공
- `201 Created`: 리소스 생성 성공
- `400 Bad Request`: 잘못된 요청
- `401 Unauthorized`: 인증 실패
- `403 Forbidden`: 권한 없음
- `404 Not Found`: 리소스를 찾을 수 없음
- `409 Conflict`: 리소스 충돌 (예: 중복된 닉네임)
- `422 Unprocessable Entity`: 요청 데이터 검증 실패
- `500 Internal Server Error`: 서버 내부 오류

### 4.2 에러 응답 형식
```json
{
    "error": "error_code",
    "message": "사용자 친화적인 에러 메시지",
    "details": {
        "field": "validation_error_details"
    }
}
```

## 5. 인증 및 보안

### 5.1 JWT 토큰 구조
```json
{
    "sub": "user_uuid",
    "platform_type": "steam",
    "platform_user_id": "steam_user_id",
    "exp": 1640995200,
    "iat": 1640908800
}
```

### 5.2 요청 헤더
```http
Authorization: Bearer {jwt_token}
Content-Type: application/json
```

## 6. 페이지네이션

### 6.1 기본 형식
```json
{
    "data": [...],
    "pagination": {
        "page": 1,
        "limit": 20,
        "total": 100,
        "total_pages": 5,
        "has_next": true,
        "has_prev": false
    }
}
```
