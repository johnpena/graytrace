package main

import (
	"graytrace/tracer"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func WorldBackgroundColor(ray *tracer.Ray) *tracer.Vec3 {
	unitDirection := ray.Direction().Unit()
	t := 0.5 * (unitDirection.Y() + 1.0)
	v1 := tracer.NewVec3(1.0, 1.0, 1.0).MulT(1.0 - t)
	v2 := tracer.NewVec3(0.5, 0.7, 1.0).MulT(t)

	return v1.Add(v2)
}

func Color(ray *tracer.Ray, world tracer.Hitable) *tracer.Vec3 {
	rec := world.Hit(ray, 0.0, math.MaxFloat64)
	if rec != nil {
		return tracer.NewVec3(rec.Normal().X()+1, rec.Normal().Y()+1, rec.Normal().Z()+1).MulT(0.5)
	}

	return WorldBackgroundColor(ray)
}

func main() {
	nx := 2000
	ny := 1000

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	lowerLeftCorner := tracer.NewVec3(-2.0, -1.0, -1.0)
	horizontal := tracer.NewVec3(4.0, 0.0, 0.0)
	vertical := tracer.NewVec3(0.0, 2.0, 0.0)
	origin := tracer.NewVec3(0.0, 0.0, 0.0)

	world := &tracer.HitableList{
		tracer.NewSphere(tracer.NewVec3(0.0, 0.0, -5.0), 0.9),
		tracer.NewSphere(tracer.NewVec3(0.0, -100.5, -100.0), 100),
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			ray := tracer.NewRay(origin, lowerLeftCorner.Add(horizontal.MulT(u)).Add(vertical.MulT(v)))

			// p := ray.PointAtParameter(2.0)
			col := Color(ray, world)

			img.Set(i, ny-j, color.RGBA{R: uint8(col.R() * 255), G: uint8(col.G() * 255), B: uint8(col.B() * 255), A: 255})
		}
	}

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	_ = png.Encode(f, img)
}
