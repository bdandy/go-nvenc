package nvenc

import "C"
import (
	"unsafe"

	"github.com/rainycape/dl"
)

// #include "include/nvEncodeAPI.h"
/*
	typedef NVENCSTATUS (NVENCAPI* PNVENCODEAPICREATEINSTANCE) (NV_ENCODE_API_FUNCTION_LIST *functionList);

	static inline NVENCSTATUS callCreateInstance(PNVENCODEAPICREATEINSTANCE fn, void* list) {
		return fn(list);
	}
*/
import "C"

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
	ret := C.callCreateInstance((*[0]byte)(createInstance), unsafe.Pointer(functions))
	return codeToError(C.int(ret))
}
