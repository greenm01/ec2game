package bbsui

// An example program demonstrating the pager component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const (
	useHighPerformanceRenderer = false
	width = 80
	height = 25
)

var (
	titleStyle = func() lg.Style {
		b := lg.RoundedBorder()
		b.Right = "├"
		return lg.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lg.Style {
		b := lg.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type Pager struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func (p *Pager) Build(s string) {
	
	p.content = s
	
	headerHeight := lg.Height(p.headerView())
	footerHeight := lg.Height(p.footerView())
	verticalMarginHeight := headerHeight + footerHeight
	// Since this program is using the full size of the viewport we
	// need to wait until we've received the window dimensions before
	// we can initialize the viewport. The initial dimensions come in
	// quickly, though asynchronously, which is why we wait for them
	// here.
	p.viewport = viewport.New(width, height-verticalMarginHeight)
	p.viewport.YPosition = headerHeight
	p.viewport.HighPerformanceRendering = useHighPerformanceRenderer
	p.viewport.SetContent(p.content)
	p.ready = true

	// This is only necessary for high performance rendering, which in
	// most cases you won't need.
	//
	// Render the viewport one line below the header.
	p.viewport.YPosition = headerHeight + 1

}

func (m *Pager) Update(msg tea.Msg) tea.Cmd {
	
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "esc" {
			return changeMenu("ftm")
		}
	}
		
	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m Pager) View() string {
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m Pager) headerView() string {
	title := titleStyle.Render("Esterian Conquest")
	line := strings.Repeat("─", max(0, m.viewport.Width-lg.Width(title)))
	return lg.JoinHorizontal(lg.Center, title, line)
}

func (m Pager) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lg.Width(info)))
	bl := lg.JoinHorizontal(lg.Center, line, info)	
	bl += "\nArrows to scroll, ESC to exit."
	return bl
}

