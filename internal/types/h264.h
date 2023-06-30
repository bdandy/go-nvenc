#include <nvEncodeAPI.h>

 static inline void EnableRepeatSPSPPS(NV_ENC_CONFIG_H264* c) {
 	c->repeatSPSPPS = 1;
 }

 static inline void DisableSPSPPS(NV_ENC_CONFIG_H264* c) {
 	c->disableSPSPPS = 1;
 }

 static inline void EnableAUD(NV_ENC_CONFIG_H264* c) {
 	c->outputAUD = 1;
 }

 static inline void EnableIntraRefresh(NV_ENC_CONFIG_H264* c) {
 	c->enableIntraRefresh = 1;
 }

 static inline void EnableVFR(NV_ENC_CONFIG_H264* c) {
 	c->enableVFR = 1;
 }

 static inline void EnableLTR(NV_ENC_CONFIG_H264* c){
 	c->enableLTR = 1;
 }

 static inline void EnableConstrainedEncoding(NV_ENC_CONFIG_H264* c) {
 	c->enableConstrainedEncoding = 1;
 }

 static inline void UseConstrainedIntraPred(NV_ENC_CONFIG_H264* c) {
 	c->useConstrainedIntraPred = 1;
 }

 static inline void EnableOutputBufferingPeriodSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputBufferingPeriodSEI = 1;
 }

 static inline void EnableOutputPictureTimingSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputPictureTimingSEI = 1;
 }

 static inline void EnableOutputFramePackingSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputFramePackingSEI = 1;
 }

 static inline void EnableOutputRecoveryPointSEI(NV_ENC_CONFIG_H264* c) {
 	c->outputRecoveryPointSEI = 1;
 }

 static inline void SetConstrainedFrame(NV_ENC_PIC_PARAMS_H264* p, int i) {
 	p->constrainedFrame = i;
 }