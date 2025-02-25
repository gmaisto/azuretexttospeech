// Code generated by "enumer -type=VoiceType -linecomment -json"; DO NOT EDIT.

package azuretexttospeech

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _VoiceTypeName = "StandardNeuralNeutralNeuralHD"

var _VoiceTypeIndex = [...]uint8{0, 8, 14, 21, 29}

const _VoiceTypeLowerName = "standardneuralneutralneuralhd"

func (i VoiceType) String() string {
	if i < 0 || i >= VoiceType(len(_VoiceTypeIndex)-1) {
		return fmt.Sprintf("VoiceType(%d)", i)
	}
	return _VoiceTypeName[_VoiceTypeIndex[i]:_VoiceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _VoiceTypeNoOp() {
	var x [1]struct{}
	_ = x[VoiceStandard-(0)]
	_ = x[VoiceNeural-(1)]
	_ = x[VoiceNeutral-(2)]
	_ = x[VoiceNeuralHD-(3)]
}

var _VoiceTypeValues = []VoiceType{VoiceStandard, VoiceNeural, VoiceNeutral, VoiceNeuralHD}

var _VoiceTypeNameToValueMap = map[string]VoiceType{
	_VoiceTypeName[0:8]:        VoiceStandard,
	_VoiceTypeLowerName[0:8]:   VoiceStandard,
	_VoiceTypeName[8:14]:       VoiceNeural,
	_VoiceTypeLowerName[8:14]:  VoiceNeural,
	_VoiceTypeName[14:21]:      VoiceNeutral,
	_VoiceTypeLowerName[14:21]: VoiceNeutral,
	_VoiceTypeName[21:29]:      VoiceNeuralHD,
	_VoiceTypeLowerName[21:29]: VoiceNeuralHD,
}

var _VoiceTypeNames = []string{
	_VoiceTypeName[0:8],
	_VoiceTypeName[8:14],
	_VoiceTypeName[14:21],
	_VoiceTypeName[21:29],
}

// VoiceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func VoiceTypeString(s string) (VoiceType, error) {
	if val, ok := _VoiceTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _VoiceTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to VoiceType values", s)
}

// VoiceTypeValues returns all values of the enum
func VoiceTypeValues() []VoiceType {
	return _VoiceTypeValues
}

// VoiceTypeStrings returns a slice of all String values of the enum
func VoiceTypeStrings() []string {
	strs := make([]string, len(_VoiceTypeNames))
	copy(strs, _VoiceTypeNames)
	return strs
}

// IsAVoiceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i VoiceType) IsAVoiceType() bool {
	for _, v := range _VoiceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for VoiceType
func (i VoiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoiceType
func (i *VoiceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("VoiceType should be a string, got %s", data)
	}

	var err error
	*i, err = VoiceTypeString(s)
	return err
}
