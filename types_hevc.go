package nvenc

/*
 #include "include/nvEncodeAPI.h"

 static inline void HEVCSetRepeatSPSPPS(NV_ENC_CONFIG_HEVC* c) {
 	c->repeatSPSPPS = 1;
 }

 static inline void HEVCSetDisableSPSPPS(NV_ENC_CONFIG_HEVC* c) {
 	c->disableSPSPPS = 1;
 }

 static inline void HEVCSetEnableAUD(NV_ENC_CONFIG_HEVC* c) {
 	c->outputAUD = 1;
 }

 static inline void HEVCSetEnableIntraRefresh(NV_ENC_CONFIG_HEVC* c) {
 	c->enableIntraRefresh = 1;
 }
*/
import "C"

/*
uint32_t displayPOCSyntax;                           /**< [in]: Specifies the display POC syntax This is required to be set if client is handling the picture type decision.
uint32_t refPicFlag;                                 /**< [in]: Set to 1 for a reference picture. This is ignored if NV_ENC_INITIALIZE_PARAMS::enablePTD is set to 1.
uint32_t temporalId;                                 /**< [in]: Specifies the temporal id of the picture
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
uint8_t* sliceTypeData;                              /**< [in]: Array which specifies the slice type used to force intra slice for a particular slice. Currently supported only for NV_ENC_CONFIG_H264::sliceMode == 3.

	Client should allocate array of size sliceModeData where sliceModeData is specified in field of ::_NV_ENC_CONFIG_H264
	Array element with index n corresponds to nth slice. To force a particular slice to intra client should set corresponding array element to NV_ENC_SLICE_TYPE_I
	all other array elements should be set to NV_ENC_SLICE_TYPE_DEFAULT

uint32_t sliceTypeArrayCnt;                          /**< [in]: Client should set this to the number of elements allocated in sliceTypeData array. If sliceTypeData is NULL then this should be set to 0
uint32_t sliceMode;                                  /**< [in]: This parameter in conjunction with sliceModeData specifies the way in which the picture is divided into slices

	sliceMode = 0 CTU based slices, sliceMode = 1 Byte based slices, sliceMode = 2 CTU row based slices, sliceMode = 3, numSlices in Picture
	When forceIntraRefreshWithFrameCnt is set it will have priority over sliceMode setting
	When sliceMode == 0 and sliceModeData == 0 whole picture will be coded with one slice

uint32_t sliceModeData;                              /**< [in]: Specifies the parameter needed for sliceMode. For:

	sliceMode = 0, sliceModeData specifies # of CTUs in each slice (except last slice)
	sliceMode = 1, sliceModeData specifies maximum # of bytes in each slice (except last slice)
	sliceMode = 2, sliceModeData specifies # of CTU rows in each slice (except last slice)
	sliceMode = 3, sliceModeData specifies number of slices in the picture. Driver will divide picture into slices optimally

uint32_t ltrMarkFrameIdx;                            /**< [in]: Specifies the long term reference frame index to use for marking this frame as LTR.
uint32_t ltrUseFrameBitmap;                          /**< [in]: Specifies the associated bitmap of LTR frame indices when encoding this frame.
uint32_t ltrUsageMode;                               /**< [in]: Specifies additional usage constraints for encoding using LTR frames from this point further. 0: no constraints, 1: no short term refs older than current, no previous LTR frames.
uint32_t seiPayloadArrayCnt;                         /**< [in]: Specifies the number of elements allocated in  seiPayloadArray array.
uint32_t reserved;                                   /**< [in]: Reserved and must be set to 0.
NV_ENC_SEI_PAYLOAD* seiPayloadArray;                 /**< [in]: Array of SEI payloads which will be inserted for this frame.
uint32_t reserved2 [244];                             /**< [in]: Reserved and must be set to 0.
void*    reserved3[61];                              /**< [in]: Reserved and must be set to NULL.
*/
type PIC_PARAMS_HEVC = C.NV_ENC_PIC_PARAMS_HEVC

