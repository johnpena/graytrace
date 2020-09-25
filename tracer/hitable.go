package tracer

import (
	"fmt"
	"math"
)

type HitRecord struct {
	t float64
	p *Vec3
	normal *Vec3
}

func (hr *HitRecord) String() string {
	return fmt.Sprintf("HitRecord t: %f, p: %s, normal: %s", hr.t, hr.p, hr.normal)
}

func (hr *HitRecord) Normal() *Vec3 {
	return hr.normal
}

type Hitable interface {
	Hit(ray *Ray, tMin, tMax float64) *HitRecord
}

type HitableList []Hitable

func (hl HitableList) Hit(r *Ray, tMin, tMax float64) *HitRecord {
	var hr *HitRecord
	closestSoFar := tMax
	for i := 0; i < len(hl); i++  {
		hit := hl[i].Hit(r, tMin, closestSoFar)
		if hit != nil {
			closestSoFar = hit.t
			hr = hit
		}
	}

	return hr
}

type Sphere struct {
	center *Vec3
	radius float64
}

func NewSphere(center *Vec3, radius float64) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
	}
}

func (s *Sphere) Hit(r *Ray, tMin, tMax float64) *HitRecord {
	oc := r.Origin().Sub(s.center)
	a := Dot(r.Direction(), r.Direction())
	b := Dot(oc, r.Direction())
	c := Dot(oc, oc) - s.radius*s.radius

	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := -b - math.Sqrt(b*b-a*c)/a
		if temp < tMax && temp > tMin {
			point := r.PointAtParameter(temp)
			return &HitRecord{
				t:      temp,
				p:      point,
				normal: point.Sub(s.center).DivT(s.radius),
			}
		}

		temp = (-b + math.Sqrt(b*b - a*c))/a
		if temp < tMax && temp > tMin {
			point := r.PointAtParameter(temp)
			return &HitRecord{
				t:      temp,
				p:      point,
				normal: point.Sub(s.center).DivT(s.radius),
			}
		}
	}

	return nil
}
