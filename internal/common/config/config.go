package config

import "time"

type Config struct {
	Cache *CacheConfig
}

type CacheConfig struct {
	CleanupInterval   time.Duration
	DefaultExpiration time.Duration
	TokenExpiration   time.Duration
}
