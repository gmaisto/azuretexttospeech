// Code generated by "enumer -type=AudioOutput"; DO NOT EDIT.

package azuretexttospeech

import (
	"fmt"
	"strings"
)

const _AudioOutputName = "AudioRIFF8Bit8kHzMonoPCMAudioRIFF16Bit16kHzMonoPCMAudioRIFF16khz16kbpsMonoSirenAudioRIFF24khz16bitMonoPcmAudioRAW8Bit8kHzMonoMulawAudioRAW16Bit16kHzMonoMulawAudioRAW24khz16bitMonoPcmAudioSsml16khz16bitMonoTtsAudio16khz16kbpsMonoSirenAudio16khz32kbitrateMonoMp3Audio6khz64kbitrateMonoMp3Audio16khz128kbitrateMonoMp3Audio24khz48kbitrateMonoMp3Audio24khz96kbitrateMonoMp3"

var _AudioOutputIndex = [...]uint16{0, 24, 50, 79, 105, 130, 157, 182, 208, 233, 260, 286, 314, 341, 368}

const _AudioOutputLowerName = "audioriff8bit8khzmonopcmaudioriff16bit16khzmonopcmaudioriff16khz16kbpsmonosirenaudioriff24khz16bitmonopcmaudioraw8bit8khzmonomulawaudioraw16bit16khzmonomulawaudioraw24khz16bitmonopcmaudiossml16khz16bitmonottsaudio16khz16kbpsmonosirenaudio16khz32kbitratemonomp3audio6khz64kbitratemonomp3audio16khz128kbitratemonomp3audio24khz48kbitratemonomp3audio24khz96kbitratemonomp3"

