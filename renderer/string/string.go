package string

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/usememos/gomark/ast"
)

type RendererContext struct {
}

// StringRenderer renders AST to raw string.
type StringRenderer struct {
	output  *bytes.Buffer
	context *RendererContext
}

// NewStringRenderer creates a new StringRender.
func NewStringRenderer() *StringRenderer {
	return &StringRenderer{
		output:  new(bytes.Buffer),
		context: &RendererContext{},
	}
}

// RenderNode renders a single AST node to raw string.
func (r *StringRenderer) RenderNode(node ast.Node) {
	switch n := node.(type) {
	case *ast.LineBreak:
		r.renderLineBreak(n)
	case *ast.Paragraph:
		r.renderParagraph(n)
	case *ast.CodeBlock:
		r.renderCodeBlock(n)
	case *ast.Heading:
		r.renderHeading(n)
	case *ast.HorizontalRule:
		r.renderHorizontalRule(n)
	case *ast.Blockquote:
		r.renderBlockquote(n)
	case *ast.List:
		r.renderList(n)
	case *ast.UnorderedListItem:
		r.renderUnorderedListItem(n)
	case *ast.OrderedListItem:
		r.renderOrderedListItem(n)
	case *ast.TaskListItem:
		r.renderTaskListItem(n)
	case *ast.MathBlock:
		r.renderMathBlock(n)
	case *ast.Table:
		r.renderTable(n)
	case *ast.EmbeddedContent:
		r.renderEmbeddedContent(n)
	case *ast.Text:
		r.renderText(n)
	case *ast.Bold:
		r.renderBold(n)
	case *ast.Italic:
		r.renderItalic(n)
	case *ast.BoldItalic:
		r.renderBoldItalic(n)
	case *ast.Code:
		r.renderCode(n)
	case *ast.Image:
		r.renderImage(n)
	case *ast.Link:
		r.renderLink(n)
	case *ast.AutoLink:
		r.renderAutoLink(n)
	case *ast.Tag:
		r.renderTag(n)
	case *ast.Function:
		r.renderFunction(n)
	case *ast.Strikethrough:
		r.renderStrikethrough(n)
	case *ast.EscapingCharacter:
		r.renderEscapingCharacter(n)
	case *ast.Math:
		r.renderMath(n)
	case *ast.Highlight:
		r.renderHighlight(n)
	case *ast.Subscript:
		r.renderSubscript(n)
	case *ast.Superscript:
		r.renderSuperscript(n)
	case *ast.ReferencedContent:
		r.renderReferencedContent(n)
	case *ast.Spoiler:
		r.renderSpoiler(n)
	case *ast.HTMLElement:
		r.renderHTMLElement(n)
	default:
		// Handle other block types if needed.
	}
}

// RenderNodes renders a slice of AST nodes to raw string.
func (r *StringRenderer) RenderNodes(nodes []ast.Node) {
	var prevNode ast.Node
	var skipNextLineBreakFlag bool
	for _, node := range nodes {
		if node.Type() == ast.LineBreakNode && skipNextLineBreakFlag {
			if prevNode != nil && ast.IsBlockNode(prevNode) {
				skipNextLineBreakFlag = false
				continue
			}
		}

		r.RenderNode(node)
		prevNode = node
		skipNextLineBreakFlag = true
	}
}

// Render renders the AST to raw string.
func (r *StringRenderer) Render(astRoot []ast.Node) string {
	r.RenderNodes(astRoot)
	return r.output.String()
}

func (r *StringRenderer) renderLineBreak(_ *ast.LineBreak) {
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderParagraph(node *ast.Paragraph) {
	r.RenderNodes(node.Children)
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderCodeBlock(node *ast.CodeBlock) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderHeading(node *ast.Heading) {
	r.RenderNodes(node.Children)
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderHorizontalRule(_ *ast.HorizontalRule) {
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderBlockquote(node *ast.Blockquote) {
	r.RenderNodes(node.Children)
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderList(node *ast.List) {
	for _, item := range node.Children {
		r.RenderNodes([]ast.Node{item})
	}
}

func (r *StringRenderer) renderUnorderedListItem(node *ast.UnorderedListItem) {
	r.output.WriteString(node.Symbol)
	r.RenderNodes(node.Children)
}

func (r *StringRenderer) renderOrderedListItem(node *ast.OrderedListItem) {
	r.output.WriteString(fmt.Sprintf("%s. ", node.Number))
	r.RenderNodes(node.Children)
}

func (r *StringRenderer) renderTaskListItem(node *ast.TaskListItem) {
	r.output.WriteString(node.Symbol)
	r.RenderNodes(node.Children)
}

func (r *StringRenderer) renderMathBlock(node *ast.MathBlock) {
	r.output.WriteString(node.Content)
	r.output.WriteString("\n")
}

func (r *StringRenderer) renderTable(node *ast.Table) {
	for _, cell := range node.Header {
		r.RenderNodes([]ast.Node{cell})
		r.output.WriteString("\t")
	}
	r.output.WriteString("\n")
	for _, row := range node.Rows {
		for _, cell := range row {
			r.RenderNodes([]ast.Node{cell})
			r.output.WriteString("\t")
		}
		r.output.WriteString("\n")
	}
}

func (*StringRenderer) renderEmbeddedContent(_ *ast.EmbeddedContent) {}

func (r *StringRenderer) renderText(node *ast.Text) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderBold(node *ast.Bold) {
	r.RenderNodes(node.Children)
}

func (r *StringRenderer) renderItalic(node *ast.Italic) {
	r.RenderNodes(node.Children)
}

func (r *StringRenderer) renderBoldItalic(node *ast.BoldItalic) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderCode(node *ast.Code) {
	r.output.WriteString(node.Content)
}

func (*StringRenderer) renderImage(_ *ast.Image) {}

func (r *StringRenderer) renderLink(node *ast.Link) {
	r.output.WriteString(node.URL)
}

func (r *StringRenderer) renderAutoLink(node *ast.AutoLink) {
	r.output.WriteString(node.URL)
}

func (r *StringRenderer) renderTag(node *ast.Tag) {
	r.output.WriteString(`#`)
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderFunction(node *ast.Function) {
	r.output.WriteString("!")
	r.output.WriteString(node.Name)
	r.output.WriteString("(")
	r.output.WriteString(strings.Join(node.Params, ","))
	r.output.WriteString(")")
}

func (r *StringRenderer) renderStrikethrough(node *ast.Strikethrough) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderEscapingCharacter(node *ast.EscapingCharacter) {
	r.output.WriteString("\\")
	r.output.WriteString(node.Symbol)
}

func (r *StringRenderer) renderMath(node *ast.Math) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderHighlight(node *ast.Highlight) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderSubscript(node *ast.Subscript) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderSuperscript(node *ast.Superscript) {
	r.output.WriteString(node.Content)
}

func (*StringRenderer) renderReferencedContent(_ *ast.ReferencedContent) {}

func (r *StringRenderer) renderSpoiler(node *ast.Spoiler) {
	r.output.WriteString(node.Content)
}

func (r *StringRenderer) renderHTMLElement(*ast.HTMLElement) {
	r.output.WriteString("\n")
}
