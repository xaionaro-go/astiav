package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFrameSideDataMotionVectors_Empty verifies that MotionVectors().Get()
// returns nil and false for a frame without motion vector side data.
// Agent-generated test.
func TestFrameSideDataMotionVectors_Empty(t *testing.T) {
	f := AllocFrame()
	require.NotNil(t, f)
	defer f.Free()

	sd := f.SideData()
	mvs, ok := sd.MotionVectors().Get()
	require.False(t, ok)
	require.Nil(t, mvs)
}
