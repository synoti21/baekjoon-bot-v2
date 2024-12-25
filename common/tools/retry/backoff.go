package retry

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

const (
	initialInterval = 10 * time.Second
	maxInterval     = 5 * time.Second
	maxElapsedTime  = 1 * time.Minute
)

func NonRetryableError(err error) *backoff.PermanentError {
	return &backoff.PermanentError{
		Err: err,
	}
}

func Do(tryFunc func() error) error {
	b := backoff.NewExponentialBackOff()

	b.InitialInterval = initialInterval
	b.MaxInterval = maxInterval
	b.MaxElapsedTime = maxElapsedTime

	return backoff.Retry(tryFunc, b)
}
