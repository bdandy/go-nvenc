package nvenc

import (
	"unsafe"

	"github.com/bdandy/go-nvenc/v8/internal/types"
)

func (e *Encoder) OpenCUDAEncodeSession(device unsafe.Pointer) error {
	var err error

	params := types.NewOpenEncodeSessionParams(types.DeviceTypeCuda, device)
	e.instance, err = e.functions.openEncodeSessionEx(params)

	return err
}
