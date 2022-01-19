package zhiyi

import (
	"github.com:CarrotVegeta/yuemiao/request"
	"log"
)

func ZhiyiGet(urlParam,BodyParam map[string]string) []byte {
	g := request.NewGet(BaseURL+Info, request.WithUrlParam(urlParam),request.WithBodyParam(BodyParam))
	g.SetHeader("Referer", Referer)
	g.SetHeader("Cookie", Cookie)
	respBody, err := g.Get()
	if err != nil {
		log.Fatal(err.Error())
	}
	return respBody
}
