package astiav

// MotionVector represents a single motion vector extracted from frame side data.
//
// https://ffmpeg.org/doxygen/8.0/structAVMotionVector.html
type MotionVector struct {
	// Source is the reference frame index: negative for past, positive for future.
	Source int32
	// W is the block width.
	W uint8
	// H is the block height.
	H uint8
	// SrcX is the absolute source X position.
	SrcX int16
	// SrcY is the absolute source Y position.
	SrcY int16
	// DstX is the absolute destination X position.
	DstX int16
	// DstY is the absolute destination Y position.
	DstY int16
	// Flags contains extra flag information (currently unused by FFmpeg).
	Flags uint64
	// MotionX is the motion delta X in 1/MotionScale units.
	MotionX int32
	// MotionY is the motion delta Y in 1/MotionScale units.
	MotionY int32
	// MotionScale is the divisor for motion vectors (typically 4 for quarter-pixel).
	MotionScale uint16
}
