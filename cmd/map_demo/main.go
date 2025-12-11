package main

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	xwidget "fyne.io/x/fyne/widget"
)

// latLonToTile converts latitude and longitude to tile coordinates at a given zoom level
func latLonToTile(lat, lon float64, zoom int) (x, y int) {
	n := math.Pow(2, float64(zoom))
	x = int((lon + 180.0) / 360.0 * n)
	latRad := lat * math.Pi / 180.0
	y = int((1.0 - math.Log(math.Tan(latRad)+1.0/math.Cos(latRad))/math.Pi) / 2.0 * n)
	return x, y
}

func main() {
	w := app.New().NewWindow("Map Widget - Sousse")

	m := xwidget.NewMapWithOptions(
		xwidget.WithOsmTiles(),
		xwidget.WithZoomButtons(true),
		xwidget.WithScrollButtons(false), // Disable arrow buttons, use mouse drag instead
	)

	// Sousse, Tunisia coordinates: ~35.83°N, 10.64°E
	// Set zoom level to 12 for a good city view
	sousseLat := 35.83
	sousseLon := 10.64
	zoomLevel := 12

	m.Zoom(zoomLevel)

	// Convert lat/lon to tile coordinates and center the map
	tileX, tileY := latLonToTile(sousseLat, sousseLon, zoomLevel)
	// Offset to center (tile coordinates are relative to center)
	count := 1 << zoomLevel
	// Calculate the offset from center
	centerX := int(float32(count)/2 - 0.5)
	centerY := int(float32(count)/2 - 0.5)
	deltaX := tileX - centerX
	deltaY := tileY - centerY

	// Pan to Sousse (4 clicks per tile since each click = 1/4 tile)
	for i := 0; i < abs(deltaX)*4; i++ {
		if deltaX > 0 {
			m.PanEast()
		} else {
			m.PanWest()
		}
	}
	for i := 0; i < abs(deltaY)*4; i++ {
		if deltaY > 0 {
			m.PanSouth()
		} else {
			m.PanNorth()
		}
	}

	w.SetContent(m)

	w.SetPadded(false)
	w.Resize(fyne.NewSize(512, 320))
	w.ShowAndRun()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
