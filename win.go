package wa

import (
	"github.com/lxn/win"
)

//GetPixel x, y
func GetPixel(x, y int32) win.COLORREF {
	return win.GetPixel(win.GetDC(0), x, y)
}
