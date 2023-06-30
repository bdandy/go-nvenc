package types

// #cgo CFLAGS: -I ../../include
// #include "types.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// CreateInputBuffer is struct with settings for input buffer
type CreateInputBuffer C.NV_ENC_CREATE_INPUT_BUFFER

// GetBufferPtr returns pointer to the buffer
func (p CreateInputBuffer) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(uintptr(p.inputBuffer))
}

// SetResolution sets target resolution for input buffer
func (p *CreateInputBuffer) SetResolution(width, height uint32) {
	p.width = C.uint32_t(width)
	p.height = C.uint32_t(height)
}

// SetFormat sets target format for input buffer
func (p *CreateInputBuffer) SetFormat(format BufferFormat) {
	p.bufferFmt = C.NV_ENC_BUFFER_FORMAT(format)
}

func (p CreateInputBuffer) CType() C.NV_ENC_CREATE_INPUT_BUFFER {
	return C.NV_ENC_CREATE_INPUT_BUFFER(p)
}

func NewCreateInputBuffer() *CreateInputBuffer {
	var params CreateInputBuffer
	params.version = C.NV_ENC_CREATE_INPUT_BUFFER_VER
	return &params
}

// BitstreamBuffer is struct contains pointer to bitstream buffer
type BitstreamBuffer C.NV_ENC_CREATE_BITSTREAM_BUFFER

func (b *BitstreamBuffer) GetBufferPtr() unsafe.Pointer {
	return unsafe.Pointer(b.bitstreamBuffer)
}

func NewBitstreamBuffer(size uint32) *BitstreamBuffer {
	buffer := new(BitstreamBuffer)
	buffer.version = C.NV_ENC_CREATE_BITSTREAM_BUFFER_VER
	buffer.size = C.uint32_t(size)
	return buffer
}

// LockBitstreamParams is a struct with params for locking bitstream buffer
type LockBitstreamParams C.NV_ENC_LOCK_BITSTREAM

func (b LockBitstreamParams) BitstreamSize() int {
	return int(b.bitstreamSizeInBytes)
}

func (b *LockBitstreamParams) MemcpyBitstream(buf []byte) error {
	if len(buf) < b.BitstreamSize() {
		return fmt.Errorf("bufSize is %d bytes, but bitstream need %d bytes for output", len(buf), b.BitstreamSize())
	}

	C.memcpy((unsafe.Pointer)(&buf[0]), b.bitstreamBufferPtr, C.size_t(b.BitstreamSize()))

	return nil
}

func (b *LockBitstreamParams) CopyBitstream(buf []byte) error {
	if len(buf) < b.BitstreamSize() {
		return fmt.Errorf("bufSize is %d bytes, but bitstream need %d bytes for output", len(buf), b.BitstreamSize())
	}

	copy(unsafe.Slice((*byte)(b.bitstreamBufferPtr), b.BitstreamSize()), buf)

	return nil
}

func (b *LockBitstreamParams) FrameId() uint32 {
	return uint32(b.frameIdx)
}

func (b *LockBitstreamParams) FrameTimeStamp() uint64 {
	return uint64(b.outputTimeStamp)
}

func NewBitstreamBufferLock(buf *BitstreamBuffer) *LockBitstreamParams {
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
func (p LockInputBufferParams) Pitch() uint32 {
	return uint32(p.pitch)
}

func (p LockInputBufferParams) buf() []byte {
	return unsafe.Slice((*byte)(p.bufferDataPtr), p.pitch)
}

// CopyBuffer copies data into buf.
func (p *LockInputBufferParams) CopyBuffer(buf []byte) error {
	if uint32(len(buf)) < p.Pitch() {
		return fmt.Errorf("buffer is too small")
	}

	copy(p.buf(), buf)

	return nil
}

// TODO: Async mode
func NewLockInputBufferParams(b *CreateInputBuffer) *LockInputBufferParams {
	params := new(LockInputBufferParams)
	params.version = C.NV_ENC_LOCK_INPUT_BUFFER_VER
	params.inputBuffer = b.inputBuffer

	return params
}
