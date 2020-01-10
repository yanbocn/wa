# wa

	$ go get github.com/yanbocn/wa

demo.go

        package main
        
        import "github.com/yanbocn/wa"
        
        func main() {
	        wa.MouseClickPos(50, 50)
	        wa.Sleep(1)
	        wa.InputStr("ahahaha")
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
    	func GetPixelColor(x, y int32) COLORREF