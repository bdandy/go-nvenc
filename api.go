package nvenc

/*
 #include "headers/nvEncodeAPI.h"

 static inline int callOpenEncodeSession(PNVENCOPENENCODESESSION f,void* device, uint32_t deviceType, void** encoder) {
 	return f(device, deviceType, encoder);
 };
 static inline int callGetGUIDCount(PNVENCGETENCODEGUIDCOUNT f,void* encoder, uint32_t* encodeGUIDCount) {
 	return f(encoder, encodeGUIDCount);
 };
 static inline int callGetProfileGUIDCount(PNVENCGETENCODEPROFILEGUIDCOUNT f,void* encoder, GUID encodeGUID, uint32_t* encodeProfileGUIDCount) {
 	return f(encoder, encodeGUID, encodeProfileGUIDCount);
 };
 static inline int callGetProfileGUIDs(PNVENCGETENCODEPROFILEGUIDS f, void* encoder, GUID encodeGUID, GUID* profileGUIDs, uint32_t guidArraySize, uint32_t* GUIDCount) {
 	return f(encoder, encodeGUID, profileGUIDs, guidArraySize, GUIDCount);
 };
 static inline int callGetGUIDs(PNVENCGETENCODEGUIDS f, void* encoder, GUID* GUIDs, uint32_t guidArraySize, uint32_t* GUIDCount) {
 	return f(encoder, GUIDs, guidArraySize, GUIDCount);
 };
 static inline int callGetInputFormatCount(PNVENCGETINPUTFORMATCOUNT f, void* encoder, GUID encodeGUID, uint32_t* inputFmtCount) {
 	return f(encoder, encodeGUID, inputFmtCount);
 };
 static inline int callGetInputFormats(PNVENCGETINPUTFORMATS f, void* encoder, GUID encodeGUID, NV_ENC_BUFFER_FORMAT* inputFmts, uint32_t inputFmtArraySize, uint32_t* inputFmtCount) {
 	return f(encoder, encodeGUID, inputFmts, inputFmtArraySize, inputFmtCount);
 };
 static inline int callGetCaps(PNVENCGETENCODECAPS f,void* encoder, GUID encodeGUID, NV_ENC_CAPS_PARAM* capsParam, int* capsVal) {
 	return f(encoder, encodeGUID, capsParam, capsVal);
 };
 static inline int callGetPresetCount(PNVENCGETENCODEPRESETCOUNT f,void* encoder, GUID encodeGUID, uint32_t* encodePresetGUIDCount) {
 	return f(encoder, encodeGUID, encodePresetGUIDCount);
 };
 static inline int callGetPresetGUIDs(PNVENCGETENCODEPRESETGUIDS f, void* encoder, GUID encodeGUID, GUID* presetGUIDs, uint32_t guidArraySize, uint32_t* encodePresetGUIDCount) {
 	return f(encoder, encodeGUID, presetGUIDs, guidArraySize, encodePresetGUIDCount);
 };
 static inline int callGetPresetConfig(PNVENCGETENCODEPRESETCONFIG f,void* encoder, GUID encodeGUID, GUID  presetGUID, NV_ENC_PRESET_CONFIG* presetConfig) {
 	return f(encoder, encodeGUID, presetGUID, presetConfig);
 };
 static inline int callInitializeEncoder(PNVENCINITIALIZEENCODER f, void* encoder, NV_ENC_INITIALIZE_PARAMS* createEncodeParams) {
 	return f(encoder, createEncodeParams);
 };
 static inline int callCreateBuffer(PNVENCCREATEINPUTBUFFER f,void* encoder, NV_ENC_CREATE_INPUT_BUFFER* createInputBufferParams) {
 	return f(encoder, createInputBufferParams);
 };
 static inline int callDestroyBuffer(PNVENCDESTROYINPUTBUFFER f,void* encoder, NV_ENC_INPUT_PTR inputBuffer) {
 	return f(encoder, inputBuffer);
 };
 static inline int callCreateBitstreamBuffer(PNVENCCREATEBITSTREAMBUFFER f,void* encoder, NV_ENC_CREATE_BITSTREAM_BUFFER* createBitstreamBufferParams) {
 	return f(encoder, createBitstreamBufferParams);
 };
 static inline int callDestroyBitstreamBuffer(PNVENCDESTROYBITSTREAMBUFFER f, void* encoder, NV_ENC_OUTPUT_PTR bitstreamBuffer) {
 	return f(encoder, bitstreamBuffer);
 };
 static inline int callEncodePicture(PNVENCENCODEPICTURE f, void* encoder, NV_ENC_PIC_PARAMS* encodePicParams) {
 	return f(encoder, encodePicParams);
 };
 static inline int callLockBitstream(PNVENCLOCKBITSTREAM f,void* encoder, NV_ENC_LOCK_BITSTREAM* lockBitstreamBufferParams) {
 	return f(encoder, lockBitstreamBufferParams);
 };
 static inline int callUnlockBitstream(PNVENCUNLOCKBITSTREAM f, void* encoder, NV_ENC_OUTPUT_PTR bitstreamBuffer) {
 	return f(encoder, bitstreamBuffer);
 };
 static inline int callLockInputBuffer(PNVENCLOCKINPUTBUFFER f, void* encoder, NV_ENC_LOCK_INPUT_BUFFER* lockInputBufferParams) {
 	return f(encoder, lockInputBufferParams);
 };
 static inline int callUnlockInputBuffer(PNVENCUNLOCKINPUTBUFFER f, void* encoder, NV_ENC_INPUT_PTR inputBuffer) {
 	return f(encoder, inputBuffer);
 };
 static inline int callEncodeStats(PNVENCGETENCODESTATS f,void* encoder, NV_ENC_STAT* encodeStats) {
 	return f(encoder, encodeStats);
 }
 static inline int callGetSequenceParams(PNVENCGETSEQUENCEPARAMS f, void* encoder, NV_ENC_SEQUENCE_PARAM_PAYLOAD* sequenceParamPayload) {
 	return f(encoder, sequenceParamPayload);
 };
 static inline int callRegisterAsyncEvent(PNVENCREGISTERASYNCEVENT f,void* encoder, NV_ENC_EVENT_PARAMS* eventParams) {
 	return f(encoder, eventParams);
 };
 static inline int callUnregisterAsyncEvent(PNVENCUNREGISTERASYNCEVENT f, void* encoder, NV_ENC_EVENT_PARAMS* eventParams) {
 	return f(encoder, eventParams);
 };
 static inline int callMapInputResource(PNVENCMAPINPUTRESOURCE f, void* encoder, NV_ENC_MAP_INPUT_RESOURCE* mapInputResParams) {
 	return f(encoder, mapInputResParams);
 };
 static inline int callUnmapInputResource(PNVENCUNMAPINPUTRESOURCE f,void* encoder, NV_ENC_INPUT_PTR mappedInputBuffer) {
 	return f(encoder, mappedInputBuffer);
 };
 static inline int callDestroyEncoder(PNVENCDESTROYENCODER f, void* encoder) {
 	return f(encoder);
 };
 static inline int callInvalidateRefFrames(PNVENCINVALIDATEREFFRAMES f, void* encoder, uint64_t invalidRefFrameTimeStamp) {
 	return f(encoder, invalidRefFrameTimeStamp);
 };
 static inline int callOpenEncodeSessionEx(PNVENCOPENENCODESESSIONEX f, NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS *openSessionExParams, void** encoder) {
 	return f(openSessionExParams, encoder);
 };
 static inline int callRegisterResource(PNVENCREGISTERRESOURCE f, void* encoder, NV_ENC_REGISTER_RESOURCE* registerResParams) {
 	return f(encoder, registerResParams);
 };
 static inline int callUnregisterResource(PNVENCUNREGISTERRESOURCE f, void* encoder, NV_ENC_REGISTERED_PTR registeredRes) {
 	return f(encoder, registeredRes);
 };
 static inline int callReconfigureEncoder(PNVENCRECONFIGUREENCODER f, void* encoder, NV_ENC_RECONFIGURE_PARAMS* reInitEncodeParams) {
 	return f(encoder, reInitEncodeParams);
 };
 static inline int callCreateMvBuffer(PNVENCCREATEMVBUFFER f,void* encoder, NV_ENC_CREATE_MV_BUFFER* createMVBufferParams) {
 	return f(encoder, createMVBufferParams);
 };
 static inline int callDestroyMvBuffer(PNVENCDESTROYMVBUFFER f,  void* encoder, NV_ENC_OUTPUT_PTR MVBuffer) {
 	return f(encoder, MVBuffer);
 };
 static inline int callRunMotionEstimateOnly(PNVENCRUNMOTIONESTIMATIONONLY f, void* encoder, NV_ENC_MEONLY_PARAMS* MEOnlyParams) {
 	return f(encoder, MEOnlyParams);
 };
*/
import "C"
import "unsafe"

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
	err := codeToError(C.callGetProfileGUIDCount(e.getProfileGUIDCountPtr, encoder, C.GUID(encodeGUID), &count))
	return uint32(count), err
}

