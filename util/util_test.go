package util

import (
	"fmt"
	"testing"
)

func TestString2json(t *testing.T) {
	byteArr := String2json([]string{"hello", "yes", "extra"}, []string{"world", "no", "gg", "g"})
	fmt.Println(string(byteArr))
}
