# Using MyMap in Your Phone App

## Quick Setup

1. **Import the module** in your phone app:

```go
import "fyne.io/x/fyne/mymap"
```

2. **Create the map** in your app:

```go
func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("My Location App")

    // Create map centered on your location
    mapWidget := mymap.NewMapWithOptions(
        mymap.WithOsmTiles(),
        mymap.WithZoomButtons(true),
        mymap.WithScrollButtons(false), // Touch drag works better on mobile
    )

    // Set initial location (e.g., Sousse, Tunisia)
    mapWidget.Zoom(14) // Higher zoom for city view
    mapWidget.CenterOnLocation(35.83, 10.64)

    myWindow.SetContent(mapWidget)
    myWindow.ShowAndRun()
}
```

## Mobile-Specific Tips

### Touch Controls

- **Drag to pan**: Already works! The mouse drag events work for touch on mobile
- **Pinch to zoom**: Use the zoom buttons for now
- **No arrows needed**: Set `WithScrollButtons(false)` - drag is better on mobile

### Zoom Levels for Mobile

- **City view**: Zoom 14-16
- **Neighborhood**: Zoom 16-18
- **Street level**: Zoom 18-19
- **Country view**: Zoom 6-8

### Performance

The module automatically handles:

- Tile caching (via mapcache.go)
- Efficient rendering
- Memory management

### Example: Location Tracking App

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    "fyne.io/x/fyne/mymap"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Location Tracker")

    // Create map
    mapWidget := mymap.NewMapWithOptions(
        mymap.WithOsmTiles(),
        mymap.WithZoomButtons(true),
        mymap.WithScrollButtons(false),
    )

    mapWidget.Zoom(15)

    // Create location buttons
    sousseBt := widget.NewButton("Go to Sousse", func() {
        mapWidget.CenterOnLocation(35.83, 10.64)
    })

    tunisBt := widget.NewButton("Go to Tunis", func() {
        mapWidget.CenterOnLocation(36.8, 10.18)
    })

    sfaxBt := widget.NewButton("Go to Sfax", func() {
        mapWidget.CenterOnLocation(34.74, 10.76)
    })

    buttons := container.NewHBox(sousseBt, tunisBt, sfaxBt)
    content := container.NewBorder(buttons, nil, nil, nil, mapWidget)

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
```

## Building for Mobile

### Android

```bash
fyne package -os android -appID com.yourcompany.locationapp
```

### iOS

```bash
fyne package -os ios -appID com.yourcompany.locationapp
```

## Integration with GPS

To integrate with device GPS, you can update the map location:

```go
// When you get GPS coordinates (lat, lon)
func updateLocation(mapWidget *mymap.Map, lat, lon float64) {
    mapWidget.CenterOnLocation(lat, lon)
}
```

## File Structure

Your app should look like:

```
your-phone-app/
├── main.go
├── go.mod
└── (fyne-x/mymap module is imported)
```

The mymap module files are in:

```
fyne-x/mymap/
├── map.go          # Main map widget
├── mapcache.go     # Tile caching
├── mapbutton.go    # UI buttons
├── helpers.go      # Helper functions (CenterOnLocation, etc.)
└── README.md       # Full documentation
```

## That's it!

The module is self-contained and ready to use in your phone app. Just import it and start using `mymap.NewMapWithOptions()`.
