//-----------------------------------------------------------------------------
/*

KeyCaps for Cherry MX key switches

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

//-----------------------------------------------------------------------------

// material shrinkage
var shrink = 1.0 / 0.999 // PLA ~0.1%
//var shrink = 1.0/0.995; // ABS ~0.5%

//-----------------------------------------------------------------------------

const stemX = 6.0
const stemY = 5.0

const crossDepth = 4.0
const crossWidth = 1.0
const crossX = 4.0
const stemRound = 0.05

// keyStem returns a keycap stem of a given length.
func keyStem(length float64) (sdf.SDF3, error) {
	ofs := length - crossDepth
	s0, err := sdf.Box3D(sdf.V3{crossX, crossWidth, length}, crossX*stemRound)
	if err != nil {
		return nil, err
	}
	s1, err := sdf.Box3D(sdf.V3{crossWidth, stemY * (1.0 + 2.0*stemRound), length}, crossX*stemRound)
	if err != nil {
		return nil, err
	}
	cavity := sdf.Transform3D(sdf.Union3D(s0, s1), sdf.Translate3d(sdf.V3{0, 0, ofs}))
	stem, err := sdf.Box3D(sdf.V3{stemX, stemY, length}, stemX*stemRound)
	if err != nil {
		return nil, err
	}
	return sdf.Difference3D(stem, cavity), nil
}

//-----------------------------------------------------------------------------

const stemLength = 15.0

// roundCap returns a round keycap.
func roundCap(diameter, height, wall float64) (sdf.SDF3, error) {
	rOuter := 0.5 * diameter
	rInner := 0.5 * (diameter - (2.0 * wall))

	outer, err := sdf.Cylinder3D(height, rOuter, 0)
	if err != nil {
		return nil, err
	}

	inner, err := sdf.Cylinder3D(height, rInner, 0)
	if err != nil {
		return nil, err
	}

	inner = sdf.Transform3D(inner, sdf.Translate3d(sdf.V3{0, 0, wall}))
	keycap := sdf.Difference3D(outer, inner)

	stem, err := keyStem(stemLength)
	if err != nil {
		return nil, err
	}
	ofs := (stemLength - height) * 0.5
	stem = sdf.Transform3D(stem, sdf.Translate3d(sdf.V3{0, 0, ofs}))

	return sdf.Union3D(keycap, stem), nil
}

//-----------------------------------------------------------------------------

func main() {
	s, err := roundCap(18, 6, 1.5)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(sdf.ScaleUniform3D(s, shrink), 150, "round_cap.stl")
}

//-----------------------------------------------------------------------------
