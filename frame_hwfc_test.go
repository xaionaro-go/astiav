package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFrame_HardwareFramesContext_NilReceiver guards the new nil
// receiver / nil-c protection on Frame.HardwareFramesContext and
// Frame.SetHardwareFramesContext. Mirrors the nil-safety added to
// Class methods in 563ba58 — defense-in-depth so misordered code
// (e.g. operating on a Free()d Frame) surfaces a benign nil instead
// of a SIGSEGV.
func TestFrame_HardwareFramesContext_NilReceiver(t *testing.T) {
	// Nil pointer receiver.
	var f *Frame
	require.Nil(t, f.HardwareFramesContext(),
		"nil receiver must return nil, not crash")
	// Setter must not panic on nil receiver.
	f.SetHardwareFramesContext(nil)

	// Zero-value Frame (non-nil pointer, nil c — possible after Free).
	fEmpty := &Frame{}
	require.Nil(t, fEmpty.HardwareFramesContext(),
		"Frame with nil c must return nil")
	fEmpty.SetHardwareFramesContext(nil)
}

// TestFrame_SetHardwareFramesContext_RoundTrip verifies the happy path
// via the new setBufferRef-backed setter: setting then reading back
// returns a context that wraps the same underlying AVBufferRef. Catches
// regressions in the cgo type cast between *C.struct_AVBufferRef
// (HardwareFramesContext.c) and *C.AVBufferRef (setBufferRef helper).
func TestFrame_SetHardwareFramesContext_RoundTrip(t *testing.T) {
	hdc := tryAnyHardwareDevice(t)
	if hdc == nil {
		t.Skip("no usable hardware device on this host")
	}
	defer hdc.Free()

	hfc := AllocHardwareFramesContext(hdc)
	require.NotNil(t, hfc)
	defer hfc.Free()

	f := AllocFrame()
	require.NotNil(t, f)
	defer f.Free()

	require.Nil(t, f.HardwareFramesContext(), "fresh frame has no hwframes ctx")

	f.SetHardwareFramesContext(hfc)
	got := f.HardwareFramesContext()
	require.NotNil(t, got, "after Set, getter must return non-nil")
	require.Equal(t, hfc.c, got.c,
		"getter must yield the same underlying AVBufferRef as the setter received")

	// Self-assign must be safe (regression for the latent UAF).
	f.SetHardwareFramesContext(f.HardwareFramesContext())
	require.NotNil(t, f.HardwareFramesContext(),
		"self-assignment must keep the context attached")

	// Clear.
	f.SetHardwareFramesContext(nil)
	require.Nil(t, f.HardwareFramesContext(),
		"setting nil must clear the context")
}
