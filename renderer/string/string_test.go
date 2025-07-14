package string

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
)

func TestStringRender(t *testing.T) {
	tests := []struct {
		text     string
		expected string
	}{
		{
			text:     "",
			expected: "",
		},
		{
			text:     "Hello world!",
			expected: "Hello world!\n",
		},
		{
			text:     "**Hello** world!",
			expected: "Hello world!\n",
		},
		{
			text:     "Test\n1. Hello\n2. World",
			expected: "Test\n1. Hello\n2. World",
		},
		{
			text:     "!test(hello)",
			expected: "!test(hello)\n",
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		nodes, err := parser.Parse(tokens)
		require.NoError(t, err)
		actual := NewStringRenderer().Render(nodes)
		require.Equal(t, test.expected, actual)
	}
}
