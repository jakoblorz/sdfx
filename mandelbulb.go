package main

import (
	"math"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

var (
	Bailout    = 2.0
	Power      = 10.0
	Iterations = 15
	Epsilon    = 0.01
)

type Mandelbulb struct {
	render.Material
}

func (m Mandelbulb) Evaluate(pos sdf.V3) float64 {
	var (
		z  = pos
		dr = 1.0
		r  = 0.0
	)

	for i := 0; i < Iterations; i++ {
		r = z.Length()
		if r > Bailout {
			break
		}

		// convert to polar coordinates
		var (
			theta = math.Acos(z.Z / r)
			phi   = math.Atan2(z.Y, z.X)
		)
		dr = math.Pow(r, Power-1.0)*Power*dr + 1.0

		// scale and rotate the point
		var (
			zr = math.Pow(r, Power)
		)
		theta = theta * Power
		phi = phi * Power

		// convert back to cartesian coordinates
		z = sdf.V3{
			X: (math.Sin(theta) * math.Cos(phi)) * zr,
			Y: (math.Sin(phi) * math.Sin(theta)) * zr,
			Z: (math.Cos(theta)) * zr,
		}.Add(pos)
	}

	return 0.5 * math.Log(r) * r / dr
}

func (m Mandelbulb) EstimateNormal(pos sdf.V3) sdf.V3 {
	return sdf.V3{
		X: m.Evaluate(pos.Add(sdf.V3{X: Epsilon})) - m.Evaluate(pos.Sub(sdf.V3{X: Epsilon})),
		Y: m.Evaluate(pos.Add(sdf.V3{Y: Epsilon})) - m.Evaluate(pos.Sub(sdf.V3{Y: Epsilon})),
		Z: m.Evaluate(pos.Add(sdf.V3{Z: Epsilon})) - m.Evaluate(pos.Sub(sdf.V3{Z: Epsilon})),
	}.Normalize()
}

func (m Mandelbulb) Hit(ray *render.Ray3, tMin float64, tMax float64) (bool, *render.HitRecord) {
	oc := ray.Origin                      // A-C
	a := ray.Direction.Dot(ray.Direction) // dot(B, B)
	b := oc.Dot(ray.Direction)            // dot(A-C, B)
	c := oc.Dot(oc) - 2*2                 // dot(A-C, A-C) - R*R
	discriminant := b*b - a*c

	if discriminant > 0 {
		discriminantSquareRoot := math.Sqrt(discriminant)

		temp := (-b - discriminantSquareRoot) / a
		if !(temp < tMax && temp > tMin) {
			temp = (-b + discriminantSquareRoot) / a
		}

		if temp < tMax && temp > tMin {
			return m.hit(&render.Ray3{
				Origin:    ray.PointAt(temp),
				Direction: ray.Direction,
			}, tMin, tMax, 0)
		}
	}

	return false, nil
}

func (m Mandelbulb) hit(ray *render.Ray3, tMin, tMax float64, depth int) (bool, *render.HitRecord) {
	if depth >= 1000 {
		return false, nil
	}

	t := m.Evaluate(ray.Origin)

	if t < tMax && t > tMin {
		p := ray.PointAt(t)

		if t < 0.01 {
			return true, render.RecordHit(t, ray.PointAt(t), m.EstimateNormal(ray.Origin), m)
		}

		return m.hit(&render.Ray3{
			Origin:    p,
			Direction: ray.Direction,
		}, tMin, tMax, depth+1)
	}

	return false, nil
}