func (e *EncoderFunctions) getProfileGUIDs(encoder unsafe.Pointer, encodeGUID GUID, guidArraySize uint32) ([]GUID, error) {
	var count C.uint32_t
	var guids = make([]GUID, guidArraySize)

	err := codeToError(C.callGetProfileGUIDs(e.getProfileGUIDsPtr, encoder, C.GUID(encodeGUID), (*C.GUID)(unsafe.Pointer(&guids[0])), C.uint32_t(guidArraySize), &count))

	return guids, err
}

func (e *EncoderFunctions) getGUIDs(encoder unsafe.Pointer, guidArraySize uint32) ([]GUID, error) {
	var count C.uint32_t
	var guids = make([]GUID, guidArraySize)

	pGuids := (*C.GUID)(unsafe.Pointer(&guids[0]))

	err := codeToError(C.callGetGUIDs(e.getGUIDsPtr, encoder, pGuids, C.uint32_t(guidArraySize), &count))

	return guids, err
}

func (e *EncoderFunctions) getInputFormatCount(encoder unsafe.Pointer, encodeGUID GUID) (uint32, error) {
	var count C.uint32_t

	err := codeToError(C.callGetInputFormatCount(e.getInputFormatCountPtr, encoder, C.GUID(encodeGUID), &count))

	return uint32(count), err
}
func (e *EncoderFunctions) getInputFormats(encoder unsafe.Pointer, encodeGUID GUID, size uint32) ([]BUFFER_FORMAT, error) {
	var count C.uint32_t
	var fmts = make([]BUFFER_FORMAT, size)

	pFmts := (*C.NV_ENC_BUFFER_FORMAT)(unsafe.Pointer(&fmts[0]))

	err := codeToError(C.callGetInputFormats(e.getInputFormatsPtr, encoder, C.GUID(encodeGUID), pFmts, C.uint32_t(size), &count))

	return fmts, err
}

