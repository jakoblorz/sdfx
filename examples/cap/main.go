//-----------------------------------------------------------------------------
/*

Tube Cap

This is a simple round cap that fits onto the outside of a tube.

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

//-----------------------------------------------------------------------------

const wallThickness = 2.0
const innerDiameter = 75.5
const innerHeight = 15.0

//-----------------------------------------------------------------------------

// material shrinkage
const shrink = 1.0 / 0.999 // PLA ~0.1%
//const shrink = 1.0/0.995; // ABS ~0.5%

//-----------------------------------------------------------------------------

func tubeCap() (sdf.SDF3, error) {

	h := innerHeight + wallThickness
	r := (innerDiameter * 0.5) + wallThickness
	outer, err := sdf.Cylinder3D(h, r, 1.0)
	if err != nil {
		return nil, err
	}

	h = innerHeight
	r = innerDiameter * 0.5
	inner, err := sdf.Cylinder3D(h, r, 1.0)
	inner = sdf.Transform3D(inner, sdf.Translate3d(sdf.V3{0, 0, wallThickness * 0.5}))
	if err != nil {
		return nil, err
	}

	return sdf.Difference3D(outer, inner), nil
}

//-----------------------------------------------------------------------------

func main() {
	c, err := tubeCap()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(sdf.ScaleUniform3D(c, shrink), 120, "cap.stl")
}

//-----------------------------------------------------------------------------
