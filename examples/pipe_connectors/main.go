//-----------------------------------------------------------------------------
/*

Pipe Connectors

*/
//-----------------------------------------------------------------------------

package main

import (
	"log"

	"github.com/jakoblorz/sdfx/obj"
	"github.com/jakoblorz/sdfx/render"
)

//-----------------------------------------------------------------------------

const name = "sch40:1"
const units = "mm"
const length = 40.0

//-----------------------------------------------------------------------------

func main() {

	// 2-way
	s, err := obj.StdPipeConnector3D(name, units, length, [6]bool{false, false, false, false, true, true})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_2a.stl")

	// 2-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, false, false, false, true, false})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_2b.stl")

	// 3-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, false, false, false, true, true})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_3a.stl")

	// 3-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, false, true, false, true, false})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_3b.stl")

	// 4-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, true, true, true, false, false})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_4a.stl")

	// 4-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, false, true, true, true, false})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_4b.stl")

	// 5-way
	s, err = obj.StdPipeConnector3D(name, units, length, [6]bool{true, true, true, true, true, false})
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	render.RenderSTL(s, 300, "pipe_connector_5a.stl")

}

//-----------------------------------------------------------------------------
