package abi_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/soroworks/soroprobe/internal/abi"
)

func typePtr(t abi.Type) *abi.Type { return &t }

func TestTypeJSONRoundTrip(t *testing.T) {
	t.Parallel()

	want := abi.Type{
		Kind:     "result",
		Display:  "Result<Vec<Map<Symbol, u32>>, Error>",
		Inner:    typePtr(scalar("option", "Option<u32>")),
		Element:  typePtr(scalar("vec", "Vec<u32>")),
		Key:      typePtr(scalar("symbol", "Symbol")),
		Value:    typePtr(scalar("u32", "u32")),
		Elements: []abi.Type{scalar("u32", "u32"), scalar("string", "String")},
		Ok:       typePtr(scalar("bool", "bool")),
		Err:      typePtr(scalar("i128", "i128")),
		Length:   32,
	}

	data, err := json.Marshal(want)
	require.NoError(t, err)

	var got abi.Type
	require.NoError(t, json.Unmarshal(data, &got))
	assert.Equal(t, want, got)
}

func TestUDTJSONRoundTrip(t *testing.T) {
	t.Parallel()

	want := abi.UDT{
		Kind: abi.UDTUnion,
		Name: "Config",
		Fields: []abi.Param{
			{Name: "owner", Type: scalar("address", "Address")},
			{Name: "fee", Type: scalar("u32", "u32")},
		},
		Cases: []abi.Case{
			{Name: "None"},
			{Name: "Value", Value: 7},
			{Name: "Pair", Payload: []abi.Type{
				scalar("u32", "u32"),
				scalar("symbol", "Symbol"),
			}},
		},
	}

	data, err := json.Marshal(want)
	require.NoError(t, err)

	var got abi.UDT
	require.NoError(t, json.Unmarshal(data, &got))
	assert.Equal(t, want, got)
}
