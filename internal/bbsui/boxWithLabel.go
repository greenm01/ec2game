package bbsui

import (
	"strings"
	lg "github.com/charmbracelet/lipgloss"
)

type boxWithLabel struct {
	BoxStyle   lg.Style
	LabelStyle lg.Style
}

func newBoxWithLabel() boxWithLabel {
	return boxWithLabel{
		BoxStyle: lg.NewStyle().
			Border(lg.RoundedBorder()).
			BorderForeground(lg.Color("63")).
			Padding(1),

		// You could, of course, also set background and foreground colors here 
		// as well.
		LabelStyle: lg.NewStyle().
			PaddingTop(0).
			PaddingBottom(0).
			PaddingLeft(1).
			PaddingRight(1),
	}
}

func (b boxWithLabel) Render(label, content string, width int) string {
	var (
		// Query the box style for some of its border properties so we can
		// essentially take the top border apart and put it around the label.
		border          lg.Border     = b.BoxStyle.GetBorderStyle()
		topBorderStyler func(string) string = lg.NewStyle().Foreground(b.BoxStyle.GetBorderTopForeground()).Render
		topLeft         string              = topBorderStyler(border.TopLeft)
		topRight        string              = topBorderStyler(border.TopRight)

		renderedLabel string = b.LabelStyle.Render(label)
	)

	// Render top row with the label
	borderWidth := b.BoxStyle.GetHorizontalBorderSize()
	cellsShort := max(0, width+borderWidth-lg.Width(topLeft+topRight+renderedLabel))
	gap := strings.Repeat(border.Top, cellsShort)
	top := topLeft + renderedLabel + topBorderStyler(gap) + topRight

	// Render the rest of the box
	bottom := b.BoxStyle.Copy().
		BorderTop(false).
		Width(width).
		Render(content)

	// Stack the pieces
	return top + "\n" + bottom
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
