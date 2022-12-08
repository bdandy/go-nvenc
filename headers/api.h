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