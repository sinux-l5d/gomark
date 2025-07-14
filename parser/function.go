package parser

import (
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser/tokenizer"
)

type FunctionParser struct{}

func NewFunctionParser() *FunctionParser {
	return &FunctionParser{}
}

func (*FunctionParser) Match(tokens []*tokenizer.Token) (ast.Node, int) {
	matchedTokens := tokenizer.GetFirstLine(tokens)
	if len(matchedTokens) < 2 {
		return nil, 0
	}

	if matchedTokens[0].Type != tokenizer.ExclamationMark {
		return nil, 0
	}

	cursor, nameTokens := 1, []*tokenizer.Token{}
	for ; cursor < len(matchedTokens); cursor++ {
		if matchedTokens[cursor].Type == tokenizer.Space || matchedTokens[cursor].Type == tokenizer.LeftParenthesis {
			break
		}
		nameTokens = append(nameTokens, matchedTokens[cursor])
	}

	if len(nameTokens) == 0 {
		return nil, 0
	}

	// if cursor is at the end or that the next element is not a (, then end
	if cursor >= len(matchedTokens) || matchedTokens[cursor].Type != tokenizer.LeftParenthesis {
		return &ast.Function{
			Name:   tokenizer.Stringify(nameTokens),
			Params: []string{},
		}, len(nameTokens) + 1
	}

	cursor += 1
	paramsIndex := 0
	// bidimensional array of tokens
	paramsTokens, matched := make([][]*tokenizer.Token, 0), false
	if matchedTokens[cursor].Type != tokenizer.RightParenthesis {
		paramsTokens = append(paramsTokens, []*tokenizer.Token{})
	}
	for _, token := range matchedTokens[cursor:] {
		if token.Type == tokenizer.RightParenthesis {
			// No empty parameter allowed
			if paramsIndex > 0 && len(paramsTokens[paramsIndex]) == 0 {
				return nil, 0
			}
			matched = true
			break
		} else if token.Type == tokenizer.Comma {
			// No empty parameter allowed
			if len(paramsTokens[paramsIndex]) == 0 {
				return nil, 0
			}
			paramsIndex++
			paramsTokens = append(paramsTokens, []*tokenizer.Token{})
		} else {
			paramsTokens[paramsIndex] = append(paramsTokens[paramsIndex], token)
		}
	}
	if !matched || paramsIndex != max(len(paramsTokens)-1, 0) {
		return nil, 0
	}

	paramsStringified := make([]string, len(paramsTokens))
	paramsTokensCount := 0
	for i, tokens := range paramsTokens {
		paramsStringified[i] = tokenizer.Stringify(tokens)
		paramsTokensCount += len(tokens)
	}

	return &ast.Function{
		Name:   tokenizer.Stringify(nameTokens),
		Params: paramsStringified,
	}, 3 + len(nameTokens) + paramsTokensCount + max(len(paramsTokens)-1, 0)
}
