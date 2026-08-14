package main

import (
	"strings"
	"testing"
)

// TestAddNodeVirtualNodesUpperBound verifies that AddNode rejects
// excessively large virtualNodes values that could exhaust memory.
func TestAddNodeVirtualNodesUpperBound(t *testing.T) {
	s := NewService()

	_, err := s.AddNode("huge", 100000)
	if err == nil {
		t.Fatal("expected error for virtualNodes > reasonable limit, got nil")
	}
	if !strings.Contains(err.Error(), "too large") && !strings.Contains(err.Error(), "exceed") && !strings.Contains(err.Error(), "max") {
		t.Fatalf("error should indicate limit exceeded, got: %v", err)
	}

	// Normal values should still work
	_, err = s.AddNode("normal", 100)
	if err != nil {
		t.Fatalf("reasonable virtualNodes should succeed: %v", err)
	}
}
