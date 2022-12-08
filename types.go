package nvenc

/*
 #include <stdio.h>
 #include "headers/nvEncodeAPI.h"

 static inline NV_ENC_CONFIG_H264* GetH264Config(NV_ENC_CODEC_CONFIG* c) {
 	return &c->h264Config;
 }

 static inline NV_ENC_CONFIG_HEVC* GetHEVCConfig(NV_ENC_CODEC_CONFIG* c) {
 	return &c->hevcConfig;
 }

 static inline NV_ENC_PIC_PARAMS_H264* GetPicParamsH264(NV_ENC_CODEC_PIC_PARAMS* p) {
 	return &p->h264PicParams;
 }

 static inline void EnableMinQP(NV_ENC_RC_PARAMS* c) {
 	c->enableMinQP = 1;
 }

 static inline void EnableMaxQP(NV_ENC_RC_PARAMS* c) {
 	c->enableMaxQP = 1;
 }

 static inline void EnableTemporalAQ(NV_ENC_RC_PARAMS* c) {
 	c->enableTemporalAQ =1;
 }

 static inline void EnableSpartialAQ(NV_ENC_RC_PARAMS* c, int s) {
 	c->enableAQ = 1;
	c->aqStrength = s;
 }

 static inline void EnableZeroReorderDelay(NV_ENC_RC_PARAMS* c) {
 	c->zeroReorderDelay = 1;
 }

 static inline void EnableNonRefP(NV_ENC_RC_PARAMS* c) {
 	c->enableNonRefP = 1;
 }

 static inline void ResetEncoder(NV_ENC_RECONFIGURE_PARAMS* p) {
 	p->resetEncoder = 1;
 }

 static inline void EnableWeightedPrediction(NV_ENC_INITIALIZE_PARAMS* p) {
 	p->enableWeightedPrediction = 1;
 }

 static inline void SetEncodeConfig(NV_ENC_INITIALIZE_PARAMS* p, NV_ENC_CONFIG* c) {
    p->encodeConfig = c;
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var (
	DEVICE_TYPE_DIRECTX = C.NV_ENC_DEVICE_TYPE_DIRECTX
	DEVICE_TYPE_CUDA    = C.NV_ENC_DEVICE_TYPE_CUDA

	BUFFER_FORMAT_NV12   = BUFFER_FORMAT(0x1)       /**< Semi-Planar YUV [Y plane followed by interleaved UV plane] */
	BUFFER_FORMAT_YV12   = BUFFER_FORMAT(0x10)      /**< Planar YUV [Y plane followed by V and U planes] */
	BUFFER_FORMAT_IYUV   = BUFFER_FORMAT(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BUFFER_FORMAT_I420   = BUFFER_FORMAT(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BUFFER_FORMAT_YUV420 = BUFFER_FORMAT(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BUFFER_FORMAT_YUV444 = BUFFER_FORMAT(0x1000)    /**< Planar YUV [Y plane followed by U and V planes] */
	BUFFER_FORMAT_ARGB   = BUFFER_FORMAT(0x1000000) /**< 8 bit Packed A8R8G8B8 */
	BUFFER_FORMAT_ARGB10 = BUFFER_FORMAT(0x2000000) /**< 10 bit Packed A2R10G10B10 */
	BUFFER_FORMAT_AYUV   = BUFFER_FORMAT(0x4000000) /**< 8 bit Packed A8Y8U8V8 */

	INFINITE_GOPLENGTH = uint32(C.NVENC_INFINITE_GOPLENGTH)
	LEVEL_AUTOSELECT   = uint32(C.NV_ENC_LEVEL_AUTOSELECT)
)

/*
uint32_t            version;                          //< [in]: Struct version. Must be set to ::NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS_VER.
NV_ENC_DEVICE_TYPE  deviceType;                       //< [in]: Specified the device Type
void*               device;                           //< [in]: Pointer to client device.
void*               reserved;                         //< [in]: Reserved and must be set to 0.
uint32_t            apiVersion;                       //< [in]: API version. Should be set to NVENCAPI_VERSION.
uint32_t            reserved1[253];                   //< [in]: Reserved and must be set to 0
void*               reserved2[64];                    //< [in]: Reserved and must be set to NULL
*/
type OPEN_ENCODE_SESSION_PARAMS = C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS

func newOpenEncodeSessionParams(devType int, device unsafe.Pointer) *OPEN_ENCODE_SESSION_PARAMS {
	params := new(OPEN_ENCODE_SESSION_PARAMS)
	params.version = C.NV_ENC_OPEN_ENCODE_SESSION_EX_PARAMS_VER
	params.deviceType = C.NV_ENC_DEVICE_TYPE(devType)
	params.device = device
	params.apiVersion = C.NVENCAPI_VERSION

	return params
}

type BUFFER_FORMAT = C.NV_ENC_BUFFER_FORMAT

func (format BUFFER_FORMAT) String() string {
	switch format {
	case BUFFER_FORMAT_NV12:
		return "NV12"
	case BUFFER_FORMAT_YV12:
		return "YV12"
	case BUFFER_FORMAT_IYUV:
		return "IYUV"
	case BUFFER_FORMAT_YUV444:
		return "YUV444"
	case BUFFER_FORMAT_ARGB:
		return "ARGB"
	case BUFFER_FORMAT_ARGB10:
		return "ARGB10"
	case BUFFER_FORMAT_AYUV:
		return "AYUV"
	}

	return fmt.Sprintf("unknown:%x", int(format))
}

/*
typedef struct _NV_ENC_INITIALIZE_PARAMS

	{
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
uint32_t                                   enableExternalMEHints     :1;    /**< [in]: Set to 1 to enable external ME hints for the current frame. For NV_ENC_INITIALIZE_PARAMS::enablePTD=1 with B frames, programming L1 hints is optional for B frames since Client doesn't know internal GOP structure.

	NV_ENC_PIC_PARAMS::meHintRefPicDist should preferably be set with enablePTD=1.

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
} NV_ENC_INITIALIZE_PARAMS;
*/
type INITIALIZE_PARAMS = C.NV_ENC_INITIALIZE_PARAMS

func (p *INITIALIZE_PARAMS) SetFrameRate(num, den uint32) {
	p.frameRateNum = C.uint32_t(num)
	p.frameRateDen = C.uint32_t(den)
}

func (p *INITIALIZE_PARAMS) SetPresetGUID(guid GUID) {
	p.presetGUID = C.GUID(guid)
}

func (p *INITIALIZE_PARAMS) EnableWeightedPrediction() {
	C.EnableWeightedPrediction(p)
}

func (p *INITIALIZE_PARAMS) SetEncodeConfig(conf *ENCODER_CONFIG) {
	C.SetEncodeConfig(p, conf)
}

func (p *INITIALIZE_PARAMS) SetEncodeGUID(guid GUID) {
	p.encodeGUID = C.GUID(guid)
}

func (p *INITIALIZE_PARAMS) SetResolution(width, height uint32) {
	p.encodeWidth = C.uint32_t(width)
	p.encodeHeight = C.uint32_t(height)
	p.darWidth = C.uint32_t(width)
	p.darHeight = C.uint32_t(height)

}

func newInitializeParams() *INITIALIZE_PARAMS {
	params := new(INITIALIZE_PARAMS)
	params.version = C.NV_ENC_INITIALIZE_PARAMS_VER
	params.enablePTD = 1

	return params
}

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
type ENCODER_CONFIG = C.NV_ENC_CONFIG

func (e *ENCODER_CONFIG) RC() *RC_PARAMS {
	return &e.rcParams
}

func (e *ENCODER_CONFIG) SetGOPLen(ln uint32) {
	e.gopLength = C.uint32_t(ln)
	if ln == INFINITE_GOPLENGTH {
		e.frameIntervalP = 1
	}
}

func (e *ENCODER_CONFIG) SetFrameInterval(i int32) {
	e.frameIntervalP = C.int32_t(i)
}

func (e *ENCODER_CONFIG) UseMonoChrome(b bool) {
	if b {
		e.monoChromeEncoding = C.uint32_t(1)
	} else {
		e.monoChromeEncoding = C.uint32_t(0)
	}
}

func (e *ENCODER_CONFIG) SetProfile(guid profileGUID) {
	e.profileGUID = C.GUID(guid)
}

func (e *ENCODER_CONFIG) SetFrameFieldMode(mode FRAME_FIELD_MODE) {
	e.frameFieldMode = C.NV_ENC_PARAMS_FRAME_FIELD_MODE(mode)
}

func (e *ENCODER_CONFIG) SetMVPrecision(mv MV_PRECISION) {
	e.mvPrecision = C.NV_ENC_MV_PRECISION(mv)
}

func (e *ENCODER_CONFIG) getCodecConfig() *CODEC_CONFIG {
	return (*CODEC_CONFIG)(&e.encodeCodecConfig)
}
func (e *ENCODER_CONFIG) GetH264Config() *CONFIG_H264 {
	return e.getCodecConfig().GetH264Config()
}

func (e *ENCODER_CONFIG) GetHEVCConfig() *CONFIG_HEVC {
	return e.getCodecConfig().GetHEVCConfig()
}

func newEncoderConfig() *ENCODER_CONFIG {
	config := new(ENCODER_CONFIG)
	config.version = C.NV_ENC_CONFIG_VER
	config.rcParams = C.NV_ENC_RC_PARAMS(newRCParams())

	return config
}

/*
uint32_t version;                                  /**< [in]: Struct version. Must be set to ::NV_ENC_CAPS_PARAM_VER
NV_ENC_CAPS  capsToQuery;                          /**< [in]: Specifies the encode capability to be queried. Client should pass a member for ::NV_ENC_CAPS enum.
uint32_t reserved[62];                             /**< [in]: Reserved and must be set to 0
*/
type CAPS_PARAM = C.NV_ENC_CAPS_PARAM

func newCapsParam(caps uint) *CAPS_PARAM {
	param := new(CAPS_PARAM)
	param.version = C.NV_ENC_CAPS_PARAM_VER
	param.capsToQuery = C.NV_ENC_CAPS(caps)
	return param
}

/*
uint32_t                  version;                 /**< [in]: Struct version. Must be set to ::NV_ENC_CREATE_INPUT_BUFFER_VER
uint32_t                  width;                   /**< [in]: Input buffer width
uint32_t                  height;                  /**< [in]: Input buffer width
NV_ENC_MEMORY_HEAP        memoryHeap;              /**< [in]: Input buffer memory heap
NV_ENC_BUFFER_FORMAT      bufferFmt;               /**< [in]: Input buffer format
uint32_t                  reserved;                /**< [in]: Reserved and must be set to 0
NV_ENC_INPUT_PTR          inputBuffer;             /**< [out]: Pointer to input buffer
void*                     pSysMemBuffer;           /**< [in]: Pointer to existing sysmem buffer
uint32_t                  reserved1[57];           /**< [in]: Reserved and must be set to 0
void*                     reserved2[63];           /**< [in]: Reserved and must be set to NULL
*/
type INPUT_BUFFER = C.NV_ENC_CREATE_INPUT_BUFFER

func (p *INPUT_BUFFER) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(p.inputBuffer)
}

func (p *INPUT_BUFFER) SetResolution(width, height uint32) {
	p.width = C.uint32_t(width)
	p.height = C.uint32_t(height)
}

func (p *INPUT_BUFFER) SetFormat(format BUFFER_FORMAT) {
	p.bufferFmt = C.NV_ENC_BUFFER_FORMAT(format)
}

func newInputBufferParams() *INPUT_BUFFER {
	params := new(INPUT_BUFFER)
	params.version = C.NV_ENC_CREATE_INPUT_BUFFER_VER
	return params
}

/*
uint32_t              version;                     /**< [in]: Struct version. Must be set to ::NV_ENC_CREATE_BITSTREAM_BUFFER_VER
uint32_t              size;                        /**< [in]: Size of the bitstream buffer to be created
NV_ENC_MEMORY_HEAP    memoryHeap;                  /**< [in]: Output buffer memory heap
uint32_t              reserved;                    /**< [in]: Reserved and must be set to 0
NV_ENC_OUTPUT_PTR     bitstreamBuffer;             /**< [out]: Pointer to the output bitstream buffer
void*                 bitstreamBufferPtr;          /**< [out]: Reserved and should not be used
uint32_t              reserved1[58];               /**< [in]: Reserved and should be set to 0
void*                 reserved2[64];               /**< [in]: Reserved and should be set to NULL
*/
type BITSTREAM_BUFFER = C.NV_ENC_CREATE_BITSTREAM_BUFFER

func (b *BITSTREAM_BUFFER) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(b.bitstreamBuffer)
}

