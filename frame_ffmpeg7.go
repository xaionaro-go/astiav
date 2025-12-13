//go:build ffmpeg7
// +build ffmpeg7

package astiav

//#include <libavutil/channel_layout.h>
//#include <libavutil/frame.h>
//#include <libavutil/imgutils.h>
//#include <libavutil/samplefmt.h>
//#include <libavutil/hwcontext.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/structAVFrame.html#afe0345882416bbb9d3a86720dcaa9252
func (f *Frame) KeyFrame() bool {
	return int(f.c.key_frame) > 0
}

// https://ffmpeg.org/doxygen/7.0/structAVFrame.html#afe0345882416bbb9d3a86720dcaa9252
func (f *Frame) SetKeyFrame(k bool) {
	i := 0
	if k {
		i = 1
	}
	f.c.key_frame = C.int(i)
}
