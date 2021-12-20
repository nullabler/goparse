package version

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/unixoff/goparse/pkg/client"
)

type Collect struct {
	conf *Config
	http *client.Http
}

func NewCollect(conf *Config) *Collect {
	return &Collect{
		conf: conf,
		http: client.NewHttp(&http.Client{
			Timeout: time.Second * time.Duration(conf.HttpTimeoutSec),
		}),
	}
}

func (c *Collect) Run() {
	bodyByte, err := c.http.GetByte("https://www.php.net/downloads.php")
	if err != nil {
		log.Fatalln(err)
	}

	body := string(bodyByte)

	reg := regexp.MustCompile(`Current\s+Stable[^>]+>\s+PHP\s+([\d\.]+)[^<]+<.+href="([^"]+)`)

	list := reg.FindAllStringSubmatch(body, -1)

	//list := reg.FindString(body)
	log.Println(list[0][1], list[0][2])

}
