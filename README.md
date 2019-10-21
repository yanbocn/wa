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

package wa // import "wa"

		func GetCursorPos() win.POINT
		func KeyBd(s string, args ...string)
		func MouseClick(args ...string)
		func MouseClickPos(x, y int32)
		func SetCursorPos(x, y int32)
		func Sleep(t int)
		func ToGBK(rd io.Reader) *mahonia.Reader
		func TypeStr(str string)
		func GetPixel(x, y int32) win.COLORREF
		