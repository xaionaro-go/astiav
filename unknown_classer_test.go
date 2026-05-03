package astiav

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestUnknownClasser_PointerString covers the diagnostic-only formatter
// that replaced the previously-exported Pointer() unsafe.Pointer
// getter. By returning a string, the API makes accidental dereference
// of a stale pointer impossible at the package boundary while still
// preserving the originating address as identification in log output.
func TestUnknownClasser_PointerString(t *testing.T) {
	// Allocate a real pointer in C heap so the address is stable and
	// distinct from any Go object the GC might move.
	p := allocBytesForTest(8)
	require.NotNil(t, p)
	defer freeBytesForTest(p)

	uc := newUnknownClasser(p)
	require.NotNil(t, uc)
	require.Nil(t, uc.Class(), "an unknown classer must not expose a usable Class — see newUnknownClasser doc")

	s := uc.PointerString()
	// Format: "0x" followed by hex digits, matching Go's %p verb.
	require.Regexp(t, regexp.MustCompile(`^0x[0-9a-f]+$`), s,
		"PointerString must use Go's %%p hex format")
}

// TestUnknownClasser_PointerString_Nil covers the safety contracts for
// nil receivers and nil pointers — both must produce "<nil>" instead
// of crashing.
func TestUnknownClasser_PointerString_Nil(t *testing.T) {
	var nilUc *UnknownClasser
	require.Equal(t, "<nil>", nilUc.PointerString(),
		"nil receiver must yield <nil>, not panic")

	emptyUc := &UnknownClasser{}
	require.Equal(t, "<nil>", emptyUc.PointerString(),
		"empty UnknownClasser (no ptr) must yield <nil>")
}
