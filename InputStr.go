package wa

import (
	"syscall"
	"unsafe"
)

//InputStr ("XX")
func InputStr(str string) {
	s := syscall.StringToUTF16(str)
	for _, a := range s {
		unicodeType(a)
	}
}

func unicodeType(value uint16) {
	var input [2]keyBdInput

	input[0].Type = inputKeyBoard
	input[0].Ki.WVk = 0
	input[0].Ki.WScan = value
	input[0].Ki.DwFlags = keyEventUnicode

	input[1].Type = inputKeyBoard
	input[1].Ki.WVk = 0
	input[1].Ki.WScan = value
	input[1].Ki.DwFlags = keyEventKeyUp | keyEventUnicode
	sendInputa(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	sendInputa(1, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
}
