// Code generated by "stringer -type=GovernmentType"; DO NOT EDIT.

package content

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GovernmentTypeDemocratic-1]
	_ = x[GovernmentTypeOligarchic-2]
	_ = x[GovernmentTypeDictatorial-3]
	_ = x[GovernmentTypeTheocratic-4]
	_ = x[GovernmentTypeAnarchic-5]
	_ = x[GovernmentTypeImperial-6]
	_ = x[GovernmentTypeHivemind-7]
	_ = x[GovernmentTypeMAX-8]
}

const _GovernmentType_name = "GovernmentTypeDemocraticGovernmentTypeOligarchicGovernmentTypeDictatorialGovernmentTypeTheocraticGovernmentTypeAnarchicGovernmentTypeImperialGovernmentTypeHivemindGovernmentTypeMAX"

var _GovernmentType_index = [...]uint8{0, 24, 48, 73, 97, 119, 141, 163, 180}

func (i GovernmentType) String() string {
	i -= 1
	if i < 0 || i >= GovernmentType(len(_GovernmentType_index)-1) {
		return "GovernmentType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _GovernmentType_name[_GovernmentType_index[i]:_GovernmentType_index[i+1]]
}
