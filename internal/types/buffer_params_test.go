package types

import (
	"testing"
	"unsafe"
)

func Benchmark_LockBitstreamParams_Memcpy(b *testing.B) {
	var data [1024]byte
	var buf [1024]byte

	var params LockBitstreamParams
	params.bitstreamBufferPtr = unsafe.Pointer(&buf[0])
	params.bitstreamSizeInBytes = 1024
	for i := 0; i < b.N; i++ {
		params.MemcpyBitstream(data[:])
	}
}

func Benchmark_LockBitstreamParams_Copy(b *testing.B) {
	var data [1024]byte
	var buf [1024]byte

	var params LockBitstreamParams
	params.bitstreamBufferPtr = unsafe.Pointer(&buf[0])
	params.bitstreamSizeInBytes = 1024
	for i := 0; i < b.N; i++ {
		params.CopyBitstream(data[:])
	}
}
