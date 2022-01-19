package request

import (
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	g := NewGet("https://cloud.cn2030.com/sc/wx/HandlerSubscribe.ashx", WithUrlParam(map[string]string{"act": "User"}))
	g.SetHeader("Referer", "https://servicewechat.com/wx2c7f0f3c30d99445/92/page-frame.html")
	g.SetHeader("Cookie", "ASP.NET_SessionId=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2NDI1OTY4MjMuODE1NzY4MiwiZXhwIjoxNjQyNjAwNDIzLjgxNTc2ODIsInN1YiI6IllOVy5WSVAiLCJqdGkiOiIyMDIyMDExOTA4NTM0MyIsInZhbCI6IkRxb3JBUUlBQUFBUU16azBNR0V3TVdGaE1UaGpObUl3TXh4dmNYSTFielZOYmxZdFZteHpUV2gwYmpSeVh6UXdNekJQV2toekFCeHZcclxuVlRJMldIUXhMVmxCVUd0U2JYVm5lVk5DZHpWVk1EZEdZVVpyRERZMExqWTBMakkwTXk0NU1nQUFBQUFBQUFBPSJ9.xMPkL2WcLFX-P_9ISxlrDHJcNbZh6vBELKpgSXWT8BI")
	respBody, err := g.Get()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(respBody))
}
