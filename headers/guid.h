//
// Created by Bohdan Dymchenko on 17.06.2017.
//

#ifndef SAM_GUID_H
#define SAM_GUID_H

#ifndef GUID
#include <stdint.h>

typedef struct
{
    uint32_t Data1;                                      /**< [in]: Specifies the first 8 hexadecimal digits of the GUID.                                */
    uint16_t Data2;                                      /**< [in]: Specifies the first group of 4 hexadecimal digits.                                   */
    uint16_t Data3;                                      /**< [in]: Specifies the second group of 4 hexadecimal digits.                                  */
    uint8_t  Data4[8];                                   /**< [in]: Array of 8 bytes. The first 2 bytes contain the third group of 4 hexadecimal digits.
                                                                    The remaining 6 bytes contain the final 12 hexadecimal digits.                       */
} GUID;
#endif // GUID

GUID CODEC_H264_GUID = { 0x6bc82762, 0x4e63, 0x4ca4, { 0xaa, 0x85, 0x1e, 0x50, 0xf3, 0x21, 0xf6, 0xbf } };
GUID CODEC_HEVC_GUID = { 0x790cdc88, 0x4522, 0x4d7b, { 0x94, 0x25, 0xbd, 0xa9, 0x97, 0x5f, 0x76, 0x3 } };
GUID CODEC_PROFILE_AUTOSELECT_GUID = { 0xbfd6f8e7, 0x233c, 0x4341, { 0x8b, 0x3e, 0x48, 0x18, 0x52, 0x38, 0x3, 0xf4 } };
GUID H264_PROFILE_BASELINE_GUID = { 0x727bcaa, 0x78c4, 0x4c83, { 0x8c, 0x2f, 0xef, 0x3d, 0xff, 0x26, 0x7c, 0x6a } };
GUID H264_PROFILE_MAIN_GUID = { 0x60b5c1d4, 0x67fe, 0x4790, { 0x94, 0xd5, 0xc4, 0x72, 0x6d, 0x7b, 0x6e, 0x6d } };
GUID H264_PROFILE_HIGH_GUID = { 0xe7cbc309, 0x4f7a, 0x4b89, { 0xaf, 0x2a, 0xd5, 0x37, 0xc9, 0x2b, 0xe3, 0x10 } };
GUID H264_PROFILE_HIGH_444_GUID = { 0x7ac663cb, 0xa598, 0x4960, { 0xb8, 0x44, 0x33, 0x9b, 0x26, 0x1a, 0x7d, 0x52 } };
GUID H264_PROFILE_STEREO_GUID = { 0x40847bf5, 0x33f7, 0x4601, { 0x90, 0x84, 0xe8, 0xfe, 0x3c, 0x1d, 0xb8, 0xb7 } };
GUID H264_PROFILE_SVC_TEMPORAL_SCALABILTY = { 0xce788d20, 0xaaa9, 0x4318, { 0x92, 0xbb, 0xac, 0x7e, 0x85, 0x8c, 0x8d, 0x36 } };
GUID H264_PROFILE_PROGRESSIVE_HIGH_GUID = { 0xb405afac, 0xf32b, 0x417b, { 0x89, 0xc4, 0x9a, 0xbe, 0xed, 0x3e, 0x59, 0x78 } };
GUID H264_PROFILE_CONSTRAINED_HIGH_GUID = { 0xaec1bd87, 0xe85b, 0x48f2, { 0x84, 0xc3, 0x98, 0xbc, 0xa6, 0x28, 0x50, 0x72 } };
GUID HEVC_PROFILE_MAIN_GUID = { 0xb514c39a, 0xb55b, 0x40fa, { 0x87, 0x8f, 0xf1, 0x25, 0x3b, 0x4d, 0xfd, 0xec } };

GUID PRESET_DEFAULT_GUID = { 0xb2dfb705, 0x4ebd, 0x4c49, { 0x9b, 0x5f, 0x24, 0xa7, 0x77, 0xd3, 0xe5, 0x87 } };
GUID PRESET_HP_GUID = { 0x60e4c59f, 0xe846, 0x4484, { 0xa5, 0x6d, 0xcd, 0x45, 0xbe, 0x9f, 0xdd, 0xf6 } };
GUID PRESET_HQ_GUID = { 0x34dba71d, 0xa77b, 0x4b8f, { 0x9c, 0x3e, 0xb6, 0xd5, 0xda, 0x24, 0xc0, 0x12 } };
GUID PRESET_BD_GUID  = { 0x82e3e450, 0xbdbb, 0x4e40, { 0x98, 0x9c, 0x82, 0xa9, 0xd, 0xf9, 0xef, 0x32 } };
GUID PRESET_LOW_LATENCY_DEFAULT_GUID  = { 0x49df21c5, 0x6dfa, 0x4feb, { 0x97, 0x87, 0x6a, 0xcc, 0x9e, 0xff, 0xb7, 0x26 } };
GUID PRESET_LOW_LATENCY_HQ_GUID  = { 0xc5f733b9, 0xea97, 0x4cf9, { 0xbe, 0xc2, 0xbf, 0x78, 0xa7, 0x4f, 0xd1, 0x5 } };
GUID PRESET_LOW_LATENCY_HP_GUID = { 0x67082a44, 0x4bad, 0x48fa, { 0x98, 0xea, 0x93, 0x5, 0x6d, 0x15, 0xa, 0x58 } };
GUID PRESET_LOSSLESS_DEFAULT_GUID = { 0xd5bfb716, 0xc604, 0x44e7, { 0x9b, 0xb8, 0xde, 0xa5, 0x51, 0xf, 0xc3, 0xac } };
GUID PRESET_LOSSLESS_HP_GUID = { 0x149998e7, 0x2364, 0x411d, { 0x82, 0xef, 0x17, 0x98, 0x88, 0x9, 0x34, 0x9 } };


#endif //SAM_ENCODER_TYPES_H_H
