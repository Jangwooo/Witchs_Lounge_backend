package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/strategy"
)

// SteamAuthStrategy Steam 플랫폼 인증 전략
type SteamAuthStrategy struct {
	httpClient *http.Client
	apiKey     string
}

// SteamAuthTicketResponse Steam 티켓 검증 API 응답
type SteamAuthTicketResponse struct {
	Response struct {
		Params struct {
			Result          string `json:"result"`
			SteamID         string `json:"steamid"`
			OwnerSteamID    string `json:"ownersteamid"`
			VacBanned       bool   `json:"vacbanned"`
			PublisherBanned bool   `json:"publisherbanned"`
		} `json:"params"`
		Error *struct {
			ErrorCode int    `json:"errorcode"`
			ErrorDesc string `json:"errordesc"`
		} `json:"error"`
	} `json:"response"`
}

// SteamPlayerSummariesResponse Steam 사용자 정보 API 응답
type SteamPlayerSummariesResponse struct {
	Response struct {
		Players []struct {
			SteamID                  string `json:"steamid"`
			CommunityVisibilityState int    `json:"communityvisibilitystate"`
			ProfileState             int    `json:"profilestate"`
			PersonaName              string `json:"personaname"`
			ProfileURL               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			AvatarMedium             string `json:"avatarmedium"`
			AvatarFull               string `json:"avatarfull"`
			PersonaState             int    `json:"personastate"`
			RealName                 string `json:"realname"`
			TimeCreated              int64  `json:"timecreated"`
			LocCountryCode           string `json:"loccountrycode"`
		} `json:"players"`
	} `json:"response"`
}

// NewSteamAuthStrategy Steam 인증 전략 생성자
func NewSteamAuthStrategy() strategy.PlatformAuthStrategy {
	apiKey := os.Getenv("STEAM_WEB_API_KEY")
	if apiKey == "" {
		// 개발 환경에서 API 키가 없으면 경고만 출력
		fmt.Println("⚠️  경고: STEAM_WEB_API_KEY 환경 변수가 설정되지 않았습니다")
	}

	return &SteamAuthStrategy{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
	}
}

// GetPlatformType Steam 플랫폼 타입 반환
func (s *SteamAuthStrategy) GetPlatformType() string {
	return "steam"
}

// VerifyTicket Steam 티켓 검증
func (s *SteamAuthStrategy) VerifyTicket(ctx context.Context, appID, ticket string) (*entity.PlatformUserInfo, error) {
	if ticket == "" {
		return nil, fmt.Errorf("steam 티켓이 비어있습니다")
	}

	// API 키가 없으면 개발 모드로 동작
	if s.apiKey == "" {
		return s.devModeVerify(appID, ticket)
	}

	// 1. Steam Web API로 티켓 검증
	steamID, err := s.authenticateUserTicket(ctx, appID, ticket)
	if err != nil {
		return nil, fmt.Errorf("steam 티켓 검증 실패: %w", err)
	}

	// 2. Steam 사용자 정보 조회
	playerInfo, err := s.getPlayerSummaries(ctx, steamID)
	if err != nil {
		// 사용자 정보 조회 실패 시에도 기본 정보로 계속 진행
		fmt.Printf("⚠️  Steam 사용자 정보 조회 실패: %v\n", err)
		return s.createBasicPlatformUserInfo(steamID, appID), nil
	}

	// 3. 플랫폼 사용자 정보 생성
	return &entity.PlatformUserInfo{
		PlatformType:   "steam",
		PlatformUserID: steamID,
		Email:          "",                    // Steam에서는 이메일을 제공하지 않음
		AvatarURL:      playerInfo.AvatarFull, // 큰 아바타 이미지
		DisplayName:    playerInfo.PersonaName,
		Language:       s.mapCountryCodeToLanguage(playerInfo.LocCountryCode),
		PlatformData: map[string]interface{}{
			"steam_id":      steamID,
			"app_id":        appID,
			"profile_url":   playerInfo.ProfileURL,
			"avatar":        playerInfo.Avatar,
			"avatar_medium": playerInfo.AvatarMedium,
			"avatar_full":   playerInfo.AvatarFull,
			"real_name":     playerInfo.RealName,
			"country_code":  playerInfo.LocCountryCode,
			"time_created":  playerInfo.TimeCreated,
			"persona_state": playerInfo.PersonaState,
			"profile_state": playerInfo.ProfileState,
		},
		IsVerified: true,
	}, nil
}