func (i AudioOutput) String() string {
	if i < 0 || i >= AudioOutput(len(_AudioOutputIndex)-1) {
		return fmt.Sprintf("AudioOutput(%d)", i)
	}
	return _AudioOutputName[_AudioOutputIndex[i]:_AudioOutputIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AudioOutputNoOp() {
	var x [1]struct{}
	_ = x[AudioRIFF8Bit8kHzMonoPCM-(0)]
	_ = x[AudioRIFF16Bit16kHzMonoPCM-(1)]
	_ = x[AudioRIFF16khz16kbpsMonoSiren-(2)]
	_ = x[AudioRIFF24khz16bitMonoPcm-(3)]
	_ = x[AudioRAW8Bit8kHzMonoMulaw-(4)]
	_ = x[AudioRAW16Bit16kHzMonoMulaw-(5)]
	_ = x[AudioRAW24khz16bitMonoPcm-(6)]
	_ = x[AudioSsml16khz16bitMonoTts-(7)]
	_ = x[Audio16khz16kbpsMonoSiren-(8)]
	_ = x[Audio16khz32kbitrateMonoMp3-(9)]
	_ = x[Audio6khz64kbitrateMonoMp3-(10)]
	_ = x[Audio16khz128kbitrateMonoMp3-(11)]
	_ = x[Audio24khz48kbitrateMonoMp3-(12)]
	_ = x[Audio24khz96kbitrateMonoMp3-(13)]
}

var _AudioOutputValues = []AudioOutput{AudioRIFF8Bit8kHzMonoPCM, AudioRIFF16Bit16kHzMonoPCM, AudioRIFF16khz16kbpsMonoSiren, AudioRIFF24khz16bitMonoPcm, AudioRAW8Bit8kHzMonoMulaw, AudioRAW16Bit16kHzMonoMulaw, AudioRAW24khz16bitMonoPcm, AudioSsml16khz16bitMonoTts, Audio16khz16kbpsMonoSiren, Audio16khz32kbitrateMonoMp3, Audio6khz64kbitrateMonoMp3, Audio16khz128kbitrateMonoMp3, Audio24khz48kbitrateMonoMp3, Audio24khz96kbitrateMonoMp3}

var _AudioOutputNameToValueMap = map[string]AudioOutput{
	_AudioOutputName[0:24]:         AudioRIFF8Bit8kHzMonoPCM,
	_AudioOutputLowerName[0:24]:    AudioRIFF8Bit8kHzMonoPCM,
	_AudioOutputName[24:50]:        AudioRIFF16Bit16kHzMonoPCM,
	_AudioOutputLowerName[24:50]:   AudioRIFF16Bit16kHzMonoPCM,
	_AudioOutputName[50:79]:        AudioRIFF16khz16kbpsMonoSiren,
	_AudioOutputLowerName[50:79]:   AudioRIFF16khz16kbpsMonoSiren,
	_AudioOutputName[79:105]:       AudioRIFF24khz16bitMonoPcm,
	_AudioOutputLowerName[79:105]:  AudioRIFF24khz16bitMonoPcm,
	_AudioOutputName[105:130]:      AudioRAW8Bit8kHzMonoMulaw,
	_AudioOutputLowerName[105:130]: AudioRAW8Bit8kHzMonoMulaw,
	_AudioOutputName[130:157]:      AudioRAW16Bit16kHzMonoMulaw,
	_AudioOutputLowerName[130:157]: AudioRAW16Bit16kHzMonoMulaw,
	_AudioOutputName[157:182]:      AudioRAW24khz16bitMonoPcm,
	_AudioOutputLowerName[157:182]: AudioRAW24khz16bitMonoPcm,
	_AudioOutputName[182:208]:      AudioSsml16khz16bitMonoTts,
	_AudioOutputLowerName[182:208]: AudioSsml16khz16bitMonoTts,
	_AudioOutputName[208:233]:      Audio16khz16kbpsMonoSiren,
	_AudioOutputLowerName[208:233]: Audio16khz16kbpsMonoSiren,
	_AudioOutputName[233:260]:      Audio16khz32kbitrateMonoMp3,
	_AudioOutputLowerName[233:260]: Audio16khz32kbitrateMonoMp3,
	_AudioOutputName[260:286]:      Audio6khz64kbitrateMonoMp3,
	_AudioOutputLowerName[260:286]: Audio6khz64kbitrateMonoMp3,
	_AudioOutputName[286:314]:      Audio16khz128kbitrateMonoMp3,
	_AudioOutputLowerName[286:314]: Audio16khz128kbitrateMonoMp3,
	_AudioOutputName[314:341]:      Audio24khz48kbitrateMonoMp3,
	_AudioOutputLowerName[314:341]: Audio24khz48kbitrateMonoMp3,
	_AudioOutputName[341:368]:      Audio24khz96kbitrateMonoMp3,
	_AudioOutputLowerName[341:368]: Audio24khz96kbitrateMonoMp3,
}

var _AudioOutputNames = []string{
	_AudioOutputName[0:24],
	_AudioOutputName[24:50],
	_AudioOutputName[50:79],
	_AudioOutputName[79:105],
	_AudioOutputName[105:130],
	_AudioOutputName[130:157],
	_AudioOutputName[157:182],
	_AudioOutputName[182:208],
	_AudioOutputName[208:233],
	_AudioOutputName[233:260],
	_AudioOutputName[260:286],
	_AudioOutputName[286:314],
	_AudioOutputName[314:341],
	_AudioOutputName[341:368],
}

// AudioOutputString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AudioOutputString(s string) (AudioOutput, error) {
	if val, ok := _AudioOutputNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AudioOutputNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AudioOutput values", s)
}

// AudioOutputValues returns all values of the enum
func AudioOutputValues() []AudioOutput {
	return _AudioOutputValues
}

// AudioOutputStrings returns a slice of all String values of the enum
func AudioOutputStrings() []string {
	strs := make([]string, len(_AudioOutputNames))
	copy(strs, _AudioOutputNames)
	return strs
}

// IsAAudioOutput returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AudioOutput) IsAAudioOutput() bool {
	for _, v := range _AudioOutputValues {
		if i == v {
			return true
		}
	}
	return false
}
