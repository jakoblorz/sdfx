package camera

import (
	"math"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

type basicFOVCamera struct {
	origin          sdf.V3
	lowerLeftCorner sdf.V3
	horizontal      sdf.V3
	vertical        sdf.V3
	u, v            sdf.V3
	lensRadius      float64
}

// NewBasicFOVCamera computes the parameters necessary for the camera...
//	vfov is expressed in degrees (not radians)
func NewBasicFOVCamera(lookFrom sdf.V3, lookAt sdf.V3, vup sdf.V3, vfov float64, aspect float64, aperture float64, focusDist float64) render.Camera {
	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Sub(lookAt).Normalize()
	u := vup.Cross(w).Normalize()
	v := w.Cross(u).Normalize()

	lowerLeftCorner := origin.Add(u.MulScalar(-(halfWidth * focusDist))).Add(v.MulScalar(-(halfHeight * focusDist))).Add(w.MulScalar(-focusDist))
	horizontal := u.MulScalar(2 * halfWidth * focusDist)
	vertical := v.MulScalar(2 * halfHeight * focusDist)

	return basicFOVCamera{origin, lowerLeftCorner, horizontal, vertical, u, v, aperture / 2.0}
}

// ray implements the main api of the Camera interface according to the book
func (c basicFOVCamera) Ray(rnd render.Rnd, u, v float64) *render.Ray3 {
	d := c.lowerLeftCorner.Add(c.horizontal.MulScalar(u)).Add(c.vertical.MulScalar(v)).Sub(c.origin)
	origin := c.origin

	if c.lensRadius > 0 {
		rd := render.RandomInUnitDisk(rnd).MulScalar(c.lensRadius)
		offset := c.u.MulScalar(rd.X).Add(c.v.MulScalar(rd.Y))
		origin = origin.Add(offset)
		d = d.Sub(offset)
	}
	return &render.Ray3{origin, d, rnd}
}
