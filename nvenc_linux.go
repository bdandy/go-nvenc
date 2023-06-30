package nvenc

// #cgo CFLAGS: -I ./include
// #include "nvenc_linux.h"
import "C"

import (
	"unsafe"

	"github.com/rainycape/dl"
)

var (
	libEncode      *dl.DL
	createInstance unsafe.Pointer
)

func init() {
	var err error
	libEncode, err = dl.Open("libnvidia-encode", dl.RTLD_LAZY)
	if err != nil {
		panic(err)
	}

	if err = libEncode.Sym("NvEncodeAPICreateInstance", &createInstance); err != nil {
		panic(err)
	}
}

func callCreateInstance(functions *EncoderFunctions) error {
	ret := C.callCreateInstance((*[0]byte)(createInstance), (*C.NV_ENCODE_API_FUNCTION_LIST)(unsafe.Pointer(functions)))
	return codeToError(C.int(ret))
}
