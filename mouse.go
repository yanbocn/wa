package wa

import (
	"unsafe"
)

//GetMousePos mouse pos
func GetMousePos() (int32, int32) {
	var pt point
	GetCursorPos(&pt)
	return pt.X, pt.Y
}

//SetMousePos mouse pos
func SetMousePos(x, y int32) {
	SetCursorPos(x, y)
}

//MouseClickPos x, y click
func MouseClickPos(x, y int32) {
	SetCursorPos(x, y)
	MouseClick()
}

//MouseClick click
func MouseClick(args ...string) {
	switch len(args) {
	case 0:
		liftClick()
		return
	case 1:
		switch args[0] {
		case "left":
			liftClick()
			return
		case "right":
			rightClick()
			return
		}
	case 2:
		if args[0] == "left" && args[1] == "double" {
			liftClick()
			liftClick()
			return
		}
		return
	}
}

func liftClick() {
	var input, input2 mouseInput
	input.Mi.DwFlags = mouseEventLeftDown
	input2.Mi.DwFlags = mouseEventLeftUp
	SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
	SendInput(1, unsafe.Pointer(&input2), int32(unsafe.Sizeof(input2)))
}

func rightClick() {
	var input, input2 mouseInput
	input.Mi.DwFlags = mouseEventRightDown
	input2.Mi.DwFlags = mouseEventRightUp
	SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
	SendInput(1, unsafe.Pointer(&input2), int32(unsafe.Sizeof(input2)))
}
