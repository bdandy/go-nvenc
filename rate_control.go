package nvenc

// #include "include/types.h"
import "C"

// RcParams is rate control settings
/*
uint32_t                        version;
NV_ENC_PARAMS_RC_MODE           rateControlMode;                             /**< [in]: Specifies the rate control mode. Check support for various rate control modes using ::NV_ENC_CAPS_SUPPORTED_RATECONTROL_MODES caps.
NV_ENC_QP                       constQP;                                     /**< [in]: Specifies the initial QP to be used for encoding, these values would be used for all frames if in Constant QP mode.
uint32_t                        averageBitRate;                              /**< [in]: Specifies the average bitrate(in bits/sec) used for encoding.
uint32_t                        maxBitRate;                                  /**< [in]: Specifies the maximum bitrate for the encoded output. This is used for VBR and ignored for CBR mode.
uint32_t                        vbvBufferSize;                               /**< [in]: Specifies the VBV(HRD) buffer size. in bits. Set 0 to use the default VBV  buffer size.
uint32_t                        vbvInitialDelay;                             /**< [in]: Specifies the VBV(HRD) initial delay in bits. Set 0 to use the default VBV  initial delay .
uint32_t                        enableMinQP          :1;                     /**< [in]: Set this to 1 if minimum QP used for rate control.
uint32_t                        enableMaxQP          :1;                     /**< [in]: Set this to 1 if maximum QP used for rate control.
uint32_t                        enableInitialRCQP    :1;                     /**< [in]: Set this to 1 if user suppplied initial QP is used for rate control.
uint32_t                        enableAQ             :1;                     /**< [in]: Set this to 1 to enable adaptive quantization.
uint32_t                        enableExtQPDeltaMap  :1;                     /**< [in]: Set this to 1 to enable additional QP modifier for each MB supplied by client though signed byte array pointed to by NV_ENC_PIC_PARAMS::qpDeltaMap
uint32_t                        reservedBitFields    :27;                    /**< [in]: Reserved bitfields and must be set to 0
NV_ENC_QP                       minQP;                                       /**< [in]: Specifies the minimum QP used for rate control. Client must set NV_ENC_CONFIG::enableMinQP to 1.
NV_ENC_QP                       maxQP;                                       /**< [in]: Specifies the maximum QP used for rate control. Client must set NV_ENC_CONFIG::enableMaxQP to 1.
NV_ENC_QP                       initialRCQP;                                 /**< [in]: Specifies the initial QP used for rate control. Client must set NV_ENC_CONFIG::enableInitialRCQP to 1.
uint32_t                        temporallayerIdxMask;                        /**< [in]: Specifies the temporal layers (as a bitmask) whose QPs have changed. Valid max bitmask is [2^NV_ENC_CAPS_NUM_MAX_TEMPORAL_LAYERS - 1]
uint8_t                         temporalLayerQP[8];                          /**< [in]: Specifies the temporal layer QPs used for rate control. Temporal layer index is used as as the array index
uint32_t                        reserved[10];
*/
type RcParams C.NV_ENC_RC_PARAMS

func (p *RcParams) cType() *C.NV_ENC_RC_PARAMS {
	return (*C.NV_ENC_RC_PARAMS)(p)
}

func (p *RcParams) setVersion() {
	p.version = C.NV_ENC_RC_PARAMS_VER
}

func (p *RcParams) EnableSpartialAQ(strength int) {
	C.EnableSpartialAQ(p.cType(), C.int(strength))
}

func (p *RcParams) EnableTemporalAQ() {
	C.EnableTemporalAQ(p.cType())
}

func (p *RcParams) EnableZeroReorderDelay() {
	C.EnableZeroReorderDelay(p.cType())
}

func (p *RcParams) SetMinQP(pQP, bQP, iQP uint32) {
	C.EnableMinQP(p.cType())
	p.minQP = C.NV_ENC_QP{C.uint32_t(pQP), C.uint32_t(bQP), C.uint32_t(iQP)}
}

func (p *RcParams) SetMaxQP(pQP, bQP, iQP uint32) {
	C.EnableMaxQP(p.cType())
	p.maxQP = C.NV_ENC_QP{C.uint32_t(pQP), C.uint32_t(bQP), C.uint32_t(iQP)}
}

func (p *RcParams) EnableNonRefP() {
	C.EnableNonRefP(p.cType())
}

func (p *RcParams) UseConstQP() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CONSTQP
}

func (p *RcParams) UseVBR() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_VBR
}

func (p *RcParams) UseVBR_HQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_VBR_HQ
}

func (p *RcParams) UseCBR() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR
}

func (p *RcParams) UseCBR_LowDelayHQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR_LOWDELAY_HQ
}

func (p *RcParams) UseCBR_HQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR_HQ
}

func (p *RcParams) SetVBVBufSize(size, delay uint32) {
	p.vbvBufferSize = C.uint32_t(size)
	p.vbvInitialDelay = C.uint32_t(delay)
}

func (p *RcParams) SetAverageBitRate(rate uint32) {
	p.averageBitRate = C.uint32_t(rate)
}

func (p *RcParams) SetMaxBitRate(rate uint32) {
	p.maxBitRate = C.uint32_t(rate)
}

func newRCParams() (param RcParams) {
	param.setVersion()
	return param
}
