package types

// #cgo CFLAGS: -I ../../include
// #include "types.h"
import "C"
import (
	"unsafe"
)

type deviceType int

var (
	DeviceTypeDirectx = deviceType(C.NV_ENC_DEVICE_TYPE_DIRECTX)
	DeviceTypeCuda    = deviceType(C.NV_ENC_DEVICE_TYPE_CUDA)

	InfiniteGOPLength = C.NVENC_INFINITE_GOPLENGTH
	LevelAutoSelect   = C.NV_ENC_LEVEL_AUTOSELECT
)

type OpenEncodeSessionParams C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS

func NewOpenEncodeSessionParams(devType deviceType, device unsafe.Pointer) *OpenEncodeSessionParams {
	params := new(OpenEncodeSessionParams)
	params.version = C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS_VER
	params.deviceType = C.NV_ENC_DEVICE_TYPE(devType)
	params.device = device
	params.apiVersion = C.NVENCAPI_VERSION

	return params
}

type CapsParam C.NV_ENC_CAPS_PARAM

func NewCapsParam(caps uint) *CapsParam {
	param := new(CapsParam)
	param.version = C.NV_ENC_CAPS_PARAM_VER
	param.capsToQuery = C.NV_ENC_CAPS(caps)
	return param
}

type ENC_STAT C.NV_ENC_STAT
type SEQUENCE_PARAM_PAYLOAD C.NV_ENC_SEQUENCE_PARAM_PAYLOAD

func (p *SEQUENCE_PARAM_PAYLOAD) Bytes() []byte {
	return C.GoBytes(p.spsppsBuffer, C.int(*p.outSPSPPSPayloadSize))
}

func NewSequenceParamPayload() *SEQUENCE_PARAM_PAYLOAD {
	var payload = new(SEQUENCE_PARAM_PAYLOAD)
	bufsize := 128
	payload.version = C.NV_ENC_SEQUENCE_PARAM_PAYLOAD_VER
	payload.inBufferSize = C.uint32_t(bufsize)
	payload.spsppsBuffer = C.malloc(C.size_t(bufsize))
	payload.outSPSPPSPayloadSize = (*C.uint32_t)(C.malloc(C.size_t(4)))
	return payload
}

type EVENT_PARAMS C.NV_ENC_EVENT_PARAMS
type MAP_INPUT_RESOURCE_PARAMS C.NV_ENC_MAP_INPUT_RESOURCE
type REGISTER_RESOURCE_PARAMS C.NV_ENC_REGISTER_RESOURCE
type RECONFIGURE_PARAMS C.NV_ENC_RECONFIGURE_PARAMS

func (r *RECONFIGURE_PARAMS) SetInitializeParams(p InitializeParams) {
	r.version = C.NV_ENC_RECONFIGURE_PARAMS_VER
	r.reInitEncodeParams = C.NV_ENC_INITIALIZE_PARAMS(p)
	C.SetResetEncoder((*C.NV_ENC_RECONFIGURE_PARAMS)(r))
}

type CREATE_MV_BUFFER_PARAMS C.NV_ENC_CREATE_MV_BUFFER
type MOTION_ESTIMATE_ONLY_PARAMS C.NV_ENC_MEONLY_PARAMS
type FRAME_FIELD_MODE C.NV_ENC_PARAMS_FRAME_FIELD_MODE
type MV_PRECISION C.NV_ENC_MV_PRECISION

type CODEC_CONFIG C.NV_ENC_CODEC_CONFIG

func (c CODEC_CONFIG) H264Config() *CONFIG_H264 {
	return C.GetH264Config((*C.NV_ENC_CODEC_CONFIG)(&c))
}

func (c CODEC_CONFIG) HEVCConfig() *CONFIG_HEVC {
	return C.GetHEVCConfig((*C.NV_ENC_CODEC_CONFIG)(&c))
}
