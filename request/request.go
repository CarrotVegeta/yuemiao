package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Get struct {
	Url       string
	header    http.Header
	BodyParam map[string]string
	UrlParam  map[string]string
}
type Option func(get *Get)

func NewGet(url string, options ...Option) *Get {
	g := &Get{
		Url:       url,
		header:    make(http.Header),
		UrlParam:  make(map[string]string),
		BodyParam: make(map[string]string),
	}
	for _, o := range options {
		o(g)
	}
	return g
}
func WithBodyParam(p map[string]string) Option {
	return func(get *Get) {
		get.BodyParam = p
	}
}
func WithUrlParam(p map[string]string) Option {
	return func(get *Get) {
		get.UrlParam = p
	}
}
func (g *Get) SetUrlParam(k, v string) {
	g.UrlParam[k] = v
}
func (g *Get) SetHeader(k, v string) {
	g.header.Set(k, v)
}
func (g *Get) Get() ([]byte, error) {
	bs, _ := json.Marshal(g.BodyParam)
	if len(g.UrlParam) > 0 {
		g.Url += "?"
		for k, v := range g.UrlParam {
			g.Url += k + "=" + v + "&"
		}
	}
	g.Url = strings.TrimSuffix(g.Url, "&")
	client := &http.Client{}
	request, err := http.NewRequest("GET", g.Url, strings.NewReader(string(bs)))
	if err != nil {
		return nil, err
	}
	request.Header = g.header
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bs, nil
}
