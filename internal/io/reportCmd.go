package io

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	bx "github.com/treilik/bubbleboxer"
)

const (
	leftAddr   = "left"
	middleAddr = "middle"
	rightAddr  = "right"
	lowerAddr  = "lower"
)

type reportCmd struct {
	text string
}

func initReportCmd(m *bx.Boxer) {

	// leaf content creation (models)
	left := stringer(leftAddr)
	middle := stringer(middleAddr)
	right := stringer(rightAddr)

	lower := stringer(fmt.Sprintf("%s: use ctrl+c to quit", lowerAddr))
	width := func(_ bx.Node, widthOrHeight int) []int {
		return []int{12, 12, 12}
	}
	// layout-tree defintion

	m.LayoutTree = bx.Node{
		// orientation
		VerticalStacked: true,
		// spacing
		SizeFunc: func(_ bx.Node, widthOrHeight int) []int {
			return []int{
				// since this node is vertical stacked return the height partioning since the width stays for all children fixed
				//widthOrHeight - 1,
				1,
				1,
				// make also sure that the amount of the returned ints match the amount of children:
				// in this case two, but in more complex cases read the amount of the chilren from the len(boxer.Node.Children)
			}
		},
		Children: []bx.Node{
			{
				SizeFunc: width,
				Children: []bx.Node{
					// make sure to encapsulate the models into a leaf with CreateLeaf:
					m.CreateLeaf(leftAddr, left),
					m.CreateLeaf(middleAddr, middle),
					m.CreateLeaf(rightAddr, right),
				},
			},
			m.CreateLeaf(lowerAddr, lower),
		},
	}
	m.UpdateSize(tea.WindowSizeMsg{Width: 132, Height: 37})
}

// satisfy the tea.Model interface
func (b reportCmd) Init() tea.Cmd                           { return nil }
func (b reportCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return b, nil }
func (b reportCmd) View() string                            { return b.text }

type stringer string

func (s stringer) String() string {
	return string(s)
}

// satisfy the tea.Model interface
func (s stringer) Init() tea.Cmd                           { return nil }
func (s stringer) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return s, nil }
func (s stringer) View() string                            { return s.String() }
