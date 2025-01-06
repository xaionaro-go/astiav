//go:build ffmpeg7
// +build ffmpeg7

package astiav

//#include <libavfilter/buffersink.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/group__lavfi__buffersink__accessors.html#gab80976e506ab88d23d94bb6d7a4051bd
func (bfc *BuffersinkFilterContext) ColorRange() ColorRange {
	return ColorRange(C.av_buffersink_get_color_range(bfc.fc.c))
}

// https://ffmpeg.org/doxygen/7.0/group__lavfi__buffersink__accessors.html#gaad817cdcf5493c385126e8e17c5717f2
func (bfc *BuffersinkFilterContext) ColorSpace() ColorSpace {
	return ColorSpace(C.av_buffersink_get_colorspace(bfc.fc.c))
}
