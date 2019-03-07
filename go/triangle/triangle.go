// Package triangle provides a function for determining what kind of triangle
// it's given.
package triangle

import "math"

const testVersion = 3

// Kind of triangle
type Kind int

// Kinds of triangles
const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides returns the kind of triangle, given the sizes of its sides.
func KindFromSides(a, b, c float64) Kind {
	if !positiveAndFinite(a) || !positiveAndFinite(b) || !positiveAndFinite(c) || !triangleInequalityHolds(a, b, c) {
		return NaT
	}

	switch {
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}

func positiveAndFinite(n float64) bool {
	return n > 0.0 && !math.IsInf(n, 1)
}

func triangleInequalityHolds(a, b, c float64) bool {
	return !(a+b < c || a+c < b || b+c < a)
}
