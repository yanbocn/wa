package wa

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

//TypeStr ("XX")
func TypeStr(str string) {
	s := syscall.StringToUTF16(str)
	for _, a := range s {
		unicodeType(a)
	}
}

func unicodeType(value uint16) {
	var input [2]win.KEYBD_INPUT

	input[0].Type = win.INPUT_KEYBOARD
	input[0].Ki.WVk = 0
	input[0].Ki.WScan = value
	input[0].Ki.DwFlags = 0x4 // KEYEVENTF_UNICODE;

	input[1].Type = win.INPUT_KEYBOARD
	input[1].Ki.WVk = 0
	input[1].Ki.WScan = value
	input[1].Ki.DwFlags = win.KEYEVENTF_KEYUP | 0x4 // KEYEVENTF_UNICODE;
	win.SendInput(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	win.SendInput(2, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
}
