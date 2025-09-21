package astiav

//#include <libavcodec/avcodec.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/group__lavu__frame__flags.html
type FrameFlag int64

const (
	FrameFlagCorrupt       = FrameFlag(C.AV_FRAME_FLAG_CORRUPT)
	FrameFlagKey           = FrameFlag(C.AV_FRAME_FLAG_KEY)
	FrameFlagDiscard       = FrameFlag(C.AV_FRAME_FLAG_DISCARD)
	FrameFlagInterlaced    = FrameFlag(C.AV_FRAME_FLAG_INTERLACED)
	FrameFlagTopFieldFirst = FrameFlag(C.AV_FRAME_FLAG_TOP_FIELD_FIRST)
)
