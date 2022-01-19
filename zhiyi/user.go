package zhiyi

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Birthday string `json:"birthday"`
	Tel      string `json:"tel"`
	Cname    string `json:"cname"`
	Sex      int    `json:"sex"`
	Idcard   string `json:"idcard"`
	Doctype  int    `json:"doctype"`
	Cookie   string `json:"cookie"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserInfo() {
	respBody := ZhiyiGet(map[string]string{"act":"User"},nil)
	type Resp struct {
		Status int   `json:"status"`
		U      []byte `json:"user"`
	}
	r := &Resp{}
	_ = json.Unmarshal(respBody, &r)
	if r.Status != http.StatusOK {
		log.Fatal(string(respBody))
	}
	_ = json.Unmarshal(r.U,&u)
}
