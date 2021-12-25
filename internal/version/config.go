package version

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	HttpClient   *http.Client
	RedisOptions *redis.Options
	List         []Setting
}

type Setting struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Regexp string `json:"regexp"`
}

func NewConfig() (*Config, error) {
	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT_SEC"))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(os.Getenv("SETTING_PATH"))
	if err != nil {
		return nil, err
	}

	list := []Setting{}
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	return &Config{
		HttpClient: &http.Client{
			Timeout: time.Second * time.Duration(httpTimeout),
		},
		RedisOptions: &redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
		List: list,
	}, nil
}
