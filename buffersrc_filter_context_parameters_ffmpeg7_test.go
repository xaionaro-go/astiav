//go:build ffmpeg7
// +build ffmpeg7

package astiav

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuffersrcFilterContextParameters(t *testing.T) {
	p := AllocBuffersrcFilterContextParameters()
	defer p.Free()
	p.SetColorRange(ColorRangeMpeg)
	require.Equal(t, ColorRangeMpeg, p.ColorRange())
	p.SetColorSpace(ColorSpaceBt470Bg)
	require.Equal(t, ColorSpaceBt470Bg, p.ColorSpace())
}
