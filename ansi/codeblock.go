package ansi

import (
	"io"

	"github.com/muesli/reflow/indent"
	rwansi "github.com/rwinkhart/go-highlite/ansi"
)

// A CodeBlockElement is used to render code blocks.
type CodeBlockElement struct {
	Code     string
	Language string
}

func (e *CodeBlockElement) Render(w io.Writer, ctx RenderContext) error {
	bs := ctx.blockStack

	var indentation uint
	var margin uint
	rules := ctx.options.Styles.CodeBlock
	if rules.Indent != nil {
		indentation = *rules.Indent
	}
	if rules.Margin != nil {
		margin = *rules.Margin
	}

	iw := indent.NewWriterPipe(w, indentation+margin, func(wr io.Writer) {
		renderText(w, ctx.options.ColorProfile, bs.Current().Style.StylePrimitive, " ")
	})

	renderText(iw, ctx.options.ColorProfile, bs.Current().Style.StylePrimitive, rules.BlockPrefix)
	highightedCode, err := rwansi.Colorize(e.Code, e.Language, 171)
	if err != nil {
		return err
	}
	iw.Write([]byte(highightedCode))
	renderText(iw, ctx.options.ColorProfile, bs.Current().Style.StylePrimitive, rules.BlockSuffix)
	return nil
}