func (e *EncoderFunctions) getCaps(encoder unsafe.Pointer, encodeGUID GUID, capsParam *CAPS_PARAM) (int, error) {
	var capsVal C.int

	err := codeToError(C.callGetCaps(e.getCapsPtr, encoder, C.GUID(encodeGUID), (*C.NV_ENC_CAPS_PARAM)(capsParam), &capsVal))

	return int(capsVal), err
}

func (e *EncoderFunctions) getPresetCount(encoder unsafe.Pointer, encodeGUID GUID) (uint32, error) {
	var count C.uint32_t

	err := codeToError(C.callGetPresetCount(e.getPresetCountPtr, encoder, C.GUID(encodeGUID), &count))

	return uint32(count), err
}

func (e *EncoderFunctions) getPresetGUIDs(encoder unsafe.Pointer, encodeGUID GUID, guidArraySize uint32) ([]GUID, error) {
	var GUIDs = make([]GUID, guidArraySize)
	var count C.uint32_t

	pGUIDs := (*C.GUID)(unsafe.Pointer(&GUIDs[0]))
	err := codeToError(C.callGetPresetGUIDs(e.getPresetGUIDsPtr, encoder, C.GUID(encodeGUID), pGUIDs, C.uint32_t(guidArraySize), &count))
	return GUIDs, err
}

