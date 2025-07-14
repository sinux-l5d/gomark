package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
)

func TestFunctionParser(t *testing.T) {
	tests := []struct {
		text string
		node ast.Node
	}{
		{
			text: "!test",
			node: &ast.Function{
				Name:   "test",
				Params: []string{},
			},
		},
		{
			text: "!test()",
			node: &ast.Function{
				Name:   "test",
				Params: []string{},
			},
		},
		{
			text: "!test(hello)",
			node: &ast.Function{
				Name:   "test",
				Params: []string{"hello"},
			},
		},
		{
			text: "!due(2025-07-14)",
			node: &ast.Function{
				Name:   "due",
				Params: []string{"2025-07-14"},
			},
		},
		{
			text: "!test/sub(hello)",
			node: &ast.Function{
				Name:   "test/sub",
				Params: []string{"hello"},
			},
		},
		{
			text: "!test(hello,  world)",
			node: &ast.Function{
				Name:   "test",
				Params: []string{"hello", "  world"},
			},
		},
		{
			text: "!test( hello, world)",
			node: &ast.Function{
				Name:   "test",
				Params: []string{" hello", " world"},
			},
		},
		{
			text: "!test(hello,world)",
			node: &ast.Function{
				Name:   "test",
				Params: []string{"hello", "world"},
			},
		},
		{
			text: "!test (hello)",
			node: &ast.Function{
				Name:   "test",
				Params: []string{},
			},
		},
		{
			text: "! test",
			node: nil,
		},
		{
			text: "! test(hello)",
			node: nil,
		},
		{
			text: "!test(,)",
			node: nil,
		},
		{
			text: "!test(hello,)",
			node: nil,
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, count := parser.NewFunctionParser().Match(tokens)
		t.Logf("`%s` as %d tokens", test.text, count)
		// t.Logf("Params excpected: %#v", test.node.(*ast.Function).Params)
		// t.Logf("Params got: %#v", node.(*ast.Function).Params)
		require.Equal(t, test.node, node)
	}
}
