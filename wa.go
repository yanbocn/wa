package wa

import (
	"time"
)

// Sleep time.Sleep t second
func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

//GetPixelColor get color
func GetPixelColor(x, y int32) uint32 {
	h := getDesktopWindowa()
	dc := getDCa(h)
	co := getPixela(dc, x, y)
	releaseDCa(0, dc)
	return co
}
