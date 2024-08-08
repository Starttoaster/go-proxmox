package proxmox

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

type testIntOrStringExample struct {
	ID IntOrString `json:"id"`
}

// TestUnmarshalIntOrString_Positive tests the unmarshalling of valid IntOrString values.
func TestUnmarshalIntOrString_Positive(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`{"id": 12345}`, "12345"},
		{`{"id": "abc123"}`, "abc123"},
	}

	for _, test := range tests {
		var example testIntOrStringExample
		err := json.Unmarshal([]byte(test.input), &example)
		require.NoError(t, err)
		require.Equal(t, IntOrString(test.expected), example.ID)
	}
}

// TestUnmarshalIntOrString_Negative tests the unmarshalling of invalid IntOrString values.
func TestUnmarshalIntOrString_Negative(t *testing.T) {
	tests := []string{
		`{"id": {"nested": "object"}}`,
		`{"id": [12345]}`,
	}

	for _, test := range tests {
		var example testIntOrStringExample
		err := json.Unmarshal([]byte(test), &example)
		require.Error(t, err)
	}
}
