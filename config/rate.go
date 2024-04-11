package config

import "golang.org/x/time/rate"

var Limiter *rate.Limiter

func InitRate() {
	Limiter = rate.NewLimiter(rate.Limit(CustomConfig.RateConfig.Limit), CustomConfig.RateConfig.Burst)
}
