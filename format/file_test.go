package format

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name  string
		file  string
		local string
	}{
		{
			name:  "File with no imports (no local)",
			file:  "no_imports",
			local: "",
		},
		{
			name:  "File with no imports",
			file:  "no_imports",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "Invalid imports (no local)",
			file:  "invalid_imports",
			local: "",
		},
		{
			name:  "Invalid imports (2)",
			file:  "invalid_imports_2",
			local: "",
		},
		{
			name:  "File with only std imports (no local)",
			file:  "only_stdlib_1",
			local: "",
		},
		{
			name:  "File with only std imports",
			file:  "only_stdlib_1",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "File with only std imports 2",
			file:  "only_stdlib_2",
			local: "",
		},
		{
			name:  "File with only std imports 3",
			file:  "only_stdlib_2",
			local: "",
		},
		{
			name:  "File with two groups of imports (stdlib and libraries)",
			file:  "two_groups_1",
			local: "",
		},
		{
			name:  "File with two groups of imports (stdlib and local)",
			file:  "two_groups_1",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "File with two groups of imports (stdlib and libraries)",
			file:  "two_groups_2",
			local: "",
		},
		{
			name:  "File with two groups of imports (stdlib and local)",
			file:  "two_groups_2",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "File with two groups of imports (stdlib and libraries)",
			file:  "two_groups_3",
			local: "",
		},
		{
			name:  "File with two groups of imports (stdlib and local)",
			file:  "two_groups_3",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "File with one import group: lib",
			file:  "one_group",
			local: "",
		},
		{
			name:  "File with one import group: local",
			file:  "one_group",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "Properly sorted file with three groups",
			file:  "three_groups_1",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "File with three group, which needs sorting",
			file:  "three_groups_2",
			local: "github.com/krzysztofdrys",
		},
		{
			name:  "Sorting file with import aliases",
			file:  "alias_sorting",
			local: "github.com/krzysztofdrys",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input, err := ioutil.ReadFile(filepath.Join("fixtures", fmt.Sprintf("%s_input", test.file)))
			if err != nil {
				t.Fatalf("failed to read input: %v", err)
			}
			expected, err := ioutil.ReadFile(filepath.Join("fixtures", fmt.Sprintf("%s_expected", test.file)))
			if err != nil {
				t.Fatalf("failed to read expected output: %v", err)
			}
			output := File(string(input), test.local)
			if string(expected) != output {
				t.Fatalf("Expected:\n%s\nGot:\n%s", string(expected), output)
			}
		})
	}
}
