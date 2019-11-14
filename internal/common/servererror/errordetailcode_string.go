// Code generated by "stringer -type ErrorDetailCode internal/common/servererror/error_type.go"; DO NOT EDIT.

package servererror

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorDetailCodeUnknown-1]
	_ = x[ErrorDetailCodeTokenRequired-2]
	_ = x[ErrorDetailCodeTokenExpired-3]
	_ = x[ErrorDetailCodeTokenInvalid-4]
}

const _ErrorDetailCode_name = "ErrorDetailCodeUnknownErrorDetailCodeTokenRequiredErrorDetailCodeTokenExpiredErrorDetailCodeTokenInvalid"

var _ErrorDetailCode_index = [...]uint8{0, 22, 50, 77, 104}

func (i ErrorDetailCode) String() string {
	i -= 1
	if i < 0 || i >= ErrorDetailCode(len(_ErrorDetailCode_index)-1) {
		return "ErrorDetailCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ErrorDetailCode_name[_ErrorDetailCode_index[i]:_ErrorDetailCode_index[i+1]]
}
