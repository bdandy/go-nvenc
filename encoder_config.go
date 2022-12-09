package nvenc

// #include "include/types.h"
import "C"

// EncoderConfig is a config type from C library
/*
uint32_t                        version;                                     /**< [in]: Struct version. Must be set to ::NV_ENC_CONFIG_VER.
GUID                            profileGUID;                                 /**< [in]: Specifies the codec profile guid. If client specifies \p NV_ENC_CODEC_PROFILE_AUTOSELECT_GUID the NvEncodeAPI interface will select the appropriate codec profile.
uint32_t                        gopLength;                                   /**< [in]: Specifies the number of pictures in one GOP. Low latency application client can set goplength to NVENC_INFINITE_GOPLENGTH so that keyframes are not inserted automatically.
int32_t                         frameIntervalP;                              /**< [in]: Specifies the GOP pattern as follows: \p frameIntervalP = 0: I, 1: IPP, 2: IBP, 3: IBBP  If goplength is set to NVENC_INFINITE_GOPLENGTH \p frameIntervalP should be set to 1.
uint32_t                        monoChromeEncoding;                          /**< [in]: Set this to 1 to enable monochrome encoding for this session.
NV_ENC_PARAMS_FRAME_FIELD_MODE  frameFieldMode;                              /**< [in]: Specifies the frame/field mode. Check support for field encoding using ::NV_ENC_CAPS_SUPPORT_FIELD_ENCODING caps.
NV_ENC_MV_PRECISION             mvPrecision;                                 /**< [in]: Specifies the desired motion vector prediction precision.
NV_ENC_RC_PARAMS                rcParams;                                    /**< [in]: Specifies the rate control parameters for the current encoding session.
NV_ENC_CODEC_CONFIG             encodeCodecConfig;                           /**< [in]: Specifies the codec specific config parameters through this union.
uint32_t                        reserved [278];                              /**< [in]: Reserved and must be set to 0
void*                           reserved2[64];                               /**< [in]: Reserved and must be set to NULL
*/
type EncoderConfig C.NV_ENC_CONFIG

func (e *EncoderConfig) cType() *C.NV_ENC_CONFIG {
	return (*C.NV_ENC_CONFIG)(e)
}

func (e *EncoderConfig) RC() *RcParams {
	return (*RcParams)(&e.rcParams)
}

func (e *EncoderConfig) SetGOPLen(ln uint32) {
	e.gopLength = C.uint32_t(ln)
	if ln == uint32(InfiniteGOPLength) {
		e.frameIntervalP = 1
	}
}

func (e *EncoderConfig) SetFrameInterval(i int32) {
	e.frameIntervalP = C.int32_t(i)
}

func (e *EncoderConfig) UseMonoChrome(b bool) {
	if b {
		e.monoChromeEncoding = C.uint32_t(1)
	} else {
		e.monoChromeEncoding = C.uint32_t(0)
	}
}

func (e *EncoderConfig) SetProfile(guid profileGUID) {
	e.profileGUID = C.GUID(guid)
}

func (e *EncoderConfig) SetFrameFieldMode(mode FRAME_FIELD_MODE) {
	e.frameFieldMode = C.NV_ENC_PARAMS_FRAME_FIELD_MODE(mode)
}

func (e *EncoderConfig) SetMVPrecision(mv MV_PRECISION) {
	e.mvPrecision = C.NV_ENC_MV_PRECISION(mv)
}

func (e *EncoderConfig) getCodecConfig() *CODEC_CONFIG {
	return (*CODEC_CONFIG)(&e.encodeCodecConfig)
}

func (e *EncoderConfig) GetH264Config() *CONFIG_H264 {
	return e.getCodecConfig().GetH264Config()
}

func (e *EncoderConfig) GetHEVCConfig() *CONFIG_HEVC {
	return e.getCodecConfig().GetHEVCConfig()
}

func newEncoderConfig() *EncoderConfig {
	config := new(EncoderConfig)
	config.version = C.NV_ENC_CONFIG_VER
	config.rcParams = C.NV_ENC_RC_PARAMS(newRCParams())

	return config
}
