package di8

//#include "dinput_wrapper.h"
import "C"

// EnumDevicesCallback returns true to continue enumerating and false to stop.
// You can use the GuidInstance variable in the DeviceInstance to create the
// device.
type EnumDevicesCallback func(device DeviceInstance) bool

var currentEnumDevicesCallback EnumDevicesCallback

//export enumDevicesCallbackGo
func enumDevicesCallbackGo(lpddi C.LPCDIDEVICEINSTANCE, pvRef C.LPVOID) C.BOOL {
	var device DeviceInstance
	device.fromC(lpddi)
	if currentEnumDevicesCallback(device) {
		return C.DIENUM_CONTINUE
	}
	return C.DIENUM_STOP
}

// EnumObjectsCallback returns true to continue enumerating and false to stop.
type EnumEffectsCallback func(effect EffectInfo) bool

var currentEnumEffectsCallback EnumEffectsCallback

//export enumEffectsCallbackGo
func enumEffectsCallbackGo(lpdei C.LPCDIEFFECTINFO, pvRef C.LPVOID) C.BOOL {
	var info EffectInfo
	info.fromC(lpdei)
	if currentEnumEffectsCallback(info) {
		return C.DIENUM_CONTINUE
	}
	return C.DIENUM_STOP
}

// EnumObjectsCallback returns true to continue enumerating and false to stop.
type EnumObjectsCallback func(object DeviceObjectInstance) bool

var currentEnumObjectsCallback EnumObjectsCallback

//export enumObjectsCallbackGo
func enumObjectsCallbackGo(lpddoi C.LPCDIDEVICEOBJECTINSTANCE, pvRef C.LPVOID) C.BOOL {
	var object DeviceObjectInstance
	object.fromC(lpddoi)
	if currentEnumObjectsCallback(object) {
		return C.DIENUM_CONTINUE
	}
	return C.DIENUM_STOP
}
