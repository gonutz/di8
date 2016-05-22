package di8

/*
#cgo LDFLAGS: -ldinput8 -ldxguid

#include "dinput_wrapper.h"

typedef HRESULT WINAPI direct_input8_create(
	HINSTANCE hinst,
	DWORD dwVersion,
	REFIID riidltf,
	void** ppvOut,
	LPUNKNOWN punkOuterr
);

HRESULT create(
		void* DirectInput8CreateFuncPtr,
		HINSTANCE instance,
		IDirectInput8** obj) {
	direct_input8_create* DirectInput8Create =
		(direct_input8_create*) DirectInput8CreateFuncPtr;
	return DirectInput8Create(
		instance,
		DIRECTINPUT_VERSION,
		&IID_IDirectInput8,
		(void**)obj,
		0
	);
}

HRESULT IDirectInput8CreateDevice(
		IDirectInput8* obj,
		GUID* rguid,
		IDirectInputDevice8** lplpDirectInputDevice,
		LPUNKNOWN pUnkOuter) {
	return obj->lpVtbl->CreateDevice(
		obj,
		rguid,
		lplpDirectInputDevice,
		pUnkOuter);
}

BOOL enumDevicesCallbackGo(LPCDIDEVICEINSTANCE, void*);

HRESULT IDirectInput8EnumDevices(
		IDirectInput8* obj,
		DWORD dwDevType,
		void* pvRef,
		DWORD dwFlags) {
	return obj->lpVtbl->EnumDevices(
		obj,
		dwDevType,
		(LPDIENUMDEVICESCALLBACK)enumDevicesCallbackGo,
		pvRef,
		dwFlags);
}

HRESULT IDirectInput8FindDevice(
		IDirectInput8* obj,
		REFGUID rguidClass,
		CHAR* ptszName,
		LPGUID pguidInstance) {
	return obj->lpVtbl->FindDevice(obj, rguidClass, ptszName, pguidInstance);
}

HRESULT IDirectInput8GetDeviceStatus(
		IDirectInput8* obj,
		REFGUID rguidInstance) {
	return obj->lpVtbl->GetDeviceStatus(obj, rguidInstance);
}

HRESULT IDirectInput8RunControlPanel(
		IDirectInput8* obj,
			HWND hwndOwner,
			DWORD dwFlags) {
	return obj->lpVtbl->RunControlPanel(obj, hwndOwner, dwFlags);
}

void IDirectInput8Release(IDirectInput8* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"

import (
	"syscall"
	"unsafe"
)

var dll syscall.Handle

func Init() (err error) {
	dll, err = syscall.LoadLibrary("dinput8.dll")
	return
}

func Close() {
	syscall.FreeLibrary(dll)
}

type DirectInput struct {
	handle *C.IDirectInput8
}

func (obj DirectInput) Release() {
	C.IDirectInput8Release(obj.handle)
}

func Create(windowInstance unsafe.Pointer) (obj DirectInput, err error) {
	DirectInput8Create, err := syscall.GetProcAddress(dll, "DirectInput8Create")
	if err != nil {
		return DirectInput{}, err
	}
	err = toError(C.create(
		unsafe.Pointer(DirectInput8Create),
		C.HINSTANCE(windowInstance),
		&obj.handle,
	))
	return
}

// ConfigureDevices is deprecated since Vista and therefore omitted.

func (obj DirectInput) CreateDevice(guid GUID) (device Device, err error) {
	cGuid := guid.toC()
	err = toError(C.IDirectInput8CreateDevice(
		obj.handle,
		&cGuid,
		&device.handle,
		nil,
	))
	return
}

// devType: DEVCLASS_* or DEVTYPE_*
// flags: EDFL_*
func (obj DirectInput) EnumDevices(
	devType uint32,
	callback EnumDevicesCallback,
	flags uint32,
) (
	err error,
) {
	currentEnumDevicesCallback = callback
	err = toError(C.IDirectInput8EnumDevices(
		obj.handle,
		C.DWORD(devType),
		nil,
		C.DWORD(flags),
	))
	return
}

// TODO EnumDevicesBySemantics

func (obj DirectInput) FindDevice(
	guid GUID,
	name string,
) (
	guidDevice GUID,
	err error,
) {
	cGuid := guid.toC()
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cGuidDevice C.GUID
	err = toError(C.IDirectInput8FindDevice(
		obj.handle,
		&cGuid,
		(*C.CHAR)(cName),
		&cGuidDevice,
	))
	guidDevice.fromC(&cGuidDevice)
	return
}

func (obj DirectInput) GetDeviceStatus(guid GUID) (err error) {
	cGuid := guid.toC()
	err = toError(C.IDirectInput8GetDeviceStatus(obj.handle, &cGuid))
	return
}

func (obj DirectInput) RunControlPanel(
	ownerWindow unsafe.Pointer,
	flags uint32,
) (
	err error,
) {
	err = toError(C.IDirectInput8RunControlPanel(
		obj.handle,
		C.HWND(ownerWindow),
		C.DWORD(flags),
	))
	return
}

func toError(hr C.HRESULT) error {
	if hr == C.DI_OK {
		return nil
	}
	return Error(hr)
}
