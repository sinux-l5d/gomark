package ast

import (
	"fmt"
	"strings"
)

type BaseInline struct {
	BaseNode
}

type Text struct {
	BaseInline

	Content string
}

func (*Text) Type() NodeType {
	return TextNode
}

func (n *Text) Restore() string {
	return n.Content
}

type Bold struct {
	BaseInline

	// Symbol is "*" or "_".
	Symbol   string
	Children []Node
}

func (*Bold) Type() NodeType {
	return BoldNode
}

func (n *Bold) Restore() string {
	symbol := n.Symbol + n.Symbol
	children := ""
	for _, child := range n.Children {
		children += child.Restore()
	}
	return fmt.Sprintf("%s%s%s", symbol, children, symbol)
}

type Italic struct {
	BaseInline

	// Symbol is "*" or "_".
	Symbol   string
	Children []Node
}

func (*Italic) Type() NodeType {
	return ItalicNode
}

func (n *Italic) Restore() string {
	content := ""
	for _, child := range n.Children {
		content += child.Restore()
	}
	return fmt.Sprintf("%s%s%s", n.Symbol, content, n.Symbol)
}

type BoldItalic struct {
	BaseInline

	// Symbol is "*" or "_".
	Symbol  string
	Content string
}

func (*BoldItalic) Type() NodeType {
	return BoldItalicNode
}

func (n *BoldItalic) Restore() string {
	symbol := n.Symbol + n.Symbol + n.Symbol
	return fmt.Sprintf("%s%s%s", symbol, n.Content, symbol)
}

type Code struct {
	BaseInline

	Content string
}

func (*Code) Type() NodeType {
	return CodeNode
}

func (n *Code) Restore() string {
	return fmt.Sprintf("`%s`", n.Content)
}

type Image struct {
	BaseInline

	AltText string
	URL     string
}

func (*Image) Type() NodeType {
	return ImageNode
}

func (n *Image) Restore() string {
	return fmt.Sprintf("![%s](%s)", n.AltText, n.URL)
}

type Link struct {
	BaseInline

	Content []Node
	URL     string
}

func (*Link) Type() NodeType {
	return LinkNode
}

func (n *Link) Restore() string {
	content := ""
	for _, child := range n.Content {
		content += child.Restore()
	}
	return fmt.Sprintf("[%s](%s)", content, n.URL)
}

type AutoLink struct {
	BaseInline

	URL       string
	IsRawText bool
}

func (*AutoLink) Type() NodeType {
	return AutoLinkNode
}

func (n *AutoLink) Restore() string {
	if n.IsRawText {
		return n.URL
	}
	return fmt.Sprintf("<%s>", n.URL)
}

type Tag struct {
	BaseInline

	Content string
}

func (*Tag) Type() NodeType {
	return TagNode
}

func (n *Tag) Restore() string {
	return fmt.Sprintf("#%s", n.Content)
}

type Strikethrough struct {
	BaseInline

	Content string
}

func (*Strikethrough) Type() NodeType {
	return StrikethroughNode
}

func (n *Strikethrough) Restore() string {
	return fmt.Sprintf("~~%s~~", n.Content)
}

type EscapingCharacter struct {
	BaseInline

	Symbol string
}

func (*EscapingCharacter) Type() NodeType {
	return EscapingCharacterNode
}

func (n *EscapingCharacter) Restore() string {
	return fmt.Sprintf("\\%s", n.Symbol)
}

type Math struct {
	BaseInline

	Content string
}

func (*Math) Type() NodeType {
	return MathNode
}

func (n *Math) Restore() string {
	return fmt.Sprintf("$%s$", n.Content)
}

type Highlight struct {
	BaseInline

	Content string
}

func (*Highlight) Type() NodeType {
	return HighlightNode
}

func (n *Highlight) Restore() string {
	return fmt.Sprintf("==%s==", n.Content)
}

type Subscript struct {
	BaseInline

	Content string
}

func (*Subscript) Type() NodeType {
	return SubscriptNode
}

func (n *Subscript) Restore() string {
	return fmt.Sprintf("~%s~", n.Content)
}

type Superscript struct {
	BaseInline

	Content string
}

func (*Superscript) Type() NodeType {
	return SuperscriptNode
}

func (n *Superscript) Restore() string {
	return fmt.Sprintf("^%s^", n.Content)
}

type ReferencedContent struct {
	BaseInline

	ResourceName string
	Params       string
}

func (*ReferencedContent) Type() NodeType {
	return ReferencedContentNode
}

func (n *ReferencedContent) Restore() string {
	params := ""
	if n.Params != "" {
		params = fmt.Sprintf("?%s", n.Params)
	}
	result := fmt.Sprintf("[[%s%s]]", n.ResourceName, params)
	return result
}

type Spoiler struct {
	BaseInline

	Content string
}

func (*Spoiler) Type() NodeType {
	return SpoilerNode
}

func (n *Spoiler) Restore() string {
	return fmt.Sprintf("||%s||", n.Content)
}

type HTMLElement struct {
	BaseInline

	TagName    string
	Attributes map[string]string
}

func (*HTMLElement) Type() NodeType {
	return HTMLElementNode
}

func (n *HTMLElement) Restore() string {
	attributes := []string{}
	for key, value := range n.Attributes {
		attributes = append(attributes, fmt.Sprintf(`%s="%s"`, key, value))
	}
	attrStr := ""
	if len(attributes) > 0 {
		attrStr = " " + strings.Join(attributes, " ")
	}
	return fmt.Sprintf("<%s%s />", n.TagName, attrStr)
}

type Function struct {
	BaseInline

	Name   string
	Params []string
}

func (*Function) Type() NodeType {
	return FunctionNode
}

func (n *Function) Restore() string {
	params := ""
	if len(n.Params) > 0 {
		params = fmt.Sprintf("(%s)", strings.Join(n.Params, ","))
	}
	return fmt.Sprintf("!%s%s", n.Name, params)
}
