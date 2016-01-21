package di8

import "strconv"

type Error int64

func (e Error) Error() string {
	switch int64(e) {
	case ERR_ACQUIRED:
		return "ERR_ACQUIRED: The operation cannot be performed while the device is acquired."
	case ERR_ALREADYINITIALIZED:
		return "ERR_ALREADYINITIALIZED: This object is already initialized"
	case ERR_BADDRIVERVER:
		return "ERR_BADDRIVERVER: The object could not be created due to an incompatible driver version or mismatched or incomplete driver components."
	case ERR_BETADIRECTINPUTVERSION:
		return "ERR_BETADIRECTINPUTVERSION: The object could not be created due to an incompatible driver version or mismatched or incomplete driver components."
	case ERR_DEVICEFULL:
		return "ERR_DEVICEFULL: The device is full."
	case ERR_DEVICENOTREG:
		return "ERR_DEVICENOTREG: The device or device instance is not registered with DirectInput. This value is equal to the REGDB_E_CLASSNOTREG standard COM return value."
	case ERR_EFFECTPLAYING:
		return "ERR_EFFECTPLAYING: The parameters were updated in memory but were not downloaded to the device because the device does not support updating an effect while it is still playing."
	case ERR_GENERIC:
		return "ERR_GENERIC: An undetermined error occurred inside the DirectInput subsystem. This value is equal to the E_FAIL standard COM return value."
	case ERR_HANDLEEXISTS:
		return "ERR_HANDLEEXISTS: The device already has an event notification associated with it. This value is equal to the E_ACCESSDENIED standard COM return value."
	case ERR_HASEFFECTS:
		return "ERR_HASEFFECTS: The device cannot be reinitialized because effects are attached to it."
	case ERR_INCOMPLETEEFFECT:
		return "ERR_INCOMPLETEEFFECT: The effect could not be downloaded because essential information is missing. For example, no axes have been associated with the effect, or no type-specific information has been supplied."
	case ERR_INPUTLOST:
		return "ERR_INPUTLOST: Access to the input device has been lost. It must be reacquired."
	case ERR_INVALIDPARAM:
		return "ERR_INVALIDPARAM: An invalid parameter was passed to the returning function, or the object was not in a state that permitted the function to be called. This value is equal to the E_INVALIDARG standard COM return value."
	case ERR_MAPFILEFAIL:
		return "ERR_MAPFILEFAIL: An error has occurred either reading the vendor-supplied action-mapping file for the device or reading or writing the user configuration mapping file for the device."
	case ERR_MOREDATA:
		return "ERR_MOREDATA: Not all the requested information fit into the buffer."
	case ERR_NOAGGREGATION:
		return "ERR_NOAGGREGATION: This object does not support aggregation."
	case ERR_NOINTERFACE:
		return "ERR_NOINTERFACE: The object does not support the specified interface. This value is equal to the E_NOINTERFACE standard COM return value."
	case ERR_NOTACQUIRED:
		return "ERR_NOTACQUIRED: The operation cannot be performed unless the device is acquired."
	case ERR_NOTBUFFERED:
		return "ERR_NOTBUFFERED: The device is not buffered. Set the DIPROP_BUFFERSIZE property to enable buffering."
	case ERR_NOTDOWNLOADED:
		return "ERR_NOTDOWNLOADED: The effect is not downloaded."
	case ERR_NOTEXCLUSIVEACQUIRED:
		return "ERR_NOTEXCLUSIVEACQUIRED: The operation cannot be performed unless the device is acquired in DISCL_EXCLUSIVE mode."
	case ERR_NOTFOUND:
		return "ERR_NOTFOUND, ERR_OBJECTNOTFOUND: The requested object does not exist."
	case ERR_NOTINITIALIZED:
		return "ERR_NOTINITIALIZED: This object has not been initialized."
	default:
		return "Unknown error code " + strconv.Itoa(int(e))
	}
}

