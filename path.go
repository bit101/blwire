// Package wire has wireframe drawing funcs
package wire

// Path is a series of 3d points
type Path struct {
	Points []*Point
}

// NewPath creates a new 3d path
func NewPath() *Path {
	return &Path{[]*Point{}}
}

// FromPoints creates a new 3d path
func FromPoints(points []*Point) *Path {
	return &Path{points}
}

// AddPoint adds a point to a path
func (p *Path) AddPoint(point *Point) {
	p.Points = append(p.Points, point)
}

// RotateX rotates a 3d path on the x-axis
func (p *Path) RotateX(t float64) *Path {
	path := NewPath()
	for _, point := range p.Points {
		path.AddPoint(point.RotateX(t))
	}
	return path
}

// RotateY rotates a 3d path on the x-axis
func (p *Path) RotateY(t float64) *Path {
	path := NewPath()
	for _, point := range p.Points {
		path.AddPoint(point.RotateY(t))
	}
	return path
}

// RotateZ rotates a 3d path on the x-axis
func (p *Path) RotateZ(t float64) *Path {
	path := NewPath()
	for _, point := range p.Points {
		path.AddPoint(point.RotateZ(t))
	}
	return path
}

// Rotate rotates a 3d path on the x, y, and z axes
func (p *Path) Rotate(x, y, z float64) *Path {
	path := NewPath()
	for _, point := range p.Points {
		path.AddPoint(point.Rotate(x, y, z))
	}
	return path
}
