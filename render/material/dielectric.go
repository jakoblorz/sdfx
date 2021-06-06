package material

import (
	"math"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

/***********************
 * Dielectric material (glass)
 ************************/
type Dielectric struct {
	RefIdx float64
}

func schlick(cosine float64, iRefIdx float64) float64 {
	r0 := (1.0 - iRefIdx) / (1.0 + iRefIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5)
}

func (die Dielectric) Scatter(r *render.Ray3, rec *render.HitRecord) (bool, *render.Color, *render.Ray3) {
	var (
		outwardNormal sdf.V3
		niOverNt      float64
		cosine        float64
	)

	dotRayNormal := r.Direction.Dot(rec.Normal())
	if dotRayNormal > 0 {
		outwardNormal = rec.Normal().MulScalar(-1)
		niOverNt = die.RefIdx
		cosine = dotRayNormal / r.Direction.Length()
		cosine = math.Sqrt(1.0 - die.RefIdx*die.RefIdx*(1.0-cosine*cosine))
	} else {
		outwardNormal = rec.Normal()
		niOverNt = 1.0 / die.RefIdx
		cosine = -dotRayNormal / r.Direction.Length()
	}

	wasRefracted, refracted := r.Direction.Refract(outwardNormal, niOverNt)

	var direction sdf.V3

	// refract only with some probability
	if wasRefracted && r.Rnd.Float64() >= schlick(cosine, die.RefIdx) {
		direction = *refracted
	} else {
		direction = r.Direction.Normalize().Reflect(rec.Normal())
	}

	return true, &render.White, &render.Ray3{Origin: rec.P(), Direction: direction, Rnd: r.Rnd}
}
