// Package wire has wireframe drawing funcs
package wire

// MultiPath is shape made by multiple 3d paths.
type MultiPath struct {
	list []*Path
}

// NewMultiPath creates a new multi path
func NewMultiPath() *MultiPath {
	return &MultiPath{
		[]*Path{},
	}
}

// FromPaths returns a multi path from the given paths
func FromPaths(paths []*Path) *MultiPath {
	return &MultiPath{
		paths,
	}
}

// Box creates a 3d box
func Box(x, y, z, w, h, d float64) *MultiPath {
	mp := NewMultiPath()

	front := NewPath()
	front.AddPoint(NewPoint(x-w/2, y-h/2, z-d/2))
	front.AddPoint(NewPoint(x+w/2, y-h/2, z-d/2))
	front.AddPoint(NewPoint(x+w/2, y+h/2, z-d/2))
	front.AddPoint(NewPoint(x-w/2, y+h/2, z-d/2))
	front.AddPoint(NewPoint(x-w/2, y-h/2, z-d/2))
	mp.AddPath(front)

	back := NewPath()
	back.AddPoint(NewPoint(x-w/2, y-h/2, z+d/2))
	back.AddPoint(NewPoint(x+w/2, y-h/2, z+d/2))
	back.AddPoint(NewPoint(x+w/2, y+h/2, z+d/2))
	back.AddPoint(NewPoint(x-w/2, y+h/2, z+d/2))
	back.AddPoint(NewPoint(x-w/2, y-h/2, z+d/2))
	mp.AddPath(back)

	a := NewPath()
	a.AddPoint(NewPoint(x-w/2, y-h/2, z-d/2))
	a.AddPoint(NewPoint(x-w/2, y-h/2, z+d/2))
	mp.AddPath(a)

	b := NewPath()
	b.AddPoint(NewPoint(x+w/2, y-h/2, z-d/2))
	b.AddPoint(NewPoint(x+w/2, y-h/2, z+d/2))
	mp.AddPath(b)

	c := NewPath()
	c.AddPoint(NewPoint(x+w/2, y+h/2, z-d/2))
	c.AddPoint(NewPoint(x+w/2, y+h/2, z+d/2))
	mp.AddPath(c)

	e := NewPath()
	e.AddPoint(NewPoint(x-w/2, y+h/2, z-d/2))
	e.AddPoint(NewPoint(x-w/2, y+h/2, z+d/2))
	mp.AddPath(e)

	return mp
}

// AddPath adds a path to a multipath
func (m *MultiPath) AddPath(path *Path) {
	m.list = append(m.list, path)
}

// RotateX rotates a multipath on the x-axis
func (m *MultiPath) RotateX(t float64) *MultiPath {
	newMP := NewMultiPath()
	for _, p := range m.list {
		newMP.AddPath(p.RotateX(t))
	}
	return newMP
}

// RotateY rotates a multipath on the y-axis
func (m *MultiPath) RotateY(t float64) *MultiPath {
	newMP := NewMultiPath()
	for _, p := range m.list {
		newMP.AddPath(p.RotateY(t))
	}
	return newMP
}

// RotateZ rotates a multipath on the z-axis
func (m *MultiPath) RotateZ(t float64) *MultiPath {
	newMP := NewMultiPath()
	for _, p := range m.list {
		newMP.AddPath(p.RotateZ(t))
	}
	return newMP
}

// Rotate rotates a multipath on the x, y and z axes
func (m *MultiPath) Rotate(x, y, z float64) *MultiPath {
	newMP := NewMultiPath()
	for _, p := range m.list {
		newMP.AddPath(p.Rotate(x, y, z))
	}
	return newMP
}
