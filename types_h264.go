package nvenc

/*
 #include "include/nvEncodeAPI.h"

 static inline void EnableRepeatSPSPPS(NV_ENC_CONFIG_H264* c) {
 	c->repeatSPSPPS = 1;
 }

 static inline void DisableSPSPPS(NV_ENC_CONFIG_H264* c) {
 	c->disableSPSPPS = 1;
 }

 static inline void EnableAUD(NV_ENC_CONFIG_H264* c) {
 	c->outputAUD = 1;
 }

 static inline void EnableIntraRefresh(NV_ENC_CONFIG_H264* c) {
 	c->enableIntraRefresh = 1;
 }

 static inline void EnableVFR(NV_ENC_CONFIG_H264* c) {
 	c->enableVFR = 1;
 }

 static inline void EnableLTR(NV_ENC_CONFIG_H264* c){
 	c->enableLTR = 1;
 }

 static inline void EnableConstrainedEncoding(NV_ENC_CONFIG_H264* c) {
 	c->enableConstrainedEncoding = 1;
 }

 static inline void UseConstrainedIntraPred(NV_ENC_CONFIG_H264* c) {
 	c->useConstrainedIntraPred = 1;
 }

 static inline void EnableOutputBufferingPeriodSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputBufferingPeriodSEI = 1;
 }

 static inline void EnableOutputPictureTimingSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputPictureTimingSEI = 1;
 }

 static inline void EnableOutputFramePackingSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputFramePackingSEI = 1;
 }

 static inline void EnableOutputRecoveryPointSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputRecoveryPointSEI = 1;
 }

 static inline void SetConstrainedFrame(NV_ENC_PIC_PARAMS_H264* p, int i) {
 	p->constrainedFrame = i;
 }
*/
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

