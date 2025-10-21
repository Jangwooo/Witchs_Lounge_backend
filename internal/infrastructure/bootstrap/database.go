package bootstrap

import (
	"context"
	"log"

	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/internal/infrastructure/database"
)

// SetupDatabase 데이터베이스 연결 초기화
func SetupDatabase(mode string) (*ent.Client, error) {
	// Load database configuration
	dbConfig, err := database.LoadConfig(&mode)
	if err != nil {
		return nil, err
	}

	// Initialize database connection
	client, err := database.NewEntClient(dbConfig)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("데이터베이스 마이그레이션 실패: %v", err)
	}

	log.Printf("✅ 데이터베이스 연결 성공 (모드: %s)", mode)
	return client, nil
}
