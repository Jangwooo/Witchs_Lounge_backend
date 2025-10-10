package strategy

import (
	"fmt"

	"github.com/witchs-lounge_backend/internal/domain/strategy"
)

// PlatformAuthFactoryImpl 플랫폼 인증 전략 팩토리 구현체
type PlatformAuthFactoryImpl struct{}

// NewPlatformAuthFactory 팩토리 생성자
func NewPlatformAuthFactory() strategy.PlatformAuthFactory {
	return &PlatformAuthFactoryImpl{}
}

// GetStrategy 플랫폼 타입에 따른 인증 전략 반환
func (f *PlatformAuthFactoryImpl) GetStrategy(platformType string) (strategy.PlatformAuthStrategy, error) {
	switch platformType {
	case "steam":
		return NewSteamAuthStrategy(), nil
	case "stove":
		return NewStoveAuthStrategy(), nil
	default:
		return nil, fmt.Errorf("지원되지 않는 플랫폼입니다: %s", platformType)
	}
}