/*
uint32_t level;                                                 /**< [in]: Specifies the level of the encoded bitstream.
uint32_t tier;                                                  /**< [in]: Specifies the level tier of the encoded bitstream.
NV_ENC_HEVC_CUSIZE minCUSize;                                   /**< [in]: Specifies the minimum size of luma coding unit.
NV_ENC_HEVC_CUSIZE maxCUSize;                                   /**< [in]: Specifies the maximum size of luma coding unit. Currently NVENC SDK only supports maxCUSize equal to NV_ENC_HEVC_CUSIZE_32x32.
uint32_t useConstrainedIntraPred               :1;              /**< [in]: Set 1 to enable constrained intra prediction.
uint32_t disableDeblockAcrossSliceBoundary     :1;              /**< [in]: Set 1 to disable in loop filtering across slice boundary.
uint32_t outputBufferingPeriodSEI              :1;              /**< [in]: Set 1 to write SEI buffering period syntax in the bitstream
uint32_t outputPictureTimingSEI                :1;              /**< [in]: Set 1 to write SEI picture timing syntax in the bitstream
uint32_t outputAUD                             :1;              /**< [in]: Set 1 to write Access Unit Delimiter syntax.
uint32_t enableLTR                             :1;              /**< [in]: Set 1 to enable use of long term reference pictures for inter prediction.
uint32_t disableSPSPPS                         :1;              /**< [in]: Set 1 to disable VPS,SPS and PPS signalling in the bitstream.
uint32_t repeatSPSPPS                          :1;              /**< [in]: Set 1 to output VPS,SPS and PPS for every IDR frame.
uint32_t enableIntraRefresh                    :1;              /**< [in]: Set 1 to enable gradual decoder refresh or intra refresh. If the GOP structure uses B frames this will be ignored
uint32_t chromaFormatIDC                       :2;              /**< [in]: Specifies the chroma format. Should be set to 1 for yuv420 input, 3 for yuv444 input.
uint32_t reserved                              :21;             /**< [in]: Reserved bitfields.
uint32_t idrPeriod;                                             /**< [in]: Specifies the IDR interval. If not set, this is made equal to gopLength in NV_ENC_CONFIG.Low latency application client can set IDR interval to NVENC_INFINITE_GOPLENGTH so that IDR frames are not inserted automatically.
uint32_t intraRefreshPeriod;                                    /**< [in]: Specifies the interval between successive intra refresh if enableIntrarefresh is set. Requires enableIntraRefresh to be set.
                                                                    Will be disabled if NV_ENC_CONFIG::gopLength is not set to NVENC_INFINITE_GOPLENGTH.
uint32_t intraRefreshCnt;                                       /**< [in]: Specifies the length of intra refresh in number of frames for periodic intra refresh. This value should be smaller than intraRefreshPeriod
uint32_t maxNumRefFramesInDPB;                                  /**< [in]: Specifies the maximum number of references frames in the DPB.
uint32_t ltrNumFrames;                                          /**< [in]: Specifies the maximum number of long term references can be used for prediction
uint32_t vpsId;                                                 /**< [in]: Specifies the VPS id of the video parameter set. Currently reserved and must be set to 0.
uint32_t spsId;                                                 /**< [in]: Specifies the SPS id of the sequence header. Currently reserved and must be set to 0.
uint32_t ppsId;                                                 /**< [in]: Specifies the PPS id of the picture header. Currently reserved and must be set to 0.
uint32_t sliceMode;                                             /**< [in]: This parameter in conjunction with sliceModeData specifies the way in which the picture is divided into slices
                                                                                sliceMode = 0 CTU based slices, sliceMode = 1 Byte based slices, sliceMode = 2 CTU row based slices, sliceMode = 3, numSlices in Picture
                                                                                When sliceMode == 0 and sliceModeData == 0 whole picture will be coded with one slice
uint32_t sliceModeData;                                         /**< [in]: Specifies the parameter needed for sliceMode. For:
                                                                                sliceMode = 0, sliceModeData specifies # of CTUs in each slice (except last slice)
                                                                                sliceMode = 1, sliceModeData specifies maximum # of bytes in each slice (except last slice)
                                                                                sliceMode = 2, sliceModeData specifies # of CTU rows in each slice (except last slice)
                                                                                sliceMode = 3, sliceModeData specifies number of slices in the picture. Driver will divide picture into slices optimally
uint32_t maxTemporalLayersMinus1;                               /**< [in]: Specifies the max temporal layer used for hierarchical coding.
NV_ENC_CONFIG_HEVC_VUI_PARAMETERS   hevcVUIParameters;          /**< [in]: Specifies the HEVC video usability info pamameters
uint32_t reserved1[218];                                        /**< [in]: Reserved and must be set to 0.
void*    reserved2[64];                                         /**< [in]: Reserved and must be set to NULL
*/

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
