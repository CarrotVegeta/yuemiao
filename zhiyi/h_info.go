package zhiyi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HInfo struct {
	ID uint
	Lat string
	Lng string
}

func NewHInfo(id uint,lat,lng string) *HInfo  {
	return &HInfo{
		ID:  id,
		Lat: lat,
		Lng: lng,
	}
}

func (h *HInfo) GetInfo() interface{} {
	m := map[string]string{
		"act":"CustomerProduct",
		"id":fmt.Sprintf("%d",h.ID),
		"lat":h.Lat,
		"lng":h.Lng,
	}
	respBody := ZhiyiGet(m,nil)
	type Resp struct {
		Status int `json:"status"`
		List []struct{
			ID uint `json:"id"`
			Text string `json:"text"`
		}
	}
	r := &Resp{}
	_ = json.Unmarshal(respBody,&r)
	if r.Status != http.StatusOK {
		log.Fatal(string(respBody))
	}
	return r.List
}