/*
uint32_t enableTemporalSVC         :1;                          /**< [in]: Set to 1 to enable SVC temporal
uint32_t enableStereoMVC           :1;                          /**< [in]: Set to 1 to enable stereo MVC
uint32_t hierarchicalPFrames       :1;                          /**< [in]: Set to 1 to enable hierarchical PFrames
uint32_t hierarchicalBFrames       :1;                          /**< [in]: Set to 1 to enable hierarchical BFrames
uint32_t outputBufferingPeriodSEI  :1;                          /**< [in]: Set to 1 to write SEI buffering period syntax in the bitstream
uint32_t outputPictureTimingSEI    :1;                          /**< [in]: Set to 1 to write SEI picture timing syntax in the bitstream
uint32_t outputAUD                 :1;                          /**< [in]: Set to 1 to write access unit delimiter syntax in bitstream
uint32_t disableSPSPPS             :1;                          /**< [in]: Set to 1 to disable writing of Sequence and Picture parameter info in bitstream
uint32_t outputFramePackingSEI     :1;                          /**< [in]: Set to 1 to enable writing of frame packing arrangement SEI messages to bitstream
uint32_t outputRecoveryPointSEI    :1;                          /**< [in]: Set to 1 to enable writing of recovery point SEI message
uint32_t enableIntraRefresh        :1;                          /**< [in]: Set to 1 to enable gradual decoder refresh or intra refresh. If the GOP structure uses B frames this will be ignored
uint32_t enableConstrainedEncoding :1;                          /**< [in]: Set this to 1 to enable constrainedFrame encoding where each slice in the constarined picture is independent of other slices

	Check support for constrained encoding using ::NV_ENC_CAPS_SUPPORT_CONSTRAINED_ENCODING caps.

uint32_t repeatSPSPPS              :1;                          /**< [in]: Set to 1 to enable writing of Sequence and Picture parameter for every IDR frame
uint32_t enableVFR                 :1;                          /**< [in]: Set to 1 to enable variable frame rate.
uint32_t enableLTR                 :1;                          /**< [in]: Currently this feature is not available and must be set to 0. Set to 1 to enable LTR support and auto-mark the first
uint32_t qpPrimeYZeroTransformBypassFlag :1;                    /**< [in]: To enable lossless encode set this to 1, set QP to 0 and RC_mode to NV_ENC_PARAMS_RC_CONSTQP and profile to HIGH_444_PREDICTIVE_PROFILE.

	Check support for lossless encoding using ::NV_ENC_CAPS_SUPPORT_LOSSLESS_ENCODE caps.

uint32_t useConstrainedIntraPred   :1;                          /**< [in]: Set 1 to enable constrained intra prediction.
uint32_t reservedBitFields         :15;                         /**< [in]: Reserved bitfields and must be set to 0
uint32_t level;                                                 /**< [in]: Specifies the encoding level. Client is recommended to set this to NV_ENC_LEVEL_AUTOSELECT in order to enable the NvEncodeAPI interface to select the correct level.
uint32_t idrPeriod;                                             /**< [in]: Specifies the IDR interval. If not set, this is made equal to gopLength in NV_ENC_CONFIG.Low latency application client can set IDR interval to NVENC_INFINITE_GOPLENGTH so that IDR frames are not inserted automatically.
uint32_t separateColourPlaneFlag;                               /**< [in]: Set to 1 to enable 4:4:4 separate colour planes
uint32_t disableDeblockingFilterIDC;                            /**< [in]: Specifies the deblocking filter mode. Permissible value range: [0,2]
uint32_t numTemporalLayers;                                     /**< [in]: Specifies max temporal layers to be used for hierarchical coding. Valid value range is [1,::NV_ENC_CAPS_NUM_MAX_TEMPORAL_LAYERS]
uint32_t spsId;                                                 /**< [in]: Specifies the SPS id of the sequence header. Currently reserved and must be set to 0.
uint32_t ppsId;                                                 /**< [in]: Specifies the PPS id of the picture header. Currently reserved and must be set to 0.
NV_ENC_H264_ADAPTIVE_TRANSFORM_MODE adaptiveTransformMode;      /**< [in]: Specifies the AdaptiveTransform Mode. Check support for AdaptiveTransform mode using ::NV_ENC_CAPS_SUPPORT_ADAPTIVE_TRANSFORM caps.
NV_ENC_H264_FMO_MODE                fmoMode;                    /**< [in]: Specified the FMO Mode. Check support for FMO using ::NV_ENC_CAPS_SUPPORT_FMO caps.
NV_ENC_H264_BDIRECT_MODE            bdirectMode;                /**< [in]: Specifies the BDirect mode. Check support for BDirect mode using ::NV_ENC_CAPS_SUPPORT_BDIRECT_MODE caps.
NV_ENC_H264_ENTROPY_CODING_MODE     entropyCodingMode;          /**< [in]: Specifies the entropy coding mode. Check support for CABAC mode using ::NV_ENC_CAPS_SUPPORT_CABAC caps.
NV_ENC_STEREO_PACKING_MODE          stereoMode;                 /**< [in]: Specifies the stereo frame packing mode which is to be signalled in frame packing arrangement SEI
uint32_t                            intraRefreshPeriod;         /**< [in]: Specifies the interval between successive intra refresh if enableIntrarefresh is set. Requires enableIntraRefresh to be set.

	Will be disabled if NV_ENC_CONFIG::gopLength is not set to NVENC_INFINITE_GOPLENGTH.

uint32_t                            intraRefreshCnt;            /**< [in]: Specifies the length of intra refresh in number of frames for periodic intra refresh. This value should be smaller than intraRefreshPeriod
uint32_t                            maxNumRefFrames;            /**< [in]: Specifies the DPB size used for encoding. Setting it to 0 will let driver use the default dpb size.

	The low latency application which wants to invalidate reference frame as an error resilience tool
	is recommended to use a large DPB size so that the encoder can keep old reference frames which can be used if recent
	frames are invalidated.

uint32_t                            sliceMode;                  /**< [in]: This parameter in conjunction with sliceModeData specifies the way in which the picture is divided into slices

	sliceMode = 0 MB based slices, sliceMode = 1 Byte based slices, sliceMode = 2 MB row based slices, sliceMode = 3, numSlices in Picture
	When forceIntraRefreshWithFrameCnt is set it will have priority over sliceMode setting
	When sliceMode == 0 and sliceModeData == 0 whole picture will be coded with one slice

uint32_t                            sliceModeData;              /**< [in]: Specifies the parameter needed for sliceMode. For:

	sliceMode = 0, sliceModeData specifies # of MBs in each slice (except last slice)
	sliceMode = 1, sliceModeData specifies maximum # of bytes in each slice (except last slice)
	sliceMode = 2, sliceModeData specifies # of MB rows in each slice (except last slice)
	sliceMode = 3, sliceModeData specifies number of slices in the picture. Driver will divide picture into slices optimally

NV_ENC_CONFIG_H264_VUI_PARAMETERS   h264VUIParameters;          /**< [in]: Specifies the H264 video usability info pamameters
uint32_t                            ltrNumFrames;               /**< [in]: Specifies the number of LTR frames used. Additionally, encoder will mark the first numLTRFrames base layer reference frames within each IDR interval as LTR
uint32_t                            ltrTrustMode;               /**< [in]: Specifies the LTR operating mode. Set to 0 to disallow encoding using LTR frames until later specified. Set to 1 to allow encoding using LTR frames unless later invalidated.
uint32_t                            chromaFormatIDC;            /**< [in]: Specifies the chroma format. Should be set to 1 for yuv420 input, 3 for yuv444 input.

	Check support for YUV444 encoding using ::NV_ENC_CAPS_SUPPORT_YUV444_ENCODE caps.

uint32_t                            maxTemporalLayers;          /**< [in]: Specifies the max temporal layer used for hierarchical coding.
uint32_t                            reserved1[270];             /**< [in]: Reserved and must be set to 0
void*                               reserved2[64];              /**< [in]: Reserved and must be set to NULL
*/
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

