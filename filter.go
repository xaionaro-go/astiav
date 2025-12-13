package astiav

//#include <libavfilter/avfilter.h>
import "C"
import (
	"unsafe"
)

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html
type Filter struct {
	c *C.AVFilter
}

func newFilterFromC(c *C.AVFilter) *Filter {
	if c == nil {
		return nil
	}
	return &Filter{c: c}
}

// https://ffmpeg.org/doxygen/7.0/group__lavfi.html#gadd774ec49e50edf00158248e1bfe4ae6
func FindFilterByName(n string) *Filter {
	cn := C.CString(n)
	defer C.free(unsafe.Pointer(cn))
	return newFilterFromC(C.avfilter_get_by_name(cn))
}

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#a632c76418742ad4f4dccbd4db40badd0
func (f *Filter) Flags() FilterFlags {
	return FilterFlags(f.c.flags)
}

// https://ffmpeg.org/doxygen/7.0/structAVFilter.html#a28a4776f344f91055f42a4c2a1b15c0c
func (f *Filter) Name() string {
	return C.GoString(f.c.name)
}

func (f *Filter) String() string {
	return f.Name()
}
