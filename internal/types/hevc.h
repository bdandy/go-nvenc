#include <nvEncodeAPI.h>


 static inline void HEVCSetRepeatSPSPPS(NV_ENC_CONFIG_HEVC* c) {
 	c->repeatSPSPPS = 1;
 }

 static inline void HEVCSetDisableSPSPPS(NV_ENC_CONFIG_HEVC* c) {
 	c->disableSPSPPS = 1;
 }

 static inline void HEVCSetEnableAUD(NV_ENC_CONFIG_HEVC* c) {
 	c->outputAUD = 1;
 }

 static inline void HEVCSetEnableIntraRefresh(NV_ENC_CONFIG_HEVC* c) {
 	c->enableIntraRefresh = 1;
 }