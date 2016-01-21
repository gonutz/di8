package main

import (
	"fmt"
	"github.com/AllenDang/gform"
	"github.com/AllenDang/w32"
	"github.com/gonutz/di8"
	"unsafe"
)

func main() {
	gform.Init()

	form := gform.NewForm(nil)
	form.Show()
	form.OnClose().Bind(func(arg *gform.EventArg) {
		w32.DestroyWindow(form.Handle())
	})

	check(di8.Init())
	defer di8.Close()

	dinput, err := di8.Create(unsafe.Pointer(w32.GetModuleHandle("")))
	defer dinput.Release()
	check(err)

	dinput.EnumDevices(
		di8.DEVCLASS_ALL,
		func(device di8.DeviceInstance) bool {
			fmt.Printf("%v (%v)\n", device.InstanceName, device.ProductName)
			return true
		},
		di8.EDFL_ALLDEVICES,
	)

	gform.RunMainLoop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
