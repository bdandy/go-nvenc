package nvenc

// #cgo CFLAGS: -I ./include
// #include <functions.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/bdandy/go-nvenc/v8/guid"
	"github.com/bdandy/go-nvenc/v8/internal/types"
)

type GUID = guid.GUID

// EncoderFunctions is a bridge to C method calls
// Contains pointers to
type EncoderFunctions struct {
	version uint32
	_       uint32

	openEncodeSessionPtr      *[0]byte
	getGUIDCountPtr           *[0]byte
	getProfileGUIDCountPtr    *[0]byte
	getProfileGUIDsPtr        *[0]byte
	getGUIDsPtr               *[0]byte
	getInputFormatCountPtr    *[0]byte
	getInputFormatsPtr        *[0]byte
	getCapsPtr                *[0]byte
	getPresetCountPtr         *[0]byte
	getPresetGUIDsPtr         *[0]byte
	getPresetConfigPtr        *[0]byte
	initializeEncoderPtr      *[0]byte
	createBufferPtr           *[0]byte
	destroyBufferPtr          *[0]byte
	createBitstreamBufferPtr  *[0]byte
	destroyBitstreamBufferPtr *[0]byte
	encodePicturePtr          *[0]byte
	lockBitstreamPtr          *[0]byte
	unlockBitstreamPtr        *[0]byte
	lockInputBufferPtr        *[0]byte
	unlockInputBufferPtr      *[0]byte
	encodeStatsPtr            *[0]byte
	getSequenceParamsPtr      *[0]byte
	registerAsyncEventPtr     *[0]byte
	unregisterAsyncEventPtr   *[0]byte
	mapInputResourcePtr       *[0]byte
	unmapInputResourcePtr     *[0]byte
	destroyEncoderPtr         *[0]byte
	invalidateRefFramesPtr    *[0]byte
	openEncodeSessionExPtr    *[0]byte
	registerResourcePtr       *[0]byte
	unregisterResourcePtr     *[0]byte
	reconfigureEncoderPtr     *[0]byte

	_ uintptr

	createMvBufferPtr        *[0]byte
	destroyMvBufferPtr       *[0]byte
	runMotionEstimateOnlyPtr *[0]byte

	_ [281]uintptr
}

func (e *EncoderFunctions) openEncodeSession(device unsafe.Pointer, deviceType uint32) (unsafe.Pointer, error) {
	var encoder unsafe.Pointer
	err := codeToError(C.callOpenEncodeSession(e.openEncodeSessionPtr, device, C.uint32_t(deviceType), &encoder))
	return encoder, err
}

func (e *EncoderFunctions) getGUIDCount(encoder unsafe.Pointer) (uint32, error) {
	var count C.uint32_t
	err := codeToError(C.callGetGUIDCount(e.getGUIDCountPtr, encoder, &count))
	return uint32(count), err
}

func (e *EncoderFunctions) getProfileGUIDCount(encoder unsafe.Pointer, encodeGUID GUID) (uint32, error) {
	var count C.uint32_t
	err := codeToError(C.callGetProfileGUIDCount(e.getProfileGUIDCountPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), &count))
	return uint32(count), err
}

func (e *EncoderFunctions) getProfileGUIDs(encoder unsafe.Pointer, encodeGUID GUID, guidArraySize uint32) ([]GUID, error) {
	var count C.uint32_t
	var guids = make([]GUID, guidArraySize)

	err := codeToError(C.callGetProfileGUIDs(e.getProfileGUIDsPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), (*C.GUID)(unsafe.Pointer(&guids[0])), C.uint32_t(guidArraySize), &count))

	return guids, err
}

// getGUIDs returns slice of supported guids for encoder
func (e *EncoderFunctions) getGUIDs(encoder unsafe.Pointer, guidArraySize uint32) ([]GUID, error) {
	var count C.uint32_t
	var guids = make([]GUID, guidArraySize)

	pGuids := (*C.GUID)(unsafe.Pointer(&guids[0]))

	err := codeToError(C.callGetGUIDs(e.getGUIDsPtr, encoder, pGuids, C.uint32_t(guidArraySize), &count))

	return guids, err
}

func (e *EncoderFunctions) getInputFormats(encoder unsafe.Pointer, encodeGUID guid.CodecGUID) ([]types.BufferFormat, error) {
	var count C.uint32_t

	err := codeToError(C.callGetInputFormatCount(e.getInputFormatCountPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), &count))
	if err != nil {
		return nil, fmt.Errorf("get input format count: %w", err)
	}

	data := make([]types.BufferFormat, count)
	err = codeToError(C.callGetInputFormats(e.getInputFormatsPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), (*C.NV_ENC_BUFFER_FORMAT)(unsafe.Pointer(&data[0])), C.uint32_t(count), &count))
	if err != nil {
		return nil, fmt.Errorf("get input formats: %w", err)
	}

	return data, err
}

