// Code generated by "enumer -type=voiceType -linecomment -json"; DO NOT EDIT.

//
package azuretexttospeech

import (
	"encoding/json"
	"fmt"
)

const _voiceTypeName = "StandardNeural"

var _voiceTypeIndex = [...]uint8{0, 8, 14}

func (i voiceType) String() string {
	if i < 0 || i >= voiceType(len(_voiceTypeIndex)-1) {
		return fmt.Sprintf("voiceType(%d)", i)
	}
	return _voiceTypeName[_voiceTypeIndex[i]:_voiceTypeIndex[i+1]]
}

var _voiceTypeValues = []voiceType{0, 1}

var _voiceTypeNameToValueMap = map[string]voiceType{
	_voiceTypeName[0:8]:  0,
	_voiceTypeName[8:14]: 1,
}

// voiceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func voiceTypeString(s string) (voiceType, error) {
	if val, ok := _voiceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to voiceType values", s)
}

// voiceTypeValues returns all values of the enum
func voiceTypeValues() []voiceType {
	return _voiceTypeValues
}

// IsAvoiceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i voiceType) IsAvoiceType() bool {
	for _, v := range _voiceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for voiceType
func (i voiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for voiceType
func (i *voiceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("voiceType should be a string, got %s", data)
	}

	var err error
	*i, err = voiceTypeString(s)
	return err
}
