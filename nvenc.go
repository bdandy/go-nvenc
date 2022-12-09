package nvenc

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"gorgonia.org/cu"
)

func (e *Encoder) OpenCUDAEncodeSession() error {
	count, err := cu.NumDevices()
	if err != nil {
		return fmt.Errorf("cuda numerate devices: %w", err)
	}
	for i := 0; i < count; i++ {
		device, err := cu.GetDevice(i)
		if err != nil {
			return fmt.Errorf("cuda get device %d: %w", i, err)
		}
		if !device.IsGPU() {
			continue
		}
		name, err := device.Name()
		if err != nil {
			return fmt.Errorf("cuda get device name: %w", err)
		}
		if !strings.Contains(name, "NVIDIA") {
			continue
		}

		ctx, err := device.MakeContext(0)
		if err != nil {
			return err
		}
		_, _ = cu.PopCurrentCtx()

		addr, _ := strconv.ParseUint(strings.TrimPrefix(ctx.String(), "0x"), 16, 0)

		params := newOpenEncodeSessionParams(DeviceTypeCuda, unsafe.Pointer(uintptr(addr)))
		e.instance, err = e.functions.openEncodeSessionEx(params)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("No NVIDIA device found")
}
