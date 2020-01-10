package wa

import (
	"time"
)

// Sleep time.Sleep t second
func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

//GetPixelColor get color
func GetPixelColor(x, y int32) COLORREF {
	h := GetDesktopWindow()
	dc := GetDC(h)
	co := GetPixel(dc, x, y)
	return co
}
