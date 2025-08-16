package config

import (
	"fmt"
	"net/url"
	"os"
)

func RedisURL() (string, error) {
	if redisURL := os.Getenv("REDIS_URL"); redisURL != "" {
		return redisURL, nil
	}

	host := getEnv("REDIS_HOST", "localhost")
	port := getEnv("REDIS_PORT", "6379")
	pass := os.Getenv("REDIS_PASSWORD")
	user := os.Getenv("REDIS_USERNAME") // optional (ACL user)
	db := getEnv("REDIS_DB", "0")
	scheme := getEnv("REDIS_SCHEME", "redis") // can be "redis" or "rediss"

	u := &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   db,
	}

	if user != "" || pass != "" {
		u.User = url.UserPassword(user, pass)
	}

	// optional query params
	q := u.Query()
	if v := os.Getenv("REDIS_DIAL_TIMEOUT"); v != "" {
		q.Set("dial_timeout", v)
	}
	if v := os.Getenv("REDIS_READ_TIMEOUT"); v != "" {
		q.Set("read_timeout", v)
	}
	if v := os.Getenv("REDIS_WRITE_TIMEOUT"); v != "" {
		q.Set("write_timeout", v)
	}
	if v := os.Getenv("REDIS_POOL_SIZE"); v != "" {
		q.Set("pool_size", v)
	}
	if v := os.Getenv("REDIS_MIN_IDLE_CONNS"); v != "" {
		q.Set("min_idle_conns", v)
	}
	if v := os.Getenv("REDIS_MAX_RETRIES"); v != "" {
		q.Set("max_retries", v)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
