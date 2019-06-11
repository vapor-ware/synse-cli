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

	synse "github.com/vapor-ware/synse-server-grpc/go"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTags(t *testing.T) {
	cases := []struct {
		description string
		tags        []string
		expected    []string
	}{
		{
			description: "empty tags list",
			tags:        []string{},
			expected:    nil,
		},
		{
			description: "single tag string, not comma separated",
			tags:        []string{"test/1"},
			expected:    []string{"test/1"},
		},
		{
			description: "single tag string, comma separated",
			tags:        []string{"test/1,test/2"},
			expected:    []string{"test/1", "test/2"},
		},
		{
			description: "multiple tag strings, not comma separated",
			tags:        []string{"test/1", "test/2"},
			expected:    []string{"test/1", "test/2"},
		},
		{
			description: "multiple tag strings, some comma separated",
			tags:        []string{"test/1", "test/2,test/3"},
			expected:    []string{"test/1", "test/2", "test/3"},
		},
		{
			description: "multiple tag strings, all comma separated",
			tags:        []string{"test/1,test/2", "test/3,test/4", "test/5,test/6,test/7"},
			expected:    []string{"test/1", "test/2", "test/3", "test/4", "test/5", "test/6", "test/7"},
		},
	}

	for _, c := range cases {
		actual := NormalizeTags(c.tags)
		assert.Equal(t, c.expected, actual, c.description)
	}
}

func TestStringToTag(t *testing.T) {
	cases := []struct {
		tag     string
		message *synse.V3Tag
	}{
		{
			tag:     "foo",
			message: &synse.V3Tag{Label: "foo"},
		},
		{
			tag:     "bar",
			message: &synse.V3Tag{Label: "bar"},
		},
		{
			tag:     "a/foo",
			message: &synse.V3Tag{Namespace: "a", Label: "foo"},
		},
		{
			tag:     "b/bar",
			message: &synse.V3Tag{Namespace: "b", Label: "bar"},
		},
		{
			tag:     "x:foo",
			message: &synse.V3Tag{Annotation: "x", Label: "foo"},
		},
		{
			tag:     "y:bar",
			message: &synse.V3Tag{Annotation: "y", Label: "bar"},
		},
		{
			tag:     "a/x:foo",
			message: &synse.V3Tag{Namespace: "a", Annotation: "x", Label: "foo"},
		},
		{
			tag:     "b/y:bar",
			message: &synse.V3Tag{Namespace: "b", Annotation: "y", Label: "bar"},
		},
		{
			tag:     "a-b/x-y:m-n",
			message: &synse.V3Tag{Namespace: "a-b", Annotation: "x-y", Label: "m-n"},
		}, {
			tag:     "a.b/x.y:m.n",
			message: &synse.V3Tag{Namespace: "a.b", Annotation: "x.y", Label: "m.n"},
		},
		{
			tag:     "  yankee/hotel:foxtrot  ",
			message: &synse.V3Tag{Namespace: "yankee", Annotation: "hotel", Label: "foxtrot"},
		},
	}

	for i, c := range cases {
		tag, err := StringToTag(c.tag)

		assert.NoError(t, err, "case: %d", i)
		assert.Equal(t, c.message.Namespace, tag.Namespace, "case: %d", i)
		assert.Equal(t, c.message.Annotation, tag.Annotation, "case: %d", i)
		assert.Equal(t, c.message.Label, tag.Label, "case: %d", i)
	}
}

func TestStringToTag_Error(t *testing.T) {
	cases := []struct {
		tag string
	}{
		{tag: ""},
		{tag: "a//b"},
		{tag: "a::b"},
		{tag: "a/b:"},
		{tag: "/"},
		{tag: "//"},
		{tag: ":"},
		{tag: "::"},
		{tag: "vaporio/contains spaces:foo"},
	}

	for i, c := range cases {
		tag, err := StringToTag(c.tag)

		assert.Error(t, err, "case: %d", i)
		assert.Nil(t, tag, "case: %d", i)
	}
}
