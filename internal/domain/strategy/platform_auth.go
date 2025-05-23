package strategy

import (
	"context"

	"github.com/witchs-lounge_backend/internal/domain/entity"
)

// PlatformAuthStrategy 플랫폼별 인증 전략 인터페이스
type PlatformAuthStrategy interface {
	// GetPlatformType 플랫폼 타입 반환
	GetPlatformType() string

	// VerifyTicket 티켓을 검증하고 플랫폼 사용자 정보 반환
	VerifyTicket(ctx context.Context, appID, ticket string) (*entity.PlatformUserInfo, error)

	// ExtractUserID 티켓에서 사용자 ID 추출 (간단한 검증용)
	ExtractUserID(ticket string) string
}

// PlatformAuthFactory 플랫폼별 인증 전략 팩토리
type PlatformAuthFactory interface {
	GetStrategy(platformType string) (PlatformAuthStrategy, error)
}
