// Package wire has wireframe drawing funcs
package wire

import cairo "github.com/bit101/blcairo"

// Wire is a Context add-on
type Wire struct {
	*cairo.Context
	Cx, Cy, Cz float64
	Persp      float64
}

// FromContext creates a new spatter
func FromContext(context *cairo.Context) *Wire {
	return &Wire{context, context.Width / 2, context.Height / 2, 1000, 500}
}

// Project projects a 3d point onto a 2d surface
func (w *Wire) Project(p *Point) (x, y float64) {
	scale := w.Persp / (w.Persp + p.Z + w.Cz)
	x = w.Cx + p.X*scale
	y = w.Cy + p.Y*scale
	return x, y
}

// Path draws a 3d path
func (w *Wire) Path(path *Path) {
	for _, p := range path.Points {
		x, y := w.Project(p)
		w.LineTo(x, y)
	}
	w.Stroke()
}

// MultiPath draws a 3d path
func (w *Wire) MultiPath(mpath *MultiPath) {
	for _, p := range mpath.list {
		w.Path(p)
	}
}
