package nvenc

// #include "headers/types.h"
import "C"
import "unsafe"

// EncoderPictureParams is picture params for encoding
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
type EncoderPictureParams C.NV_ENC_PIC_PARAMS

// SetTimestamp sets timestamp of current frame
func (p *EncoderPictureParams) SetTimestamp(t uint64) {
	p.inputTimeStamp = C.uint64_t(t)
}

func (p *EncoderPictureParams) SetH264Params(param *PIC_PARAMS_H264) {
	cast := (*C.NV_ENC_CODEC_PIC_PARAMS)(unsafe.Pointer(param))
	p.codecPicParams = C.NV_ENC_CODEC_PIC_PARAMS(*cast)
}

func (p *EncoderPictureParams) SetHEVCParams(param *PIC_PARAMS_HEVC) {
	cast := (*C.NV_ENC_CODEC_PIC_PARAMS)(unsafe.Pointer(param))
	p.codecPicParams = C.NV_ENC_CODEC_PIC_PARAMS(*cast)
}

func (p *EncoderPictureParams) SetInputBuffer(ptr unsafe.Pointer) {
	p.inputBuffer = C.NV_ENC_INPUT_PTR(ptr)
}

func (p *EncoderPictureParams) SetOutputBuffer(ptr unsafe.Pointer) {
	p.outputBitstream = C.NV_ENC_OUTPUT_PTR(ptr)
}

// ForceIntra tells encoder to force this frame as an Intra picture
func (p *EncoderPictureParams) ForceIntra(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_FORCEINTRA
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_FORCEINTRA
	}
}

// ForceIDR tells encoder to encode the current picture as an IDR picture.
// This flag is only valid when Picture type decision is taken by the Encoder.
// See InitializeParams.SetEnablePTD()
func (p *EncoderPictureParams) ForceIDR(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_FORCEIDR
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_FORCEIDR
	}
}

// ForceSPSPPS tells encoder to write the sequence and picture header in encoded bitstream of the current picture
func (p *EncoderPictureParams) ForceSPSPPS(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_OUTPUT_SPSPPS
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_OUTPUT_SPSPPS
	}
}

// ForceEOS tells encoder if it's end of the input stream
func (p *EncoderPictureParams) ForceEOS(b bool) {
	if b {
		p.encodePicFlags |= C.NV_ENC_PIC_FLAG_EOS
	} else {
		p.encodePicFlags &^= C.NV_ENC_PIC_FLAG_EOS
	}
}

// SetResolution tells encoder frame resolution. If optional pitch is not set, pitch is equal to width
func (p *EncoderPictureParams) SetResolution(width, height uint32, pitch ...uint32) {
	p.inputWidth = C.uint32_t(width)
	p.inputHeight = C.uint32_t(height)

	if len(pitch) == 0 {
		p.inputPitch = C.uint32_t(width)
	} else {
		p.inputPitch = C.uint32_t(pitch[0])
	}
}

// SetInputFormat tells encoder in which format is input
func (p *EncoderPictureParams) SetInputFormat(format bufferFormat) {
	p.bufferFmt = C.NV_ENC_BUFFER_FORMAT(format)
}

// PicParamsH264 returns H264 params struct
func (p *EncoderPictureParams) PicParamsH264() *PIC_PARAMS_H264 {
	return (*PIC_PARAMS_H264)(C.GetPicParamsH264(&p.codecPicParams))
}

func newEncPicParams() *EncoderPictureParams {
	params := new(EncoderPictureParams)
	params.version = C.NV_ENC_PIC_PARAMS_VER
	params.pictureStruct = C.NV_ENC_PIC_STRUCT_FRAME

	return params
}
