package astiav

//#include <libavutil/log.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// dispatchLogCallbackForTest invokes the Go-side log callback dispatch
// with the given ptr, simulating an FFmpeg log emission from the C
// side. It exists so tests can exercise goAstiavLogCallback without
// having cgo in *_test.go files (which Go's test toolchain rejects).
func dispatchLogCallbackForTest(ptr unsafe.Pointer, level LogLevel, fmtStr, msg string) {
	fmtC := C.CString(fmtStr)
	defer C.free(unsafe.Pointer(fmtC))
	msgC := C.CString(msg)
	defer C.free(unsafe.Pointer(msgC))
	goAstiavLogCallback(ptr, C.int(level), fmtC, msgC)
}

// allocBytesForTest allocates size bytes from the C heap and returns a
// pointer the caller must release via freeBytesForTest. It is used by
// tests to simulate raw memory left behind by a freed AVFormatContext.
func allocBytesForTest(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

// freeBytesForTest releases memory obtained from allocBytesForTest.
func freeBytesForTest(p unsafe.Pointer) {
	C.free(p)
}
