package types

// #cgo CFLAGS: -I ../../include
// #include "hevc.h"
import "C"

type PIC_PARAMS_HEVC = C.NV_ENC_PIC_PARAMS_HEVC
type CONFIG_HEVC = C.NV_ENC_CONFIG_HEVC

func (c *CONFIG_HEVC) SetLevel(level uint32) {
	c.level = C.uint32_t(level)
}

func (c *CONFIG_HEVC) SetTier(tier uint32) {
	c.tier = C.uint32_t(tier)
}

func (c *CONFIG_HEVC) RepeatSPSPPS() {
	C.HEVCSetRepeatSPSPPS((*C.NV_ENC_CONFIG_HEVC)(c))
}

func (c *CONFIG_HEVC) SetMaxRefFrames(max uint32) {
	c.maxNumRefFramesInDPB = C.uint32_t(max)
}
