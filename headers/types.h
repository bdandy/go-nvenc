#include <stdio.h>
#include "headers/nvEncodeAPI.h"

 static inline NV_ENC_CONFIG_H264* GetH264Config(NV_ENC_CODEC_CONFIG* c) {
 	return &c->h264Config;
 }

 static inline NV_ENC_CONFIG_HEVC* GetHEVCConfig(NV_ENC_CODEC_CONFIG* c) {
 	return &c->hevcConfig;
 }

 static inline NV_ENC_PIC_PARAMS_H264* GetPicParamsH264(NV_ENC_CODEC_PIC_PARAMS* p) {
 	return &p->h264PicParams;
 }

 static inline void EnableMinQP(NV_ENC_RC_PARAMS* c) {
 	c->enableMinQP = 1;
 }

 static inline void EnableMaxQP(NV_ENC_RC_PARAMS* c) {
 	c->enableMaxQP = 1;
 }

 static inline void EnableTemporalAQ(NV_ENC_RC_PARAMS* c) {
 	c->enableTemporalAQ =1;
 }

 static inline void EnableSpartialAQ(NV_ENC_RC_PARAMS* c, int s) {
 	c->enableAQ = 1;
	c->aqStrength = s;
 }

 static inline void EnableZeroReorderDelay(NV_ENC_RC_PARAMS* c) {
 	c->zeroReorderDelay = 1;
 }

 static inline void EnableNonRefP(NV_ENC_RC_PARAMS* c) {
 	c->enableNonRefP = 1;
 }

 static inline void ResetEncoder(NV_ENC_RECONFIGURE_PARAMS* p) {
 	p->resetEncoder = 1;
 }

 static inline void EnableWeightedPrediction(NV_ENC_INITIALIZE_PARAMS* p) {
 	p->enableWeightedPrediction = 1;
 }

 static inline void SetEncodeConfig(NV_ENC_INITIALIZE_PARAMS* p, NV_ENC_CONFIG* c) {
    p->encodeConfig = c;
 }