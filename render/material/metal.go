package material

import "github.com/jakoblorz/sdfx/render"

/***********************
 * Metal material
 ************************/
type Metal struct {
	Albedo render.Color
	Fuzz   float64
}

func (mat Metal) Scatter(r *render.Ray3, rec *render.HitRecord) (bool, *render.Color, *render.Ray3) {
	reflected := r.Direction.Normalize().Reflect(rec.Normal())
	if mat.Fuzz < 1 {
		reflected = reflected.Add(render.RandomInUnitSphere(r.Rnd).MulScalar(mat.Fuzz))
	}
	scattered := &render.Ray3{Origin: rec.P(), Direction: reflected, Rnd: r.Rnd}
	attenuation := &mat.Albedo

	if scattered.Direction.Dot(rec.Normal()) > 0 {
		return true, attenuation, scattered
	}

	return false, nil, nil
}
