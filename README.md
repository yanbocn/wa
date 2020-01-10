# wa

	$ go get github.com/yanbocn/wa

demo.go

        package main
        
        import "github.com/yanbocn/wa"
        
        func main() {
	        wa.MouseClickPos(50, 50)  
	        wa.Sleep(1)
	        wa.TypeStr("ahahaha")
        }


package wa // import "github.com/yanbocn/wa"

		func GetCursorPos(lpPoint *point) bool
		func GetMousePos() (int32, int32)
		func InputKey(s string, args ...string)
		func InputStr(str string)
		func MouseClick(args ...string)
		func MouseClickPos(x, y int32)
		func SendInput(nInput uint32, pInput unsafe.Pointer, cbSiz int32) uint32
		func SetCursorPos(x, y int32) bool
		func SetMousePos(x, y int32)
		func Sleep(t int)
		type COLORREF uint32
    	func GetPixel(hdc HDC, nXPos, nYPos int32) COLORREF
    	func GetPixelColor(x, y int32) COLORREF
		type HDC uint32
    	func GetDC(hWnd HWND) HDC
		type HWND uint32
    	func GetActiveWindow() HWND
    	func GetDesktopWindow() HWND