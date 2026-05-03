package astiav

//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html
type HardwareFramesContext struct {
	c *C.struct_AVBufferRef
}

func newHardwareFramesContextFromC(c *C.struct_AVBufferRef) *HardwareFramesContext {
	if c == nil {
		return nil
	}
	return &HardwareFramesContext{c: c}
}

// https://ffmpeg.org/doxygen/8.0/hwcontext_8c.html#ac45a7c039eb4e084b692f69ff5f2e217
func AllocHardwareFramesContext(hdc *HardwareDeviceContext) *HardwareFramesContext {
	return newHardwareFramesContextFromC(C.av_hwframe_ctx_alloc(hdc.c))
}

func (hfc *HardwareFramesContext) Free() {
	if hfc.c != nil {
		C.av_buffer_unref(&hfc.c)
	}
}

// data returns the underlying AVHWFramesContext, or nil if the buffer
// reference has not been allocated (e.g. zero-value receiver) or its
// data pointer is unset (allocator failed, refcount dropped to zero,
// etc.). Callers that read fields must nil-check the result; setters
// rely on this returning non-nil when the receiver was constructed via
// AllocHardwareFramesContext.
func (hfc *HardwareFramesContext) data() *C.AVHWFramesContext {
	if hfc == nil || hfc.c == nil || hfc.c.data == nil {
		return nil
	}
	return (*C.AVHWFramesContext)(unsafe.Pointer(hfc.c.data))
}

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a9e6f29d0f744930cdd0e8bdff8771520
func (hfc *HardwareFramesContext) SetWidth(width int) {
	hfc.data().width = C.int(width)
}

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#ae61bbe1d8645a0c573085e29f1d0a58f
func (hfc *HardwareFramesContext) SetHeight(height int) {
	hfc.data().height = C.int(height)
}

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a045bc1713932804f6ceef170a5578e0e
func (hfc *HardwareFramesContext) SetHardwarePixelFormat(format PixelFormat) {
	hfc.data().format = C.enum_AVPixelFormat(format)
}

// HardwarePixelFormat returns the configured hardware pixel format
// (e.g. AV_PIX_FMT_MEDIACODEC, AV_PIX_FMT_CUDA). Returns PixelFormatNone
// when called before AllocHardwareFramesContext (or with a zero-value
// receiver) so misordered code surfaces a benign sentinel rather than
// a SIGSEGV in C.
//
// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a045bc1713932804f6ceef170a5578e0e
func (hfc *HardwareFramesContext) HardwarePixelFormat() PixelFormat {
	d := hfc.data()
	if d == nil {
		return PixelFormatNone
	}
	return PixelFormat(d.format)
}

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a663a9aceca97aa7b2426c9aba6543e4a
func (hfc *HardwareFramesContext) SetSoftwarePixelFormat(swFormat PixelFormat) {
	hfc.data().sw_format = C.enum_AVPixelFormat(swFormat)
}

// SoftwarePixelFormat returns the configured software pixel format used
// for SW<->HW transfers (e.g. AV_PIX_FMT_NV12 for mediacodec/cuda).
// Returns PixelFormatNone when called before AllocHardwareFramesContext
// (or with a zero-value receiver); see HardwarePixelFormat.
//
// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a663a9aceca97aa7b2426c9aba6543e4a
func (hfc *HardwareFramesContext) SoftwarePixelFormat() PixelFormat {
	d := hfc.data()
	if d == nil {
		return PixelFormatNone
	}
	return PixelFormat(d.sw_format)
}

// https://ffmpeg.org/doxygen/8.0/structAVHWFramesContext.html#a9c3a94dcd9c96e19059b56a6bae9c764
func (hfc *HardwareFramesContext) SetInitialPoolSize(initialPoolSize int) {
	hfc.data().initial_pool_size = C.int(initialPoolSize)
}

// https://ffmpeg.org/doxygen/8.0/hwcontext_8c.html#a66a7e1ebc7e459ce07d3de6639ac7e38
func (hfc *HardwareFramesContext) Initialize() error {
	return newError(C.av_hwframe_ctx_init(hfc.c))
}
