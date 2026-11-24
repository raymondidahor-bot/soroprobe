package abi_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/abi"
)

// TestTypeJSONRoundTrip verifies that a Type value survives a JSON
// marshal/unmarshal cycle, including nested pointer fields and omitempty
// fields. The field values are chosen to exercise every JSON-tagged field.
func TestTypeJSONRoundTrip(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ty   abi.Type
	}{
		{
			name: "empty type round-trips",
			ty:   abi.Type{},
		},
		{
			name: "scalar type round-trips",
			ty:   abi.Type{Kind: "u32", Display: "u32"},
		},
		{
			name: "nested type round-trips",
			ty: abi.Type{
				Kind:    "map",
				Display: "Map<String, Tuple<BytesN<32>, Result<u64, bool>>>",
				Key: &abi.Type{Kind: "string", Display: "String"},
				Value: &abi.Type{
					Kind:    "tuple",
					Display: "Tuple<BytesN<32>, Result<u64, bool>>",
					Elements: []abi.Type{
						{Kind: "bytes_n", Display: "BytesN<32>", Length: 32},
						{Kind: "result", Display: "Result<u64, bool>",
							Ok:  &abi.Type{Kind: "u64", Display: "u64"},
							Err: &abi.Type{Kind: "bool", Display: "bool"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tt.ty)
			require.NoError(t, err)

			var got abi.Type
			require.NoError(t, json.Unmarshal(data, &got))
			assert.Equal(t, tt.ty, got)
		})
	}
}