func newBitstreamBuffer(size uint32) *BITSTREAM_BUFFER {
	buffer := new(BITSTREAM_BUFFER)
	buffer.version = C.NV_ENC_CREATE_BITSTREAM_BUFFER_VER
	buffer.size = C.uint32_t(size)
	return buffer
}

/*
uint32_t                                    version;                        /**< [in]: Struct version. Must be set to ::NV_ENC_PIC_PARAMS_VER.
uint32_t                                    inputWidth;                     /**< [in]: Specifies the input buffer width
uint32_t                                    inputHeight;                    /**< [in]: Specifies the input buffer height
uint32_t                                    inputPitch;                     /**< [in]: Specifies the input buffer pitch. If pitch value is not known, set this to inputWidth.
uint32_t                                    encodePicFlags;                 /**< [in]: Specifies bit-wise OR`ed encode pic flags. See ::NV_ENC_PIC_FLAGS enum.
uint32_t                                    frameIdx;                       /**< [in]: Specifies the frame index associated with the input frame [optional].
uint64_t                                    inputTimeStamp;                 /**< [in]: Specifies presentation timestamp associated with the input picture.
uint64_t                                    inputDuration;                  /**< [in]: Specifies duration of the input picture
NV_ENC_INPUT_PTR                            inputBuffer;                    /**< [in]: Specifies the input buffer pointer. Client must use a pointer obtained from ::NvEncCreateInputBuffer() or ::NvEncMapInputResource() APIs.
NV_ENC_OUTPUT_PTR                           outputBitstream;                /**< [in]: Specifies the pointer to output buffer. Client should use a pointer obtained from ::NvEncCreateBitstreamBuffer() API.
void*                                       completionEvent;                /**< [in]: Specifies an event to be signalled on completion of encoding of this Frame [only if operating in Asynchronous mode]. Each output buffer should be associated with a distinct event pointer.
NV_ENC_BUFFER_FORMAT                        bufferFmt;                      /**< [in]: Specifies the input buffer format.
NV_ENC_PIC_STRUCT                           pictureStruct;                  /**< [in]: Specifies structure of the input picture.
NV_ENC_PIC_TYPE                             pictureType;                    /**< [in]: Specifies input picture type. Client required to be set explicitly by the client if the client has not set NV_ENC_INITALIZE_PARAMS::enablePTD to 1 while calling NvInitializeEncoder.
NV_ENC_CODEC_PIC_PARAMS                     codecPicParams;                 /**< [in]: Specifies the codec specific per-picture encoding parameters.
NVENC_EXTERNAL_ME_HINT_COUNTS_PER_BLOCKTYPE meHintCountsPerBlock[2];        /**< [in]: Specifies the number of hint candidates per block per direction for the current frame. meHintCountsPerBlock[0] is for L0 predictors and meHintCountsPerBlock[1] is for L1 predictors.

	The candidate count in NV_ENC_PIC_PARAMS::meHintCountsPerBlock[lx] must never exceed NV_ENC_INITIALIZE_PARAMS::maxMEHintCountsPerBlock[lx] provided during encoder intialization.

NVENC_EXTERNAL_ME_HINT                     *meExternalHints;                /**< [in]: Specifies the pointer to ME external hints for the current frame. The size of ME hint buffer should be equal to number of macroblocks multiplied by the total number of candidates per macroblock.

	The total number of candidates per MB per direction = 1*meHintCountsPerBlock[Lx].numCandsPerBlk16x16 + 2*meHintCountsPerBlock[Lx].numCandsPerBlk16x8 + 2*meHintCountsPerBlock[Lx].numCandsPerBlk8x8
	+ 4*meHintCountsPerBlock[Lx].numCandsPerBlk8x8. For frames using bidirectional ME , the total number of candidates for single macroblock is sum of total number of candidates per MB for each direction (L0 and L1)

uint32_t                                    reserved1[6];                    /**< [in]: Reserved and must be set to 0
void*                                       reserved2[2];                    /**< [in]: Reserved and must be set to NULL
int8_t                                     *qpDeltaMap;                      /**< [in]: Specifies the pointer to signed byte array containing QP delta value per MB in raster scan order in the current picture. This QP modifier is applied on top of the QP chosen by rate control.
uint32_t                                    qpDeltaMapSize;                  /**< [in]: Specifies the size in bytes of qpDeltaMap surface allocated by client and pointed to by NV_ENC_PIC_PARAMS::qpDeltaMap. Surface (array) should be picWidthInMbs * picHeightInMbs
uint32_t                                    reservedBitFields;               /**< [in]: Reserved bitfields and must be set to 0
uint32_t                                    reserved3[287];                  /**< [in]: Reserved and must be set to 0
void*                                       reserved4[60];                   /**< [in]: Reserved and must be set to NULL
*/
type ENC_PIC_PARAMS = C.NV_ENC_PIC_PARAMS

