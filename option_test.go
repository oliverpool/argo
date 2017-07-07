package argo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeOption(t *testing.T) {
	a := assert.New(t)

	cases := []struct {
		Input    []Option
		Expected Option
	}{
		{
			Input:    nil,
			Expected: Option{},
		},
		{
			Input:    []Option{},
			Expected: Option{},
		},
		{
			Input:    []Option{{}},
			Expected: Option{},
		},
		{
			Input:    []Option{nil},
			Expected: Option{},
		},
		{
			Input:    []Option{nil, {"a": "1"}, nil},
			Expected: Option{"a": "1"},
		},
		{
			Input:    []Option{{"a": "1"}},
			Expected: Option{"a": "1"},
		},
		{
			Input:    []Option{{"a": "1"}, {"a": "2"}},
			Expected: Option{"a": "2"},
		},
		{
			Input:    []Option{{"a": "1"}, {"b": "2"}},
			Expected: Option{"a": "1", "b": "2"},
		},
	}

	for _, c := range cases {
		actual := mergeOptions(c.Input...)
		a.Equal(c.Expected, actual)
		a.True(reflect.DeepEqual(c.Expected, actual))
	}

}
