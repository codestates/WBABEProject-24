package util

import (
	"fmt"
	"time"
)

/*
단순하게 true, false bool 값을 리턴하는 것을 굳이 유틸함수로 만들 필요가 있나요?
*/
func NewTrue() *bool {
	b := true
	return &b
}

func NewFalse() *bool {
	b := false
	return &b
}

func CreateSeqStr(count uint32) string {
	return fmt.Sprintf("%d-%010d", time.Now().Unix(), count)
}
