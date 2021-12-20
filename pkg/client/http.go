package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Http struct {
	cli *http.Client
}

func NewHttp(cli *http.Client) *Http {
	return &Http{
		cli,
	}
}

func (h *Http) GetStruct(url string, value interface{}) error {
	resp, err := h.cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(value)
}

func (h *Http) GetByte(url string) ([]byte, error) {
	resp, err := h.cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
