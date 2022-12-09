package nvenc

// #include "include/types.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// CreateInputBuffer is struct with settings for input buffer
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
type CreateInputBuffer C.NV_ENC_CREATE_INPUT_BUFFER

func (p *CreateInputBuffer) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(uintptr(p.inputBuffer))
}

func (p *CreateInputBuffer) SetResolution(width, height uint32) {
	p.width = C.uint32_t(width)
	p.height = C.uint32_t(height)
}

func (p *CreateInputBuffer) SetFormat(format bufferFormat) {
	p.bufferFmt = C.NV_ENC_BUFFER_FORMAT(format)
}

func newCreateInputBuffer() *CreateInputBuffer {
	params := new(CreateInputBuffer)
	params.version = C.NV_ENC_CREATE_INPUT_BUFFER_VER
	return params
}

// BitstreamBuffer is struct contains pointer to bitstream buffer
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
type BitstreamBuffer C.NV_ENC_CREATE_BITSTREAM_BUFFER

func (b *BitstreamBuffer) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(b.bitstreamBuffer)
}

func newBitstreamBuffer(size uint32) *BitstreamBuffer {
	buffer := new(BitstreamBuffer)
	buffer.version = C.NV_ENC_CREATE_BITSTREAM_BUFFER_VER
	buffer.size = C.uint32_t(size)
	return buffer
}

// LockBitstreamParams is a struct with params for locking bitstream buffer
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
type LockBitstreamParams C.NV_ENC_LOCK_BITSTREAM

func (b *LockBitstreamParams) BitstreamSize() int {
	return int(b.bitstreamSizeInBytes)
}

func (b *LockBitstreamParams) CopyBitstream(buf []byte) error {
	if len(buf) < b.BitstreamSize() {
		return fmt.Errorf("bufSize is %d bytes, but bitstream need %d bytes for output", len(buf), b.BitstreamSize())
	}

	C.memcpy((unsafe.Pointer)(&buf[0]), b.bitstreamBufferPtr, C.size_t(b.BitstreamSize()))

	return nil
}

func (b *LockBitstreamParams) FrameId() uint32 {
	return uint32(b.frameIdx)
}

func (b *LockBitstreamParams) FrameTimeStamp() uint64 {
	return uint64(b.outputTimeStamp)
}

func newBitstreamBufferLock(buf *BitstreamBuffer) *LockBitstreamParams {
	params := new(LockBitstreamParams)
	params.version = C.NV_ENC_LOCK_BITSTREAM_VER
	params.outputBitstream = unsafe.Pointer(buf.GetBufferPtr())
	return params
}

// LockInputBufferParams
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
type LockInputBufferParams C.NV_ENC_LOCK_INPUT_BUFFER

// Pitch returns pitch of input buffer
func (p *LockInputBufferParams) Pitch() uint32 {
	return uint32(p.pitch)
}

// CopyBuffer copies data into buf.
func (p *LockInputBufferParams) CopyBuffer(buf []byte) error {
	if uint32(len(buf)) < p.Pitch() {
		return fmt.Errorf("buffer is too small")
	}

	C.memcpy((unsafe.Pointer)(&buf[0]), p.bufferDataPtr, C.size_t(p.pitch))

	return nil
}

// TODO: Async mode
func newLockInputBufferParams(b *CreateInputBuffer) *LockInputBufferParams {
	params := new(LockInputBufferParams)
	params.version = C.NV_ENC_LOCK_INPUT_BUFFER_VER
	params.inputBuffer = b.inputBuffer

	return params
}
