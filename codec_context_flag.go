package astiav

//#include <libavcodec/avcodec.h>
import "C"

// https://ffmpeg.org/doxygen/7.0/group__lavc__core.html#gaa52d62f5dbfc4529388f0454ae671359
type CodecContextFlag int64

const (
	CodecContextFlagUnaligned     = CodecContextFlag(C.AV_CODEC_FLAG_UNALIGNED)
	CodecContextFlagQscale        = CodecContextFlag(C.AV_CODEC_FLAG_QSCALE)
	CodecContextFlag4Mv           = CodecContextFlag(C.AV_CODEC_FLAG_4MV)
	CodecContextFlagOutputCorrupt = CodecContextFlag(C.AV_CODEC_FLAG_OUTPUT_CORRUPT)
	CodecContextFlagQpel          = CodecContextFlag(C.AV_CODEC_FLAG_QPEL)
	CodecContextFlagDropChanged   = CodecContextFlag(C.AV_CODEC_FLAG_DROPCHANGED)
	CodecContextFlagReconFrame    = CodecContextFlag(C.AV_CODEC_FLAG_RECON_FRAME)
	CodecContextFlagCopyOpaque    = CodecContextFlag(C.AV_CODEC_FLAG_COPY_OPAQUE)
	CodecContextFlagFrameDuration = CodecContextFlag(C.AV_CODEC_FLAG_FRAME_DURATION)
	CodecContextFlagPass1         = CodecContextFlag(C.AV_CODEC_FLAG_PASS1)
	CodecContextFlagPass2         = CodecContextFlag(C.AV_CODEC_FLAG_PASS2)
	CodecContextFlagLoopFilter    = CodecContextFlag(C.AV_CODEC_FLAG_LOOP_FILTER)
	CodecContextFlagGray          = CodecContextFlag(C.AV_CODEC_FLAG_GRAY)
	CodecContextFlagPsnr          = CodecContextFlag(C.AV_CODEC_FLAG_PSNR)
	CodecContextFlagInterlacedDct = CodecContextFlag(C.AV_CODEC_FLAG_INTERLACED_DCT)
	CodecContextFlagLowDelay      = CodecContextFlag(C.AV_CODEC_FLAG_LOW_DELAY)
	CodecContextFlagGlobalHeader  = CodecContextFlag(C.AV_CODEC_FLAG_GLOBAL_HEADER)
	CodecContextFlagBitexact      = CodecContextFlag(C.AV_CODEC_FLAG_BITEXACT)
	CodecContextFlagAcPred        = CodecContextFlag(C.AV_CODEC_FLAG_AC_PRED)
	CodecContextFlagInterlacedMe  = CodecContextFlag(C.AV_CODEC_FLAG_INTERLACED_ME)
	CodecContextFlagClosedGop     = CodecContextFlag(C.AV_CODEC_FLAG_CLOSED_GOP)
)

type CodecContextFlag2 int64

// https://ffmpeg.org/doxygen/7.0/group__lavc__core.html#ga1a6a486e686ab6c581ffffcb88cb31b3
const (
	CodecFlag2Fast        = CodecContextFlag2(C.AV_CODEC_FLAG2_FAST)
	CodecFlag2NoOutput    = CodecContextFlag2(C.AV_CODEC_FLAG2_NO_OUTPUT)
	CodecFlag2LocalHeader = CodecContextFlag2(C.AV_CODEC_FLAG2_LOCAL_HEADER)
	CodecFlag2Chunks      = CodecContextFlag2(C.AV_CODEC_FLAG2_CHUNKS)
	CodecFlag2IgnoreCrop  = CodecContextFlag2(C.AV_CODEC_FLAG2_IGNORE_CROP)
	CodecFlag2ShowAll     = CodecContextFlag2(C.AV_CODEC_FLAG2_SHOW_ALL)
	CodecFlag2ExportMvs   = CodecContextFlag2(C.AV_CODEC_FLAG2_EXPORT_MVS)
	CodecFlag2SkipManual  = CodecContextFlag2(C.AV_CODEC_FLAG2_SKIP_MANUAL)
	CodecFlag2RoFlushNoop = CodecContextFlag2(C.AV_CODEC_FLAG2_RO_FLUSH_NOOP)
	CodecFlag2IccProfiles = CodecContextFlag2(C.AV_CODEC_FLAG2_ICC_PROFILES)
)
