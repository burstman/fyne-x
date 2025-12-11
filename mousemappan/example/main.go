package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/x/fyne/mousemappan"
)

func main() {
	a := app.New()
	w := a.NewWindow("MouseMapPan - Sousse, Tunisia")

	// Create map with options
	m := mousemappan.NewMapWithOptions(
		mousemappan.WithOsmTiles(),
		mousemappan.WithZoomButtons(true),
		mousemappan.WithScrollButtons(false), // Use mouse drag instead of arrows
	)

	// Set zoom level and center on Sousse, Tunisia
	m.Zoom(12)
	m.CenterOnLocation(35.83, 10.64)

	w.SetContent(m)
	w.SetPadded(false)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
