package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLog_UnknownPtrIsNotDereferenced reproduces the UAF that triggers
// SIGSEGV at astiavClassCategory+24 on rapid input add->remove->add
// cycles (e.g. android_camera): an FFmpeg background thread fires
// av_log with a pointer to an already-freed AVFormatContext.
//
// goAstiavLogCallback then receives a ptr that is no longer in the
// classers pool. Pre-fix, it falls back to newUnknownClasser(ptr),
// which dereferences ptr to read *AVClass — reading freed memory.
// The garbage *AVClass survives until user code (e.g. logrus
// formatter) calls cl.Category(), which crashes inside C trying to
// load c->get_category.
//
// Contract under test: when ptr is non-nil but not registered in
// classers, the log callback must NOT dereference ptr. The Classer
// delivered to user code must yield Class() == nil so that downstream
// formatters cannot trigger a UAF on a stale AVFormatContext.
func TestLog_UnknownPtrIsNotDereferenced(t *testing.T) {
	defer ResetLogCallback()

	// Allocate a buffer in C heap (so Go GC won't move or scan it).
	// The first 8 bytes are read by the buggy code as *AVClass — set
	// them to a non-null sentinel that mimics garbage left by a freed
	// context.
	const sentinel uintptr = 0xdeadbeefcafebabe
	ptr := allocBytesForTest(64)
	require.NotNil(t, ptr)
	defer freeBytesForTest(ptr)
	*(*uintptr)(ptr) = sentinel

	// Ensure ptr is not in classers (it should not be — it's a raw
	// buffer not registered with any astiav object).
	_, ok := classers.get(ptr)
	require.False(t, ok, "test setup: ptr must not be registered")

	var captured Classer
	var captureCount int
	SetLogCallback(func(c Classer, l LogLevel, fmt, msg string) {
		captured = c
		captureCount++
	})

	dispatchLogCallbackForTest(ptr, LogLevelWarning, "", "test")

	require.Equal(t, 1, captureCount, "log callback must be invoked exactly once")

	// Critical contract: the Classer delivered for an unregistered
	// ptr must not expose a Class object backed by a dereferenced
	// ptr. Either captured is nil (no classer info) or
	// captured.Class() is nil — both signal "do not deref ptr".
	// Anything else means the binding read *ptr to build a Class,
	// which is the UAF.
	if captured == nil {
		return
	}
	cl := captured.Class()
	require.Nil(t, cl,
		"Class() for an unregistered ptr must be nil — otherwise "+
			"downstream callers can deref a stale AVFormatContext "+
			"(read garbage as *AVClass — sentinel %#x)", sentinel)
}
