package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestHardwareFramesContext_GettersPreInitialize is the regression test
// for the nil-data crash: HardwarePixelFormat and SoftwarePixelFormat
// were dereferencing hfc.data() unconditionally, so a zero-value or
// uninitialized HardwareFramesContext would SIGSEGV inside the cgo
// shim before the user even saw the alloc failure.
//
// Contract under test: pre-Initialize getters return PixelFormatNone
// instead of crashing, so misordered code surfaces a benign sentinel
// the caller can detect.
func TestHardwareFramesContext_GettersPreInitialize(t *testing.T) {
	// Zero-value receiver: hfc.c == nil.
	var hfcZero HardwareFramesContext
	require.Equal(t, PixelFormatNone, hfcZero.HardwarePixelFormat(),
		"zero-value receiver must return PixelFormatNone, not crash")
	require.Equal(t, PixelFormatNone, hfcZero.SoftwarePixelFormat(),
		"zero-value receiver must return PixelFormatNone, not crash")

	// nil receiver via newHardwareFramesContextFromC's nil-input path.
	require.Nil(t, newHardwareFramesContextFromC(nil),
		"nil C ref must yield nil receiver (existing contract)")
}

// TestHardwareFramesContext_GettersAfterAlloc verifies the happy path:
// after AllocHardwareFramesContext succeeds (data is non-nil) the
// getters return the values previously set by SetHardwarePixelFormat /
// SetSoftwarePixelFormat. This complements the pre-init test by
// proving the new nil-data guard does not regress the normal flow.
func TestHardwareFramesContext_GettersAfterAlloc(t *testing.T) {
	// Pick the first available HW device — any will do; we never call
	// Initialize so the hwframe state stays in the "allocated, not
	// initialized" stage where data is non-nil and writable but no
	// frames have been allocated.
	t.Helper()
	hdc := tryAnyHardwareDevice(t)
	if hdc == nil {
		t.Skip("no usable hardware device on this host; skipping happy-path getter test")
	}
	defer hdc.Free()

	hfc := AllocHardwareFramesContext(hdc)
	require.NotNil(t, hfc)
	defer hfc.Free()

	hfc.SetHardwarePixelFormat(PixelFormatCuda)
	hfc.SetSoftwarePixelFormat(PixelFormatNv12)
	hfc.SetWidth(64)
	hfc.SetHeight(64)
	hfc.SetInitialPoolSize(1)

	require.Equal(t, PixelFormatCuda, hfc.HardwarePixelFormat(),
		"getter must return value previously written by setter")
	require.Equal(t, PixelFormatNv12, hfc.SoftwarePixelFormat(),
		"getter must return value previously written by setter")
}

// tryAnyHardwareDevice attempts to create any hardware device context,
// preferring CUDA → DRM → VAAPI → MediaCodec. Returns nil if none can
// be created on the current host (e.g. CI without GPU access).
func tryAnyHardwareDevice(t *testing.T) *HardwareDeviceContext {
	t.Helper()
	candidates := []HardwareDeviceType{
		HardwareDeviceTypeCUDA,
		HardwareDeviceTypeDRM,
		HardwareDeviceTypeVAAPI,
		HardwareDeviceTypeMediaCodec,
	}
	for _, ty := range candidates {
		hdc, err := CreateHardwareDeviceContext(ty, "", nil, 0)
		if err == nil && hdc != nil {
			return hdc
		}
	}
	return nil
}
