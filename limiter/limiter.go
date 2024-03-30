package limiter

import (
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
	"math"
	"net/http"
	"telegraph/config"
)

type RateLimiter struct {
	throttled.HTTPRateLimiterCtx
}

func (r *RateLimiter) RateLimitFunc(fn func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return r.RateLimit(http.HandlerFunc(fn))
}

func GetRateLimiter() *RateLimiter {
	store, err := memstore.NewCtx(65535)
	if err != nil {
		panic(err)
	}

	maxConn := 100
	if config.Conf.MaxConn > 0 {
		maxConn = config.Conf.MaxConn
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(maxConn),
		MaxBurst: int(math.Floor(float64(maxConn / 10))),
	}
	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)
	if err != nil {
		panic(err)
	}

	limiter := throttled.HTTPRateLimiterCtx{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	return &RateLimiter{limiter}
}