func (e *EncoderFunctions) getCaps(encoder unsafe.Pointer, encodeGUID GUID, capsParam *types.CapsParam) (int, error) {
	var capsVal C.int

	err := codeToError(C.callGetCaps(e.getCapsPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), (*C.NV_ENC_CAPS_PARAM)(unsafe.Pointer(capsParam)), &capsVal))

	return int(capsVal), err
}

func (e *EncoderFunctions) getPresetGUIDs(encoder unsafe.Pointer, codecGUID guid.CodecGUID) ([]guid.PresetGUID, error) {
	var count C.uint32_t
	err := codeToError(C.callGetPresetCount(e.getPresetCountPtr, encoder, *(*C.GUID)(unsafe.Pointer(&codecGUID)), &count))
	if err != nil {
		return nil, fmt.Errorf("get preset count: %w", err)
	}

	data := make(guid.GUIDs, count)
	err = codeToError(C.callGetPresetGUIDs(e.getPresetGUIDsPtr, encoder, *(*C.GUID)(unsafe.Pointer(&codecGUID)), (*C.GUID)(unsafe.Pointer(&data[0])), C.uint32_t(count), &count))
	if err != nil {
		return nil, fmt.Errorf("get preset GUIDs: %w", err)
	}

	return data.ToPreset(), err
}

func (e *EncoderFunctions) getPresetConfig(encoder unsafe.Pointer, encodeGUID guid.CodecGUID, presetGUID guid.PresetGUID) (*types.EncoderConfig, error) {
	var config C.NV_ENC_PRESET_CONFIG
	config.version = C.NV_ENC_PRESET_CONFIG_VER
	config.presetCfg.version = C.NV_ENC_CONFIG_VER

	err := codeToError(C.callGetPresetConfig(e.getPresetConfigPtr, encoder, *(*C.GUID)(unsafe.Pointer(&encodeGUID)), *(*C.GUID)(unsafe.Pointer(&presetGUID)), &config))

	return (*types.EncoderConfig)(unsafe.Pointer(&config.presetCfg)), err
}

func (e *EncoderFunctions) initializeEncoder(encoder unsafe.Pointer, createEncodeParams *types.InitializeParams) error {
	err := codeToError(C.callInitializeEncoder(e.initializeEncoderPtr, encoder, (*C.NV_ENC_INITIALIZE_PARAMS)(unsafe.Pointer(createEncodeParams))))
	return err
}