func (p *ENC_PIC_PARAMS) SetTimeStamp(t uint64) {
	p.inputTimeStamp = C.uint64_t(t)
}

func (p *ENC_PIC_PARAMS) SetH264Params(param *PIC_PARAMS_H264) {
	cast := (*C.NV_ENC_CODEC_PIC_PARAMS)(unsafe.Pointer(param))
	p.codecPicParams = C.NV_ENC_CODEC_PIC_PARAMS(*cast)
}

func (p *ENC_PIC_PARAMS) SetHEVCParams(param *PIC_PARAMS_HEVC) {
	cast := (*C.NV_ENC_CODEC_PIC_PARAMS)(unsafe.Pointer(param))
	p.codecPicParams = C.NV_ENC_CODEC_PIC_PARAMS(*cast)
}

func (p *ENC_PIC_PARAMS) SetInputBuffer(point unsafe.Pointer) {
	p.inputBuffer = C.NV_ENC_INPUT_PTR(point)
}

func (p *ENC_PIC_PARAMS) SetOutputBuffer(point unsafe.Pointer) {
	p.outputBitstream = C.NV_ENC_OUTPUT_PTR(point)
}

// NV_ENC_PIC_FLAG_FORCEINTRA         = 0x1,
// Encode the current picture as an Intra picture
func (p *ENC_PIC_PARAMS) ForceIntra(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_FORCEINTRA
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_FORCEINTRA
	}
}

