package nvenc

// #include "headers/nvEncodeAPI.h"
import "C"
import "fmt"

var (
	BufferFormatNV12   = bufferFormat(0x1)       /**< Semi-Planar YUV [Y plane followed by interleaved UV plane] */
	BufferFormatYV12   = bufferFormat(0x10)      /**< Planar YUV [Y plane followed by V and U planes] */
	BufferFormatIYUV   = bufferFormat(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BufferFormatI420   = bufferFormat(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BufferFormatYUV420 = bufferFormat(0x100)     /**< Planar YUV [Y plane followed by U and V planes] */
	BufferFormatYUV444 = bufferFormat(0x1000)    /**< Planar YUV [Y plane followed by U and V planes] */
	BufferFormatARGB   = bufferFormat(0x1000000) /**< 8 bit Packed A8R8G8B8 */
	BufferFormatARGB10 = bufferFormat(0x2000000) /**< 10 bit Packed A2R10G10B10 */
	BufferFormatAYUV   = bufferFormat(0x4000000) /**< 8 bit Packed A8Y8U8V8 */
)

type bufferFormat = C.NV_ENC_BUFFER_FORMAT

func (format bufferFormat) String() string {
	switch format {
	case BufferFormatNV12:
		return "NV12"
	case BufferFormatYV12:
		return "YV12"
	case BufferFormatIYUV:
		return "IYUV"
	case BufferFormatI420:
		return "I420"
	case BufferFormatYUV420:
		return "YUV420"
	case BufferFormatYUV444:
		return "YUV444"
	case BufferFormatARGB:
		return "ARGB"
	case BufferFormatARGB10:
		return "ARGB10"
	case BufferFormatAYUV:
		return "AYUV"
	}

	return fmt.Sprintf("unknown:%x", int(format))
}
