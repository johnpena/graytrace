package tracer

import (
	"fmt"
	"math"
)

type Vec3 struct {
	x float64
	y float64
	z float64
}

func (v *Vec3) String() string {
	return fmt.Sprintf("<x: %f, y: %f, z: %f>", v.x, v.y, v.z)
}

func NewVec3(a, b, c float64) *Vec3 {
	return &Vec3{
		x: a,
		y: b,
		z: c,
	}
}

func (v *Vec3) X() float64 {
	return v.x
}

func (v *Vec3) Y() float64 {
	return v.y
}

func (v *Vec3) Z() float64 {
	return v.z
}

func (v *Vec3) R() float64 {
	return v.x
}

func (v *Vec3) G() float64 {
	return v.y
}

func (v *Vec3) B() float64 {
	return v.z
}

func (v *Vec3) Add(other *Vec3) *Vec3 {
	return &Vec3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v *Vec3) Sub(other *Vec3) *Vec3 {
	return &Vec3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v *Vec3) Mul(other *Vec3) *Vec3 {
	return &Vec3{
		x: v.x * other.x,
		y: v.y * other.y,
		z: v.z * other.z,
	}
}

func (v *Vec3) MulT(t float64) *Vec3 {
	return &Vec3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v *Vec3) Div(other *Vec3) Vec3 {
	return Vec3{
		x: v.x / other.x,
		y: v.y / other.y,
		z: v.z / other.z,
	}
}

func (v *Vec3) DivT(t float64) *Vec3 {
	return &Vec3{
		x: v.x / t,
		y: v.y / t,
		z: v.z / t,
	}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vec3) Normalize() *Vec3 {
	v2 := Dot(v, v)
	if v2 == 0 {
		return &Vec3{}
	}
	return v.MulT(1 / math.Sqrt(v2))
}

func (v *Vec3) Unit() *Vec3 {
	return v.Normalize()
}

func Dot(v1, v2 *Vec3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func Cross(v1, v2 *Vec3) Vec3 {
	return Vec3{
		x: v1.y*v2.z - v1.z*v2.y,
		y: v1.z*v2.x - v1.x*v2.z,
		z: v1.x*v2.y - v1.y*v2.x,
	}
}
