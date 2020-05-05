package elastic

import (
	"encoding/json"
	"testing"
)

func TestMissingFilter(t *testing.T) {
	f := NewMissingFilter("user").FilterName("_my_filter")
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"missing":{"_name":"_my_filter","field":"user"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
