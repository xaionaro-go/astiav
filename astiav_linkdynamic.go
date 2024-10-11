//go:build !libav_static
// +build !libav_static

package astiav

//#cgo pkg-config: libavcodec libavdevice libavfilter libavformat libswscale libavutil
import "C"
