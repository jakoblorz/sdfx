//-----------------------------------------------------------------------------
/*

Voronoi Diagram and Delaunay Triangulation

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/sdf"
)

//-----------------------------------------------------------------------------

func main() {
	// create a random set of vertices
	b := sdf.NewBox2(sdf.V2{0, 0}, sdf.V2{20, 20})
	s := b.RandomSet(20)
	pixels := sdf.V2i{800, 800}
	k := 1.5
	path := "voronoi.png"

	// use a 0 radius circle as a point
	point, err := sdf.Circle2D(0.0)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// build an SDF for the points
	var s0 sdf.SDF2
	for i := range s {
		s0 = sdf.Union2D(s0, sdf.Transform2D(point, sdf.Translate2d(s[i])))
	}

	// work out the region we will sample
	bb := s0.BoundingBox().ScaleAboutCenter(k)
	log.Printf("rendering %s (%dx%d)\n", path, pixels[0], pixels[1])
	d, err := render.NewPNG(path, bb, pixels)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	d.RenderSDF2(s0)

	// create the delaunay triangulation
	ts, _ := render.Delaunay2d(s)
	// render the triangles
	for _, t := range ts {
		d.Triangle(t.ToTriangle2(s))
	}

	d.Save()
}

//-----------------------------------------------------------------------------
