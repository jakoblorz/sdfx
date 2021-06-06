//-----------------------------------------------------------------------------
/*

Replacement Cap for Plastic Gas/Oil Can

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/obj"
	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

//-----------------------------------------------------------------------------

const capRadius = 56.0 / 2.0
const capHeight = 28.0
const capThickness = 4.0
const threadPitch = 6.0
const holeRadius = 0.0 // 33.0 / 2.0

//var threadDiameter = 48.0 // tight
const threadDiameter = 48.5 // just right
//var threadDiameter = 49.0 // loose
const threadRadius = threadDiameter / 2.0

//-----------------------------------------------------------------------------

func capOuter() (sdf.SDF3, error) {
	return obj.KnurledHead3D(capRadius, capHeight, capRadius*0.25)
}

func capInner() (sdf.SDF3, error) {
	tp, err := sdf.PlasticButtressThread(threadRadius, threadPitch)
	if err != nil {
		return nil, err
	}
	screw, err := sdf.Screw3D(tp, capHeight, threadPitch, 1)
	if err != nil {
		return nil, err
	}
	return sdf.Transform3D(screw, sdf.Translate3d(sdf.V3{0, 0, -capThickness})), nil
}

func capHole() (sdf.SDF3, error) {
	if holeRadius == 0 {
		// no hole
		return nil, nil
	}
	return sdf.Cylinder3D(capHeight, holeRadius, 0)
}

func gasCap() (sdf.SDF3, error) {
	// hole
	hole, err := capHole()
	if err != nil {
		return nil, err
	}
	// inner
	inner, err := capInner()
	if err != nil {
		return nil, err
	}
	inner = sdf.Union3D(inner, hole)
	// outer
	outer, err := capOuter()
	if err != nil {
		return nil, err
	}
	return sdf.Difference3D(outer, inner), nil
}

//-----------------------------------------------------------------------------

func main() {
	gasCap, err := gasCap()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTLSlow(gasCap, 200, "cap.stl")
}

//-----------------------------------------------------------------------------
