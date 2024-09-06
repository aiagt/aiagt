package db

import (
	"fmt"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	t.Log(strings.ReplaceAll(fmt.Sprint(arr), " ", ","))
}
