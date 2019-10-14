package wa

import (
	"io"
	"time"

	"github.com/axgle/mahonia"
)

// Sleep time.Sleep t second
func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

//ToGBK ~
func ToGBK(rd io.Reader) *mahonia.Reader {
	decoder := mahonia.NewDecoder("gbk")
	return decoder.NewReader(rd)
}
