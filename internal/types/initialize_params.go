package types

// #cgo CFLAGS: -I ../../include
// #include "types.h"
import "C"
import (
	"unsafe"

	"github.com/bdandy/go-nvenc/v8/guid"
)

type InitializeParams C.NV_ENC_INITIALIZE_PARAMS

func (p *InitializeParams) SetFrameRate(num, den uint32) *InitializeParams {
	p.frameRateNum = C.uint32_t(num)
	p.frameRateDen = C.uint32_t(den)
	return p
}

func (p *InitializeParams) SetPresetGUID(guid guid.PresetGUID) *InitializeParams {
	t := guid.CType()
	p.presetGUID = *(*C.GUID)(unsafe.Pointer(&t))
	return p
}

func (p *InitializeParams) EnableWeightedPrediction() *InitializeParams {
	C.EnableWeightedPrediction((*C.NV_ENC_INITIALIZE_PARAMS)(p))
	return p
}

func (p *InitializeParams) SetEncodeConfig(conf *EncoderConfig) *InitializeParams {
	C.SetEncodeConfig((*C.NV_ENC_INITIALIZE_PARAMS)(p), (*C.NV_ENC_CONFIG)(conf))
	return p
}

func (p *InitializeParams) SetEncodeGUID(guid guid.CodecGUID) *InitializeParams {
	t := guid.CType()
	p.encodeGUID = *(*C.GUID)(unsafe.Pointer(&t))
	return p
}

func (p *InitializeParams) SetResolution(width, height uint32) *InitializeParams {
	p.encodeWidth = C.uint32_t(width)
	p.encodeHeight = C.uint32_t(height)
	return p
}

func (p *InitializeParams) SetAspectRatio(width, height uint32) *InitializeParams {
	p.darWidth = C.uint32_t(width)
	p.darHeight = C.uint32_t(height)
	return p
}

func (p *InitializeParams) SetEnablePTD(b bool) *InitializeParams {
	if b {
		p.enablePTD = 1
	} else {
		p.enablePTD = 0
	}
	return p
}

func NewInitializeParams() *InitializeParams {
	params := new(InitializeParams)
	params.version = C.NV_ENC_INITIALIZE_PARAMS_VER
	return params
}
