package wa

import (
	"github.com/lxn/win"
)

//GetWindowHWND x
func GetWindowHWND() win.HWND {
	return win.GetForegroundWindow()
}

//GetActivePid get pid
func GetActivePid() uint32 {
	a := win.GetForegroundWindow()
	var pid uint32
	win.GetWindowThreadProcessId(a, &pid)
	return pid
}

//GetWindowPid get pid
func GetWindowPid(a win.HWND) uint32 {
	var pid uint32
	win.GetWindowThreadProcessId(a, &pid)
	return pid
}