/*
NV_ENC_PIC_FLAG_FORCEIDR           = 0x2
Encode the current picture as an IDR picture.
This flag is only valid when Picture type decision is taken by the Encoder
[_NV_ENC_INITIALIZE_PARAMS::enablePTD == 1].
*/
func (p *ENC_PIC_PARAMS) ForceIDR(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_FORCEIDR
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_FORCEIDR
	}
}

// NV_ENC_PIC_FLAG_OUTPUT_SPSPPS      = 0x4
// Write the sequence and picture header in encoded bitstream of the current picture
func (p *ENC_PIC_PARAMS) ForceSPSPPS(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_OUTPUT_SPSPPS
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_OUTPUT_SPSPPS
	}
}

// NV_ENC_PIC_FLAG_EOS                = 0x8
// Indicates end of the input stream
func (p *ENC_PIC_PARAMS) ForceEOS(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_EOS
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_EOS
	}
}

func (p *ENC_PIC_PARAMS) SetResolution(width, height uint32) {
	p.inputWidth = C.uint32_t(width)
	p.inputHeight = C.uint32_t(height)
	p.inputPitch = C.uint32_t(width)
}

func (p *ENC_PIC_PARAMS) SetPitch(pitch uint32) {
	p.inputPitch = C.uint32_t(pitch)
}

