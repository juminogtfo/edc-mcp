package runtime

import (
	"testing"
)

func TestHealthIncrementsCallCount(t *testing.T) {
	before := callCount
	a := Health()
	b := Health()
	if a.CallCount != before+1 {
		t.Fatalf("first call: got %d, want %d", a.CallCount, before+1)
	}
	if b.CallCount != before+2 {
		t.Fatalf("second call: got %d, want %d", b.CallCount, before+2)
	}
	if a.Status != "ok" || b.Status != "ok" {
		t.Fatal("expected status ok")
	}
}