func (e *EncoderFunctions) getPresetConfig(encoder unsafe.Pointer, encodeGUID GUID, presetGUID GUID) (*ENCODER_CONFIG, error) {
	var config C.NV_ENC_PRESET_CONFIG
	config.version = C.NV_ENC_PRESET_CONFIG_VER
	config.presetCfg.version = C.NV_ENC_CONFIG_VER

	err := codeToError(C.callGetPresetConfig(e.getPresetConfigPtr, encoder, C.GUID(encodeGUID), C.GUID(presetGUID), &config))
	return (*ENCODER_CONFIG)(&config.presetCfg), err
}

func (e *EncoderFunctions) initializeEncoder(encoder unsafe.Pointer, createEncodeParams *INITIALIZE_PARAMS) error {
	err := codeToError(C.callInitializeEncoder(e.initializeEncoderPtr, encoder, (*C.NV_ENC_INITIALIZE_PARAMS)(createEncodeParams)))
	return err
}

func (e *EncoderFunctions) createBuffer(encoder unsafe.Pointer, params *INPUT_BUFFER) error {
	err := codeToError(C.callCreateBuffer(e.createBufferPtr, encoder, (*C.NV_ENC_CREATE_INPUT_BUFFER)(params)))
	return err
}
func (e *EncoderFunctions) destroyBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyBuffer(e.destroyBufferPtr, encoder, (C.NV_ENC_INPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) createBitstreamBuffer(encoder unsafe.Pointer, params *BITSTREAM_BUFFER) error {
	err := codeToError(C.callCreateBitstreamBuffer(e.createBitstreamBufferPtr, encoder, (*C.NV_ENC_CREATE_BITSTREAM_BUFFER)(params)))
	return err
}

func (e *EncoderFunctions) destroyBitstreamBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyBitstreamBuffer(e.destroyBitstreamBufferPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) encodePicture(encoder unsafe.Pointer, encodePicParams *ENC_PIC_PARAMS) error {
	err := codeToError(C.callEncodePicture(e.encodePicturePtr, encoder, (*C.NV_ENC_PIC_PARAMS)(encodePicParams)))
	return err
}

func (e *EncoderFunctions) lockBitstream(encoder unsafe.Pointer, params *LOCK_BITSTREAM_PARAMS) error {
	err := codeToError(C.callLockBitstream(e.lockBitstreamPtr, encoder, (*C.NV_ENC_LOCK_BITSTREAM)(params)))
	return err
}

func (e *EncoderFunctions) unlockBitstream(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callUnlockBitstream(e.unlockBitstreamPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) lockInputBuffer(encoder unsafe.Pointer, params *LOCK_INPUT_BUFFER_PARAMS) error {
	err := codeToError(C.callLockInputBuffer(e.lockInputBufferPtr, encoder, (*C.NV_ENC_LOCK_INPUT_BUFFER)(params)))
	return err
}

func (e *EncoderFunctions) unlockInputBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callUnlockInputBuffer(e.unlockInputBufferPtr, encoder, (C.NV_ENC_INPUT_PTR)(buffer)))
	return err
}

func (e *EncoderFunctions) encodeStats(encoder unsafe.Pointer, encodeStats *ENC_STAT) error {
	err := codeToError(C.callEncodeStats(e.encodeStatsPtr, encoder, (*C.NV_ENC_STAT)(encodeStats)))
	return err
}

