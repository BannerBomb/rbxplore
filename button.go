package main

import (
	"github.com/anaminus/gxui"
	"github.com/anaminus/gxui/math"
)

var ButtonSize = math.Size{60, 26}

func CreateButton(theme gxui.Theme, text string) gxui.Button {
	button := theme.CreateButton()
	button.SetDesiredSize(ButtonSize)
	button.SetText(text)
	return button
}