func (p *ENC_PIC_PARAMS) SetInputFormat(format BUFFER_FORMAT) {
	p.bufferFmt = C.NV_ENC_BUFFER_FORMAT(format)
}

func (p *ENC_PIC_PARAMS) PicParamsH264() *PIC_PARAMS_H264 {
	return (*PIC_PARAMS_H264)(C.GetPicParamsH264(&p.codecPicParams))
}

func newEncPicParams() *ENC_PIC_PARAMS {
	params := new(ENC_PIC_PARAMS)
	params.version = C.NV_ENC_PIC_PARAMS_VER
	params.pictureStruct = C.NV_ENC_PIC_STRUCT_FRAME

	return params
}

/*
uint32_t                version;                     /**< [in]: Struct version. Must be set to ::NV_ENC_LOCK_BITSTREAM_VER.
uint32_t                doNotWait         :1;        /**< [in]: If this flag is set, the NvEncodeAPI interface will return buffer pointer even if operation is not completed. If not set, the call will block until operation completes.
uint32_t                ltrFrame          :1;        /**< [out]: Flag indicating this frame is marked as LTR frame
uint32_t                reservedBitFields :30;       /**< [in]: Reserved bit fields and must be set to 0
void*                   outputBitstream;             /**< [in]: Pointer to the bitstream buffer being locked.
uint32_t*               sliceOffsets;                /**< [in,out]: Array which receives the slice offsets. This is not supported if NV_ENC_CONFIG_H264::sliceMode is 1 on Kepler GPUs. Array size must be equal to size of frame in MBs.
uint32_t                frameIdx;                    /**< [out]: Frame no. for which the bitstream is being retrieved.
uint32_t                hwEncodeStatus;              /**< [out]: The NvEncodeAPI interface status for the locked picture.
uint32_t                numSlices;                   /**< [out]: Number of slices in the encoded picture. Will be reported only if NV_ENC_INITIALIZE_PARAMS::reportSliceOffsets set to 1.
uint32_t                bitstreamSizeInBytes;        /**< [out]: Actual number of bytes generated and copied to the memory pointed by bitstreamBufferPtr.
uint64_t                outputTimeStamp;             /**< [out]: Presentation timestamp associated with the encoded output.
uint64_t                outputDuration;              /**< [out]: Presentation duration associates with the encoded output.
void*                   bitstreamBufferPtr;          /**< [out]: Pointer to the generated output bitstream.
NV_ENC_PIC_TYPE         pictureType;                 /**< [out]: Picture type of the encoded picture.
NV_ENC_PIC_STRUCT       pictureStruct;               /**< [out]: Structure of the generated output picture.
uint32_t                frameAvgQP;                  /**< [out]: Average QP of the frame.
uint32_t                frameSatd;                   /**< [out]: Total SATD cost for whole frame.
uint32_t                ltrFrameIdx;                 /**< [out]: Frame index associated with this LTR frame.
uint32_t                ltrFrameBitmap;              /**< [out]: Bitmap of LTR frames indices which were used for encoding this frame. Value of 0 if no LTR frames were used.
uint32_t                reserved [236];              /**< [in]: Reserved and must be set to 0
void*                   reserved2[64];               /**< [in]: Reserved and must be set to NULL
*/
type LOCK_BITSTREAM_PARAMS = C.NV_ENC_LOCK_BITSTREAM

