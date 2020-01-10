package wa

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type point struct {
	X, Y int32
}

type (
	HWND     uint32
	HDC      uint32
	COLORREF uint32
)

// INPUT Type
const (
	inputMouse    = 0
	inputKeyBoard = 1
	inputHardware = 2
)

// MOUSEINPUT DwFlags
const (
	mouseEventAbsolute       = 0x8000
	mouseEventHwheel         = 0x1000
	mouseEventMove           = 0x0001
	mouseEventMoveNoCoalesce = 0x2000
	mouseEventLeftDown       = 0x0002
	mouseEventLeftUp         = 0x0004
	mouseEventRightDown      = 0x0008
	mouseEventRightUp        = 0x0010
	mouseEventMiddleDown     = 0x0020
	mouseEventMiddleUp       = 0x0040
	mouseEventVirtualdesk    = 0x4000
	mouseEventWheel          = 0x0800
	mouseEventXDown          = 0x0080
	mouseEventXUp            = 0x0100
)

// KEYBDINPUT DwFlags
const (
	keyEventExtendedKey = 0x0001
	keyEventKeyUp       = 0x0002
	keyEventScancode    = 0x0008
	keyEventUnicode     = 0x0004
)

//KEYBD_INPUT
type keyBdInput struct {
	Type uint32
	Ki   keyBdIn
}

//KEYBDINPUT
type keyBdIn struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
	Unused      [8]byte
}

//MOUSE_INPUT
type mouseInput struct {
	Type uint32
	Mi   mouseIn
}

//MOUSEINPUT
type mouseIn struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

var (
	libuser32        *windows.LazyDLL
	sendInput        *windows.LazyProc
	setCursorPos     *windows.LazyProc
	getCursorPos     *windows.LazyProc
	getActiveWindow  *windows.LazyProc
	getDC            *windows.LazyProc
	getDesktopWindow *windows.LazyProc

	libgdi32 *windows.LazyDLL
	getPixel *windows.LazyProc
)

func init() {
	libuser32 = windows.NewLazySystemDLL("user32.dll")
	sendInput = libuser32.NewProc("SendInput")
	setCursorPos = libuser32.NewProc("SetCursorPos")
	getCursorPos = libuser32.NewProc("GetCursorPos")
	getActiveWindow = libuser32.NewProc("GetActiveWindow")
	getDC = libuser32.NewProc("GetDC")
	getDesktopWindow = libuser32.NewProc("GetDesktopWindow")

	libgdi32 = windows.NewLazySystemDLL("gdi32.dll")
	getPixel = libgdi32.NewProc("GetPixel")
}

//SendInput xx
func SendInput(nInput uint32, pInput unsafe.Pointer, cbSize int32) uint32 {
	ret, _, _ := syscall.Syscall(sendInput.Addr(), 3,
		uintptr(nInput),
		uintptr(pInput),
		uintptr(cbSize))
	return uint32(ret)
}

//SetCursorPos x, y
func SetCursorPos(x, y int32) bool {
	ret, _, _ := syscall.Syscall(setCursorPos.Addr(), 2,
		uintptr(x),
		uintptr(y),
		0)
	return ret != 0
}

//GetCursorPos x
func GetCursorPos(lpPoint *point) bool {
	ret, _, _ := syscall.Syscall(getCursorPos.Addr(), 1,
		uintptr(unsafe.Pointer(lpPoint)),
		0,
		0)
	return ret != 0
}

func GetPixel(hdc HDC, nXPos, nYPos int32) COLORREF {
	ret, _, _ := syscall.Syscall(getPixel.Addr(), 3,
		uintptr(hdc),
		uintptr(nXPos),
		uintptr(nYPos))
	return COLORREF(ret)
}

func GetDesktopWindow() HWND {
	ret, _, _ := syscall.Syscall(getDesktopWindow.Addr(), 0,
		0,
		0,
		0)
	return HWND(ret)
}

func GetActiveWindow() HWND {
	ret, _, _ := syscall.Syscall(getActiveWindow.Addr(), 0,
		0,
		0,
		0)
	return HWND(ret)
}

func GetDC(hWnd HWND) HDC {
	ret, _, _ := syscall.Syscall(getDC.Addr(), 1,
		uintptr(hWnd),
		0,
		0)
	return HDC(ret)
}
