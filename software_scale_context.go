package astiav

//#include <libswscale/swscale.h>
import "C"
import (
	"errors"
	"unsafe"
)

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html
type SoftwareScaleContext struct {
	c *C.struct_SwsContext
	// Store attributes in Go since SwsContext is opaque in some builds.
	dstFormat C.enum_AVPixelFormat
	dstH      C.int
	dstW      C.int
	flags     C.int
	srcFormat C.enum_AVPixelFormat
	srcH      C.int
	srcW      C.int
}

type softwareScaleContextUpdate struct {
	dstFormat *PixelFormat
	dstH      *int
	dstW      *int
	flags     *SoftwareScaleContextFlags
	srcFormat *PixelFormat
	srcH      *int
	srcW      *int
}

func newSoftwareScaleContextFromC(c *C.struct_SwsContext) *SoftwareScaleContext {
	if c == nil {
		return nil
	}
	return &SoftwareScaleContext{c: c}
}

// https://ffmpeg.org/doxygen/8.0/group__libsws.html#ga59cc19eff0434e7ec11676dc5e222ff3
func CreateSoftwareScaleContext(srcW, srcH int, srcFormat PixelFormat, dstW, dstH int, dstFormat PixelFormat, flags SoftwareScaleContextFlags) (*SoftwareScaleContext, error) {
	ssc := &SoftwareScaleContext{
		dstFormat: C.enum_AVPixelFormat(dstFormat),
		dstH:      C.int(dstH),
		dstW:      C.int(dstW),
		flags:     C.int(flags),
		srcFormat: C.enum_AVPixelFormat(srcFormat),
		srcH:      C.int(srcH),
		srcW:      C.int(srcW),
	}

	ssc.c = C.sws_getContext(
		ssc.srcW,
		ssc.srcH,
		ssc.srcFormat,
		ssc.dstW,
		ssc.dstH,
		ssc.dstFormat,
		ssc.flags,
		nil, nil, nil,
	)
	if ssc.c == nil {
		return nil, errors.New("astiav: empty new context")
	}

	classers.set(ssc)
	return ssc, nil
}

// https://ffmpeg.org/doxygen/8.0/group__libsws.html#gad90b463ceeafdfd526994742f9954dbb
func (ssc *SoftwareScaleContext) Free() {
	if ssc.c != nil {
		// Make sure to clone the classer before freeing the object since
		// the C free method may reset the pointer
		c := newClonedClasser(ssc)
		C.sws_freeContext(ssc.c)
		ssc.c = nil
		// Make sure to remove from classers after freeing the object since
		// the C free method may use methods needing the classer
		if c != nil {
			classers.del(c)
		}
	}
}

var _ Classer = (*SoftwareScaleContext)(nil)

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a6866f52574bc730833d2580abc806261
func (ssc *SoftwareScaleContext) Class() *Class {
	if ssc.c == nil {
		return nil
	}
	return newClassFromC(unsafe.Pointer(ssc.c))
}

// https://ffmpeg.org/doxygen/8.0/swscale-v2_8txt.html#a20ffff3ac1378332422b93ed70264f4c
func (ssc *SoftwareScaleContext) ScaleFrame(src, dst *Frame) error {
	return newError(C.sws_scale_frame(ssc.c, dst.c, src.c))
}

