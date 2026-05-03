package astiav

//#include <libavutil/buffer.h>
import "C"

// setBufferRef replaces *dst with a new reference to src in a way that
// is safe even when *dst and src ultimately reference the same underlying
// AVBuffer (i.e. self-assignment, e.g. f.SetHardwareFramesContext(f.HardwareFramesContext())).
//
// The order is intentional: take the new reference FIRST, drop the old
// one SECOND. The naive opposite order (unref-old, then ref-new) is a
// latent UAF: if *dst was the only reference and src points at the same
// underlying buffer, av_buffer_unref drops the refcount to zero and
// frees the buffer; the subsequent av_buffer_ref then dereferences freed
// memory.
//
// Passing src == nil clears *dst to nil after unref-ing the old value.
//
// Centralized here so all setters that own an AVBufferRef field
// (Frame.SetHardwareFramesContext, CodecContext.SetHardwareFramesContext,
// CodecContext.SetHardwareDeviceContext, FilterContext.SetHardwareDeviceContext,
// BuffersrcFilterContextParameters.SetHardwareFramesContext) share one
// proven-correct implementation. Updating callers must keep using this
// helper rather than re-implementing the unref/ref dance inline.
func setBufferRef(dst **C.AVBufferRef, src *C.AVBufferRef) {
	var newRef *C.AVBufferRef
	if src != nil {
		newRef = C.av_buffer_ref(src)
	}
	if *dst != nil {
		C.av_buffer_unref(dst)
	}
	*dst = newRef
}
