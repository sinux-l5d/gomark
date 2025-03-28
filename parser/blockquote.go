package parser

import (
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
)

type BlockquoteParser struct{}

func NewBlockquoteParser() *BlockquoteParser {
	return &BlockquoteParser{}
}

func (*BlockquoteParser) Match(tokens []*tokenizer.Token) (ast.Node, int) {
	rows := tokenizer.Split(tokens, tokenizer.NewLine)
	contentRows := [][]*tokenizer.Token{}
	for _, row := range rows {
		if len(row) < 2 || row[0].Type != tokenizer.GreaterThan || row[1].Type != tokenizer.Space {
			break
		}
		contentRows = append(contentRows, row)
	}
	if len(contentRows) == 0 {
		return nil, 0
	}

	children := []ast.Node{}
	size := 0

	for index, row := range contentRows {
		contentTokens := row[2:]
		var node ast.Node
		if len(contentTokens) == 0 {
			node = &ast.Paragraph{
				Children: []ast.Node{&ast.Text{Content: " "}},
			}
		} else {
			nodes, err := ParseBlockWithParsers(contentTokens, []BlockParser{NewBlockquoteParser(), NewParagraphParser()})
			if err != nil {
				return nil, 0
			}
			if len(nodes) != 1 {
				return nil, 0
			}
			node = nodes[0]
		}
		children = append(children, node)
		size += len(row)
		if index != len(contentRows)-1 {
			size++ // NewLine.
		}
	}
	if len(children) == 0 {
		return nil, 0
	}
	return &ast.Blockquote{
		Children: children,
	}, size
}
