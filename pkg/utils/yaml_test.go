// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

type testObj1 struct {
	Foo            string
	Bar            map[string]string
	SomeLongerName int
}

type testObj2 struct {
	Foo            string            `json:"foo"`
	Bar            map[string]string `json:"bar"`
	SomeLongerName int               `json:"some_longer_name"`
}

type testObj3 struct {
	Foo            string            `json:"foo" yaml:"yfoo"`
	Bar            map[string]string `json:"bar" yaml:"ybar"`
	SomeLongerName int               `json:"some_longer_name" yaml:"ysome_longer_name"`
}

type testObj4 struct {
	Foo            string            `json:"foo"`
	Bar            map[string]string `json:"-"`
	SomeLongerName int               `json:"some_longer_name"`
}

func TestObjToYAML(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		expected string
	}{
		{
			name: "simple map",
			obj: map[string]interface{}{
				"foo": "bar",
				"a":   1,
			},
			expected: heredoc.Doc(`
				a: 1
				foo: bar
			`),
		},
		{
			name: "nested map",
			obj: map[string]interface{}{
				"foo": "bar",
				"a": map[string]interface{}{
					"b": []int{
						1, 2, 3,
					},
				},
			},
			expected: heredoc.Doc(`
				a:
				  b:
				  - 1
				  - 2
				  - 3
				foo: bar
			`),
		},
		{
			name: "obj with no tags",
			obj: &testObj1{
				Foo: "foo",
				Bar: map[string]string{
					"abc": "def",
				},
				SomeLongerName: 3,
			},
			expected: heredoc.Doc(`
				Bar:
				  abc: def
				Foo: foo
				SomeLongerName: 3
			`),
		},
		{
			name: "obj with json tags",
			obj: &testObj2{
				Foo: "foo",
				Bar: map[string]string{
					"abc": "def",
				},
				SomeLongerName: 3,
			},
			expected: heredoc.Doc(`
				bar:
				  abc: def
				foo: foo
				some_longer_name: 3
			`),
		},
		{
			name: "obj with json and yaml tags",
			obj: &testObj3{
				Foo: "foo",
				Bar: map[string]string{
					"abc": "def",
				},
				SomeLongerName: 3,
			},
			// The expected here should take the keys specified by the JSON
			// tags, since we are first marshalling to JSON.
			expected: heredoc.Doc(`
				bar:
				  abc: def
				foo: foo
				some_longer_name: 3
			`),
		},
		{
			name: "obj with some json tags",
			obj: &testObj4{
				Foo: "foo",
				Bar: map[string]string{
					"abc": "def",
				},
				SomeLongerName: 3,
			},
			expected: heredoc.Doc(`
				foo: foo
				some_longer_name: 3
			`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ObjToYAML(test.obj)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(res))
		})
	}
}
