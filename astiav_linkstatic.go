//go:build libav_static
// +build libav_static

package astiav

//#cgo LDFLAGS: -Wl,-Bstatic -lavcodec -lavdevice -lavfilter -lavformat -lswscale -lavutil -Wl,-Bdynamic
import "C"
