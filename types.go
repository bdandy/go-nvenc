package nvenc

// #include "headers/types.h"
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

// OpenEncodeSessionParams is a C type from library
/*
uint32_t            version;                          //< [in]: Struct version. Must be set to ::NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS_VER.
NV_ENC_DEVICE_TYPE  deviceType;                       //< [in]: Specified the device Type
void*               device;                           //< [in]: Pointer to client device.
void*               reserved;                         //< [in]: Reserved and must be set to 0.
uint32_t            apiVersion;                       //< [in]: API version. Should be set to NVENCAPI_VERSION.
uint32_t            reserved1[253];                   //< [in]: Reserved and must be set to 0
void*               reserved2[64];                    //< [in]: Reserved and must be set to NULL
*/
type OpenEncodeSessionParams C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS

func newOpenEncodeSessionParams(devType deviceType, device unsafe.Pointer) *OpenEncodeSessionParams {
	params := new(OpenEncodeSessionParams)
	params.version = C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS_VER
	params.deviceType = C.NV_ENC_DEVICE_TYPE(devType)
	params.device = device
	params.apiVersion = C.NVENCAPI_VERSION

	return params
}

/*
uint32_t version;                                  /**< [in]: Struct version. Must be set to ::NV_ENC_CAPS_PARAM_VER
NV_ENC_CAPS  capsToQuery;                          /**< [in]: Specifies the encode capability to be queried. Client should pass a member for ::NV_ENC_CAPS enum.
uint32_t reserved[62];                             /**< [in]: Reserved and must be set to 0
*/
type CapsParam = C.NV_ENC_CAPS_PARAM

func newCapsParam(caps uint) *CapsParam {
	param := new(CapsParam)
	param.version = C.NV_ENC_CAPS_PARAM_VER
	param.capsToQuery = C.NV_ENC_CAPS(caps)
	return param
}

type ENC_STAT = C.NV_ENC_STAT

/*
uint32_t            version;                         /**< [in]:  Struct version. Must be set to ::NV_ENC_INITIALIZE_PARAMS_VER.
uint32_t            inBufferSize;                    /**< [in]:  Specifies the size of the spsppsBuffer provied by the client
uint32_t            spsId;                           /**< [in]:  Specifies the SPS id to be used in sequence header. Default value is 0.
uint32_t            ppsId;                           /**< [in]:  Specifies the PPS id to be used in picture header. Default value is 0.
void*               spsppsBuffer;                    /**< [in]:  Specifies bitstream header pointer of size NV_ENC_SEQUENCE_PARAM_PAYLOAD::inBufferSize. It is the client's responsibility to manage this memory.
uint32_t*           outSPSPPSPayloadSize;            /**< [out]: Size of the sequence and picture header in  bytes written by the NvEncodeAPI interface to the SPSPPSBuffer.
uint32_t            reserved [250];                  /**< [in]:  Reserved and must be set to 0
void*               reserved2[64];                   /**< [in]:  Reserved and must be set to NULL
*/
type SEQUENCE_PARAM_PAYLOAD = C.NV_ENC_SEQUENCE_PARAM_PAYLOAD

func (p *SEQUENCE_PARAM_PAYLOAD) Bytes() []byte {
	return C.GoBytes(p.spsppsBuffer, C.int(*p.outSPSPPSPayloadSize))
}

func newSequenceParamPayload() *SEQUENCE_PARAM_PAYLOAD {
	var payload = new(SEQUENCE_PARAM_PAYLOAD)
	bufsize := 128
	payload.version = C.NV_ENC_SEQUENCE_PARAM_PAYLOAD_VER
	payload.inBufferSize = C.uint32_t(bufsize)
	payload.spsppsBuffer = C.malloc(C.size_t(bufsize))
	payload.outSPSPPSPayloadSize = (*C.uint32_t)(C.malloc(C.size_t(4)))
	return payload
}

type EVENT_PARAMS = C.NV_ENC_EVENT_PARAMS
type MAP_INPUT_RESOURCE_PARAMS = C.NV_ENC_MAP_INPUT_RESOURCE
type REGISTER_RESOURCE_PARAMS = C.NV_ENC_REGISTER_RESOURCE
type RECONFIGURE_PARAMS = C.NV_ENC_RECONFIGURE_PARAMS

func (r *RECONFIGURE_PARAMS) SetInitializeParams(p InitializeParams) {
	r.version = C.NV_ENC_RECONFIGURE_PARAMS_VER
	r.reInitEncodeParams = C.NV_ENC_INITIALIZE_PARAMS(p)
	C.ResetEncoder(r)
}

type CREATE_MV_BUFFER_PARAMS = C.NV_ENC_CREATE_MV_BUFFER
type MOTION_ESTIMATE_ONLY_PARAMS = C.NV_ENC_MEONLY_PARAMS
type FRAME_FIELD_MODE = C.NV_ENC_PARAMS_FRAME_FIELD_MODE
type MV_PRECISION = C.NV_ENC_MV_PRECISION

type CODEC_CONFIG = C.NV_ENC_CODEC_CONFIG

func (c *CODEC_CONFIG) GetH264Config() *CONFIG_H264 {
	return C.GetH264Config(c)
}

func (c *CODEC_CONFIG) GetHEVCConfig() *CONFIG_HEVC {
	return C.GetHEVCConfig(c)
}
