package main

import (
	"flag"
	"fmt"
	"image"
	clr "image/color"
	"image/png"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	"github.com/jakoblorz/sdfx/render"
	"github.com/jakoblorz/sdfx/render/camera"
	"github.com/jakoblorz/sdfx/render/material"
	"github.com/jakoblorz/sdfx/sdf"
	"github.com/veandco/go-sdl2/sdl"
)

// RaysPerPixelList is used on the command line (flag) to define the number of rays per pixel per phase (hence a list)
// Example: ray-tracing -r 1 -r 10
type RaysPerPixelList []int

func (r *RaysPerPixelList) String() string {
	return fmt.Sprint(*r)
}

func (r *RaysPerPixelList) Set(value string) error {
	for _, e := range strings.Split(value, ",") {
		i, err := strconv.Atoi(e)
		if err != nil {
			return err
		}
		*r = append(*r, i)
	}
	return nil
}

// Options defines all the command line options available (all have a default value)
type Options struct {
	Width        int
	Height       int
	RaysPerPixel RaysPerPixelList
	Output       string
	Seed         int64
	CPU          int
}

// display will update the screen with the pixels provided
// note that there is no synchronization required on the array of pixels since it is an array of 32 bits integers
// that only gets updated to a final value by 1 goroutine at a time
func display(window *sdl.Window, screen *sdl.Surface, scene *render.Scene, pixels render.Pixels) {
	// create an image from the pixels generated
	image, err := sdl.CreateRGBSurfaceFrom(unsafe.Pointer(&pixels[0]), int32(scene.Width()), int32(scene.Height()), 32, scene.Width()*int(unsafe.Sizeof(pixels[0])), 0, 0, 0, 0)
	if err != nil {
		panic(err)
	}
	defer image.Free()
	// copy it into the screen
	err = image.Blit(nil, screen, nil)
	if err != nil {
		panic(err)
	}

	// update the surface to show it
	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

// buildWorld is the end result chapter 7
func buildWorldChapter7(width, height int) (render.Camera, render.HitableList) {

	world := render.HitableList{
		Sphere{center: sdf.V3{Z: -1.0}, radius: 0.5, material: material.Lambertian{render.Color{R: 1.0}}},
		Sphere{center: sdf.V3{Y: -100.5, Z: -1.0}, radius: 100, material: material.Lambertian{render.Color{G: 1.0}}},
	}

	lookFrom := sdf.V3{0, 0.0, 3.0}
	lookAt := sdf.V3{Z: -1.0}
	aperture := 0.0
	distToFocus := 1.0

	return camera.NewBasicFOVCamera(lookFrom, lookAt, sdf.V3{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus), world
}

// buildWorldMetalSpheres is the end result chapter 8
func buildWorldMetalSpheres(width, height int) (render.Camera, render.HitableList) {

	world := render.HitableList{
		Sphere{center: sdf.V3{Z: -1.0}, radius: 0.5, material: material.Lambertian{render.Color{R: 0.8, G: 0.3, B: 0.3}}},
		Sphere{center: sdf.V3{Y: -100.5, Z: -1.0}, radius: 100, material: material.Lambertian{render.Color{R: 0.8, G: 0.8}}},
		Sphere{center: sdf.V3{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: material.Metal{render.Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		Sphere{center: sdf.V3{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: material.Metal{render.Color{R: 0.8, G: 0.8, B: 0.8}, 0.3}},
	}

	lookFrom := sdf.V3{0, 0.0, 3.0}
	lookAt := sdf.V3{Z: -1.0}
	aperture := 0.0
	distToFocus := 1.0

	return camera.NewBasicFOVCamera(lookFrom, lookAt, sdf.V3{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus), world
}

// buildWorldDielectrics is the end result chapter 10
func buildWorldDielectrics(width, height int) (render.Camera, render.HitableList) {

	world := render.HitableList{
		Sphere{center: sdf.V3{Z: -1.0}, radius: 0.5, material: material.Lambertian{render.Color{R: 0.1, G: 0.2, B: 0.5}}},
		Sphere{center: sdf.V3{Y: -100.5, Z: -1.0}, radius: 100, material: material.Lambertian{render.Color{R: 0.8, G: 0.8}}},
		Sphere{center: sdf.V3{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: material.Metal{render.Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		Sphere{center: sdf.V3{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: material.Dielectric{1.5}},
		Sphere{center: sdf.V3{X: -1.0, Y: 0, Z: -1.0}, radius: -0.45, material: material.Dielectric{1.5}},
	}

	lookFrom := sdf.V3{-2.0, 2.0, 1.0}
	lookAt := sdf.V3{Z: -1.0}
	aperture := 0.0
	distToFocus := 1.0

	return camera.NewBasicFOVCamera(lookFrom, lookAt, sdf.V3{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus), world
}

func buildWorldMandelbulb(width, height int) (render.Camera, render.HitableList) {

	world := render.HitableList{
		// Sphere{center: sdf.V3{X: -1, Z: -1}, radius: 2, material: Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		// Sphere{center: sdf.V3{Y: -100.5, Z: -1.0}, radius: 100, material: Lambertian{Color{R: 0.8, G: 0.8}}},
		// Sphere{center: sdf.V3{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		// Sphere{center: sdf.V3{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Dielectric{1.5}},
		// Sphere{center: sdf.V3{X: -1.0, Y: 0, Z: -1.0}, radius: -0.45, material: Dielectric{1.5}},
		Mandelbulb{material.Metal{render.Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
	}

	lookFrom := sdf.V3{-4.0, 4.0, 3.0}
	lookAt := sdf.V3{}
	aperture := 0.0
	distToFocus := 1.0

	return camera.NewBasicFOVCamera(lookFrom, lookAt, sdf.V3{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus), world
}

// saveImage saves the image (if requested) to a file in png format
func saveImage(pixels render.Pixels, options Options) (error, bool) {
	if options.Output != "" {
		f, err := os.OpenFile(options.Output, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err, true
		}

		img := image.NewNRGBA(image.Rect(0, 0, options.Width, options.Height))

		k := 0
		for y := 0; y < options.Height; y++ {
			for x := 0; x < options.Width; x++ {
				p := pixels[k]
				img.Set(x, y, clr.NRGBA{
					R: uint8(p >> 16 & 0xFF),
					G: uint8(p >> 8 & 0xFF),
					B: uint8(p & 0xFF),
					A: 255,
				})
				k++
			}
		}

		if err := png.Encode(f, img); err != nil {
			f.Close()
			return err, true
		}

		if err := f.Close(); err != nil {
			return err, true
		}

		return nil, true
	}

	return nil, false

}

// main parses the options, set up the Window/Screen, builds the world and renders the scene.
// As the scene gets rendered the screen gets refreshed regularly to show progress. When the image is fully
// rendered, it saves it to a file (if the output option is set)
func main() {
	options := Options{}

	flag.IntVar(&options.Width, "w", 800, "width in pixel")
	flag.IntVar(&options.Height, "h", 400, "height in pixel")
	flag.IntVar(&options.CPU, "cpu", runtime.NumCPU(), "number of CPU to use (default to number of CPU available)")
	flag.Int64Var(&options.Seed, "seed", 2017, "seed for random number generator")
	flag.Var(&options.RaysPerPixel, "r", "comma separated list (or multiple) rays per pixel")
	flag.StringVar(&options.Output, "o", "", "path to file for saving (do not save if not defined)")

	flag.Parse()

	if len(options.RaysPerPixel) == 0 {
		options.RaysPerPixel = []int{1, 99}
	}

	// initializes the random number generator (since the scene has random spheres... to be reproducible)
	rand.Seed(options.Seed)

	// initializes SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// create (and show) window
	window, err := sdl.CreateWindow("Ray Tracing", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(options.Width), int32(options.Height), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// retrieves the screen
	screen, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// clear the screen (otherwise there is garbage...)
	err = screen.FillRect(&sdl.Rect{W: int32(options.Width), H: int32(options.Height)}, 0x00000000)
	if err != nil {
		panic(err)
	}

	//camera, world := buildWorldChapter7(options.Width, options.Height)
	//camera, world := buildWorldMetalSpheres(options.Width, options.Height)
	//camera, world := buildWorldDielectrics(options.Width, options.Height)
	camera, world := buildWorldMandelbulb(options.Width, options.Height)
	//camera, world := buildWorldOneWeekend(options.Width, options.Height)

	scene := render.NewScene(options.Width, options.Height, options.RaysPerPixel, camera, world)
	pixels, completed := scene.Render(options.CPU)

	// update the surface to show it
	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}

	updateDisplay := true

	// poll for quit event
	for running := true; running; {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		// wait a few between iterations
		sdl.Delay(16)

		if updateDisplay {

			display(window, screen, scene, pixels)

			// check (non blocking thanks to select) that the image is completely rendered
			select {
			case <-completed:
				updateDisplay = false
				fmt.Println("Render complete.")
				err, saved := saveImage(pixels, options)
				switch {
				case err != nil:
					fmt.Printf("Error while saving the image [%v]\n", err)
				case saved:
					fmt.Printf("Image saved to %v\n", options.Output)
				}
			default:
				break
			}

		}
	}
}
