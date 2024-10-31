package interfaces

import "time"

type RedisService interface {
	BlacklistToken(token string, expiration time.Duration) error
	IsBlacklistedToken(token string) (bool, error)
}