func (b *LOCK_BITSTREAM_PARAMS) BitstreamSize() int {
	return int(b.bitstreamSizeInBytes)
}

func (b *LOCK_BITSTREAM_PARAMS) CopyBitstream(buf []byte) error {
	if len(buf) < b.BitstreamSize() {
		return fmt.Errorf("bufSize is %d bytes, but bitstream need %d bytes for output", len(buf), b.BitstreamSize())
	}

	C.memcpy((unsafe.Pointer)(&buf[0]), b.bitstreamBufferPtr, C.size_t(b.BitstreamSize()))

	return nil
}

func (b *LOCK_BITSTREAM_PARAMS) FrameId() uint32 {
	return uint32(b.frameIdx)
}

func (b *LOCK_BITSTREAM_PARAMS) FrameTimeStamp() uint64 {
	return uint64(b.outputTimeStamp)
}

func newBitstreamBufferLock(buf *BITSTREAM_BUFFER) *LOCK_BITSTREAM_PARAMS {
	params := new(LOCK_BITSTREAM_PARAMS)
	params.version = C.NV_ENC_LOCK_BITSTREAM_VER
	params.outputBitstream = buf.GetBufferPtr()
	return params
}

/*
uint32_t                  version;                   /**< [in]:  Struct version. Must be set to ::NV_ENC_LOCK_INPUT_BUFFER_VER.
uint32_t                  doNotWait         :1;      /**< [in]:  Set to 1 to make ::NvEncLockInputBuffer() a unblocking call. If the encoding is not completed, driver will return ::NV_ENC_ERR_ENCODER_BUSY error code.
uint32_t                  reservedBitFields :31;     /**< [in]:  Reserved bitfields and must be set to 0
NV_ENC_INPUT_PTR          inputBuffer;               /**< [in]:  Pointer to the input buffer to be locked, client should pass the pointer obtained from ::NvEncCreateInputBuffer() or ::NvEncMapInputResource API.
void*                     bufferDataPtr;             /**< [out]: Pointed to the locked input buffer data. Client can only access input buffer using the \p bufferDataPtr.
uint32_t                  pitch;                     /**< [out]: Pitch of the locked input buffer.
uint32_t                  reserved1[251];            /**< [in]:  Reserved and must be set to 0
void*                     reserved2[64];             /**< [in]:  Reserved and must be set to NULL
*/
type LOCK_INPUT_BUFFER_PARAMS = C.NV_ENC_LOCK_INPUT_BUFFER

