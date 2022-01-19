package zhiyi

import (
	"fmt"
	"testing"
)

func TestGetHInfo(t *testing.T)  {
	h := NewHInfo(2560,"30.57447","103.92377")
	temp := h.GetInfo()
	fmt.Println(temp)

}
