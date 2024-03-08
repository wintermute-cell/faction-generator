package util

import (
	"math"
)

type number interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}

type signed_number interface {
	int | int16 | int32 | int64 | float32 | float64
}

// Sign returns the sign of x (-1 if x < 0, 1 if x > 0, 0 if x == 0)
func Sign[T signed_number](x T) T {
	if x < 0 {
		return T(-1)
	} else if x > 0 {
		return T(1)
	} else {
		return T(0)
	}
}

// Abs returns the absolute value of x
func Abs[T number](x T) T {
	ret := x
	if x < 0 {
		ret = -x
	}
	return ret
}

// Max will return the maximum value between x and y
func Max[T number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

// Min will return the minimum value between x and y.
func Min[T number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

// Clamps x between lower_bound and upper_bound, both inclusive.
// (Clamp will return at least lower_bound and at most upper_bound)
func Clamp[T number](x, lower_bound, upper_bound T) T {
	v := Min[T](x, upper_bound)
	v = Max[T](v, lower_bound)
	return v
}

// Round x to the nearest integer, either down if x < .5 or up if x >= .5.
func Round[T number](x T) T {
	integer, fraction := math.Modf(float64(x))
	v := x
	if fraction >= 0.5 {
		v = T(integer) + T(1.0)
	} else {
		v = T(integer)
	}
	return v
}

// Vector2Lerp returns the linear interpolation between two numbers.
func Lerp[T number](a, b T, factor float32) T {
	return T(float32(a)*(1.0-factor) + (float32(b) * factor))
}

// ShortestLerp returns the shortest linear interpolation between two numbers.
func ShortestLerp(current, target, factor float32) float32 {
	// Calculate the difference
	difference := target - current

	// Calculate possible wrapped differences
	wrappedDifferencePlus := float64(difference + 360)
	wrappedDifferenceMinus := float64(difference - 360)

	// Check which one is the smallest in terms of absolute value
	if math.Abs(wrappedDifferencePlus) < math.Abs(float64(difference)) {
		difference = float32(wrappedDifferencePlus)
	} else if math.Abs(wrappedDifferenceMinus) < math.Abs(float64(difference)) {
		difference = float32(wrappedDifferenceMinus)
	}

	// Compute the lerped value
	lerped := current + difference*factor

	// Adjust the lerped value to be within the 0-360 range
	for lerped < 0 {
		lerped += 360
	}
	for lerped >= 360 {
		lerped -= 360
	}

	return lerped
}
