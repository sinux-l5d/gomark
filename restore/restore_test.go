package restore

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/usememos/gomark/ast"
)

func TestRestore(t *testing.T) {
	tests := []struct {
		nodes   []ast.Node
		rawText string
	}{
		{
			nodes:   nil,
			rawText: "",
		},
		{
			nodes: []ast.Node{
				&ast.Text{
					Content: "Hello world!",
				},
			},
			rawText: "Hello world!",
		},
		{
			nodes: []ast.Node{
				&ast.Paragraph{
					Children: []ast.Node{
						&ast.Text{
							Content: "Code: ",
						},
						&ast.Code{
							Content: "Hello world!",
						},
					},
				},
			},
			rawText: "Code: `Hello world!`",
		},
		{
			nodes: []ast.Node{
				&ast.Function{
					Name:   "test",
					Params: []string{"hello", " world"},
				},
			},
			rawText: "!test(hello, world)",
		},
		{
			nodes: []ast.Node{
				&ast.Function{
					Name:   "test",
					Params: []string{},
				},
			},
			rawText: "!test",
		},
	}

	for _, test := range tests {
		require.Equal(t, Restore(test.nodes), test.rawText)
	}
}
