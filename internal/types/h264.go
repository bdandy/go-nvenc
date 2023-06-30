package types

// #cgo CFLAGS: -I ../../include
// #include "h264.h"
import "C"

var (
	LEVEL_H264_1  = uint32(C.NV_ENC_LEVEL_H264_1)
	LEVEL_H264_1b = uint32(C.NV_ENC_LEVEL_H264_1b)
	LEVEL_H264_11 = uint32(C.NV_ENC_LEVEL_H264_11)
	LEVEL_H264_12 = uint32(C.NV_ENC_LEVEL_H264_12)
	LEVEL_H264_13 = uint32(C.NV_ENC_LEVEL_H264_13)
	LEVEL_H264_2  = uint32(C.NV_ENC_LEVEL_H264_2)
	LEVEL_H264_21 = uint32(C.NV_ENC_LEVEL_H264_21)
	LEVEL_H264_22 = uint32(C.NV_ENC_LEVEL_H264_22)
	LEVEL_H264_3  = uint32(C.NV_ENC_LEVEL_H264_3)
	LEVEL_H264_31 = uint32(C.NV_ENC_LEVEL_H264_31)
	LEVEL_H264_32 = uint32(C.NV_ENC_LEVEL_H264_32)
	LEVEL_H264_4  = uint32(C.NV_ENC_LEVEL_H264_4)
	LEVEL_H264_41 = uint32(C.NV_ENC_LEVEL_H264_41)
	LEVEL_H264_42 = uint32(C.NV_ENC_LEVEL_H264_42)
	LEVEL_H264_5  = uint32(C.NV_ENC_LEVEL_H264_5)
	LEVEL_H264_51 = uint32(C.NV_ENC_LEVEL_H264_51)
	LEVEL_H264_52 = uint32(C.NV_ENC_LEVEL_H264_52)
)

type CONFIG_H264 = C.NV_ENC_CONFIG_H264

func (c *CONFIG_H264) EnableVariableFPS() {
	C.EnableVFR(c)
}

func (c *CONFIG_H264) EnableConstrainedEncoding() {
	C.EnableConstrainedEncoding(c)
}

func (c *CONFIG_H264) SetChromaFormatIDC(f uint32) {
	c.chromaFormatIDC = C.uint32_t(f)
}

func (c *CONFIG_H264) EnableIntraRefresh() {
	C.EnableIntraRefresh(c)
}

func (c *CONFIG_H264) SetIntraRefreshPeriod(period uint32) {
	c.intraRefreshPeriod = C.uint32_t(period)
}

func (c *CONFIG_H264) SetIntraRefreshCount(count uint32) {
	c.intraRefreshCnt = C.uint32_t(count)
}

func (c *CONFIG_H264) SetLevel(level uint32) {
	c.level = C.uint32_t(level)
}

func (c *CONFIG_H264) RepeatSPSPPS() {
	C.EnableRepeatSPSPPS(c)
}

func (c *CONFIG_H264) DisableSPSPPS() {
	C.DisableSPSPPS(c)
}

func (c *CONFIG_H264) SetMaxRefFrames(num uint32) {
	c.maxNumRefFrames = C.uint32_t(num)
}

func (c *CONFIG_H264) EnableAUD() {
	C.EnableAUD(c)
}

func (c *CONFIG_H264) EnableLTRTrustMode(ltrNumFrames uint32) {
	C.EnableLTR(c)
	c.ltrTrustMode = 1
	c.ltrNumFrames = C.uint32_t(ltrNumFrames)
}

func (c *CONFIG_H264) UseConstrainedIntraPred() {
	C.UseConstrainedIntraPred(c)
}

func (c *CONFIG_H264) EnableOutputBufferingPeriodSEI() {
	C.EnableOutputBufferingPeriodSEI(c)
}

func (c *CONFIG_H264) EnableOutputPictureTimingSEI() {
	C.EnableOutputPictureTimingSEI(c)
}

func (c *CONFIG_H264) EnableOutputFramePackingSEI() {
	C.EnableOutputFramePackingSEI(c)
}

func (c *CONFIG_H264) EnableOutputRecoveryPointSEI() {
	C.EnableOutputRecoveryPointSEI(c)
}

func (c *CONFIG_H264) SliceCount(num int) {
	c.sliceMode = C.uint32_t(3)
	c.sliceModeData = C.uint32_t(num)
}

type PIC_PARAMS_H264 C.NV_ENC_PIC_PARAMS_H264

func (p *PIC_PARAMS_H264) ForceIntraRefresh(frames uint32) {
	p.forceIntraRefreshWithFrameCnt = C.uint32_t(frames)
}

func (p *PIC_PARAMS_H264) SetConstrainedFrame(enable int) {
	C.SetConstrainedFrame((*C.NV_ENC_PIC_PARAMS_H264)(p), C.int(enable))
}
