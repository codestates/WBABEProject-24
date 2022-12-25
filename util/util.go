package util

import (
	"fmt"
	"time"
)

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
