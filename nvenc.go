package nvenc

import (
	"fmt"
	"unsafe"
)

type Encoder struct {
	instance unsafe.Pointer

	encodeGUID GUID
	presetGUID GUID

	initializeParams *INITIALIZE_PARAMS
	encoderConfig    *ENCODER_CONFIG
	pictureParams    *ENC_PIC_PARAMS
	inputBuffer      *INPUT_BUFFER
	inputBufferLock  *LOCK_INPUT_BUFFER_PARAMS
	outputBuffer     *BITSTREAM_BUFFER
	outputBufferLock *LOCK_BITSTREAM_PARAMS
	spsPayload       *SEQUENCE_PARAM_PAYLOAD

	functions *EncoderFunctions
}

func (e *Encoder) InitializeParams() *INITIALIZE_PARAMS {
	return e.initializeParams
}

func (e *Encoder) Config() *ENCODER_CONFIG {
	return e.encoderConfig
}

func (e *Encoder) ForceIDR(b bool) {
	e.pictureParams.ForceIDR(b)
}

func (e *Encoder) ForceIntraRefresh(frames uint32) {
	e.pictureParams.PicParamsH264().ForceIntraRefresh(frames)
}

func (e *Encoder) getGUIDs() ([]GUID, error) {
	count, err := e.functions.getGUIDCount(e.instance)
	if err != nil {
		return nil, fmt.Errorf("getGUIDCount error: %w", err)
	}
	return e.functions.getGUIDs(e.instance, count)
}

func (e *Encoder) UseH264() error {
	guids, err := e.getGUIDs()
	if err != nil {
		return err
	}
	if !containsGUID(guids, CODEC_H264_GUID) {
		return fmt.Errorf("NvEncoder doesn't support H264 profile")
	}
	e.encodeGUID = CODEC_H264_GUID

	return nil
}

func (e *Encoder) UseHEVC() error {
	guids, err := e.getGUIDs()
	if err != nil {
		return err
	}
	if !containsGUID(guids, CODEC_HEVC_GUID) {
		return fmt.Errorf("NvEncoder doesn't support HEVC profile")
	}
	e.encodeGUID = CODEC_HEVC_GUID

	return nil
}

func (e *Encoder) SetPreset(guid GUID) error {
	e.presetGUID = guid
	conf, err := e.functions.getPresetConfig(e.instance, e.encodeGUID, guid)
	if err != nil {
		return err
	}
	e.encoderConfig = conf
	// bugfix missing RC struct version
	e.encoderConfig.RC().setVersion()

	return nil
}

func (e *Encoder) GetInputFormats() ([]BUFFER_FORMAT, error) {
	count, err := e.functions.getInputFormatCount(e.instance, e.encodeGUID)
	if err != nil {
		return nil, fmt.Errorf("getInputFormatCount error: %w", err)
	}

	return e.functions.getInputFormats(e.instance, e.encodeGUID, count)
}

func (e *Encoder) GetPresets() ([]GUID, error) {
	count, err := e.functions.getPresetCount(e.instance, e.encodeGUID)
	if err != nil {
		return nil, err
	}
	return e.functions.getPresetGUIDs(e.instance, e.encodeGUID, count)
}

func (e *Encoder) SetResolution(width, height uint32) {
	e.initializeParams.SetResolution(width, height)
	e.inputBuffer.SetResolution(width, height)
	e.pictureParams.SetResolution(width, height)
}

func (e *Encoder) SetFrameRate(num, den uint32) {
	e.initializeParams.SetFrameRate(num, den)
}

