package main

//#cgo CFLAGS: -DUNICODE
//#include "win.h"
import "C"
import (
	"fmt"
	"github.com/gonutz/di8"
	"unsafe"
)

func main() {
	var window C.HWND
	if errorCode := C.openWindow(&window); errorCode != C.OK {
		panic(errorCode)
	}

	check(di8.Init())
	defer di8.Close()

	dinput, err := di8.Create(unsafe.Pointer(C.GetModuleHandle(nil)))
	check(err)
	defer dinput.Release()

	dinput.EnumDevices(
		di8.DEVCLASS_ALL,
		func(device di8.DeviceInstance) bool {
			fmt.Printf("%v (%v)\n", device.InstanceName, device.ProductName)
			return true
		},
		di8.EDFL_ALLDEVICES,
	)

	var msg C.MSG
	C.PeekMessage(&msg, nil, 0, 0, C.PM_NOREMOVE)
	for msg.message != C.WM_QUIT {
		if C.PeekMessage(&msg, nil, 0, 0, C.PM_REMOVE) != 0 {
			C.TranslateMessage(&msg)
			C.DispatchMessage(&msg)
		} else {
		}
	}
}

//export messageCallbackGo
func messageCallbackGo(window C.HWND, message C.UINT, w C.WPARAM, l C.LPARAM) C.LRESULT {
	if message == C.WM_DESTROY {
		C.PostQuitMessage(0)
	}
	return C.DefWindowProc(window, message, w, l)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
