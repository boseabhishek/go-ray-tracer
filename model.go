package model

import "fmt"

/* // TupleType type
type TupleType float64

const (
	Vector TupleType = 0.0
	Point  TupleType = 1.0
) */

type tuple struct {
	x, y, z, w float64
}

func point(x, y, z float64) tuple {
	return tuple{
		x: x,
		y: y,
		z: z,
		w: 1.0,
	}
}

func vector(x, y, z float64) tuple {
	return tuple{
		x: x,
		y: y,
		z: z,
		w: 0.0,
	}
}

// EPSILON denotes the ignorable marginal difference between two floating point numbers
const EPSILON float64 = 0.00000001

func floatCompare(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func (t tuple) isEqualTo(t1 tuple) bool {
	if !floatCompare(t.x, t1.x) {
		return false
	}
	if !floatCompare(t.y, t1.y) {
		return false
	}
	if !floatCompare(t.z, t1.z) {
		return false
	}
	if !floatCompare(t.w, t1.w) {
		return false
	}
	return true
}

func (t tuple) isVector() bool {
	if t.w == 0.0 {
		return true
	}
	return false
}

func newTuple(t, t1 tuple, f func(float64, float64) float64) tuple {
	return tuple{
		x: f(t.x, t1.x),
		y: f(t.y, t1.y),
		z: f(t.z, t1.z),
		w: f(t.w, t1.w),
	}
}

func (t tuple) Add(t1 tuple) (tuple, error) {
	addFunc := func(a, b float64) float64 {
		return a + b
	}
	if !t.isVector() && !t1.isVector() {
		return tuple{}, fmt.Errorf("two points %+v and %+v can't be added", t, t1)
	}
	return newTuple(t, t1, addFunc), nil
}

func (t tuple) Subtract(t1 tuple) (tuple, error) {
	subFunc := func(a, b float64) float64 {
		return a - b
	}
	return newTuple(t, t1, subFunc), nil
}

func (t tuple) Negate() (tuple, error) {
	zt := tuple{0, 0, 0, 0}
	return zt.Subtract(t)
}
