# MyMap - Interactive Map Widget for Fyne

A reusable map widget module for Fyne applications with mouse drag support and location-based centering.

## Features

- 🗺️ OpenStreetMap tile rendering
- 🖱️ Mouse drag to pan
- 🔍 Zoom in/out controls
- 📍 Center on specific latitude/longitude coordinates
- 📱 Mobile-ready (works on phone apps)

## Installation

```go
import "fyne.io/x/fyne/mymap"
```

## Usage

### Basic Map

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2"
    "fyne.io/x/fyne/mymap"
)

func main() {
    a := app.New()
    w := a.NewWindow("My Map")

    // Create map with OpenStreetMap tiles
    m := mymap.NewMapWithOptions(
        mymap.WithOsmTiles(),
        mymap.WithZoomButtons(true),
        mymap.WithScrollButtons(false), // Use mouse drag instead
    )

    w.SetContent(m)
    w.Resize(fyne.NewSize(800, 600))
    w.ShowAndRun()
}
```

### Center on a Location

```go
// Center on Sousse, Tunisia
m := mymap.NewMapWithOptions(
    mymap.WithOsmTiles(),
    mymap.WithZoomButtons(true),
    mymap.WithScrollButtons(false),
)

// Set zoom level first
m.Zoom(12)

// Then center on coordinates
m.CenterOnLocation(35.83, 10.64) // Sousse: 35.83°N, 10.64°E

w.SetContent(m)
```

### Available Options

- `WithOsmTiles()` - Use OpenStreetMap tiles (default)
- `WithTileSource(url)` - Use custom tile source
- `WithZoomButtons(bool)` - Show/hide zoom buttons
- `WithScrollButtons(bool)` - Show/hide arrow buttons
- `WithAttribution(bool, label, url)` - Configure attribution display
- `WithHTTPClient(client)` - Use custom HTTP client

## API

### Methods

- `Zoom(level int)` - Set zoom level (0-19)
- `ZoomIn()` - Zoom in one level
- `ZoomOut()` - Zoom out one level
- `PanEast()`, `PanNorth()`, `PanSouth()`, `PanWest()` - Pan by 1/4 tile
- `CenterOnLocation(lat, lon float64)` - Center map on coordinates

### Helper Functions

- `LatLonToTile(lat, lon float64, zoom int) (x, y int)` - Convert coordinates to tile positions

## Example: Mobile App Integration

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2"
    "fyne.io/x/fyne/mymap"
)

func main() {
    a := app.New()
    w := a.NewWindow("Location Tracker")

    m := mymap.NewMapWithOptions(
        mymap.WithOsmTiles(),
        mymap.WithZoomButtons(true),
        mymap.WithScrollButtons(false),
    )

    // Start at a specific location
    m.Zoom(15)
    m.CenterOnLocation(35.83, 10.64)

    w.SetContent(m)
    w.ShowAndRun()
}
```

## Coordinates Reference

Some Tunisian cities:

- Tunis: 36.8°N, 10.18°E
- Sousse: 35.83°N, 10.64°E
- Sfax: 34.74°N, 10.76°E
- Bizerte: 37.27°N, 9.87°E

## License

See parent project license.
