package strategy

import (
	"context"
	"fmt"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/strategy"
)

// SteamAuthStrategy Steam 플랫폼 인증 전략
type SteamAuthStrategy struct{}

// NewSteamAuthStrategy Steam 인증 전략 생성자
func NewSteamAuthStrategy() strategy.PlatformAuthStrategy {
	return &SteamAuthStrategy{}
}

// GetPlatformType Steam 플랫폼 타입 반환
func (s *SteamAuthStrategy) GetPlatformType() string {
	return "steam"
}

// VerifyTicket Steam 티켓 검증 (개발 단계에서는 간단히 처리)
func (s *SteamAuthStrategy) VerifyTicket(ctx context.Context, appID, ticket string) (*entity.PlatformUserInfo, error) {
	// 실제로는 Steam Web API를 호출하여 티켓을 검증해야 함
	// 개발 단계에서는 티켓을 steam_id로 간주
	steamID := ticket

	// 간단한 검증 로직 (실제로는 Steam API 호출)
	if steamID == "" {
		return nil, fmt.Errorf("Steam 티켓이 비어있습니다")
	}

	// Steam 사용자 정보 반환 (실제로는 Steam API에서 가져와야 함)
	return &entity.PlatformUserInfo{
		PlatformType:   "steam",
		PlatformUserID: steamID,
		Email:          "",                          // Steam에서는 이메일을 제공하지 않음
		AvatarURL:      "",                          // 기본값, 실제로는 Steam API에서 가져옴
		DisplayName:    "Steam User " + steamID[:5], // 기본 닉네임
		Language:       "ko",
		PlatformData: map[string]interface{}{
			"steam_id": steamID,
			"app_id":   appID,
		},
		IsVerified: true, // Steam 티켓이 유효하면 인증됨
	}, nil
}

// ExtractUserID 티켓에서 사용자 ID 추출
func (s *SteamAuthStrategy) ExtractUserID(ticket string) string {
	// 개발 단계에서는 티켓 자체를 steam_id로 사용
	return ticket
}
