package abi_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/abi"
)

// TestUDTJSONRoundTrip verifies that a UDT value survives a JSON
// marshal/unmarshal cycle, including its Fields and Cases slices.
func TestUDTJSONRoundTrip(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		udt  abi.UDT
	}{
		{
			name: "empty udt round-trips",
			udt:  abi.UDT{},
		},
		{
			name: "struct udt round-trips",
			udt: abi.UDT{
				Kind: abi.UDTStruct,
				Name: "Config",
				Fields: []abi.Param{
					{Name: "admin", Type: abi.Type{Kind: "address", Display: "Address"}},
					{Name: "fee", Type: abi.Type{Kind: "u32", Display: "u32"}},
				},
			},
		},
		{
			name: "enum udt round-trips",
			udt: abi.UDT{
				Kind: abi.UDTEnum,
				Name: "Level",
				Cases: []abi.Case{
					{Name: "Low", Value: 1},
					{Name: "High", Value: 2},
				},
			},
		},
		{
			name: "union udt with payload round-trips",
			udt: abi.UDT{
				Kind: abi.UDTUnion,
				Name: "Key",
				Cases: []abi.Case{
					{Name: "Admin"},
					{Name: "Balance", Payload: []abi.Type{{Kind: "u64", Display: "u64"}}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tt.udt)
			require.NoError(t, err)

			var got abi.UDT
			require.NoError(t, json.Unmarshal(data, &got))
			assert.Equal(t, tt.udt, got)
		})
	}
}
