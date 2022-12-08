package nvenc

// #include "headers/guid.h"
import "C"
import "fmt"

type (
	GUID = C.GUID

	codecGUID   GUID
	profileGUID GUID
	presetGUID  GUID
)

var (
	CodecH264Guid = codecGUID(C.CODEC_H264_GUID)
	CodecHevcGuid = codecGUID(C.CODEC_HEVC_GUID)

	ProfileAutoSelectGuid            = profileGUID(C.CODEC_PROFILE_AUTOSELECT_GUID)
	ProfileH264BaselineGuid          = profileGUID(C.H264_PROFILE_BASELINE_GUID)
	ProfileH264MainGuid              = profileGUID(C.H264_PROFILE_MAIN_GUID)
	ProfileH264HighGuid              = profileGUID(C.H264_PROFILE_HIGH_GUID)
	ProfileH264High444Guid           = profileGUID(C.H264_PROFILE_HIGH_444_GUID)
	ProfileH264StereoGuid            = profileGUID(C.H264_PROFILE_STEREO_GUID)
	ProfileH264SvcTemporalScalabilty = profileGUID(C.H264_PROFILE_SVC_TEMPORAL_SCALABILTY)
	ProfileH264ProgressiveHighGuid   = profileGUID(C.H264_PROFILE_PROGRESSIVE_HIGH_GUID)
	ProfileH264ConstrainedHighGuid   = profileGUID(C.H264_PROFILE_CONSTRAINED_HIGH_GUID)
	ProfileHevcMainGuid              = profileGUID(C.HEVC_PROFILE_MAIN_GUID)

	PresetDefaultGuid           = presetGUID(C.PRESET_DEFAULT_GUID)
	PresetHpGuid                = presetGUID(C.PRESET_HP_GUID)
	PresetHqGuid                = presetGUID(C.PRESET_HQ_GUID)
	PresetBdGuid                = presetGUID(C.PRESET_BD_GUID)
	PresetLowLatencyDefaultGuid = presetGUID(C.PRESET_LOW_LATENCY_DEFAULT_GUID)
	PresetLowLatencyHqGuid      = presetGUID(C.PRESET_LOW_LATENCY_HQ_GUID)
	PresetLowLatencyHpGuid      = presetGUID(C.PRESET_LOW_LATENCY_HP_GUID)
	PresetLosslessDefaultGuid   = presetGUID(C.PRESET_LOSSLESS_DEFAULT_GUID)
	PresetLosslessHpGuid        = presetGUID(C.PRESET_LOSSLESS_HP_GUID)
)

func (guid codecGUID) String() string {
	switch guid {
	case CodecH264Guid:
		return "h264"
	case CodecHevcGuid:
		return "hevc"
	default:
		return "unknown"
	}
}

func (guid profileGUID) String() string {
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

func (guid presetGUID) String() string {
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

func (e *Encoder) getGUIDs() ([]GUID, error) {
	count, err := e.functions.getGUIDCount(e.instance)
	if err != nil {
		return nil, fmt.Errorf("getGUIDCount error: %w", err)
	}
	return e.functions.getGUIDs(e.instance, count)
}

type guids interface {
	profileGUID | presetGUID | codecGUID
}

func hasGUID[T guids](slice []GUID, g T) bool {
	for _, r := range slice {
		if r == GUID(g) {
			return true
		}
	}

	return false
}
