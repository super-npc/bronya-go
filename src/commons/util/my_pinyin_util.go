package util

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func ToPinyin(zh string) string {
	lazyPinyin := pinyin.LazyPinyin(zh, pinyin.NewArgs())
	return strings.Join(lazyPinyin, "_")
}
