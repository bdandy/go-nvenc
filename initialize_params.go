package nvenc

// #include "include/types.h"
import "C"

// InitializeParams is a struct from C library.
/*
uint32_t                                   version;                         /**< [in]: Struct version. Must be set to ::NV_ENC_INITIALIZE_PARAMS_VER.
GUID                                       encodeGUID;                      /**< [in]: Specifies the Encode GUID for which the encoder is being created. ::NvEncInitializeEncoder() API will fail if this is not set, or set to unsupported value.
GUID                                       presetGUID;                      /**< [in]: Specifies the preset for encoding. If the preset GUID is set then , the preset configuration will be applied before any other parameter.
uint32_t                                   encodeWidth;                     /**< [in]: Specifies the encode width. If not set ::NvEncInitializeEncoder() API will fail.
uint32_t                                   encodeHeight;                    /**< [in]: Specifies the encode height. If not set ::NvEncInitializeEncoder() API will fail.
uint32_t                                   darWidth;                        /**< [in]: Specifies the display aspect ratio Width.
uint32_t                                   darHeight;                       /**< [in]: Specifies the display aspect ratio height.
uint32_t                                   frameRateNum;                    /**< [in]: Specifies the numerator for frame rate used for encoding in frames per second ( Frame rate = frameRateNum / frameRateDen ).
uint32_t                                   frameRateDen;                    /**< [in]: Specifies the denominator for frame rate used for encoding in frames per second ( Frame rate = frameRateNum / frameRateDen ).
uint32_t                                   enableEncodeAsync;               /**< [in]: Set this to 1 to enable asynchronous mode and is expected to use events to get picture completion notification.
uint32_t                                   enablePTD;                       /**< [in]: Set this to 1 to enable the Picture Type Decision is be taken by the NvEncodeAPI interface.
uint32_t                                   reportSliceOffsets        :1;    /**< [in]: Set this to 1 to enable reporting slice offsets in ::_NV_ENC_LOCK_BITSTREAM. NV_ENC_INITIALIZE_PARAMS::enableEncodeAsync must be set to 0 to use this feature. Client must set this to 0 if NV_ENC_CONFIG_H264::sliceMode is 1 on Kepler GPUs
uint32_t                                   enableSubFrameWrite       :1;    /**< [in]: Set this to 1 to write out available bitstream to memory at subframe intervals
uint32_t                                   enableExternalMEHints     :1;    /**< [in]: Set to 1 to enable external ME hints for the current frame. For NV_ENC_INITIALIZE_PARAMS::enablePTD=1 with B frames, programming L1 hints is optional for B frames since Client doesn't know internal GOP structure. NV_ENC_PIC_PARAMS::meHintRefPicDist should preferably be set with enablePTD=1.
uint32_t                                   enableMEOnlyMode          :1;    /**< [in]: Set to 1 to enable ME Only Mode .
uint32_t                                   enableWeightedPrediction  :1;    /**< [in]: Set this to 1 to enable weighted prediction. Not supported if encode session is configured for B-Frames( 'frameIntervalP' in NV_ENC_CONFIG is greater than 1).
uint32_t                                   reservedBitFields         :27;   /**< [in]: Reserved bitfields and must be set to 0
uint32_t                                   privDataSize;                    /**< [in]: Reserved private data buffer size and must be set to 0
void*                                      privData;                        /**< [in]: Reserved private data buffer and must be set to NULL
NV_ENC_CONFIG*                             encodeConfig;                    /**< [in]: Specifies the advanced codec specific structure. If client has sent a valid codec config structure, it will override parameters set by the NV_ENC_INITIALIZE_PARAMS::presetGUID parameter. If set to NULL the NvEncodeAPI interface will use the NV_ENC_INITIALIZE_PARAMS::presetGUID to set the codec specific parameters.

	Client can also optionally query the NvEncodeAPI interface to get codec specific parameters for a presetGUID using ::NvEncGetEncodePresetConfig() API. It can then modify (if required) some of the codec config parameters and send down a custom config structure as part of ::_NV_ENC_INITIALIZE_PARAMS.
	Even in this case client is recommended to pass the same preset guid it has used in ::NvEncGetEncodePresetConfig() API to query the config structure; as NV_ENC_INITIALIZE_PARAMS::presetGUID. This will not override the custom config structure but will be used to determine other Encoder HW specific parameters not exposed in the API.

uint32_t                                   maxEncodeWidth;                  /**< [in]: Maximum encode width to be used for current Encode session.

	Client should allocate output buffers according to this dimension for dynamic resolution change. If set to 0, Encoder will not allow dynamic resolution change.

uint32_t                                   maxEncodeHeight;                 /**< [in]: Maximum encode height to be allowed for current Encode session.

	Client should allocate output buffers according to this dimension for dynamic resolution change. If set to 0, Encode will not allow dynamic resolution change.

NVENC_EXTERNAL_ME_HINT_COUNTS_PER_BLOCKTYPE maxMEHintCountsPerBlock[2];      /**< [in]: If Client wants to pass external motion vectors in NV_ENC_PIC_PARAMS::meExternalHints buffer it must specify the maximum number of hint candidates per block per direction for the encode session.

	The NV_ENC_INITIALIZE_PARAMS::maxMEHintCountsPerBlock[0] is for L0 predictors and NV_ENC_INITIALIZE_PARAMS::maxMEHintCountsPerBlock[1] is for L1 predictors.
	This client must also set NV_ENC_INITIALIZE_PARAMS::enableExternalMEHints to 1.

uint32_t                                   reserved [289];                  /**< [in]: Reserved and must be set to 0
void*                                      reserved2[64];                   /**< [in]: Reserved and must be set to NULL
*/
type InitializeParams C.NV_ENC_INITIALIZE_PARAMS

func (p *InitializeParams) cType() *C.NV_ENC_INITIALIZE_PARAMS {
	return (*C.NV_ENC_INITIALIZE_PARAMS)(p)
}

func (p *InitializeParams) SetFrameRate(num, den uint32) *InitializeParams {
	p.frameRateNum = C.uint32_t(num)
	p.frameRateDen = C.uint32_t(den)
	return p
}

func (p *InitializeParams) SetPresetGUID(guid presetGUID) *InitializeParams {
	p.presetGUID = C.GUID(guid)
	return p
}

func (p *InitializeParams) EnableWeightedPrediction() *InitializeParams {
	C.EnableWeightedPrediction(p.cType())
	return p
}

func (p *InitializeParams) SetEncodeConfig(conf *EncoderConfig) *InitializeParams {
	C.SetEncodeConfig(p.cType(), conf.cType())
	return p
}

func (p *InitializeParams) SetEncodeGUID(guid codecGUID) *InitializeParams {
	p.encodeGUID = C.GUID(guid)
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

func newInitializeParams() *InitializeParams {
	params := new(InitializeParams)
	params.version = C.NV_ENC_INITIALIZE_PARAMS_VER
	return params
}
