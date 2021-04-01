package utils

import (
	"testing"
)

func test(got, expect bool, t *testing.T) {
	if got != expect {
		t.Errorf("want true, got %t.", got)
	}
}

func TestContains(t *testing.T) {
	test(Contains([]string{"123", "234"}, "123"), true, t)
	test(Contains([]string{"222", "234"}, "123"), false, t)
}
