package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSetBufferRef_SelfAssignmentNoUAF is the regression test for the
// latent UAF in the original Frame.SetHardwareFramesContext /
// CodecContext.SetHardwareFramesContext / CodecContext.SetHardwareDeviceContext
// implementations: when src and *dst aliased the same underlying
// AVBuffer at refcount=1, the unref-then-ref order would drop the
// refcount to zero (freeing the buffer) before re-referencing it,
// reading freed memory.
//
// Contract under test: after setBufferRef(&dst, src) where src and *dst
// alias the same buffer, the buffer must still be live (refcount >= 1)
// and *dst must still point to it.
func TestSetBufferRef_SelfAssignmentNoUAF(t *testing.T) {
	// Allocate a buffer; refcount=1, *dst owns the only reference.
	dst := allocBufferRefForTest(64)
	require.NotNil(t, dst)
	require.Equal(t, 1, dst.refcountForTest(), "fresh buffer must start at refcount=1")

	dataBefore := dst.dataPtrForTest()
	require.NotNil(t, dataBefore, "fresh buffer must have non-nil data")

	// Self-assign: src aliases the same underlying AVBuffer as *dst.
	// Pre-fix order (unref-first) would drop refcount=1 → 0, freeing
	// the buffer, then ref-new would dereference freed memory. The
	// fixed order (ref-new first → refcount becomes 2; unref-old →
	// refcount returns to 1) keeps the buffer live throughout.
	setBufferRefForTest(dst, dst)

	require.Equal(t, 1, dst.refcountForTest(),
		"self-assign must leave refcount unchanged at 1")
	require.Equal(t, dataBefore, dst.dataPtrForTest(),
		"self-assign must not relocate or free the underlying buffer")

	freeBufferRefForTest(dst)
}

// TestSetBufferRef_NewAssignment covers the common case: dst initially
// nil, src holds a freshly allocated buffer. Result: dst takes a
// reference (refcount goes from 1 to 2 on the underlying buffer);
// neither side is freed.
func TestSetBufferRef_NewAssignment(t *testing.T) {
	src := allocBufferRefForTest(64)
	require.NotNil(t, src)
	require.Equal(t, 1, src.refcountForTest())

	dst := &bufferRefHandle{}
	require.Equal(t, 0, dst.refcountForTest(), "empty handle has refcount 0")

	setBufferRefForTest(dst, src)

	require.Equal(t, 2, src.refcountForTest(),
		"src takes a new reference, so refcount = original 1 + new 1 = 2")
	require.Equal(t, 2, dst.refcountForTest(),
		"dst now references the same underlying buffer")
	require.Equal(t, src.dataPtrForTest(), dst.dataPtrForTest(),
		"dst and src must reference the same data after assignment")

	freeBufferRefForTest(dst)
	require.Equal(t, 1, src.refcountForTest(),
		"freeing dst must drop the count back to 1")
	freeBufferRefForTest(src)
}

// TestSetBufferRef_NilSrcClearsDst covers the explicit clearing path:
// passing src=nil to a non-nil dst must unref dst's buffer and leave
// the field nil.
func TestSetBufferRef_NilSrcClearsDst(t *testing.T) {
	dst := allocBufferRefForTest(64)
	require.NotNil(t, dst)
	require.Equal(t, 1, dst.refcountForTest())

	setBufferRefForTest(dst, nil)

	require.Equal(t, 0, dst.refcountForTest(),
		"nil src must unref dst → refcount 0 (buffer freed)")
	require.Nil(t, dst.dataPtrForTest(), "dst's underlying ref must be cleared")
}

// TestSetBufferRef_NilSrcOnNilDst covers the trivial no-op: both nil.
// Must not crash.
func TestSetBufferRef_NilSrcOnNilDst(t *testing.T) {
	dst := &bufferRefHandle{}
	setBufferRefForTest(dst, nil)
	require.Nil(t, dst.dataPtrForTest())
	require.Equal(t, 0, dst.refcountForTest())
}

// TestSetBufferRef_ReplacesOldRef covers transitive replacement:
// dst already holds buffer A; src holds buffer B; after the call dst
// must hold a reference to B and A must be freed (no leak).
func TestSetBufferRef_ReplacesOldRef(t *testing.T) {
	dstStart := allocBufferRefForTest(64)
	src := allocBufferRefForTest(64)
	require.NotNil(t, dstStart)
	require.NotNil(t, src)
	require.NotEqual(t, dstStart.dataPtrForTest(), src.dataPtrForTest(),
		"two fresh buffers must occupy distinct allocations")

	// dst is a separate handle pointing at dstStart's buffer; we
	// mutate its .c field via setBufferRefForTest.
	dst := &bufferRefHandle{c: dstStart.c}
	dstStart.c = nil // dst now owns the reference; dstStart is empty.

	dataA := dst.dataPtrForTest()
	dataB := src.dataPtrForTest()
	require.NotNil(t, dataA)
	require.NotNil(t, dataB)
	require.Equal(t, 1, src.refcountForTest())

	setBufferRefForTest(dst, src)

	// dst now references buffer B; buffer A has been freed.
	require.Equal(t, dataB, dst.dataPtrForTest(),
		"dst must point to src's buffer after replacement")
	require.Equal(t, 2, src.refcountForTest(),
		"src and dst share the same underlying buffer (refcount 2)")

	freeBufferRefForTest(dst)
	freeBufferRefForTest(src)
}
