package nvenc

import "C"

const (
	Success                  int = iota // This indicates that API call returned with no errors.
	ErrNoEncodeDevice                   // This indicates that no encode capable devices were detected.
	ErrUnsupportedDevice                // This indicates that devices pass by the client is not supported.
	ErrInvalidEncoderDevice             // This indicates that the encoder device supplied by the client is not valid.
	ErrInvalidDevice                    // This indicates that device passed to the API call is invalid.
	ErrDeviceNotExist                   // This indicates that device passed to the API call is no longer available and needs to be reinitialized. The clients need to destroy the current encoder session by freeing the allocated input output buffers and destroying the device and create a new encoding session.
	ErrInvalidPtr                       // This indicates that one or more of the pointers passed to the API call is invalid.
	ErrInvalidEvent                     // This indicates that completion event passed in ::NvEncEncodePicture() call is invalid.
	ErrInvalidParam                     // This indicates that one or more of the parameter passed to the API call is invalid.
	ErrInvalidCall                      // This indicates that an API call was made in wrong sequence/order.
	ErrOutOfMemory                      // This indicates that the API call failed because it was unable to allocate enough memory to perform the requested operation.
	ErrEncoderNotInitialized            // This indicates that the encoder has not been initialized with ::NvEncInitializeEncoder() or that initialization has failed. The client cannot allocate input or output buffers or do any encoding related operation before successfully initializing the encoder.
	ErrUnsupportedParam                 // This indicates that an unsupported parameter was passed by the client.
	ErrLockBusy                         // This indicates that the ::NvEncLockBitstream() failed to lock the output buffer. This happens when the client makes a non blocking lock call to access the output bitstream by passing NV_ENC_LOCK_BITSTREAM::doNotWait flag. This is not a fatal error and client should retry the same operation after few milliseconds.
	ErrNotEnoughBuffer                  // This indicates that the size of the user buffer passed by the client is insufficient for the requested operation.
	ErrInvalidVersion                   // ErrInvalidVersion indicates that an invalid struct version was used by the client.
	ErrMapFailed                        // ErrMapFailed indicates that ::NvEncMapInputResource() API failed to map the client provided input resource.
	/*
	 * ErrNeedMoreInput indicates encode driver requires more input buffers to produce an output
	 * bitstream. If this error is returned from ::NvEncEncodePicture() API, this
	 * is not a fatal error. If the client is encoding with B frames then,
	 * ::NvEncEncodePicture() API might be buffering the input frame for re-ordering.
	 *
	 * A client operating in synchronous mode cannot call ::NvEncLockBitstream()
	 * API on the output bitstream buffer if ::NvEncEncodePicture() returned the
	 * ::NV_ENC_ERR_NEED_MORE_INPUT error code.
	 * The client must continue providing input frames until encode driver returns
	 * ::NV_ENC_SUCCESS. After receiving ::NV_ENC_SUCCESS status the client can call
	 * ::NvEncLockBitstream() API on the output buffers in the same order in which
	 * it has called ::NvEncEncodePicture().
	 */
	ErrNeedMoreInput
	ErrEncoderBusy            // ErrEncoderBusy indicates that the HW encoder is busy encoding and is unable to encode the input. The client should call ::NvEncEncodePicture() again after few milliseconds.
	ErrEventNotRegistered     // This indicates that the completion event passed in ::NvEncEncodePicture() API has not been registered with encoder driver using ::NvEncRegisterAsyncEvent().
	ErrGeneric                // This indicates that an unknown internal error has occurred.
	ErrIncompatibleClientKey  // This indicates that the client is attempting to use a feature that is not available for the license type for the current system.
	ErrUnimplemented          // This indicates that the client is attempting to use a feature that is not implemented for the current version.
	ErrResourceRegisterFailed // This indicates that the ::NvEncRegisterResource API failed to register the resource.
	ErrResourceNotRegistered  // This indicates that the client is attempting to unregister a resource that has not been successfully registered.
	ErrResourceNotMapped      // This indicates that the client is attempting to unmap a resource that has not been successfully mapped.
)

type apiError int

func (a apiError) Error() string {
	return errorMessages[int(a)]
}

var errorMessages = map[int]string{
	ErrNoEncodeDevice:         "no encode capable devices were detected",
	ErrUnsupportedDevice:      "the device pass by the client is not supported",
	ErrInvalidEncoderDevice:   "the encoder device supplied by the client is not valid",
	ErrInvalidDevice:          "the device passed to the API call is invalid",
	ErrDeviceNotExist:         "device passed to the API call is no longer available and needs to be reinitialized",
	ErrInvalidPtr:             "one or more pointers passed are invalid",
	ErrInvalidEvent:           "completion event passed in ::NvEncEncodePicture() call is invalid.",
	ErrInvalidParam:           "one or more param passed to API are invalid",
	ErrInvalidCall:            "an API call was made in wrong sequence/order",
	ErrOutOfMemory:            "the API call failed because it was unable to allocate enough memory to perform the requested operation",
	ErrEncoderNotInitialized:  "the encoder has not been initialized with NvEncInitializeEncoder() or that initialization has failed",
	ErrUnsupportedParam:       "an unsupported parameter was passed by the client",
	ErrLockBusy:               "the ::NvEncLockBitstream() failed to lock the output buffer",
	ErrNotEnoughBuffer:        "the size of the user buffer passed by the client is insufficient for the requested operation",
	ErrInvalidVersion:         "an invalid struct version was used by the client",
	ErrMapFailed:              "::NvEncMapInputResource() API failed to map the client provided input resource",
	ErrNeedMoreInput:          "encode driver requires more input buffers to produce an output bitstream.",
	ErrEncoderBusy:            "the HW encoder is busy encoding and is unable to encode the input.",
	ErrEventNotRegistered:     "the completion event passed in ::NvEncEncodePicture() API has not been registered with encoder driver using ::NvEncRegisterAsyncEvent()",
	ErrGeneric:                "an unknown internal error has occurred",
	ErrIncompatibleClientKey:  "the client is attempting to use a feature that is not available for the license type for the current system",
	ErrUnimplemented:          "that the client is attempting to use a feature that is not implemented for the current version",
	ErrResourceRegisterFailed: "the ::NvEncRegisterResource API failed to register the resource",
	ErrResourceNotRegistered:  "the client is attempting to unregister a resource that has not been successfully registered",
	ErrResourceNotMapped:      "the client is attempting to unmap a resource that has not been successfully mapped",
}

func codeToError(code C.int) error {
	if int(code) == Success {
		return nil
	}

	return apiError(int(code))
}
