//-----------------------------------------------------------------------------
/*

Axoloti Board Mounting Kit

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

// material shrinkage
var shrink = 1.0 / 0.999 // PLA ~0.1%
//var shrink = 1.0/0.995; // ABS ~0.5%

//-----------------------------------------------------------------------------

const frontPanelThickness = 3.0
const frontPanelLength = 170.0
const frontPanelHeight = 50.0
const frontPanelYOffset = 15.0

const baseWidth = 50.0
const baseLength = 170.0
const baseThickness = 3.0

const baseFootWidth = 10.0
const baseFootCornerRadius = 3.0

const pcbWidth = 50.0
const pcbLength = 160.0

const pillarHeight = 16.8

//-----------------------------------------------------------------------------

// multiple standoffs
func standoffs() (sdf.SDF3, error) {

	k := &obj.StandoffParms{
		PillarHeight:   pillarHeight,
		PillarDiameter: 6.0,
		HoleDepth:      10.0,
		HoleDiameter:   2.4,
	}

	zOfs := 0.5 * (pillarHeight + baseThickness)

	// from the board mechanicals
	positions := sdf.V3Set{
		{3.5, 10.0, zOfs},   // H1
		{3.5, 40.0, zOfs},   // H2
		{54.0, 40.0, zOfs},  // H3
		{156.5, 10.0, zOfs}, // H4
		//{54.0, 10.0, zOfs},  // H5
		{156.5, 40.0, zOfs}, // H6
		{44.0, 10.0, zOfs},  // H7
		{116.0, 10.0, zOfs}, // H8
	}

	s, err := obj.Standoff3D(k)
	if err != nil {
		return nil, err
	}
	return sdf.Multi3D(s, positions), nil
}

//-----------------------------------------------------------------------------

// base returns the base mount.
func base() (sdf.SDF3, error) {
	// base
	pp := &obj.PanelParms{
		Size:         sdf.V2{baseLength, baseWidth},
		CornerRadius: 5.0,
		HoleDiameter: 3.5,
		HoleMargin:   [4]float64{7.0, 20.0, 7.0, 20.0},
		HolePattern:  [4]string{"xx", "x", "xx", "x"},
	}
	s0, err := obj.Panel2D(pp)
	if err != nil {
		return nil, err
	}

	// cutout
	l := baseLength - (2.0 * baseFootWidth)
	w := 18.0
	s1 := sdf.Box2D(sdf.V2{l, w}, baseFootCornerRadius)
	yOfs := 0.5 * (baseWidth - pcbWidth)
	s1 = sdf.Transform2D(s1, sdf.Translate2d(sdf.V2{0, yOfs}))

	s2 := sdf.Extrude3D(sdf.Difference2D(s0, s1), baseThickness)
	xOfs := 0.5 * pcbLength
	yOfs = pcbWidth - (0.5 * baseWidth)
	s2 = sdf.Transform3D(s2, sdf.Translate3d(sdf.V3{xOfs, yOfs, 0}))

	// standoffs
	s3, err := standoffs()
	if err != nil {
		return nil, err
	}

	s4 := sdf.Union3D(s2, s3)
	s4.(*sdf.UnionSDF3).SetMin(sdf.PolyMin(3.0))

	return s4, nil
}

//-----------------------------------------------------------------------------
// front panel cutouts

type panelHole struct {
	center sdf.V2   // center of hole
	hole   sdf.SDF2 // 2d hole
}

// button positions
const pbX = 53.0

var pb0 = sdf.V2{pbX, 0.8}
var pb1 = sdf.V2{pbX + 5.334, 0.8}

// panelCutouts returns the 2D front panel cutouts
func panelCutouts() (sdf.SDF2, error) {

	sMidi, err := sdf.Circle2D(0.5 * 17.0)
	if err != nil {
		return nil, err
	}
	sJack0, err := sdf.Circle2D(0.5 * 11.5)
	if err != nil {
		return nil, err
	}
	sJack1, err := sdf.Circle2D(0.5 * 5.5)
	if err != nil {
		return nil, err
	}

	sLed := sdf.Box2D(sdf.V2{1.6, 1.6}, 0)

	k := obj.FingerButtonParms{
		Width:  4.0,
		Gap:    0.6,
		Length: 20.0,
	}
	fb, err := obj.FingerButton2D(&k)
	if err != nil {
		return nil, err
	}
	sButton := sdf.Transform2D(fb, sdf.Rotate2d(sdf.DtoR(-90)))

	jackX := 123.0
	midiX := 18.8
	ledX := 62.9

	holes := []panelHole{
		{sdf.V2{midiX, 10.2}, sMidi},                         // MIDI DIN Jack
		{sdf.V2{midiX + 20.32, 10.2}, sMidi},                 // MIDI DIN Jack
		{sdf.V2{jackX, 8.14}, sJack0},                        // 1/4" Stereo Jack
		{sdf.V2{jackX + 19.5, 8.14}, sJack0},                 // 1/4" Stereo Jack
		{sdf.V2{107.6, 2.3}, sJack1},                         // 3.5 mm Headphone Jack
		{sdf.V2{ledX, 0.5}, sLed},                            // LED
		{sdf.V2{ledX + 3.635, 0.5}, sLed},                    // LED
		{pb0, sButton},                                       // Push Button
		{pb1, sButton},                                       // Push Button
		{sdf.V2{84.1, 1.0}, sdf.Box2D(sdf.V2{16.0, 7.5}, 0)}, // micro SD card
		{sdf.V2{96.7, 1.0}, sdf.Box2D(sdf.V2{11.0, 7.5}, 0)}, // micro USB connector
		{sdf.V2{73.1, 7.1}, sdf.Box2D(sdf.V2{7.5, 15.0}, 0)}, // fullsize USB connector
	}

	s := make([]sdf.SDF2, len(holes))
	for i, k := range holes {
		s[i] = sdf.Transform2D(k.hole, sdf.Translate2d(k.center))
	}

	return sdf.Union2D(s...), nil
}

//-----------------------------------------------------------------------------

// frontPanel returns the front panel mount.
func frontPanel() (sdf.SDF3, error) {

	// overall panel
	pp := &obj.PanelParms{
		Size:         sdf.V2{frontPanelLength, frontPanelHeight},
		CornerRadius: 5.0,
		HoleDiameter: 3.5,
		HoleMargin:   [4]float64{5.0, 5.0, 5.0, 5.0},
		HolePattern:  [4]string{"xx", "x", "xx", "x"},
	}
	panel, err := obj.Panel2D(pp)
	if err != nil {
		return nil, err
	}

	xOfs := 0.5 * pcbLength
	yOfs := (0.5 * frontPanelHeight) - frontPanelYOffset
	panel = sdf.Transform2D(panel, sdf.Translate2d(sdf.V2{xOfs, yOfs}))

	// extrude to 3d
	panelCutouts, err := panelCutouts()
	if err != nil {
		return nil, err
	}
	fp := sdf.Extrude3D(sdf.Difference2D(panel, panelCutouts), frontPanelThickness)

	// Add buttons to the finger button
	bHeight := 4.0
	b, _ := sdf.Cylinder3D(bHeight, 1.4, 0)
	b0 := sdf.Transform3D(b, sdf.Translate3d(pb0.ToV3(-0.5*bHeight)))
	b1 := sdf.Transform3D(b, sdf.Translate3d(pb1.ToV3(-0.5*bHeight)))

	return sdf.Union3D(fp, b0, b1), nil
}

//-----------------------------------------------------------------------------

func main() {

	// front panel
	s0, err := frontPanel()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	sx := sdf.Transform3D(s0, sdf.RotateY(sdf.DtoR(180.0)))
	render.RenderSTL(sdf.ScaleUniform3D(sx, shrink), 400, "panel.stl")

	// base
	s1, err := base()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(sdf.ScaleUniform3D(s1, shrink), 400, "base.stl")

	// both together
	s0 = sdf.Transform3D(s0, sdf.Translate3d(sdf.V3{0, 80, 0}))
	s3 := sdf.Union3D(s0, s1)
	render.RenderSTL(sdf.ScaleUniform3D(s3, shrink), 400, "panel_and_base.stl")
}

//-----------------------------------------------------------------------------
