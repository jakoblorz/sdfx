package material

import "github.com/jakoblorz/sdfx/render"

/***********************
 * Lambertian material (diffuse only)
 ************************/
type Lambertian struct {
	Albedo render.Color
}

func (mat Lambertian) Scatter(r *render.Ray3, rec *render.HitRecord) (bool, *render.Color, *render.Ray3) {
	target := rec.P().Add(rec.Normal()).Add(render.RandomInUnitSphere(r.Rnd))
	scattered := &render.Ray3{Origin: rec.P(), Direction: target.Sub(rec.P()), Rnd: r.Rnd}
	attenuation := &mat.Albedo
	return true, attenuation, scattered
}