func (p *LOCK_INPUT_BUFFER_PARAMS) Pitch() uint32 {
	return uint32(p.pitch)
}

func (p *LOCK_INPUT_BUFFER_PARAMS) Buffer(size int) []byte {
	return C.GoBytes(unsafe.Pointer(p.bufferDataPtr), C.int(size))
}

func (p *LOCK_INPUT_BUFFER_PARAMS) CopyBuffer(buf []byte) {
	bufferData := p.Buffer(len(buf))
	copy(bufferData, buf)
}

func newLockInputBufferParams(b *INPUT_BUFFER) *LOCK_INPUT_BUFFER_PARAMS {
	params := new(LOCK_INPUT_BUFFER_PARAMS)
	params.version = C.NV_ENC_LOCK_INPUT_BUFFER_VER
	params.inputBuffer = b.inputBuffer

	return params
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

func (r *RECONFIGURE_PARAMS) SetInitializeParams(p INITIALIZE_PARAMS) {
	r.version = C.NV_ENC_RECONFIGURE_PARAMS_VER
	r.reInitEncodeParams = C.NV_ENC_INITIALIZE_PARAMS(p)
	C.ResetEncoder(r)
}

type CREATE_MV_BUFFER_PARAMS = C.NV_ENC_CREATE_MV_BUFFER
type MOTION_ESTIMATE_ONLY_PARAMS = C.NV_ENC_MEONLY_PARAMS
type FRAME_FIELD_MODE = C.NV_ENC_PARAMS_FRAME_FIELD_MODE
type MV_PRECISION = C.NV_ENC_MV_PRECISION

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
type RC_PARAMS = C.NV_ENC_RC_PARAMS

func (p *RC_PARAMS) setVersion() {
	p.version = C.NV_ENC_RC_PARAMS_VER
}

func (p *RC_PARAMS) EnableSpartialAQ(strength int) {
	C.EnableSpartialAQ(p, C.int(strength))
}

func (p *RC_PARAMS) EnableTemporalAQ() {
	C.EnableTemporalAQ(p)
}

func (p *RC_PARAMS) EnableZeroReorderDelay() {
	C.EnableZeroReorderDelay(p)
}

func (p *RC_PARAMS) SetMinQP(pQP, bQP, iQP uint32) {
	C.EnableMinQP(p)
	p.minQP = C.NV_ENC_QP{C.uint32_t(pQP), C.uint32_t(bQP), C.uint32_t(iQP)}
}

func (p *RC_PARAMS) SetMaxQP(pQP, bQP, iQP uint32) {
	C.EnableMaxQP(p)
	p.maxQP = C.NV_ENC_QP{C.uint32_t(pQP), C.uint32_t(bQP), C.uint32_t(iQP)}
}

func (p *RC_PARAMS) EnableNonRefP() {
	C.EnableNonRefP(p)
}

/*
NV_ENC_PARAMS_RC_CONSTQP                = 0x0,       /**< Constant QP mode
NV_ENC_PARAMS_RC_VBR                    = 0x1,       /**< Variable bitrate mode
NV_ENC_PARAMS_RC_CBR                    = 0x2,       /**< Constant bitrate mode
NV_ENC_PARAMS_RC_CBR_LOWDELAY_HQ        = 0x8,       /**< low-delay CBR, high quality
NV_ENC_PARAMS_RC_CBR_HQ                 = 0x10,      /**< CBR, high quality (slower)
NV_ENC_PARAMS_RC_VBR_HQ                 = 0x20       /**< VBR, high quality (slower)
*/
func (p *RC_PARAMS) UseConstQP() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CONSTQP
}

