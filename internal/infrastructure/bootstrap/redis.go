package bootstrap

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	redisClient "github.com/witchs-lounge_backend/internal/infrastructure/redis"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
)

// SetupRedis Redis 연결 초기화 및 세션 스토어 생성
func SetupRedis(addr string, password string, db int) (*redis.Client, session.SessionStore, error) {
	log.Printf("Redis 연결 중: %s", addr)

	// Redis 클라이언트 생성
	client := redisClient.NewRedisClient(addr, password, db)

	// Redis 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis 연결 실패: %v", err)
		log.Printf("세션 기능이 작동하지 않을 수 있습니다.")
		return client, nil, err
	}

	log.Printf("Redis 연결 성공!")

	// 세션 스토어 초기화
	sessionStore := session.NewRedisSessionStore(client, 24*time.Hour)

	return client, sessionStore, nil
}
