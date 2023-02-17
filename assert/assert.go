package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, a T, b T) {
	if a != b {
		t.Fatalf("\nWant: %s\nGot: %s", a, b)
	}
}

func StringContains(t *testing.T, s string, substr string) {
	if !strings.Contains(s, substr) {
		t.Fatalf("\nString '%s' does not contain '%s'", s, substr)
	}
}

func SliceContains[T comparable](t *testing.T, slice []T, item T) {
	for _, v := range slice {
		if v == item {
			return
		}
	}

	t.Fatalf("\nSlice '%s' does not contain '%s'", slice, item)
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("\nWanted no error, but got '%s'\nMessage: '%s'", err, err.Error())
	}
}
