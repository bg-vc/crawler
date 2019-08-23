package service

import (
	"regexp"
	"strings"
	"testing"
)

func TestAAA(t *testing.T) {
	compileNum = regexp.MustCompile("[0-9]")
	kilometer := strings.Join(compileNum.FindAllString("3333", -1), "")
	year := strings.Join(compileNum.FindAllString("2018", -1), "")
	t.Logf("kilometer:%v, year:%v\n", kilometer, year)
}
