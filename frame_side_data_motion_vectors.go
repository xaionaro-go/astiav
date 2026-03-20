package astiav

//#include <libavutil/frame.h>
//#include <libavutil/motion_vector.h>
//#include "frame_side_data.h"
import "C"
import (
	"math"
	"unsafe"
)

type frameSideDataMotionVectors struct {
	d *FrameSideData
}

func newFrameSideDataMotionVectors(d *FrameSideData) *frameSideDataMotionVectors {
	return &frameSideDataMotionVectors{d: d}
}

func (d *frameSideDataMotionVectors) data(sd *C.AVFrameSideData) *[(math.MaxInt32 - 1) / C.sizeof_AVMotionVector]C.AVMotionVector {
	return (*[(math.MaxInt32 - 1) / C.sizeof_AVMotionVector]C.AVMotionVector)(unsafe.Pointer(C.astiavConvertMotionVectorsFrameSideData(sd)))
}

// Get returns all motion vectors from the frame side data.
// Returns nil and false if no motion vector side data is present.
func (d *frameSideDataMotionVectors) Get() ([]MotionVector, bool) {
	sd := C.av_frame_side_data_get(*d.d.sd, *d.d.size, C.AV_FRAME_DATA_MOTION_VECTORS)
	if sd == nil {
		return nil, false
	}

	cmvs := d.data(sd)
	count := int(sd.size / C.sizeof_AVMotionVector)
	mvs := make([]MotionVector, count)
	for i := range mvs {
		mvs[i] = MotionVector{
			Source:      int32(cmvs[i].source),
			W:           uint8(cmvs[i].w),
			H:           uint8(cmvs[i].h),
			SrcX:        int16(cmvs[i].src_x),
			SrcY:        int16(cmvs[i].src_y),
			DstX:        int16(cmvs[i].dst_x),
			DstY:        int16(cmvs[i].dst_y),
			Flags:       uint64(cmvs[i].flags),
			MotionX:     int32(cmvs[i].motion_x),
			MotionY:     int32(cmvs[i].motion_y),
			MotionScale: uint16(cmvs[i].motion_scale),
		}
	}
	return mvs, true
}