// authenticateUserTicket Steam 티켓 검증 API 호출
func (s *SteamAuthStrategy) authenticateUserTicket(ctx context.Context, appID, ticket string) (string, error) {
	apiURL := "https://api.steampowered.com/ISteamUserAuth/AuthenticateUserTicket/v1/"

	params := url.Values{}
	params.Add("key", s.apiKey)
	params.Add("appid", appID)
	params.Add("ticket", ticket)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL+"?"+params.Encode(), nil)
	if err != nil {
		return "", fmt.Errorf("http 요청 생성 실패: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("steam API 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("응답 본문 읽기 실패: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("steam API 오류 (상태 코드: %d): %s", resp.StatusCode, string(body))
	}

	var authResp SteamAuthTicketResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", fmt.Errorf("응답 파싱 실패: %w", err)
	}

	// 에러 체크
	if authResp.Response.Error != nil {
		return "", fmt.Errorf("steam API 오류: %s (코드: %d)",
			authResp.Response.Error.ErrorDesc,
			authResp.Response.Error.ErrorCode)
	}

	// 티켓 검증 결과 확인
	if authResp.Response.Params.Result != "OK" {
		return "", fmt.Errorf("티켓 검증 실패: %s", authResp.Response.Params.Result)
	}

	return authResp.Response.Params.SteamID, nil
}

// getPlayerSummaries Steam 사용자 정보 조회
func (s *SteamAuthStrategy) getPlayerSummaries(ctx context.Context, steamID string) (*struct {
	SteamID        string
	PersonaName    string
	ProfileURL     string
	Avatar         string
	AvatarMedium   string
	AvatarFull     string
	RealName       string
	LocCountryCode string
	TimeCreated    int64
	PersonaState   int
	ProfileState   int
}, error) {
	apiURL := "https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/"

	params := url.Values{}
	params.Add("key", s.apiKey)
	params.Add("steamids", steamID)

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("http 요청 생성 실패: %w", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("steam API 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 본문 읽기 실패: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("steam API 오류 (상태 코드: %d): %s", resp.StatusCode, string(body))
	}

	var summariesResp SteamPlayerSummariesResponse
	if err := json.Unmarshal(body, &summariesResp); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	if len(summariesResp.Response.Players) == 0 {
		return nil, fmt.Errorf("사용자 정보를 찾을 수 없습니다")
	}

	player := summariesResp.Response.Players[0]
	return &struct {
		SteamID        string
		PersonaName    string
		ProfileURL     string
		Avatar         string
		AvatarMedium   string
		AvatarFull     string
		RealName       string
		LocCountryCode string
		TimeCreated    int64
		PersonaState   int
		ProfileState   int
	}{
		SteamID:        player.SteamID,
		PersonaName:    player.PersonaName,
		ProfileURL:     player.ProfileURL,
		Avatar:         player.Avatar,
		AvatarMedium:   player.AvatarMedium,
		AvatarFull:     player.AvatarFull,
		RealName:       player.RealName,
		LocCountryCode: player.LocCountryCode,
		TimeCreated:    player.TimeCreated,
		PersonaState:   player.PersonaState,
		ProfileState:   player.ProfileState,
	}, nil
}

// createBasicPlatformUserInfo 기본 플랫폼 사용자 정보 생성
func (s *SteamAuthStrategy) createBasicPlatformUserInfo(steamID, appID string) *entity.PlatformUserInfo {
	nickname := "Steam User"
	if len(steamID) > 5 {
		nickname = "Steam User " + steamID[:5]
	}

	return &entity.PlatformUserInfo{
		PlatformType:   "steam",
		PlatformUserID: steamID,
		Email:          "",
		AvatarURL:      "",
		DisplayName:    nickname,
		Language:       "ko",
		PlatformData: map[string]interface{}{
			"steam_id": steamID,
			"app_id":   appID,
		},
		IsVerified: true,
	}
}

// devModeVerify 개발 모드 티켓 검증 (API 키 없을 때)
func (s *SteamAuthStrategy) devModeVerify(appID, ticket string) (*entity.PlatformUserInfo, error) {
	fmt.Println("🔧 개발 모드: Steam API 키가 없어 티켓을 SteamID로 간주합니다")
	return s.createBasicPlatformUserInfo(ticket, appID), nil
}

// mapCountryCodeToLanguage 국가 코드를 언어로 매핑
func (s *SteamAuthStrategy) mapCountryCodeToLanguage(countryCode string) string {
	languageMap := map[string]string{
		"KR": "ko",
		"US": "en",
		"GB": "en",
		"JP": "ja",
		"CN": "zh",
		"TW": "zh",
	}

	if lang, ok := languageMap[countryCode]; ok {
		return lang
	}
	return "en" // 기본값
}

// ExtractUserID 티켓에서 사용자 ID 추출
func (s *SteamAuthStrategy) ExtractUserID(ticket string) string {
	// Steam 티켓은 복잡한 바이너리 데이터이므로 API 호출 필요
	// 개발 모드에서는 티켓을 steam_id로 간주
	return ""
}
