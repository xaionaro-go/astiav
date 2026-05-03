package astiav

//#include <libavutil/buffer.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// allocBufferRefForTest allocates a refcounted AVBuffer of size bytes
// from the C heap and returns a *C.AVBufferRef wrapped in an opaque
// handle. Tests use it to exercise setBufferRef without depending on a
// real hardware device.
type bufferRefHandle struct {
	c *C.AVBufferRef
}

func allocBufferRefForTest(size int) *bufferRefHandle {
	ref := C.av_buffer_alloc(C.size_t(size))
	if ref == nil {
		return nil
	}
	return &bufferRefHandle{c: ref}
}

// refcountForTest returns the number of references currently held on
// the underlying AVBuffer (i.e. av_buffer_get_ref_count).
func (h *bufferRefHandle) refcountForTest() int {
	if h == nil || h.c == nil {
		return 0
	}
	return int(C.av_buffer_get_ref_count(h.c))
}

// dataPtrForTest returns the address of the buffer's data, useful for
// verifying that two refs point at the same underlying allocation.
func (h *bufferRefHandle) dataPtrForTest() unsafe.Pointer {
	if h == nil || h.c == nil {
		return nil
	}
	return unsafe.Pointer(h.c.data)
}

// setBufferRefForTest exposes setBufferRef to *_test.go files via a
// pointer-to-pointer the test can manipulate. It mirrors the in-tree
// usage at struct field call sites (e.g. &cc.c.hw_frames_ctx).
func setBufferRefForTest(dst *bufferRefHandle, src *bufferRefHandle) {
	var srcRaw *C.AVBufferRef
	if src != nil {
		srcRaw = src.c
	}
	setBufferRef(&dst.c, srcRaw)
}

// freeBufferRefForTest unrefs the underlying AVBuffer if still held;
// safe to call multiple times. Tests must call this on every handle
// they create (after the test asserts succeed) to keep refcount
// accounting clean across tests.
func freeBufferRefForTest(h *bufferRefHandle) {
	if h == nil || h.c == nil {
		return
	}
	C.av_buffer_unref(&h.c)
}
