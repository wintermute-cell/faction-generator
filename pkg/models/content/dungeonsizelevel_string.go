// Code generated by "stringer -type=DungeonSizeLevel"; DO NOT EDIT.

package content

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DungeonSizeLevelSmall-1]
	_ = x[DungeonSizeLevelMedium-2]
	_ = x[DungeonSizeLevelLarge-3]
	_ = x[DungeonSizeLevelGreater-4]
	_ = x[DungeonSizeLevelMega-5]
	_ = x[DungeonSizeLevelMAX-6]
}

const _DungeonSizeLevel_name = "DungeonSizeLevelSmallDungeonSizeLevelMediumDungeonSizeLevelLargeDungeonSizeLevelGreaterDungeonSizeLevelMegaDungeonSizeLevelMAX"

var _DungeonSizeLevel_index = [...]uint8{0, 21, 43, 64, 87, 107, 126}

func (i DungeonSizeLevel) String() string {
	i -= 1
	if i < 0 || i >= DungeonSizeLevel(len(_DungeonSizeLevel_index)-1) {
		return "DungeonSizeLevel(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DungeonSizeLevel_name[_DungeonSizeLevel_index[i]:_DungeonSizeLevel_index[i+1]]
}