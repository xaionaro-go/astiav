package astiav

//#include <libavcodec/avcodec.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/group__lavu__frame.html#ga893dbd60c5ebe415600523fbae202880
type FrameCropFlag int64

const (
	FrameCropFlagUnaligned = FrameCropFlag(C.AV_FRAME_CROP_UNALIGNED)
)
