package util

import (
	"fmt"
	"time"
)

/*
단순하게 true, false bool 값을 리턴하는 것을 굳이 유틸함수로 만들 필요가 있나요?
--------------------
true/false 값을 갖는 bool 포인터가 필요한 경우에 사용하고자 만든 함수입니다.
상수 true/false를 bool 포인터로 만드는 깔끔한 방법을 찾지 못했습니다...
*/
func NewBool(b bool) *bool {
	return &b
}

func CreateSeqStr(count uint32) string {
	return fmt.Sprintf("%d-%010d", time.Now().Unix(), count)
}
