package wa

import (
	"unsafe"
)

const (
	ent  = 13
	back = 8
	ctrl = 17
	alt  = 18
)

//InputKey key
func InputKey(s string, args ...string) {
	if len(args) == 0 {
		var input [2]keyBdInput
		switch s {
		case "enter":
			input[0].Ki.WVk, input[1].Ki.WVk = ent, ent
			break
		case "backspace":
			input[0].Ki.WVk, input[1].Ki.WVk = back, back
			break
		default:
			t := []byte(s)
			input[0].Ki.WVk, input[1].Ki.WVk = uint16(t[0]), uint16(t[0])
			break
		}
		input[0].Type = inputKeyBoard
		input[1].Type = inputKeyBoard
		input[1].Ki.DwFlags = keyEventKeyUp
		sendInputa(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
		sendInputa(1, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
		return
	}
	t := []byte(s)
	var input [4]keyBdInput
	input[0].Ki.WVk, input[1].Ki.WVk = uint16(t[0]), uint16(t[0])
	switch args[0] {
	case "ctrl":
		input[2].Ki.WVk, input[3].Ki.WVk = ctrl, ctrl
		break
	case "alt":
		input[2].Ki.WVk, input[3].Ki.WVk = alt, alt
		break
	}
	input[0].Type = inputKeyBoard
	input[1].Type = inputKeyBoard
	input[2].Type = inputKeyBoard
	input[3].Type = inputKeyBoard
	sendInputa(1, unsafe.Pointer(&input[2]), int32(unsafe.Sizeof(input[2])))
	sendInputa(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	input[1].Ki.DwFlags = keyEventKeyUp
	input[3].Ki.DwFlags = keyEventKeyUp
	sendInputa(1, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
	sendInputa(1, unsafe.Pointer(&input[3]), int32(unsafe.Sizeof(input[3])))
}
