package model

import "testing"

var tests = []struct {
	input    tuple
	valx     float64
	valy     float64
	valz     float64
	valw     float64
	isVector bool
}{
	{tuple{4.3, -4.2, 3.1, 1.0}, 4.3, -4.2, 3.1, 1.0, false},
	{tuple{4.3, -4.2, 3.1, 0.0}, 4.3, -4.2, 3.1, 0.0, true},
}

func TestTuples(t *testing.T) {
	for _, e := range tests {
		if e.valx != e.input.x {
			t.Errorf("not matching %v values of x parameter %v", e.valx, e.input.x)
		}
		if e.valy != e.input.y {
			t.Errorf("not matching %v values of y parameter %v", e.valy, e.input.y)
		}
		if e.valz != e.input.z {
			t.Errorf("not matching %v values of z parameter %v", e.valz, e.input.z)
		}
		if e.valw != e.input.w {
			t.Errorf("not matching %v values of w parameter %v", e.valw, e.input.w)
		}
		if e.isVector != e.input.isVector() {
			t.Errorf("not matching %v with actual isVector() output %v", e.isVector, e.input.isVector())
		}
	}
}

func TestPoints(t *testing.T) {
	p := point(4, -4, 3)
	got := tuple{4, -4, 3, 1.0}

	if p != got {
		t.Errorf("want %v, got %v", p, got)
	}

}

func TestVectors(t *testing.T) {
	p := vector(4, -4, 3)
	got := tuple{4, -4, 3, 0.0}

	if p != got {
		t.Errorf("want %v, got %v", p, got)
	}

}

func TestIsEqualTo(t *testing.T) {
	p1 := point(4, -4, 3)
	p2 := point(4, -4, 3)

	got := p1.isEqualTo(p2)

	if !got {
		t.Errorf("expected true, got false")
	}

	v1 := vector(4, -4, 3)

	got = p1.isEqualTo(v1)

	if got {
		t.Errorf("expected false, got true")
	}
}

func TestAddTuple(t *testing.T) {
	p1 := point(4, -4, 3)
	v1 := vector(4, 4, -4)

	got, err := p1.Add(v1)
	if err != nil {
		t.Errorf("error received %v", err)
	}
	want := point(8, 0, -1)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSubtractTuple(t *testing.T) {
	p1 := point(3, 2, 1)
	p2 := point(5, 6, 7)

	got, err := p1.Subtract(p2)
	if err != nil {
		t.Errorf("error received %v", err)
	}
	want := vector(-2, -4, -6)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}

	v1 := vector(5, 6, 7)

	got, err = p1.Subtract(v1)
	if err != nil {
		t.Errorf("error received %v", err)
	}
	want = point(-2, -4, -6)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}

	v2 := vector(3, 2, 1)

	got, err = v2.Subtract(v1)

	if err != nil {
		t.Errorf("error received %v", err)
	}
	want = vector(-2, -4, -6)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}

}

func TestZeroVectorSubtractTuple(t *testing.T) {
	v1 := vector(0, 0, 0)
	v2 := vector(1, -2, 3)

	got, err := v1.Subtract(v2)
	if err != nil {
		t.Errorf("error received %v", err)
	}
	want := vector(-1, 2, -3)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNegate(t *testing.T) {
	v1 := vector(1, -2, 3)

	got, err := v1.Negate()
	if err != nil {
		t.Errorf("error received %v", err)
	}
	want := vector(-1, 2, -3)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
