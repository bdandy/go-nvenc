package nvenc

import "C"
import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/bdandy/go-nvenc/v8/d3d11"
)

var (
	dll            *syscall.LazyDLL
	createInstance *syscall.LazyProc
)

func init() {
	switch runtime.GOARCH {
	case "amd64":
		dll = syscall.NewLazyDLL("nvEncodeAPI64.dll")
	default:
		dll = syscall.NewLazyDLL("nvEncodeAPI.dll")
	}

	createInstance = dll.NewProc("NvEncodeAPICreateInstance")
}

func callCreateInstance(functions *EncoderFunctions) error {
	ret, _, _ := createInstance.Call(uintptr(unsafe.Pointer(functions)))
	err := codeToError(C.int(ret))
	return err
}

func (e *Encoder) OpenDX11EncodeSession() error {
	device, err := d3d11.CreateDevice()
	if err != nil {
		return err
	}

	params := newOpenEncodeSessionParams(DEVICE_TYPE_DIRECTX, unsafe.Pointer(device))
	e.instance, err = e.functions.openEncodeSessionEx(params)
	if err != nil {
		return err
	}

	return nil
}
