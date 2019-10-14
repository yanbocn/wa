package wa

import (
	"unsafe"

	"github.com/lxn/win"
)

const (
	ent  = 0x0d
	back = 8
	ctrl = 17
	alt  = 18
)

//KeyBd key
func KeyBd(s string, args ...string) {
	if len(args) == 0 {
		var input [2]win.KEYBD_INPUT
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
		input[0].Type = win.INPUT_KEYBOARD
		input[1].Type = win.INPUT_KEYBOARD
		input[1].Ki.DwFlags = win.KEYEVENTF_KEYUP
		win.SendInput(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
		win.SendInput(1, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
		return
	}
	t := []byte(s)
	var input [4]win.KEYBD_INPUT
	input[0].Ki.WVk, input[1].Ki.WVk = uint16(t[0]), uint16(t[0])
	switch args[0] {
	case "ctrl":
		input[2].Ki.WVk, input[3].Ki.WVk = ctrl, ctrl
		break
	case "alt":
		input[2].Ki.WVk, input[3].Ki.WVk = alt, alt
		break
	}
	input[0].Type = win.INPUT_KEYBOARD
	input[1].Type = win.INPUT_KEYBOARD
	input[2].Type = win.INPUT_KEYBOARD
	input[3].Type = win.INPUT_KEYBOARD
	win.SendInput(1, unsafe.Pointer(&input[2]), int32(unsafe.Sizeof(input[2])))
	win.SendInput(1, unsafe.Pointer(&input[0]), int32(unsafe.Sizeof(input[0])))
	input[1].Ki.DwFlags = win.KEYEVENTF_KEYUP
	input[3].Ki.DwFlags = win.KEYEVENTF_KEYUP
	win.SendInput(1, unsafe.Pointer(&input[1]), int32(unsafe.Sizeof(input[1])))
	win.SendInput(1, unsafe.Pointer(&input[3]), int32(unsafe.Sizeof(input[3])))

}
