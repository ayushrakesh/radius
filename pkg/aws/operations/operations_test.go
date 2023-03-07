// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package operations

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wI2L/jsondiff"
)

func TestFlattenProperties(t *testing.T) {
	properties := map[string]any{
		"A": map[string]any{
			"B": map[string]any{
				"C": "D",
			},
			"E": "F",
		},
		"G": "H",
	}

	flattened := FlattenProperties(properties)
	require.Equal(t, map[string]any{
		"A/B/C": "D",
		"A/E":   "F",
		"G":     "H",
	}, flattened)
}

func TestUnflattenProperties(t *testing.T) {
	properties := map[string]any{
		"A/B/C": "D",
		"A/E":   "F",
		"G":     "H",
	}

	unflattened := UnflattenProperties(properties)
	require.Equal(t, map[string]any{
		"A": map[string]any{
			"B": map[string]any{
				"C": "D",
			},
			"E": "F",
		},
		"G": "H",
	}, unflattened)
}

func TestFlattenUnflattenInverses(t *testing.T) {
	properties := map[string]any{
		"A": map[string]any{
			"B": map[string]any{
				"C": "D",
			},
			"E": "F",
		},
		"G": "H",
	}

	flattened := FlattenProperties(properties)
	unflattened := UnflattenProperties(flattened)
	require.Equal(t, properties, unflattened)
}

func TestFlattenUnflattenRealData(t *testing.T) {
	properties := map[string]any{
		"ClusterEndpoint:": map[string]any{
			"Address": "https://A1B2C3D4E5F6.gr7.us-west-2.eks.amazonaws.com",
			"Port":    443,
		},
		"ClusterName": "my-cluster",
	}

	flattened := FlattenProperties(properties)
	unflattened := UnflattenProperties(flattened)
	require.Equal(t, properties, unflattened)
}

func Test_GeneratePatch(t *testing.T) {
	testCases := []struct {
		name          string
		currentState  map[string]any
		desiredState  map[string]any
		schema        map[string]any
		expectedPatch jsondiff.Patch
	}{
		{
			"No updates creates empty patch",
			map[string]any{
				"A": "B",
				"C": map[string]any{
					"D": map[string]any{
						"E": "F",
					},
					"G": map[string]any{
						"I": "J",
					},
					"K": "L",
				},
			},
			map[string]any{
				"A": "B",
				"C": map[string]any{
					"G": map[string]any{
						"I": "J",
					},
				},
			},
			map[string]any{
				"properties": map[string]any{
					"A": map[string]any{},
					"C": map[string]any{},
				},
				"readOnlyProperties": []any{
					"/properties/C/D/E",
				},
				"createOnlyProperties": []any{
					"/properties/C/K",
				},
			},
			nil,
		},
		{
			"Update creates patch",
			map[string]any{
				"A": "B",
				"C": map[string]any{
					"D": map[string]any{
						"E": "F",
					},
					"G": map[string]any{
						"I": "J",
					},
					"K": "L",
				},
			},
			map[string]any{
				"A": "Test",
				"C": map[string]any{
					"G": map[string]any{
						"I": "Test2",
					},
				},
			},
			map[string]any{
				"properties": map[string]any{
					"A": map[string]any{},
					"C": map[string]any{},
				},
				"readOnlyProperties": []any{
					"/properties/C/D/E",
				},
				"createOnlyProperties": []any{
					"/properties/C/K",
				},
			},
			jsondiff.Patch{
				{
					Type:     "replace",
					Path:     "/A",
					OldValue: "B",
					Value:    "Test",
				},
				{
					Type:     "replace",
					Path:     "/C/G/I",
					OldValue: "J",
					Value:    "Test2",
				},
			},
		},
		{
			"Specify create-only properties",
			map[string]any{
				"A": map[string]any{
					"B": map[string]any{
						"C": "D",
						"E": "F",
					},
				},
			},
			map[string]any{
				"A": map[string]any{
					"B": map[string]any{
						"C": "D",
						"E": "Test",
					},
				},
			},
			map[string]any{
				"properties": map[string]any{
					"A": map[string]any{},
				},
				"createOnlyProperties": []any{
					"/properties/A/B/C",
				},
			},
			jsondiff.Patch{
				{
					Type:     "replace",
					Path:     "/A/B/E",
					OldValue: "F",
					Value:    "Test",
				},
			},
		},
		{
			"Remove object",
			map[string]any{
				"A": map[string]any{
					"B": map[string]any{
						"C": "D",
						"E": "F",
					},
				},
				"G": "H",
			},
			map[string]any{
				"G": "H",
			},
			map[string]any{
				"properties": map[string]any{
					"A": map[string]any{},
					"G": map[string]any{},
				},
			},
			jsondiff.Patch{
				{
					Type: "remove",
					Path: "/A",
					OldValue: map[string]any{
						"B": map[string]any{
							"C": "D",
							"E": "F",
						},
					},
					Value: nil,
				},
			},
		},
		{
			"Updating create-and-write-only property noops",
			map[string]any{
				"A": "B",
			},
			map[string]any{
				"A": "C",
			},
			map[string]any{
				"properties": map[string]any{
					"A": map[string]any{},
				},
				"createOnlyProperties": []any{
					"/properties/A",
				},
				"writeOnlyProperties": []any{
					"/properties/A",
				},
			},
			nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			desiredStateBytes, err := json.Marshal(testCase.desiredState)
			require.NoError(t, err)

			currentStateBytes, err := json.Marshal(testCase.currentState)
			require.NoError(t, err)

			schemaBytes, err := json.Marshal(testCase.schema)
			require.NoError(t, err)

			patch, err := GeneratePatch(currentStateBytes, desiredStateBytes, schemaBytes)
			require.NoError(t, err)

			require.Equal(t, testCase.expectedPatch, patch)
		})
	}
}

func Test_ParsePropertyName(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			"ParsePropertyName successfully parses single property",
			"/properties/propertyName",
			"propertyName",
			nil,
		},
		{
			"ParsePropertyName successfully parses sub-properties",
			"/properties/propertyName/subProperty/subSubProperty",
			"propertyName/subProperty/subSubProperty",
			nil,
		},
		{
			"ParsePropertyName returns an error if input is invalid",
			"propertyName",
			"",
			fmt.Errorf("property identifier propertyName is not in the format /properties/<propertyName>"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := ParsePropertyName(testCase.input)
			require.Equal(t, testCase.output, actual)
			require.Equal(t, err, testCase.err)
		})
	}
}