/*
//#include "include.h"
import "C"
import "strconv"

type Error string

func (e Error) Error() string {
	return string(e)
}

var (
	ErrAcquired               Error = "The operation cannot be performed while the device is acquired."
	ErrAlreadyInitialized     Error = "This object is already initialized"
	ErrBadDriverVersion       Error = "The object could not be created due to an incompatible driver version or mismatched or incomplete driver components."
	ErrBetaDirectInputVersion Error = "The application was written for an unsupported prerelease version of DirectInput."
	ErrDeviceFull             Error = "The device is full."
	ErrDeviceNotRegistered    Error = "The device or device instance is not registered with DirectInput. This value is equal to the REGDB_E_CLASSNOTREG standard COM return value."
	ErrEffectPlaying          Error = "The parameters were updated in memory but were not downloaded to the device because the device does not support updating an effect while it is still playing."
	ErrGeneric                Error = "An undetermined error occurred inside the DirectInput subsystem. This value is equal to the E_FAIL standard COM return value."
	ErrHandleExists           Error = "The device already has an event notification associated with it. This value is equal to the E_ACCESSDENIED standard COM return value."
	ErrHasEffects             Error = "The device cannot be reinitialized because effects are attached to it."
	ErrIncompleteEffect       Error = "The effect could not be downloaded because essential information is missing. For example, no axes have been associated with the effect, or no type-specific information has been supplied."
	ErrInputLost              Error = "Access to the input device has been lost. It must be reacquired."
	ErrInvalidParam           Error = "An invalid parameter was passed to the returning function, or the object was not in a state that permitted the function to be called. This value is equal to the E_INVALIDARG standard COM return value."
	ErrMapFileFail            Error = "An error has occurred either reading the vendor-supplied action-mapping file for the device or reading or writing the user configuration mapping file for the device."
	ErrMoreData               Error = "Not all the requested information fit into the buffer."
	ErrNoAggregation          Error = "This object does not support aggregation."
	ErrNoInterface            Error = "The object does not support the specified interface. This value is equal to the E_NOINTERFACE standard COM return value."
	ErrNotAcquired            Error = "The operation cannot be performed unless the device is acquired."
	ErrNotBuffered            Error = "The device is not buffered. Set the DIPROP_BUFFERSIZE property to enable buffering."
	ErrNotDownloaded          Error = "The effect is not downloaded."
	ErrNotExclusiveAcquired   Error = "The operation cannot be performed unless the device is acquired in DISCL_EXCLUSIVE mode."
	ErrNotFound               Error = "The requested object does not exist."
	ErrNotInitialized         Error = "This object has not been initialized."
	ErrObjectNotFound               = ErrNotFound
	ErrOldDirectInputVersion  Error = "The application requires a newer version of DirectInput."
	// TODO this is the same numerical value as for ErrHandleExists
	//ErrOtherAppHasPriority    Error = "Another application has a higher priority level, preventing this call from succeeding. This value is equal to the E_ACCESSDENIED standard DirectInput return value. This error can be returned when an application has only foreground access to a device but is attempting to acquire the device while in the background."
	ErrOutOfMemory Error = "The DirectInput subsystem could not allocate sufficient memory to complete the call. This value is equal to the E_OUTOFMEMORY standard COM return value."
	// TODO this is the same numerical value as for ErrHandleExists
	//ErrReadOnly    Error = "The specified property cannot be changed. This value is equal to the E_ACCESSDENIED standard COM return value."
	ErrReportFull  Error = "More information was requested to be sent than can be sent to the device."
	ErrUnplugged   Error = "The operation could not be completed because the device is not plugged in."
	ErrUnsupported Error = "The function called is not supported at this time. This value is equal to the E_NOTIMPL standard COM return value."
	ErrHandle      Error = "The HWND parameter is not a valid top-level window that belongs to the process."
	ErrPending     Error = "Data is not yet available."
	ErrPointer     Error = "An invalid pointer, usually NULL, was passed as a parameter. "

	ErrBufferOverflow  Error = "The device buffer overflowed and some input was lost. This value is equal to the S_FALSE standard COM return value."
	ErrDownloadSkipped Error = "The parameters of the effect were successfully updated, but the effect could not be downloaded because the associated device was not acquired in exclusive mode."
	ErrEffectRestarted Error = "The effect was stopped, the parameters were updated, and the effect was restarted."
	ErrNotAttached     Error = "The device exists but is not currently attached to the user's computer. This value is equal to the S_FALSE standard COM return value."
)

func toError(result C.HRESULT) error {
	if result == C.DI_OK {
		return nil
	}
	switch result {
	case C.DIERR_ACQUIRED:
		return ErrAcquired
	case C.DIERR_ALREADYINITIALIZED:
		return ErrAlreadyInitialized
	case C.DIERR_BADDRIVERVER:
		return ErrBadDriverVersion
	case C.DIERR_BETADIRECTINPUTVERSION:
		return ErrBetaDirectInputVersion
	//case C.DIERR_DEVICEFULL:
	//return ErrDeviceFull
	case C.DIERR_DEVICENOTREG:
		return ErrDeviceNotRegistered
	//case C.DIERR_EFFECTPLAYING:
	//	return ErrEffectPlaying
	case C.DIERR_GENERIC:
		return ErrGeneric
	case C.DIERR_HANDLEEXISTS:
		return ErrHandleExists
	//case C.DIERR_HASEFFECTS:
	//	return ErrHasEffects
	//case C.DIERR_INCOMPLETEEFFECT:
	//	return ErrIncompleteEffect
	case C.DIERR_INPUTLOST:
		return ErrInputLost
	case C.DIERR_INVALIDPARAM:
		return ErrInvalidParam
	//case C.DIERR_MAPFILEFAIL:
	//	return ErrMapFileFail
	//case C.DIERR_MOREDATA:
	//return ErrMoreData
	case C.DIERR_NOAGGREGATION:
		return ErrNoAggregation
	case C.DIERR_NOINTERFACE:
		return ErrNoInterface
	case C.DIERR_NOTACQUIRED:
		return ErrNotAcquired
	//case C.DIERR_NOTBUFFERED:
	//return ErrNotBuffered
	//case C.DIERR_NOTDOWNLOADED:
	//return ErrNotDownloaded
	//case C.DIERR_NOTEXCLUSIVEACQUIRED:
	//return ErrNotExclusiveAcquired
	case C.DIERR_NOTFOUND:
		return ErrNotFound
	case C.DIERR_NOTINITIALIZED:
		return ErrNotInitialized
	case C.DIERR_OLDDIRECTINPUTVERSION:
		return ErrOldDirectInputVersion
		// TODO this is the same numerical value as for ErrHandleExists
	//case C.DIERR_OTHERAPPHASPRIO:
	//return ErrOtherAppHasPriority
	case C.DIERR_OUTOFMEMORY:
		return ErrOutOfMemory
		// TODO this is the same numerical value as for ErrHandleExists
	//case C.DIERR_READONLY:
	//return ErrReadOnly
	//case C.DIERR_REPORTFULL:
	//return ErrReportFull
	//case C.DIERR_UNPLUGGED:
	//return ErrUnplugged
	case C.DIERR_UNSUPPORTED:
		return ErrUnsupported
	case C.E_HANDLE:
		return ErrHandle
	case C.E_PENDING:
		return ErrPending
	case C.E_POINTER:
		return ErrPointer

	//case C.DI_BUFFEROVERFLOW:
	//return ErrBufferOverflow
	case C.DI_DOWNLOADSKIPPED:
		return ErrDownloadSkipped
	case C.DI_EFFECTRESTARTED:
		return ErrEffectRestarted
	case C.DI_NOTATTACHED:
		return ErrNotAttached

	default:
		return Error("Unknown error code: " + strconv.Itoa(int(result)))
	}
}
*/
