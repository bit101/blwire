// Package main renders an image, gif or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

func main() {
	renderTarget := target.Video
	makePath()

	switch renderTarget {
	case target.Image:
		render.Image(800, 800, "out/out.png", renderFrame, 0.0)
		render.ViewImage("out/out.png")
		break

	case target.Video:
		render.Frames(400, 400, 240, "out/frames", renderFrame)
		render.ConvertToVideo("out/frames", "out/out.mp4", 400, 400, 30)
		render.PlayVideo("out/out.mp4")
		break
	}
}

var path *wire.Path
var box *wire.MultiPath

func makePath() {
	points := []*wire.Point{}

	for i := 0.0; i < blmath.Tau*200; i++ {

		p := wire.NewPoint(
			math.Cos(i/20)*400,
			-blmath.Tau*100+i,
			math.Sin(i/20)*400,
		)
		points = append(points, p)
	}
	path = wire.FromPoints(points)

	box = wire.Box(0, 0, 0, 1200, 1400, 1200)
}

//revive:disable-next-line:unused-parameter
func renderFrame(context *cairo.Context, width, height, percent float64) {
	random.Seed(0)
	context.WhiteOnBlack()
	for i := 0; i < 1000; i++ {
		p := geom.RandomPointInRect(0, 0, width, height)
		context.FillCircle(p.X, p.Y, random.FloatRange(0.1, 0.8))
	}
	w := wire.FromContext(context)
	w.Cz = 1300
	w.Persp = 200
	// w.SetLineWidth(0.7)

	w.SetLineWidth(1)
	rPath := path.Rotate(percent*blmath.Tau, percent*blmath.Tau, percent*blmath.Tau)
	w.Path(rPath)

	w.SetLineWidth(0.5)
	rBox := box.Rotate(percent*blmath.Tau, percent*blmath.Tau, percent*blmath.Tau)
	w.MultiPath(rBox)
}
