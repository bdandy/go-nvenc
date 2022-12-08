package d3d11

import "C"
import (
	"fmt"
)

// #cgo LDFLAGS: -ld3d11 -ldxgi
// #include "get_adapter.h"
import "C"

type Device = C.ID3D11Device

func CreateDevice() (*Device, error) {
	var device *Device

	adapter := C.GetNvidiaAdapter()
	if adapter == nil {
		return nil, fmt.Errorf("No NVIDIA adapters was found")
	}

	ret := C.D3D11CreateDevice(adapter, C.D3D_DRIVER_TYPE_UNKNOWN, nil, 0, nil, 0, C.D3D11_SDK_VERSION, &device, nil, nil)
	if ret != C.S_OK {
		return nil, fmt.Errorf("D3D11CreateDevice error code: %d", int(ret))
	}

	return device, nil
}
