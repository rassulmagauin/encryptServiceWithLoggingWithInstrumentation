package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentationMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           EncryptService
}

func (mw InstrumentationMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "encrypt", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(float64(time.Since(begin).Seconds()))
	}(time.Now())
	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw InstrumentationMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "decrypt", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
