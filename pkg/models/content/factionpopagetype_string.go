// Code generated by "stringer -type=FactionPopAgeType"; DO NOT EDIT.

package content

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FactionPopAgeTypeYoungling-1]
	_ = x[FactionPopAgeTypeAdult-2]
	_ = x[FactionPopAgeTypeElder-3]
	_ = x[FactionPopAgeTypeMAX-4]
}

const _FactionPopAgeType_name = "FactionPopAgeTypeYounglingFactionPopAgeTypeAdultFactionPopAgeTypeElderFactionPopAgeTypeMAX"

var _FactionPopAgeType_index = [...]uint8{0, 26, 48, 70, 90}

func (i FactionPopAgeType) String() string {
	i -= 1
	if i < 0 || i >= FactionPopAgeType(len(_FactionPopAgeType_index)-1) {
		return "FactionPopAgeType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _FactionPopAgeType_name[_FactionPopAgeType_index[i]:_FactionPopAgeType_index[i+1]]
}
