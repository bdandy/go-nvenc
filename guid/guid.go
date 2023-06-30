package guid

// #cgo CFLAGS: -I ../include
// #include <guid.h>
import "C"
import "unsafe"

type (
	GUID  C.GUID
	GUIDs []GUID

	CodecGUID   GUID
	ProfileGUID GUID
	PresetGUID  GUID
)

var (
	CodecH264Guid = CodecGUID(C.CODEC_H264_GUID)
	CodecHevcGuid = CodecGUID(C.CODEC_HEVC_GUID)

	ProfileAutoSelectGuid            = ProfileGUID(C.CODEC_PROFILE_AUTOSELECT_GUID)
	ProfileH264BaselineGuid          = ProfileGUID(C.H264_PROFILE_BASELINE_GUID)
	ProfileH264MainGuid              = ProfileGUID(C.H264_PROFILE_MAIN_GUID)
	ProfileH264HighGuid              = ProfileGUID(C.H264_PROFILE_HIGH_GUID)
	ProfileH264High444Guid           = ProfileGUID(C.H264_PROFILE_HIGH_444_GUID)
	ProfileH264StereoGuid            = ProfileGUID(C.H264_PROFILE_STEREO_GUID)
	ProfileH264SvcTemporalScalabilty = ProfileGUID(C.H264_PROFILE_SVC_TEMPORAL_SCALABILTY)
	ProfileH264ProgressiveHighGuid   = ProfileGUID(C.H264_PROFILE_PROGRESSIVE_HIGH_GUID)
	ProfileH264ConstrainedHighGuid   = ProfileGUID(C.H264_PROFILE_CONSTRAINED_HIGH_GUID)
	ProfileHevcMainGuid              = ProfileGUID(C.HEVC_PROFILE_MAIN_GUID)

	PresetDefaultGuid           = PresetGUID(C.PRESET_DEFAULT_GUID)
	PresetHpGuid                = PresetGUID(C.PRESET_HP_GUID)
	PresetHqGuid                = PresetGUID(C.PRESET_HQ_GUID)
	PresetBdGuid                = PresetGUID(C.PRESET_BD_GUID)
	PresetLowLatencyDefaultGuid = PresetGUID(C.PRESET_LOW_LATENCY_DEFAULT_GUID)
	PresetLowLatencyHqGuid      = PresetGUID(C.PRESET_LOW_LATENCY_HQ_GUID)
	PresetLowLatencyHpGuid      = PresetGUID(C.PRESET_LOW_LATENCY_HP_GUID)
	PresetLosslessDefaultGuid   = PresetGUID(C.PRESET_LOSSLESS_DEFAULT_GUID)
	PresetLosslessHpGuid        = PresetGUID(C.PRESET_LOSSLESS_HP_GUID)
)

func (guid CodecGUID) String() string {
	switch guid {
	case CodecH264Guid:
		return "h264"
	case CodecHevcGuid:
		return "hevc"
	default:
		return "unknown"
	}
}

func (guid ProfileGUID) String() string {
	switch guid {
	case ProfileH264BaselineGuid:
		return "baseline"
	case ProfileH264MainGuid, ProfileHevcMainGuid:
		return "main"
	case ProfileH264HighGuid:
		return "high"
	case ProfileH264High444Guid:
		return "high444"
	case ProfileH264StereoGuid:
		return "stereo"
	case ProfileH264SvcTemporalScalabilty:
		return "svc_temporal_scalability"
	case ProfileH264ProgressiveHighGuid:
		return "progressive_high"
	case ProfileH264ConstrainedHighGuid:
		return "constrained_high"
	case ProfileAutoSelectGuid:
		return "auto-select"
	default:
		return "unknown"
	}
}

func (guid PresetGUID) String() string {
	switch guid {
	case PresetDefaultGuid:
		return "default"
	case PresetHpGuid:
		return "hp"
	case PresetHqGuid:
		return "hq"
	case PresetBdGuid:
		return "bd"
	case PresetLowLatencyDefaultGuid:
		return "low_latency_default"
	case PresetLowLatencyHpGuid:
		return "low_latency_hp"
	case PresetLowLatencyHqGuid:
		return "low_latency_hq"
	case PresetLosslessDefaultGuid:
		return "lossless_default"
	case PresetLosslessHpGuid:
		return "lossless_hp"
	default:
		return "unknown"
	}
}

func (guid CodecGUID) CType() C.GUID {
	return C.GUID(guid)
}

func (guid PresetGUID) CType() C.GUID {
	return C.GUID(guid)
}

func (guid ProfileGUID) CType() C.GUID {
	return C.GUID(guid)
}

func (guids GUIDs) ToPreset() []PresetGUID {
	return unsafe.Slice((*PresetGUID)(&guids[0]), len(guids))
}