// https://ffmpeg.org/doxygen/8.0/group__libsws.html#ga9fd74ceaf0f126f762b81e3677f70c75
func (ssc *SoftwareScaleContext) update(u softwareScaleContextUpdate) error {
	dstW := ssc.dstW
	if u.dstW != nil {
		dstW = C.int(*u.dstW)
	}

	dstH := ssc.dstH
	if u.dstH != nil {
		dstH = C.int(*u.dstH)
	}

	dstFormat := ssc.dstFormat
	if u.dstFormat != nil {
		dstFormat = C.enum_AVPixelFormat(*u.dstFormat)
	}

	srcW := ssc.srcW
	if u.srcW != nil {
		srcW = C.int(*u.srcW)
	}

	srcH := ssc.srcH
	if u.srcH != nil {
		srcH = C.int(*u.srcH)
	}

	srcFormat := ssc.srcFormat
	if u.srcFormat != nil {
		srcFormat = C.enum_AVPixelFormat(*u.srcFormat)
	}

	flags := ssc.flags
	if u.flags != nil {
		flags = C.int(*u.flags)
	}

	c := C.sws_getCachedContext(
		ssc.c,
		srcW,
		srcH,
		srcFormat,
		dstW,
		dstH,
		dstFormat,
		flags,
		nil, nil, nil,
	)
	if c == nil {
		return errors.New("astiav: empty new context")
	}

	ssc.c = c
	ssc.dstW = dstW
	ssc.dstH = dstH
	ssc.dstFormat = dstFormat
	ssc.srcW = srcW
	ssc.srcH = srcH
	ssc.srcFormat = srcFormat
	ssc.flags = flags

	return nil
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#aef45de443b59978fd38ad1531c618574
func (ssc *SoftwareScaleContext) Flags() SoftwareScaleContextFlags {
	return SoftwareScaleContextFlags(ssc.flags)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#aef45de443b59978fd38ad1531c618574
func (ssc *SoftwareScaleContext) SetFlags(swscf SoftwareScaleContextFlags) error {
	return ssc.update(softwareScaleContextUpdate{flags: &swscf})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a883a891c8a2d4ea7a15a3a7055f64386
func (ssc *SoftwareScaleContext) DestinationWidth() int {
	return int(ssc.dstW)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a883a891c8a2d4ea7a15a3a7055f64386
func (ssc *SoftwareScaleContext) SetDestinationWidth(i int) error {
	return ssc.update(softwareScaleContextUpdate{dstW: &i})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a7facd34608c9258dae8c2942e3dce78f
func (ssc *SoftwareScaleContext) DestinationHeight() int {
	return int(ssc.dstH)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a7facd34608c9258dae8c2942e3dce78f
func (ssc *SoftwareScaleContext) SetDestinationHeight(i int) error {
	return ssc.update(softwareScaleContextUpdate{dstH: &i})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a0ff71c9ef5ab6dabf90378fa7bf836ec
func (ssc *SoftwareScaleContext) DestinationPixelFormat() PixelFormat {
	return PixelFormat(ssc.dstFormat)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a0ff71c9ef5ab6dabf90378fa7bf836ec
func (ssc *SoftwareScaleContext) SetDestinationPixelFormat(p PixelFormat) error {
	return ssc.update(softwareScaleContextUpdate{dstFormat: &p})
}

func (ssc *SoftwareScaleContext) DestinationResolution() (width int, height int) {
	return int(ssc.dstW), int(ssc.dstH)
}

func (ssc *SoftwareScaleContext) SetDestinationResolution(w int, h int) error {
	return ssc.update(softwareScaleContextUpdate{dstW: &w, dstH: &h})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#aa7dc7a4f9ec57a7c37957259a51cd920
func (ssc *SoftwareScaleContext) SourceWidth() int {
	return int(ssc.srcW)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a0ff71c9ef5ab6dabf90378fa7bf836ec
func (ssc *SoftwareScaleContext) SetSourceWidth(i int) error {
	return ssc.update(softwareScaleContextUpdate{srcW: &i})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a0dbc8c02bd3b4cd472e07008009751ff
func (ssc *SoftwareScaleContext) SourceHeight() int {
	return int(ssc.srcH)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#a0ff71c9ef5ab6dabf90378fa7bf836ec
func (ssc *SoftwareScaleContext) SetSourceHeight(i int) error {
	return ssc.update(softwareScaleContextUpdate{srcH: &i})
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#aab113373f157ee3b255ad97481af0cd9
func (ssc *SoftwareScaleContext) SourcePixelFormat() PixelFormat {
	return PixelFormat(ssc.srcFormat)
}

// https://ffmpeg.org/doxygen/8.0/structSwsContext.html#aab113373f157ee3b255ad97481af0cd9
func (ssc *SoftwareScaleContext) SetSourcePixelFormat(p PixelFormat) error {
	return ssc.update(softwareScaleContextUpdate{srcFormat: &p})
}

func (ssc *SoftwareScaleContext) SourceResolution() (int, int) {
	return int(ssc.srcW), int(ssc.srcH)
}

func (ssc *SoftwareScaleContext) SetSourceResolution(w int, h int) error {
	return ssc.update(softwareScaleContextUpdate{srcW: &w, srcH: &h})
}