func (e *EncoderFunctions) getSequenceParams(encoder unsafe.Pointer, payload *SEQUENCE_PARAM_PAYLOAD) error {
	err := codeToError(C.callGetSequenceParams(e.getSequenceParamsPtr, encoder, (*C.NV_ENC_SEQUENCE_PARAM_PAYLOAD)(payload)))
	return err
}

func (e *EncoderFunctions) registerAsyncEvent(encoder unsafe.Pointer, eventParams *EVENT_PARAMS) error {
	err := codeToError(C.callRegisterAsyncEvent(e.registerAsyncEventPtr, encoder, (*C.NV_ENC_EVENT_PARAMS)(eventParams)))
	return err
}

func (e *EncoderFunctions) unregisterAsyncEvent(encoder unsafe.Pointer, eventParams *EVENT_PARAMS) error {
	err := codeToError(C.callUnregisterAsyncEvent(e.unregisterAsyncEventPtr, encoder, (*C.NV_ENC_EVENT_PARAMS)(eventParams)))
	return err
}

func (e *EncoderFunctions) mapInputResource(encoder unsafe.Pointer, params *MAP_INPUT_RESOURCE_PARAMS) error {
	err := codeToError(C.callMapInputResource(e.mapInputResourcePtr, encoder, (*C.NV_ENC_MAP_INPUT_RESOURCE)(params)))
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

func (e *EncoderFunctions) openEncodeSessionEx(params *OPEN_ENCODE_SESSION_PARAMS) (unsafe.Pointer, error) {
	var encoder unsafe.Pointer
	err := codeToError(C.callOpenEncodeSessionEx(e.openEncodeSessionExPtr, (*C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS)(params), &encoder))
	return encoder, err
}

func (e *EncoderFunctions) registerResource(encoder unsafe.Pointer, registerResParams *REGISTER_RESOURCE_PARAMS) error {
	err := codeToError(C.callRegisterResource(e.registerResourcePtr, encoder, (*C.NV_ENC_REGISTER_RESOURCE)(registerResParams)))
	return err
}

func (e *EncoderFunctions) unregisterResource(encoder unsafe.Pointer, resource unsafe.Pointer) error {
	err := codeToError(C.callUnregisterResource(e.unregisterResourcePtr, encoder, (C.NV_ENC_REGISTERED_PTR)(resource)))
	return err
}

func (e *EncoderFunctions) reconfigureEncoder(encoder unsafe.Pointer, params *RECONFIGURE_PARAMS) error {
	err := codeToError(C.callReconfigureEncoder(e.reconfigureEncoderPtr, encoder, (*C.NV_ENC_RECONFIGURE_PARAMS)(params)))
	return err
}

func (e *EncoderFunctions) createMvBuffer(encoder unsafe.Pointer, createMVBufferParams *CREATE_MV_BUFFER_PARAMS) error {
	err := codeToError(C.callCreateMvBuffer(e.createMvBufferPtr, encoder, (*C.NV_ENC_CREATE_MV_BUFFER)(createMVBufferParams)))
	return err
}

func (e *EncoderFunctions) destroyMvBuffer(encoder unsafe.Pointer, buffer unsafe.Pointer) error {
	err := codeToError(C.callDestroyMvBuffer(e.destroyMvBufferPtr, encoder, (C.NV_ENC_OUTPUT_PTR)(buffer)))
	return err
}
func (e *EncoderFunctions) runMotionEstimateOnly(encoder unsafe.Pointer, meOnlyParams *MOTION_ESTIMATE_ONLY_PARAMS) error {
	err := codeToError(C.callRunMotionEstimateOnly(e.runMotionEstimateOnlyPtr, encoder, (*C.NV_ENC_MEONLY_PARAMS)(meOnlyParams)))
	return err
}

func newEncoderFunctions() *EncoderFunctions {
	e := new(EncoderFunctions)
	e.version = C.NV_ENCODE_API_FUNCTION_LIST_VER
	return e
}