/*
uint32_t displayPOCSyntax;                           /**< [in]: Specifies the display POC syntax This is required to be set if client is handling the picture type decision.
uint32_t reserved3;                                  /**< [in]: Reserved and must be set to 0
uint32_t refPicFlag;                                 /**< [in]: Set to 1 for a reference picture. This is ignored if NV_ENC_INITIALIZE_PARAMS::enablePTD is set to 1.
uint32_t colourPlaneId;                              /**< [in]: Specifies the colour plane ID associated with the current input.
uint32_t forceIntraRefreshWithFrameCnt;              /**< [in]: Forces an intra refresh with duration equal to intraRefreshFrameCnt.

	When outputRecoveryPointSEI is set this is value is used for recovery_frame_cnt in recovery point SEI message
	forceIntraRefreshWithFrameCnt cannot be used if B frames are used in the GOP structure specified

uint32_t constrainedFrame           :1;              /**< [in]: Set to 1 if client wants to encode this frame with each slice completely independent of other slices in the frame.

	NV_ENC_INITIALIZE_PARAMS::enableConstrainedEncoding should be set to 1

uint32_t sliceModeDataUpdate        :1;              /**< [in]: Set to 1 if client wants to change the sliceModeData field to specify new sliceSize Parameter

	When forceIntraRefreshWithFrameCnt is set it will have priority over sliceMode setting

uint32_t ltrMarkFrame               :1;              /**< [in]: Set to 1 if client wants to mark this frame as LTR
uint32_t ltrUseFrames               :1;              /**< [in]: Set to 1 if client allows encoding this frame using the LTR frames specified in ltrFrameBitmap
uint32_t reservedBitFields          :28;             /**< [in]: Reserved bit fields and must be set to 0
uint8_t* sliceTypeData;                              /**< [in]: Deprecated.
uint32_t sliceTypeArrayCnt;                          /**< [in]: Deprecated.
uint32_t seiPayloadArrayCnt;                         /**< [in]: Specifies the number of elements allocated in  seiPayloadArray array.
NV_ENC_SEI_PAYLOAD* seiPayloadArray;                 /**< [in]: Array of SEI payloads which will be inserted for this frame.
uint32_t sliceMode;                                  /**< [in]: This parameter in conjunction with sliceModeData specifies the way in which the picture is divided into slices

	sliceMode = 0 MB based slices, sliceMode = 1 Byte based slices, sliceMode = 2 MB row based slices, sliceMode = 3, numSlices in Picture
	When forceIntraRefreshWithFrameCnt is set it will have priority over sliceMode setting
	When sliceMode == 0 and sliceModeData == 0 whole picture will be coded with one slice

uint32_t sliceModeData;                              /**< [in]: Specifies the parameter needed for sliceMode. For:

	sliceMode = 0, sliceModeData specifies # of MBs in each slice (except last slice)
	sliceMode = 1, sliceModeData specifies maximum # of bytes in each slice (except last slice)
	sliceMode = 2, sliceModeData specifies # of MB rows in each slice (except last slice)
	sliceMode = 3, sliceModeData specifies number of slices in the picture. Driver will divide picture into slices optimally

uint32_t ltrMarkFrameIdx;                            /**< [in]: Specifies the long term referenceframe index to use for marking this frame as LTR.
uint32_t ltrUseFrameBitmap;                          /**< [in]: Specifies the the associated bitmap of LTR frame indices when encoding this frame.
uint32_t ltrUsageMode;                               /**< [in]: Specifies additional usage constraints for encoding using LTR frames from this point further. 0: no constraints, 1: no short term refs older than current, no previous LTR frames.
uint32_t reserved [243];                             /**< [in]: Reserved and must be set to 0.
void*    reserved2[62];                              /**< [in]: Reserved and must be set to NULL.
*/
type PIC_PARAMS_H264 = C.NV_ENC_PIC_PARAMS_H264

func (p *PIC_PARAMS_H264) ForceIntraRefresh(frames uint32) {
	p.forceIntraRefreshWithFrameCnt = C.uint32_t(frames)
}

func (p *PIC_PARAMS_H264) SetConstrainedFrame(enable int) {
	C.SetConstrainedFrame(p, C.int(enable))
}
