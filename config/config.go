package config

import (
	"time"
)

var JWTSecret = "jwt_secret_key"
var TokenExpiryDuration = 24 * time.Hour
