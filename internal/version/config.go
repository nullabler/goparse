package version

import (
	"os"
	"strconv"
)

type Config struct {
	HttpTimeoutSec int
}

func NewConfig() (*Config, error) {
	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT_SEC"))
	if err != nil {
		return nil, err
	}

	return &Config{
		HttpTimeoutSec: httpTimeout,
	}, nil
}
