package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/strategy"
)

// StoveAuthStrategy STOVE 플랫폼 인증 전략
type StoveAuthStrategy struct {
	httpClient *http.Client
}

// StoveTokenValidationResponse STOVE 토큰 검증 API 응답 구조체
type StoveTokenValidationResponse struct {
	Result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"result"`
	Member struct {
		MemberNo   string `json:"memberNo"`
		GameUserNo string `json:"gameUserNo"`
	} `json:"member"`
}

// NewStoveAuthStrategy STOVE 인증 전략 생성자
func NewStoveAuthStrategy() strategy.PlatformAuthStrategy {
	return &StoveAuthStrategy{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetPlatformType STOVE 플랫폼 타입 반환
func (s *StoveAuthStrategy) GetPlatformType() string {
	return "stove"
}

// VerifyTicket STOVE access token 검증
func (s *StoveAuthStrategy) VerifyTicket(ctx context.Context, appID, ticket string) (*entity.PlatformUserInfo, error) {
	// STOVE의 경우 ticket은 access token을 의미함
	accessToken := ticket

	if accessToken == "" {
		return nil, fmt.Errorf("stove access token이 비어있습니다")
	}

	// STOVE API를 통해 토큰 유효성 검증
	userInfo, err := s.validateAccessToken(ctx, appID, accessToken)
	if err != nil {
		return nil, fmt.Errorf("stove 토큰 검증 실패: %w", err)
	}

	return userInfo, nil
}

// validateAccessToken STOVE API를 호출하여 access token 유효성 검증
func (s *StoveAuthStrategy) validateAccessToken(ctx context.Context, appID, accessToken string) (*entity.PlatformUserInfo, error) {
	// STOVE Game User Access Token 유효성 검증 API
	// 실제 API URL은 STOVE 개발자 센터에서 제공하는 엔드포인트로 교체해야 함
	apiURL := "https://api.onstove.com/auth/v1/game/validate"

	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("http 요청 생성 실패: %w", err)
	}

	// STOVE API 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("X-Stove-App-Id", appID)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("stove API 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("stove API 오류 (상태 코드: %d): %s", resp.StatusCode, string(body))
	}

	// 응답 파싱
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 본문 읽기 실패: %w", err)
	}

	var validationResp StoveTokenValidationResponse
	if err := json.Unmarshal(body, &validationResp); err != nil {
		return nil, fmt.Errorf("응답 파싱 실패: %w", err)
	}

	// 응답 코드 확인
	if validationResp.Result.Code != 0 {
		return nil, fmt.Errorf("stove API 오류: %s (코드: %d)",
			validationResp.Result.Message,
			validationResp.Result.Code)
	}

	// 플랫폼 사용자 정보 생성
	return &entity.PlatformUserInfo{
		PlatformType:   "stove",
		PlatformUserID: validationResp.Member.MemberNo,
		Email:          "",                                                 // STOVE에서 이메일 정보가 제공되지 않는 경우
		AvatarURL:      "",                                                 // 기본값
		DisplayName:    "STOVE User " + validationResp.Member.MemberNo[:5], // 기본 닉네임
		Language:       "ko",
		PlatformData: map[string]interface{}{
			"member_no":    validationResp.Member.MemberNo,
			"game_user_no": validationResp.Member.GameUserNo,
			"app_id":       appID,
		},
		IsVerified: true, // STOVE 토큰이 유효하면 인증됨
	}, nil
}

// ExtractUserID 티켓(access token)에서 사용자 ID 추출
// STOVE의 경우 토큰을 디코딩하지 않고 API 호출이 필요하므로 빈 문자열 반환
func (s *StoveAuthStrategy) ExtractUserID(ticket string) string {
	// STOVE access token은 JWT가 아니므로 직접 추출 불가
	// API 호출을 통해서만 사용자 정보를 얻을 수 있음
	return ""
}
