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
	"bytes"
	"fmt"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

type testOutput struct {
	Foo string `json:"foo" yaml:"foo"`
	Bar int    `json:"bar" yaml:"bar"`
}

func TestNewPrinter_Table(t *testing.T) {
	p := NewPrinter(&bytes.Buffer{}, false, false, false)

	assert.True(t, p.table)
	assert.False(t, p.json)
	assert.False(t, p.yaml)
	assert.False(t, p.noHeader)
}

func TestNewPrinter_JSON(t *testing.T) {
	p := NewPrinter(&bytes.Buffer{}, true, false, false)

	assert.False(t, p.table)
	assert.True(t, p.json)
	assert.False(t, p.yaml)
	assert.False(t, p.noHeader)
}

func TestNewPrinter_Yaml(t *testing.T) {
	p := NewPrinter(&bytes.Buffer{}, false, true, false)

	assert.False(t, p.table)
	assert.False(t, p.json)
	assert.True(t, p.yaml)
	assert.False(t, p.noHeader)
}

func TestPrinter_SetRowFunc(t *testing.T) {
	p := Printer{}
	assert.Nil(t, p.rowFunc)

	p.SetRowFunc(func(data interface{}) (i []interface{}, e error) {
		return nil, nil
	})
	assert.NotNil(t, p.rowFunc)
}

func TestPrinter_SetHeader(t *testing.T) {
	p := Printer{}
	assert.Nil(t, p.header)

	p.SetHeader("FOO", "BAR")
	assert.NotNil(t, p.header)
	assert.Equal(t, []string{"FOO", "BAR"}, p.header)
}

func TestPrinter_Write_err(t *testing.T) {
	p := Printer{}

	err := p.Write("test")
	assert.Error(t, err)
	assert.Equal(t, ErrNoOutputMode, err)
}

func TestPrinter_toJSON(t *testing.T) {
	data := testOutput{
		Foo: "test",
		Bar: 2,
	}

	out := &bytes.Buffer{}
	p := Printer{
		out: out,
	}

	err := p.toJSON(&data)
	assert.NoError(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			{
			  "foo": "test",
			  "bar": 2
			}
		`),
		out.String(),
	)
}

func TestPrinter_toYAML(t *testing.T) {
	data := testOutput{
		Foo: "test",
		Bar: 2,
	}

	out := &bytes.Buffer{}
	p := Printer{
		out: out,
	}

	err := p.toYAML(&data)
	assert.NoError(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			bar: 2
			foo: test
		`),
		out.String(),
	)
}

func TestPrinter_toTable_sliceOK(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		out:    out,
		header: []string{"FOO"},
		rowFunc: func(data interface{}) (i []interface{}, e error) {
			return []interface{}{data}, nil
		},
	}

	err := p.toTable([]string{"1", "2"})
	assert.NoError(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			FOO
			1
			2
		`),
		out.String(),
	)
}

func TestPrinter_toTable_defaultOK(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		out:    out,
		header: []string{"FOO"},
		rowFunc: func(data interface{}) (i []interface{}, e error) {
			return []interface{}{data}, nil
		},
	}

	err := p.toTable("1")
	assert.NoError(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			FOO
			1
		`),
		out.String(),
	)
}

func TestPrinter_toTable_sliceErr(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		out:    out,
		header: []string{"FOO"},
		rowFunc: func(data interface{}) (i []interface{}, e error) {
			return nil, fmt.Errorf("test error")
		},
	}

	err := p.toTable([]string{"1", "2"})
	assert.Error(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			FOO
		`),
		out.String(),
	)
}

func TestPrinter_toTable_defaultErr(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		out:    out,
		header: []string{"FOO"},
		rowFunc: func(data interface{}) (i []interface{}, e error) {
			return nil, fmt.Errorf("test error")
		},
	}

	err := p.toTable("1")
	assert.Error(t, err)
	assert.Equal(
		t,
		heredoc.Doc(`
			FOO
		`),
		out.String(),
	)
}

func TestPrinter_toTable_noRowFn(t *testing.T) {
	p := Printer{}

	err := p.toTable("test data")
	assert.Error(t, err)
	assert.Equal(t, ErrNoRowFunc, err)
}

func TestPrinter_writeHeader_noHeader(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		header:   []string{"FOO", "BAR"},
		noHeader: true,
	}

	err := p.writeHeader(out)
	assert.NoError(t, err)
	assert.Empty(t, out.String())
}

func TestPrinter_writeHeader_withHeader(t *testing.T) {
	out := &bytes.Buffer{}
	p := Printer{
		header:   []string{"FOO", "BAR"},
		noHeader: false,
	}

	err := p.writeHeader(out)
	assert.NoError(t, err)
	assert.Equal(t, "FOO\tBAR\n", out.String())
}
