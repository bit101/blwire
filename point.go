// Package wire has wireframe drawing funcs
package wire

import "math"

// Point is a 3d point
type Point struct {
	X, Y, Z float64
}

// NewPoint creates a new 3d point
func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

// RotateX rotates a point around the y-axis
func (p *Point) RotateX(t float64) *Point {
	c := math.Cos(t)
	s := math.Sin(t)
	y := c*p.Y - s*p.Z
	z := c*p.Z + s*p.Y
	return &Point{p.X, y, z}
}

// RotateY rotates a point around the y-axis
func (p *Point) RotateY(t float64) *Point {
	c := math.Cos(t)
	s := math.Sin(t)
	x := c*p.X - s*p.Z
	z := c*p.Z + s*p.X
	return &Point{x, p.Y, z}
}

// RotateZ rotates a point around the z-axis
func (p *Point) RotateZ(t float64) *Point {
	c := math.Cos(t)
	s := math.Sin(t)
	x := c*p.X - s*p.Y
	y := c*p.Y + s*p.X
	return &Point{x, y, p.Z}
}

// Rotate rotates a point around the x, y and z axes
func (p *Point) Rotate(x, y, z float64) *Point {
	np := p.RotateX(x)
	np = np.RotateY(y)
	return np.RotateZ(z)
}