func (e *EncoderFunctions) createBuffer(encoder unsafe.Pointer, params *types.CreateInputBuffer) error {
	err := codeToError(C.callCreateBuffer(e.createBufferPtr, encoder, (*C.NV_ENC_CREATE_INPUT_BUFFER)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) destroyBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyBuffer(e.destroyBufferPtr, encoder, (C.NV_ENC_INPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) createBitstreamBuffer(encoder unsafe.Pointer, params *types.BitstreamBuffer) error {
	err := codeToError(C.callCreateBitstreamBuffer(e.createBitstreamBufferPtr, encoder, (*C.NV_ENC_CREATE_BITSTREAM_BUFFER)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) destroyBitstreamBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyBitstreamBuffer(e.destroyBitstreamBufferPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) encodePicture(encoder unsafe.Pointer, encodePicParams *types.EncoderPictureParams) error {
	err := codeToError(C.callEncodePicture(e.encodePicturePtr, encoder, (*C.NV_ENC_PIC_PARAMS)(unsafe.Pointer(encodePicParams))))
	return err
}

func (e *EncoderFunctions) lockBitstream(encoder unsafe.Pointer, params *types.LockBitstreamParams) error {
	err := codeToError(C.callLockBitstream(e.lockBitstreamPtr, encoder, (*C.NV_ENC_LOCK_BITSTREAM)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) unlockBitstream(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callUnlockBitstream(e.unlockBitstreamPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) lockInputBuffer(encoder unsafe.Pointer, params *types.LockInputBufferParams) error {
	err := codeToError(C.callLockInputBuffer(e.lockInputBufferPtr, encoder, (*C.NV_ENC_LOCK_INPUT_BUFFER)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) unlockInputBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callUnlockInputBuffer(e.unlockInputBufferPtr, encoder, (C.NV_ENC_INPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) encodeStats(encoder unsafe.Pointer, encodeStats *types.ENC_STAT) error {
	err := codeToError(C.callEncodeStats(e.encodeStatsPtr, encoder, (*C.NV_ENC_STAT)(unsafe.Pointer(encodeStats))))
	return err
}

func (e *EncoderFunctions) getSequenceParams(encoder unsafe.Pointer, payload *types.SEQUENCE_PARAM_PAYLOAD) error {
	err := codeToError(C.callGetSequenceParams(e.getSequenceParamsPtr, encoder, (*C.NV_ENC_SEQUENCE_PARAM_PAYLOAD)(unsafe.Pointer(payload))))
	return err
}

func (e *EncoderFunctions) registerAsyncEvent(encoder unsafe.Pointer, eventParams *types.EVENT_PARAMS) error {
	err := codeToError(C.callRegisterAsyncEvent(e.registerAsyncEventPtr, encoder, (*C.NV_ENC_EVENT_PARAMS)(unsafe.Pointer(eventParams))))
	return err
}

func (e *EncoderFunctions) unregisterAsyncEvent(encoder unsafe.Pointer, eventParams *types.EVENT_PARAMS) error {
	err := codeToError(C.callUnregisterAsyncEvent(e.unregisterAsyncEventPtr, encoder, (*C.NV_ENC_EVENT_PARAMS)(unsafe.Pointer(eventParams))))
	return err
}

func (e *EncoderFunctions) mapInputResource(encoder unsafe.Pointer, params *types.MAP_INPUT_RESOURCE_PARAMS) error {
	err := codeToError(C.callMapInputResource(e.mapInputResourcePtr, encoder, (*C.NV_ENC_MAP_INPUT_RESOURCE)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) unmapInputResource(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callUnmapInputResource(e.unmapInputResourcePtr, encoder, (C.NV_ENC_INPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) destroyEncoder(encoder unsafe.Pointer) error {
	err := codeToError(C.callDestroyEncoder(e.destroyEncoderPtr, encoder))
	return err
}

func (e *EncoderFunctions) invalidateRefFrames(encoder unsafe.Pointer, invalidRefFrameTimeStamp uint64) error {
	err := codeToError(C.callInvalidateRefFrames(e.invalidateRefFramesPtr, encoder, C.uint64_t(invalidRefFrameTimeStamp)))
	return err
}

func (e *EncoderFunctions) openEncodeSessionEx(params *types.OpenEncodeSessionParams) (unsafe.Pointer, error) {
	var encoder unsafe.Pointer
	err := codeToError(C.callOpenEncodeSessionEx(e.openEncodeSessionExPtr, (*C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS)(unsafe.Pointer(params)), &encoder))
	return encoder, err
}

func (e *EncoderFunctions) registerResource(encoder unsafe.Pointer, registerResParams *types.REGISTER_RESOURCE_PARAMS) error {
	err := codeToError(C.callRegisterResource(e.registerResourcePtr, encoder, (*C.NV_ENC_REGISTER_RESOURCE)(unsafe.Pointer(registerResParams))))
	return err
}

func (e *EncoderFunctions) unregisterResource(encoder unsafe.Pointer, resource unsafe.Pointer) error {
	err := codeToError(C.callUnregisterResource(e.unregisterResourcePtr, encoder, (C.NV_ENC_REGISTERED_PTR)(resource)))
	return err
}

func (e *EncoderFunctions) reconfigureEncoder(encoder unsafe.Pointer, params *types.RECONFIGURE_PARAMS) error {
	err := codeToError(C.callReconfigureEncoder(e.reconfigureEncoderPtr, encoder, (*C.NV_ENC_RECONFIGURE_PARAMS)(unsafe.Pointer(params))))
	return err
}

func (e *EncoderFunctions) createMvBuffer(encoder unsafe.Pointer, createMVBufferParams *types.CREATE_MV_BUFFER_PARAMS) error {
	err := codeToError(C.callCreateMvBuffer(e.createMvBufferPtr, encoder, (*C.NV_ENC_CREATE_MV_BUFFER)(unsafe.Pointer(createMVBufferParams))))
	return err
}

func (e *EncoderFunctions) destroyMvBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyMvBuffer(e.destroyMvBufferPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}
func (e *EncoderFunctions) runMotionEstimateOnly(encoder unsafe.Pointer, meOnlyParams *types.MOTION_ESTIMATE_ONLY_PARAMS) error {
	err := codeToError(C.callRunMotionEstimateOnly(e.runMotionEstimateOnlyPtr, encoder, (*C.NV_ENC_MEONLY_PARAMS)(unsafe.Pointer(meOnlyParams))))
	return err
}

func newEncoderFunctions() *EncoderFunctions {
	e := new(EncoderFunctions)
	e.version = C.NV_ENCODE_API_FUNCTION_LIST_VER
	return e
}
