// Code generated by "enumer --values --type=UserDataType --output type_user_data.go --transform=snake-upper"; DO NOT EDIT.

package farcaster

import (
	"fmt"
	"strings"
)

const (
	_UserDataTypeName_0      = "USER_DATA_TYPE_NONEUSER_DATA_TYPE_PFPUSER_DATA_TYPE_DISPLAYUSER_DATA_TYPE_BIO"
	_UserDataTypeLowerName_0 = "user_data_type_noneuser_data_type_pfpuser_data_type_displayuser_data_type_bio"
	_UserDataTypeName_1      = "USER_DATA_TYPE_URLUSER_DATA_TYPE_USERNAME"
	_UserDataTypeLowerName_1 = "user_data_type_urluser_data_type_username"
)

var (
	_UserDataTypeIndex_0 = [...]uint8{0, 19, 37, 59, 77}
	_UserDataTypeIndex_1 = [...]uint8{0, 18, 41}
)

func (i UserDataType) String() string {
	switch {
	case 0 <= i && i <= 3:
		return _UserDataTypeName_0[_UserDataTypeIndex_0[i]:_UserDataTypeIndex_0[i+1]]
	case 5 <= i && i <= 6:
		i -= 5
		return _UserDataTypeName_1[_UserDataTypeIndex_1[i]:_UserDataTypeIndex_1[i+1]]
	default:
		return fmt.Sprintf("UserDataType(%d)", i)
	}
}

func (UserDataType) Values() []string {
	return UserDataTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UserDataTypeNoOp() {
	var x [1]struct{}
	_ = x[UserDataTypeNone-(0)]
	_ = x[UserDataTypePfp-(1)]
	_ = x[UserDataTypeDisplay-(2)]
	_ = x[UserDataTypeBio-(3)]
	_ = x[UserDataTypeURL-(5)]
	_ = x[UserDataTypeUsername-(6)]
}

var _UserDataTypeValues = []UserDataType{UserDataTypeNone, UserDataTypePfp, UserDataTypeDisplay, UserDataTypeBio, UserDataTypeURL, UserDataTypeUsername}

var _UserDataTypeNameToValueMap = map[string]UserDataType{
	_UserDataTypeName_0[0:19]:       UserDataTypeNone,
	_UserDataTypeLowerName_0[0:19]:  UserDataTypeNone,
	_UserDataTypeName_0[19:37]:      UserDataTypePfp,
	_UserDataTypeLowerName_0[19:37]: UserDataTypePfp,
	_UserDataTypeName_0[37:59]:      UserDataTypeDisplay,
	_UserDataTypeLowerName_0[37:59]: UserDataTypeDisplay,
	_UserDataTypeName_0[59:77]:      UserDataTypeBio,
	_UserDataTypeLowerName_0[59:77]: UserDataTypeBio,
	_UserDataTypeName_1[0:18]:       UserDataTypeURL,
	_UserDataTypeLowerName_1[0:18]:  UserDataTypeURL,
	_UserDataTypeName_1[18:41]:      UserDataTypeUsername,
	_UserDataTypeLowerName_1[18:41]: UserDataTypeUsername,
}

var _UserDataTypeNames = []string{
	_UserDataTypeName_0[0:19],
	_UserDataTypeName_0[19:37],
	_UserDataTypeName_0[37:59],
	_UserDataTypeName_0[59:77],
	_UserDataTypeName_1[0:18],
	_UserDataTypeName_1[18:41],
}

// UserDataTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UserDataTypeString(s string) (UserDataType, error) {
	if val, ok := _UserDataTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _UserDataTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UserDataType values", s)
}

// UserDataTypeValues returns all values of the enum
func UserDataTypeValues() []UserDataType {
	return _UserDataTypeValues
}

// UserDataTypeStrings returns a slice of all String values of the enum
func UserDataTypeStrings() []string {
	strs := make([]string, len(_UserDataTypeNames))
	copy(strs, _UserDataTypeNames)
	return strs
}

// IsAUserDataType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UserDataType) IsAUserDataType() bool {
	for _, v := range _UserDataTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
