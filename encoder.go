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

func (e *Encoder) SetPreset(guid presetGUID) error {
	e.presetGUID = GUID(guid)
	conf, err := e.functions.getPresetConfig(e.instance, e.encodeGUID, GUID(guid))
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

// Encode is main function for encoding raw image into video bitstream.
// It allocates output buffer for each frame, which will be garbage collected
func (e *Encoder) Encode(data []byte) ([]byte, error) {
	var buf *[]byte

	err := e.EncodeTo(data, buf)

	return *buf, err
}

// EncodeTo encodes picture into pre-allocated buffer. If buf is nil - new buffer will be allocated
func (e *Encoder) EncodeTo(data []byte, buf *[]byte) error {
	if err := e.encode(data); err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	if err := e.copyOutput(buf); err != nil {
		return fmt.Errorf("copy output: %w", err)
	}
	return nil
}

func (e *Encoder) encode(data []byte) error {
	if err := e.functions.lockInputBuffer(e.instance, e.inputBufferLock); err != nil {
		return fmt.Errorf("lockInputBuffer: %w", err)
	}
	e.inputBufferLock.CopyBuffer(data)

	if err := e.functions.unlockInputBuffer(e.instance, e.inputBuffer.GetBufferPtr()); err != nil {
		return fmt.Errorf("unlockInputBuffer: %w", err)
	}

	if err := e.functions.encodePicture(e.instance, e.pictureParams); err != nil {
		return fmt.Errorf("encodePicture: %w", err)
	}

	return nil
}

func (e *Encoder) copyOutput(buf *[]byte) error {
	if err := e.functions.lockBitstream(e.instance, e.outputBufferLock); err != nil {
		return fmt.Errorf("lockBitstream: %w", err)
	}

	size := e.outputBufferLock.BitstreamSize()

	if buf == nil {
		b := make([]byte, size)
		buf = &b
	}

	if err := e.outputBufferLock.CopyBitstream(*buf); err != nil {
		return fmt.Errorf("copy bistream: %w", err)
	}

	if err := e.functions.unlockBitstream(e.instance, e.outputBuffer.GetBufferPtr()); err != nil {
		return fmt.Errorf("unlockBitstream: %w", err)
	}

	return nil
}

// InvalidateRefFrames invalidates reference for timestamp
func (e *Encoder) InvalidateRefFrames(timestamp uint64) error {
	return e.functions.invalidateRefFrames(e.instance, timestamp)
}

func (e *Encoder) Picture() *ENC_PIC_PARAMS {
	return e.pictureParams
}

// GetSequence returns SPS\PPS header
func (e *Encoder) GetSequence() ([]byte, error) {
	err := e.functions.getSequenceParams(e.instance, e.spsPayload)
	if err != nil {
		return nil, fmt.Errorf("get sequence params: %w", err)
	}

	return e.spsPayload.Bytes(), nil
}

// Destroy clean ups encoder
func (e *Encoder) Destroy() error {
	err := e.functions.destroyBitstreamBuffer(e.instance, e.outputBuffer.GetBufferPtr())
	if err != nil {
		return fmt.Errorf("destroy bitsteam buffer: %w", err)
	}

	err = e.functions.destroyBuffer(e.instance, e.inputBuffer.GetBufferPtr())
	if err != nil {
		return fmt.Errorf("destroy buffer: %w", err)
	}

	err = e.functions.destroyEncoder(e.instance)
	if err != nil {
		return fmt.Errorf("destroy encoder: %w", err)
	}

	e.instance = nil
	e.functions = nil
	e.encoderConfig = nil
	e.initializeParams = nil
	e.pictureParams = nil
	e.inputBuffer = nil
	e.outputBuffer = nil

	return nil
}

// NewEncoder returns initialized encoder instance for chosen Codec (h264,hevc) with output buffer allocated to bufSize
func NewEncoder(codec codecGUID, bufSize uint32) (*Encoder, error) {
	enc, err := newEncoder(bufSize)
	if err != nil {
		return nil, fmt.Errorf("new encoder: %w", err)
	}

	guids, err := enc.getGUIDs()
	if err != nil {
		return nil, fmt.Errorf("get GUIDs: %w", err)
	}

	if !hasGUID(guids, codec) {
		return nil, fmt.Errorf("NvEncoder doesn't support %s codec", codec)
	}

	return enc, nil
}

func newEncoder(bufSize uint32) (*Encoder, error) {
	functions := newEncoderFunctions()

	err := callCreateInstance(functions)
	if err != nil {
		return nil, fmt.Errorf("callCreateInstance: %w", err)
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
