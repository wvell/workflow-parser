package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsesStringer(t *testing.T) {
	cases := []struct {
		uses     fmt.Stringer
		expected string
	}{
		{
			uses:     &UsesDockerImage{Image: "alpine"},
			expected: "docker://alpine",
		},
		{
			uses:     &UsesRepository{Repository: "wvell/workflow-parser", Ref: "master"},
			expected: "wvell/workflow-parser@master",
		},
		{
			uses:     &UsesRepository{Repository: "wvell/workflow-parser", Path: "path", Ref: "master"},
			expected: "wvell/workflow-parser/path@master",
		},
		{
			uses:     &UsesPath{Path: "path"},
			expected: "./path",
		},
		{
			uses:     &UsesInvalid{},
			expected: "",
		},
		{
			uses:     &UsesInvalid{Raw: "foo"},
			expected: "foo",
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, tc.uses.String())
	}
}
