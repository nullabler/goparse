package version

import (
	"context"
	"encoding/json"
	"log"
	"regexp"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/unixoff/goparse/pkg/client"
)

type Collect struct {
	ctx  context.Context
	wg   sync.WaitGroup
	conf *Config
	http *client.Http
	rcli *redis.Client
}

func NewCollect(conf *Config) *Collect {
	return &Collect{
		ctx:  context.Background(),
		conf: conf,
		http: client.NewHttp(conf.HttpClient),
		rcli: redis.NewClient(conf.RedisOptions),
	}
}

func (c *Collect) Run() {
	for _, item := range c.conf.List {
		c.wg.Add(1)
		go func(setting Setting) {
			defer c.wg.Done()
			bodyByte, err := c.http.GetByte(setting.Url)
			if err != nil {
				log.Fatalln("HttClient:", err)
			}

			body := string(bodyByte)
			reg := regexp.MustCompile(setting.Regexp)
			list := reg.FindAllStringSubmatch(body, -1)

			val, err := json.Marshal(list)
			if err != nil {
				log.Fatalln("Encode:", err)
			}

			if err := c.rcli.Set(c.ctx, setting.Id, string(val), 0).Err(); err != nil {
				log.Fatalln(err)
			}
			log.Println("Save:", setting.Id)
		}(item)
	}
	c.wg.Wait()
}
