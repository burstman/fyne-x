package mousemappan

import "math"

// LatLonToTile converts latitude and longitude to tile coordinates at a given zoom level
func LatLonToTile(lat, lon float64, zoom int) (x, y int) {
	n := math.Pow(2, float64(zoom))
	x = int((lon + 180.0) / 360.0 * n)
	latRad := lat * math.Pi / 180.0
	y = int((1.0 - math.Log(math.Tan(latRad)+1.0/math.Cos(latRad))/math.Pi) / 2.0 * n)
	return x, y
}

// CenterOnLocation centers the map on the given latitude/longitude coordinates
func (m *Map) CenterOnLocation(lat, lon float64) {
	tileX, tileY := LatLonToTile(lat, lon, m.zoom)
	count := 1 << m.zoom
	centerX := int(float32(count)/2 - 0.5)
	centerY := int(float32(count)/2 - 0.5)
	deltaX := tileX - centerX
	deltaY := tileY - centerY

	// Pan to location (4 clicks per tile since each click = 1/4 tile)
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
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