func (e *Encoder) InitializeEncoder(inputFormat, outputFormat BUFFER_FORMAT) (err error) {
	e.initializeParams.SetEncodeGUID(e.encodeGUID)
	e.initializeParams.SetPresetGUID(e.presetGUID)
	e.initializeParams.SetEncodeConfig(e.encoderConfig)

	if err = e.functions.initializeEncoder(e.instance, e.initializeParams); err != nil {
		return err
	}

	e.inputBuffer.SetFormat(inputFormat)
	e.pictureParams.SetInputFormat(inputFormat)

	if err := e.functions.createBuffer(e.instance, e.inputBuffer); err != nil {
		return err
	}
	e.pictureParams.SetInputBuffer(e.inputBuffer.GetBufferPtr())

	e.inputBufferLock = newLockInputBufferParams(e.inputBuffer)

	if err := e.functions.createBitstreamBuffer(e.instance, e.outputBuffer); err != nil {
		return err
	}
	e.pictureParams.SetOutputBuffer(e.outputBuffer.GetBufferPtr())
	e.outputBufferLock = newBitstreamBufferLock(e.outputBuffer)

	return nil
}

func (e *Encoder) Reset() error {
	var params RECONFIGURE_PARAMS
	e.initializeParams.SetPresetGUID(e.presetGUID)
	e.initializeParams.SetEncodeConfig(e.encoderConfig)
	params.SetInitializeParams(*e.initializeParams)
	return e.functions.reconfigureEncoder(e.instance, &params)
}

func (e *Encoder) Encode(data []byte) ([]byte, error) {
	if err := e.functions.lockInputBuffer(e.instance, e.inputBufferLock); err != nil {
		return nil, fmt.Errorf("lockInputBuffer: %w", err)
	}
	e.inputBufferLock.CopyBuffer(data)

	if err := e.functions.unlockInputBuffer(e.instance, e.inputBuffer.GetBufferPtr()); err != nil {
		return nil, fmt.Errorf("unlockInputBuffer: %w", err)
	}

	if err := e.functions.encodePicture(e.instance, e.pictureParams); err != nil {
		return nil, fmt.Errorf("encodePicture: %w", err)
	}

	if err := e.functions.lockBitstream(e.instance, e.outputBufferLock); err != nil {
		return nil, fmt.Errorf("lockBitstream: %w", err)
	}

	result := e.outputBufferLock.GetData()
	if err := e.functions.unlockBitstream(e.instance, e.outputBuffer.GetBufferPtr()); err != nil {
		return nil, fmt.Errorf("unlockBitstream: %w", err)
	}

	return result, nil
}

func (e *Encoder) InvalidateRefFrames(timestamp uint64) error {
	return e.functions.invalidateRefFrames(e.instance, timestamp)
}

func (e *Encoder) Picture() *ENC_PIC_PARAMS {
	return e.pictureParams
}

func (e *Encoder) GetSPSPPS() ([]byte, error) {
	err := e.functions.getSequenceParams(e.instance, e.spsPayload)
	if err != nil {
		return nil, fmt.Errorf("GetSPSPPS: %w", err)
	}

	return e.spsPayload.Bytes(), nil
}

func (e *Encoder) Destroy() error {
	err := e.functions.destroyBitstreamBuffer(e.instance, e.outputBuffer.GetBufferPtr())
	err = e.functions.destroyBuffer(e.instance, e.inputBuffer.GetBufferPtr())
	err = e.functions.destroyEncoder(e.instance)

	e.instance = nil
	e.functions = nil
	e.encoderConfig = nil
	e.initializeParams = nil
	e.pictureParams = nil
	e.inputBuffer = nil
	e.outputBuffer = nil
	return err
}

func NewEncoder(bufSize uint32) (*Encoder, error) {
	functions := newEncoderFunctions()

	err := callCreateInstance(functions)
	if err != nil {
		return nil, err
	}

	enc := &Encoder{
		functions:        functions,
		inputBuffer:      newInputBufferParams(),
		outputBuffer:     newBitstreamBuffer(bufSize),
		initializeParams: newInitializeParams(),
		encoderConfig:    newEncoderConfig(),
		pictureParams:    newEncPicParams(),
		spsPayload:       newSequenceParamPayload(),
	}

	return enc, nil
}
