package main

import (
	"math"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

type Sphere struct {
	center   sdf.V3
	radius   float64
	material render.Material
}

// hit implements the hit interface for a Sphere
func (s Sphere) Hit(r *render.Ray3, tMin float64, tMax float64) (bool, *render.HitRecord) {
	oc := r.Origin.Sub(s.center)        // A-C
	a := r.Direction.Dot(r.Direction)   // dot(B, B)
	b := oc.Dot(r.Direction)            // dot(A-C, B)
	c := oc.Dot(oc) - s.radius*s.radius // dot(A-C, A-C) - R*R
	discriminant := b*b - a*c

	if discriminant > 0 {
		discriminantSquareRoot := math.Sqrt(discriminant)

		temp := (-b - discriminantSquareRoot) / a
		if temp < tMax && temp > tMin {
			hitPoint := r.PointAt(temp)
			return true, render.RecordHit(temp, hitPoint, hitPoint.Sub(s.center).MulScalar(1/s.radius), s.material)
		}

		temp = (-b + discriminantSquareRoot) / a
		if temp < tMax && temp > tMin {
			hitPoint := r.PointAt(temp)
			return true, render.RecordHit(temp, hitPoint, hitPoint.Sub(s.center).MulScalar(1/s.radius), s.material)
		}
	}

	return false, nil
}
