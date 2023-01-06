package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var GlobalCache *cache.Cache

func InitCache() {
	GlobalCache = cache.New(4*time.Hour, 5*time.Hour)
}
