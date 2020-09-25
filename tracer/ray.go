package tracer

type Ray struct {
	origin    *Vec3
	direction *Vec3
}

func NewRay(origin, direction *Vec3) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r *Ray) Origin() *Vec3 {
	return r.origin
}

func (r *Ray) Direction() *Vec3 {
	return r.direction
}

func (r *Ray) PointAtParameter(t float64) *Vec3 {
	return r.origin.Add(r.direction.MulT(t))
}