func (p *RC_PARAMS) UseVBR() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_VBR
}

func (p *RC_PARAMS) UseVBR_HQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_VBR_HQ
}

func (p *RC_PARAMS) UseCBR() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR
}

func (p *RC_PARAMS) UseCBR_LowDelayHQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR_LOWDELAY_HQ
}

func (p *RC_PARAMS) UseCBR_HQ() {
	p.rateControlMode = C.NV_ENC_PARAMS_RC_CBR_HQ
}

func (p *RC_PARAMS) SetVBVBufSize(size, delay uint32) {
	p.vbvBufferSize = C.uint32_t(size)
	p.vbvInitialDelay = C.uint32_t(delay)
}

func (p *RC_PARAMS) SetAverageBitRate(rate uint32) {
	p.averageBitRate = C.uint32_t(rate)
}

func (p *RC_PARAMS) SetMaxBitRate(rate uint32) {
	p.maxBitRate = C.uint32_t(rate)
}

func newRCParams() (param RC_PARAMS) {
	param.setVersion()
	return param
}

type CODEC_CONFIG = C.NV_ENC_CODEC_CONFIG

func (c *CODEC_CONFIG) GetH264Config() *CONFIG_H264 {
	return C.GetH264Config(c)
}

func (c *CODEC_CONFIG) GetHEVCConfig() *CONFIG_HEVC {
	return C.GetHEVCConfig(c)
}
