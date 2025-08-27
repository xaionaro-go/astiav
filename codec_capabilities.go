package astiav

type CodecCapabilities uint32

// https://ffmpeg.org/doxygen/trunk/codec_8h_source.html
const (
	CodecCapabilityDrawHorizBand          = CodecCapabilities(1 << 0)
	CodecCapabilityDr1                    = CodecCapabilities(1 << 1)
	CodecCapabilityDelay                  = CodecCapabilities(1 << 5)
	CodecCapabilitySmallLastFrame         = CodecCapabilities(1 << 6)
	CodecCapabilitySubframes              = CodecCapabilities(1 << 8)
	CodecCapabilityExperimental           = CodecCapabilities(1 << 9)
	CodecCapabilityChannelConf            = CodecCapabilities(1 << 10)
	CodecCapabilityFrameThreads           = CodecCapabilities(1 << 12)
	CodecCapabilitySliceThreads           = CodecCapabilities(1 << 13)
	CodecCapabilityParamChange            = CodecCapabilities(1 << 14)
	CodecCapabilityOtherThreads           = CodecCapabilities(1 << 15)
	CodecCapabilityVariableFrameSize      = CodecCapabilities(1 << 16)
	CodecCapabilityAvoidProbing           = CodecCapabilities(1 << 17)
	CodecCapabilityHardware               = CodecCapabilities(1 << 18)
	CodecCapabilityHybrid                 = CodecCapabilities(1 << 19)
	CodecCapabilityEncoderReorderedOpaque = CodecCapabilities(1 << 20)
	CodecCapabilityEncoderFlush           = CodecCapabilities(1 << 21)
	CodecCapabilityEncoderReconFrame      = CodecCapabilities(1 << 22)
)
