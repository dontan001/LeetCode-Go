package main

import "testing"

func TestInclude(t *testing.T) {
	t.Run("test true", func(t *testing.T) {
		result := include("hello", "lo")
		if result != true {
			t.Error("Expected true, got ", result)
		}
	})

	t.Run("test false", func(t *testing.T) {
		result := include("hello", "le")
		if result == true {
			t.Error("Expected false, got ", result)
		}
	})
}
