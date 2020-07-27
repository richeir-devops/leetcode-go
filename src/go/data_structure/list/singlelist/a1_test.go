package singlelist

import "testing"

// Test1 11
func Test1(t *testing.T) {
	l1 := New()
	if l1 != nil {
		t.Log("123")
	}
}
