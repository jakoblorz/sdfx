package render

import "github.com/jakoblorz/sdfx/sdf"

type Ray3 struct {
	Origin    sdf.V3
	Direction sdf.V3
	Rnd       Rnd
}

func (r *Ray3) PointAt(s float64) sdf.V3 {
	return r.Origin.Add(r.Direction.MulScalar(s))
}
